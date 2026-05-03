import axios from 'axios'
import { showToast } from 'vant'

const request = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

request.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

request.interceptors.response.use(
  (response) => {
    const res = response.data
    
    if (res.code === 200) {
      return res
    }
    
    showToast(res.message || '请求失败')
    return Promise.reject(new Error(res.message || '请求失败'))
  },
  (error) => {
    const status = error.response?.status
    
    if (status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      showToast('登录已过期，请重新登录')
      window.location.href = '/login'
    } else if (status === 403) {
      showToast('无权限访问')
    } else if (status === 500) {
      showToast('服务器错误')
    } else {
      showToast(error.message || '网络错误')
    }
    
    return Promise.reject(error)
  }
)

export default request
