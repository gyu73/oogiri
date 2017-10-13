package model

import (
	"github.com/jinzhu/gorm"
)

func (a *Answer) CreateAnswer(db *gorm.DB) {
	db.Create(&a)
}

func (a *Answer) ScanAnswers(db *gorm.DB, id string) ([]Answer, error) {
	answers := make([]Answer, 0, 16)
	db.Where("oogiri_id = ?", id).Find(&answers)
	return answers, nil
}
