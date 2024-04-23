package repositories

import (
	"TruthOrDare/database"
	"TruthOrDare/models"
	"math/rand"
)

func GetCategoriesForUser(name string) []models.Category {
	var user models.User
	database.DB().Where("name = ?", name).First(&user)

	var categories []models.Category
	database.DB().Model(&user).Association("Categories").Find(&categories)

	return categories
}

func GetCategoryById(id string) models.Category {
	var category models.Category
	database.DB().Where("id = ?", id).First(&category)
	return category
}

func GetItemsForCategory(category models.Category) []models.Item {
	var items []models.Item
	database.DB().Model(&category).Association("Items").Find(&items)
	category.Items = items
	return items
}

func GetRandomItemForCategory(categoryId string) models.Item {
	var item models.Item

	var totalItems int64
	database.DB().Where("category_id = ?", categoryId).Model(&models.Item{}).Count(&totalItems)

	randomIndex := rand.Intn(int(totalItems))
	database.DB().Where("category_id = ?", categoryId).Offset(randomIndex).Limit(1).Find(&item)

	return item
}

func AddCategoryForUser(userName string, categoryName string) {
	var user models.User
	database.DB().Where("name = ?", userName).Preload("Categories").First(&user)

	category := models.Category{Name: categoryName, UserID: user.ID}
	user.Categories = append(user.Categories, category)

	database.DB().Save(&category)
	database.DB().Save(&user)
}

func DeleteCategoryForUser(username string, categoryName string) {
	var user models.User
	database.DB().Where("name = ?", username).Preload("Categories").First(&user)

	var todoList models.Category
	database.DB().Where("name = ?", categoryName).First(&todoList)

	database.DB().Model(&user).Association("Categories").Delete(&todoList)
	database.DB().Delete(&todoList)
}

func AddItemForUser(categoryId string, itemName string) {
	var category models.Category
	database.DB().Where("id = ?", categoryId).Preload("Items").First(&category)

	item := models.Item{Description: itemName, CategoryID: category.ID}
	category.Items = append(category.Items, item)

	database.DB().Save(&category)
}

func DeleteItem(itemId string) {
	var item models.Item
	database.DB().Where("id = ?", itemId).First(&item)

	database.DB().Delete(&item)
}
