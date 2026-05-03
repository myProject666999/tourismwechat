<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="路线详情"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <div v-if="route" class="detail-content">
        <van-swipe class="swipe" indicator-color="#fff">
          <van-swipe-item>
            <van-image
              :src="route.cover || getPlaceholderImage()"
              width="100%"
              height="200"
              fit="cover"
            />
          </van-swipe-item>
        </van-swipe>

        <div class="info-section">
          <h2 class="title">{{ route.name }}</h2>
          <div class="meta-row">
            <span v-if="route.difficulty" class="difficulty">
              <van-tag :type="getDifficultyType(route.difficulty)">
                {{ getDifficultyText(route.difficulty) }}
              </van-tag>
            </span>
          </div>
          <div class="path-row">
            <span class="point start">{{ route.start_point || '起点' }}</span>
            <van-icon name="arrow" />
            <span class="point end">{{ route.end_point || '终点' }}</span>
          </div>
          <div class="stats-row">
            <div class="stat-item" v-if="route.distance">
              <van-icon name="location-o" />
              <span>{{ route.distance }}km</span>
            </div>
            <div class="stat-item" v-if="route.duration">
              <van-icon name="clock-o" />
              <span>{{ route.duration }}</span>
            </div>
            <div class="stat-item">
              <van-icon name="eye-o" />
              <span>{{ route.view_count || 0 }} 浏览</span>
            </div>
          </div>
        </div>

        <div class="section" v-if="route.description">
          <div class="section-title">路线介绍</div>
          <div class="section-content">{{ route.description }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast } from 'vant'
import request from '@/utils/request'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const routeData = ref(null)

const onClickLeft = () => {
  router.back()
}

const getPlaceholderImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=scenic%20mountain%20road%20highway%20beautiful%20view&image_size=landscape_16_9'
}

const getDifficultyText = (difficulty) => {
  const map = { 1: '简单', 2: '中等', 3: '困难' }
  return map[difficulty] || '未知'
}

const getDifficultyType = (difficulty) => {
  const map = { 1: 'success', 2: 'warning', 3: 'danger' }
  return map[difficulty] || 'default'
}

const fetchDetail = async () => {
  loading.value = true
  try {
    const res = await request.get(`/routes/${route.params.id}`)
    routeData.value = res.data
  } catch (error) {
    console.error('Failed to fetch detail:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchDetail()
})
</script>

<style scoped>
.page-container {
  min-height: 100vh;
  background-color: #f7f8fa;
}

.page-content {
  padding-bottom: 20px;
}

.swipe {
  height: 200px;
}

.info-section {
  background: #fff;
  padding: 16px;
}

.title {
  font-size: 20px;
  font-weight: bold;
  color: #333;
  margin: 0 0 12px 0;
}

.meta-row {
  margin-bottom: 12px;
}

.path-row {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.path-row .point {
  font-size: 14px;
  color: #333;
  max-width: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.stats-row {
  display: flex;
  border-top: 1px solid #eee;
  padding-top: 12px;
}

.stat-item {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  color: #999;
}

.stat-item span {
  margin-left: 4px;
}

.section {
  background: #fff;
  margin-top: 12px;
  padding: 16px;
}

.section-title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin-bottom: 12px;
}

.section-content {
  font-size: 14px;
  color: #666;
  line-height: 1.8;
}
</style>
