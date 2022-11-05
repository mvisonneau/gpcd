package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func List(ctx *cli.Context) error {
	c := NewClientFromCLIContext(ctx)

	medias, err := c.ListMedias(ctx.Timestamp("from"), ctx.Timestamp("to"))
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, m := range medias {
		fmt.Println(m)
	}

	return nil
}

func (c *Client) ListMedias(from, to *time.Time) (medias []*Media, err error) {
	currentPage := int64(1)

	type listMediasResponse struct {
		DefaultResponse
		Embedded struct {
			Media []*Media `json:"media"`
		} `json:"_embedded"`
	}

	for {
		var (
			req            *http.Request
			resp           *http.Response
			parsedResponse listMediasResponse
		)

		req, err = http.NewRequest("GET", c.APIEndpoint+fmt.Sprintf("search?fields=captured_at,filename,type,resolution&processing_states=ready&order_by=captured_at&per_page=200&page=%d&type=Photo,Video,TimeLapse,TimeLapseVideo,Burst,Livestream,LoopedVideo,BurstVideo,Continuous,ExternalVideo,Session", currentPage), nil)
		if err != nil {
			err = errors.Wrap(err, "creating request")
			return
		}

		c.addAuthorizationHeaders(req)

		resp, err = c.Do(req)
		if err != nil {
			err = errors.Wrap(err, "processing response")
			return
		}

		if err = json.NewDecoder(resp.Body).Decode(&parsedResponse); err != nil {
			err = errors.Wrap(err, "unmarshalling response")
			resp.Body.Close()
			return
		}

		resp.Body.Close()

		for _, m := range parsedResponse.Embedded.Media {
			if from != nil && from.After(m.CapturedAt) {
				continue
			}

			if to != nil && to.Before(m.CapturedAt) {
				continue
			}

			medias = append(medias, m)
		}

		if currentPage == parsedResponse.Pages.TotalPages {
			break
		}
		currentPage++
	}

	return
}
