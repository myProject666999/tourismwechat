import { createRouter, createWebHistory } from 'vue-router'
import { useAdminStore } from '@/stores/admin'
import Layout from '@/layout/index.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录' }
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: { title: '仪表盘', requiresAuth: true }
      },
      {
        path: 'user',
        name: 'UserManage',
        component: () => import('@/views/user/Index.vue'),
        meta: { title: '用户管理', requiresAuth: true }
      },
      {
        path: 'admin',
        name: 'AdminManage',
        component: () => import('@/views/admin/Index.vue'),
        meta: { title: '管理员管理', requiresAuth: true }
      },
      {
        path: 'announcement',
        name: 'AnnouncementManage',
        component: () => import('@/views/announcement/Index.vue'),
        meta: { title: '公告管理', requiresAuth: true }
      },
      {
        path: 'scenic',
        name: 'ScenicManage',
        component: () => import('@/views/scenic/Index.vue'),
        meta: { title: '景点管理', requiresAuth: true }
      },
      {
        path: 'hotel',
        name: 'HotelManage',
        component: () => import('@/views/hotel/Index.vue'),
        meta: { title: '酒店管理', requiresAuth: true }
      },
      {
        path: 'route',
        name: 'RouteManage',
        component: () => import('@/views/route/Index.vue'),
        meta: { title: '线路管理', requiresAuth: true }
      },
      {
        path: 'booking',
        name: 'BookingManage',
        component: () => import('@/views/booking/Index.vue'),
        meta: { title: '预订管理', requiresAuth: true }
      },
      {
        path: 'itinerary',
        name: 'ItineraryManage',
        component: () => import('@/views/itinerary/Index.vue'),
        meta: { title: '行程管理', requiresAuth: true }
      },
      {
        path: 'daily-recommend',
        name: 'DailyRecommendManage',
        component: () => import('@/views/daily-recommend/Index.vue'),
        meta: { title: '每日推荐管理', requiresAuth: true }
      },
      {
        path: 'hotel-stat',
        name: 'HotelStatManage',
        component: () => import('@/views/hotel-stat/Index.vue'),
        meta: { title: '酒店统计管理', requiresAuth: true }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const adminStore = useAdminStore()
  
  document.title = to.meta.title ? `${to.meta.title} - 旅游小程序管理后台` : '旅游小程序管理后台'

  if (to.meta.requiresAuth && !adminStore.isLoggedIn) {
    next({
      path: '/login',
      query: { redirect: to.fullPath }
    })
  } else {
    next()
  }
})

export default router
