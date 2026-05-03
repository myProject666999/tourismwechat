<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="酒店详情"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <div v-if="hotel" class="detail-content">
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
          <h2 class="title">{{ hotel.name }}</h2>
          <div class="stars" v-if="hotel.star">
            <van-icon v-for="i in hotel.star" :key="i" name="star" color="#ffc107" />
          </div>
          <div class="price-row">
            <span class="price-unit">¥</span>
            <span class="price">{{ hotel.price || 0 }}</span>
            <span class="price-label">起/晚</span>
          </div>
          <div class="address-row">
            <van-icon name="location-o" />
            <span>{{ hotel.address || '暂无地址' }}</span>
          </div>
        </div>

        <div class="section" v-if="hotel.description">
          <div class="section-title">酒店介绍</div>
          <div class="section-content">{{ hotel.description }}</div>
        </div>

        <div class="section" v-if="hotel.facilities">
          <div class="section-title">设施服务</div>
          <div class="section-content">{{ hotel.facilities }}</div>
        </div>
      </div>
    </div>

    <div class="bottom-bar">
      <van-button
        type="primary"
        size="large"
        block
        @click="goToBooking"
      >
        立即预订
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
const hotel = ref(null)

const images = computed(() => {
  if (!hotel.value) return []
  const result = []
  if (hotel.value.cover) {
    result.push(hotel.value.cover)
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
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=luxury%20hotel%20interior%20room%20photo&image_size=landscape_16_9'
}

const fetchDetail = async () => {
  loading.value = true
  try {
    const res = await request.get(`/hotels/${route.params.id}`)
    hotel.value = res.data
  } catch (error) {
    console.error('Failed to fetch detail:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const goToBooking = () => {
  router.push(`/hotel/${route.params.id}/booking`)
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

.stars {
  margin-bottom: 12px;
}

.price-row {
  display: flex;
  align-items: baseline;
  margin-bottom: 12px;
}

.price-unit {
  font-size: 14px;
  color: #ff6b6b;
}

.price {
  font-size: 24px;
  font-weight: bold;
  color: #ff6b6b;
}

.price-label {
  font-size: 12px;
  color: #999;
  margin-left: 4px;
}

.address-row {
  display: flex;
  align-items: center;
  font-size: 14px;
  color: #666;
}

.address-row span {
  margin-left: 6px;
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
  padding: 10px 16px;
  background: #fff;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.05);
}
</style>
