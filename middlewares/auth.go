package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/goddie9x/social_utils_go/constants"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	UserId   string         `json:"userId"`
	UserName string         `json:"username"`
	Role     constants.Role `json:"role"`
}

// verifySignature checks the HMAC-SHA256 signature the API gateway attaches
// as X-Internal-Signature over the X-Current-User payload. Without this,
// anyone can forge X-Current-User and impersonate any user (including Admin).
func verifySignature(payload, signature string) bool {
	secret := os.Getenv("INTERNAL_HMAC_SECRET")
	if secret == "" || signature == "" || payload == "" {
		return false
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expected))
}

func PutAuthToContext(c *gin.Context) {
	currentUserJson := c.GetHeader("X-Current-User")
	if currentUserJson == "" {
		c.Next()
		return
	}

	// If the HMAC secret is configured, only trust the header when it carries
	// a valid gateway signature.
	if os.Getenv("INTERNAL_HMAC_SECRET") != "" {
		signature := c.GetHeader("X-Internal-Signature")
		if signature == "" {
			// Not coming through the gateway: treat as unauthenticated
			// instead of trusting a forgeable header.
			c.Next()
			return
		}
		if !verifySignature(currentUserJson, signature) {
			log.Printf("Rejected forged X-Current-User header: invalid HMAC signature")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid internal auth signature"})
			return
		}
	} else {
		log.Printf("WARNING: INTERNAL_HMAC_SECRET is not set - trusting X-Current-User without verification")
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
