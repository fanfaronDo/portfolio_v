package domain

import "time"

type Project struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
