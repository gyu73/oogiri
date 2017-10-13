package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func (o *Oogiri) ScanAllOogiris(db *gorm.DB) ([]Oogiri, error) {
	oogiris := make([]Oogiri, 0, 16)
	db.Find(&oogiris)
	return oogiris, nil
}

func (o *Oogiri) ScanOneOogiri(db *gorm.DB, id string) ([]Oogiri, error) {
	oogiri := make([]Oogiri, 0, 16)
	db.Where("id = ?", id).Find(&oogiri)
	return oogiri, nil
}

func (o *Oogiri) ScanUserOogiri(db *gorm.DB, id string) ([]Oogiri, error) {
	var user User
	var oogiri []Oogiri
	db.Where("id = ?", id).Find(&user)
	db.Model(&user).Related(&user.Oogiris)
	oogiri = user.Oogiris
	return oogiri, nil
}

func (o *Oogiri) CreateOogiri(db *gorm.DB) {
	db.Create(&o)
}

func (o *Oogiri) DeleteOogiri(db *gorm.DB, id string) {
	var err error
	o.ID, err  = strconv.Atoi(id)
	db.Find(&o)
	db.Model(&o).Related(&o.Answers)
	if err != nil{
		log.Fatal(err)
	}
	db.Delete(&o)
}

//func (o *Oogiri) GetUserName(db *gorm.DB) (*User.Name, error) {
//	db.Find
//}
