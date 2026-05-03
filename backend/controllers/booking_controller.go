package controllers

import (
	"net/http"
	"strconv"
	"time"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req models.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	checkInDate, err := time.Parse("2006-01-02", req.CheckInDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "入住日期格式错误",
		})
		return
	}

	checkOutDate, err := time.Parse("2006-01-02", req.CheckOutDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "退房日期格式错误",
		})
		return
	}

	booking := &models.Booking{
		UserID:       userID,
		HotelID:      req.HotelID,
		RoomType:     req.RoomType,
		CheckInDate:  checkInDate,
		CheckOutDate: checkOutDate,
		GuestCount:   req.GuestCount,
		TotalPrice:   req.TotalPrice,
		ContactName:  req.ContactName,
		ContactPhone: req.ContactPhone,
		Status:       0,
		Remark:       req.Remark,
	}

	if err := database.DB.Create(booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "预订失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "预订成功",
		"data":    booking,
	})
}

func GetUserBookings(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var bookings []models.Booking
	var total int64

	query := database.DB.Model(&models.Booking{}).Where("user_id = ?", userID)
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		query = query.Where("status = ?", statusInt)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&bookings)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      bookings,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetBookingDetail(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var booking models.Booking
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "预订不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    booking,
	})
}

func CancelBooking(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var booking models.Booking
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "预订不存在",
		})
		return
	}

	if booking.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "当前状态无法取消",
		})
		return
	}

	database.DB.Model(&booking).Update("status", 3)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "取消成功",
	})
}

func AdminGetBookingList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var bookings []models.Booking
	var total int64

	query := database.DB.Model(&models.Booking{})
	if keyword != "" {
		query = query.Where("contact_name LIKE ? OR contact_phone LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		query = query.Where("status = ?", statusInt)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&bookings)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      bookings,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func UpdateBookingStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateBookingStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Model(&models.Booking{}).Where("id = ?", id).Update("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}
