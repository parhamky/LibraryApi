package User

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

func (s *InstanceService) Add(user *User) (uint64, error) {
	dataBase := s.data

	err := dataBase.Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *InstanceService) Read(userId *uint64) (*User, error) {
	dataBase := s.data
	cachdb := cache.GetRedisClient()
	key := strconv.FormatUint(*userId, 10)
	//check if data is exist in cache retrive it from cache
	if cachdb.Exists(key).Val() == 1 {
		res, err := cachdb.Get(key).Result()
		if err != nil {
			return nil, err
		}
		var book User
		err = json.Unmarshal([]byte(res), &book)
		if err != nil {
			return nil, err
		}
		return &book, nil
	}
	var user User
	err := dataBase.Where("id = ?", userId).First(&User{}).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = cachdb.Set(key, user, 1*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *InstanceService) Update(user *User, id *uint64) error {
	dataBase := s.data
	user.ID = *id
	err := dataBase.Model(User{}).Where("id = ?", *id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *InstanceService) Delete(userId *uint64) error {
	dataBase := s.data

	err := dataBase.Where("id = ?", userId).Delete(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetService() *InstanceService {
	return &InstanceService{data: db.GetDB()}
}
