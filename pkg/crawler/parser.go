package crawler

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"github.com/nktauserum/crawler-service/pkg/format"
)

type Page struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Sitename string `json:"sitename"`
	Content  string `json:"content"`
	HTML     string `json:"html"`
}

func ParseHTML(ctx context.Context, url string) (*Page, error) {
	var crawler Crawler

	err := crawler.SetDestination(url)
	if err != nil {
		return nil, err
	}

	article, err := crawler.Readable()
	if err != nil {
		return nil, err
	}

	return &Page{
		Title:    article.Title,
		URL:      url,
		Sitename: article.SiteName,
		HTML:     article.Content,
	}, nil
}

func ParsePDF(ctx context.Context, link string) (*Page, error) {
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

	return &Page{
		Content: text,
	}, nil
}

func GetContent(ctx context.Context, link string) (Page, error) {
	page_url, err := url.Parse(link)
	if err != nil {
		return Page{}, err
	}

	ext := filepath.Ext(page_url.Path)
	if ext == ".pdf" {
		content, err := ParsePDF(ctx, link)
		if err != nil {
			return Page{}, err
		}
		return *content, nil
	}

	content, err := ParseHTML(ctx, link)
	if err != nil {
		return Page{}, err
	}

	content.Content, err = format.HTMLtoMarkdown(&content.HTML)
	if err != nil {
		return Page{}, err
	}

	return *content, nil
}
