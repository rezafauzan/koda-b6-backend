package dto

type Response struct {
	Success  bool   `json:"success"`
	Messages string `json:"messages"`
	Results  any    `json:"results"`
}
