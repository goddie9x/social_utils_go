package exceptions

import "github.com/gin-gonic/gin"

func HandleExceptionByGin(c *gin.Context, exception CommonExceptionInterface) {
	c.JSON(exception.GetHTTPErrorCode(), exception.GetErrorMessage())
}
