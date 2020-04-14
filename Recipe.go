package main

import (
	"time"

	"github.com/google/uuid"
)

// Recipe define the basic structure of a recipe
type Recipe struct {
	Name         string    `json:"name"`
	ID           uuid.UUID `json:"uuid"`
	Date         time.Time `json:"date"`
	Difficulty   int       `json:"difficulty"`
	Time         int       `json:"time"`
	Text         string    `json:"text"`
	Rating       int       `json:"rating"`
	RecipebookID uuid.UUID `json:"recipebookID"`
}

// RecipeString is the same as Recipe but with strings as time and uuid
type RecipeString struct {
	Name         string `json:"name"`
	ID           string `json:"uuid"`
	Date         int    `json:"date"`
	Difficulty   int    `json:"difficulty"`
	Time         int    `json:"time"`
	Text         string `json:"text"`
	Rating       int    `json:"rating"`
	RecipebookID string `json:"recipebookID"`
}

// Recipes is a list of Recipes
type Recipes []Recipe

// Recipebook define the basic structure of a Recipebook
type Recipebook struct {
	Name   string `json:"name"`
	UUID   string `json:"uuid"` // uuid.UUID `json:"uuid"`
	Author string `json:"author"`
}

// Recipebooks is a list of Recipebooks
type Recipebooks []Recipebook

// Ingredient define the basic structure of an ingrediant
type Ingredient struct {
	Name     string    `json:"name"`
	Quantity string    `json:"quantity"`
	Unit     string    `json:"unit"`
	RecipeID uuid.UUID `json:"recipe-id"`
	ID       uuid.UUID `json:"id"`
}

// IngredientString is the same as a Ingredient struct, but only with strings
type IngredientString struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
	// RecipeID string `json:"recipe-id"`
	ID string `json:"id"`
}

// Ingredients is a list with ingrediants
type Ingredients []Ingredient

// IngredientsString is a list with ingredients as string structs
type IngredientsString []IngredientString

// RecipeWithIngrediants is the structure that the server will get from the Frontend
type RecipeWithIngrediants struct {
	Recipe      RecipeString       `json:"recipe"`
	Ingrediants []IngredientString `json:"ingrediants"`
}
