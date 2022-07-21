package util

import "github.com/gin-gonic/gin"

func Res200(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}
func Res401(c *gin.Context) {
	c.JSON(401, gin.H{"err": "permission denied"})
}
func Res401str(c *gin.Context, str string) {
	c.JSON(401, gin.H{"err": str})
}
func Res403(c *gin.Context) {
	c.JSON(403, gin.H{"err": "forbidden"})
}
func Res403str(c *gin.Context, str string) {
	c.JSON(403, gin.H{"err": str})
}
func Res404(c *gin.Context) {
	c.JSON(404, gin.H{"err": "not found"})
}
func Res404str(c *gin.Context, str string) {
	c.JSON(404, gin.H{"err": str})
}

func Res406(c *gin.Context) {
	c.JSON(406, gin.H{"err": "request type error"})
}
func Res406str(c *gin.Context, str string) {
	c.JSON(406, gin.H{"err": str})
}

func Res500(c *gin.Context) {
	c.JSON(500, gin.H{"err": "server error"})
}

func Res500str(c *gin.Context, str string) {
	c.JSON(500, gin.H{"err": str})
}
