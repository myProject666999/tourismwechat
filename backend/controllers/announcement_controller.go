package controllers

import (
	"net/http"
	"strconv"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
)

func GetAnnouncementList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")

	offset := (page - 1) * pageSize

	var announcements []models.Announcement
	var total int64

	query := database.DB.Model(&models.Announcement{}).Where("status = ?", 1)
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&announcements)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      announcements,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func GetAnnouncementDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var announcement models.Announcement
	if err := database.DB.First(&announcement, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "公告不存在",
		})
		return
	}

	database.DB.Model(&announcement).Update("view_count", announcement.ViewCount+1)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    announcement,
	})
}

func LikeAnnouncement(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var announcement models.Announcement
	if err := database.DB.First(&announcement, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "公告不存在",
		})
		return
	}

	var existingLike models.AnnouncementLike
	result := database.DB.Where("user_id = ? AND announcement_id = ?", userID, id).First(&existingLike)

	if result.Error == nil {
		if existingLike.Type == 1 {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "已点赞",
			})
			return
		}

		database.DB.Model(&models.Announcement{}).Where("id = ?", id).Updates(map[string]interface{}{
			"like_count":    announcement.LikeCount + 1,
			"dislike_count": announcement.DislikeCount - 1,
		})
		database.DB.Model(&existingLike).Update("type", 1)
	} else {
		like := &models.AnnouncementLike{
			UserID:         userID,
			AnnouncementID: uint(id),
			Type:           1,
		}
		database.DB.Create(like)
		database.DB.Model(&models.Announcement{}).Where("id = ?", id).Update("like_count", announcement.LikeCount+1)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "点赞成功",
	})
}

func DislikeAnnouncement(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var announcement models.Announcement
	if err := database.DB.First(&announcement, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "公告不存在",
		})
		return
	}

	var existingLike models.AnnouncementLike
	result := database.DB.Where("user_id = ? AND announcement_id = ?", userID, id).First(&existingLike)

	if result.Error == nil {
		if existingLike.Type == 2 {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "已踩",
			})
			return
		}

		database.DB.Model(&models.Announcement{}).Where("id = ?", id).Updates(map[string]interface{}{
			"like_count":    announcement.LikeCount - 1,
			"dislike_count": announcement.DislikeCount + 1,
		})
		database.DB.Model(&existingLike).Update("type", 2)
	} else {
		like := &models.AnnouncementLike{
			UserID:         userID,
			AnnouncementID: uint(id),
			Type:           2,
		}
		database.DB.Create(like)
		database.DB.Model(&models.Announcement{}).Where("id = ?", id).Update("dislike_count", announcement.DislikeCount+1)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "操作成功",
	})
}

func AdminGetAnnouncementList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	offset := (page - 1) * pageSize

	var announcements []models.Announcement
	var total int64

	query := database.DB.Model(&models.Announcement{})
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&announcements)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":      announcements,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func CreateAnnouncement(c *gin.Context) {
	var req models.CreateAnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	announcement := &models.Announcement{
		Title:   req.Title,
		Content: req.Content,
		Cover:   req.Cover,
		Status:  req.Status,
	}

	if err := database.DB.Create(announcement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    announcement,
	})
}

func UpdateAnnouncement(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req models.UpdateAnnouncementRequest
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
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.Cover != "" {
		updates["cover"] = req.Cover
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}

	if err := database.DB.Model(&models.Announcement{}).Where("id = ?", id).Updates(updates).Error; err != nil {
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

func DeleteAnnouncement(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	if err := database.DB.Delete(&models.Announcement{}, id).Error; err != nil {
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
