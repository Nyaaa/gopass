package models

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Item struct {
	Article	string	`json:"article" gorm:"primary_key"`
	Name	string
	IsSet	bool	`gorm:"default:False"`	
}

func GetAllItems(c *gin.Context) {
	var items []Item
	DB.Find(&items)
	c.HTML(http.StatusOK,
		"table.html",
		gin.H{
			"title": "Items",
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

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/items")
	routes.GET("/", GetAllItems)
	routes.DELETE("/:article", DeleteItem)
}
