package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"alexandria.backend/mappers"
	"alexandria.backend/services"
	"github.com/gin-gonic/gin"
)

// GetBookByIsbnRequest Get book by ISBN request object
type GetBookByIsbnRequest struct {
	Isbn string `uri:"isbn" binding:"required"`
}

// RemoveDashesFromIsbn Removes any dashes from ISBN
func (req *GetBookByIsbnRequest) RemoveDashesFromIsbn() {
	req.Isbn = strings.ReplaceAll(req.Isbn, "-", "")
}

// GetBookByIsbn godoc
// @Summary Retrieves book based on ISBN
// @Produce json
// @Param isbn path string true "ISBN"
// @Success 200 {object} models.Book
// @Router /books/{isbn} [get]
func GetBookByIsbn(context *gin.Context) {
	googleVolumeService := services.NewGoogleVolumeService()

	var req GetBookByIsbnRequest
	if err := context.BindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	req.RemoveDashesFromIsbn()
	volumes, err := googleVolumeService.GetVolumesByIsbn(req.Isbn)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get volumes by ISBN: isbn=%v, err=%s", req.Isbn, err)})
		return
	}

	if volumes.TotalItems == 0 {
		context.JSON(http.StatusOK, gin.H{})
		return
	}

	if volumes.TotalItems > 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("More than one volume returned by google volume service: isbn=%v", req.Isbn)})
		return
	}

	book := mappers.VolumeToBook(volumes.Items[0])
	context.JSON(http.StatusOK, gin.H{"data": book})
}
