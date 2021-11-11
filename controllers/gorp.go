package controllers

import (
	"database/sql"
	"fmt"
	"gin-login/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres://isuexdaqctjrhz:b260e378c4e6549a2e74e9d13bc947b99c0b856d8803fa4aef60b9fcb5fb36b5@ec2-107-20-127-127.compute-1.amazonaws.com:5432/dcgs73i6lmq0nc"
	port     = 5432
	username = "isuexdaqctjrhz"
	password = "b260e378c4e6549a2e74e9d13bc947b99c0b856d8803fa4aef60b9fcb5fb36b5"
	dbname   = "dcgs73i6lmq0nc"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {
	log.Println("test")
	psqlconn := fmt.Sprintf("host= %s port = %d user= %s password = %s dbname= %s sslmode=enable", host, port, username, password, dbname)
	log.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	log.Println(db)
	checkErr(err, "sql.Open failed")
	dialect := gorp.PostgresDialect{}
	dbmap := &gorp.DbMap{Db: db, Dialect: dialect}

	log.Println(err, "asd")
	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "id")
	err = dbmap.CreateTablesIfNotExists()

	checkErr(err, "Create tables failed")
	log.Println(dbmap)
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
