package controllers

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"
	"tourismwechat/database"
	"tourismwechat/models"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ExportUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "用户名")
	f.SetCellValue(sheet, "C1", "昵称")
	f.SetCellValue(sheet, "D1", "手机号")
	f.SetCellValue(sheet, "E1", "邮箱")
	f.SetCellValue(sheet, "F1", "状态")
	f.SetCellValue(sheet, "G1", "注册时间")

	for i, user := range users {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), user.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), user.Username)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), user.Nickname)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), user.Phone)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), user.Email)
		status := "正常"
		if user.Status == 0 {
			status = "禁用"
		}
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), status)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), user.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=users_"+time.Now().Format("20060102")+".xlsx")

	buf, err := f.WriteToBuffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "导出失败",
		})
		return
	}

	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}

func ImportUsers(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请上传文件",
		})
		return
	}
	defer file.Close()

	if header.Size == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文件为空",
		})
		return
	}

	tempFile, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "读取文件失败",
		})
		return
	}

	f, err := excelize.OpenReader(bytes.NewReader(tempFile))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文件格式错误",
		})
		return
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "读取表格失败",
		})
		return
	}

	var successCount, failCount int
	var errors []string

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 3 {
			failCount++
			errors = append(errors, "第"+strconv.Itoa(i+1)+"行数据不完整")
			continue
		}

		username := row[1]
		password := row[2]

		if username == "" || password == "" {
			failCount++
			errors = append(errors, "第"+strconv.Itoa(i+1)+"行用户名或密码为空")
			continue
		}

		var existingUser models.User
		if database.DB.Where("username = ?", username).First(&existingUser).Error == nil {
			failCount++
			errors = append(errors, "第"+strconv.Itoa(i+1)+"行用户名已存在")
			continue
		}

		nickname := ""
		if len(row) > 3 {
			nickname = row[3]
		}
		phone := ""
		if len(row) > 4 {
			phone = row[4]
		}
		email := ""
		if len(row) > 5 {
			email = row[5]
		}

		user := &models.User{
			Username: username,
			Password: password,
			Nickname: nickname,
			Phone:    phone,
			Email:    email,
			Status:   1,
		}

		if err := database.DB.Create(user).Error; err != nil {
			failCount++
			errors = append(errors, "第"+strconv.Itoa(i+1)+"行创建失败")
			continue
		}

		successCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "导入完成",
		"data": gin.H{
			"success_count": successCount,
			"fail_count":    failCount,
			"errors":        errors,
		},
	})
}

func ExportHotels(c *gin.Context) {
	var hotels []models.Hotel
	database.DB.Find(&hotels)

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "酒店名称")
	f.SetCellValue(sheet, "C1", "星级")
	f.SetCellValue(sheet, "D1", "价格")
	f.SetCellValue(sheet, "E1", "地址")
	f.SetCellValue(sheet, "F1", "状态")
	f.SetCellValue(sheet, "G1", "创建时间")

	for i, hotel := range hotels {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), hotel.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), hotel.Name)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), hotel.Star)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), hotel.Price)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), hotel.Address)
		status := "正常"
		if hotel.Status == 0 {
			status = "禁用"
		}
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), status)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), hotel.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=hotels_"+time.Now().Format("20060102")+".xlsx")

	buf, err := f.WriteToBuffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "导出失败",
		})
		return
	}

	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}

func ExportBookings(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	status := c.Query("status")

	query := database.DB.Model(&models.Booking{})
	if startDate != "" {
		sDate, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("created_at >= ?", sDate)
		}
	}
	if endDate != "" {
		eDate, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("created_at <= ?", eDate)
		}
	}
	if status != "" {
		statusInt, _ := strconv.Atoi(status)
		query = query.Where("status = ?", statusInt)
	}

	var bookings []models.Booking
	query.Find(&bookings)

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "酒店ID")
	f.SetCellValue(sheet, "C1", "房型")
	f.SetCellValue(sheet, "D1", "入住日期")
	f.SetCellValue(sheet, "E1", "退房日期")
	f.SetCellValue(sheet, "F1", "总金额")
	f.SetCellValue(sheet, "G1", "联系人")
	f.SetCellValue(sheet, "H1", "联系电话")
	f.SetCellValue(sheet, "I1", "状态")
	f.SetCellValue(sheet, "J1", "创建时间")

	statusMap := map[int]string{
		0: "待确认",
		1: "已确认",
		2: "已完成",
		3: "已取消",
	}

	for i, booking := range bookings {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), booking.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), booking.HotelID)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), booking.RoomType)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), booking.CheckInDate.Format("2006-01-02"))
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), booking.CheckOutDate.Format("2006-01-02"))
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), booking.TotalPrice)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), booking.ContactName)
		f.SetCellValue(sheet, "H"+strconv.Itoa(row), booking.ContactPhone)
		f.SetCellValue(sheet, "I"+strconv.Itoa(row), statusMap[booking.Status])
		f.SetCellValue(sheet, "J"+strconv.Itoa(row), booking.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=bookings_"+time.Now().Format("20060102")+".xlsx")

	buf, err := f.WriteToBuffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "导出失败",
		})
		return
	}

	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}
