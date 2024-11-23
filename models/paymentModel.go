package models

import (
	"time"
)

type Payment struct {

	ID            string        `gorm:"type:text;primaryKey"`
	OrderID       string        `gorm:"type:text;notNull"`
	Order         Order         `gorm:"foreignKey:OrderID;references:ID"`
	UserID        string        `gorm:"type:text;notNull"`
	User          User          `gorm:"foreignKey:UserID;references:ID"`
	Amount        float64       `gorm:"type:decimal(10,2);notNull"`
	PaymentMethod PaymentMethod `gorm:"type:text;notNull"`
	CreatedAt     time.Time     `gorm:"type:timestamp;notNull;default:current_timestamp"`
}
