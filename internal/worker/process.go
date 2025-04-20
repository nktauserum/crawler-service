package worker

import (
	"context"

	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/pkg/crawler"
)

func ProcessURL(url string) (common.Page, error) {
	page, err := crawler.GetContent(context.Background(), url)
	if err != nil {
		return common.Page{}, err
	}

	return page, nil
}
