package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "lok1415161213"
	DB_NAME     = "gobase1"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//ROUTES SETUP
	initializeRoutes(router)
	err1 := router.Run()
	if err1 != nil {
		return
	}
}

func index(conn *gin.Context) {
	conn.HTML(http.StatusOK, "index.html", gin.H{"title": "Home page"})
}

type UserInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func dbConn() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	checkErr(err)
	return db
}

func getUsers(c *gin.Context) {
	db := dbConn()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var users []UserInfo
	rows, err := db.Query("select * from userinfo")
	if err != nil {
		return
	}

	for rows.Next() {
		var user UserInfo
		err := rows.Scan(&user.ID, &user.Name, &user.Position)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	db := dbConn()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	var newUsers []UserInfo

	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(x, &newUsers)
	if err != nil {
		return
	}
	fmt.Println(newUsers)

	err = c.BindJSON(&newUsers)
	if err != nil {
		return
	}

	for _, val := range newUsers {
		_, err := db.Exec("insert into userInfo(id, name, position) values ($1, $2, $3)", val.ID, val.Name, val.Position)
		if err != nil {
			return
		}
	}
	c.IndentedJSON(http.StatusCreated, newUsers)
}
