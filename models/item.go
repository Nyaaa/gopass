package models

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Article string `form:"article" gorm:"unique;not null;primary_key"`
	Name    string `form:"name"  gorm:"not null; default:null"`
	IsSet   bool   `form:"isSet" gorm:"default:False"`
}

func GetAllItems(c *gin.Context) {
	var items []Item
	DB.Find(&items)
	c.HTML(http.StatusOK,
		"table.html",
		gin.H{
			"title":   "Items",
			"payload": items,
			"columns": reflect.VisibleFields(reflect.TypeOf(Item{})),
		},
	)
}

func DeleteItem(c *gin.Context) {
	article := c.Param("article")
	var item Item

	if result := DB.First(&item, "article = ?", article); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	DB.Delete(&item)
	c.Status(http.StatusOK)
}

func CreateItem(c *gin.Context) {
	var newItem Item

	if err := c.Bind(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	GetAllItems(c)
}

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/items")
	routes.GET("/", GetAllItems)
	routes.DELETE("/:article", DeleteItem)
	routes.POST("/", CreateItem)
}
