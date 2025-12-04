import { defineStore } from 'pinia';
import { ref, watch } from 'vue';

export const useThemeStore = defineStore('theme', () => {
  const THEME_STORAGE_KEY = 'app:isDarkTheme';
  const isDark = ref<boolean>(false);

  // 初始化主题（使用系统偏好或持久化配置）
  const initTheme = () => {
    const persisted = window.localStorage.getItem(THEME_STORAGE_KEY);
    if (persisted === 'true' || persisted === 'false') {
      isDark.value = persisted === 'true';
    } else if (window.matchMedia) {
      isDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches;
    } else {
      isDark.value = false;
    }
    document.documentElement.classList.toggle('dark', isDark.value);
  };

  const toggleTheme = () => {
    isDark.value = !isDark.value;
  };

  const setTheme = (dark: boolean) => {
    isDark.value = dark;
  };

  watch(isDark, (val) => {
    window.localStorage.setItem(THEME_STORAGE_KEY, String(val));
    document.documentElement.classList.toggle('dark', val);
  });

  return {
    isDark,
    initTheme,
    toggleTheme,
    setTheme,
  };
});
