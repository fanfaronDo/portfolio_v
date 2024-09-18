package domain

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Image       []byte `json:"image"`
}
