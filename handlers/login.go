package handlers

import (
	"TruthOrDare/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNameFromCookie(c echo.Context) string {
	username := "Not logged in"
	cookie, err := c.Cookie("username")
	if err == nil {
		username = cookie.Value
	}
	return username
}

func GetPasswordFromCookie(c echo.Context) string {
	password := "nil"
	cookie, err := c.Cookie("password")
	if err == nil {
		password = cookie.Value
	}
	return password
}

func TryCookiesLogin(c echo.Context) error {
	username := GetNameFromCookie(c)
	password := GetPasswordFromCookie(c)
	if username == "Not logged in" || password == "nil" {
		return c.Redirect(http.StatusFound, "login")
	}

	if repositories.Login(username, password) {
		return c.Redirect(http.StatusFound, "items")
	}

	return Login(c)
}

func RegisterUser(c echo.Context) error {
	name := c.FormValue("username")
	c.SetCookie(&http.Cookie{Name: "username", Value: name})
	password := c.FormValue("password")
	c.SetCookie(&http.Cookie{Name: "password", Value: password})

	if !repositories.DoesUserExist(name) {
		repositories.CreateUser(name, password)
	}

	return c.Redirect(http.StatusFound, "/")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if repositories.Login(username, password) {
		c.SetCookie(&http.Cookie{Name: "username", Value: username})
		c.SetCookie(&http.Cookie{Name: "password", Value: password})
	} else {
		return c.Redirect(http.StatusFound, "login")
	}

	return c.Redirect(http.StatusFound, "items")
}

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "username", Value: "Not logged in"})
	c.SetCookie(&http.Cookie{Name: "password", Value: "nil"})
	return c.Redirect(http.StatusFound, "/")
}
