package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // Add the omitempty directive
	Runtime   Runtime   `json:"runtime,omitempty"` // Customizing the Runtime field
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
