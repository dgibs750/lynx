package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/dgibs750/lynx/util/validator"
)

func (a *API) GetUserBy(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		stringId := c.Query("id")
		id, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is not valid"})
			return
		}
		user, err := a.repository.UserById(id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.IndentedJSON(http.StatusFound, user)
		return
	}

	user, err := a.repository.UserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusFound, user)
}

func (a *API) PostAddUser(c *gin.Context) {
	var newUser *NewUser

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := a.validator.Struct(newUser); err != nil {
		resp := validator.ToErrResponse(err)
		if resp == nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": resp})
		return
	}

	user, err := a.repository.AddUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}

func (a *API) PutUpdateUser(c *gin.Context) {
	var updateUserData *UpdateUserData

	if err := c.BindJSON(&updateUserData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := a.validator.Struct(updateUserData); err != nil {
		resp := validator.ToErrResponse(err)
		if resp == nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": resp})
		return
	}

	updatedUser, err := a.repository.UpdateUser(updateUserData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusAccepted, updatedUser)
}
