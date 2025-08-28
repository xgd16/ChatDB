<template>
  <svg
    :width="size"
    :height="size"
    viewBox="0 0 40 40"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <!-- 背景圆形 -->
    <circle cx="20" cy="20" r="20" fill="#10B981" />
    
    <!-- 机器人头部 -->
    <rect x="12" y="12" width="16" height="16" rx="2" fill="white" />
    
    <!-- 眼睛 -->
    <circle cx="16" cy="18" r="1.5" fill="#10B981" />
    <circle cx="24" cy="18" r="1.5" fill="#10B981" />
    
    <!-- 嘴巴 -->
    <rect x="17" y="22" width="6" height="2" rx="1" fill="#10B981" />
    
    <!-- 天线 -->
    <line x1="18" y1="12" x2="18" y2="8" stroke="white" stroke-width="2" />
    <line x1="22" y1="12" x2="22" y2="8" stroke="white" stroke-width="2" />
    <circle cx="18" cy="8" r="1" fill="white" />
    <circle cx="22" cy="8" r="1" fill="white" />
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
