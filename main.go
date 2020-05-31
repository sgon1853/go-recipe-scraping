package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/sgon1853/go-recipe-scraping/pkg/scraper"
)

type Recipe struct {
	Title       string       `selector:"p"`
	Ingredients []Ingredient `selector:"ul li"`
}

type Ingredient struct {
}

func main() {
	url := "https://www.monsieur-cuisine.com/es/recetas/detalle/batido-matutino/"
	htmlSelector := ".recipe--ingredients-html-item .col-md-8"

	s := scraper.Scraper{}
	result, _ := s.ScrapeUrl(url, htmlSelector, parserfunc)

	fmt.Println("Scrapped!", result)
}

func parserfunc(e *colly.HTMLElement) (interface{}, error) {
	r := &Recipe{}
	if err := e.Unmarshal(r); err != nil {
		return nil, err
	}

	return r, nil
}
