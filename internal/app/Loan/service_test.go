package Loan

import (
	"LibraryApi/internal/config"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
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
	loan := Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	tx := dataBase.Begin()
	id, err := serv.Add(&loan)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}
	var res Loan
	err = dataBase.First(&res, id).Error
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}

	assert.NoError(t, tx.Commit().Error)
	assert.NoError(t, err)
	assert.Equal(t, loan.UserID, res.UserID)
	assert.Equal(t, loan.BookID, res.BookID)
}

func TestInstanceService_Read(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	loan := Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	err := dataBase.Create(&loan).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := serv.Read(&loan.ID)
	log.Println(res)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, loan.UserID, res.UserID)
	assert.Equal(t, loan.BookID, res.BookID)
}

func TestInstanceService_Update(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	loan := Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	err := dataBase.Create(&loan).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	updatedLoan := Loan{
		UserID:     1,
		BookID:     3,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(48 * time.Hour),
		ReturnDate: time.Time{},
	}
	err = serv.Update(&updatedLoan, &loan.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	var res Loan
	dataBase.First(&res, loan.ID)
	assert.Equal(t, updatedLoan.UserID, res.UserID)
	assert.Equal(t, updatedLoan.BookID, res.BookID)
}

func TestInstanceService_Delete(t *testing.T) {
	serv := setup()
	conf := config.LoadDBConfig()
	dataBase := connectDB(&conf)
	serv.data = dataBase
	loan := Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	err := dataBase.Create(&loan).Error

	err = serv.Delete(&loan.ID)

	var res Loan
	err = dataBase.First(&res, loan.ID).Error
	assert.NotNil(t, err)
}
