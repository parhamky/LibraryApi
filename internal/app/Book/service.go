package Book

import (
	"LibraryApi/internal/cache"
	"LibraryApi/internal/db"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type InstanceService struct {
	data *gorm.DB
}

func (s *InstanceService) Add(book *Book) (uint64, error) {
	dataBase := s.data

	err := dataBase.Create(book).Error

	if err != nil {
		return 0, err
	}

	return book.Id, nil
}

func (s *InstanceService) Read(bookId *uint64) (*Book, error) {
	dataBase := s.data
	cachdb := cache.GetRedisClient()
	key := strconv.FormatUint(*bookId, 10)
	//check if data is exist in cache retrive it from cache
	if cachdb.Exists(key).Val() == 1 {
		res, err := cachdb.Get(key).Result()
		if err != nil {
			return nil, err
		}
		var book Book
		err = json.Unmarshal([]byte(res), &book)
		if err != nil {
			return nil, err
		}
		return &book, nil
	}

	var book Book
	err := dataBase.Where("id = ?", bookId).First(&Book{}).Take(&book).Error
	if err != nil {
		return nil, err
	}
	//create a record of data on cache
	err = cachdb.Set(key, book, 1*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *InstanceService) Update(book *Book, id *uint64) error {
	dataBase := s.data
	book.Id = *id
	err := dataBase.Model(Book{}).Where("id = ?", *id).Updates(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *InstanceService) Delete(bookId *uint64) error {
	dataBase := s.data

	err := dataBase.Where("id = ?", bookId).Delete(&Book{}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetService() *InstanceService {
	return &InstanceService{data: db.GetDB()}
}
