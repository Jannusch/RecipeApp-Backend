package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func AllRecipeBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Print("NOOOOOO!!!!")
}

func AllSpecificRecipes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars["id"])
}

func RecipebookDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Print("no!")
}

func RecipeAdd(w http.ResponseWriter, r *http.Request) {
	var f RecipeWithIngrediants
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err := json.Unmarshal(body, &f); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	var recipe Recipe = f.Recipe
	var ingredients Ingredients = f.Ingrediants

	succesInsertRecipe := insertRecipe(recipe)

	succesInsertIngrediants := insertIngrediants(ingredients)

	if succesInsertIngrediants && succesInsertRecipe {
		w.WriteHeader(200)
	}
	w.WriteHeader(422)
}
