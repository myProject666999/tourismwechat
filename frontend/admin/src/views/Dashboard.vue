<template>
  <div class="dashboard">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background: #409EFF;">
              <el-icon :size="32"><User /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.userCount || 0 }}</div>
              <div class="stat-label">用户总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background: #67C23A;">
              <el-icon :size="32"><OfficeBuilding /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.hotelCount || 0 }}</div>
              <div class="stat-label">酒店总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background: #E6A23C;">
              <el-icon :size="32"><Ticket /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.bookingCount || 0 }}</div>
              <div class="stat-label">预订总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background: #F56C6C;">
              <el-icon :size="32"><List /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.itineraryCount || 0 }}</div>
              <div class="stat-label">行程总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="16">
        <el-card>
          <template #header>
            <span>最新预订</span>
          </template>
          <el-table :data="recentBookings" style="width: 100%">
            <el-table-column prop="id" label="ID" width="80" />
            <el-table-column prop="hotel_name" label="酒店" />
            <el-table-column prop="room_type" label="房型" />
            <el-table-column prop="check_in_date" label="入住日期" />
            <el-table-column prop="total_price" label="价格">
              <template #default="{ row }">
                ¥{{ row.total_price }}
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态">
              <template #default="{ row }">
                <el-tag :type="getStatusType(row.status)">
                  {{ getStatusText(row.status) }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <span>系统信息</span>
          </template>
          <div class="system-info">
            <div class="info-item">
              <span class="info-label">系统名称:</span>
              <span class="info-value">旅游小程序管理后台</span>
            </div>
            <div class="info-item">
              <span class="info-label">版本号:</span>
              <span class="info-value">1.0.0</span>
            </div>
            <div class="info-item">
              <span class="info-label">技术栈:</span>
              <span class="info-value">Vue 3 + Element Plus</span>
            </div>
            <div class="info-item">
              <span class="info-label">后端框架:</span>
              <span class="info-value">Gin + GORM</span>
            </div>
            <div class="info-item">
              <span class="info-label">当前时间:</span>
              <span class="info-value">{{ currentTime }}</span>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>热门景点 Top 5</span>
          </template>
          <el-table :data="topScenics" style="width: 100%">
            <el-table-column type="index" label="排名" width="60" />
            <el-table-column prop="name" label="景点名称" />
            <el-table-column prop="likes" label="点赞数" />
            <el-table-column prop="dislikes" label="踩数" />
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>最新公告</span>
          </template>
          <el-table :data="recentAnnouncements" style="width: 100%">
            <el-table-column prop="id" label="ID" width="60" />
            <el-table-column prop="title" label="标题" />
            <el-table-column prop="likes" label="点赞" width="70" />
            <el-table-column prop="dislikes" label="踩" width="60" />
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import dayjs from 'dayjs'
import request from '@/utils/request'

const stats = ref({
  userCount: 0,
  hotelCount: 0,
  bookingCount: 0,
  itineraryCount: 0
})

const recentBookings = ref([])
const recentAnnouncements = ref([])
const topScenics = ref([])

const currentTime = computed(() => {
  return dayjs().format('YYYY-MM-DD HH:mm:ss')
})

const getStatusType = (status) => {
  const types = {
    0: 'info',
    1: 'success',
    2: 'warning',
    3: 'danger'
  }
  return types[status] || 'info'
}

const getStatusText = (status) => {
  const texts = {
    0: '待确认',
    1: '已确认',
    2: '已取消',
    3: '已完成'
  }
  return texts[status] || '未知'
}

const fetchStats = async () => {
  try {
    const [usersRes, hotelsRes, bookingsRes, itinerariesRes] = await Promise.all([
      request.get('/admin/users'),
      request.get('/admin/hotels'),
      request.get('/admin/bookings'),
      request.get('/admin/itineraries')
    ])
    
    stats.value = {
      userCount: usersRes.data?.total || 0,
      hotelCount: hotelsRes.data?.total || 0,
      bookingCount: bookingsRes.data?.total || 0,
      itineraryCount: itinerariesRes.data?.total || 0
    }
    
    recentBookings.value = bookingsRes.data?.items?.slice(0, 5) || []
  } catch (error) {
    console.error('Failed to fetch stats:', error)
  }
}

const fetchRecentAnnouncements = async () => {
  try {
    const res = await request.get('/announcements')
    recentAnnouncements.value = res.data?.items?.slice(0, 5) || []
  } catch (error) {
    console.error('Failed to fetch announcements:', error)
  }
}

const fetchTopScenics = async () => {
  try {
    const res = await request.get('/scenics')
    const list = res.data?.items || []
    topScenics.value = list
      .sort((a, b) => (b.likes - b.dislikes) - (a.likes - a.dislikes))
      .slice(0, 5)
  } catch (error) {
    console.error('Failed to fetch scenics:', error)
  }
}

onMounted(() => {
  fetchStats()
  fetchRecentAnnouncements()
  fetchTopScenics()
  
  setInterval(() => {
    currentTime.value = dayjs().format('YYYY-MM-DD HH:mm:ss')
  }, 1000)
})
</script>

<style scoped>
.dashboard {
  padding: 0;
}

.stat-card {
  border-radius: 8px;
}

.stat-content {
  display: flex;
  align-items: center;
}

.stat-icon {
  width: 80px;
  height: 80px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.stat-info {
  margin-left: 20px;
}

.stat-value {
  font-size: 32px;
  font-weight: bold;
  color: #303133;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-top: 4px;
}

.system-info {
  padding: 10px 0;
}

.info-item {
  display: flex;
  padding: 12px 0;
  border-bottom: 1px solid #ebeef5;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  color: #909399;
  width: 90px;
  flex-shrink: 0;
}

.info-value {
  color: #303133;
  font-weight: 500;
}
</style>
