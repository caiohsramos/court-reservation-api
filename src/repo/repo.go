package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres drivers
)

type repo struct {
	db *gorm.DB
}

type RepoAccess interface {
	GetAllReservations() []Reservation
	GetReservationID(int) (Reservation, error)
	CloseRepo()
}

func CreateRepo() RepoAccess {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("DB error")
	}
	db.AutoMigrate(&Reservation{})

	return &repo{db}
}

func (repo *repo) CloseRepo() {
	repo.db.Close()
}
