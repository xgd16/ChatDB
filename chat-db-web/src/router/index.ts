import { createRouter, createWebHashHistory } from 'vue-router'
import LoginView from '../views/user/Login.vue'
import HomeView from '../views/Home.vue'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {
        title: '登录 - AI 智能 SQL 助手',
        requiresAuth: false
      }
    },
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        title: 'AI 智能 SQL 助手',
        requiresAuth: true
      }
    }
  ],
})

// 全局路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  if (to.meta.title) {
    document.title = to.meta.title as string
  }
  next()
})

export default router
