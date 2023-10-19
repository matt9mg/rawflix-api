package entities

import "gorm.io/gorm"

type InteractionType = string

var (
	InteractionTypeDetailView InteractionType = "detail_view"
	InteractionTypeBookmark   InteractionType = "bookmaark"
)

type Interaction struct {
	gorm.Model
	User          *User `gorm:"embedded,foreignKey:UserID"`
	UserID        uint
	MovieID       uint
	Movie         *Movie `gorm:"embedded,foreignKey:MovieID"`
	Recombee      bool   `gorm:"default:false"`
	Type          InteractionType
	RecommenderID string
}
