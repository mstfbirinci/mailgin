package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyMiddleware() *gin.Engine {

	return gin.Default()
}

func initRoutes(router *gin.Engine) {

	router.POST("/send", func(c *gin.Context) {

		mailMessage := &Message{}

		if bindingError := c.BindJSON(mailMessage); bindingError != nil {

			failureResponse := Exception{
				Error:            "Failed to deserialize request content. If you are sending HTML content make sure you properly escape double-quotes or just use single-quotes in your content",
				ExceptionMessage: bindingError.Error(),
			}

			c.JSON(http.StatusInternalServerError, failureResponse)
			return
		}

		SendMail(mailMessage)

		c.Writer.WriteHeader(http.StatusAccepted)

	})

}

//RunServer initialize app
func RunServer(port string) {

	router := applyMiddleware()
	initRoutes(router)

	router.Run(":" + port)

}
