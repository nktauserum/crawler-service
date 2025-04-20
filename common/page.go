package common

type Page struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Sitename string `json:"sitename"`
	Content  string `json:"content"`
	HTML     string `json:"html"`
}
