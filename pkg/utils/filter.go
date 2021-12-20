package utils

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Filter struct {
	Field string `form:"field"`                 // 指定要从数据库检索的字段，默认情况下，将选择所有字段, 逗号隔开 "name, age"
	Order string `form:"order"`                 // 在从数据库检索记录时指定顺序, 逗号隔开 "age desc, name"
	Limit int    `form:"limit" binding:"min=0"` // 指定要检索的记录数，单页记录数
	Page  int    `form:"page" binding:"min=0"`  // 指定在开始返回记录之前要跳过的记录数, 页码
	Where string `form:"where"`                 // 过滤语句 ["age >= ? and role <> ?",20,"admin"]
}

func GetCommonFilterDB(db *gorm.DB, c *gin.Context) (dbRet *gorm.DB, err error) {
	var filter Filter
	if err = c.ShouldBindQuery(&filter); err != nil {
		return
	}

	field := filter.Field
	if field != "" {
		db = db.Select(field)
	}
	order := filter.Order
	if order != "" {
		db = db.Order(order)
	}
	limit := filter.Limit
	if limit != 0 {
		db = db.Limit(limit)
	} else {
		// 默认最多给50条记录
		db = db.Limit(50)
	}
	page := filter.Page
	if page != 0 {
		db = db.Offset((page - 1) * limit)
	}
	where := c.Query("where")
	if where != "" {
		var whereFilter []string
		err = json.Unmarshal([]byte(where), &whereFilter)
		if err != nil {
			return
		}

		var args []interface{}
		for _, v := range whereFilter[1:] {
			args = append(args, v)
		}
		db = db.Where(whereFilter[0], args...)
	}
	return db, err
}

func HandleName(name string) string {
	if strings.Trim(name, " ") == "" {
		return "-"
	}
	return name
}
