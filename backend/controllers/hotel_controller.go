package controllers

import (
	"net/http"
	"strconv"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func GetHotelList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	star := c.Query("star")

	offset := (page - 1) * pageSize

	var hotels []models.Hotel
	var total int64

	query := database.DB.Model(&models.Hotel{}).Where("status = ?", 1)
	if keyword != "" {
		query = query.Where("name LIKE ? OR address LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if star != "" {
		query = query.Where("star = ?", star)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&hotels)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      hotels,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetHotelDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var hotel models.Hotel
	if err := database.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "酒店不存在",
		})
		return
	}

	var rooms []models.HotelRoom
	database.DB.Where("hotel_id = ? AND status = ?", id, 1).Find(&rooms)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"hotel": hotel,
			"rooms": rooms,
		},
	})
}

func AdminGetHotelList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var hotels []models.Hotel
	var total int64

	query := database.DB.Model(&models.Hotel{})
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&hotels)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      hotels,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func CreateHotel(c *gin.Context) {
	var req models.CreateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	hotel := &models.Hotel{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Cover:       req.Cover,
		Images:      req.Images,
		Star:        req.Star,
		Price:       req.Price,
		RoomTypes:   req.RoomTypes,
		Facilities:  req.Facilities,
		Status:      req.Status,
	}

	if err := database.DB.Create(hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    hotel,
	})
}

func UpdateHotel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}
	if req.Latitude != 0 {
		updates["latitude"] = req.Latitude
	}
	if req.Longitude != 0 {
		updates["longitude"] = req.Longitude
	}
	if req.Cover != "" {
		updates["cover"] = req.Cover
	}
	if req.Images != "" {
		updates["images"] = req.Images
	}
	if req.Star != 0 {
		updates["star"] = req.Star
	}
	if req.Price != 0 {
		updates["price"] = req.Price
	}
	if req.RoomTypes != "" {
		updates["room_types"] = req.RoomTypes
	}
	if req.Facilities != "" {
		updates["facilities"] = req.Facilities
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}

	if err := database.DB.Model(&models.Hotel{}).Where("id = ?", id).Updates(updates).Error; err != nil {
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

func DeleteHotel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Delete(&models.Hotel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
