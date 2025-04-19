package format

import (
	md "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func HTMLtoMarkdown(html *string) (string, error) {
	markdown, err := md.ConvertString(*html)
	if err != nil {
		return "", err
	}

	return markdown, nil
}
