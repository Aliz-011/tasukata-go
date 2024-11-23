package models

import (
	"time"
)

type Order struct {
	ID         string     `gorm:"type:text;primaryKey" json:"id"`
	Address    string     `gorm:"type:text;notNull" json:"address"`
	Latitude   float64    `gorm:"type:numeric(10,6);notNull" json:"latitude"`
	Longitude  float64    `gorm:"type:numeric(10,6);notNull" json:"longitude"`
	Total      int        `gorm:"type:integer;notNull" json:"total"`
	Status     StatusEnum `gorm:"type:text;notNull" json:"status"`
	CustomerID string     `gorm:"type:text;notNull" json:"customerId"`
	Customer   Profile    `gorm:"foreignKey:CustomerID;references:ID;constraint:onDelete:Cascade" json:"customer"`
	CreatedAt  time.Time  `gorm:"type:timestamp;notNull" json:"createdAt"`
	OrderItems []OrderItem  `gorm:"foreignKey:OrderID;references:ID" json:"orderItems"`
}

type OrderResponse struct {
	ID 			string 				`json:"id"`
	Address   	string              `json:"address"`
	Total     	int                 `json:"total"`
	CustomerID 	string 				`json:"customerId"`
	Customer  	ProfileResponse     `json:"customer"`
	OrderItems  []OrderItemResponse `json:"orderItems"`
	Latitude 	float64				`json:"latitude"`
	Longitude 	float64				`json:"longitude"`
	Status		string				`json:"status"`
	CreatedAt	time.Time			`json:"createdAt"`
}