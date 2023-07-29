package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saketharshraj/email-authenticator-golang/controllers"
	"github.com/saketharshraj/email-authenticator-golang/middlewares"
)

func EmailRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/send-otp", middlewares.ValidateEmailPayload(), controllers.SendOtp())
	incomingRoutes.POST("/verify-otp", controllers.VerifyMail())

}