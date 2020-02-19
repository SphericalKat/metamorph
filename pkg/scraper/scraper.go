package scraper

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
)

type Scraper struct{}

func (s *Scraper) Get(url string) (*goquery.Document, error) {
	status, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}

	if status < 200 || status >= 400 {
		return nil, fmt.Errorf("couldn't retrieve document: status %d", status)
	}
	return goquery.NewDocumentFromReader(bytes.NewReader(body))
}

func (s *Scraper) GetSelection(url string, selector string) (*goquery.Selection, error) {
	doc, err := s.Get(url)
	if err != nil {
		return nil, err
	}
	return doc.Find(selector), nil
}
