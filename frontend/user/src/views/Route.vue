<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="自驾路线"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <van-search
        v-model="searchForm.keyword"
        placeholder="搜索路线"
        @search="handleSearch"
      />
      
      <div class="route-list">
        <van-empty v-if="list.length === 0 && !loading" description="暂无数据" />
        
        <div
          v-for="item in list"
          :key="item.id"
          class="route-item"
          @click="goToDetail(item.id)"
        >
          <div class="route-cover">
            <van-image
              :src="item.cover || getPlaceholderImage()"
              width="120"
              height="80"
              fit="cover"
            />
            <van-tag v-if="item.difficulty" type="warning" size="mini">
              {{ getDifficultyText(item.difficulty) }}
            </van-tag>
          </div>
          <div class="route-info">
            <h3 class="route-name">{{ item.name }}</h3>
            <div class="route-meta" v-if="item.start_point || item.end_point">
              <span class="start">{{ item.start_point || '起点' }}</span>
              <van-icon name="arrow" />
              <span class="end">{{ item.end_point || '终点' }}</span>
            </div>
            <div class="route-stats">
              <span v-if="item.distance">
                <van-icon name="location-o" />
                {{ item.distance }}km
              </span>
              <span v-if="item.duration">
                <van-icon name="clock-o" />
                {{ item.duration }}
              </span>
              <span>
                <van-icon name="eye-o" />
                {{ item.view_count || 0 }}
              </span>
            </div>
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

const searchForm = ref({
  keyword: ''
})

const onClickLeft = () => {
  router.back()
}

const goToDetail = (id) => {
  router.push(`/route/${id}`)
}

const getPlaceholderImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=beautiful%20road%20trip%20landscape%20mountain%20highway&image_size=square'
}

const getDifficultyText = (difficulty) => {
  const map = { 1: '简单', 2: '中等', 3: '困难' }
  return map[difficulty] || '未知'
}

const fetchList = async () => {
  loading.value = true
  try {
    const res = await request.get('/routes', {
      params: {
        keyword: searchForm.value.keyword
      }
    })
    list.value = res.data?.items || []
  } catch (error) {
    console.error('Failed to fetch list:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  fetchList()
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

.route-list {
  margin-top: 12px;
}

.route-item {
  display: flex;
  background: #fff;
  border-radius: 8px;
  padding: 12px;
  margin-bottom: 12px;
}

.route-cover {
  position: relative;
  flex-shrink: 0;
}

.route-cover .van-tag {
  position: absolute;
  top: 4px;
  left: 4px;
}

.route-info {
  flex: 1;
  margin-left: 12px;
  display: flex;
  flex-direction: column;
}

.route-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0 0 8px 0;
}

.route-meta {
  display: flex;
  align-items: center;
  font-size: 13px;
  color: #666;
  margin-bottom: 8px;
}

.route-meta .start,
.route-meta .end {
  max-width: 80px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.route-stats {
  display: flex;
  gap: 16px;
  font-size: 12px;
  color: #999;
  margin-top: auto;
}
</style>
