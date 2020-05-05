package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

func init() {
	var err error
	db, err := gorm.Open("mysql", "root:123456@(localhost)/gintodo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connet databases")
	}
	//	迁移数据库结构
	db.AutoMigrate(&todoMolet{})
}

type todoMolet struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

func (todoMolet) TableName() string {
	return "todo"
}

// createTodo 添加一条数据
func createTodo(c *gin.Context) {
	title := c.PostForm("title")
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoMolet{
		Title:     title,
		Completed: completed,
	}

	db.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created status",
		"todoId":  todo.ID,
	})
}
func fetchAllTodo(c *gin.Context) {

}
func fetchTodo(c *gin.Context) {

}
func updateTodo(c *gin.Context) {

}
func deleteTodo(c *gin.Context) {

}

func main() {
	router := gin.Default()
	todoV1 := router.Group("/api/v1/todos")
	{
		todoV1.POST("/", createTodo)
		todoV1.GET("/", fetchAllTodo)
		todoV1.GET("/:id", fetchTodo)
		todoV1.PUT("/:id", updateTodo)
		todoV1.DELETE("/:id", deleteTodo)
	}
	router.Run()

}
