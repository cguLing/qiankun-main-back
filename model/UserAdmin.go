package model

type UserAdmin struct {
	GormModel
	Ldap    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"ldap" form:"ldap"`
	MenuWay bool `gorm:"type:bool;default:true" json:"menuWay" form:"menuWay"`
	Creator string `gorm:"type:varchar(255);default:'root'" json:"creator" form:"creator"`
}

