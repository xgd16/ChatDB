<template>
  <div class="app-container">
    <!-- 顶部导航栏 -->
    <header class="app-header">
      <div class="header-left">
        <div class="logo">
          <div class="logo-icon">
            <i class="ri-database-2-fill"></i>
          </div>
          <div class="logo-text">
            <h1 class="app-title">问库</h1>
            <span class="app-desc">AI SQL Assistant</span>
          </div>
        </div>
      </div>
      
      <div class="header-right">
        <el-button circle text @click="toggleTheme">
          <i :class="isDark ? 'ri-sun-line' : 'ri-moon-line'"></i>
        </el-button>
        <el-button circle text @click="showSettings = true">
          <i class="ri-settings-3-line"></i>
        </el-button>
        <el-divider direction="vertical" />
        <el-button text bg class="logout-btn" @click="handleLogout">
          <i class="ri-logout-box-r-line"></i>
          <span>退出</span>
        </el-button>
      </div>
    </header>

    <!-- 主体内容区 -->
    <main class="app-main">
      <!-- 聊天区域 -->
      <div class="chat-section" :class="{ 'has-sidebar': showToolPanel }">
        <div class="chat-messages" ref="chatContainerRef">
          <div v-if="messages.length === 0" class="welcome-screen">
            <div class="welcome-content">
              <div class="welcome-icon">
                <i class="ri-robot-2-line"></i>
              </div>
              <h2>有什么可以帮您的吗？</h2>
              <p>我可以帮您查询数据库、分析数据结构或生成 SQL 语句。</p>
              <div class="suggestion-chips">
                <div class="chip" @click="quickAsk('查询最近注册的用户')">查询最近注册的用户</div>
                <div class="chip" @click="quickAsk('统计各地区订单数量')">统计各地区订单数量</div>
                <div class="chip" @click="quickAsk('分析商品销售趋势')">分析商品销售趋势</div>
              </div>
            </div>
          </div>

          <div v-else class="message-list">
            <template v-for="(message, index) in messages" :key="index">
              <!-- 用户消息始终显示；助手消息只有有内容才显示（加载状态单独处理） -->
              <div
                v-if="message.role === 'user' || (message.role === 'assistant' && message.content)"
                class="message-row"
                :class="message.role"
              >
                <div class="avatar">
                  <UserAvatar v-if="message.role === 'user'" />
                  <AIAvatar v-else />
                </div>
                <div class="message-bubble">
                  <!-- 当内容为空时显示一个空格占位，或者直接显示空div，避免高度坍塌 -->
                  <div class="bubble-content markdown-body" v-html="renderMarkdown(message.content || ' ')"></div>
                  <div class="message-meta" v-if="message.content || message.role === 'user'">
                    <span class="time">{{ formatTime(message.timestamp) }}</span>
                  </div>
                </div>
              </div>
            </template>

            <!-- Loading State - 只在AI还没有任何输出时显示 -->
            <div v-if="isLoading && messages.length > 0 && messages[messages.length - 1].role === 'assistant' && !messages[messages.length - 1].content" class="message-row assistant loading">
              <div class="avatar">
                <AIAvatar />
              </div>
              <div class="message-bubble">
                <div class="typing-indicator">
                  <span></span>
                  <span></span>
                  <span></span>
                </div>
                <span class="loading-text">AI 正在思考...</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 输入区域 -->
        <div class="input-section">
          <div class="input-wrapper">
            <el-input
              v-model="inputMessage"
              type="textarea"
              :autosize="{ minRows: 1, maxRows: 8 }"
              placeholder="输入您的问题..."
              class="chat-input"
              @keydown.ctrl.enter="handleSendMessage"
            />
            <div class="input-actions">
              <el-tooltip content="工具面板" placement="top">
                <el-button circle text @click="showToolPanel = !showToolPanel" :type="showToolPanel ? 'primary' : ''">
                  <i class="ri-tools-line"></i>
                </el-button>
              </el-tooltip>
              <el-button type="primary" circle :disabled="!inputMessage.trim() || isLoading" @click="handleSendMessage">
                <i class="ri-send-plane-fill"></i>
              </el-button>
            </div>
          </div>
          <div class="input-footer">
            <span>Ctrl + Enter 发送，Enter 换行</span>
          </div>
        </div>
      </div>

      <!-- 右侧工具面板 -->
      <transition name="slide-right">
        <aside v-if="showToolPanel" class="tool-sidebar">
          <ToolUsagePanel
            :events="toolEvents"
            :loading="isLoading"
            @close="showToolPanel = false"
          />
        </aside>
      </transition>
    </main>

    <!-- 设置抽屉 -->
    <ChatSettings v-model:show="showSettings" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick, watch, computed } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage, ElButton, ElInput, ElDivider, ElTooltip } from 'element-plus';
import { marked } from 'marked';
import ChatSettings from '@/components/ChatSettings.vue';
import ToolUsagePanel from '@/components/ToolUsagePanel.vue';
import UserAvatar from '@/components/icons/UserAvatar.vue';
import AIAvatar from '@/components/icons/AIAvatar.vue';
import { chatApi } from '@/api/ai';
import { useSettingsStore } from '@/stores/settings';
import { useUserStore } from '@/stores/counter';
import { useThemeStore } from '@/stores/theme';
import type { ChatMessage, ChatStreamEvent } from '@/types/ai';

const router = useRouter();
const settingsStore = useSettingsStore();
const userStore = useUserStore();
const themeStore = useThemeStore();

const chatContainerRef = ref<HTMLElement>();
const inputMessage = ref('');
const messages = ref<ChatMessage[]>([]);
const isLoading = ref(false);
const showSettings = ref(false);
const showToolPanel = ref(true);
const toolEvents = ref<Array<{ name: string; output?: any; timestamp?: number }>>([]);

const isDark = computed(() => themeStore.isDark);

const toggleTheme = () => {
  themeStore.toggleTheme();
};

const quickAsk = (text: string) => {
  inputMessage.value = text;
  handleSendMessage();
};

// 初始化设置
onMounted(() => {
  settingsStore.initSettings();
  themeStore.initTheme();
  loadDatabaseConfigs();
  loadAIModels();
});

// 监听消息变化，自动滚动到底部
watch(messages, () => {
  nextTick(() => {
    scrollToBottom();
  });
}, { deep: true });

// 监听设置面板关闭，刷新数据库配置列表
watch(showSettings, (newVal, oldVal) => {
  // 当设置面板从打开变为关闭时，刷新数据库配置列表
  if (oldVal === true && newVal === false) {
    loadDatabaseConfigs();
  }
});

// 滚动到底部
const scrollToBottom = () => {
  if (chatContainerRef.value) {
    chatContainerRef.value.scrollTo({
      top: chatContainerRef.value.scrollHeight,
      behavior: 'smooth'
    });
  }
};

// 渲染 Markdown
const renderMarkdown = (content: string) => {
  try {
    return marked(content || '');
  } catch {
    return content || '';
  }
};

// 格式化时间
const formatTime = (timestamp: number) => {
  const date = new Date(timestamp);
  return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' });
};

// 发送消息
const handleSendMessage = async () => {
  if (!inputMessage.value.trim() || isLoading.value) return;

  const userMessage: ChatMessage = {
    role: 'user',
    content: inputMessage.value.trim(),
    timestamp: Date.now()
  };

  messages.value.push(userMessage);
  const currentInput = inputMessage.value.trim();
  inputMessage.value = '';
  isLoading.value = true;
  toolEvents.value = [];

  // 如果工具面板被隐藏且有新消息，是否自动打开？暂时不自动打开，由用户控制

  try {
    // 检查必要的配置：当前选择的数据库是否存在于列表中
    const hasValidDatabase =
      settingsStore.databaseOptions.length > 0 &&
      settingsStore.databaseOptions.some(
        (opt) => opt.value === settingsStore.settings.databaseId
      );

    if (!hasValidDatabase) {
      ElMessage.warning('请先选择数据库配置');
      isLoading.value = false;
      return;
    }

    const response = await chatApi({
      ai: settingsStore.settings.ai,
      model: settingsStore.settings.model,
      message: currentInput,
      databaseId: settingsStore.settings.databaseId
    });

    if (!response.ok) {
      const errorText = await response.text().catch(() => '请求失败');
      throw new Error(errorText || '请求失败');
    }

    const reader = response.body?.getReader();
    const decoder = new TextDecoder();
    let assistantMessage: ChatMessage = {
      role: 'assistant',
      content: '',
      timestamp: Date.now()
    };
    messages.value.push(assistantMessage);
    const assistantMessageIndex = messages.value.length - 1;

    let buffer = '';
    let traceId = '';

    while (reader) {
      const { done, value } = await reader.read();
      if (done) break;

      buffer += decoder.decode(value, { stream: true });
      const lines = buffer.split('\n');
      buffer = lines.pop() || '';

      for (const line of lines) {
        if (line.trim() === '') continue;
        
        if (line.startsWith('data: ')) {
          const dataStr = line.slice(6);
          if (dataStr === '[DONE]') continue;

          try {
            const event: ChatStreamEvent = JSON.parse(dataStr);
            
            if (event.traceId) {
              traceId = event.traceId;
            }

            if (event.event === 'message' && event.content) {
              // 将 SSE 新内容直接添加到消息（不使用打字机效果，以便实时显示）
              const message = messages.value[assistantMessageIndex];
              if (message && message.role === 'assistant') {
                message.content += event.content;
              }
              nextTick(() => {
                scrollToBottom();
              });
            } else if (event.event === 'tool_call') {
              // 处理工具调用事件
              if (event.data) {
                const toolEvent: any = {
                  name: event.data.name || '未知工具',
                  output: event.data.output,
                  timestamp: event.createTime || Date.now()
                };
                
                // 如果是 SQL_Actuator 工具，从输出中提取 SQL 语句
                if (event.data.name === 'SQL_Actuator' && event.data.output) {
                  const outputText = String(event.data.output);
                  // 尝试从输出中提取 SQL 语句（格式：```sql\nSELECT ...```）
                  const sqlMatch = outputText.match(/```sql\n([\s\S]*?)\n```/);
                  if (sqlMatch && sqlMatch[1]) {
                    toolEvent.sql = sqlMatch[1].trim();
                    // 移除 SQL 部分，只保留结果
                    const resultMatch = outputText.match(/\*\*执行结果：\*\*\n\n([\s\S]*)/);
                    if (resultMatch && resultMatch[1]) {
                      toolEvent.output = resultMatch[1];
                    }
                  }
                }
                
                // 如果是导出工具，从输出中提取下载链接
                if ((event.data.name === 'ExportToExcel' || event.data.name === 'ExportData') && event.data.output) {
                  try {
                    const outputText = String(event.data.output);
                    const exportData = JSON.parse(outputText);
                    if (exportData.downloadUrl) {
                      toolEvent.downloadUrl = exportData.downloadUrl;
                      toolEvent.fileName = exportData.fileName;
                      toolEvent.message = exportData.message;
                    }
                  } catch (e) {
                    console.error('解析导出数据失败:', e);
                  }
                }
                
                toolEvents.value.push(toolEvent);
              }
            }
          } catch (e) {
            console.error('解析事件失败:', e);
          }
        }
      }
    }

    scrollToBottom();
  } catch (error: any) {
    console.error('发送消息失败:', error);
    ElMessage.error(error.message || '发送消息失败，请重试');
    messages.value.pop(); // 移除用户消息
  } finally {
    isLoading.value = false;
  }
};

// 加载数据库配置列表
const loadDatabaseConfigs = async () => {
  try {
    const { getDatabaseConfigList } = await import('@/api/ai');
    const res = await getDatabaseConfigList({ page: 1, size: 100 });
    if (res && res.code === 0 && res.data) {
      const list = res.data.list || [];
      if (list.length > 0) {
        const options = list.map((db: any) => ({
          label: `${db.dbName || '未命名'} (${db.dbType || 'unknown'})`,
          value: db.databaseId || 0
        }));
        settingsStore.updateDatabaseOptions(options);
      } else {
        // 如果没有数据库配置，清空选项
        settingsStore.updateDatabaseOptions([]);
      }
    }
  } catch (error) {
    console.error('加载数据库配置失败:', error);
    ElMessage.warning('加载数据库配置失败，请稍后重试');
  }
};

// 加载 AI 模型列表
const loadAIModels = async () => {
  try {
    const { getAiModelList } = await import('@/api/ai');
    const res = await getAiModelList();
    if (res && res.code === 0 && res.data) {
      const list = res.data.list || [];
      if (list.length > 0) {
        // 更新 AI 模型选项
        const models = list.map((item: any) => ({
          label: item.id || '未知模型',
          value: item.id || ''
        }));
        
        // 如果当前 AI 提供商是 openai，更新模型列表
        if (settingsStore.settings.ai === 'openai') {
          const openaiProvider = settingsStore.aiProviders.find(p => p.value === 'openai');
          if (openaiProvider && models.length > 0) {
            openaiProvider.models = models;
            // 如果当前模型不在新列表中，选择第一个
            if (!models.find((m: any) => m.value === settingsStore.settings.model)) {
              settingsStore.updateSettings({ model: models[0].value });
            }
          }
        }
      }
    }
  } catch (error) {
    console.error('加载 AI 模型列表失败:', error);
    // AI 模型列表加载失败不影响使用，使用默认模型
  }
};

// 退出登录
const handleLogout = () => {
  userStore.clearUserInfo();
  router.push('/login');
};
</script>

<style scoped lang="scss">
.app-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--el-bg-color);
  color: var(--el-text-color-primary);
}

/* Header Styles */
.app-header {
  height: 64px;
  padding: 0 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid var(--el-border-color-light);
  background-color: var(--el-bg-color-overlay);
  z-index: 10;
}

.header-left {
  display: flex;
  align-items: center;
}

.logo {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-icon {
  width: 36px;
  height: 36px;
  background: linear-gradient(135deg, var(--el-color-primary), var(--el-color-primary-light-3));
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.logo-text {
  display: flex;
  flex-direction: column;
}

.app-title {
  font-size: 18px;
  font-weight: 700;
  line-height: 1.2;
  margin: 0;
  background: linear-gradient(to right, var(--el-text-color-primary), var(--el-text-color-secondary));
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.app-desc {
  font-size: 12px;
  color: var(--el-text-color-secondary);
}

.header-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.logout-btn {
  display: flex;
  align-items: center;
  gap: 4px;
}

/* Main Content */
.app-main {
  flex: 1;
  display: flex;
  overflow: hidden;
  position: relative;
}

/* Chat Section */
.chat-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  position: relative;
  transition: all 0.3s ease;
  max-width: 100%;
}

.chat-section.has-sidebar {
  margin-right: 0; /* Controlled by flexbox */
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
  scroll-behavior: smooth;
}

/* Welcome Screen */
.welcome-screen {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--el-text-color-primary);
}

.welcome-content {
  text-align: center;
  max-width: 600px;
  padding: 0 20px;
}

.welcome-icon {
  font-size: 64px;
  color: var(--el-color-primary);
  margin-bottom: 24px;
  opacity: 0.8;
}

.welcome-content h2 {
  font-size: 24px;
  margin-bottom: 12px;
  font-weight: 600;
}

.welcome-content p {
  color: var(--el-text-color-secondary);
  margin-bottom: 32px;
  font-size: 16px;
}

.suggestion-chips {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 12px;
}

.chip {
  padding: 8px 16px;
  background: var(--el-fill-color-light);
  border: 1px solid var(--el-border-color-lighter);
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 14px;
  color: var(--el-text-color-regular);
}

.chip:hover {
  background: var(--el-color-primary-light-9);
  border-color: var(--el-color-primary-light-5);
  color: var(--el-color-primary);
  transform: translateY(-2px);
}

/* Message List */
.message-list {
  max-width: 900px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.message-row {
  display: flex;
  gap: 16px;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.message-row.user {
  flex-direction: row-reverse;
}

.avatar {
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  overflow: hidden;
  background: var(--el-fill-color);
  display: flex;
  align-items: center;
  justify-content: center;
}

.message-bubble {
  max-width: 85%;
  padding: 12px 16px;
  border-radius: 16px;
  background: var(--el-fill-color-light);
  position: relative;
  overflow: hidden;
}

.message-row.user .message-bubble {
  background: var(--el-color-primary);
  color: white;
  border-bottom-right-radius: 4px;
}

.message-row.assistant .message-bubble {
  background: var(--el-bg-color-overlay);
  border: 1px solid var(--el-border-color-lighter);
  border-bottom-left-radius: 4px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}

.bubble-content {
  line-height: 1.6;
  font-size: 15px;
  word-break: break-word;
  overflow: hidden;
}

/* Markdown Styles Override within bubble */
.bubble-content :deep(p) {
  margin-bottom: 8px;
  margin-top: 0;
}

.bubble-content :deep(p:last-child) {
  margin-bottom: 0;
}

/* 下载链接按钮样式（针对导出接口地址） */
.bubble-content :deep(a[href^="/api/ai_chat/v1/export/download"]) {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 14px;
  margin-top: 4px;
  background: var(--el-color-primary);
  color: #fff;
  border-radius: 999px;
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  white-space: nowrap;
}

.bubble-content :deep(a[href^="/api/ai_chat/v1/export/download"]::before) {
  content: "⇩";
  font-size: 14px;
}

.bubble-content :deep(a[href^="/api/ai_chat/v1/export/download"]:hover) {
  background: var(--el-color-primary-dark-2);
}

/* Code blocks */
.bubble-content :deep(pre) {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 8px;
  padding: 12px;
  overflow-x: auto;
  margin: 8px 0;
  max-width: 100%;
}

.user .bubble-content :deep(pre) {
  background: rgba(255, 255, 255, 0.15);
}

/* Tables */
.bubble-content :deep(table) {
  border-collapse: collapse;
  width: 100%;
  margin: 12px 0;
  font-size: 13px;
}

.bubble-content :deep(th),
.bubble-content :deep(td) {
  border: 1px solid var(--el-border-color-light);
  padding: 8px 12px;
  text-align: left;
  word-break: break-word;
}

.bubble-content :deep(th) {
  background: var(--el-fill-color-light);
  font-weight: 600;
}

.message-row.user .bubble-content :deep(th) {
  background: rgba(255, 255, 255, 0.2);
}

.bubble-content :deep(tr:nth-child(even)) {
  background: rgba(0, 0, 0, 0.02);
}

.message-row.user .bubble-content :deep(tr:nth-child(even)) {
  background: rgba(255, 255, 255, 0.1);
}

/* Lists */
.bubble-content :deep(ul),
.bubble-content :deep(ol) {
  margin: 8px 0;
  padding-left: 24px;
}

.bubble-content :deep(li) {
  margin: 4px 0;
  word-break: break-word;
}

.bubble-content :deep(ul ul),
.bubble-content :deep(ol ol),
.bubble-content :deep(ul ol),
.bubble-content :deep(ol ul) {
  margin: 4px 0;
}

/* Headings */
.bubble-content :deep(h1),
.bubble-content :deep(h2),
.bubble-content :deep(h3),
.bubble-content :deep(h4),
.bubble-content :deep(h5),
.bubble-content :deep(h6) {
  margin-top: 12px;
  margin-bottom: 8px;
  font-weight: 600;
  line-height: 1.4;
}

.bubble-content :deep(h1:first-child),
.bubble-content :deep(h2:first-child),
.bubble-content :deep(h3:first-child),
.bubble-content :deep(h4:first-child),
.bubble-content :deep(h5:first-child),
.bubble-content :deep(h6:first-child) {
  margin-top: 0;
}

/* Inline code */
.bubble-content :deep(code) {
  background: rgba(0, 0, 0, 0.05);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: 'Monaco', 'Menlo', 'Courier New', monospace;
  font-size: 0.9em;
  word-break: break-word;
}

.user .bubble-content :deep(code) {
  background: rgba(255, 255, 255, 0.2);
}

/* Blockquotes */
.bubble-content :deep(blockquote) {
  border-left: 4px solid var(--el-color-primary);
  padding-left: 12px;
  margin: 8px 0;
  opacity: 0.8;
  font-style: italic;
}

.message-meta {
  margin-top: 6px;
  font-size: 11px;
  opacity: 0.6;
  text-align: right;
}

.typing-indicator {
  display: flex;
  gap: 4px;
  padding: 4px 0;
  align-items: center;
}

.typing-indicator span {
  width: 6px;
  height: 6px;
  background: var(--el-text-color-secondary);
  border-radius: 50%;
  animation: bounce 1.4s infinite ease-in-out both;
}

.typing-indicator span:nth-child(1) { animation-delay: -0.32s; }
.typing-indicator span:nth-child(2) { animation-delay: -0.16s; }

@keyframes bounce {
  0%, 80%, 100% { transform: scale(0); }
  40% { transform: scale(1); }
}

.loading-text {
  margin-left: 8px;
  font-size: 13px;
  color: var(--el-text-color-secondary);
}

/* Input Section */
.input-section {
  flex-shrink: 0;
  padding: 16px 24px 20px;
  background: var(--el-bg-color);
  border-top: 1px solid var(--el-border-color-lighter);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  z-index: 5;
}

.input-actions :deep(.el-button) {
  --el-button-size: 32px;
  width: 32px;
  height: 32px;
  padding: 0;
  min-width: unset;
}

.input-actions :deep(.el-button--text) {
  color: var(--el-text-color-secondary);
}

.input-actions :deep(.el-button--text:hover) {
  color: var(--el-color-primary);
}

.input-wrapper {
  width: 100%;
  max-width: 800px;
  position: relative;
  background: var(--el-bg-color-overlay);
  border: 1px solid var(--el-border-color);
  border-radius: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.06);
  padding: 10px 12px;
  display: flex;
  align-items: flex-end;
  gap: 8px;
  transition: all 0.3s;
}

.input-wrapper:focus-within {
  border-color: var(--el-color-primary);
  box-shadow: 0 2px 16px rgba(var(--el-color-primary-rgb), 0.12);
}

.chat-input {
  flex: 1;
}

.chat-input :deep(.el-textarea__inner) {
  box-shadow: none !important;
  background: transparent;
  border: none;
  resize: none;
  padding: 6px 4px;
  font-size: 15px;
  line-height: 1.5;
}

.input-actions {
  display: flex;
  gap: 4px;
  align-items: center;
  flex-shrink: 0;
}

.input-footer {
  margin-top: 8px;
  font-size: 12px;
  color: var(--el-text-color-secondary);
  opacity: 0.6;
}

/* Tool Sidebar */
.tool-sidebar {
  width: 320px;
  border-left: 1px solid var(--el-border-color-light);
  background: var(--el-bg-color-page);
  display: flex;
  flex-direction: column;
}

/* Transitions */
.slide-right-enter-active,
.slide-right-leave-active {
  transition: transform 0.3s ease, width 0.3s ease, opacity 0.3s ease;
}

.slide-right-enter-from,
.slide-right-leave-to {
  width: 0;
  transform: translateX(100%);
  opacity: 0;
}

@media (max-width: 768px) {
  .chat-section {
    margin-right: 0;
  }
  
  .tool-sidebar {
    position: absolute;
    right: 0;
    top: 0;
    bottom: 0;
    z-index: 20;
    box-shadow: -4px 0 16px rgba(0, 0, 0, 0.1);
  }

  .input-wrapper {
    border-radius: 16px;
  }

  .message-list {
    max-width: 100%;
  }

  .message-bubble {
    max-width: 100%;
  }

  .bubble-content :deep(table) {
    font-size: 12px;
  }

  .bubble-content :deep(th),
  .bubble-content :deep(td) {
    padding: 6px 8px;
  }

  .bubble-content :deep(pre) {
    padding: 8px;
    font-size: 12px;
  }
}
</style>