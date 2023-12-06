package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/Danitilahun/OTPVerification-GO/model"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel = context.WithTimeout(context.Background(), appTimeout)

		var payload model.OTPData;
		defer cancel();
		app.validateBody(c, &payload);
		// send sms
		
		_, err := app.twilioSendOTP(payload.PhoneNumber);
		if err != nil {
			app.errorJSON(c, err);
			return;
		}
		app.successJSON(c, "OTP send Successful", http.StatusAccepted);
	}
};
func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)

		var payload model.VerifyOTPData;
		defer cancel();
		app.validateBody(c, &payload);
		// verify sms
		_, err := app.twilioVerifyOTP(payload.User.PhoneNumber, payload.Code);
		if err != nil {
			app.errorJSON(c, err);
			return;
		}
		app.successJSON(c, "OTP verified Successful", http.StatusAccepted);
	}
};