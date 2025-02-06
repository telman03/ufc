package models

import "gorm.io/gorm"


type Fighter struct {
	gorm.Model

	Name 		string 	`gorm:"unique;not null"`
	Rank		int
	Division 	string
}