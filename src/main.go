package main

import (
	"github.com/caiohsramos/court-reservation-api/src/handlers"
	"github.com/caiohsramos/court-reservation-api/src/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repo.CreateRepo()
	defer repo.CloseRepo()
	r := gin.Default()
	r.GET("/", handlers.Index)
	r.GET("/reservation", handlers.GetAllReservations(repo))
	r.GET("/reservation/:id", handlers.GetReservationID(repo))
	r.Run(":3001")
}
