package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/naren/og-auth/pkg/config"
	"github.com/naren/og-auth/pkg/utils"
)

var db *gorm.DB

type Reg struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Reg{})
}

func (reg *Reg) CreateUser() *Reg {
	db.NewRecord(reg)
	db.Create(&reg)
	return reg
}

func GetAllUsers() []Reg {
	var Users []Reg
	db.Find(&Users)
	return Users
}

func GetUserByEmail(email string) *Reg {
	var user Reg
	db := db.Where("email = ?", email).Find(&user)
	if db.RecordNotFound() {
		return nil
	}
	if db.Error != nil {
		fmt.Printf("Error finding email")
	}
	return &user
}

func GetUserByEmailLogin(email string, password string) string {
	var user Reg
	db := db.Where("email = ?", email).Find(&user)

	if db.RecordNotFound() {
		return "fail"
	}

	if db.Error != nil {
		fmt.Printf("Error finding email")
		return "fail"
	}

	if !db.RecordNotFound() {
		passwordID := user.Password

		err := utils.CheckPasswordHash(password, passwordID)
		if err != nil {
			return "fail"
		}
		return "success"
	}

	return "fail"
}

func DeleteUser(ID int64) Reg {
	var reg Reg
	db.Where("ID = ?", ID).Delete(&reg)
	return reg
}
