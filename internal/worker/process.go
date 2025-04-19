package worker

import (
	"context"

	"github.com/nktauserum/crawler-service/pkg/crawler"
)

func ProcessURL(url string) (crawler.Page, error) {
	page, err := crawler.GetContent(context.Background(), url)
	if err != nil {
		return crawler.Page{}, err
	}

	return page, nil
}
