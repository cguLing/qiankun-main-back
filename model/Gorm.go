package model

import (
	"gorm.io/gorm"
	"time"
)

type Page struct {
	Perpage int   `json:"per_page" form:"per_page"`
	Page    int   `json:"page" form:"page"`
	Count   int64 `json:"count" form:"count"`
}

type GormModel struct {
	ID        uint           `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt time.Time      `json:"create_at" form:"create_at"`
	UpdatedAt time.Time      `json:"update_at" form:"update_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at" form:"delete_at"`
}

// 数据库分页
func Paginate(page Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.Page <= 0 {
			page.Page = 1
		}
		//page.Page +=1
		if page.Perpage > 1000 {
			page.Perpage = 1000
		}
		if page.Perpage <= 0 {
			page.Perpage = 10
		}
		offset := (page.Page - 1) * page.Perpage
		return db.Offset(offset).Limit(page.Perpage)
	}
}
