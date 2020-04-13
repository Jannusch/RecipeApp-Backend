package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var dbCredentials = DBcredentials{"", "", ""}

func readDBCredentials() {

	data, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	if err := json.Unmarshal(data, &dbCredentials); err != nil {
		log.Fatal(err)
	}
}

func convertRecipe(recipestring RecipeString) Recipe {

	var recipe Recipe
	recipe.ID, _ = uuid.Parse(recipestring.ID)
	recipe.Name = recipestring.Name
	recipe.Date = time.Now()
	recipe.Text = recipestring.Text
	recipe.Difficulty = recipestring.Difficulty
	recipe.Time, _ = strconv.Atoi(recipestring.Time)
	recipe.Rating = recipestring.Rating
	recipe.RecipebookID, _ = uuid.Parse(recipestring.RecipebookID)
	return recipe
}

func convertIngredient(ingredientString IngredientString, recipeID uuid.UUID) Ingredient {
	var ingredient Ingredient
	ingredient.Name = ingredientString.Name
	ingredient.Quantity = ingredientString.Quantity
	ingredient.Unit = ingredientString.Unit
	ingredient.ID, _ = uuid.Parse(ingredientString.ID)
	ingredient.RecipeID = recipeID
	return ingredient
}

func insertIngredients(ingrediants IngredientsString, recipeID uuid.UUID) bool {
	for _, ingredientStrings := range ingrediants {
		var ingredient = convertIngredient(ingredientStrings, recipeID)
		succes := insertIngredient(ingredient)
		if !succes {
			return false
		}
	}
	return true
}

func insertIngredient(ingredient Ingredient) bool {
	db, err := sql.Open("postgres", "dbname="+dbCredentials.DBname+" user="+dbCredentials.DBuser+" password="+dbCredentials.DBpassword+" sslmode=disable")
	if err != nil {
		fmt.Print("Now I'm there")
		fmt.Print(err)
		log.Fatal(err)
		return false
	}

	var emptyUUID uuid.UUID
	if ingredient.ID == emptyUUID {
		ingredient.ID = uuid.New()
	}

	ins := "INSERT INTO recipes (name, quantity, recipe-id, id, unit) VALUES ($1, $2, $3, $4, $5);"
	_, err = db.Exec(ins, ingredient.Name, ingredient.Quantity, ingredient.RecipeID, ingredient.ID, ingredient.Unit)

	if err != nil {
		return false
	}

	return true
}

func getAllRecipebooks() (Recipebooks, bool) {
	sel := "SELECT * FROM recipebooks;"
	if dbCredentials.DBname == "" {
		readDBCredentials()
	}

	db, err := sql.Open("postgres", "dbname="+dbCredentials.DBname+" user="+dbCredentials.DBuser+" password="+dbCredentials.DBpassword+" sslmode=disable")
	if err != nil {
		fmt.Print("Now I'm there")
		fmt.Print(err)
		log.Fatal(err)
		return Recipebooks{}, false
	}

	rows, err := db.Query(sel)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var recipebooks Recipebooks

	for rows.Next() {
		var recipebook Recipebook
		if err := rows.Scan(&recipebook.UUID, &recipebook.Name, &recipebook.Author); err != nil {
			log.Fatal(err)
		}
		recipebooks = append(recipebooks, recipebook)
	}
	fmt.Print(recipebooks)
	return recipebooks, true
}

func insertRecipe(recipeString RecipeString) (bool, uuid.UUID) {
	var recipe = convertRecipe(recipeString)
	if dbCredentials.DBname == "" {
		readDBCredentials()
	}

	var emptyUUID uuid.UUID
	if recipe.ID == emptyUUID {
		recipe.ID = uuid.New()
	}

	var emptyDate time.Time
	if recipe.Date == emptyDate {
		recipe.Date = time.Now()
	}

	db, err := sql.Open("postgres", "dbname="+dbCredentials.DBname+" user="+dbCredentials.DBuser+" password="+dbCredentials.DBpassword+" sslmode=disable")
	if err != nil {
		fmt.Print("Now I'm there")
		fmt.Print(err)
		log.Fatal(err)
		// Only for testing
		return false, recipe.ID
		// return false, uuid.New()
	}

	ins := "INSERT INTO recipes (id, name, date, difficulty, time, text, rating, recipebook-id) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	_, err = db.Exec(ins, recipe.ID, recipe.Name, recipe.Date, recipe.Difficulty, recipe.Time, recipe.Text, recipe.Rating, recipe.RecipebookID)

	if err != nil {
		fmt.Print(err)
		// Only for testing
		return false, recipe.ID
		// return false, uuid.New()
	}
	return false, recipe.ID
}

// DBcredentials contains the crendetials for the database that stores the recipe values
type DBcredentials struct {
	DBname     string `json:"dbname"`
	DBuser     string `json:"dbuser"`
	DBpassword string `json:"dbpassword"`
}
