package entities

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Id uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	//Id uint64 `gorm:"primaryKey"`
	Firstname string `gorm:"not null;type:string;size:50;column:firstname1"`
	Lastname  string `gorm:"not null;type:string;size:50;column:lastname1"`
	Email     string `gorm:"unique;not null;type:string;size:50"`
	Age       uint8  `gorm:"type:smallserial"`
	Gender    string `gorm:"type:varchar(10);default:'gender1'"` //varchar(10) => "type:string;size(10)"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

type Tabler interface {
	TableName() string
}

func (Customer) TableName() string { //HOOK
	return "customers1"
}

func (customerEntity *Customer) BeforeCreate(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("BeforeCreateHook %s", serializedCustomer)
	return nil
}
func (customerEntity *Customer) BeforeSave(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("BeforeSaveHook %s", serializedCustomer)
	return nil
}
func (customerEntity *Customer) AfterCreate(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("AfterCreateHook %s", serializedCustomer)
	return nil
}
func (customerEntity *Customer) AfterSave(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("AfterSaveHook %s", serializedCustomer)
	return nil
}

func NewCustomerEntity(firstname string, lastname string, email string, age uint8, gender string) (Customer, error) {
	if age < 1 {
		return Customer{}, errors.New("age value can't be smaller than 1")
	} else if age > 125 {
		return Customer{}, errors.New("age value can't be greater than 125")
	}
	return Customer{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Age:       age,
		Gender:    gender,
	}, nil
}
