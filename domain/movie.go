package domain

type Movie struct{
	Name string `json:"name"`
	Path string `json:"movie_path"`
	ImgPath string `json:"image_path"`
	Description string `json:"description"`
}
