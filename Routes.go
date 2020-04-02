package main

import (
	"net/http"
)

// Route is the definition of one specific route for the Router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list with all routes that the Router can handle
type Routes []Route

var routers = Routes{
	Route{
		"AllRecipeBooks",
		"GET",
		"/recipebooks",
		AllRecipeBooks,
	},
	{
		"AllRecipesSpecificBook",
		"GET",
		"/recipebooks/all/{id}",
		AllSpecificRecipes,
	},
	{
		"RecipebookDetails",
		"GET",
		"/recipebooks/{id}",
		RecipebookDetails,
	},
	{
		"AddRecipe",
		"POST",
		"/recipe",
		RecipeAdd,
	},
	{
		"Index",
		"GET",
		"/",
		Index,
	},
}
