package main

import (
	"database/sql"
	"encoding/json"
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

func insertRecipe(recipe Recipe) bool {
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
		return false
	}

	ins := "INSERT INTO recipes (id, name, date, difficulty, time, text, rating, recipebook-id) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	_, err = db.Exec(ins, recipe.ID, recipe.Name, recipe.Date, recipe.Difficulty, recipe.Time, recipe.Text, recipe.Rating, recipe.RecipebookID)

	if err != nil {
		fmt.Print(err)
		return false
	}

	return true
}

func insertIngrediants(ingrediants Ingredients) bool {
	for _, ingredient := range ingrediants {
		succes := insertIngrediant(ingredient)
		if !succes {
			return false
		}
	}
	return true
}

func insertIngrediant(ingredient Ingredient) bool {
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

// DBcredentials contains the crendetials for the database that stores the recipe values
type DBcredentials struct {
	DBname     string `json:"dbname"`
	DBuser     string `json:"dbuser"`
	DBpassword string `json:"dbpassword"`
}
