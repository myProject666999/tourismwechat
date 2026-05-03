package controllers

import (
	"net/http"
	"strconv"
	"time"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func GetDailyRecommendations(c *gin.Context) {
	now := time.Now()

	var recommendations []models.DailyRecommend
	query := database.DB.Model(&models.DailyRecommend{}).
		Where("status = ?", 1).
		Where("start_date <= ?", now).
		Order("priority DESC, created_at DESC")

	if err := query.Find(&recommendations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    recommendations,
	})
}

func AdminGetDailyRecommendList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var recommendations []models.DailyRecommend
	var total int64

	query := database.DB.Model(&models.DailyRecommend{})
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		query = query.Where("status = ?", statusInt)
	}

	query.Count(&total)
	query.Order("priority DESC, created_at DESC").Offset(offset).Limit(pageSize).Find(&recommendations)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      recommendations,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func CreateDailyRecommend(c *gin.Context) {
	var req models.CreateDailyRecommendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	recommendation := &models.DailyRecommend{
		Title:         req.Title,
		Description:   req.Description,
		Cover:         req.Cover,
		RecommendType: req.RecommendType,
		TargetID:      req.TargetID,
		TargetName:    req.TargetName,
		Priority:      req.Priority,
		Status:        req.Status,
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err == nil {
			recommendation.StartDate = startDate
		}
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err == nil {
			recommendation.EndDate = endDate
		}
	}

	if err := database.DB.Create(recommendation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    recommendation,
	})
}

func UpdateDailyRecommend(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateDailyRecommendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Cover != "" {
		updates["cover"] = req.Cover
	}
	if req.RecommendType != "" {
		updates["recommend_type"] = req.RecommendType
	}
	if req.TargetID != 0 {
		updates["target_id"] = req.TargetID
	}
	if req.TargetName != "" {
		updates["target_name"] = req.TargetName
	}
	if req.Priority != 0 {
		updates["priority"] = req.Priority
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}
	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err == nil {
			updates["start_date"] = startDate
		}
	}
	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err == nil {
			updates["end_date"] = endDate
		}
	}

	if err := database.DB.Model(&models.DailyRecommend{}).Where("id = ?", id).Updates(updates).Error; err != nil {
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

func DeleteDailyRecommend(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Delete(&models.DailyRecommend{}, id).Error; err != nil {
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
