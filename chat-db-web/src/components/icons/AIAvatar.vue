<template>
  <div class="ai-avatar">
    <i class="ri-robot-2-fill"></i>
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
.ai-avatar {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #10B981, #059669);
  border-radius: 50%;
  color: white;
  font-size: 22px;
}
</style>
