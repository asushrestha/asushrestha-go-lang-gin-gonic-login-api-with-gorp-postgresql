package controllers

import (
	"fmt"
	"gin-login/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from users")

	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func GetUserDetail(c *gin.Context) {

	id := c.Params.ByName("id")
	var user models.User
	parsed_id, err1 := strconv.ParseInt(id, 10, 64)
	if err1 != nil {
		c.JSON(400, gin.H{"message": "Not a valid Id"})
	}
	sqlQuery := fmt.Sprintf("SELECT * FROM users WHERE id= %d", parsed_id)
	err := dbmap.SelectOne(&user, sqlQuery)
	log.Println(err)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)

		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := dbmap.SelectOne(&user, fmt.Sprintf("select * from users where username='%s'", user.Username))

	if err == nil {
		user_id := user.Id

		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func PostUser(c *gin.Context) {
	var user models.User
	var users []models.User
	c.Bind(&user)

	log.Println(user)

	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
		// log.Printf("here")
		sqlQuery0 := fmt.Sprintf(`SELECT * from users where username='%s'`, user.Username)
		log.Println(sqlQuery0)
		checkQuery, _ := dbmap.Select(&users, sqlQuery0)
		log.Println(checkQuery)
		if checkQuery != nil {
			if len(users) > 0 {
				c.Abort()
				c.JSON(400, gin.H{"error": "username already exists"})
				return
			}

		}
		sqlQuery := fmt.Sprintf(`INSERT INTO users (username, user_password, firstname, lastname) VALUES ('%s', '%s', '%s', '%s')`, user.Username, user.Password, user.Firstname, user.Lastname)
		log.Println(sqlQuery)
		// insert, error := dbmap.Exec(sqlQuery)
		// log.Println(insert)
		// log.Println(error)
		if insert, _ := dbmap.Exec(sqlQuery); insert != nil {
			// log.Printf("here1")
			// user_id, err := insert.LastInsertId()
			// if err == nil {

			content := &models.User{
				Id:        user.Id,
				Username:  user.Username,
				Password:  user.Password,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}
			c.JSON(201, content)
		} else {
			c.JSON(500, gin.H{"error": "GIN DB insertion failed"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	parsed_id, err1 := strconv.ParseInt(id, 0, 64)
	if err1 != nil {
		c.JSON(400, gin.H{"message": err1})
	}
	var user models.User
	sqlQuery := fmt.Sprintf("SELECT * FROM users WHERE id=%d", parsed_id)
	log.Println(sqlQuery)
	err := dbmap.SelectOne(&user, sqlQuery)
	log.Println(err)
	if err == nil {
		var json models.User
		c.Bind(&json)

		user_id := parsed_id

		user := models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}

		if user.Firstname != "" && user.Lastname != "" {
			_, err = dbmap.Update(&user)

			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
