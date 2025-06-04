package crawler

import (
	"strings"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type ContentType string

var Unknown ContentType = ""
var PDF ContentType = "pdf"
var Text ContentType = "text"

func CheckContentType(url string) (ContentType, error) {
	client := cycletls.Init()

	resp, err := client.Do(url, cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		Timeout:   5,
	}, "HEAD")
	if err != nil {
		return Unknown, err
	}

	header := resp.Headers["Content-Type"]

	if strings.Contains(header, "application/pdf") {
		return PDF, nil
	} else if strings.Contains(header, "text/html") || strings.Contains(header, "text/plain") {
		return Text, nil
	} else {
		return Unknown, nil
	}
}
