<template>
  <div class="login-page">
    <van-nav-bar
      title="登录"
      :left-arrow="true"
      @click-left="handleBack"
    />

    <div class="logo-container">
      <van-icon name="location" size="60" color="#1989fa" />
      <h2>旅游小程序</h2>
    </div>

    <van-cell-group inset>
      <van-field
        v-model="form.username"
        name="username"
        label="用户名"
        placeholder="请输入用户名"
        :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <van-field
        v-model="form.password"
        type="password"
        name="password"
        label="密码"
        placeholder="请输入密码"
        :rules="[{ required: true, message: '请填写密码' }]"
      />
    </van-cell-group>

    <div class="btn-group">
      <van-button
        type="primary"
        block
        :loading="loading"
        @click="handleLogin"
      >
        登录
      </van-button>
      <div class="link-text">
        <span @click="goToRegister">还没有账号？立即注册</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast } from 'vant'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const loading = ref(false)
const form = ref({
  username: '',
  password: ''
})

const handleBack = () => {
  router.back()
}

const handleLogin = async () => {
  if (!form.value.username || !form.value.password) {
    showToast('请输入用户名和密码')
    return
  }

  loading.value = true
  try {
    await userStore.login(form.value.username, form.value.password)
    showToast('登录成功')
    
    const redirect = route.query.redirect || '/home'
    router.push(redirect)
  } catch (error) {
    console.error('登录失败:', error)
  } finally {
    loading.value = false
  }
}

const goToRegister = () => {
  router.push('/register')
}
</script>

<style scoped>
.login-page {
  background-color: #f7f8fa;
  min-height: 100vh;
}

.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 0;
}

.logo-container h2 {
  margin-top: 16px;
  font-size: 20px;
  color: #333;
}

.btn-group {
  padding: 24px 16px;
}

.link-text {
  text-align: center;
  margin-top: 16px;
  color: #1989fa;
  font-size: 14px;
}

.link-text span {
  cursor: pointer;
}
</style>
