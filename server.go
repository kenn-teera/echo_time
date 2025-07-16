package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Text string `json:"hello"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sex  bool   `json:"sex"`
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
	// Assign a new ID to the user (simple increment for demonstration)
	if len(users) > 0 {
		u.ID = users[len(users)-1].ID + 1
	} else {
		u.ID = 1
	}

	users = append(users, u)
	return c.JSON(http.StatusCreated, u)
}

// deleteUser handles the deletion of a user by ID
func deleteUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Convert string ID to integer

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid user ID"})
	}

	// Find the user and remove them from the slice
	found := false
	for i, user := range users {
		if user.ID == id {
			// Remove the user by slicing the array
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return c.JSON(http.StatusNotFound, Err{Message: "User not found"})
	}

	return c.NoContent(http.StatusNoContent) // Return 204 No Content on successful deletion
}

// updateUser handles the update of a user by ID
func updateUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid user ID"})
	}

	var updatedUser User
	err = c.Bind(&updatedUser) // Bind the request body to the updatedUser struct

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	found := false
	for i, user := range users {
		if user.ID == id {
			// Update the user's fields. We maintain the original ID.
			updatedUser.ID = user.ID // Ensure the ID from the URL is preserved
			users[i] = updatedUser
			found = true
			break
		}
	}

	if !found {
		return c.JSON(http.StatusNotFound, Err{Message: "User not found"})
	}

	return c.JSON(http.StatusOK, updatedUser) // Return the updated user
}

func main() {
	e := echo.New()
	e.GET("/", helloHandler)
	e.GET("/user", getUser)
	e.POST("/user", addUser)
	e.DELETE("/user/:id", deleteUser)
	e.PUT("/user/:id", updateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
