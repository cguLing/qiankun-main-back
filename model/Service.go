package model


type ServiceType struct {
	GormModel
	Index uint `gorm:"type:int" json:"index" form:"index"` // 排序
	ClassName string `gorm:"type:varchar(128);not null;unique" json:"class_name" form:"class_name"` // 分类名称
}

type MicroList struct{
	GormModel
	Name string `gorm:"type:varchar(128);not null;unique_index" json:"name" form:"name"` // 名称
	Entry string `gorm:"type:varchar(255);not null" json:"entry" form:"entry"` // 接入链接
	Icon string `gorm:"type:varchar(128);not null" json:"icon" form:"icon"` // 图标
	Title string `gorm:"type:varchar(128);not null" json:"title" form:"title"` // 标题
	ActiveRule string `gorm:"type:varchar(128);not null;unique_index" json:"activeRule" form:"activeRule"` // 接入路由
}

type ServiceList struct{
	GormModel
	Name        string      `gorm:"type:varchar(128);not null" json:"name" form:"name"`   // 名称
	Url         string      `gorm:"type:varchar(255);not null" json:"url" form:"url"` // 接入链接
	PoPo         string      `gorm:"type:varchar(255)" json:"popo" form:"popo"` // 接入链接
	Doc         string      `gorm:"type:varchar(255)" json:"Doc" form:"Doc"` // 接入链接
	Status      uint      `gorm:"type:int;not null" json:"status" form:"status"`   // 是否接入
	Desc        string      `gorm:"type:varchar(128);not null" json:"desc" form:"desc"` // 描述
	ServiceType ServiceType `gorm:"foreignKey:ClassId;ASSOCIATION_FOREIGNKEY:ID;"`
	MicroList   MicroList   `gorm:"foreignKey:MenuId;ASSOCIATION_FOREIGNKEY:ID;"`
	ClassId     uint        `gorm:"unique_index" json:"class_id" form:"class_id"` // 分类外键
	MenuId      uint        `gorm:"unique_index" json:"menu_id" form:"menu_id"` // 菜单外键
}

type EnshrineList struct{
	GormModel
	ServiceList ServiceList `gorm:"foreignKey:ServiceId;association_foreignkey:ID;"`
	//UserAdmin UserAdmin `gorm:"foreignkey:UserName;association_foreignkey:Ldap"`
	ServiceId uint `gorm:"unique_index" json:"service_id" form:"service_id"` // 分类外键
	UserName string `gorm:"type:varchar(128);not null;unique_index" json:"user_name" form:"user_name"`
}