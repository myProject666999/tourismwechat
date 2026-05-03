<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="预订管理"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <van-tabs v-model="activeTab" @change="onTabChange">
        <van-tab title="全部"></van-tab>
        <van-tab title="待确认"></van-tab>
        <van-tab title="已确认"></van-tab>
        <van-tab title="已完成"></van-tab>
        <van-tab title="已取消"></van-tab>
      </van-tabs>

      <div class="booking-list">
        <van-empty v-if="list.length === 0 && !loading" description="暂无预订记录" />
        
        <div
          v-for="item in list"
          :key="item.id"
          class="booking-item"
          @click="goToDetail(item.id)"
        >
          <div class="booking-header">
            <span class="booking-number">预订编号: {{ item.id }}</span>
            <van-tag :type="getStatusType(item.status)">
              {{ getStatusText(item.status) }}
            </van-tag>
          </div>
          <div class="booking-content">
            <div class="hotel-info">
              <h3 class="hotel-name">{{ item.hotel_name || '酒店预订' }}</h3>
              <div class="room-info">
                <span>{{ item.room_type || '房型' }}</span>
              </div>
            </div>
            <div class="booking-dates">
              <div class="date-item">
                <span class="date-label">入住</span>
                <span class="date-value">{{ item.check_in_date || '-' }}</span>
              </div>
              <van-icon name="arrow" />
              <div class="date-item">
                <span class="date-label">退房</span>
                <span class="date-value">{{ item.check_out_date || '-' }}</span>
              </div>
            </div>
            <div class="booking-footer">
              <span class="total-label">总价</span>
              <span class="total-price">
                <span class="unit">¥</span>
                <span class="value">{{ item.total_price || 0 }}</span>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <van-tabbar v-model="activeNavTab" route>
      <van-tabbar-item replace to="/home" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item replace to="/scenic" icon="photo-o">景点</van-tabbar-item>
      <van-tabbar-item replace to="/hotel" icon="shop-o">酒店</van-tabbar-item>
      <van-tabbar-item replace to="/profile" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import request from '@/utils/request'

const router = useRouter()

const activeTab = ref(0)
const activeNavTab = ref(0)
const loading = ref(false)
const list = ref([])

const statusMap = {
  0: { text: '待确认', type: 'warning' },
  1: { text: '已确认', type: 'primary' },
  2: { text: '已完成', type: 'success' },
  3: { text: '已取消', type: 'default' }
}

const onClickLeft = () => {
  router.back()
}

const goToDetail = (id) => {
  router.push(`/booking/${id}`)
}

const getStatusText = (status) => {
  return statusMap[status]?.text || '未知'
}

const getStatusType = (status) => {
  return statusMap[status]?.type || 'default'
}

const onTabChange = () => {
  fetchList()
}

const fetchList = async () => {
  loading.value = true
  try {
    const params = {}
    if (activeTab.value > 0) {
      params.status = activeTab.value - 1
    }
    const res = await request.get('/user/bookings', { params })
    list.value = res.data?.items || []
  } catch (error) {
    console.error('Failed to fetch list:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchList()
})
</script>

<style scoped>
.page-container {
  min-height: 100vh;
  background-color: #f7f8fa;
  padding-bottom: 50px;
}

.page-content {
  padding: 0;
}

.booking-list {
  padding: 12px;
}

.booking-item {
  background: #fff;
  border-radius: 8px;
  margin-bottom: 12px;
  overflow: hidden;
}

.booking-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #fafafa;
  border-bottom: 1px solid #eee;
}

.booking-number {
  font-size: 12px;
  color: #999;
}

.booking-content {
  padding: 12px 16px;
}

.hotel-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0 0 8px 0;
}

.room-info {
  font-size: 13px;
  color: #666;
  margin-bottom: 12px;
}

.booking-dates {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px;
  background: #f7f8fa;
  border-radius: 8px;
  margin-bottom: 12px;
}

.date-item {
  text-align: center;
}

.date-label {
  display: block;
  font-size: 12px;
  color: #999;
  margin-bottom: 4px;
}

.date-value {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.booking-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid #eee;
}

.total-label {
  font-size: 13px;
  color: #666;
}

.total-price .unit {
  font-size: 14px;
  color: #ff6b6b;
}

.total-price .value {
  font-size: 20px;
  font-weight: bold;
  color: #ff6b6b;
}
</style>
