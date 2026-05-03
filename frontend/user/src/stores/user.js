import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import request from '@/utils/request'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('userInfo') || 'null'))

  const isLoggedIn = computed(() => !!token.value)

  async function login(username, password) {
    const res = await request.post('/login', {
      username,
      password
    })
    
    token.value = res.data.token
    userInfo.value = res.data.user
    
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('userInfo', JSON.stringify(res.data.user))
    
    return res
  }

  async function register(username, password, nickname, phone) {
    const res = await request.post('/register', {
      username,
      password,
      nickname,
      phone
    })
    return res
  }

  async function getProfile() {
    const res = await request.get('/user/profile')
    userInfo.value = res.data
    localStorage.setItem('userInfo', JSON.stringify(res.data))
    return res
  }

  async function updateProfile(data) {
    const res = await request.put('/user/profile', data)
    if (res.data) {
      userInfo.value = { ...userInfo.value, ...res.data }
      localStorage.setItem('userInfo', JSON.stringify(userInfo.value))
    }
    return res
  }

  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    login,
    register,
    getProfile,
    updateProfile,
    logout
  }
})
