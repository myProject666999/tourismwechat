package controllers

import (
	"net/http"
	"strconv"
	"time"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func CreateItinerary(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req models.CreateItineraryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	itinerary := &models.Itinerary{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Status:      0,
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err == nil {
			itinerary.StartDate = startDate
		}
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err == nil {
			itinerary.EndDate = endDate
		}
	}

	if err := database.DB.Create(itinerary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    itinerary,
	})
}

func GetUserItineraries(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var itineraries []models.Itinerary
	var total int64

	query := database.DB.Model(&models.Itinerary{}).Where("user_id = ?", userID)
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		query = query.Where("status = ?", statusInt)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&itineraries)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      itineraries,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetItineraryDetail(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var itinerary models.Itinerary
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&itinerary).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "行程不存在",
		})
		return
	}

	var items []models.ItineraryItem
	database.DB.Where("itinerary_id = ?", id).Order("day_index ASC, order_num ASC").Find(&items)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"itinerary": itinerary,
			"items":     items,
		},
	})
}

func AddItineraryItem(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req models.CreateItineraryItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var itinerary models.Itinerary
	if err := database.DB.Where("id = ? AND user_id = ?", req.ItineraryID, userID).First(&itinerary).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "行程不存在",
		})
		return
	}

	item := &models.ItineraryItem{
		ItineraryID: req.ItineraryID,
		DayIndex:    req.DayIndex,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Location:    req.Location,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		ItemType:    req.ItemType,
		OrderNum:    req.OrderNum,
	}

	if err := database.DB.Create(item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "添加成功",
		"data":    item,
	})
}

func UpdateItineraryItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.CreateItineraryItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var item models.ItineraryItem
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "行程项不存在",
		})
		return
	}

	var itinerary models.Itinerary
	if err := database.DB.Where("id = ? AND user_id = ?", item.ItineraryID, userID).First(&itinerary).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "无权操作",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.DayIndex != 0 {
		updates["day_index"] = req.DayIndex
	}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.StartTime != "" {
		updates["start_time"] = req.StartTime
	}
	if req.EndTime != "" {
		updates["end_time"] = req.EndTime
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Latitude != 0 {
		updates["latitude"] = req.Latitude
	}
	if req.Longitude != 0 {
		updates["longitude"] = req.Longitude
	}
	if req.ItemType != "" {
		updates["item_type"] = req.ItemType
	}
	if req.OrderNum != 0 {
		updates["order_num"] = req.OrderNum
	}

	if err := database.DB.Model(&models.ItineraryItem{}).Where("id = ?", id).Updates(updates).Error; err != nil {
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

func DeleteItineraryItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var item models.ItineraryItem
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "行程项不存在",
		})
		return
	}

	var itinerary models.Itinerary
	if err := database.DB.Where("id = ? AND user_id = ?", item.ItineraryID, userID).First(&itinerary).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "无权操作",
		})
		return
	}

	if err := database.DB.Delete(&item).Error; err != nil {
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

func AdminGetItineraryList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var itineraries []models.Itinerary
	var total int64

	query := database.DB.Model(&models.Itinerary{})
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		query = query.Where("status = ?", statusInt)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&itineraries)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      itineraries,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func AdminUpdateItineraryStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateItineraryStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Model(&models.Itinerary{}).Where("id = ?", id).Update("status", req.Status).Error; err != nil {
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
