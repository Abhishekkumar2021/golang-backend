package controllers

import (
	"io"
	"net/http"
	"time"

	"github.com/Abhishekkumar2021/golang-backend/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/loremipsum.v1"
)

// Healthcheck controllers
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Health check passed",
	})
}


// User controllers
func AddUser(c *gin.Context) {
	// AddUser logic
	user := models.User{}
	c.BindJSON(&user)
	result, err := models.AddUser(user)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error adding user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User added successfully",
		"user":    result,
	})

}

func GetUserByID(c *gin.Context) {
	// GetUserByID logic
	// Read the `id` path parameter from the request
	id := c.Param("id")
	user, err := models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error fetching user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

func GetAllUsers(c *gin.Context) {
	// GetAllUsers logic
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error fetching users",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func GetUserByEmail(c *gin.Context) {
	// GetUserByEmail logic
	email := c.Param("email")
	user, err := models.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error fetching user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

func GetUserByUsername(c *gin.Context) {
	// GetUserByUsername logic
	username := c.Param("username")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error fetching user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

func UpdateUser(c *gin.Context) {
	// UpdateUser logic
	id := c.Param("id")
	user := models.User{}
	c.BindJSON(&user)
	result, err := models.PatchUser(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating user",
		})
		return
	}
	// Update the user object with the inserted ID, the following line will convert the InsertedID to an ObjectID
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    result,
	})
}

func DeleteUser(c *gin.Context) {
	// DeleteUser logic
	id := c.Param("id")
	result, err := models.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error deleting user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"user":    result,
	})
}

// Stream controllers
func Stream(c *gin.Context) {
	generator := loremipsum.New()
	loremIpsum := generator.Sentences(5)
	var idx int = 0
	n := len(loremIpsum)
	c.Stream(func(w io.Writer) bool {
		if idx < n {
			w.Write([]byte(string(loremIpsum[idx])))
			time.Sleep(1 * time.Millisecond)
			idx++
			return true
		}
		w.Write([]byte("\n\nThat's all folks!\n"))
		return false
	})
}

