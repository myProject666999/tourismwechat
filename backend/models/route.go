package models

import (
	"time"

	"gorm.io/gorm"
)

type Route struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Description string         `json:"description" gorm:"type:text"`
	StartPoint  string         `json:"start_point" gorm:"size:100"`
	EndPoint    string         `json:"end_point" gorm:"size:100"`
	StartLat    float64        `json:"start_lat"`
	StartLng    float64        `json:"start_lng"`
	EndLat      float64        `json:"end_lat"`
	EndLng      float64        `json:"end_lng"`
	Distance    float64        `json:"distance"`
	Duration    string         `json:"duration" gorm:"size:50"`
	Cover       string         `json:"cover" gorm:"size:255"`
	WayPoints   string         `json:"way_points" gorm:"type:text"`
	Difficulty  int            `json:"difficulty"`
	ViewCount   int            `json:"view_count" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateRouteRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	StartPoint  string  `json:"start_point"`
	EndPoint    string  `json:"end_point"`
	StartLat    float64 `json:"start_lat"`
	StartLng    float64 `json:"start_lng"`
	EndLat      float64 `json:"end_lat"`
	EndLng      float64 `json:"end_lng"`
	Distance    float64 `json:"distance"`
	Duration    string  `json:"duration"`
	Cover       string  `json:"cover"`
	WayPoints   string  `json:"way_points"`
	Difficulty  int     `json:"difficulty"`
	Status      int     `json:"status"`
}

type UpdateRouteRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	StartPoint  string  `json:"start_point"`
	EndPoint    string  `json:"end_point"`
	StartLat    float64 `json:"start_lat"`
	StartLng    float64 `json:"start_lng"`
	EndLat      float64 `json:"end_lat"`
	EndLng      float64 `json:"end_lng"`
	Distance    float64 `json:"distance"`
	Duration    string  `json:"duration"`
	Cover       string  `json:"cover"`
	WayPoints   string  `json:"way_points"`
	Difficulty  int     `json:"difficulty"`
	Status      int     `json:"status"`
}
