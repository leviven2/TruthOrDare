package main

import (
	"TruthOrDare/database"
	"TruthOrDare/handlers"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	e.GET("/", handlers.TryCookiesLogin)
	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	})
	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.RegisterUser)
	e.POST("/logout", handlers.Logout)
	e.GET("/items", handlers.ShowHomePage)
	e.POST("/category/add", handlers.AddCategory)
	e.POST("/category/:id/post", handlers.GambleCategory)
	e.POST("/category/:id/edit", handlers.EditCategory)
	e.POST("/category/:id/additem", handlers.AddItem)
	e.POST("/items/deleteitem", func(c echo.Context) error {
		id := c.QueryParam("id")
		return handlers.DeleteItem(c, id)
	})
	e.POST("/items/delete", func(c echo.Context) error {
		listname := c.QueryParam("name")
		return handlers.DeleteCategory(c, listname)
	})

	e.Debug = true

	database.DB()

	log.Fatal(e.Start(":1323"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
