package model

type Page struct {
	Path        string        `json:"path"`
	Url         string        `json:"url"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	AuthorName  string        `json:"author_name,omitempty"`
	ImageUrl    string        `json:"image_url,omitempty"`
	Content     []NodeElement `json:"content,omitempty"`
	Views       int           `json:"views"`
	CanEdit     bool          `json:"can_edit,omitempty"`
}
