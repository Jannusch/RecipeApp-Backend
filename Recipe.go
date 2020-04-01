package main

import (
	"time"

	"github.com/google/uuid"
)

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

type Recipes []Recipe

type Recipebook struct {
	Name   string    `json:"name"`
	UUID   uuid.UUID `json:"uuid"`
	Author string    `json:"author"`
}

type Recipebooks []Recipebook

type Ingredient struct {
	Name     string    `json:"name`
	Quantity string    `json:"quantity"`
	Unit    string    `json:"unit"`
	RecipeID uuid.UUID `json:"recipe-id"`
	ID       uuid.UUID `json:"id"`
}

type Ingredients []Ingredient

type RecipeWithIngrediants struct {
	Recipe      Recipe       `json:"recipe"`
	Ingrediants []Ingredient `json:"ingrediants"`
}
