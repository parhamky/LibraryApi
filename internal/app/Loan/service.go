package Loan

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

func (s *InstanceService) Add(loan *Loan) (uint64, error) {
	dataBase := s.data

	err := dataBase.Create(loan).Error

	if err != nil {
		return 0, err
	}

	return loan.ID, nil
}

func (s *InstanceService) Read(loanId *uint64) (*Loan, error) {
	dataBase := s.data
	cachdb := cache.GetRedisClient()
	key := strconv.FormatUint(*loanId, 10)
	//check if data is exist in cache retrive it from cache
	if cachdb.Exists(key).Val() == 1 {
		res, err := cachdb.Get(key).Result()
		if err != nil {
			return nil, err
		}
		var book Loan
		err = json.Unmarshal([]byte(res), &book)
		if err != nil {
			return nil, err
		}
		return &book, nil
	}
	var loan Loan
	err := dataBase.Where("id = ?", loanId).First(&Loan{}).Take(&loan).Error

	if err != nil {
		return nil, err
	}
	err = cachdb.Set(key, loan, 1*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (s *InstanceService) Update(loan *Loan, id *uint64) error {
	dataBase := s.data
	loan.ID = *id
	err := dataBase.Model(Loan{}).Where("id = ?", *id).Updates(&loan).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *InstanceService) Delete(loanId *uint64) error {
	dataBase := s.data

	err := dataBase.Where("id = ?", loanId).Delete(&Loan{}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetService() *InstanceService {
	return &InstanceService{data: db.GetDB()}
}
