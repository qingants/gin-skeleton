package xgorm

import "github.com/jinzhu/gorm"

type Page struct {
	Data  interface{}
	Page  int
	Size  int
	Count int
}

// 1. 提供单表查询的分页，如果要复合查询的分页要另外实现
// 2. 不需要的参数用该类型的默认参数，string类型为""， int类型为0
// 3. fields为""时返回所有字段, fields 应该和out匹配
//	SELECT
//		fields
//	FROM
//		table
//	WHERE
//		where
//	ORDER BY
//		order
//	LIMIT size OFFSET offset
func Search(db *gorm.DB, table, fields string, where *gorm.DB, order string, page, size int, out interface{}) (error, *Page) {
	if fields == "" {
		fields = "*"
	}
	var count int
	if err := db.Table(table).Select("count(1)").Where(where).Limit(-1).Offset(0).Count(&count).Error; err != nil {
		return err, nil
	}
	offset := (page - 1) * size
	// default = 50 and max = 50
	if size == 0 || size > 50 {
		size = 50
	}
	if err := db.Table(table).Select(fields).Where(where).Order(order).Offset(offset).Limit(size).Find(out).Error; err != nil {
		return err, nil
	}
	p := Page{
		Data:  out,
		Page:  page,
		Size:  size,
		Count: count,
	}
	return nil, &p
}
