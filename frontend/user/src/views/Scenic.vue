<template>
  <div class="page-container">
    <div class="page-header">
      <van-nav-bar
        title="景点推荐"
        left-arrow
        @click-left="onClickLeft"
      />
    </div>
    <div class="page-content" v-loading="loading">
      <van-search
        v-model="searchForm.keyword"
        placeholder="搜索景点"
        @search="handleSearch"
      />
      
      <div class="scenic-list">
        <van-empty v-if="list.length === 0 && !loading" description="暂无数据" />
        
        <div
          v-for="item in list"
          :key="item.id"
          class="scenic-item"
          @click="goToDetail(item.id)"
        >
          <div class="scenic-cover">
            <van-image
              :src="item.cover || getPlaceholderImage()"
              width="120"
              height="80"
              fit="cover"
            />
            <van-tag v-if="item.level" type="primary" size="mini">
              {{ item.level }}
            </van-tag>
          </div>
          <div class="scenic-info">
            <h3 class="scenic-name">{{ item.name }}</h3>
            <div class="scenic-desc" v-if="item.description">
              {{ item.description.substring(0, 50) }}...
            </div>
            <div class="scenic-meta">
              <van-icon name="location-o" />
              <span class="address">{{ item.address || '暂无地址' }}</span>
            </div>
            <div class="scenic-footer">
              <div class="likes">
                <van-icon name="like-o" />
                <span>{{ item.like_count || 0 }}</span>
              </div>
              <div class="price" v-if="item.price">
                <span class="price-label">¥</span>
                <span class="price-value">{{ item.price }}</span>
                <span class="price-unit">起</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
        <van-list
          v-model:loading="moreLoading"
          :finished="finished"
          finished-text="没有更多了"
          @load="onLoad"
        >
        </van-list>
      </van-pull-refresh>
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

const activeTab = ref(1)
const loading = ref(false)
const refreshing = ref(false)
const moreLoading = ref(false)
const finished = ref(false)
const list = ref([])
const pagination = ref({
  page: 1,
  pageSize: 10
})

const searchForm = ref({
  keyword: ''
})

const onClickLeft = () => {
  router.back()
}

const goToDetail = (id) => {
  router.push(`/scenic/${id}`)
}

const getPlaceholderImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=beautiful%20tourist%20attraction%20landscape%20photo&image_size=square'
}

const fetchList = async (isRefresh = false) => {
  if (isRefresh) {
    pagination.value.page = 1
    finished.value = false
  }

  const params = {
    page: pagination.value.page,
    page_size: pagination.value.pageSize,
    keyword: searchForm.value.keyword
  }

  try {
    const res = await request.get('/scenics', { params })
    const items = res.data?.items || []
    
    if (isRefresh) {
      list.value = items
    } else {
      list.value = [...list.value, ...items]
    }

    pagination.value.page++
    finished.value = items.length < pagination.value.pageSize
  } catch (error) {
    console.error('Failed to fetch list:', error)
    showToast('加载失败')
  } finally {
    loading.value = false
    refreshing.value = false
    moreLoading.value = false
  }
}

const handleSearch = () => {
  fetchList(true)
}

const onRefresh = () => {
  pagination.value.page = 1
  fetchList(true)
}

const onLoad = () => {
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

.scenic-list {
  margin-top: 12px;
}

.scenic-item {
  display: flex;
  background: #fff;
  border-radius: 8px;
  padding: 12px;
  margin-bottom: 12px;
}

.scenic-cover {
  position: relative;
  flex-shrink: 0;
}

.scenic-cover .van-tag {
  position: absolute;
  top: 4px;
  left: 4px;
}

.scenic-info {
  flex: 1;
  margin-left: 12px;
  display: flex;
  flex-direction: column;
}

.scenic-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0 0 6px 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.scenic-desc {
  font-size: 12px;
  color: #999;
  margin-bottom: 6px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.scenic-meta {
  display: flex;
  align-items: center;
  font-size: 12px;
  color: #999;
  margin-bottom: 8px;
}

.scenic-meta .address {
  margin-left: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.scenic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: auto;
}

.likes {
  display: flex;
  align-items: center;
  font-size: 12px;
  color: #999;
}

.likes span {
  margin-left: 4px;
}

.price {
  display: flex;
  align-items: baseline;
}

.price-label {
  font-size: 12px;
  color: #ff6b6b;
}

.price-value {
  font-size: 18px;
  font-weight: bold;
  color: #ff6b6b;
}

.price-unit {
  font-size: 12px;
  color: #999;
}
</style>
