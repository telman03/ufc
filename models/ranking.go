package models

import (
    "math/rand"
    "time"
    "gorm.io/gorm"
)

type Ranking struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    FighterID uint           `gorm:"not null" json:"fighter_id"`
    Fighter   Fighter        `gorm:"foreignKey:FighterID" json:"fighter"`
    Rank      int            `gorm:"not null" json:"rank"`
    Division  string         `gorm:"not null" json:"division"`
}

// BeforeCreate is a GORM hook that sets a random ID before creating a record
func (r *Ranking) BeforeCreate(tx *gorm.DB) (err error) {
    r.ID = uint(rand.Intn(1000000)) // Generate a random ID
    return
}
