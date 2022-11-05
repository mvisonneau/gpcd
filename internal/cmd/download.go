package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	downloader "github.com/mostafa-asg/go-dl"
)

func Download(ctx *cli.Context) error {
	c := NewClientFromCLIContext(ctx)

	medias, err := c.ListMedias(ctx.Timestamp("from"), ctx.Timestamp("to"))
	if err != nil {
		return err
	}

	for _, m := range medias {
		fmt.Println(m)
		if err = c.DownloadMedia(m, ctx.String("local-path")); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) DownloadMedia(m *Media, localPath string) (err error) {
	type listMediaURLsResponse struct {
		Embedded struct {
			Variations []struct {
				GetURL  string `json:"url"`
				HeadURL string `json:"head"`
				Type    string `json:"type"`
				Quality string `json:"quality"`
			} `json:"variations"`
		} `json:"_embedded"`
	}

	var (
		req            *http.Request
		resp           *http.Response
		parsedResponse listMediaURLsResponse
	)

	req, err = http.NewRequest("GET", fmt.Sprintf("%s%s/download", c.APIEndpoint, m.ID), nil)
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

	defer resp.Body.Close()

	var (
		getURL  string
		headURL string
	)

Variations:
	for _, v := range parsedResponse.Embedded.Variations {
		switch m.Type {
		case MediaTypePhoto:
			getURL = v.GetURL
			headURL = v.HeadURL
			break Variations
		case MediaTypeVideo:
			if v.Quality == m.Resolution {
				getURL = v.GetURL
				headURL = v.HeadURL
				break Variations
			}
		}
	}

	if len(getURL) == 0 {
		err = errors.New("unable to find a downloable URL for this media")
		return
	}

	if err = os.MkdirAll(localPath, os.ModePerm); err != nil {
		return
	}

	var d *downloader.Downloader
	d, err = downloader.NewFromConfig(&downloader.Config{
		URL:            getURL,
		HeadURL:        headURL,
		Concurrency:    c.Concurrency,
		OutFilename:    localPath + "/" + m.Filename,
		CopyBufferSize: 32 * 1024, // 32k
		Resume:         true,
	})
	if err != nil {
		return err
	}

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		d.Pause()
		os.Exit(1)
	}()

	d.Download()

	return
}
