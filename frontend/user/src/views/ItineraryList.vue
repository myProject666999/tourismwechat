<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="行程管理"
        left-arrow
        @click-left="onClickLeft"
        :right-arrow="rightArrow"
        @click-right="goToCreate"
      />
      <template #right>
        <van-icon name="plus" size="20" />
      </template>
    </div>
    <div class="page-content" v-loading="loading">
      <div class="itinerary-list">
        <van-empty v-if="list.length === 0 && !loading" description="暂无行程">
          <template #description>
            <p>还没有创建行程</p>
            <van-button type="primary" size="small" @click="goToCreate">
              创建行程
            </van-button>
          </template>
        </van-empty>
        
        <div
          v-for="item in list"
          :key="item.id"
          class="itinerary-item"
          @click="goToDetail(item.id)"
        >
          <div class="itinerary-header">
            <h3 class="itinerary-title">{{ item.title }}</h3>
            <van-tag :type="getStatusType(item.status)">
              {{ getStatusText(item.status) }}
            </van-tag>
          </div>
          <div class="itinerary-dates" v-if="item.start_date || item.end_date">
            <span class="date">{{ item.start_date || '待定' }}</span>
            <van-icon name="arrow" />
            <span class="date">{{ item.end_date || '待定' }}</span>
          </div>
          <div class="itinerary-desc" v-if="item.description">
            {{ item.description.substring(0, 50) }}...
          </div>
        </div>
      </div>
    </div>

    <van-tabbar v-model="activeTab" route>
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
const loading = ref(false)
const list = ref([])
const rightArrow = ref(true)

const statusMap = {
  0: { text: '规划中', type: 'warning' },
  1: { text: '进行中', type: 'primary' },
  2: { text: '已完成', type: 'success' }
}

const onClickLeft = () => {
  router.back()
}

const goToCreate = () => {
  router.push('/itinerary/create')
}

const goToDetail = (id) => {
  router.push(`/itinerary/${id}`)
}

const getStatusText = (status) => {
  return statusMap[status]?.text || '未知'
}

const getStatusType = (status) => {
  return statusMap[status]?.type || 'default'
}

const fetchList = async () => {
  loading.value = true
  try {
    const res = await request.get('/user/itineraries')
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
  padding: 12px;
}

.itinerary-list {
  margin-top: 12px;
}

.itinerary-item {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 12px;
}

.itinerary-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.itinerary-title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.itinerary-dates {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  padding: 12px;
  background: #f7f8fa;
  border-radius: 8px;
  margin-bottom: 12px;
}

.itinerary-dates .date {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.itinerary-desc {
  font-size: 13px;
  color: #999;
  line-height: 1.6;
}
</style>
