package models

// Book general information about a book
type Book struct {
	ID            string
	Title         string
	Subtitle      string
	Description   string
	Authors       []string
	Publisher     string
	PublishedDate string
	PageCount     int64
	ImageURL      string
	Isbn          string
}
