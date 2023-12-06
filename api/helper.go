package api;

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

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


func (app *Config) errorJSON(c *gin.Context,err error ,  status ...int) {
	statusCode := http.StatusBadRequest;
	if len(status) > 0 {
		statusCode = status[0];
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()});
}

func (app *Config) successJSON(c *gin.Context, data any, status int) {
	c.JSON(status, jsonResponse{Status: status, Data: data});}