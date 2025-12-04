<template>
  <el-drawer v-model="show" title="聊天设置" size="600px" direction="rtl">
    <el-tabs v-model="activeTab" type="border-card">
      <!-- 基础设置 -->
      <el-tab-pane label="基础设置" name="basic">
        <el-form label-position="left" label-width="auto" style="margin-top: 16px;">
          <!-- AI 提供商设置 -->
          <el-form-item label="AI 提供商">
            <el-select
              v-model="settingsStore.settings.ai"
              placeholder="选择 AI 提供商"
              style="width: 100%;"
              @change="handleAIProviderChange"
            >
              <el-option
                v-for="item in settingsStore.aiProviderOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          
          <!-- 模型版本设置 -->
          <el-form-item label="模型版本">
            <el-select
              v-model="settingsStore.settings.model"
              placeholder="选择模型"
              style="width: 100%;"
              @change="handleModelChange"
            >
              <el-option
                v-for="item in settingsStore.currentModelOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          
          <el-divider />
          
          <!-- 数据库配置设置 -->
          <el-form-item label="当前数据库">
            <el-select
              v-model="settingsStore.settings.databaseId"
              placeholder="选择数据库"
              style="width: 100%;"
              @change="handleDatabaseChange"
            >
              <el-option
                v-for="item in settingsStore.databaseOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          
          <el-divider />
          
          <!-- 操作按钮 -->
          <el-form-item>
            <div style="display: flex; gap: 12px;">
              <el-button
                plain
                @click="handleResetSettings"
                style="flex: 1;"
              >
                重置设置
              </el-button>
              <el-button
                type="primary"
                @click="handleSaveSettings"
                style="flex: 1;"
              >
                保存设置
              </el-button>
            </div>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <!-- 数据库配置管理 -->
      <el-tab-pane label="数据库配置" name="database">
        <DatabaseConfigManager @refresh="handleDatabaseRefresh" />
      </el-tab-pane>
    </el-tabs>
  </el-drawer>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import {
  ElDrawer,
  ElTabs,
  ElTabPane,
  ElForm,
  ElFormItem,
  ElDivider,
  ElSelect,
  ElOption,
  ElButton,
  ElMessage
} from 'element-plus';
import { useSettingsStore } from '@/stores/settings';
import DatabaseConfigManager from './DatabaseConfigManager.vue';

interface Props {
  show: boolean;
}

interface Emits {
  (e: 'update:show', value: boolean): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const settingsStore = useSettingsStore();
const activeTab = ref('basic');

// 计算属性：控制抽屉显示
const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

// 处理 AI 提供商变化
const handleAIProviderChange = (value: string) => {
  settingsStore.updateAIProvider(value);
  ElMessage.success('AI 提供商已更新');
};

// 处理模型变化
const handleModelChange = (value: string) => {
  settingsStore.updateSettings({ model: value });
  ElMessage.success('模型版本已更新');
};

// 处理数据库变化
const handleDatabaseChange = (value: number) => {
  settingsStore.updateSettings({ databaseId: value });
  ElMessage.success('数据库配置已更新');
};

// 处理数据库配置刷新
const handleDatabaseRefresh = () => {
  // 数据库配置管理组件已经自动刷新了数据库选项列表
  // 这里可以添加其他刷新逻辑，比如通知父组件
  ElMessage.success('数据库配置已更新');
};

// 重置设置
const handleResetSettings = () => {
  settingsStore.resetSettings();
  ElMessage.success('设置已重置为默认值');
};

// 保存设置
const handleSaveSettings = () => {
  // 设置已经自动保存，这里只是给用户反馈
  ElMessage.success('设置已保存');
  show.value = false;
};
</script>
