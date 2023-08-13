package user

import (
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
)

type API struct {
	logger     *log.Logger
	validator  *validator.Validate
	repository *Repository
}

func New(logger *log.Logger, validator *validator.Validate, db *sql.DB) *API {
	return &API{
		logger:     logger,
		validator:  validator,
		repository: NewRepository(db),
	}
}
