package model

type SuperAdmin struct {
	GormModel
	Ldap    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"ldap" form:"ldap"` // 超级管理员的ldap账号 唯一
	Creator string `gorm:"type:varchar(255);default:'root'" json:"creator" form:"creator"` // 添加这条数据的人
}

