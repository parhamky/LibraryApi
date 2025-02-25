package Book

import (
	"errors"
	"log"
)

type MockService struct {
	Books map[uint64]*Book
}

func (m *MockService) Add(book *Book) (uint64, error) {
	newID := uint64(len(m.Books) + 1)
	book.Id = newID
	m.Books[newID] = book
	log.Println("++++++++++++++++++++++++++log mock+++++++++++++++++++++++++++++++")
	log.Println(m.Books[newID])
	log.Println("++++++++++++++++++++++++++end++++++++++++++++++++++++++++++++++++")
	return newID, nil
}

func (m *MockService) Read(id *uint64) (*Book, error) {
	if book, ok := m.Books[*id]; ok {
		return book, nil
	}
	return nil, errors.New("Book not found")
}

func (m *MockService) Update(book *Book, id *uint64) error {
	if _, ok := m.Books[*id]; ok {
		book.Id = *id
		m.Books[*id] = book
		return nil
	} else {
		return errors.New("Book not found")
	}
}

func (m *MockService) Delete(bookid *uint64) error {
	if _, ok := m.Books[*bookid]; ok {
		delete(m.Books, *bookid)
		return nil
	} else {
		return errors.New("Book not found")
	}
}

func GetMockService() *MockService {
	return &MockService{Books: make(map[uint64]*Book)}
}
