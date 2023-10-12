package entities

import "gorm.io/gorm"

type Interaction struct {
	gorm.Model
	User     *User `gorm:"embedded,foreignKey:UserID"`
	UserID   uint
	MovieID  uint
	Movie    *Movie `gorm:"embedded,foreignKey:MovieID"`
	Recombee bool   `gorm:"default:false"`
}
