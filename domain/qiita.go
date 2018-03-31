package domain

import "time"

type Qiita struct {
	ID string `json:"id"`
	TiTle string `json:"title"`
	Description string `json:"body"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
