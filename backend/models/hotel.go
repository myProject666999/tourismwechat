package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotel struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Address     string         `json:"address" gorm:"size:255"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Cover       string         `json:"cover" gorm:"size:255"`
	Images      string         `json:"images" gorm:"type:text"`
	Star        int            `json:"star"`
	Price       float64        `json:"price"`
	RoomTypes   string         `json:"room_types" gorm:"type:text"`
	Facilities  string         `json:"facilities" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type HotelRoom struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	HotelID     uint      `json:"hotel_id" gorm:"index;not null"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Images      string    `json:"images" gorm:"type:text"`
	Status      int       `json:"status" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateHotelRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Cover       string  `json:"cover"`
	Images      string  `json:"images"`
	Star        int     `json:"star"`
	Price       float64 `json:"price"`
	RoomTypes   string  `json:"room_types"`
	Facilities  string  `json:"facilities"`
	Status      int     `json:"status"`
}

type UpdateHotelRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Cover       string  `json:"cover"`
	Images      string  `json:"images"`
	Star        int     `json:"star"`
	Price       float64 `json:"price"`
	RoomTypes   string  `json:"room_types"`
	Facilities  string  `json:"facilities"`
	Status      int     `json:"status"`
}
