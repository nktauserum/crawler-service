package crawler

import (
	"io"
	"net/url"
	"strings"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/go-shiori/go-readability"
)

type Crawler struct {
	url    url.URL
	source io.ReadCloser
}

func (c *Crawler) SetDestination(rawURL string) error {
	url, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	c.url = *url
	return nil
}

func (c *Crawler) Readable() (*readability.Article, error) {
	err := c.request()
	if err != nil {
		return nil, err
	}

	article, err := readability.FromReader(c.source, &c.url)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (c *Crawler) request() error {
	client := cycletls.Init()

	resp, err := client.Do(c.url.String(), cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
	}, "GET")
	if err != nil {
		return err
	}

	c.Close()
	c.source = io.NopCloser(strings.NewReader(resp.Body))
	return nil
}

func (c *Crawler) Close() {
	if c.source != nil {
		c.source.Close()
		c.source = nil
	}
}
