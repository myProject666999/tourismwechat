package models

import (
	"time"

	"gorm.io/gorm"
)

type DailyRecommend struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Cover       string         `json:"cover" gorm:"size:255"`
	RecommendType string       `json:"recommend_type" gorm:"size:50"`
	TargetID    uint           `json:"target_id"`
	TargetName  string         `json:"target_name" gorm:"size:200"`
	Priority    int            `json:"priority" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateDailyRecommendRequest struct {
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description"`
	Cover         string `json:"cover"`
	RecommendType string `json:"recommend_type"`
	TargetID      uint   `json:"target_id"`
	TargetName    string `json:"target_name"`
	Priority      int    `json:"priority"`
	Status        int    `json:"status"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}

type UpdateDailyRecommendRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Cover         string `json:"cover"`
	RecommendType string `json:"recommend_type"`
	TargetID      uint   `json:"target_id"`
	TargetName    string `json:"target_name"`
	Priority      int    `json:"priority"`
	Status        int    `json:"status"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}
