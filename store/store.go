package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

type Store interface {
	Create(book *Book) error
	Update(book *Book) error
	Get(id string) (Book, error)
	GetAll() ([]Book, error)
	Delete(id string) error
}
