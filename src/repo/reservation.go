package repo

import "github.com/jinzhu/gorm"

type Reservation struct {
	gorm.Model
	NumberOfPeople int    `json:"number_people"`
	Apartment      int    `json:"apt"`
	Name           string `json:"name"`
	Use            string `json:"use"`
}

func (repo *repo) GetAllReservations() []Reservation {
	var result []Reservation
	repo.db.Find(&result)
	return result
}

func (repo *repo) GetReservationID(id int) (Reservation, error) {
	var result Reservation
	err := repo.db.First(&result, id).Error

	return result, err
}
