package main

import (

	// "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//================================================================
//MODELS
//================================================================

//Users Struct
type users struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Post Struct
type posts struct {
	ID        string    `json:"id"`
	Caption   string    `json:"caption"`
	ImageUrl  string    `json:"image"`
	Timestamp time.Time `json:"timestamp"`
}

var u = []users{
	{ID: "1", Name: "Harshit", Email: "harshit0300@gmail.com", Password: "ABCDEF"},
}

//===============================
//MAIN Router and Config
//===============================
func main() {
	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/user/{id}", getUsers)

	router.Run("localhost:8080")
}

//==============================
//Route Functions
//==============================

//HOME
func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, u)
}
