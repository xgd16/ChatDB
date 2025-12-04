<template>
  <router-view />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';

onMounted(() => {
  // 应用启动日志
  console.log('AI Chat SQL 应用已启动');
});
</script>

<style>
/* 全局基础样式，Element Plus 默认风格 */
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
  width: 100%;
  min-height: 100vh;
  overflow-x: hidden;
  overflow-y: auto;
}

#app {
  width: 100%;
  min-height: 100vh;
}


</style>
