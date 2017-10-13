package controller

import (
	"strconv"

	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/project/gyu73/oogiri/model"
)

type Answer struct {
	DB *gorm.DB
}

func (a *Answer) Create(c echo.Context) error {
	var answer model.Answer
	var err error
	answer.Body = c.FormValue("answer_body")
	oogiri_id := c.FormValue("oogiri_id")
	answer.OogiriId, err = strconv.ParseInt(oogiri_id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	answer.UserId = 1
	answer.CreateAnswer(a.DB)
	return c.Redirect(http.StatusSeeOther, "/oogiris")
	return nil
}
