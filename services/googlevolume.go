package services

import (
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/api/books/v1"
	"google.golang.org/api/googleapi"
)

// GoogleVolumeService service interface to the google books api
type GoogleVolumeService struct {
	booksService *books.Service
}

// NewGoogleVolumeService creates a new GoogleVolumeService, the interface to the google books api
func NewGoogleVolumeService() *GoogleVolumeService {
	service, err := books.NewService(oauth2.NoContext)
	if err != nil {
		log.Fatalf("Failed to get google books service: %v", err)
	}

	return &GoogleVolumeService{service}
}

// GetByIsbn gets volumes through BooksService based on ISBN
func (service *GoogleVolumeService) GetVolumesByIsbn(isbn string) (*books.Volumes, error) { //does the parameter need the star????/
	callOptions := isbnCallOptions(isbn)
	return service.booksService.Volumes.List().Do(callOptions)
}

// Isbn returns a CallOption that will set the "isbn" parameter of a call
func isbnCallOptions(bookIsbn string) googleapi.CallOption { return isbn(bookIsbn) }

type isbn string

func (volumeIsbn isbn) Get() (string, string) { return "q", string("isbn:" + volumeIsbn) }
