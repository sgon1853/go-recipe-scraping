package main

import (
	"fmt"
	"os"
	"path"

	csvexport "github.com/sgon1853/go-recipe-scraping/pkg/csvExport"
	"github.com/sgon1853/go-recipe-scraping/pkg/scraper"
)

func main() {
	// url := "https://www.monsieur-cuisine.com/es/recetas/detalle/lahmacun-pizza-turca/"
	url := "https://www.monsieur-cuisine.com/es/recetas/detalle/tarta-de-calabaza-con-hierbas-y-avellanas/"

	res, err := scraper.ScrapeMonsieurUrl(url)
	if err != nil {
		panic(fmt.Sprintf("Error scraping url %s, Error: %s", url, err))
	}

	fmt.Println("Scrapped successfully", url)

	csvPath := getCsvPath()
	fmt.Println("Saving into csv", csvPath)
	csvexport.ExportRecipe(csvPath, res)
}

func getCsvPath() string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "downloaded")
}
