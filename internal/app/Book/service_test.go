package Book

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
	book := Book{
		Title:       "test book",
		Author:      "test author",
		ISBN:        "1245678-1288397",
		IsAvailable: true,
	}
	tx := dataBase.Begin()
	id, err := serv.Add(&book)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}
	var res Book
	err = dataBase.First(&res, id).Error
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}

	assert.NoError(t, tx.Commit().Error)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, res.Title)
	assert.Equal(t, book.Author, res.Author)
	assert.Equal(t, book.ISBN, res.ISBN)
	assert.Equal(t, book.IsAvailable, res.IsAvailable)
}

func TestInstanceService_Read(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	book := Book{
		Title:       "test book2",
		Author:      "test author2",
		ISBN:        "1245678-1248347",
		IsAvailable: true,
	}
	err := dataBase.Create(&book).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := serv.Read(&book.Id)
	log.Println(res)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, book.Title, res.Title)
	assert.Equal(t, book.Author, res.Author)
	assert.Equal(t, book.ISBN, res.ISBN)
	assert.Equal(t, book.IsAvailable, res.IsAvailable)
}

func TestInstanceService_Update(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	book := Book{
		Title:       "test book3",
		Author:      "test author3",
		ISBN:        "1245638-3248347",
		IsAvailable: true,
	}
	err := dataBase.Create(&book).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	updatedBook := Book{
		Title:       "test book4",
		Author:      "test author4",
		ISBN:        "1245674-1444347",
		IsAvailable: true,
	}
	err = serv.Update(&updatedBook, &book.Id)
	if err != nil {
		log.Fatal(err.Error())
	}
	var res Book
	dataBase.First(&res, book.Id)
	assert.Equal(t, updatedBook.Title, res.Title)
	assert.Equal(t, updatedBook.Author, res.Author)
	assert.Equal(t, updatedBook.ISBN, res.ISBN)
	assert.Equal(t, updatedBook.IsAvailable, res.IsAvailable)
}

func TestInstanceService_Delete(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	book := Book{
		Title:       "test book5",
		Author:      "test author5",
		ISBN:        "1245538-3548357",
		IsAvailable: true,
	}
	err := dataBase.Create(&book).Error

	err = serv.Delete(&book.Id)

	var res Book
	err = dataBase.First(&res, book.Id).Error
	assert.NotNil(t, err)
}
