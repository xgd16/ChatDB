<template>
  <router-view />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';

onMounted(() => {
  // 可以在这里添加全局初始化逻辑
  console.log('AI Chat SQL 应用已启动');
});
</script>

<style>
/* 全局样式 */
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
