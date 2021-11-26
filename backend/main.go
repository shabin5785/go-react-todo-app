package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shabin5785/go-react-todo/api"
)

func indexView(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "todo app"})
}

func setUpRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", indexView)

	router.GET("/items", api.GetAllTodoItems)
	router.POST("/createItem", api.CreateTodoItem)
	router.POST("/updateItem", api.UpdateTodoItem)
	router.GET("/deleteItem/:id", api.DeleteTodoItem)

	return router

}

func main() {

	api.SetupPostgres()
	router := setUpRoutes()
	router.Run(":8080")

}
