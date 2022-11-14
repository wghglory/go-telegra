package model

type Page struct {
	Path        string        `json:"path"`
	Url         string        `json:"url"`
	Title       string        `json:"title" xml:"title" form:"title" query:"title"`
	Description string        `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty" query:"description,omitempty"`
	AuthorName  string        `json:"author_name,omitempty" xml:"author_name,omitempty" form:"author_name,omitempty" query:"author_name,omitempty"`
	AuthorUrl   string        `json:"author_url,omitempty" xml:"author_url,omitempty" form:"author_url,omitempty" query:"author_url,omitempty"`
	ImageUrl    string        `json:"image_url,omitempty" xml:"image_url,omitempty" form:"image_url,omitempty" query:"image_url,omitempty"`
	Content     []NodeElement `json:"content,omitempty" xml:"content,omitempty" form:"content,omitempty" query:"content,omitempty"`
	Views       int           `json:"views"`
	CanEdit     bool          `json:"can_edit,omitempty"`
}

type PagePayload struct {
	AccessToken   string `json:"access_token" xml:"access_token" form:"access_token" query:"access_token"`
	Title         string `json:"title" xml:"title" form:"title" query:"title"`
	Description   string `json:"description" xml:"description" form:"description" query:"description"`
	AuthorName    string `json:"author_name" xml:"author_name" form:"author_name" query:"author_name"`
	AuthorUrl     string `json:"author_url" xml:"author_url" form:"author_url" query:"author_url"`
	ImageUrl      string `json:"image_url" xml:"image_url" form:"image_url" query:"image_url"`
	Content       string `json:"content" xml:"content" form:"content" query:"content"`
	ReturnContent string `json:"return_content" xml:"return_content" form:"return_content" query:"return_content"`
}
