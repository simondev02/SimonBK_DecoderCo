package models

import (
	"gorm.io/gorm"
)

type AvlRecord struct {
	gorm.Model
	Imei      *string  `json:"imei"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
