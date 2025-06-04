package common

type Page struct {
	URL      string `json:"url,omitempty"`
	Title    string `json:"title,omitempty"`
	Sitename string `json:"sitename,omitempty"`
	Content  string `json:"content,omitempty"`
	HTML     string `json:"html,omitempty"`
}

type Task struct {
	UUID   string `json:"uuid"`
	URL    string `json:"url"`
	Status string `json:"status"`
	Result string `json:"result"`
}
