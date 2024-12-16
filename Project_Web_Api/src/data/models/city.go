package models


type City struct {
	BaseModel
	Name string `gorm:"size:15;type:string;not null"`
	CountryId int `gorm:"not null"`
	Country Country `gorm:"foreignkey:CountryId"`

}