package model

import (
	
)

type Item struct {
	ItemId int `json:"itemId"`
	ItemName string	`json:"itemName"`
	Brand string	`json:"brand"`
	Quantity int	`json:"quantity"`
	Price float64	`json:"price"`
	Category string	`json:"category"`
}