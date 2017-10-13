package model

import "github.com/jinzhu/gorm"

func (o *User) ScanUser(db *gorm.DB, id string) (User, error) {
	var user User
	db.Where("id = ?", id).Find(&user)
	return user, nil
}

//
//func (u *User) Auth(db *gorm.DB, User User) (User ,error) {
//	User.GetUSer(db)
//	return u, nil
//}
//
//func (u *User) GetUser(db *gorm.DB)(User, error){
//	user := db.Find(&u)
//	if user != nil{
//		return nil, error
//	}
//}
