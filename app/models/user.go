package models

import (
	"log"
	"strconv"
	"time"

	"github.com/WymA/go-restful/app"
	"github.com/gin-gonic/gin"
)

// User struct
type User struct {
	ID        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	UpdateAt  string `db:"updateAt" json:"updateAt"`
	CreateAt  string `db:"createAt" json:"createAt"`
}

// GetUsers func
func GetUsers(c *gin.Context) {

	var users []User
	_, err := app.GetDB().Select(&users, "SELECT * FROM user")

	if err == nil {
		c.JSON(200, users)
	} else {
		log.Printf("XXXX%v", err.Error())
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/users
}

// GetUser func:
func GetUser(c *gin.Context) {

	id := c.Params.ByName("id")
	var user User
	err := app.GetDB().SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)

	if err == nil {

		userID, _ := strconv.ParseInt(id, 0, 64)

		content := &User{
			ID:        userID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {

		c.JSON(404, gin.H{"error": "user not found"})
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

// PostUser func
func PostUser(c *gin.Context) {

	var user User
	c.Bind(&user)

	log.Println(user)

	if user.Firstname != "" && user.Lastname != "" {

		if insert, _ :=
			app.GetDB().Exec(`INSERT INTO user (firstname, lastname, createAt, updateAt)
			VALUES (?, ?, ?, ?)`,
				user.Firstname, user.Lastname, time.Now(), time.Now()); insert != nil {

			userID, err := insert.LastInsertId()
			if err == nil {

				content := &User{
					ID:        userID,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
				}
				c.JSON(201, content)
			} else {
				//checkErr(err, "Insert failed")
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

// UpdateUser func
func UpdateUser(c *gin.Context) {

	id := c.Params.ByName("id")
	var user User
	err := app.GetDB().SelectOne(&user, "SELECT * FROM user WHERE id=?", id)

	if err == nil {
		var json User
		c.Bind(&json)

		userID, _ := strconv.ParseInt(id, 0, 64)

		user := User{
			ID:        userID,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}

		if user.Firstname != "" && user.Lastname != "" {
			_, err = app.GetDB().Update(&user)

			if err == nil {
				c.JSON(200, user)
			} else {
				//app.checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

// DeleteUser func
func DeleteUser(c *gin.Context) {

	id := c.Params.ByName("id")

	var user User
	err := app.GetDB().SelectOne(&user, "SELECT * FROM user WHERE id=?", id)

	if err == nil {
		_, err = app.GetDB().Delete(&user)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			//app.checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}
