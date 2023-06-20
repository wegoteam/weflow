package base

import "gorm.io/gorm"

// Paginate
// @Description: 分页封装
// @param: pageNum 页码
// @param: pageSize 每页条数
// @return func(db *gorm.DB) *gorm.DB
func Paginate(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 5000:
			pageSize = 5000
		case pageSize <= 0:
			pageSize = 30
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// Page
// @Description: 分页实体
type Page[T comparable] struct {
	Total    int64 `json:"total"`    // 总条数
	PageNum  int   `json:"pageNum"`  // 页码
	PageSize int   `json:"pageSize"` // 每页条数
	Records  []T   `json:"records"`  // 数据
}
