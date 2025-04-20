package crawler

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/pkg/cache"
	"github.com/nktauserum/crawler-service/pkg/format"
)

func ParseHTML(ctx context.Context, url string) (*common.Page, error) {
	var crawler Crawler

	err := crawler.SetDestination(url)
	if err != nil {
		return nil, err
	}

	article, err := crawler.Readable()
	if err != nil {
		return nil, err
	}

	return &common.Page{
		Title:    article.Title,
		URL:      url,
		Sitename: article.SiteName,
		HTML:     article.Content,
	}, nil
}

func ParsePDF(ctx context.Context, link string) (*common.Page, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("ParsePDF took %s\n", elapsed)
	}()

	downloadStart := time.Now()
	pdf_file, err := DownloadFile(link)
	if err != nil {
		return nil, err
	}
	fmt.Printf("DownloadFile took %s\n", time.Since(downloadStart))

	processStart := time.Now()
	text, err := ProcessPDF(pdf_file)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ProcessPDF took %s\n", time.Since(processStart))

	return &common.Page{
		Content: text,
	}, nil
}

func GetContent(ctx context.Context, link string) (common.Page, error) {
	cache := cache.NewCache()
	cached_page, exists := cache.Get(link)
	if exists {
		return cached_page.Page, nil
	}

	page_url, err := url.Parse(link)
	if err != nil {
		return common.Page{}, err
	}

	ext := filepath.Ext(page_url.Path)
	if ext == ".pdf" {
		content, err := ParsePDF(ctx, link)
		if err != nil {
			return common.Page{}, err
		}
		return *content, nil
	}

	content, err := ParseHTML(ctx, link)
	if err != nil {
		return common.Page{}, err
	}

	content.Content, err = format.HTMLtoMarkdown(&content.HTML)
	if err != nil {
		return common.Page{}, err
	}

	cache.Set(link, *content)

	return *content, nil
}
