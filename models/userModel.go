package models

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {

	ID            string         `gorm:"type:text;primaryKey" json:"id"`
	Name          string         `gorm:"type:text;notNull" json:"name"`
	Email         string         `gorm:"type:text;unique;notNull" json:"email"`
	Password      string         `gorm:"type:text;notNull" json:"password"`
	Role          UserRole       `gorm:"type:text;notNull;default:'customer'" json:"role"`
	EmailVerified int            `gorm:"type:integer;notNull;default:0" json:"emailVerified"`
	TOTPKey       datatypes.JSON `gorm:"type:json" json:"totpKey"`
	RecoveryCode  datatypes.JSON `gorm:"type:json;notNull" json:"recoveryCode"`
	CreatedAt     time.Time      `gorm:"type:timestamp;notNull" json:"createdAt"`
}