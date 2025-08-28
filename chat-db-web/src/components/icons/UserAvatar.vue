<template>
  <svg
    :width="size"
    :height="size"
    viewBox="0 0 40 40"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <!-- 背景圆形 -->
    <circle cx="20" cy="20" r="20" fill="#3B82F6" />
    
    <!-- 头部 -->
    <circle cx="20" cy="16" r="6" fill="white" />
    
    <!-- 身体 -->
    <path
      d="M8 32C8 26.4772 12.4772 22 18 22H22C27.5228 22 32 26.4772 32 32V40H8V32Z"
      fill="white"
    />
  </svg>
</template>

<script setup lang="ts">
interface Props {
  size?: number;
}

withDefaults(defineProps<Props>(), {
  size: 40
});
</script>
