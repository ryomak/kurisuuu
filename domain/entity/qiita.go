package entity

import "time"

type Qiita struct {
	ID string `json:id`
	TiTle string `json:title`
	Description string
	Url string
	CreatedAt time.Time
}
