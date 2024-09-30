package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Role int

const (
	Admin Role = iota
	Mod
	User
)

type UserAuth struct {
	UserId   string `json:"userId"`
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}

func GetAuthAndPutCurrentUserAuthToBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUserJson := c.GetHeader("X-Current-User")
		if currentUserJson == "" {
			log.Printf("Not have auth")
			c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "Not have auth"})
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
}

func GetUserAuthFromContext(c *gin.Context) UserAuth {
	currentUserInfo, _ := c.Get("currentUser")
	currentUser, _ := currentUserInfo.(UserAuth)

	return currentUser
}
