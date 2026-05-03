<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="创建行程"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content">
      <van-form @submit="onSubmit">
        <van-cell-group inset>
          <van-field
            v-model="formData.title"
            name="title"
            label="行程标题"
            placeholder="请输入行程标题"
            :rules="[{ required: true, message: '请输入行程标题' }]"
          />
          <van-field
            v-model="formData.description"
            name="description"
            label="行程描述"
            type="textarea"
            placeholder="请输入行程描述"
          />
          <van-cell
            title="开始日期"
            is-link
            @click="showStartDate = true"
          >
            <template #value>
              <span :class="{ 'placeholder': !formData.start_date }">
                {{ formData.start_date || '请选择' }}
              </span>
            </template>
          </van-cell>
          <van-cell
            title="结束日期"
            is-link
            @click="showEndDate = true"
          >
            <template #value>
              <span :class="{ 'placeholder': !formData.end_date }">
                {{ formData.end_date || '请选择' }}
              </span>
            </template>
          </van-cell>
        </van-cell-group>

        <div style="margin: 16px">
          <van-button
            round
            block
            type="primary"
            native-type="submit"
            :loading="submitting"
          >
            创建行程
          </van-button>
        </div>
      </van-form>
    </div>

    <van-calendar
      v-model:show="showStartDate"
      type="single"
      color="#1989fa"
      @confirm="onStartDateConfirm"
    />
    <van-calendar
      v-model:show="showEndDate"
      type="single"
      color="#1989fa"
      @confirm="onEndDateConfirm"
    />
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import dayjs from 'dayjs'
import request from '@/utils/request'

const router = useRouter()

const submitting = ref(false)
const showStartDate = ref(false)
const showEndDate = ref(false)

const formData = reactive({
  title: '',
  description: '',
  start_date: '',
  end_date: '',
  status: 0
})

const onClickLeft = () => {
  router.back()
}

const onStartDateConfirm = (value) => {
  formData.start_date = dayjs(value).format('YYYY-MM-DD')
  showStartDate.value = false
}

const onEndDateConfirm = (value) => {
  formData.end_date = dayjs(value).format('YYYY-MM-DD')
  showEndDate.value = false
}

const onSubmit = async (values) => {
  submitting.value = true
  try {
    await request.post('/user/itineraries', {
      title: formData.title,
      description: formData.description,
      start_date: formData.start_date,
      end_date: formData.end_date,
      status: 0
    })
    showToast('创建成功')
    router.push('/itinerary')
  } catch (error) {
    console.error('Failed to create:', error)
    showToast('创建失败')
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.page-container {
  min-height: 100vh;
  background-color: #f7f8fa;
}

.placeholder {
  color: #999;
}
</style>
