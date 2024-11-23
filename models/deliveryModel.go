package models

import (
	"time"
)

type Delivery struct {

	ID                string         `gorm:"type:text;primaryKey"`
	OrderID           string         `gorm:"type:text;notNull"`
	Order             Order          `gorm:"foreignKey:OrderID;references:ID"`
	DeliveryPersonID  string         `gorm:"type:text;notNull"`
	DeliveryPerson    User           `gorm:"foreignKey:DeliveryPersonID;references:ID"`
	Status            DeliveryStatus `gorm:"type:text;notNull"`
	EstDeliveryTime   time.Time      `gorm:"type:timestamp;notNull"`
	CreatedAt         time.Time      `gorm:"type:timestamp;notNull;default:current_timestamp"`
	UpdatedAt         *time.Time     `gorm:"type:timestamp"`
}
