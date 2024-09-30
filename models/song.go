package models

type Song struct {
	ID          int    `json:"id"`
	Group       string `json:"group_name"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Lyrics      string `json:"lyrics"`
	Link        string `json:"link"`
}
