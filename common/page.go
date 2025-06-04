package common

type Page struct {
	URL      string `json:"url,omitempty"`
	Title    string `json:"title,omitempty"`
	Sitename string `json:"sitename,omitempty"`
	Content  string `json:"content,omitempty"`
	HTML     string `json:"html,omitempty"`
}
