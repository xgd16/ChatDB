import { defineStore } from 'pinia';
import { ref, computed, watch } from 'vue';

// 设置项类型定义
export interface ChatSettings {
    ai: string;
    model: string;
    databaseId: number;
}

// AI 提供商配置
export interface AIProvider {
    label: string;
    value: string;
    models: Array<{
        label: string;
        value: string;
    }>;
}

// 数据库配置
export interface DatabaseOption {
    label: string;
    value: number;
}

export const useSettingsStore = defineStore('settings', () => {
    const SETTINGS_STORAGE_KEY = 'app:chatSettings';

    // 默认设置
    const defaultSettings: ChatSettings = {
        ai: 'openai',
        model: 'gpt-5-mini',
        databaseId: 1
    };

    // 当前设置
    const settings = ref<ChatSettings>({ ...defaultSettings });

    // AI 提供商配置
    const aiProviders = ref<AIProvider[]>([
        {
            label: 'OpenAI',
            value: 'openai',
            models: [
                { label: 'GPT-5 Mini', value: 'gpt-5-mini' },
                { label: 'GPT-5 Nano', value: 'gpt-5-nano' },
                { label: 'GPT-4o Mini', value: 'gpt-4o-mini' },
            ]
        },
        {
            label: 'DeepSeek',
            value: 'deepseek',
            models: [
                { label: 'DeepSeek Chat', value: 'deepseek-chat' },
                { label: 'DeepSeek Coder', value: 'deepseek-coder' }
            ]
        }
    ]);

    // 数据库选项
    const databaseOptions = ref<DatabaseOption[]>([
        { label: '默认数据库', value: 1 }
    ]);

    // 计算属性：当前 AI 提供商的模型选项
    const currentModelOptions = computed(() => {
        const provider = aiProviders.value.find(p => p.value === settings.value.ai);
        return provider?.models || [];
    });

    // 计算属性：AI 提供商选项
    const aiProviderOptions = computed(() => {
        return aiProviders.value.map(p => ({
            label: p.label,
            value: p.value
        }));
    });

    // 初始化设置
    const initSettings = () => {
        const persisted = localStorage.getItem(SETTINGS_STORAGE_KEY);
        if (persisted) {
            try {
                const savedSettings = JSON.parse(persisted);
                settings.value = { ...defaultSettings, ...savedSettings };
            } catch (error) {
                console.error('Failed to parse settings:', error);
                settings.value = { ...defaultSettings };
            }
        }
    };

    // 更新设置
    const updateSettings = (newSettings: Partial<ChatSettings>) => {
        settings.value = { ...settings.value, ...newSettings };
    };

    // 更新 AI 提供商
    const updateAIProvider = (ai: string) => {
        settings.value.ai = ai;
        // 自动选择该提供商的第一个模型
        const provider = aiProviders.value.find(p => p.value === ai);
        if (provider && provider.models.length > 0) {
            settings.value.model = provider.models[0].value;
        }
    };

    // 更新数据库选项
    const updateDatabaseOptions = (options: DatabaseOption[]) => {
        databaseOptions.value = options;
        // 如果当前选择的数据库不在新选项中，选择第一个
        if (!options.find(opt => opt.value === settings.value.databaseId) && options.length > 0) {
            settings.value.databaseId = options[0].value;
        }
    };

    // 重置设置
    const resetSettings = () => {
        settings.value = { ...defaultSettings };
    };

    // 监听设置变化并持久化
    watch(settings, (newSettings) => {
        localStorage.setItem(SETTINGS_STORAGE_KEY, JSON.stringify(newSettings));
    }, { deep: true });

    return {
        // 状态
        settings,
        aiProviders,
        databaseOptions,

        // 计算属性
        currentModelOptions,
        aiProviderOptions,

        // 方法
        initSettings,
        updateSettings,
        updateAIProvider,
        updateDatabaseOptions,
        resetSettings
    };
});
