<template>
  <el-container class="layout-container">
    <el-aside width="220px" class="aside">
      <div class="logo">
        <el-icon :size="28"><Location /></el-icon>
        <span>旅游小程序</span>
      </div>
      <el-menu
        :default-active="activeMenu"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
        router
      >
        <el-menu-item index="/dashboard">
          <el-icon><Odometer /></el-icon>
          <span>仪表盘</span>
        </el-menu-item>
        
        <el-sub-menu index="1">
          <template #title>
            <el-icon><User /></el-icon>
            <span>用户管理</span>
          </template>
          <el-menu-item index="/user">用户列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="2">
          <template #title>
            <el-icon><Avatar /></el-icon>
            <span>管理员管理</span>
          </template>
          <el-menu-item index="/admin">管理员列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="3">
          <template #title>
            <el-icon><ChatLineSquare /></el-icon>
            <span>公告管理</span>
          </template>
          <el-menu-item index="/announcement">公告列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="4">
          <template #title>
            <el-icon><PictureFilled /></el-icon>
            <span>景点管理</span>
          </template>
          <el-menu-item index="/scenic">景点列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="5">
          <template #title>
            <el-icon><OfficeBuilding /></el-icon>
            <span>酒店管理</span>
          </template>
          <el-menu-item index="/hotel">酒店列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="6">
          <template #title>
            <el-icon><Guide /></el-icon>
            <span>线路管理</span>
          </template>
          <el-menu-item index="/route">线路列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="7">
          <template #title>
            <el-icon><Ticket /></el-icon>
            <span>预订管理</span>
          </template>
          <el-menu-item index="/booking">预订列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="8">
          <template #title>
            <el-icon><List /></el-icon>
            <span>行程管理</span>
          </template>
          <el-menu-item index="/itinerary">行程列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="9">
          <template #title>
            <el-icon><Star /></el-icon>
            <span>每日推荐</span>
          </template>
          <el-menu-item index="/daily-recommend">推荐列表</el-menu-item>
        </el-sub-menu>
        
        <el-sub-menu index="10">
          <template #title>
            <el-icon><DataAnalysis /></el-icon>
            <span>统计管理</span>
          </template>
          <el-menu-item index="/hotel-stat">酒店统计</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    
    <el-container>
      <el-header class="header">
        <div class="header-left">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item v-if="currentRoute.meta.title">{{ currentRoute.meta.title }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" icon="UserFilled" />
              <span class="username">{{ adminInfo?.nickname || adminInfo?.username }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item divided command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      
      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { useAdminStore } from '@/stores/admin'

const route = useRoute()
const router = useRouter()
const adminStore = useAdminStore()

const currentRoute = computed(() => route)
const activeMenu = computed(() => route.path)
const adminInfo = computed(() => adminStore.adminInfo)

const handleCommand = (command) => {
  if (command === 'logout') {
    ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      adminStore.logout()
      router.push('/login')
    }).catch(() => {})
  }
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.aside {
  background-color: #304156;
  transition: width 0.3s;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 60px;
  background-color: #263445;
  color: #fff;
  font-size: 18px;
  font-weight: bold;
}

.logo span {
  margin-left: 10px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
}

.header-right {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.username {
  margin: 0 8px;
  font-size: 14px;
}

.main {
  background-color: #f0f2f5;
  padding: 20px;
  overflow-y: auto;
}
</style>
