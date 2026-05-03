<template>
  <div class="home-page page-container">
    <van-nav-bar title="旅游小程序" />
    
    <van-search 
      v-model="searchValue" 
      placeholder="搜索景点、酒店..." 
      @search="handleSearch"
    />

    <div class="quick-nav">
      <div class="nav-item" @click="navTo('/announcement')">
        <van-icon name="bulletin" size="28" color="#1989fa" />
        <span>公告</span>
      </div>
      <div class="nav-item" @click="navTo('/scenic')">
        <van-icon name="location" size="28" color="#ff976a" />
        <span>景点</span>
      </div>
      <div class="nav-item" @click="navTo('/hotel')">
        <van-icon name="home-o" size="28" color="#39b54a" />
        <span>酒店</span>
      </div>
      <div class="nav-item" @click="navTo('/route')">
        <van-icon name="direction" size="28" color="#ee0a24" />
        <span>路线</span>
      </div>
    </div>

    <div class="section">
      <div class="section-title">每日推荐</div>
      <div v-if="recommendationsLoading" class="loading-container">
        <van-loading type="spinner" size="24" />
        <span class="loading-text">加载中...</span>
      </div>
      <div v-else-if="dailyRecommendations.length > 0">
        <van-card
          v-for="item in dailyRecommendations"
          :key="item.id"
          :thumb="item.cover || 'https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg'"
          :title="item.title"
          :desc="item.description"
          @click="handleRecommendClick(item)"
        >
          <template #tags>
            <van-tag plain type="primary">{{ getRecommendTypeText(item.recommend_type) }}</van-tag>
          </template>
        </van-card>
      </div>
      <div v-else class="empty-container">
        <van-empty description="暂无推荐" />
      </div>
    </div>

    <div class="section">
      <div class="section-title">热门景点</div>
      <div v-if="scenicsLoading" class="loading-container">
        <van-loading type="spinner" size="24" />
        <span class="loading-text">加载中...</span>
      </div>
      <div v-else-if="scenicList.length > 0">
        <van-card
          v-for="item in scenicList"
          :key="item.id"
          :thumb="item.cover || 'https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg'"
          :title="item.name"
          :desc="item.address"
          :price="item.price"
          price-text="¥"
          @click="navTo(`/scenic/${item.id}`)"
        >
          <template #tags>
            <van-tag v-if="item.level" plain type="warning">{{ item.level }}</van-tag>
            <van-tag type="danger" size="small">
              <van-icon name="good-job-o" /> {{ item.like_count || 0 }}
            </van-tag>
          </template>
        </van-card>
      </div>
      <div v-else class="empty-container">
        <van-empty description="暂无景点" />
      </div>
    </div>

    <van-tabbar v-model="activeTab">
      <van-tabbar-item icon="home-o" @click="navTo('/home')">首页</van-tabbar-item>
      <van-tabbar-item icon="orders-o" @click="navTo('/booking')">预订</van-tabbar-item>
      <van-tabbar-item icon="todo-list-o" @click="navTo('/itinerary')">行程</van-tabbar-item>
      <van-tabbar-item icon="user-o" @click="navTo('/profile')">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import request from '@/utils/request'

const router = useRouter()
const searchValue = ref('')
const activeTab = ref(0)

const recommendationsLoading = ref(false)
const dailyRecommendations = ref([])

const scenicsLoading = ref(false)
const scenicList = ref([])

const navTo = (path) => {
  router.push(path)
}

const handleSearch = () => {
  if (searchValue.value) {
    showToast(`搜索: ${searchValue.value}`)
  }
}

const getRecommendTypeText = (type) => {
  const map = {
    'scenic': '景点',
    'hotel': '酒店',
    'route': '路线'
  }
  return map[type] || type
}

const loadRecommendations = async () => {
  recommendationsLoading.value = true
  try {
    const res = await request.get('/daily-recommendations')
    dailyRecommendations.value = res.data || []
  } catch (error) {
    console.error('加载推荐失败:', error)
  } finally {
    recommendationsLoading.value = false
  }
}

const loadScenics = async () => {
  scenicsLoading.value = true
  try {
    const res = await request.get('/scenics', {
      params: {
        page: 1,
        page_size: 10
      }
    })
    scenicList.value = res.data?.list || []
  } catch (error) {
    console.error('加载景点失败:', error)
  } finally {
    scenicsLoading.value = false
  }
}

const handleRecommendClick = (item) => {
  if (item.recommend_type === 'scenic' && item.target_id) {
    router.push(`/scenic/${item.target_id}`)
  } else if (item.recommend_type === 'hotel' && item.target_id) {
    router.push(`/hotel/${item.target_id}`)
  } else if (item.recommend_type === 'route' && item.target_id) {
    router.push(`/route/${item.target_id}`)
  }
}

onMounted(() => {
  loadRecommendations()
  loadScenics()
})
</script>

<style scoped>
.home-page {
  background-color: #f7f8fa;
}

.quick-nav {
  display: flex;
  justify-content: space-around;
  padding: 16px;
  background-color: #fff;
  margin-bottom: 12px;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #333;
}

.section {
  background-color: #fff;
  margin-bottom: 12px;
  padding: 0 12px 12px;
}

.section-title {
  font-size: 16px;
  font-weight: bold;
  padding: 12px 0;
  border-bottom: 1px solid #eee;
  margin-bottom: 12px;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 0;
  gap: 12px;
}

.loading-text {
  font-size: 14px;
  color: #999;
}

.empty-container {
  padding: 20px 0;
}
</style>
