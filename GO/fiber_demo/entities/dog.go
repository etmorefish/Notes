package entities

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Breed     string `json:"breed"`
	IsGoodBoy bool   `json:"is_good_boy" gorm:"default:true"`
}
