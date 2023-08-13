package api

// import (
// "net/http"

// "github.com/gin-gonic/gin"
// )

type user struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	UserPassword string `json:"userPassword"`
}

type account struct {
	ID              int    `json:"id"`
	UserID          int    `json:"userID"`
	AccountName     string `json:"accountName"`
	Username        string `json:"username"`
	AccountPassword string `json:"accountPassword"`
	Website         string `json:"website"`
	Category        string `json:"category"`
}

// var creds = []cred {}

// getCrefsByUID locates all credentials for the given user.
// func getCredsByUID(c * gin.Context) {
// 	id := c.Param("uid")

// }
