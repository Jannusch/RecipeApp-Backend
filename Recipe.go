package main

import (
	"time"

	"github.com/google/uuid"
)

// Recipe define the basic structure of a recipe
type Recipe struct {
	Name         string    `json:"name"`
	ID           uuid.UUID `json:"id"`
	Date         time.Time `json:"date"`
	Difficulty   int       `json:"difficulty"`
	Time         int       `json:"time"`
	Text         string    `json:"text"`
	Rating       int       `json:"rating"`
	RecipebookID uuid.UUID `json:"recipebook-id"`
}

// Recipes is a list of Recipes
type Recipes []Recipe

// Recipebook define the basic structure of a Recipebook
type Recipebook struct {
	Name   string    `json:"name"`
	UUID   uuid.UUID `json:"uuid"`
	Author string    `json:"author"`
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

// Ingredients is a list with ingrediants
type Ingredients []Ingredient

// RecipeWithIngrediants is the structure that the server will get from the Frontend
type RecipeWithIngrediants struct {
	Recipe      Recipe       `json:"recipe"`
	Ingrediants []Ingredient `json:"ingrediants"`
}
