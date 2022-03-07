package model


type ServiceType struct {
	GormModel
	Index uint `gorm:"type:int" json:"index" form:"index"` // 排序
	ClassName string `gorm:"type:varchar(128);not null;unique" json:"class_name" form:"class_name"` // 分类名称
}

type MicroList struct{
	GormModel
	name string `gorm:"type:varchar(128);not null;unique_index" json:"name" form:"name"` // 名称
	entry string `gorm:"type:varchar(255);not null" json:"entry" form:"entry"` // 接入链接
	icon string `gorm:"type:varchar(128);not null" json:"icon" form:"icon"` // 图标
	title string `gorm:"type:varchar(128);not null" json:"title" form:"title"` // 标题
	activeRule string `gorm:"type:varchar(128);not null;unique_index" json:"activeRule" form:"activeRule"` // 接入路由
}

type ServiceList struct{
	GormModel
	name string `gorm:"type:varchar(128);not null" json:"name" form:"name"` // 名称
	url string `gorm:"type:varchar(255);not null" json:"entry" form:"entry"` // 接入链接
	status string `gorm:"type:varchar(128);not null" json:"icon" form:"icon"` // 是否接入
	desc string `gorm:"type:varchar(128);not null" json:"title" form:"title"` // 描述
	ServiceType ServiceType `gorm:"foreignKey:ID;association_foreignkey:classId;"`
	MicroList MicroList `gorm:"foreignKey:ID;association_foreignkey:menuId;"`
	classId uint `gorm:"unique_index" json:"classId" form:"classId"` // 分类外键
	menuId uint `gorm:"unique_index" json:"menuId" form:"menuId"` // 菜单外键
}

type EnshrineList struct{
	GormModel
	ServiceList ServiceList `gorm:"foreignKey:ID;association_foreignkey:classId;"`
	SuperAdmin SuperAdmin `gorm:"foreignkey:ID;association_foreignkey:userId"`
	serviceId uint `gorm:"unique_index" json:"serviceId" form:"serviceId"` // 分类外键
	userId uint `gorm:"unique_index" json:"userId" form:"userId"` // 菜单外键
}