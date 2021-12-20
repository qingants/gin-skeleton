package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qingants/gin-skeleton/pkg/errs"
	"go.uber.org/zap"
	"net/http"
)

type LoginModel struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


func Login(c *gin.Context)  {
	var params LoginModel
	if err := c.ShouldBind(&params); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, ""))
		return
	}
}
