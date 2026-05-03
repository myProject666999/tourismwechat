package controllers

import (
	"net/http"
	"strconv"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func GetScenicList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	level := c.Query("level")

	offset := (page - 1) * pageSize

	var scenics []models.Scenic
	var total int64

	query := database.DB.Model(&models.Scenic{}).Where("status = ?", 1)
	if keyword != "" {
		query = query.Where("name LIKE ? OR address LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if level != "" {
		query = query.Where("level = ?", level)
	}

	query.Count(&total)
	query.Order("like_count DESC, view_count DESC").Offset(offset).Limit(pageSize).Find(&scenics)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      scenics,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetScenicDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var scenic models.Scenic
	if err := database.DB.First(&scenic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "景点不存在",
		})
		return
	}

	database.DB.Model(&scenic).Update("view_count", scenic.ViewCount+1)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    scenic,
	})
}

func LikeScenic(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var scenic models.Scenic
	if err := database.DB.First(&scenic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "景点不存在",
		})
		return
	}

	var existingLike models.ScenicLike
	result := database.DB.Where("user_id = ? AND scenic_id = ?", userID, id).First(&existingLike)

	if result.Error == nil {
		if existingLike.Type == 1 {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "已点赞",
			})
			return
		}

		database.DB.Model(&models.Scenic{}).Where("id = ?", id).Updates(map[string]interface{}{
			"like_count":    scenic.LikeCount + 1,
			"dislike_count": scenic.DislikeCount - 1,
		})
		database.DB.Model(&existingLike).Update("type", 1)
	} else {
		like := &models.ScenicLike{
			UserID:   userID,
			ScenicID: uint(id),
			Type:     1,
		}
		database.DB.Create(like)
		database.DB.Model(&models.Scenic{}).Where("id = ?", id).Update("like_count", scenic.LikeCount+1)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "点赞成功",
	})
}

func DislikeScenic(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var scenic models.Scenic
	if err := database.DB.First(&scenic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "景点不存在",
		})
		return
	}

	var existingLike models.ScenicLike
	result := database.DB.Where("user_id = ? AND scenic_id = ?", userID, id).First(&existingLike)

	if result.Error == nil {
		if existingLike.Type == 2 {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "已踩",
			})
			return
		}

		database.DB.Model(&models.Scenic{}).Where("id = ?", id).Updates(map[string]interface{}{
			"like_count":    scenic.LikeCount - 1,
			"dislike_count": scenic.DislikeCount + 1,
		})
		database.DB.Model(&existingLike).Update("type", 2)
	} else {
		like := &models.ScenicLike{
			UserID:   userID,
			ScenicID: uint(id),
			Type:     2,
		}
		database.DB.Create(like)
		database.DB.Model(&models.Scenic{}).Where("id = ?", id).Update("dislike_count", scenic.DislikeCount+1)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "操作成功",
	})
}

func AdminGetScenicList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var scenics []models.Scenic
	var total int64

	query := database.DB.Model(&models.Scenic{})
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&scenics)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      scenics,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func CreateScenic(c *gin.Context) {
	var req models.CreateScenicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	scenic := &models.Scenic{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Cover:       req.Cover,
		Images:      req.Images,
		Price:       req.Price,
		OpenTime:    req.OpenTime,
		Level:       req.Level,
		Status:      req.Status,
	}

	if err := database.DB.Create(scenic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    scenic,
	})
}

func UpdateScenic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateScenicRequest
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
	if req.Price != 0 {
		updates["price"] = req.Price
	}
	if req.OpenTime != "" {
		updates["open_time"] = req.OpenTime
	}
	if req.Level != "" {
		updates["level"] = req.Level
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}

	if err := database.DB.Model(&models.Scenic{}).Where("id = ?", id).Updates(updates).Error; err != nil {
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

func DeleteScenic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Delete(&models.Scenic{}, id).Error; err != nil {
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
