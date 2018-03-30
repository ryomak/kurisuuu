package domain

type GithubApi struct {
	Name string `json:"name"`
	Url string `json:"html_url"`
	Language string `json:"language"`
	UpdatedAt string `json:"updated_at"`
}

type Github struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Language string `json:"language"`
	UpdatedAt string `json:"updated_at"`
}

