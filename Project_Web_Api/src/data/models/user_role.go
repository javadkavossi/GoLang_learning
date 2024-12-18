package models

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignkey:UserId; constraint:OnUpdate:NO ACTION;OnDELETE:NO ACTION"`
	Role   Role `gorm:"foreignkey:RoleId; constraint:OnUpdate:NO ACTION;OnDELETE:NO ACTION"`
	UserId int
	RoleId int
}
