package main

import (
	"html/template"
	"io"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/project/gyu73/oogiri/controller"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Run() {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/oogiri?parseTime=true&collation=utf8_general_ci&interpolateParams=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	oogiri := &controller.Oogiri{DB: db}
	user := &controller.User{DB: db}
	answer := &controller.Answer{DB: db}

	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	e.GET("/", oogiri.Get)
	e.GET("/oogiris", oogiri.Get)
	e.GET("/user/:id", user.Show)
	//e.GET("/login", user.Login)

	//e.GET("/oogiris", oogiri.Get)
	e.GET("/oogiri/new", oogiri.New)
	e.GET("/oogiri/:id", oogiri.Show)
	e.POST("/oogiri", oogiri.Create)
	e.DELETE("/oogiri/:id/delete", oogiri.Delete)

	e.POST("/oogiri/answer", answer.Create)
	e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	Run()
}
