package models

type Category struct {	
	ID   string    `gorm:"type:text;primaryKey" json:"id"`
	Name string    `gorm:"type:text;notNull" json:"name"`
}