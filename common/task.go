package common

type Task struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Result Page   `json:"result,omitempty"`
	Status string `json:"status"`
	Time   string `json:"time"`
}
