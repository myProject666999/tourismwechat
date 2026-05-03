package models

import (
	"time"

	"gorm.io/gorm"
)

type Itinerary struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"index;not null"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Description string         `json:"description" gorm:"type:text"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	Status      int            `json:"status" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type ItineraryItem struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ItineraryID uint      `json:"itinerary_id" gorm:"index;not null"`
	DayIndex    int       `json:"day_index"`
	Title       string    `json:"title" gorm:"size:200"`
	Description string    `json:"description" gorm:"type:text"`
	StartTime   string    `json:"start_time" gorm:"size:20"`
	EndTime     string    `json:"end_time" gorm:"size:20"`
	Location    string    `json:"location" gorm:"size:200"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	ItemType    string    `json:"item_type" gorm:"size:50"`
	OrderNum    int       `json:"order_num"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateItineraryRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type CreateItineraryItemRequest struct {
	ItineraryID uint    `json:"itinerary_id" binding:"required"`
	DayIndex    int     `json:"day_index"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Location    string  `json:"location"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ItemType    string  `json:"item_type"`
	OrderNum    int     `json:"order_num"`
}

type UpdateItineraryStatusRequest struct {
	Status int `json:"status" binding:"required"`
}
