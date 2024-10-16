package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// "time"

	// "os"

	"github.com/SohomSaha045/POS_backend/model"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var database *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	dsn := os.Getenv("URL")

	// fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	database = db
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// res, err := db.Exec("CREATE TABLE cars (  brand VARCHAR(255),model VARCHAR(255), year INT);")
	fmt.Printf("%T \n", db)

	// fmt.Println(rows)
}

func ItemController(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("SELECT * from cars;")
	if err != nil {
		log.Fatal("failed to execute query", err)
	}
	var cars []model.Item
	for rows.Next() {
		var newCar model.Item
		err := rows.Scan(&newCar.Brand, &newCar.Model, &newCar.Year)
		if err != nil {
			log.Fatal(err)
		}
		cars = append(cars, newCar)
		// fmt.Printf("Brand: %s, Model: %s, Year: %d\n", brand, model, year)
	}
	fmt.Println(cars)

	json.NewEncoder(w).Encode("All movies Deleted.")
}
