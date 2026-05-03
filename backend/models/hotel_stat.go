package models

import (
	"time"
)

type HotelStat struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	HotelID    uint      `json:"hotel_id" gorm:"uniqueIndex:idx_hotel_date;not null"`
	HotelName  string    `json:"hotel_name" gorm:"size:100"`
	StatDate   time.Time `json:"stat_date" gorm:"uniqueIndex:idx_hotel_date;not null"`
	BookingCount int     `json:"booking_count" gorm:"default:0"`
	TotalRevenue float64 `json:"total_revenue" gorm:"default:0"`
	RoomOccupancy float64 `json:"room_occupancy" gorm:"default:0"`
	ViewCount  int       `json:"view_count" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type HotelStatsSummary struct {
	HotelID       uint    `json:"hotel_id"`
	HotelName     string  `json:"hotel_name"`
	TotalBookings int     `json:"total_bookings"`
	TotalRevenue  float64 `json:"total_revenue"`
	AvgOccupancy  float64 `json:"avg_occupancy"`
}
