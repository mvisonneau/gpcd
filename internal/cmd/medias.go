package cmd

import (
	"fmt"
	"time"
)

type Media struct {
	ID         string    `json:"id"`
	CapturedAt time.Time `json:"captured_at"`
	Filename   string    `json:"filename"`
	Type       MediaType `json:"type"`
	Resolution string    `json:"resolution"`
}

type MediaType string

const (
	MediaTypeVideo MediaType = "Video"
	MediaTypePhoto MediaType = "Photo"

	// TODO: Add missed ones.
)

func (m *Media) String() string {
	return fmt.Sprintf("%s | %s | %s - %s (%s)", m.ID, m.CapturedAt.String(), m.Filename, m.Type, m.Resolution)
}
