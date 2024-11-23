package models

import (
	"gorm.io/datatypes"
)

type Profile struct {
	ID       string         `gorm:"type:text;primaryKey" json:"id"`
	UserID   string         `gorm:"type:text" json:"userId"`
	User     User           `gorm:"foreignKey:UserID;references:ID;constraint:onDelete:CASCADE" json:"user"`
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
}

type ProfileResponse struct {
	ID		string `json:"id"`
	UserID	string	`json:"userId"`
	Metadata datatypes.JSON `json:"metadata"`
}