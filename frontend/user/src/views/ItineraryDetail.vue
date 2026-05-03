<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="行程详情"
        left-arrow
        @click-left="onClickLeft"
        :right-arrow="rightArrow"
        @click-right="showActionSheet = true"
      />
      <template #right>
        <van-icon name="more-o" size="20" />
      </template>
    </div>
    <div class="page-content" v-loading="loading">
      <div v-if="itinerary" class="detail-content">
        <div class="info-card">
          <div class="title-row">
            <h2 class="title">{{ itinerary.title }}</h2>
            <van-tag :type="getStatusType(itinerary.status)">
              {{ getStatusText(itinerary.status) }}
            </van-tag>
          </div>
          <div class="dates-row" v-if="itinerary.start_date || itinerary.end_date">
            <span class="date">{{ itinerary.start_date || '待定' }}</span>
            <van-icon name="arrow" />
            <span class="date">{{ itinerary.end_date || '待定' }}</span>
          </div>
          <div class="desc-row" v-if="itinerary.description">
            {{ itinerary.description }}
          </div>
        </div>

        <div class="section">
          <div class="section-header">
            <span class="section-title">行程安排</span>
            <van-button type="primary" size="small" @click="showAddItem = true">
              添加行程
            </van-button>
          </div>

          <div class="items-list">
            <van-empty v-if="items.length === 0" description="暂无行程安排" />
            
            <div
              v-for="(item, index) in items"
              :key="item.id"
              class="item-card"
            >
              <div class="item-header">
                <van-tag type="warning">第 {{ item.day_index || 1 }} 天</van-tag>
                <div class="item-actions">
                  <van-button type="primary" link size="small" @click="editItem(item)">
                    编辑
                  </van-button>
                  <van-button type="danger" link size="small" @click="deleteItem(item)">
                    删除
                  </van-button>
                </div>
              </div>
              <div class="item-content">
                <h3 class="item-title">{{ item.title }}</h3>
                <div class="item-time" v-if="item.start_time || item.end_time">
                  {{ item.start_time || '待定' }} - {{ item.end_time || '待定' }}
                </div>
                <div class="item-location" v-if="item.location">
                  <van-icon name="location-o" />
                  <span>{{ item.location }}</span>
                </div>
                <div class="item-desc" v-if="item.description">
                  {{ item.description }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <van-action-sheet
      v-model:show="showActionSheet"
      :actions="actions"
      cancel-text="取消"
      @select="onActionSelect"
    />

    <van-dialog
      v-model:show="showAddItem"
      :title="isEditItem ? '编辑行程' : '添加行程'"
      show-cancel-button
      @confirm="submitItem"
    >
      <van-form>
        <van-field
          v-model="itemForm.day_index"
          label="第几天"
          type="number"
          placeholder="请输入天数"
        />
        <van-field
          v-model="itemForm.title"
          label="行程标题"
          placeholder="请输入行程标题"
        />
        <van-field
          v-model="itemForm.start_time"
          label="开始时间"
          placeholder="如: 08:00"
        />
        <van-field
          v-model="itemForm.end_time"
          label="结束时间"
          placeholder="如: 12:00"
        />
        <van-field
          v-model="itemForm.location"
          label="地点"
          placeholder="请输入地点"
        />
        <van-field
          v-model="itemForm.description"
          label="描述"
          type="textarea"
          placeholder="请输入描述"
        />
      </van-form>
    </van-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast, showConfirmDialog } from 'vant'
import request from '@/utils/request'

const route = useRoute()
const router = useRouter()

const rightArrow = ref(true)
const loading = ref(false)
const itinerary = ref(null)
const items = ref([])
const showActionSheet = ref(false)
const showAddItem = ref(false)
const isEditItem = ref(false)
const editingItemId = ref(0)

const statusMap = {
  0: { text: '规划中', type: 'warning' },
  1: { text: '进行中', type: 'primary' },
  2: { text: '已完成', type: 'success' }
}

const actions = computed(() => [
  { name: '修改状态', value: 'update-status' },
  { name: '删除行程', value: 'delete', color: '#ee0a24' }
])

const itemForm = reactive({
  day_index: 1,
  title: '',
  start_time: '',
  end_time: '',
  location: '',
  description: '',
  item_type: 'activity'
})

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
    const res = await request.get(`/user/itineraries/${route.params.id}`)
    itinerary.value = res.data
  } catch (error) {
    console.error('Failed to fetch detail:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
  }
}

const onActionSelect = (action) => {
  if (action.value === 'delete') {
    deleteItinerary()
  } else if (action.value === 'update-status') {
    showToast('功能开发中')
  }
}

const deleteItinerary = async () => {
  try {
    await showConfirmDialog({
      title: '提示',
      message: '确定要删除该行程吗？'
    })
    showToast('删除成功')
    router.back()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('Failed to delete:', error)
      showToast('操作失败')
    }
  }
}

const editItem = (item) => {
  isEditItem.value = true
  editingItemId.value = item.id
  Object.assign(itemForm, {
    day_index: item.day_index || 1,
    title: item.title || '',
    start_time: item.start_time || '',
    end_time: item.end_time || '',
    location: item.location || '',
    description: item.description || '',
    item_type: item.item_type || 'activity'
  })
  showAddItem.value = true
}

const deleteItem = async (item) => {
  try {
    await showConfirmDialog({
      title: '提示',
      message: '确定要删除该行程安排吗？'
    })
    showToast('删除成功')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('Failed to delete:', error)
      showToast('操作失败')
    }
  }
}

const submitItem = async () => {
  if (!itemForm.title) {
    showToast('请输入行程标题')
    return
  }
  
  if (isEditItem.value) {
    showToast('修改成功')
  } else {
    showToast('添加成功')
    items.value.push({
      id: Date.now(),
      ...itemForm
    })
  }
  
  showAddItem.value = false
  resetItemForm()
}

const resetItemForm = () => {
  Object.assign(itemForm, {
    day_index: 1,
    title: '',
    start_time: '',
    end_time: '',
    location: '',
    description: '',
    item_type: 'activity'
  })
  isEditItem.value = false
  editingItemId.value = 0
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
  padding: 12px;
}

.info-card {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
}

.title-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.title {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin: 0;
  flex: 1;
  margin-right: 12px;
}

.dates-row {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px;
  background: #f7f8fa;
  border-radius: 8px;
  margin-bottom: 12px;
}

.dates-row .date {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.desc-row {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
}

.section {
  background: #fff;
  border-radius: 8px;
  margin-top: 12px;
  padding: 16px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.section-title {
  font-size: 15px;
  font-weight: bold;
  color: #333;
}

.items-list {
  margin-top: 12px;
}

.item-card {
  background: #f7f8fa;
  border-radius: 8px;
  padding: 12px;
  margin-bottom: 12px;
}

.item-card:last-child {
  margin-bottom: 0;
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.item-actions {
  display: flex;
  gap: 8px;
}

.item-title {
  font-size: 15px;
  font-weight: bold;
  color: #333;
  margin: 0 0 8px 0;
}

.item-time {
  font-size: 13px;
  color: #1989fa;
  margin-bottom: 8px;
}

.item-location {
  display: flex;
  align-items: center;
  font-size: 13px;
  color: #666;
  margin-bottom: 8px;
}

.item-location span {
  margin-left: 4px;
}

.item-desc {
  font-size: 13px;
  color: #999;
  line-height: 1.5;
}
</style>
