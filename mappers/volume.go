package mappers

import (
	"alexandria.backend/models"

	"google.golang.org/api/books/v1"
)

// VolumeToBook maps a google volume to an Alexandria book
func VolumeToBook(volume *books.Volume) *models.Book {
	if volume.VolumeInfo == nil {
		return nil
	}

	isbn := ""
	if volume.VolumeInfo.IndustryIdentifiers != nil && len(volume.VolumeInfo.IndustryIdentifiers) > 0 {
		for _, industryIdentifiers := range volume.VolumeInfo.IndustryIdentifiers {
			if industryIdentifiers.Type == "ISBN_13" {
				isbn = industryIdentifiers.Identifier
				break
			}

			if industryIdentifiers.Type == "ISBN_10" {
				isbn = industryIdentifiers.Identifier
			}
		}
	}

	return &models.Book{
		ID:            volume.Id,
		Title:         volume.VolumeInfo.Title,
		Subtitle:      volume.VolumeInfo.Subtitle,
		Description:   volume.VolumeInfo.Description,
		Authors:       volume.VolumeInfo.Authors,
		Publisher:     volume.VolumeInfo.Publisher,
		PublishedDate: volume.VolumeInfo.PublishedDate,
		PageCount:     volume.VolumeInfo.PageCount,
		ImageURL:      volume.VolumeInfo.ImageLinks.Large,
		Isbn:          isbn,
	}
}
