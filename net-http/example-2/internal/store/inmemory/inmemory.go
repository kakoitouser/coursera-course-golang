package inmemory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/models"
	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/store"
)

type DB struct {
	mu   *sync.RWMutex
	data map[int]*models.Book
}

func NewDB() store.BookStore {
	db := &DB{
		data: make(map[int]*models.Book),
		mu:   new(sync.RWMutex),
	}
	db.data[1] = &models.Book{
		ID:          1,
		Name:        "Book",
		Author:      "Author",
		Description: "",
		Pages:       559,
	}
	return db
}

func (db *DB) Create(ctx context.Context, book *models.Book) error {
	if book == nil {
		return errors.New("error")
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[book.ID] = book
	return nil
}
func (db *DB) GetAll(ctx context.Context) ([]*models.Book, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	books := make([]*models.Book, 0, len(db.data))
	for _, book := range db.data {
		books = append(books, book)
	}
	return books, nil
}
func (db *DB) GetById(ctx context.Context, id int) (*models.Book, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	book, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No book by id:%d", id)
	}
	return book, nil
}
func (db *DB) Update(ctx context.Context, book *models.Book) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[book.ID] = book
	return nil
}
func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, id)
	return nil
}
