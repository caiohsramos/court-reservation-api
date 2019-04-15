package handlers

import (
	"fmt"
	"strconv"

	"github.com/caiohsramos/court-reservation-api/src/repo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAllReservations is a handler to get all reservations
func GetAllReservations(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result repo.Reservation
		db.Find(&result)
		c.JSON(200, result)
	}
}

// GetReservationID is a handler to get a reservation by ID
func GetReservationID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var result repo.Reservation
		err := db.First(&result, id).Error
		if err != nil {
			c.JSON(401, gin.H{
				"status": fmt.Sprint("Not Found ", id),
				"error":  err,
			})
			return
		}

		c.JSON(200, result)
	}
}
