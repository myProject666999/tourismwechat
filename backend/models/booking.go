package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id" gorm:"index;not null"`
	HotelID       uint           `json:"hotel_id" gorm:"index;not null"`
	RoomType      string         `json:"room_type" gorm:"size:100"`
	CheckInDate   time.Time      `json:"check_in_date"`
	CheckOutDate  time.Time      `json:"check_out_date"`
	GuestCount    int            `json:"guest_count"`
	TotalPrice    float64        `json:"total_price"`
	ContactName   string         `json:"contact_name" gorm:"size:50"`
	ContactPhone  string         `json:"contact_phone" gorm:"size:20"`
	Status        int            `json:"status" gorm:"default:0"`
	Remark        string         `json:"remark" gorm:"type:text"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateBookingRequest struct {
	HotelID      uint    `json:"hotel_id" binding:"required"`
	RoomType     string  `json:"room_type"`
	CheckInDate  string  `json:"check_in_date" binding:"required"`
	CheckOutDate string  `json:"check_out_date" binding:"required"`
	GuestCount   int     `json:"guest_count"`
	TotalPrice   float64 `json:"total_price"`
	ContactName  string  `json:"contact_name"`
	ContactPhone string  `json:"contact_phone"`
	Remark       string  `json:"remark"`
}

type UpdateBookingStatusRequest struct {
	Status int `json:"status" binding:"required"`
}
