package controllers

import (
	"fmt"
	"net/http"

	"alexandria.backend/mappers"
	"alexandria.backend/services"
	"github.com/gin-gonic/gin"
)

// GET /books/:isbn
// Retrieve a book based on ISBN
func GetBookByIsbn(context *gin.Context) {
	googleVolumeService := services.NewGoogleVolumeService()

	isbn := context.Param("isbn")
	volumes, err := googleVolumeService.GetVolumesByIsbn(isbn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to get volumes by ISBN: isbn=%v, err=%s", isbn, err)})
		return
	}

	if volumes.TotalItems > 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("More than one volume returned by google volume service: isbn=%v", isbn)})
		return
	}

	book := mappers.VolumeToBook(volumes.Items[0])
	context.JSON(http.StatusOK, gin.H{"data": book})
}
