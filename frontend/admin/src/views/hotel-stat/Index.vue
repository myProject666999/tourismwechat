<template>
  <div class="common-manage">
    <el-card>
      <template #header>
        <div class="page-header">
          <span class="page-title">{{ pageTitle }}</span>
        </div>
      </template>

      <el-row :gutter="20" class="stats-cards">
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-icon" style="background: #e6f7ff;">
              <el-icon :size="24" style="color: #1890ff;"><DataLine /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.totalBookings }}</div>
              <div class="stat-label">总预订数</div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-icon" style="background: #f6ffed;">
              <el-icon :size="24" style="color: #52c41a;"><Wallet /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">¥{{ stats.totalRevenue }}</div>
              <div class="stat-label">总收入</div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-icon" style="background: #fff7e6;">
              <el-icon :size="24" style="color: #faad14;"><Building /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.totalHotels }}</div>
              <div class="stat-label">酒店总数</div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-icon" style="background: #f9f0ff;">
              <el-icon :size="24" style="color: #722ed1;"><Odometer /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.avgOccupancy }}%</div>
              <div class="stat-label">平均入住率</div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="请输入关键词" clearable @keyup.enter="handleSearch" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="hotel_name" label="酒店" min-width="150" />
        <el-table-column prop="booking_count" label="预订数" width="100" sortable />
        <el-table-column prop="total_revenue" label="总收入" width="120" sortable>
          <template #default="{ row }">
            <span class="revenue">¥{{ row.total_revenue || 0 }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="room_occupancy" label="入住率" width="100" sortable>
          <template #default="{ row }">
            <el-progress 
              :percentage="row.room_occupancy || 0" 
              :color="getOccupancyColor(row.room_occupancy)"
              :stroke-width="10"
            />
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="浏览数" width="100" />
        <el-table-column prop="avg_rating" label="评分" width="100">
          <template #default="{ row }">
            <el-rate 
              :model-value="row.avg_rating || 0" 
              disabled 
              :max="5"
              show-score
              text-color="#ff9900"
            />
          </template>
        </el-table-column>
        <el-table-column prop="updated_at" label="更新时间" width="180" />
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleView(row)">
              <el-icon><View /></el-icon>
              详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="fetchData"
        @current-change="fetchData"
      />
    </el-card>

    <el-dialog v-model="detailVisible" title="酒店统计详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="酒店名称">
          {{ currentHotel.hotel_name || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="酒店ID">
          {{ currentHotel.hotel_id || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="预订总数">
          <el-tag type="primary">{{ currentHotel.booking_count || 0 }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="总收入">
          <span class="revenue">¥{{ currentHotel.total_revenue || 0 }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="入住率">
          <el-progress 
            :percentage="currentHotel.room_occupancy || 0" 
            :color="getOccupancyColor(currentHotel.room_occupancy)"
          />
        </el-descriptions-item>
        <el-descriptions-item label="浏览数">
          {{ currentHotel.view_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item label="平均评分">
          <el-rate 
            :model-value="currentHotel.avg_rating || 0" 
            disabled 
            :max="5"
            show-score
          />
        </el-descriptions-item>
        <el-descriptions-item label="统计周期">
          近30天
        </el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, getCurrentInstance } from 'vue'
import { ElMessage } from 'element-plus'
import request from '@/utils/request'

const loading = ref(false)
const detailVisible = ref(false)
const currentHotel = ref({})

const tableData = ref([])
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const searchForm = reactive({
  keyword: ''
})

const stats = reactive({
  totalBookings: 0,
  totalRevenue: 0,
  totalHotels: 0,
  avgOccupancy: 0
})

const pageTitle = computed(() => '酒店统计管理')

const getOccupancyColor = (percentage) => {
  if (percentage >= 80) return '#52c41a'
  if (percentage >= 50) return '#1890ff'
  if (percentage >= 30) return '#faad14'
  return '#ff4d4f'
}

const fetchData = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword
    }
    const res = await request.get('/admin/hotel-stats', { params })
    tableData.value = res.data?.items || []
    pagination.total = res.data?.total || 0
    
    if (res.data?.summary) {
      Object.assign(stats, res.data.summary)
    }
  } catch (error) {
    console.error('Failed to fetch data:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  fetchData()
}

const handleReset = () => {
  searchForm.keyword = ''
  handleSearch()
}

const handleView = (row) => {
  currentHotel.value = row
  detailVisible.value = true
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.stats-cards {
  margin-bottom: 20px;
}

.stat-card {
  display: flex;
  align-items: center;
  padding: 10px !important;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.revenue {
  color: #67c23a;
  font-weight: bold;
  font-size: 16px;
}
</style>
