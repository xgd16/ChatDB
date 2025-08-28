<template>
  <n-drawer v-model:show="show" :width="400" placement="right">
    <n-drawer-content title="聊天设置" closable>
      <n-form label-placement="left" label-width="auto" require-mark-placement="right-hanging">
        <!-- AI 提供商设置 -->
        <n-form-item label="AI 提供商">
          <n-select
            :value="settingsStore.settings.ai"
            :options="settingsStore.aiProviderOptions"
            placeholder="选择AI模型"
            style="width: 100%;"
            @update:value="handleAIProviderChange"
          />
        </n-form-item>
        
        <!-- 模型版本设置 -->
        <n-form-item label="模型版本">
          <n-select
            :value="settingsStore.settings.model"
            :options="settingsStore.currentModelOptions"
            placeholder="选择模型"
            style="width: 100%;"
            @update:value="handleModelChange"
          />
        </n-form-item>
        
        <n-divider />
        
        <!-- 数据库配置设置 -->
        <n-form-item label="数据库配置">
          <n-select
            :value="settingsStore.settings.databaseId"
            :options="settingsStore.databaseOptions"
            placeholder="选择数据库"
            style="width: 100%;"
            @update:value="handleDatabaseChange"
          />
        </n-form-item>
        
        <n-divider />
        
        <!-- 主题模式设置 -->
        <n-form-item label="主题模式">
          <n-button
            quaternary
            @click="handleThemeToggle"
            style="width: 100%; justify-content: flex-start;"
          >
            <template #icon>
              <i :class="themeStore.isDark ? 'ri-sun-line' : 'ri-moon-line'"></i>
            </template>
            {{ themeStore.isDark ? '切换到浅色模式' : '切换到深色模式' }}
          </n-button>
        </n-form-item>
        
        <n-divider />
        
        <!-- 操作按钮 -->
        <n-form-item>
          <div style="display: flex; gap: 12px;">
            <n-button
              secondary
              @click="handleResetSettings"
              style="flex: 1;"
            >
              重置设置
            </n-button>
            <n-button
              type="primary"
              @click="handleSaveSettings"
              style="flex: 1;"
            >
              保存设置
            </n-button>
          </div>
        </n-form-item>
      </n-form>
    </n-drawer-content>
  </n-drawer>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { 
  NDrawer, 
  NDrawerContent, 
  NForm, 
  NFormItem, 
  NDivider,
  NSelect,
  NButton,
  useMessage
} from 'naive-ui';
import { useSettingsStore } from '@/stores/settings';
import { useThemeStore } from '@/stores/theme';

interface Props {
  show: boolean;
}

interface Emits {
  (e: 'update:show', value: boolean): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const message = useMessage();
const settingsStore = useSettingsStore();
const themeStore = useThemeStore();

// 计算属性：控制抽屉显示
const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

// 处理 AI 提供商变化
const handleAIProviderChange = (value: string) => {
  settingsStore.updateAIProvider(value);
  message.success('AI 提供商已更新');
};

// 处理模型变化
const handleModelChange = (value: string) => {
  settingsStore.updateSettings({ model: value });
  message.success('模型版本已更新');
};

// 处理数据库变化
const handleDatabaseChange = (value: number) => {
  settingsStore.updateSettings({ databaseId: value });
  message.success('数据库配置已更新');
};

// 处理主题切换
const handleThemeToggle = () => {
  themeStore.toggleTheme();
  message.success(`已切换到${themeStore.isDark ? '深色' : '浅色'}模式`);
};

// 重置设置
const handleResetSettings = () => {
  settingsStore.resetSettings();
  message.success('设置已重置为默认值');
};

// 保存设置
const handleSaveSettings = () => {
  // 设置已经自动保存，这里只是给用户反馈
  message.success('设置已保存');
  show.value = false;
};
</script>
