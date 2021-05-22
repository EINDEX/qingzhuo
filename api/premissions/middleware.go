package premissions

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Premission string

const (
	POST_EDITOR Premission = "post.editor"
)

func PremissionApplyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-Api-Key")
		if key != os.Getenv("API_KEY") {
			c.Next()
			return
		}
		premissionMap := make(map[Premission]bool)
		premissionMap[POST_EDITOR] = true
		c.Set("premissions", premissionMap)
		c.Next()
	}
}

func checkingPremissions(requirePremissions []Premission, havePremissions map[Premission]bool) bool {
	if len(requirePremissions) == 0 {
		return true
	}
	for _, requirePremission := range requirePremissions {
		if _, ok := havePremissions[requirePremission]; !ok {
			return false
		}
	}
	return true
}

func PremissionCheck(requirePremissions ...Premission) gin.HandlerFunc {
	return func(c *gin.Context) {
		havePremissionsInterface, exists := c.Get("premissions")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if havePremissions, ok := havePremissionsInterface.(map[Premission]bool); ok {
			if checkingPremissions(requirePremissions, havePremissions) {
				c.Next()
				return
			}
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}
