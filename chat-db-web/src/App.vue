<template>
  <n-config-provider :theme="currentTheme">
    <n-global-style />
    <n-message-provider>
      <router-view />
    </n-message-provider>
  </n-config-provider>
  
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue';
import { NConfigProvider, NGlobalStyle, NMessageProvider, darkTheme } from 'naive-ui';
import { useThemeStore } from '@/stores/theme';

const themeStore = useThemeStore();

onMounted(() => {
  themeStore.initTheme();
  // 应用启动日志
  console.log('AI Chat SQL 应用已启动');
});

const currentTheme = computed(() => (themeStore.isDark ? darkTheme : null));
</script>

<style>
/* 全局样式（不设定品牌色，使用 naive-ui 默认） */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}


</style>
