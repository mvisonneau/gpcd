package cmd

import (
	"net/http"

	"github.com/urfave/cli/v2"
)

type DefaultResponse struct {
	Pages struct {
		CurrentPage int64 `json:"current_page"`
		PerPage     int64 `json:"per_page"`
		TotalItems  int64 `json:"total_items"`
		TotalPages  int64 `json:"total_pages"`
	} `json:"_pages"`
}

type Client struct {
	http.Client
	APIEndpoint    string
	RequestHeaders map[string]string
	Concurrency    int
}

func NewClient(apiEndpoint, bearerToken, userAgent string, concurrency int) *Client {
	return &Client{
		Client:      *http.DefaultClient,
		APIEndpoint: apiEndpoint,
		RequestHeaders: map[string]string{
			"Accept":        "application/vnd.gopro.jk.media+json; version=2.0.0",
			"Authorization": "Bearer " + bearerToken,
			"User-Agent":    userAgent,
		},
		Concurrency: concurrency,
	}
}

func NewClientFromCLIContext(ctx *cli.Context) *Client {
	return NewClient(
		ctx.String("api-endpoint"),
		ctx.String("bearer-token"),
		ctx.String("user-agent"),
		ctx.Int("max-concurrent-downloads"),
	)
}

func (c *Client) addAuthorizationHeaders(req *http.Request) {
	for k, v := range c.RequestHeaders {
		req.Header.Add(k, v)
	}
}
