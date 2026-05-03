package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Cover       string         `json:"cover" gorm:"size:255"`
	LikeCount   int            `json:"like_count" gorm:"default:0"`
	DislikeCount int           `json:"dislike_count" gorm:"default:0"`
	ViewCount   int            `json:"view_count" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type AnnouncementLike struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id" gorm:"uniqueIndex:idx_user_announcement;not null"`
	AnnouncementID uint      `json:"announcement_id" gorm:"uniqueIndex:idx_user_announcement;not null"`
	Type           int       `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateAnnouncementRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
	Status  int    `json:"status"`
}

type UpdateAnnouncementRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
	Status  int    `json:"status"`
}
