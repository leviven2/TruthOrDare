package handlers

import (
	"TruthOrDare/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ShowHomePage(c echo.Context) error {
	username := GetNameFromCookie(c)

	categories := repositories.GetCategoriesForUser(username)

	for i := range categories {
		categories[i].Items = repositories.GetItemsForCategory(categories[i])
	}

	data := echo.Map{
		"Name":       username,
		"Categories": categories,
	}

	return c.Render(http.StatusOK, "items", data)
}

func AddCategory(c echo.Context) error {
	username := GetNameFromCookie(c)
	if !repositories.DoesUserExist(username) {
		return c.NoContent(http.StatusNotFound)
	}

	repositories.AddCategoryForUser(username, c.FormValue("name"))
	return c.Redirect(http.StatusFound, "/items")
}

func DeleteCategory(c echo.Context, categoryName string) error {
	username := GetNameFromCookie(c)
	if !repositories.DoesUserExist(username) {
		return c.NoContent(http.StatusNotFound)
	}

	repositories.DeleteCategoryForUser(username, categoryName)
	return c.Redirect(http.StatusFound, "/items")
}

func AddItem(c echo.Context) error {
	username := GetNameFromCookie(c)
	if !repositories.DoesUserExist(username) {
		return c.NoContent(http.StatusNotFound)
	}

	repositories.AddItemForUser(c.Param("id"), c.FormValue("name"))
	return EditCategory(c)
}

func DeleteItem(c echo.Context, id string) error {
	repositories.DeleteItem(id)
	return c.Redirect(http.StatusFound, "/items")
}

func GambleCategory(c echo.Context) error {
	id := c.Param("id")
	item := repositories.GetRandomItemForCategory(id)

	data := echo.Map{
		"Item":         item,
		"CategoryName": repositories.GetCategoryById(id).Name,
	}

	return c.Render(http.StatusOK, "gamble", data)
}

func EditCategory(c echo.Context) error {
	id := c.Param("id")
	category := repositories.GetCategoryById(id)
	category.Items = repositories.GetItemsForCategory(category)

	data := echo.Map{
		"Category": category,
	}

	return c.Render(http.StatusOK, "editcategory", data)
}
