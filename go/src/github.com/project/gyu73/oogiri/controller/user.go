package controller

import (
	"github.com/gyu73/projects/treasure2017-pre/answers/gyu73/go/treasure2017/yagyu/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"fmt"
	"net/http"
)

type User struct {
	DB *gorm.DB
}

//
//func UserByEmail(db *gorm.DB, email string) (User, error) {
//
//
//}

func (u *User) Show(c echo.Context) error {
	var us model.User
	var oogiri model.Oogiri
	id := c.Param("id")
	user, _ := us.ScanUser(u.DB, id)
	user_oogiris, _ := oogiri.ScanUserOogiri(u.DB, id)
	fmt.Println(user)
	fmt.Println(user_oogiris)
	return c.Render(http.StatusOK, "mypage", map[string]interface{}{
		"user_oogiris": user_oogiris,
		"user":         user},
	)
}

//
//func (u *User) Login(c echo.Context) error {
//	var User model.User
//	User.Name = c.FormValue("name")
//	User.Email = c.FormValue("email")
//	User.Salt = c.FormValue("password")
//	m, err := User.Auth(u.DB,User)
//}
//
//
//func ScanUser(db *sql.DB, email string) (User, error) {
//
//	return ScanUser(db.QueryRow(`select * from users where email = ?`, email))
//}

//
//funcd (u *User)Login (c echo.Context) error {
//	loginForm := LoginForm(
//	UserId: c.FormValue("userId"),
//	Password: c.FormValue("password"),
//	)
//loginForm
//}

//type LoginForm struct {
//	UserId       string
//	Password     string
//	ErrorMessage string
//}
//
//
//func (u *User)ShowLoginHtml(c echo.Context) error {
//session := session.Default(c)

//loginId := session.Get("loginCompleted")
//if loginId != nil && loginId == "completed" {
//	completeJson := CompleteJson{
//		Success: true,
//	}
//	return c.JSON(http.StatusOK, completeJson)
//}

//	return c.Render(http.StatusOK, "login", "")
//}
//
//
//func (u *User)Login(c echo.Context) error {
//	loginForm := LoginForm{
//		UserId:   c.FormValue("userId"),
//		Password: c.FormValue("password"),
//	}
//
//	userId := html.EscapeString(loginForm.UserId)
//	password := html.EscapeString(loginForm.Password)
//
//	if userId != "userId" && password != "password" {
//		loginForm.ErrorMessage = "ユーザーID または パスワードが間違っています"
//		return c.Render(http.StatusOK, "login", loginForm)
//	}

//セッションにデータを保存する
//session := session.Default(c)
//session.Set("loginCompleted", "completed")
//session.Save()
//
//completeJson := CompleteJson{
//	Success: true,
//}
//
//return c.JSON(http.StatusOK, completeJson)
//}
