import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/Home.vue'),
    meta: { title: '首页' }
  },
  {
    path: '/announcement',
    name: 'Announcement',
    component: () => import('@/views/Announcement.vue'),
    meta: { title: '公告' }
  },
  {
    path: '/announcement/:id',
    name: 'AnnouncementDetail',
    component: () => import('@/views/AnnouncementDetail.vue'),
    meta: { title: '公告详情' }
  },
  {
    path: '/scenic',
    name: 'Scenic',
    component: () => import('@/views/Scenic.vue'),
    meta: { title: '景点推荐' }
  },
  {
    path: '/scenic/:id',
    name: 'ScenicDetail',
    component: () => import('@/views/ScenicDetail.vue'),
    meta: { title: '景点详情' }
  },
  {
    path: '/hotel',
    name: 'Hotel',
    component: () => import('@/views/Hotel.vue'),
    meta: { title: '酒店预订' }
  },
  {
    path: '/hotel/:id',
    name: 'HotelDetail',
    component: () => import('@/views/HotelDetail.vue'),
    meta: { title: '酒店详情' }
  },
  {
    path: '/hotel/:id/booking',
    name: 'HotelBooking',
    component: () => import('@/views/HotelBooking.vue'),
    meta: { title: '酒店预订', requiresAuth: true }
  },
  {
    path: '/route',
    name: 'Route',
    component: () => import('@/views/Route.vue'),
    meta: { title: '自驾路线' }
  },
  {
    path: '/route/:id',
    name: 'RouteDetail',
    component: () => import('@/views/RouteDetail.vue'),
    meta: { title: '路线详情' }
  },
  {
    path: '/booking',
    name: 'BookingList',
    component: () => import('@/views/BookingList.vue'),
    meta: { title: '预订管理', requiresAuth: true }
  },
  {
    path: '/booking/:id',
    name: 'BookingDetail',
    component: () => import('@/views/BookingDetail.vue'),
    meta: { title: '预订详情', requiresAuth: true }
  },
  {
    path: '/itinerary',
    name: 'ItineraryList',
    component: () => import('@/views/ItineraryList.vue'),
    meta: { title: '行程管理', requiresAuth: true }
  },
  {
    path: '/itinerary/create',
    name: 'ItineraryCreate',
    component: () => import('@/views/ItineraryCreate.vue'),
    meta: { title: '创建行程', requiresAuth: true }
  },
  {
    path: '/itinerary/:id',
    name: 'ItineraryDetail',
    component: () => import('@/views/ItineraryDetail.vue'),
    meta: { title: '行程详情', requiresAuth: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/Profile.vue'),
    meta: { title: '个人中心', requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录' }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register.vue'),
    meta: { title: '注册' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  document.title = to.meta.title || '旅游小程序'

  if (to.meta.requiresAuth && !userStore.isLoggedIn) {
    next({
      path: '/login',
      query: { redirect: to.fullPath }
    })
  } else {
    next()
  }
})

export default router
