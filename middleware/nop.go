package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func NopCloser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body != nil {
			if body, err := c.GetRawData(); err == nil {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		c.Next()
	}
}
