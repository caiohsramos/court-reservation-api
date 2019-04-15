package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres drivers
)

// CreateDB creates a db connection a migrations
func CreateDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("DB error")
	}
	db.AutoMigrate(&Reservation{})

	return db
}
