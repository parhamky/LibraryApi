package Book

type Book struct {
	Id          uint64 `gorm:"primaryKey" `
	Title       string `json:"title" `
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	IsAvailable bool   `json:"isavailable"`
}

type Service interface {
	Add(book *Book) (uint64, error)
	Read(bookId *uint64) (*Book, error)
	Update(book *Book, id *uint64) error
	Delete(bookId *uint64) error
}

func (Book) TableName() string {
	return "book"
}
