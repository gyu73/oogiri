package controller

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/ipfans/echo-session"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/project/gyu73/oogiri/model"
)

type Oogiri struct {
	DB *gorm.DB
}

func (o *Oogiri) Get(c echo.Context) error {

	// "user_information"という名前のcookieのValueを取得
	_, err := c.Cookie("user_information")
	if err != nil {
		c.Logger().Print(err)
		return c.Render(http.StatusOK, "new", "")
	}

	// sessions.Get

	var store = sessions.NewCookieStore([]byte("something-very-secret"))
	//user_informationという名前のクッキーからsessionを取得
	sess, err := store.Get(c.Request(), "user_information")

	sess.Values["id"] = "1"
	sess.Values["password"] = "pass"
	sess.Save(c.Request(), c.Response())

	session, nil = session.Get("user_information", c)

	var oogiri model.Oogiri
	oogiris, err := oogiri.ScanAllOogiris(o.DB)
	if err != nil {
		log.Fatal(err)
	}
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"oogiris": oogiris},
	)
}

func (o *Oogiri) Show(c echo.Context) error {
	//store, err := c.Cookie("user_information")
	//sess, err := store.Get(c.Request(), "user_information")
	var answer model.Answer
	var oogiri model.Oogiri
	id := c.Param("id")
	current_oogiri, err := oogiri.ScanOneOogiri(o.DB, id)
	if err != nil {
		log.Fatal(err)
	}
	current_oogiri_answers, err := answer.ScanAnswers(o.DB, id)
	return c.Render(http.StatusOK, "show", map[string]interface{}{
		"oogiri":  current_oogiri,
		"answers": current_oogiri_answers},
	)
}

func (o *Oogiri) New(c echo.Context) error {
	return c.Render(http.StatusOK, "new", "")
}

func (o *Oogiri) Create(c echo.Context) error {
	var oogiri model.Oogiri
	oogiri.Title = c.FormValue("title")
	oogiri.Body = c.FormValue("body")
	oogiri.Image = c.FormValue("image")
	oogiri.UserId = 1
	oogiri.CreateOogiri(o.DB)
	return c.Redirect(http.StatusSeeOther, "/oogiris")
	return nil
}

func (o *Oogiri) Delete(c echo.Context) error {
	var oogiri model.Oogiri
	id := c.Param("id")
	oogiri.DeleteOogiri(o.DB, id)
	return c.Redirect(http.StatusSeeOther, "/oogiris")
	return nil
}
