import { createRouter, createWebHashHistory } from 'vue-router'
import LoginView from '../views/user/Login.vue'
import HomeView from '../views/Home.vue'
import { useUserStore } from '@/stores/counter'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {
        title: '登录',
        requiresAuth: false
      }
    },
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        title: '主页',
        requiresAuth: true
      }
    }
  ],
})

// 全局路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  // 设置页面标题
  if (to.meta.title) {
    document.title = '问库 - ' + (to.meta.title as string)
  }
  // 判断没有登录信息的话返回到登录页面
  if (!userStore.isLogin() && to.name !== 'login') {
    router.push('/login')
    return
  }
  next()
})

export default router
