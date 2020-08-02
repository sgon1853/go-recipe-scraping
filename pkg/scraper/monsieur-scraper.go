package scraper

import (
	"strings"
	"unicode"

	"github.com/gocolly/colly/v2"
	"github.com/sgon1853/go-recipe-scraping/pkg/models"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func ScrapeMonsieurUrl(url string) (*models.Recipe, error) {

	htmlSelector := "div#recipe-detail"

	s := Scraper{}
	result, _ := s.ScrapeUrl(url, htmlSelector, parserfunc)

	return (result).(*models.Recipe), nil
}

func parserfunc(e *colly.HTMLElement) (interface{}, error) {

	titSelector := ".recipe--header div>div>h1"
	ingSelector := ".recipe--ingredients-html-item.col-md-8"

	recTitle := removeAccents(e.ChildText(titSelector))

	r := &models.Recipe{
		Title:  recTitle,
		Dishes: []*models.Dish{}}

	var currDish *models.Dish

	e.ForEach(ingSelector, func(_ int, elout *colly.HTMLElement) {
		elout.ForEach("*", func(_ int, el *colly.HTMLElement) {
			// fmt.Println("Name", el.Name)
			// fmt.Println("Text", el.Text)
			// fmt.Println()

			if el.Name == "p" {
				currDish = &models.Dish{
					Title:       removeAccents(el.Text),
					Ingredients: []models.Ingredient{},
				}

				r.Dishes = append(r.Dishes, currDish)
			} else if el.Name == "li" {
				currIng := models.Ingredient{
					Description: strings.TrimSpace(removeAccents(el.Text)),
				}

				if currDish == nil {
					currDish = &models.Dish{
						Title:       recTitle,
						Ingredients: []models.Ingredient{},
					}
					r.Dishes = append(r.Dishes, currDish)
				}

				currDish.Ingredients = append(currDish.Ingredients, currIng)
			}
		})
	})

	return r, nil
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}
