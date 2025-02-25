package User

import (
	"LibraryApi/internal/config"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func connectDB(config *config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		config.User,
		config.Password,
		"postgres",
		"5432",
	)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDB
}

func setup() *InstanceService {
	var serv = GetService()
	err := godotenv.Load("../../../.env")
	if err != nil {
		msg := err.Error()
		log.Fatalf("Error loading .env file:%s", msg)
	}
	return serv
}

func TestInstanceService_Add(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	log.Println(conf)
	dataBase := connectDB(&conf)
	serv.data = dataBase
	user := User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
	}
	tx := dataBase.Begin()
	id, err := serv.Add(&user)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}
	var res User
	err = dataBase.First(&res, id).Error
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}

	assert.NoError(t, tx.Commit().Error)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, res.Name)
	assert.Equal(t, user.Email, res.Email)
}

func TestInstanceService_Read(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	user := User{
		Name:     "test2",
		Email:    "test@test2.com",
		Password: "test2",
		Role:     "admin2",
	}
	err := dataBase.Create(&user).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := serv.Read(&user.ID)
	log.Println(res)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, user.Name, res.Name)
	assert.Equal(t, user.Email, res.Email)
}

func TestInstanceService_Update(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	user := User{
		Name:     "test3",
		Email:    "test3@test2.com",
		Password: "test3",
		Role:     "admin3",
	}
	err := dataBase.Create(&user).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	updatedUser := User{
		Name:     "testUpdated4",
		Email:    "testUpdated@test.com",
		Password: "test",
		Role:     "user",
	}
	err = serv.Update(&updatedUser, &user.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	var res User
	dataBase.First(&res, user.ID)
	assert.Equal(t, updatedUser.Name, res.Name)
	assert.Equal(t, updatedUser.Email, res.Email)
}

func TestInstanceService_Delete(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	user := User{
		Name:     "test5",
		Email:    "test6@test.com",
		Password: "test5",
		Role:     "admin5",
	}
	err := dataBase.Create(&user).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	err = serv.Delete(&user.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	var res User
	err = dataBase.First(&res, user.ID).Error
	assert.NotNil(t, err)
}
