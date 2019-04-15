package main

import (
	"github.com/caiohsramos/court-reservation-api/src/handlers"
	"github.com/caiohsramos/court-reservation-api/src/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	db := repo.CreateDB()
	defer db.Close()
	r := gin.Default()
	r.GET("/", handlers.Index)
	r.GET("/reservation", handlers.GetAllReservations(db))
	r.GET("/reservation/:id", handlers.GetReservationID(db))

	r.Run(":3001")
}
