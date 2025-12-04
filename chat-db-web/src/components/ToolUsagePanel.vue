<template>
  <div class="tool-usage-panel">
    <div class="panel-header">
      <div class="header-title">
        <i class="ri-cpu-line"></i>
        <span>工具执行流程</span>
      </div>
      <div class="header-actions">
        <el-tag v-if="events && events.length" size="small" round effect="plain">{{ events.length }} 步</el-tag>
        <el-button link size="small" @click="$emit('close')">
          <i class="ri-close-line"></i>
        </el-button>
      </div>
    </div>

    <div class="panel-content" ref="contentRef">
      <div v-if="!events || events.length === 0" class="empty-state">
        <div class="empty-icon">
          <i class="ri-flow-chart"></i>
        </div>
        <p>暂无工具调用</p>
        <span class="sub-text">AI 调用工具查询数据时会显示在这里</span>
      </div>

      <div v-else class="timeline-container">
        <div v-for="(item, idx) in events" :key="idx" class="timeline-item">
          <div class="timeline-left">
            <div class="line-top" v-if="idx > 0"></div>
            <div class="dot" :class="{ active: loading && idx === events.length - 1, finished: idx < events.length - 1 }">
              <i v-if="idx < events.length - 1" class="ri-check-fill"></i>
              <span v-else>{{ idx + 1 }}</span>
            </div>
            <div class="line-bottom" v-if="idx < events.length - 1"></div>
          </div>
          
          <div class="timeline-content">
            <div class="step-header">
              <span class="tool-name">{{ item.name || '未知工具' }}</span>
              <span class="time" v-if="item.timestamp">{{ formatTime(item.timestamp) }}</span>
              <el-icon v-if="loading && idx === events.length - 1" class="is-loading">
                <Loading />
              </el-icon>
            </div>
            
            <div v-if="item.output" class="step-output">
              <div class="output-label">输出结果:</div>
              <div class="code-block markdown-body" v-html="renderMarkdown(normalizeOutput(item.output))"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, nextTick } from 'vue';
import { marked } from 'marked';
import { ElButton, ElTag, ElIcon } from 'element-plus';
import { Loading } from '@element-plus/icons-vue';

interface ToolEventItem {
  name: string;
  output?: any;
  timestamp?: number;
}

const props = defineProps<{ events: ToolEventItem[]; loading: boolean }>();
defineEmits<{ (e: 'close'): void }>();

const contentRef = ref<HTMLElement>();

watch(() => props.events.length, () => {
  nextTick(() => {
    if (contentRef.value) {
      contentRef.value.scrollTop = contentRef.value.scrollHeight;
    }
  });
});

const formatTime = (timestamp: number) => {
  return new Date(timestamp).toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit', second: '2-digit' });
};

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
  height: 100%;
  display: flex;
  flex-direction: column;
  background: var(--el-bg-color);
}

.panel-header {
  height: 56px;
  padding: 0 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid var(--el-border-color-lighter);
  flex-shrink: 0;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  font-size: 15px;
  color: var(--el-text-color-primary);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.panel-content {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  position: relative;
}

.empty-state {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: var(--el-text-color-secondary);
  opacity: 0.7;
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 12px;
  color: var(--el-border-color-darker);
}

.sub-text {
  font-size: 12px;
  margin-top: 4px;
}

/* Timeline Styles */
.timeline-container {
  display: flex;
  flex-direction: column;
}

.timeline-item {
  display: flex;
  gap: 12px;
  min-height: 60px;
}

.timeline-left {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 24px;
  flex-shrink: 0;
}

.dot {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: var(--el-fill-color);
  border: 2px solid var(--el-border-color);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
  color: var(--el-text-color-secondary);
  z-index: 1;
  transition: all 0.3s;
}

.dot.active {
  border-color: var(--el-color-primary);
  color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
}

.dot.finished {
  background: var(--el-color-success);
  border-color: var(--el-color-success);
  color: white;
}

.line-top, .line-bottom {
  width: 2px;
  background: var(--el-border-color-lighter);
  flex: 1;
}

.timeline-content {
  flex: 1;
  padding-bottom: 24px;
  min-width: 0; /* prevent overflow */
}

.step-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
  height: 24px; /* Align with dot */
}

.tool-name {
  font-weight: 600;
  font-size: 14px;
}

.time {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin-left: auto;
}

.step-output {
  background: var(--el-fill-color-light);
  border-radius: 8px;
  padding: 10px;
  font-size: 13px;
}

.output-label {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin-bottom: 4px;
}

.code-block {
  overflow-x: auto;
  line-height: 1.5;
}

.code-block :deep(pre) {
  margin: 0;
  background: transparent;
  padding: 0;
}

.code-block :deep(code) {
  font-family: 'Monaco', 'Menlo', monospace;
  font-size: 12px;
}
</style>