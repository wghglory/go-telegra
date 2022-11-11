package model

type Account struct {
	ShortName   string `json:"short_name,omitempty"`
	AuthorName  string `json:"author_name,omitempty"`
	AuthorUrl   string `json:"author_url,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	AuthUrl     string `json:"auth_url,omitempty"`
	PageCount   int    `json:"page_count,omitempty"`
}
