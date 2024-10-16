package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
}

func AddItemController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var items []model.Item
	e := json.NewDecoder(r.Body).Decode(&items)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(items)

	for _, item := range items {
		_, err := database.Query(`Insert into item values ($1,$2,$3,$4,$5,$6)`, item.ItemId, item.ItemName, item.Brand, item.Quantity, item.Price, item.Category)
		if err != nil {
			json.NewEncoder(w).Encode("status : SomeThing went wrong")
			return
		}
	}

	json.NewEncoder(w).Encode("status : Success")
}

func GetAllItemController(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("SELECT * from cars;")
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	// rows, err := database.Query("SELECT * from cars;")
	// if err != nil {
	// 	log.Fatal("failed to execute query", err)
	// }
	// var cars []model.Item
	// for rows.Next() {
	// 	var newCar model.Item
	// 	err := rows.Scan()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	cars = append(cars, newCar)
	// 	// fmt.Printf("Brand: %s, Model: %s, Year: %d\n", brand, model, year)
	// }
	// fmt.Println(cars)

	var cars []model.Item
	for rows.Next() {
		var newCar model.Item
		err := rows.Scan()
		if err != nil {
			log.Fatal(err)
		}
		cars = append(cars, newCar)
	}
	fmt.Println(cars)

	json.NewEncoder(w).Encode("Item has been Added.")
}
