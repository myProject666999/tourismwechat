<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="酒店预订"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <div class="hotel-info-card" v-if="hotel">
        <div class="hotel-cover">
          <van-image
            :src="hotel.cover || getPlaceholderImage()"
            width="80"
            height="60"
            fit="cover"
          />
        </div>
        <div class="hotel-info">
          <h3 class="hotel-name">{{ hotel.name }}</h3>
          <div class="hotel-price">
            <span class="price-unit">¥</span>
            <span class="price-value">{{ hotel.price || 0 }}</span>
            <span class="price-label">起/晚</span>
          </div>
        </div>
      </div>

      <van-form @submit="onSubmit">
        <van-cell-group inset>
          <van-cell
            title="入住日期"
            is-link
            @click="showCheckIn = true"
          >
            <template #value>
              <span :class="{ 'placeholder': !formData.checkInDate }">
                {{ formData.checkInDate || '请选择' }}
              </span>
            </template>
          </van-cell>
          <van-cell
            title="退房日期"
            is-link
            @click="showCheckOut = true"
          >
            <template #value>
              <span :class="{ 'placeholder': !formData.checkOutDate }">
                {{ formData.checkOutDate || '请选择' }}
              </span>
            </template>
          </van-cell>
          <van-field
            v-model="formData.roomType"
            name="roomType"
            label="房型"
            placeholder="请输入房型"
            :rules="[{ required: true, message: '请输入房型' }]"
          />
          <van-field
            v-model="formData.guestCount"
            name="guestCount"
            label="入住人数"
            type="number"
            placeholder="请输入入住人数"
            :rules="[{ required: true, message: '请输入入住人数' }]"
          />
          <van-field
            v-model="formData.contactName"
            name="contactName"
            label="联系人"
            placeholder="请输入联系人姓名"
            :rules="[{ required: true, message: '请输入联系人姓名' }]"
          />
          <van-field
            v-model="formData.contactPhone"
            name="contactPhone"
            label="联系电话"
            placeholder="请输入联系电话"
            :rules="[{ required: true, message: '请输入联系电话' }]"
          />
          <van-field
            v-model="formData.remark"
            name="remark"
            label="备注"
            type="textarea"
            placeholder="请输入备注信息"
          />
        </van-cell-group>

        <div class="total-price">
          <span class="total-label">总价</span>
          <span class="total-value">
            <span class="unit">¥</span>
            <span class="value">{{ calculateTotalPrice() }}</span>
          </span>
        </div>

        <div style="margin: 16px">
          <van-button
            round
            block
            type="primary"
            native-type="submit"
            :loading="submitting"
          >
            提交预订
          </van-button>
        </div>
      </van-form>
    </div>

    <van-calendar
      v-model:show="showCheckIn"
      type="single"
      color="#1989fa"
      @confirm="onCheckInConfirm"
    />
    <van-calendar
      v-model:show="showCheckOut"
      type="single"
      color="#1989fa"
      @confirm="onCheckOutConfirm"
    />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast } from 'vant'
import dayjs from 'dayjs'
import request from '@/utils/request'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const submitting = ref(false)
const hotel = ref(null)
const showCheckIn = ref(false)
const showCheckOut = ref(false)

const formData = reactive({
  checkInDate: '',
  checkOutDate: '',
  roomType: '',
  guestCount: 1,
  contactName: '',
  contactPhone: '',
  remark: ''
})

const onClickLeft = () => {
  router.back()
}

const getPlaceholderImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=luxury%20hotel%20room%20interior%20photo&image_size=square'
}

const fetchDetail = async () => {
  loading.value = true
  try {
    const res = await request.get(`/hotels/${route.params.id}`)
    hotel.value = res.data
  } catch (error) {
    console.error('Failed to fetch hotel:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const onCheckInConfirm = (value) => {
  formData.checkInDate = dayjs(value).format('YYYY-MM-DD')
  showCheckIn.value = false
}

const onCheckOutConfirm = (value) => {
  formData.checkOutDate = dayjs(value).format('YYYY-MM-DD')
  showCheckOut.value = false
}

const calculateTotalPrice = () => {
  if (!hotel.value || !formData.checkInDate || !formData.checkOutDate) {
    return hotel.value?.price || 0
  }
  const days = dayjs(formData.checkOutDate).diff(dayjs(formData.checkInDate), 'day')
  if (days <= 0) {
    return hotel.value.price
  }
  return hotel.value.price * days
}

const onSubmit = async (values) => {
  submitting.value = true
  try {
    await request.post('/user/bookings', {
      hotel_id: parseInt(route.params.id),
      check_in_date: formData.checkInDate,
      check_out_date: formData.checkOutDate,
      room_type: formData.roomType,
      guest_count: formData.guestCount,
      contact_name: formData.contactName,
      contact_phone: formData.contactPhone,
      remark: formData.remark,
      total_price: calculateTotalPrice()
    })
    showToast('预订成功')
    router.push('/booking')
  } catch (error) {
    console.error('Failed to create booking:', error)
    showToast('预订失败')
  } finally {
    submitting.value = false
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

.hotel-info-card {
  display: flex;
  background: #fff;
  margin: 12px;
  padding: 12px;
  border-radius: 8px;
}

.hotel-info {
  flex: 1;
  margin-left: 12px;
  display: flex;
  flex-direction: column;
}

.hotel-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0 0 8px 0;
}

.hotel-price {
  display: flex;
  align-items: baseline;
  margin-top: auto;
}

.price-unit {
  font-size: 12px;
  color: #ff6b6b;
}

.price-value {
  font-size: 18px;
  font-weight: bold;
  color: #ff6b6b;
}

.price-label {
  font-size: 12px;
  color: #999;
  margin-left: 4px;
}

.placeholder {
  color: #999;
}

.total-price {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #fff;
  margin: 12px;
  border-radius: 8px;
}

.total-label {
  font-size: 14px;
  color: #333;
}

.total-value .unit {
  font-size: 14px;
  color: #ff6b6b;
}

.total-value .value {
  font-size: 24px;
  font-weight: bold;
  color: #ff6b6b;
}
</style>
