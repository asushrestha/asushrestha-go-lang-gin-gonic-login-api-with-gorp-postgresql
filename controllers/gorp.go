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
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "root"
	dbname   = "test_go"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {
	log.Println("test")
	psqlconn := fmt.Sprintf("host= %s port = %d user= %s password = %s dbname= %s sslmode=disable", host, port, username, password, dbname)
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
