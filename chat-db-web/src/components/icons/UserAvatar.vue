<template>
  <div class="user-avatar">
    <i class="ri-user-fill"></i>
  </div>
</template>

<script setup lang="ts">
interface Props {
  size?: number;
}

withDefaults(defineProps<Props>(), {
  size: 40
});
</script>

<style scoped>
.user-avatar {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #3B82F6, #1E40AF);
  border-radius: 50%;
  color: white;
  font-size: 22px;
}
</style>
