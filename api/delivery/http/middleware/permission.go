package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Permission string

const (
	POST_EDITOR Permission = "post.editor"
)

func PermissionApplyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-Api-Key")
		if key != os.Getenv("API_KEY") {
			c.Next()
			return
		}
		PermissionMap := make(map[Permission]bool)
		PermissionMap[POST_EDITOR] = true
		c.Set("Permissions", PermissionMap)
		c.Next()
	}
}

func checkingPermissions(requirePermissions []Permission, havePermissions map[Permission]bool) bool {
	if len(requirePermissions) == 0 {
		return true
	}
	for _, requirePermission := range requirePermissions {
		if _, ok := havePermissions[requirePermission]; !ok {
			return false
		}
	}
	return true
}

func PermissionCheck(requirePermissions ...Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		havePermissionsInterface, exists := c.Get("Permissions")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if havePermissions, ok := havePermissionsInterface.(map[Permission]bool); ok {
			if checkingPermissions(requirePermissions, havePermissions) {
				c.Next()
				return
			}
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}
