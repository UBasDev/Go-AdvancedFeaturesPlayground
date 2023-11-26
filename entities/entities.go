package entities

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleEntity struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Key         string    `gorm:"unique;not null;type:string;size:50;column:key1"`
	Code        string    `gorm:"unique;not null;type:string;size:50;column:code1"`
	Value       uint8     `gorm:"unique;not null;type:smallserial;column:value1"`
	Description string    `gorm:"type:string;size:100;column:description1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt    `gorm:"index"`
	Customers1  []*CustomerEntity `gorm:"foreignKey:RoleId1;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Screens1    []*ScreenEntity   `gorm:"many2many:screens1_roles1;foreignKey:Id;joinForeignKey:RoleId1;References:Id;joinReferences:ScreenId1"`
}

type ScreenEntity struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Key         string    `gorm:"unique;not null;type:string;size:50;column:key1"`
	Value       string    `gorm:"unique;not null;type:string;size:50;column:value1"`
	Description string    `gorm:"type:string;size:100;column:description1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt    `gorm:"index"`
	Customers1  []*CustomerEntity `gorm:"many2many:customers1_screens1;foreignKey:Id;joinForeignKey:ScreenId1;References:Id;joinReferences:CustomerId1"`
	Roles1      []*RoleEntity     `gorm:"many2many:screens1_roles1;foreignKey:Id;joinForeignKey:ScreenId1;References:Id;joinReferences:RoleId1"`
}

type CustomerEntity struct {
	gorm.Model
	Id uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	//Id uint64 `gorm:"primaryKey"`
	Firstname      string `gorm:"not null;type:string;size:50;column:firstname1"`
	Lastname       string `gorm:"not null;type:string;size:50;column:lastname1"`
	Email          string `gorm:"unique;not null;type:string;size:50"`
	Age            uint8  `gorm:"type:smallint"`
	Gender         string `gorm:"type:varchar(10);default:'gender1'"` //varchar(10) => "type:string;size(10)"
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt  `gorm:"index"`
	RoleId1        *uuid.UUID      `gorm:"type:uuid"`
	Screens1       []*ScreenEntity `gorm:"many2many:customers1_screens1;foreignKey:Id;joinForeignKey:CustomerId1;References:Id;joinReferences:ScreenId1"`
	ProfileEntity1 *ProfileEntity  `gorm:"foreignKey:CustomerId1;references:Id"`
}

type ProfileEntity struct {
	gorm.Model
	Id                 uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	TokenCount         uint32    `gorm:"type:integer"`
	BalanceIntegerPart uint32    `gorm:"type:integer"`
	BalanceDecimalPart uint32    `gorm:"type:integer"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	CustomerId1        uuid.UUID      `gorm:"type:uuid"`
}

type Tabler interface {
	TableName() string
}

func (ProfileEntity) TableName() string {
	return "profiles1"
}

func (RoleEntity) TableName() string {
	return "roles1"
}
func NewRoleEntity(key string, code string, value uint8, description string) (RoleEntity, error) {
	if value < 1 {
		return RoleEntity{}, errors.New("role value can't be smaller than 1")
	}
	return RoleEntity{
		Key:         key,
		Code:        code,
		Value:       value,
		Description: description,
	}, nil
}

func (ScreenEntity) TableName() string {
	return "screens1"
}
func NewScreenEntity(key string, value string, description string) (ScreenEntity, error) {
	return ScreenEntity{
		Key:         key,
		Value:       value,
		Description: description,
	}, nil
}

func (CustomerEntity) TableName() string { //HOOK
	return "customers1"
}

func (customerEntity *CustomerEntity) BeforeCreate(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("BeforeCreateHook %s", serializedCustomer)
	return nil
}
func (customerEntity *CustomerEntity) AfterCreate(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("AfterCreateHook %s", serializedCustomer)
	return nil
}
func (customerEntity *CustomerEntity) BeforeUpdate(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("BeforeUpdateHook %s", serializedCustomer)
	return nil
}
func (customerEntity *CustomerEntity) AfterUpdate(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("AfterUpdateHook %s", serializedCustomer)
	return nil
}
func (customerEntity *CustomerEntity) BeforeSave(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("BeforeSaveHook %s", serializedCustomer)
	return nil
}

func (customerEntity *CustomerEntity) AfterSave(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("AfterSaveHook %s", serializedCustomer)
	return nil
}

func (customerEntity *CustomerEntity) BeforeDelete(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("BeforeDeleteHook %s", serializedCustomer)
	return nil
}
func (customerEntity *CustomerEntity) AfterDelete(dbContext *gorm.DB) (err error) {
	serializedCustomer, err := json.Marshal(customerEntity)
	if err != nil {
		return err
	}
	log.Printf("AfterDeleteHook %s", serializedCustomer)
	return nil
}

func NewCustomerEntity(firstname string, lastname string, email string, age uint8, gender string) (CustomerEntity, error) {
	if age < 1 {
		return CustomerEntity{}, errors.New("age value can't be smaller than 1")
	} else if age > 125 {
		return CustomerEntity{}, errors.New("age value can't be greater than 125")
	}
	return CustomerEntity{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Age:       age,
		Gender:    gender,
	}, nil
}

func NewProfileEntity(token_count uint32, balance_integer_part uint32, balance_decimal_part uint32) (ProfileEntity, error) {
	if token_count < 1 {
		return ProfileEntity{}, errors.New("token count can't be smaller than 1")
	} else if balance_integer_part < 1 {
		return ProfileEntity{}, errors.New("balance integer value can't be smaller than 1")
	} else if balance_decimal_part < 1 {
		return ProfileEntity{}, errors.New("balance decimal value can't be smaller than 1")
	}
	return ProfileEntity{
		TokenCount:         token_count,
		BalanceIntegerPart: balance_integer_part,
		BalanceDecimalPart: balance_decimal_part,
	}, nil
}
