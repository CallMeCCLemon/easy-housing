package model

import "gorm.io/gorm"

type Home struct {
	gorm.Model
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	OwnerID   int     `json:"ownerId"`
	Owner     User    `gorm:"foreignKey:OwnerID"`
}
