package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"post_service/pkg/constants"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	UserId   string         `json:"userId"`
	UserName string         `json:"username"`
	Role     constants.Role `json:"role"`
}

func PutAuthToContext(c *gin.Context) {
	currentUserJson := c.GetHeader("X-Current-User")
	if currentUserJson == "" {
		c.Next()
		return
	}
	var currentUser UserAuth
	if err := json.Unmarshal([]byte(currentUserJson), &currentUser); err != nil {
		log.Printf("Failed to parse X-Current-User header: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid X-Current-User header"})
		return
	}
	c.Set("currentUser", currentUser)
	c.Next()
}

func GetUserAuthFromContext(c *gin.Context) UserAuth {
	currentUserInfo, _ := c.Get("currentUser")
	currentUser, _ := currentUserInfo.(UserAuth)

	return currentUser
}
