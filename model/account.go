package model

type Account struct {
	ShortName   string `json:"short_name"`
	AuthorName  string `json:"author_name"`
	AuthorUrl   string `json:"author_url"`
	AccessToken string `json:"access_token,omitempty"`
	AuthUrl     string `json:"auth_url,omitempty"`
	PageCount   int    `json:"page_count,omitempty"`
}
