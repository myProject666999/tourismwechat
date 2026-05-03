package controllers

import (
	"net/http"
	"strconv"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func GetRouteList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	difficulty := c.Query("difficulty")

	offset := (page - 1) * pageSize

	var routes []models.Route
	var total int64

	query := database.DB.Model(&models.Route{}).Where("status = ?", 1)
	if keyword != "" {
		query = query.Where("name LIKE ? OR start_point LIKE ? OR end_point LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}

	query.Count(&total)
	query.Order("view_count DESC, created_at DESC").Offset(offset).Limit(pageSize).Find(&routes)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      routes,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetRouteDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var route models.Route
	if err := database.DB.First(&route, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "路线不存在",
		})
		return
	}

	database.DB.Model(&route).Update("view_count", route.ViewCount+1)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    route,
	})
}

func AdminGetRouteList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var routes []models.Route
	var total int64

	query := database.DB.Model(&models.Route{})
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&routes)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      routes,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func CreateRoute(c *gin.Context) {
	var req models.CreateRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	route := &models.Route{
		Name:        req.Name,
		Description: req.Description,
		StartPoint:  req.StartPoint,
		EndPoint:    req.EndPoint,
		StartLat:    req.StartLat,
		StartLng:    req.StartLng,
		EndLat:      req.EndLat,
		EndLng:      req.EndLng,
		Distance:    req.Distance,
		Duration:    req.Duration,
		Cover:       req.Cover,
		WayPoints:   req.WayPoints,
		Difficulty:  req.Difficulty,
		Status:      req.Status,
	}

	if err := database.DB.Create(route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    route,
	})
}

func UpdateRoute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateRouteRequest
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
	if req.StartPoint != "" {
		updates["start_point"] = req.StartPoint
	}
	if req.EndPoint != "" {
		updates["end_point"] = req.EndPoint
	}
	if req.StartLat != 0 {
		updates["start_lat"] = req.StartLat
	}
	if req.StartLng != 0 {
		updates["start_lng"] = req.StartLng
	}
	if req.EndLat != 0 {
		updates["end_lat"] = req.EndLat
	}
	if req.EndLng != 0 {
		updates["end_lng"] = req.EndLng
	}
	if req.Distance != 0 {
		updates["distance"] = req.Distance
	}
	if req.Duration != "" {
		updates["duration"] = req.Duration
	}
	if req.Cover != "" {
		updates["cover"] = req.Cover
	}
	if req.WayPoints != "" {
		updates["way_points"] = req.WayPoints
	}
	if req.Difficulty != 0 {
		updates["difficulty"] = req.Difficulty
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}

	if err := database.DB.Model(&models.Route{}).Where("id = ?", id).Updates(updates).Error; err != nil {
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

func DeleteRoute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Delete(&models.Route{}, id).Error; err != nil {
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
