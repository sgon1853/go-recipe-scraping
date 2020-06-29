package models

type Recipe struct {
	Title  string
	Dishes []*Dish
}

type Dish struct {
	Title       string
	Ingredients []Ingredient
}

type Ingredient struct {
	Description string
}

type IngredientCsvOutput struct {
	Ingredient string `csv:"Ingredient"`
	Dish       string `csv:"Dish"`
	Recipe     string `csv:"Recipe"`
}
