package store

import (
	"context"

	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/models"
)

type BookStore interface {
	BookRepository
}

type BookRepository interface {
	Create(ctx context.Context, books *models.Book) error
	GetAll(ctx context.Context) ([]*models.Book, error)
	GetById(ctx context.Context, id int) (*models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int) error
}
