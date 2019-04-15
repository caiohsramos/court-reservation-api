package repo

import "github.com/jinzhu/gorm"

// Reservation is a model for the DB
type Reservation struct {
	gorm.Model
	NumberOfPeople int    `json:"number_people"`
	Apartment      int    `json:"apt"`
	Name           string `json:"name"`
	Use            string `json:"use"`
}
