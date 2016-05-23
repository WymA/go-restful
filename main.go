package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// User struct
type User struct {
	ID        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

func main() {

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", PostUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	router.Run(":8080")
}

// GetUsers func
func GetUsers(c *gin.Context) {

	type Users []User
	var users = Users{
		User{ID: 1, Firstname: "Oliver", Lastname: "Queen"},
		User{ID: 2, Firstname: "Malcom", Lastname: "Merlyn"},
	}
	c.JSON(200, users)
	// curl -i http://localhost:8080/api/v1/users
}

// GetUser func
func GetUser(c *gin.Context) {

	id := c.Params.ByName("id")
	userID, _ := strconv.ParseInt(id, 0, 64)

	if userID == 1 {

		content := gin.H{"id": userID, "firstname": "Oliver", "lastname": "Queen"}
		c.JSON(200, content)
	} else if userID == 2 {

		content := gin.H{"id": userID, "firstname": "Malcom", "lastname": "Merlyn"}
		c.JSON(200, content)
	} else {

		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

// PostUser func
func PostUser(c *gin.Context) {
	// The futur code…
}

// UpdateUser func
func UpdateUser(c *gin.Context) {
	// The futur code…
}

// DeleteUser func
func DeleteUser(c *gin.Context) {
	// The futur code…
}
