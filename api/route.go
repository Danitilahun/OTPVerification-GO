package api;

import (
	"github.com/gin-gonic/gin"
);

type Config struct {
	Router *gin.Engine;
};


func (app *Config) Routes() {
	app.Router.POST("/otp", app.OTP);
	app.Router.POST("/verify", app.VerifyOTP);
}
