<template>
  <div class="tool-usage-panel">
    <div class="panel-header">
      <div class="left">
        <i class="ri-cpu-line"></i>
        <span class="title">工具流程</span>
        <n-tag v-if="events && events.length" size="small" type="primary">{{ events.length }} 步</n-tag>
      </div>
      <div class="right">
        <n-button quaternary size="tiny" @click="$emit('close')">
          <template #icon>
            <i class="ri-eye-off-line"></i>
          </template>
          隐藏
        </n-button>
      </div>
    </div>
    <div class="panel-body" :class="{ empty: !events || events.length === 0 }">
      <div v-if="!events || events.length === 0" class="empty">
        <i class="ri-information-line"></i>
        <span>暂无工具执行记录</span>
      </div>
      <div v-else class="timeline">
        <div v-for="(item, idx) in events" :key="idx" class="timeline-item">
          <div class="bullet" :class="{ active: loading && idx === events.length - 1 }"></div>
          <div class="content">
            <div class="row">
              <n-tag size="small" type="success">Step {{ idx + 1 }}</n-tag>
              <span class="name">{{ item.name || '未知工具' }}</span>
              <n-spin v-if="loading && idx === events.length - 1" size="small" style="margin-left: 6px" />
            </div>
            <div v-if="item.output" class="output markdown" v-html="renderMarkdown(normalizeOutput(item.output))"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { marked } from 'marked';
import { NButton, NTag, NSpin } from 'naive-ui';

interface ToolEventItem {
  name: string;
  output?: any;
  timestamp?: number;
}

const props = defineProps<{ events: ToolEventItem[]; loading: boolean }>();
defineEmits<{ (e: 'close'): void }>();

const normalizeOutput = (output: any) => {
  try {
    if (output && typeof output === 'object') {
      if ('text' in output && typeof output.text === 'string') return output.text as string;
      if ('content' in (output as any) && typeof (output as any).content === 'string') return (output as any).content as string;
      return JSON.stringify(output, null, 2);
    }
    return String(output ?? '');
  } catch {
    return String(output ?? '');
  }
};

const renderMarkdown = (content: string) => {
  try {
    return marked(content ?? '');
  } catch {
    return content ?? '';
  }
};
</script>

<style scoped lang="scss">
.tool-usage-panel {
  border: 1px solid var(--n-border-color);
  border-radius: 12px;
  background: var(--n-card-color);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
}
.panel-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 12px;
  border-bottom: 1px solid var(--n-border-color);
}
.panel-header .left {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}
.panel-body {
  max-height: 220px;
  overflow: auto;
  padding: 10px 12px;
}
.panel-body.empty {
  color: var(--n-text-color-3);
}
.empty {
  display: flex;
  align-items: center;
  gap: 8px;
}
.timeline {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.timeline-item {
  display: flex;
  gap: 10px;
}
.bullet {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: var(--n-border-color);
  margin-top: 6px;
}
.bullet.active {
  background: var(--n-primary-color);
}
.content .row {
  display: flex;
  align-items: center;
  gap: 8px;
}
.content .name {
  font-weight: 500;
}
.output {
  margin: 6px 0 0 0;
  border-left: 3px solid var(--n-border-color);
  padding-left: 10px;
  font-size: 13px;
  line-height: 1.55;
}
.output.markdown :deep(code) {
  background: var(--n-code-color);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 0.9em;
}
.output.markdown :deep(pre) {
  background: var(--n-code-color);
  padding: 10px;
  border-radius: 6px;
  overflow: auto;
}
.output.markdown :deep(ul),
.output.markdown :deep(ol) {
  padding-left: 18px;
}
</style>


