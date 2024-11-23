package models

import "time"

type Product struct {
	ID          string  `gorm:"type:text;primaryKey" json:"id"`
	Name        string  `gorm:"type:text;notNull" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);notNull" json:"price"`
	Image       string  `gorm:"type:text;notNull" json:"img"`
	CreatedAt 	time.Time `gorm:"type:timestamp;notNull" json:"createdAt"`
	CategoryID  string  `gorm:"type:text;notNull" json:"categoryId"`
	Category    Category `gorm:"foreignKey:CategoryID;references:ID;constraint:onDelete:SET NULL" json:"category"`
}

type ProductResponse struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Price		int			`json:"price"`
	Image		string		`json:"img"`
	CreatedAt	time.Time	`json:"createdAt"`
	CategoryID	string		`json:"categoryId"`
	Category    Category 	`json:"category"`
}