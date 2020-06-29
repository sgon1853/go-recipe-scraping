package csvexport

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/sgon1853/go-recipe-scraping/pkg/models"
)

func ExportRecipe(csvPath string, recipesInput *models.Recipe) {

	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		if err = os.MkdirAll(csvPath, os.ModePerm); err != nil {
			panic(err)
		}
	}

	fileName := strings.ReplaceAll(recipesInput.Title, " ", "-")
	fileName = strings.ToLower(fileName)
	fullFileName := path.Join(csvPath, fmt.Sprintf("%s.csv", fileName))

	ingredientsOut := []*models.IngredientCsvOutput{}

	for _, d := range recipesInput.Dishes {
		for _, ing := range d.Ingredients {
			ingredientsOut = append(ingredientsOut, &models.IngredientCsvOutput{Recipe: recipesInput.Title,
				Dish: d.Title, Ingredient: ing.Description})
		}
	}

	dataFile, err := os.OpenFile(fullFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = ';'
		return gocsv.NewSafeCSVWriter(writer)
	})

	err = gocsv.MarshalFile(ingredientsOut, dataFile)
	if err != nil {
		panic(err)
	}
}
