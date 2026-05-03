package controllers

import (
	"net/http"
	"strconv"
	"time"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func GetHotelStats(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	hotelID := c.Query("hotel_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	offset := (page - 1) * pageSize

	var stats []models.HotelStat
	var total int64

	query := database.DB.Model(&models.HotelStat{})
	if hotelID != "" {
		hotelIDInt, _ := strconv.Atoi(hotelID)
		query = query.Where("hotel_id = ?", hotelIDInt)
	}
	if startDate != "" {
		sDate, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("stat_date >= ?", sDate)
		}
	}
	if endDate != "" {
		eDate, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("stat_date <= ?", eDate)
		}
	}

	query.Count(&total)
	query.Order("stat_date DESC").Offset(offset).Limit(pageSize).Find(&stats)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      stats,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetHotelStatsSummary(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := database.DB.Model(&models.HotelStat{}).
		Select("hotel_id, hotel_name, COUNT(*) as total_bookings, SUM(total_revenue) as total_revenue, AVG(room_occupancy) as avg_occupancy").
		Group("hotel_id, hotel_name")

	if startDate != "" {
		sDate, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("stat_date >= ?", sDate)
		}
	}
	if endDate != "" {
		eDate, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("stat_date <= ?", eDate)
		}
	}

	var summaries []models.HotelStatsSummary
	query.Scan(&summaries)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": summaries,
	})
}

func GenerateHotelStat(c *gin.Context) {
	hotelID := c.Query("hotel_id")
	statDate := c.Query("stat_date")

	if hotelID == "" || statDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少必要参数",
		})
		return
	}

	hotelIDInt, _ := strconv.Atoi(hotelID)
	date, err := time.Parse("2006-01-02", statDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "日期格式错误",
		})
		return
	}

	var hotel models.Hotel
	if err := database.DB.First(&hotel, hotelIDInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "酒店不存在",
		})
		return
	}

	var bookingCount int64
	startOfDay := date
	endOfDay := date.Add(24 * time.Hour)
	database.DB.Model(&models.Booking{}).
		Where("hotel_id = ? AND created_at >= ? AND created_at < ? AND status = ?", hotelIDInt, startOfDay, endOfDay, 1).
		Count(&bookingCount)

	var totalRevenue float64
	var bookings []models.Booking
	database.DB.Where("hotel_id = ? AND created_at >= ? AND created_at < ? AND status = ?", hotelIDInt, startOfDay, endOfDay, 1).
		Find(&bookings)
	for _, b := range bookings {
		totalRevenue += b.TotalPrice
	}

	var existingStat models.HotelStat
	result := database.DB.Where("hotel_id = ? AND stat_date = ?", hotelIDInt, date).First(&existingStat)

	if result.Error == nil {
		database.DB.Model(&existingStat).Updates(map[string]interface{}{
			"booking_count":   bookingCount,
			"total_revenue":   totalRevenue,
			"room_occupancy":  75.0,
		})
	} else {
		stat := &models.HotelStat{
			HotelID:      uint(hotelIDInt),
			HotelName:    hotel.Name,
			StatDate:     date,
			BookingCount: int(bookingCount),
			TotalRevenue: totalRevenue,
			RoomOccupancy: 75.0,
		}
		database.DB.Create(stat)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "统计生成成功",
	})
}
