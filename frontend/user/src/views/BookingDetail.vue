<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="预订详情"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <div v-if="booking" class="detail-content">
        <div class="status-card">
          <van-tag :type="getStatusType(booking.status)" size="large">
            {{ getStatusText(booking.status) }}
          </van-tag>
        </div>

        <div class="info-card">
          <div class="info-title">预订信息</div>
          <div class="info-row">
            <span class="info-label">预订编号</span>
            <span class="info-value">{{ booking.id }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">酒店名称</span>
            <span class="info-value">{{ booking.hotel_name || '-' }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">房型</span>
            <span class="info-value">{{ booking.room_type || '-' }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">入住日期</span>
            <span class="info-value">{{ booking.check_in_date || '-' }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">退房日期</span>
            <span class="info-value">{{ booking.check_out_date || '-' }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">入住人数</span>
            <span class="info-value">{{ booking.guest_count || 1 }} 人</span>
          </div>
        </div>

        <div class="info-card">
          <div class="info-title">联系人信息</div>
          <div class="info-row">
            <span class="info-label">联系人</span>
            <span class="info-value">{{ booking.contact_name || '-' }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">联系电话</span>
            <span class="info-value">{{ booking.contact_phone || '-' }}</span>
          </div>
          <div class="info-row" v-if="booking.remark">
            <span class="info-label">备注</span>
            <span class="info-value">{{ booking.remark }}</span>
          </div>
        </div>

        <div class="price-card">
          <div class="price-row">
            <span class="price-label">总价</span>
            <span class="price-value">
              <span class="unit">¥</span>
              <span class="number">{{ booking.total_price || 0 }}</span>
            </span>
          </div>
        </div>
      </div>
    </div>

    <div class="bottom-bar" v-if="booking && booking.status === 0">
      <van-button
        type="warning"
        size="large"
        block
        @click="handleCancel"
      >
        取消预订
      </van-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast, showConfirmDialog } from 'vant'
import request from '@/utils/request'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const booking = ref(null)

const statusMap = {
  0: { text: '待确认', type: 'warning' },
  1: { text: '已确认', type: 'primary' },
  2: { text: '已完成', type: 'success' },
  3: { text: '已取消', type: 'default' }
}

const onClickLeft = () => {
  router.back()
}

const getStatusText = (status) => {
  return statusMap[status]?.text || '未知'
}

const getStatusType = (status) => {
  return statusMap[status]?.type || 'default'
}

const fetchDetail = async () => {
  loading.value = true
  try {
    const res = await request.get(`/user/bookings/${route.params.id}`)
    booking.value = res.data
  } catch (error) {
    console.error('Failed to fetch detail:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const handleCancel = async () => {
  try {
    await showConfirmDialog({
      title: '提示',
      message: '确定要取消预订吗？'
    })
    await request.post(`/user/bookings/${route.params.id}/cancel`)
    showToast('取消成功')
    fetchDetail()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('Failed to cancel:', error)
      showToast('操作失败')
    }
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
  padding: 12px;
}

.status-card {
  background: linear-gradient(135deg, #1989fa, #07c160);
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  margin-bottom: 12px;
}

.info-card {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 12px;
}

.info-title {
  font-size: 15px;
  font-weight: bold;
  color: #333;
  margin-bottom: 12px;
  padding-bottom: 12px;
  border-bottom: 1px solid #eee;
}

.info-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid #f5f5f5;
}

.info-row:last-child {
  border-bottom: none;
}

.info-label {
  font-size: 14px;
  color: #999;
}

.info-value {
  font-size: 14px;
  color: #333;
}

.price-card {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
}

.price-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price-label {
  font-size: 15px;
  color: #333;
}

.price-value .unit {
  font-size: 14px;
  color: #ff6b6b;
}

.price-value .number {
  font-size: 24px;
  font-weight: bold;
  color: #ff6b6b;
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
