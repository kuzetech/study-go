package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

func main() {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, users)
	})

	router.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		found := false
		//类似于数据库的SQL查询
		for _, u := range users {
			if strings.EqualFold(id, strconv.Itoa(u.ID)) {
				user = u
				found = true
				break
			}
		}
		if found {
			c.JSON(200, user)
		} else {
			c.JSON(404, gin.H{
				"message": "用户不存在",
			})
		}
	})

	router.POST("/user", func(c *gin.Context) {
		name := c.DefaultPostForm("name", "")
		if name != "" {
			u := User{ID: len(users) + 1, Name: name}
			users = append(users, u)
			c.JSON(http.StatusCreated, u)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "请输入用户名称",
			})
		}
	})

	router.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		i := -1

		for index, u := range users {
			if strings.EqualFold(id, strconv.Itoa(u.ID)) {
				i = index
				break
			}
		}
		if i >= 0 {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusNoContent, "")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "用户不存在",
			})
		}
	})

	router.Run(":8080")
}
