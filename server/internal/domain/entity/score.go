package entity

import "time"

// Score represents a drum score entity
type Score struct {
	ID        string
	Title     string
	Artist    string
	Tempo     int
	TimeSign  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewScore creates a new Score entity
func NewScore(title, artist string, tempo int, timeSign string) *Score {
	return &Score{
		Title:     title,
		Artist:    artist,
		Tempo:     tempo,
		TimeSign:  timeSign,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
