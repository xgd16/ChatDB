import { defineStore } from 'pinia';
import { ref, watch } from 'vue';
import { useOsTheme } from 'naive-ui';

export const useThemeStore = defineStore('theme', () => {
    const THEME_STORAGE_KEY = 'app:isDarkTheme';
    const osThemeRef = useOsTheme();
    const isDark = ref<boolean>(false);

    // 初始化主题
    const initTheme = () => {
        const persisted = window.localStorage.getItem(THEME_STORAGE_KEY);
        if (persisted === 'true' || persisted === 'false') {
            isDark.value = persisted === 'true';
        } else {
            isDark.value = osThemeRef.value === 'dark';
        }
    };

    // 切换主题
    const toggleTheme = () => {
        isDark.value = !isDark.value;
    };

    // 设置主题
    const setTheme = (dark: boolean) => {
        isDark.value = dark;
    };

    // 监听主题变化并持久化
    watch(isDark, (val) => {
        window.localStorage.setItem(THEME_STORAGE_KEY, String(val));
    });

    return {
        isDark,
        initTheme,
        toggleTheme,
        setTheme
    };
});
