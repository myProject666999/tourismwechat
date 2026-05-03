<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="景点详情"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <div v-if="scenic" class="detail-content">
        <van-swipe class="swipe" indicator-color="#fff">
          <van-swipe-item
            v-for="(img, index) in images"
            :key="index"
          >
            <van-image
              :src="img"
              width="100%"
              height="200"
              fit="cover"
            />
          </van-swipe-item>
        </van-swipe>

        <div class="info-section">
          <h2 class="title">{{ scenic.name }}</h2>
          <div class="tags" v-if="scenic.level">
            <van-tag type="primary">{{ scenic.level }}</van-tag>
          </div>
          <div class="price-row" v-if="scenic.price">
            <span class="price-label">门票</span>
            <span class="price">¥{{ scenic.price }}</span>
          </div>
          <div class="location-row">
            <van-icon name="location-o" />
            <span>{{ scenic.address || '暂无地址' }}</span>
          </div>
          <div class="stats-row">
            <div class="stat-item">
              <van-icon name="eye-o" />
              <span>{{ scenic.view_count || 0 }} 浏览</span>
            </div>
            <div class="stat-item">
              <van-icon name="like-o" />
              <span>{{ scenic.like_count || 0 }} 点赞</span>
            </div>
            <div class="stat-item">
              <van-icon name="dislike-o" />
              <span>{{ scenic.dislike_count || 0 }} 踩</span>
            </div>
          </div>
        </div>

        <div class="section" v-if="scenic.description">
          <div class="section-title">景点介绍</div>
          <div class="section-content">{{ scenic.description }}</div>
        </div>

        <div class="section" v-if="scenic.open_time">
          <div class="section-title">开放时间</div>
          <div class="section-content">{{ scenic.open_time }}</div>
        </div>
      </div>
    </div>

    <div class="bottom-bar">
      <van-button
        type="primary"
        size="large"
        :icon="isLiked ? 'good-job' : 'like-o'"
        :class="{ 'van-button--primary--plain': isLiked }"
        @click="handleLike"
      >
        点赞
      </van-button>
      <van-button
        type="danger"
        size="large"
        :icon="isDisliked ? 'cross' : 'dislike-o'"
        :class="{ 'van-button--danger--plain': isDisliked }"
        @click="handleDislike"
      >
        踩
      </van-button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast } from 'vant'
import request from '@/utils/request'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const scenic = ref(null)
const isLiked = ref(false)
const isDisliked = ref(false)

const images = computed(() => {
  if (!scenic.value) return []
  const result = []
  if (scenic.value.cover) {
    result.push(scenic.value.cover)
  }
  if (scenic.value.images) {
    try {
      const imgs = JSON.parse(scenic.value.images)
      if (Array.isArray(imgs)) {
        result.push(...imgs)
      }
    } catch (e) {
      // ignore
    }
  }
  if (result.length === 0) {
    result.push(getPlaceholderImage())
  }
  return result
})

const onClickLeft = () => {
  router.back()
}

const getPlaceholderImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=beautiful%20tourist%20attraction%20landscape%20photo&image_size=landscape_16_9'
}

const fetchDetail = async () => {
  loading.value = true
  try {
    const res = await request.get(`/scenics/${route.params.id}`)
    scenic.value = res.data
  } catch (error) {
    console.error('Failed to fetch detail:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const handleLike = async () => {
  try {
    await request.post(`/user/scenics/${route.params.id}/like`)
    isLiked.value = true
    isDisliked.value = false
    scenic.value.like_count = (scenic.value.like_count || 0) + 1
    showToast('点赞成功')
  } catch (error) {
    console.error('Failed to like:', error)
    showToast('操作失败')
  }
}

const handleDislike = async () => {
  try {
    await request.post(`/user/scenics/${route.params.id}/dislike`)
    isDisliked.value = true
    isLiked.value = false
    scenic.value.dislike_count = (scenic.value.dislike_count || 0) + 1
    showToast('操作成功')
  } catch (error) {
    console.error('Failed to dislike:', error)
    showToast('操作失败')
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
  padding-bottom: 60px;
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

.tags {
  margin-bottom: 12px;
}

.price-row {
  display: flex;
  align-items: baseline;
  margin-bottom: 12px;
}

.price-label {
  font-size: 14px;
  color: #999;
  margin-right: 8px;
}

.price {
  font-size: 24px;
  font-weight: bold;
  color: #ff6b6b;
}

.location-row {
  display: flex;
  align-items: center;
  font-size: 14px;
  color: #666;
  margin-bottom: 12px;
}

.location-row span {
  margin-left: 6px;
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

.bottom-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  padding: 10px 16px;
  background: #fff;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.05);
  gap: 12px;
}

.bottom-bar .van-button {
  flex: 1;
}
</style>
