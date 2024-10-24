package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Text string `json:"hello"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sex  bool   `json:"sex 0:M 1:FM"`
}

type Err struct {
	Message string `json:"err"`
}

var users = []User{
	{ID: 1, Name: "Kenn", Sex: false},
	{ID: 2, Name: "Noon", Sex: true},
}

var messages = []Message{
	{Text: "Hello World"},
}

func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, messages)
}

func addUser(c echo.Context) error {
	var u User
	err := c.Bind(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	users = append(users, u)
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	e.GET("/", helloHandler)
	e.GET("/user", getUser)
	e.POST("/user", addUser)
	e.Logger.Fatal(e.Start(":1323"))
}
