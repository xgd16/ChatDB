<template>
  <div class="chat-container">
    <BackgroundOrbs :orb-count="10" />
    <!-- 顶部工具栏 -->
    <div class="chat-header">
      <div class="header-left">
        <h2>问库</h2>
        <span class="status" :class="{ online: isConnected }">
          {{ isConnected ? '已连接' : '未连接' }}
        </span>
      </div>
      <div class="header-right">
        <n-button
          quaternary
          circle
          @click="showSettings = true"
          class="settings-btn"
        >
          <template #icon>
            <i class="ri-settings-3-line"></i>
          </template>
        </n-button>
      </div>
    </div>

    <!-- 聊天消息区域 -->
    <div class="chat-messages" ref="messagesContainer">
      <div
        v-for="(message, index) in messages"
        :key="index"
        class="message"
        :class="message.role"
      >
        <div class="message-avatar">
          <UserAvatar v-if="message.role === 'user'" :size="40" />
          <AIAvatar v-else :size="40" />
        </div>
        <div class="message-content">
          <div class="message-header">
            <span class="message-author">
              {{ message.role === 'user' ? '我' : 'AI 助手' }}
            </span>
            <span class="message-time">{{ formatTime(message.timestamp) }}</span>
            <n-button
              v-if="message.role === 'assistant' && (message as any).toolEvents && (message as any).toolEvents.length"
              size="tiny"
              quaternary
              style="margin-left: 8px"
              @click="(message as any).showToolPanel = !(message as any).showToolPanel"
            >
              <template #icon>
                <i :class="(message as any).showToolPanel === false ? 'ri-eye-line' : 'ri-eye-off-line'"></i>
              </template>
              {{ (message as any).showToolPanel === false ? '显示流程' : '隐藏流程' }}
            </n-button>
          </div>
          <ToolUsagePanel
            v-if="message.role === 'assistant' && (message as any).toolEvents && (message as any).toolEvents.length && (message as any).showToolPanel !== false"
            :events="(message as any).toolEvents"
            :loading="isLoading && index === messages.length - 1"
            @close="(message as any).showToolPanel = false"
            style="margin: 6px 0 10px 0;"
          />
          <div class="message-text">
            <template v-if="message.role === 'assistant'">
              <!-- 思考中：假进度条（内嵌在消息气泡内） -->
              <div v-if="index === messages.length - 1 && isLoading && (!message.content || message.content.length === 0) && (!(message as any).toolEvents || (message as any).toolEvents.length === 0)" class="thinking-inline">
                <div class="thinking-bar">
                  <div class="thinking-bar-inner" :style="{ width: Math.round(thinkingProgress) + '%' }"></div>
                  <div class="thinking-bar-percent">{{ Math.round(thinkingProgress) }}%</div>
                </div>
                <span class="thinking-text">正在思考...</span>
              </div>
              <div v-else class="markdown-content" v-html="renderMarkdown(message.content)"></div>
            </template>
            <div v-else class="user-message">{{ message.content }}</div>
          </div>
        </div>
      </div>
      
      
    </div>

    <!-- 悬浮输入区域 -->
    <div class="floating-input-container">
      <div class="glassmorphism-input">
        <div class="input-wrapper" id="message-enter">
          <n-input
            v-model:value="inputMessage"
            type="textarea"
            :autosize="{ minRows: 1, maxRows: 8 }"
            placeholder="输入你的问题... (Shift+Enter 换行，Enter 发送)"
            @keydown="handleKeyDown"
            :disabled="isLoading"
            class="glass-input"
            :show-count="false"
            clearable
          />
          <n-button
            type="primary"
            :disabled="!inputMessage.trim() || isLoading"
            @click="handleSendMessage"
            class="glass-send-button"
            size="large"
          >
            <template #icon>
              <i class="ri-send-plane-fill"></i>
            </template>
            发送
          </n-button>
        </div>
      </div>
    </div>
  </div>

  <!-- 设置抽屉 -->
  <ChatSettings v-model:show="showSettings" />
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick, watch } from 'vue';
import { 
  NAvatar, 
  NInput, 
  NButton, 
  NSpin,
  useMessage
} from 'naive-ui';
import UserAvatar from '@/components/icons/UserAvatar.vue';
import AIAvatar from '@/components/icons/AIAvatar.vue';
import ChatSettings from '@/components/ChatSettings.vue';
import BackgroundOrbs from '@/components/BackgroundOrbs.vue';
import ToolUsagePanel from '@/components/ToolUsagePanel.vue';
import { marked } from 'marked';
import { chatApi, getDatabaseConfigList } from '@/api/ai';
import type { ChatMessage } from '@/types/ai';
import { useSettingsStore } from '@/stores/settings';

// 消息提示
const message = useMessage();

// 设置相关
const settingsStore = useSettingsStore();

// 响应式数据
type ToolEvent = { name: string; output?: any; timestamp: number };
type LocalChatMessage = ChatMessage & { toolEvents?: ToolEvent[]; showToolPanel?: boolean };
const messages = ref<LocalChatMessage[]>([]);

const inputMessage = ref('');
const isLoading = ref(false);
const isConnected = ref(true);
const messagesContainer = ref<HTMLElement>();

// 思考中进度条（假进度）
const thinkingProgress = ref(0);
let thinkingTimer: number | undefined;
let hasReceivedFirstChunk = false;

const startThinkingProgress = () => {
  hasReceivedFirstChunk = false;
  thinkingProgress.value = 5;
  if (thinkingTimer) window.clearInterval(thinkingTimer);
  thinkingTimer = window.setInterval(() => {
    // 递增到 92% 封顶，避免太快结束
    if (thinkingProgress.value < 92) {
      thinkingProgress.value = Math.min(92, thinkingProgress.value + Math.random() * 4 + 2);
    }
  }, 120);
};

const finishThinkingProgress = () => {
  if (thinkingTimer) {
    window.clearInterval(thinkingTimer);
    thinkingTimer = undefined;
  }
  thinkingProgress.value = 100;
};

// 设置抽屉
const showSettings = ref(false);

// 工具面板相关（改为每条消息内部控制 showToolPanel 与 toolEvents）

// 配置 marked 选项
marked.setOptions({
  breaks: true,
  gfm: true
});

// 渲染 Markdown
const renderMarkdown = (content: string) => {
  try {
    return marked(content);
  } catch (error) {
    console.error('Markdown 渲染错误:', error);
    return content;
  }
};

// 格式化时间
const formatTime = (timestamp: number) => {
  const date = new Date(timestamp);
  return date.toLocaleTimeString('zh-CN', {
    hour: '2-digit',
    minute: '2-digit'
  });
};

// 滚动到底部
const scrollToBottom = async () => {
  await nextTick();
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
  }
};

// 键盘事件处理
const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    if (event.shiftKey) {
      // Shift+Enter: 换行
      return;
    } else {
      // Enter: 发送消息
      event.preventDefault();
      handleSendMessage();
    }
  }
};

// 发送消息
const handleSendMessage = async () => {
  if (!inputMessage.value.trim() || isLoading.value) return;

  const userMessage = inputMessage.value.trim();
  inputMessage.value = '';
  
  // 添加用户消息
  messages.value.push({
    role: 'user',
    content: userMessage,
    timestamp: Date.now()
  });

  // 预先添加占位的助手消息，用于显示“思考中”进度条
  messages.value.push({
    role: 'assistant',
    content: '',
    timestamp: Date.now(),
    toolEvents: [],
    showToolPanel: false
  } as LocalChatMessage);

  await scrollToBottom();
  isLoading.value = true;
  startThinkingProgress();

  try {
    // 调用 AI 聊天 API
    const response = await chatApi({
      ai: settingsStore.settings.ai,
      model: settingsStore.settings.model,
      message: userMessage,
      databaseId: settingsStore.settings.databaseId,
      prompt: ''
    });

    if (!response.ok) {
      throw new Error('网络请求失败');
    }

    const reader = response.body?.getReader();
    if (!reader) {
      throw new Error('无法读取响应流');
    }

    let aiResponse = '';
    const decoder = new TextDecoder();

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;

      const chunk = decoder.decode(value);
      const lines = chunk.split('\n');

      for (const line of lines) {
        if (line.startsWith('data: ')) {
          try {
            const data = JSON.parse(line.slice(6));
            if (data.event === 'tool_call') {
              // 确保存在当前的助手消息
              let lastMessage = messages.value[messages.value.length - 1];
              if (!lastMessage || lastMessage.role !== 'assistant') {
                lastMessage = {
                  role: 'assistant',
                  content: '',
                  timestamp: Date.now(),
                  toolEvents: [],
                  showToolPanel: true
                } as LocalChatMessage;
                messages.value.push(lastMessage);
              }
              if (!lastMessage.toolEvents) lastMessage.toolEvents = [];
              lastMessage.showToolPanel = true;
              lastMessage.toolEvents.push({
                name: data?.data?.name ?? '未知工具',
                output: data?.data?.output,
                timestamp: Date.now()
              });
              await scrollToBottom();
              continue;
            }
            if (data.event === 'message' && data.content) {
              if (!hasReceivedFirstChunk) {
                finishThinkingProgress();
                hasReceivedFirstChunk = true;
              }
              aiResponse += data.content;
              // 更新最后一条消息或创建新消息
              const lastMessage = messages.value[messages.value.length - 1];
              if (lastMessage && lastMessage.role === 'assistant') {
                lastMessage.content = aiResponse;
              } else {
                messages.value.push({
                  role: 'assistant',
                  content: aiResponse,
                  timestamp: Date.now(),
                  toolEvents: [],
                  showToolPanel: true
                });
              }
              await scrollToBottom();
            } else if (data.event === 'end') {
              break;
            }
          } catch (e) {
            // 忽略解析错误
          }
        }
      }
    }

  } catch (error) {
    console.error('发送消息失败:', error);
    message.error('发送消息失败，请重试');
    
    // 添加错误消息
    messages.value.push({
      role: 'assistant',
      content: '抱歉，我遇到了一些问题。请稍后重试。',
      timestamp: Date.now()
    });
  } finally {
    isLoading.value = false;
    if (thinkingTimer) window.clearInterval(thinkingTimer);
    await scrollToBottom();
  }
};

// 组件挂载时初始化设置和加载数据库配置
onMounted(async () => {
  // 初始化设置
  settingsStore.initSettings();
  
  try {
    const response = await getDatabaseConfigList();
    if (response.code === 0 && response.data.list && response.data.list.length > 0) {
      const dbOptions = response.data.list.map((db: any, index: number) => ({
        label: `${db.dbName} (${db.host}:${db.port})`,
        value: index + 1 // 使用索引+1作为databaseId
      }));
      settingsStore.updateDatabaseOptions(dbOptions);
    }
  } catch (error) {
    console.error('加载数据库配置失败:', error);
  }
  
  // 确保输入框正确初始化
  await nextTick();
});
</script>

<style lang="scss" scoped>
.chat-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--n-color);
  position: relative; // 让绝对定位的光球容器定位于此
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  border-bottom: 1px solid var(--n-border-color);
  background: var(--n-card-color);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;

  h2 {
    margin: 0;
    font-size: 18px;
    font-weight: 600;
  }
}

.status {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  background: var(--n-error-color);

  &.online {
    background: var(--n-success-color);
  }
}

.header-right {
  display: flex;
  align-items: center;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  padding-bottom: 120px; /* 为悬浮输入框留出空间 */
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.message {
  display: flex;
  gap: 12px;
  max-width: 80%;

  &.user {
    align-self: flex-end;
    flex-direction: row-reverse;

    .message-text {
      background: var(--n-primary-color);
      border-color: var(--n-primary-color);
    }
  }

  &.assistant {
    align-self: flex-start;
  }
}

.message-content {
  flex: 1;
  min-width: 0;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 12px;
  color: var(--n-text-color-3);
}

.message-author {
  font-weight: 500;
}

.message-time {
  opacity: 0.7;
}

.message-text {
  padding: 12px 16px;
  border-radius: 12px;
  background: var(--n-color-modal);
  backdrop-filter: blur(10px);
  border: 1px solid var(--n-border-color);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  opacity: 0.96;
  word-wrap: break-word;
  transition: box-shadow 0.3s ease, transform 0.2s ease, background 0.3s ease;
}

/* 思考中假进度条（内联） */
.thinking-inline {
  display: inline-flex;
  align-items: center;
  gap: 8px;
}
.thinking-bar {
  position: relative;
  width: 160px;
  height: 10px;
  background: var(--n-border-color);
  border-radius: 6px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06) inset, 0 2px 6px rgba(0, 0, 0, 0.06);
}
.thinking-bar-inner {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  background: linear-gradient(90deg, var(--n-primary-color), var(--n-primary-color-hover));
  transition: width 0.2s ease;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.08), 0 0 12px rgba(0, 0, 0, 0.06);
}
.thinking-bar-percent {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  color: #fff;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.35);
  pointer-events: none;
}
.thinking-text {
  display: inline-block;
  font-size: 12px;
  color: var(--n-text-color-3);
}
/* 不再使用滑动动画，使用基于宽度的百分比 */

.markdown-content {
  line-height: 1.6;

  :deep(p) {
    margin: 0 0 12px 0;

    &:last-child {
      margin-bottom: 0;
    }
  }

  :deep(code) {
    background: var(--n-code-color);
    padding: 2px 6px;
    border-radius: 4px;
    font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
    font-size: 0.9em;
  }

  :deep(pre) {
    background: var(--n-code-color);
    padding: 12px;
    border-radius: 8px;
    overflow-x: auto;
    margin: 12px 0;

    code {
      background: none;
      padding: 0;
    }
  }

  :deep(blockquote) {
    border-left: 4px solid var(--n-border-color);
    padding-left: 12px;
    margin: 12px 0;
    color: var(--n-text-color-2);
  }

  :deep(ul), :deep(ol) {
    margin: 12px 0;
    padding-left: 20px;
  }

  :deep(li) {
    margin: 4px 0;
  }

  /* 表格样式优化 */
  :deep(table) {
    width: 100%;
    border-collapse: collapse;
    margin: 16px 0;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 
      0 4px 20px rgba(0, 0, 0, 0.08),
      0 2px 8px rgba(0, 0, 0, 0.06),
      0 0 0 1px rgba(255, 255, 255, 0.1);
    background: var(--n-card-color);
    border: 2px solid var(--n-border-color);
    position: relative;
    backdrop-filter: blur(10px);
    transform: translateZ(0);
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    isolation: isolate;

    &::before {
      content: '';
      position: absolute;
      top: -2px;
      left: -2px;
      right: -2px;
      bottom: -2px;
      background: linear-gradient(45deg, 
        var(--n-primary-color), 
        var(--n-primary-color-hover), 
        var(--n-primary-color));
      border-radius: 14px;
      z-index: -1;
      opacity: 0.1;
      transition: opacity 0.3s ease;
    }

    &:hover {
      box-shadow: 
        0 8px 32px rgba(0, 0, 0, 0.12),
        0 4px 16px rgba(0, 0, 0, 0.08),
        0 0 0 1px rgba(255, 255, 255, 0.15);
      transform: translateY(-2px);

      &::before {
        opacity: 0.2;
      }
    }
  }

  :deep(thead) {
    background: linear-gradient(135deg, var(--n-primary-color), var(--n-primary-color-hover));
    position: relative;

    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: linear-gradient(135deg, rgba(255,255,255,0.1), rgba(255,255,255,0.05));
      pointer-events: none;
    }
  }

  :deep(th) {
    padding: 12px 16px;
    text-align: left;
    font-weight: 600;
    font-size: 14px;
    border-bottom: 2px solid var(--n-primary-color);
    position: relative;

    &:not(:last-child) {
      border-right: 1px solid rgba(255, 255, 255, 0.2);
    }
  }

  :deep(td) {
    padding: 12px 16px;
    border-bottom: 2px solid var(--n-border-color);
    font-size: 14px;
    line-height: 1.5;
    vertical-align: top;
    position: relative;

    &:not(:last-child) {
      border-right: 2px solid var(--n-border-color);
    }

    &::after {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: linear-gradient(135deg, 
        rgba(255, 255, 255, 0.05), 
        rgba(255, 255, 255, 0.02));
      pointer-events: none;
      opacity: 0;
      transition: opacity 0.2s ease;
    }

    code {
      background: var(--n-code-color);
      padding: 2px 6px;
      border-radius: 4px;
      font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
      font-size: 0.85em;
      color: var(--n-text-color);
    }
  }

  :deep(tr) {
    &:last-child td {
      border-bottom: none;
    }

    &:nth-child(even) {
      background: var(--n-color-modal);
      opacity: 0.9;
    }

    &:nth-child(odd) {
      background: var(--n-card-color);
      opacity: 0.95;
    }

    &:hover {
      background: var(--n-color-hover);
      transition: all 0.2s ease;
      transform: translateX(4px);
      box-shadow: 
        inset 0 0 0 1px var(--n-primary-color),
        0 2px 8px rgba(0, 0, 0, 0.1);
      opacity: 1;

      td::after {
        opacity: 1;
      }
    }
  }

  /* 表格响应式设计 */
  @media (max-width: 768px) {
    :deep(table) {
      font-size: 12px;
      margin: 12px 0;
      overflow-x: auto;
      display: block;
    }
    
    :deep(th),
    :deep(td) {
      padding: 8px 10px;
    }
  }
}

/* 悬浮输入区域 */
.floating-input-container {
  position: fixed;
  bottom: 30px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 800px;
  z-index: 1000;
  pointer-events: none;
}

.glassmorphism-input {
  background: var(--n-card-color);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  border: 1px solid var(--n-border-color);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 16px;
  pointer-events: auto;
  transition: all 0.3s ease;
  opacity: 0.95;

  &:hover {
    box-shadow: 0 6px 25px rgba(0, 0, 0, 0.15);
    transform: translateY(-1px);
  }
}

.input-wrapper {
  display: flex;
  gap: 12px;
  align-items: flex-end;
  min-height: 48px;
}

.glass-input {
  flex: 1;
  display: flex;
  align-items: stretch;

  :deep(.n-input__textarea) {
    height: 48px;
  }

  :deep(.n-input-wrapper) {
    display: flex;
    align-items: stretch;
    width: 100%;
  }

  :deep(.n-input) {
    background: transparent;
    border: none;
    box-shadow: none;
    width: 100%;
  }

  :deep(.n-input__textarea-el) {
    background: var(--n-color-modal);
    border: none;
    font-size: 16px;
    color: var(--n-text-color);
    resize: none;
    outline: none;
    width: 100%;
    min-height: 48px;
    max-height: 200px;
    padding: 12px 16px;
    transition: all 0.3s ease;
    line-height: 1.5;
    font-family: inherit;
    caret-color: var(--n-text-color);
    text-align: left;
    vertical-align: top;
    position: relative;
    text-indent: 0;

    &::placeholder {
      color: var(--n-text-color-3);
      text-align: left;
      vertical-align: top;
    }

    &:focus {
      background: var(--n-color-hover);
      caret-color: var(--n-text-color);
      text-indent: 0;
      padding-left: 0;
      padding-right: 0;
      padding-top: 5px;
      padding-bottom: 12px;
    }
  }

  :deep(.n-input__input-el) {
    background: transparent;
    border: none;
  }
}

/* 深色模式特殊处理 */
[data-theme="dark"] {
  .glass-input :deep(.n-input__textarea-el) {
    background: rgba(60, 60, 60, 0.9);
    color: #fff;
    caret-color: #fff;
    text-indent: 0;

    &::placeholder {
      color: #aaa;
      text-align: left;
      vertical-align: top;
    }

    &:focus {
      background: rgba(80, 80, 80, 0.95);
      caret-color: #fff;
    }
  }
}

.glass-send-button {
  height: 48px;
  padding: 0 24px;
  border-radius: 12px;
  background: #18a058;
  color: #fff;
  border: none;
  box-shadow: 0 2px 8px rgba(24, 160, 88, 0.2);
  transition: all 0.3s ease;
  flex-shrink: 0;
  align-self: flex-end;

  &:hover {
    background: #36ad6a;
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(24, 160, 88, 0.3);
  }

  &:disabled {
    opacity: 0.6;
    transform: none;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }
}

/* 深色模式按钮样式 */
[data-theme="dark"] {
  .glass-send-button {
    background: #36ad6a;
    color: #fff;
    box-shadow: 0 2px 8px rgba(54, 173, 106, 0.4);

    &:hover {
      background: #42b883;
      box-shadow: 0 4px 12px rgba(54, 173, 106, 0.5);
    }

    &:disabled {
      background: #555;
      opacity: 0.6;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    }
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .chat-header {
    flex-direction: column;
    gap: 12px;
    align-items: stretch;
  }
  
  .header-right {
    justify-content: center;
  }
  
  .message {
    max-width: 95%;
  }
  
  .floating-input-container {
    width: 95%;
    bottom: 20px;
  }
  
  .glassmorphism-input {
    padding: 12px;
  }
  
  .input-wrapper {
    gap: 8px;
  }
  
  .glass-input :deep(.n-input__textarea-el) {
    font-size: 14px;
    padding: 10px 12px;
    min-height: 44px;
    max-height: 150px;
  }
  
  .glass-send-button {
    height: 44px;
    padding: 0 16px;
    font-size: 14px;
  }
}

.settings-btn {
  transition: all 0.3s ease;
  margin-left: 8px;

  &:hover {
    transform: scale(1.1);
  }
}
</style>