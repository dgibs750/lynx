package router

import (
	"database/sql"
	"log"

	"github.com/dgibs750/lynx/api/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func New(logger *log.Logger, v *validator.Validate, db *sql.DB) *gin.Engine {
	userAPI := user.New(logger, v, db)
	r := gin.Default()
	r.GET("/user", userAPI.GetUserBy)
	r.POST("/user", userAPI.PostAddUser)
	r.PUT("/user", userAPI.PutUpdateUser)

	return r
}
