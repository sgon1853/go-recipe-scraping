package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/sgon1853/go-recipe-scraping/pkg/scraper"
)

type Recipe struct {
	Dishes []*Dish
}

type Dish struct {
	Title       string
	Ingredients []Ingredient
}

type Ingredient struct {
	Description string
}

func main() {
	// url := "https://www.monsieur-cuisine.com/es/recetas/detalle/lahmacun-pizza-turca/"
	url := "https://www.monsieur-cuisine.com/es/recetas/detalle/tarta-de-calabaza-con-hierbas-y-avellanas/"

	htmlSelector := ".recipe--ingredients-html-item.col-md-8"

	s := scraper.Scraper{}
	result, _ := s.ScrapeUrl(url, htmlSelector, parserfunc)

	fmt.Println("Scrapped!", result)
}

func parserfunc(e *colly.HTMLElement) (interface{}, error) {
	r := &Recipe{[]*Dish{}}

	if err := e.Unmarshal(r); err != nil {
		return nil, err
	}

	var currDish *Dish
	e.ForEach("*", func(_ int, el *colly.HTMLElement) {
		fmt.Println("Name", el.Name)
		fmt.Println("Text", el.Text)
		fmt.Println()

		if el.Name == "p" {
			currDish = &Dish{
				Title:       el.Text,
				Ingredients: []Ingredient{},
			}

			r.Dishes = append(r.Dishes, currDish)
		} else if el.Name == "li" {
			currIng := Ingredient{
				Description: el.Text,
			}

			if currDish != nil {
				currDish.Ingredients = append(currDish.Ingredients, currIng)
			}
		}
	})

	return r, nil
}
