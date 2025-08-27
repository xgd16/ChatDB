import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { LoginRes } from '@/types/user'

export const useUserStore = defineStore("userInfo", {
  state: () => ({} as LoginRes),
  actions: {
    setUserInfo(data: LoginRes) {
      this.$state = data
    }
  },
  persist: true
})