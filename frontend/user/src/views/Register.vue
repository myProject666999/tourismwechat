<template>
  <div class="register-page">
    <van-nav-bar
      title="注册"
      :left-arrow="true"
      @click-left="handleBack"
    />

    <div class="logo-container">
      <van-icon name="plus" size="60" color="#1989fa" />
      <h2>注册新账号</h2>
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
      <van-field
        v-model="form.confirmPassword"
        type="password"
        name="confirmPassword"
        label="确认密码"
        placeholder="请再次输入密码"
        :rules="[{ required: true, message: '请确认密码' }]"
      />
      <van-field
        v-model="form.nickname"
        name="nickname"
        label="昵称"
        placeholder="请输入昵称（可选）"
      />
      <van-field
        v-model="form.phone"
        name="phone"
        label="手机号"
        placeholder="请输入手机号（可选）"
      />
    </van-cell-group>

    <div class="btn-group">
      <van-button
        type="primary"
        block
        :loading="loading"
        @click="handleRegister"
      >
        注册
      </van-button>
      <div class="link-text">
        <span @click="goToLogin">已有账号？立即登录</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const form = ref({
  username: '',
  password: '',
  confirmPassword: '',
  nickname: '',
  phone: ''
})

const handleBack = () => {
  router.back()
}

const handleRegister = async () => {
  if (!form.value.username || !form.value.password) {
    showToast('请输入用户名和密码')
    return
  }

  if (form.value.password !== form.value.confirmPassword) {
    showToast('两次密码输入不一致')
    return
  }

  loading.value = true
  try {
    await userStore.register(
      form.value.username,
      form.value.password,
      form.value.nickname,
      form.value.phone
    )
    showToast('注册成功')
    router.push('/login')
  } catch (error) {
    console.error('注册失败:', error)
  } finally {
    loading.value = false
  }
}

const goToLogin = () => {
  router.push('/login')
}
</script>

<style scoped>
.register-page {
  background-color: #f7f8fa;
  min-height: 100vh;
}

.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 0 20px;
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
