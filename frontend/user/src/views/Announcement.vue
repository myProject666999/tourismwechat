<template>
  <div class="announcement-page page-container">
    <van-nav-bar
      title="公告"
      :left-arrow="true"
      @click-left="$router.back()"
    />

    <van-search
      v-model="searchValue"
      placeholder="搜索公告"
      @search="handleSearch"
    />

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="loadAnnouncements"
      >
        <van-card
          v-for="item in announcements"
          :key="item.id"
          :thumb="item.cover || 'https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg'"
          :title="item.title"
          :desc="`浏览 ${item.view_count} | 点赞 ${item.like_count}`"
          @click="$router.push(`/announcement/${item.id}`)"
        >
          <template #tags>
            <van-tag type="danger" size="small">
              <van-icon name="good-job-o" /> {{ item.like_count }}
            </van-tag>
            <van-tag type="primary" size="small">
              <van-icon name="eye-o" /> {{ item.view_count }}
            </van-tag>
          </template>
          <template #footer>
            <span style="color: #999; font-size: 12px;">
              {{ formatDate(item.created_at) }}
            </span>
          </template>
        </van-card>
      </van-list>
    </van-pull-refresh>

    <van-empty v-if="!loading && !refreshing && announcements.length === 0" description="暂无公告" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { showToast } from 'vant'
import dayjs from 'dayjs'
import request from '@/utils/request'

const searchValue = ref('')
const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)
const announcements = ref([])
const page = ref(1)

const formatDate = (date) => {
  return dayjs(date).format('YYYY-MM-DD HH:mm')
}

const handleSearch = () => {
  page.value = 1
  announcements.value = []
  finished.value = false
  loadAnnouncements()
}

const onRefresh = async () => {
  page.value = 1
  announcements.value = []
  finished.value = false
  await loadAnnouncements()
  refreshing.value = false
}

const loadAnnouncements = async () => {
  try {
    const res = await request.get('/announcements', {
      params: {
        page: page.value,
        page_size: 10,
        keyword: searchValue.value
      }
    })
    const newItems = res.data?.list || []
    announcements.value = [...announcements.value, ...newItems]
    page.value++
    finished.value = newItems.length < 10
  } catch (error) {
    showToast('加载失败')
    console.error('加载公告失败:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.announcement-page {
  background-color: #f7f8fa;
}
</style>
