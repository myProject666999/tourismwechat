package database

import (
	"time"
	"tourismwechat/config"
	"tourismwechat/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	sqlite "github.com/glebarez/sqlite"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	
	DB, err = gorm.Open(sqlite.Open(config.AppConfig.Database.Sqlite.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	err = migrateTables()
	if err != nil {
		return err
	}

	err = createDefaultAdmin()
	if err != nil {
		return err
	}

	err = createTestData()
	if err != nil {
		return err
	}

	return nil
}

func migrateTables() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Announcement{},
		&models.AnnouncementLike{},
		&models.Scenic{},
		&models.ScenicLike{},
		&models.Hotel{},
		&models.HotelRoom{},
		&models.Route{},
		&models.Booking{},
		&models.Itinerary{},
		&models.ItineraryItem{},
		&models.DailyRecommend{},
		&models.HotelStat{},
	)
}

func createDefaultAdmin() error {
	var count int64
	DB.Model(&models.Admin{}).Count(&count)
	if count > 0 {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.AppConfig.Admin.DefaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &models.Admin{
		Username: config.AppConfig.Admin.DefaultUsername,
		Password: string(hashedPassword),
		Nickname: "超级管理员",
		Role:     0,
		Status:   1,
	}

	return DB.Create(admin).Error
}

func createTestData() error {
	var scenicCount int64
	DB.Model(&models.Scenic{}).Count(&scenicCount)
	if scenicCount > 0 {
		return nil
	}

	scenics := []models.Scenic{
		{
			Name:        "长城",
			Description: "万里长城是中国古代的军事防御工程，是世界文化遗产之一。",
			Address:     "北京市延庆区八达岭镇",
			Price:       40,
			Level:       "5A",
			OpenTime:    "06:30-19:00",
			ViewCount:   12580,
			LikeCount:   3560,
			DislikeCount: 120,
			Status:      1,
		},
		{
			Name:        "故宫博物院",
			Description: "故宫博物院是中国最大的古代文化艺术博物馆，建立于1925年10月10日。",
			Address:     "北京市东城区景山前街4号",
			Price:       60,
			Level:       "5A",
			OpenTime:    "08:30-17:00",
			ViewCount:   25600,
			LikeCount:   8900,
			DislikeCount: 230,
			Status:      1,
		},
		{
			Name:        "西湖",
			Description: "西湖位于浙江省杭州市西湖区，景区总面积49平方千米。",
			Address:     "浙江省杭州市西湖区",
			Price:       0,
			Level:       "5A",
			OpenTime:    "全天开放",
			ViewCount:   35800,
			LikeCount:   12500,
			DislikeCount: 180,
			Status:      1,
		},
		{
			Name:        "黄山",
			Description: "黄山位于安徽省南部黄山市境内，有72峰，主峰莲花峰海拔1864米。",
			Address:     "安徽省黄山市黄山区",
			Price:       190,
			Level:       "5A",
			OpenTime:    "06:00-17:30",
			ViewCount:   18900,
			LikeCount:   7800,
			DislikeCount: 150,
			Status:      1,
		},
	}

	for _, s := range scenics {
		if err := DB.Create(&s).Error; err != nil {
			return err
		}
	}

	hotels := []models.Hotel{
		{
			Name:        "北京国际饭店",
			Description: "北京国际饭店是一家五星级豪华酒店，提供高品质的住宿服务。",
			Address:     "北京市东城区建国门内大街9号",
			Star:        5,
			Price:       580,
			Status:      1,
		},
		{
			Name:        "杭州西湖国宾馆",
			Description: "杭州西湖国宾馆位于西湖西岸，环境优美，服务一流。",
			Address:     "浙江省杭州市西湖区杨公堤18号",
			Star:        5,
			Price:       880,
			Status:      1,
		},
		{
			Name:        "上海外滩华尔道夫酒店",
			Description: "上海外滩华尔道夫酒店位于上海外滩，是一家奢华五星级酒店。",
			Address:     "上海市黄浦区中山东一路2号",
			Star:        5,
			Price:       1280,
			Status:      1,
		},
	}

	for _, h := range hotels {
		if err := DB.Create(&h).Error; err != nil {
			return err
		}
	}

	announcements := []models.Announcement{
		{
			Title:   "国庆黄金周旅游通知",
			Content: "国庆黄金周即将到来，各景点将迎来旅游高峰。请游客提前规划行程，注意出行安全。",
			ViewCount: 15600,
			LikeCount:   2800,
			DislikeCount: 85,
			Status:      1,
		},
		{
			Title:   "冬季旅游特惠活动",
			Content: "为回馈广大游客，本平台推出冬季旅游特惠活动。活动期间，部分景点门票享受5折优惠。",
			ViewCount: 8900,
			LikeCount:   3200,
			DislikeCount: 45,
			Status:      1,
		},
		{
			Title:   "景区开放时间调整通知",
			Content: "因季节变化，部分景区开放时间将进行调整。请游客合理安排出行时间。",
			ViewCount: 6500,
			LikeCount:   1200,
			DislikeCount: 200,
			Status:      1,
		},
	}

	for _, a := range announcements {
		if err := DB.Create(&a).Error; err != nil {
			return err
		}
	}

	routes := []models.Route{
		{
			Name:        "北京经典三日游",
			Description: "游览北京的经典景点，包括故宫、长城、颐和园、天坛等。",
			StartPoint:  "北京市区",
			EndPoint:    "北京市区",
			Duration:    "3天",
			Distance:    200,
			Difficulty:  2,
			ViewCount:   12500,
			Status:      1,
		},
		{
			Name:        "杭州西湖两日游",
			Description: "漫步西湖边，感受江南水乡的诗情画意。",
			StartPoint:  "杭州市区",
			EndPoint:    "杭州市区",
			Duration:    "2天",
			Distance:    100,
			Difficulty:  1,
			ViewCount:   9800,
			Status:      1,
		},
		{
			Name:        "黄山深度游",
			Description: "征服黄山，感受壮丽景色。游览黄山四绝：奇松、怪石、云海、温泉。",
			StartPoint:  "黄山市",
			EndPoint:    "黄山市",
			Duration:    "2天",
			Distance:    80,
			Difficulty:  3,
			ViewCount:   7600,
			Status:      1,
		},
	}

	for _, r := range routes {
		if err := DB.Create(&r).Error; err != nil {
			return err
		}
	}

	now := time.Now()
	recommendations := []models.DailyRecommend{
		{
			Title:         "长城金秋游",
			Description:   "秋季是游览长城的最佳时节，金黄的树叶与古老的城墙交相辉映。",
			Cover:         "",
			RecommendType: "scenic",
			TargetID:      1,
			TargetName:    "长城",
			Priority:      1,
			Status:        1,
			StartDate:     now.AddDate(0, 0, -10),
		},
		{
			Title:         "西湖荷花盛开季",
			Description:   "夏季西湖荷花盛开，接天莲叶无穷碧，映日荷花别样红。",
			Cover:         "",
			RecommendType: "scenic",
			TargetID:      3,
			TargetName:    "西湖",
			Priority:      2,
			Status:        1,
			StartDate:     now.AddDate(0, 0, -5),
		},
		{
			Title:         "北京国际饭店特惠",
			Description:   "入住北京国际饭店，享受五星级服务，体验宾至如归的感觉。",
			Cover:         "",
			RecommendType: "hotel",
			TargetID:      1,
			TargetName:    "北京国际饭店",
			Priority:      3,
			Status:        1,
			StartDate:     now.AddDate(0, 0, -3),
		},
	}

	for _, r := range recommendations {
		if err := DB.Create(&r).Error; err != nil {
			return err
		}
	}

	return nil
}
