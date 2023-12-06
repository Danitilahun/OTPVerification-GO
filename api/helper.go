package api;

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New();
func (app *Config) validateBody(c *gin.Context , data any) error{
	err := c.BindJSON(&data);
	if err != nil {
		return err
	}

	err = validate.Struct(data);
	if err != nil {
		return err
	}
	return nil
}

