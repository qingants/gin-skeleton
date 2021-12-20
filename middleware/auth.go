package middleware

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)

//func CheckToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if c.GetHeader("token") != setting.AuthKey || c.GetHeader("Token") != setting.AuthKey {
//			c.AbortWithStatusJSON(http.StatusOK, models.GetResponse(errs.ErrorForbidden, ""))
//			return
//		}
//		c.Next()
//		return
//	}
//}
