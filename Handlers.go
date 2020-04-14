package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// AllRecipeBooks return all Recipebooks that stored in the database
func AllRecipeBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	recipebooks, err := getAllRecipebooks()
	if err != nil {
		w.WriteHeader(302)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(recipebooks)

	/*var allRecipeBooksVariable = Recipebooks{
		Recipebook{
			"50 Soßen eine Nudel",
			"24ba3f2e-dd11-4153-8318-b75de9310ee3",
			"Margarete",
		},
		Recipebook{
			"Ruhrpotesser",
			"24ba3f2e-dd11-4153-8318-wefewfwef234",
			"Frank Weintraube",
		},
		Recipebook{
			"Wokgemuche",
			"24ba3f2e-dd11-4153-8318-fgrg3456346ef",
			"Anton Antonson",
		},
	}
	*/

}

// AllSpecificRecipes return all recipes to one given recipebook
func AllSpecificRecipes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars["id"])
}

// RecipebookDetails returns a json with details to a given recipebook
func RecipebookDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]
	fmt.Print(uuid)

	recipebook, err := getRecipebookDetails(uuid)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(400)
		fmt.Print(w, "I'm Sorry but the requested uuid is not valid")
	}

	json.NewEncoder(w).Encode(recipebook)
}

// RecipeAdd handle a new incomming recipe
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

	var recipe RecipeString = f.Recipe
	recipeID, err := insertRecipe(recipe)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(422)
		return
	}

	var ingredientsString IngredientsString = f.Ingrediants
	err = insertIngredients(ingredientsString, recipeID)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(422)
		return
	}

	w.WriteHeader(200)

}

// Index is the Indexpage \(°^°)/
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Print("Haleluja")
}
