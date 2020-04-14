package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
	// layout := "2006-01-02"

	var recipe Recipe
	recipe.ID, _ = uuid.Parse(recipestring.ID)
	recipe.Name = recipestring.Name
	recipe.Date = time.Now()
	recipe.Difficulty = recipestring.Difficulty
	recipe.Time = recipestring.Time
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

func insertIngredients(ingrediants IngredientsString, recipeID uuid.UUID) error {
	for _, ingredientStrings := range ingrediants {
		var ingredient = convertIngredient(ingredientStrings, recipeID)
		succes := insertIngredient(ingredient)
		if !succes {
			return errors.New("Unabel to insert Ingredient")
		}
	}
	return nil
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

func getRecipe(uuid uuid.UUID) (Recipe, error) {
	sel := "SELECT * FROM recipes WHERE id=$1"
	if dbCredentials.DBname == "" {
		readDBCredentials()
	}

	db, err := sql.Open("postgres", "dbname="+dbCredentials.DBname+" user="+dbCredentials.DBuser+" password="+dbCredentials.DBpassword+" sslmode=disable")
	if err != nil {
		fmt.Print("Now I'm there")
		fmt.Print(err)
		log.Fatal(err)
		return Recipe{}, errors.New("unabel to open DB")
	}

	rows, err := db.Query(sel)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var recipes Recipes

	for rows.Next() {
		var recipe Recipe
		if err := rows.Scan(&recipe.Name, &recipe.ID, &recipe.Date, &recipe.Difficulty, &recipe.Time, &recipe.Text, &recipe.Rating, &recipe.RecipebookID); err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, recipe)
	}

	return recipes[0], nil
}

func getRecipebookDetails(uuidString string) (Recipebook, error) {
	uuid, err := uuid.Parse(uuidString)

	if err != nil {
		return Recipebook{}, errors.New("Unabel to parse uuid")
	}
	sel := "SELECT * FROM recipebooks WHERE id=$1"
	if dbCredentials.DBname == "" {
		readDBCredentials()
	}

	db, err := sql.Open("postgres", "dbname="+dbCredentials.DBname+" user="+dbCredentials.DBuser+" password="+dbCredentials.DBpassword+" sslmode=disable")
	if err != nil {
		fmt.Print("Now I'm there")
		fmt.Print(err)
		log.Fatal(err)
		return Recipebook{}, errors.New("unabel to open DB")
	}

	rows, err := db.Query(sel, uuid)
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
	if len(recipebooks) != 1 {
		return Recipebook{}, errors.New("To many recibooks for given uuid")
	}
	return recipebooks[0], nil
}

func getAllRecipebooks() (Recipebooks, error) {
	sel := "SELECT * FROM recipebooks;"
	if dbCredentials.DBname == "" {
		readDBCredentials()
	}

	db, err := sql.Open("postgres", "dbname="+dbCredentials.DBname+" user="+dbCredentials.DBuser+" password="+dbCredentials.DBpassword+" sslmode=disable")
	if err != nil {
		fmt.Print("Now I'm there")
		fmt.Print(err)
		log.Fatal(err)
		return Recipebooks{}, errors.New("Unabel to open DB")
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
	return recipebooks, nil
}

func insertRecipe(recipeString RecipeString) (uuid.UUID, error) {
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
		return recipe.ID, errors.New("Unabel to open DB")
		// return false, uuid.New()
	}

	fmt.Print(recipe)

	// I know the stupidity with the recipebook-id. If you know a nice solution -> pull request, please
	ins := "INSERT INTO recipes (id, name, date, difficulty, time, text, rating, \"recipebook-id\") VALUES ($1, $2, $3, $4, $5, $6, $7, (SELECT id from recipebooks WHERE id=$8));"
	_, err = db.Exec(ins, recipe.ID, recipe.Name, recipe.Date, recipe.Difficulty, recipe.Time, recipe.Text, recipe.Rating, recipe.RecipebookID)

	if err != nil {
		fmt.Print(err)
		// Only for testing
		return recipe.ID, errors.New("Unabel to insert recipe")
		// return false, uuid.New()
	}
	return recipe.ID, nil
}

// DBcredentials contains the crendetials for the database that stores the recipe values
type DBcredentials struct {
	DBname     string `json:"dbname"`
	DBuser     string `json:"dbuser"`
	DBpassword string `json:"dbpassword"`
}
