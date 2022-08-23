package request

import "github.com/gin-gonic/gin"

func GetRequestAdminToken(c *gin.Context) string {
	return c.GetHeader("crmt")
}

func GetRequestAdminUserId(c *gin.Context) string {
	return c.GetHeader("crmu")
}
