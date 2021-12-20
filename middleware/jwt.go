package middleware

import (
	//"github.com/gin-gonic/gin"
	//"github.com/pkthigh/k1/pkg/errs"
	//"github.com/pkthigh/k1/pkg/utils"
	//"net/http"
	//"time"
)

//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var code int
//		var data interface{}
//
//		code = errs.Success
//		token := c.Query("token")
//		if token == "" {
//			code = errs.InvalidParams
//		} else {
//			claim, err := utils.ParseToken(token)
//			if err != nil {
//				code = errs.ErrAuthCheckTokenFail
//			} else if time.Now().Unix() > claim.ExpiresAt {
//				code = errs.ErrAuthCheckTokenTimeout
//			}
//		}
//		if code != errs.Success {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"code": code,
//				"msg":  errs.GetMsg(code),
//				"data": data,
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
