package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	csvexport "github.com/sgon1853/go-recipe-scraping/pkg/csvExport"
	"github.com/sgon1853/go-recipe-scraping/pkg/scraper"
	"github.com/spf13/viper"
)

func main() {
	var rsConfigPath string
	defaultConfigPath := getDefaultConfigPath()
	flag.StringVar(&rsConfigPath, "config", defaultConfigPath, "The path where the config is located")
	flag.Parse()

	fmt.Println("Reading config file", rsConfigPath)

	data, err := ioutil.ReadFile(rsConfigPath)
	check(err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data)) // Find and read the config file
	if err != nil {                               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	recUrls := viper.GetStringSlice("recipes")

	for _, url := range recUrls {
		res, err := scraper.ScrapeMonsieurUrl(url)
		if err != nil {
			panic(fmt.Sprintf("Error scraping url %s, Error: %s", url, err))
		}

		fmt.Println("Scrapped successfully", url)

		csvPath := getCsvPath()
		fmt.Println("Saving into csv", csvPath)
		csvexport.ExportRecipe(csvPath, res)
	}
}

func getDefaultConfigPath() string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "config.yaml")
}

func getCsvPath() string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "downloaded")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
