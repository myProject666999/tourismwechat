<template>
  <div class="profile-page page-container">
    <van-nav-bar title="个人中心" />

    <div class="user-card">
      <van-image
        round
        :src="userInfo?.avatar || 'https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg'"
        class="avatar"
      />
      <div class="user-info">
        <div class="nickname">{{ userInfo?.nickname || userInfo?.username || '未设置昵称' }}</div>
        <div class="username">@{{ userInfo?.username }}</div>
      </div>
      <van-icon name="arrow" size="20" color="#999" @click="$router.push('/profile/edit')" />
    </div>

    <van-cell-group>
      <van-cell
        title="我的预订"
        is-link
        icon="orders-o"
        @click="$router.push('/booking')"
      />
      <van-cell
        title="我的行程"
        is-link
        icon="todo-list-o"
        @click="$router.push('/itinerary')"
      />
    </van-cell-group>

    <van-cell-group>
      <van-cell
        title="设置"
        is-link
        icon="setting-o"
      />
      <van-cell
        title="关于我们"
        is-link
        icon="info-o"
      />
    </van-cell-group>

    <div class="logout-btn" v-if="isLoggedIn">
      <van-button type="danger" block @click="handleLogout">退出登录</van-button>
    </div>

    <van-tabbar v-model="activeTab">
      <van-tabbar-item icon="home-o" @click="$router.push('/home')">首页</van-tabbar-item>
      <van-tabbar-item icon="orders-o" @click="$router.push('/booking')">预订</van-tabbar-item>
      <van-tabbar-item icon="todo-list-o" @click="$router.push('/itinerary')">行程</van-tabbar-item>
      <van-tabbar-item icon="user-o" @click="$router.push('/profile')">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const activeTab = ref(3)
const userInfo = computed(() => userStore.userInfo)
const isLoggedIn = computed(() => userStore.isLoggedIn)

const handleLogout = () => {
  showConfirmDialog({
    title: '提示',
    message: '确定要退出登录吗？',
  })
    .then(() => {
      userStore.logout()
      showToast('已退出登录')
      router.push('/login')
    })
    .catch(() => {})
}

onMounted(() => {
  if (userStore.isLoggedIn) {
    userStore.getProfile()
  }
})
</script>

<style scoped>
.profile-page {
  background-color: #f7f8fa;
}

.user-card {
  display: flex;
  align-items: center;
  padding: 20px 16px;
  background-color: #1989fa;
  color: #fff;
}

.avatar {
  width: 60px;
  height: 60px;
  border: 2px solid #fff;
}

.user-info {
  flex: 1;
  margin-left: 12px;
}

.nickname {
  font-size: 18px;
  font-weight: bold;
}

.username {
  font-size: 14px;
  opacity: 0.8;
  margin-top: 4px;
}

.logout-btn {
  padding: 24px 16px;
}
</style>
