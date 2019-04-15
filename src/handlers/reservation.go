package handlers

import (
	"fmt"
	"strconv"

	"github.com/caiohsramos/court-reservation-api/src/repo"
	"github.com/gin-gonic/gin"
)

func GetAllReservations(repo repo.RepoAccess) gin.HandlerFunc {
	return func(c *gin.Context) {
		reservations := repo.GetAllReservations()
		c.JSON(200, reservations)
	}
}

func GetReservationID(repo repo.RepoAccess) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		reservation, err := repo.GetReservationID(id)
		if err != nil {
			c.JSON(401, gin.H{"status": fmt.Sprint("Not Found ", id)})
			return
		}

		c.JSON(200, reservation)
	}
}
