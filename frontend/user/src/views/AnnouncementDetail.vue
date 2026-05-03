<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="公告详情"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <div v-if="announcement" class="detail-content">
        <h2 class="detail-title">{{ announcement.title }}</h2>
        <div class="detail-meta">
          <van-icon name="eye-o" />
          <span>{{ announcement.view_count || 0 }} 浏览</span>
          <van-icon name="clock-o" />
          <span>{{ announcement.created_at }}</span>
        </div>
        <div class="detail-body" v-if="announcement.content">
          {{ announcement.content }}
        </div>
        <div class="action-bar">
          <van-button
            type="primary"
            size="large"
            :icon="isLiked ? 'good-job' : 'like-o'"
            :class="{ 'van-button--primary--plain': isLiked }"
            @click="handleLike"
          >
            点赞 ({{ announcement.like_count || 0 }})
          </van-button>
          <van-button
            type="danger"
            size="large"
            :icon="isDisliked ? 'cross' : 'dislike-o'"
            :class="{ 'van-button--danger--plain': isDisliked }"
            @click="handleDislike"
          >
            踩 ({{ announcement.dislike_count || 0 }})
          </van-button>
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
const announcement = ref(null)
const isLiked = ref(false)
const isDisliked = ref(false)

const onClickLeft = () => {
  router.back()
}

const fetchDetail = async () => {
  loading.value = true
  try {
    const res = await request.get(`/announcements/${route.params.id}`)
    announcement.value = res.data
  } catch (error) {
    console.error('Failed to fetch detail:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const handleLike = async () => {
  try {
    await request.post(`/user/announcements/${route.params.id}/like`)
    isLiked.value = true
    isDisliked.value = false
    announcement.value.like_count = (announcement.value.like_count || 0) + 1
    showToast('点赞成功')
  } catch (error) {
    console.error('Failed to like:', error)
    showToast('操作失败')
  }
}

const handleDislike = async () => {
  try {
    await request.post(`/user/announcements/${route.params.id}/dislike`)
    isDisliked.value = true
    isLiked.value = false
    announcement.value.dislike_count = (announcement.value.dislike_count || 0) + 1
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
}

.page-content {
  padding: 16px;
}

.detail-content {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
}

.detail-title {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 12px;
}

.detail-meta {
  display: flex;
  align-items: center;
  gap: 16px;
  font-size: 12px;
  color: #999;
  margin-bottom: 16px;
}

.detail-body {
  font-size: 14px;
  color: #666;
  line-height: 1.8;
  margin-bottom: 20px;
}

.action-bar {
  display: flex;
  gap: 12px;
  margin-top: 20px;
}

.action-bar .van-button {
  flex: 1;
}
</style>
