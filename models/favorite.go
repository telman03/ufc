package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model

	UserID    uint		`gorm:"not null" json:"user_id"`
	FighterID uint		`gorm:"not null" json:"fighter_id"`
	Fighter  Fighter 	`gorm:"foreignKey:FighterID" json:"fighter"` // Add this line

}


type FavoriteInput struct {
	FighterID uint `json:"fighter_id" binding:"required"`
}