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
      <van-search
        v-model="searchForm.keyword"
        placeholder="搜索酒店"
        @search="handleSearch"
      />
      
      <div class="hotel-list">
        <van-empty v-if="list.length === 0 && !loading" description="暂无数据" />
        
        <div
          v-for="item in list"
          :key="item.id"
          class="hotel-item"
          @click="goToDetail(item.id)"
        >
          <div class="hotel-cover">
            <van-image
              :src="item.cover || getPlaceholderImage()"
              width="100"
              height="70"
              fit="cover"
            />
          </div>
          <div class="hotel-info">
            <h3 class="hotel-name">{{ item.name }}</h3>
            <div class="hotel-stars" v-if="item.star">
              <van-icon v-for="i in item.star" :key="i" name="star" color="#ffc107" />
            </div>
            <div class="hotel-address">
              <van-icon name="location-o" />
              <span>{{ item.address || '暂无地址' }}</span>
            </div>
            <div class="hotel-footer">
              <div class="price">
                <span class="price-unit">¥</span>
                <span class="price-value">{{ item.price || 0 }}</span>
                <span class="price-label">起/晚</span>
              </div>
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

const activeTab = ref(2)
const loading = ref(false)
const list = ref([])

const searchForm = ref({
  keyword: ''
})

const onClickLeft = () => {
  router.back()
}

const goToDetail = (id) => {
  router.push(`/hotel/${id}`)
}

const getPlaceholderImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=luxury%20hotel%20exterior%20architecture%20photo&image_size=square'
}

const fetchList = async () => {
  loading.value = true
  try {
    const res = await request.get('/hotels', {
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

.hotel-list {
  margin-top: 12px;
}

.hotel-item {
  display: flex;
  background: #fff;
  border-radius: 8px;
  padding: 12px;
  margin-bottom: 12px;
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
  margin: 0 0 6px 0;
}

.hotel-stars {
  margin-bottom: 6px;
}

.hotel-address {
  display: flex;
  align-items: center;
  font-size: 12px;
  color: #999;
  margin-bottom: 8px;
}

.hotel-address span {
  margin-left: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.hotel-footer {
  margin-top: auto;
}

.price {
  display: flex;
  align-items: baseline;
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
</style>
