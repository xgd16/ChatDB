import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { LoginRes } from '@/types/user'

export const useUserStore = defineStore("userInfo", {
  state: () => ({
    token: '',
    expire: 0,
    userId: 0,
    username: '',
    ruleLevel: 0,
    lastLoginTme: 0,
    createTime: 0
  } as LoginRes),
  actions: {
    setUserInfo(data: LoginRes) {
      this.token = data.token;
      this.expire = data.expire;
      this.userId = data.userId;
      this.username = data.username;
      this.ruleLevel = data.ruleLevel;
      this.lastLoginTme = data.lastLoginTme;
      this.createTime = data.createTime;
    },
    isLogin() {
      return !!this.token;
    },
    clearUserInfo() {
      this.token = '';
      this.expire = 0;
      this.userId = 0;
      this.username = '';
      this.ruleLevel = 0;
      this.lastLoginTme = 0;
      this.createTime = 0;
    }
  },
  persist: true
})