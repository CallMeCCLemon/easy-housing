package model

import "gorm.io/gorm"

type Home struct {
	gorm.Model
	Address string `json:"address"`
	OwnerID int    `json:"ownerId"`
	Owner   User   `gorm:"foreignKey:OwnerID"`
}
