package model

import "gorm.io/gorm"

type SupportingDoc struct {
	gorm.Model
	HomeID           int                 `json:"homeId"`
	Home             Home                `json:"home"`
	DocumentLocation string              `json:"documentLocation"`
	DocumentDate     string              `json:"documentDate"`
	Status           SupportingDocStatus `json:"status"`
}

type SupportingDocStatus string

const (
	SupportingDocStatusPending  SupportingDocStatus = "pending"
	SupportingDocStatusApproved SupportingDocStatus = "approved"
	SupportingDocStatusRejected SupportingDocStatus = "rejected"
)
