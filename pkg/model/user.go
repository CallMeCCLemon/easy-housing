package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OidcUserID string `json:"oidcUserId"`
	Email      string `json:"email"`
	Name       string `json:"name"`
}
