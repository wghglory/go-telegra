package model

type Account struct {
	ShortName   string `json:"short_name,omitempty" xml:"short_name,omitempty" form:"short_name,omitempty" query:"short_name,omitempty"`
	AuthorName  string `json:"author_name,omitempty" xml:"author_name,omitempty" form:"author_name,omitempty" query:"author_name,omitempty"`
	AuthorUrl   string `json:"author_url,omitempty" xml:"author_url,omitempty" form:"author_url,omitempty" query:"author_url,omitempty"`
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty" form:"access_token,omitempty" query:"access_token,omitempty"`
	AuthUrl     string `json:"auth_url,omitempty" xml:"auth_url,omitempty" form:"auth_url,omitempty" query:"auth_url,omitempty"`
	PageCount   int    `json:"page_count,omitempty" xml:"page_count,omitempty" form:"page_count,omitempty" query:"page_count,omitempty"`
}

type AccountPayload struct {
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty" form:"access_token,omitempty" query:"access_token,omitempty"`
	Fields      string `json:"fields,omitempty" xml:"fields,omitempty" form:"fields,omitempty" query:"fields,omitempty"`
}
