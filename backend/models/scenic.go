package models

import (
	"time"

	"gorm.io/gorm"
)

type Scenic struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"size:100;not null"`
	Description  string         `json:"description" gorm:"type:text"`
	Address      string         `json:"address" gorm:"size:255"`
	Latitude     float64        `json:"latitude"`
	Longitude    float64        `json:"longitude"`
	Cover        string         `json:"cover" gorm:"size:255"`
	Images       string         `json:"images" gorm:"type:text"`
	Price        float64        `json:"price"`
	OpenTime     string         `json:"open_time" gorm:"size:100"`
	Level        string         `json:"level" gorm:"size:20"`
	LikeCount    int            `json:"like_count" gorm:"default:0"`
	DislikeCount int            `json:"dislike_count" gorm:"default:0"`
	ViewCount    int            `json:"view_count" gorm:"default:0"`
	Status       int            `json:"status" gorm:"default:1"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type ScenicLike struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"uniqueIndex:idx_user_scenic;not null"`
	ScenicID  uint      `json:"scenic_id" gorm:"uniqueIndex:idx_user_scenic;not null"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateScenicRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Cover       string  `json:"cover"`
	Images      string  `json:"images"`
	Price       float64 `json:"price"`
	OpenTime    string  `json:"open_time"`
	Level       string  `json:"level"`
	Status      int     `json:"status"`
}

type UpdateScenicRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Cover       string  `json:"cover"`
	Images      string  `json:"images"`
	Price       float64 `json:"price"`
	OpenTime    string  `json:"open_time"`
	Level       string  `json:"level"`
	Status      int     `json:"status"`
}
