package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/WymA/go-restful/app"
	"github.com/WymA/go-restful/app/models"
)

// Cors func
func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// main func
//
func main() {

	app.InitDB()
	app.AddTables(models.User{}, "user")

	router := gin.Default()

	router.Use(Cors())

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", models.GetUsers)
		v1.GET("/users/:id", models.GetUser)
		v1.POST("/users", models.PostUser)
		v1.PUT("/users/:id", models.UpdateUser)
		v1.DELETE("/users/:id", models.DeleteUser)
		// v1.OPTIONS("/users", models.OptionsUser)     // POST
		// v1.OPTIONS("/users/:id", models.OptionsUser) // PUT, DELETE
	}

	router.Run(":8080")
}
