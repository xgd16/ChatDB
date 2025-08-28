<template>
  <div class="login-container">
    <n-card size="large" class="login-form-wrapper" :bordered="false">
      <div class="logo-area">
        <h1 class="logo-title">问库</h1>
        <p class="logo-subtitle">让数据查询变得简单高效</p>
      </div>
      
      <n-form ref="formRef" :model="formData" :rules="rules" class="login-form">
        <n-form-item label="用户名" path="username" class="form-item">
          <n-input
            v-model:value="formData.username"
            placeholder="请输入用户名"
            clearable
            class="input-item"
          >
            <template #prefix>
              <n-icon><i class="ri-user-line"></i></n-icon>
            </template>
          </n-input>
        </n-form-item>
        
        <n-form-item label="密码" path="password" class="form-item">
          <n-input
            v-model:value="formData.password"
            type="password"
            placeholder="请输入密码"
            clearable
            class="input-item"
          >
            <template #prefix>
              <n-icon><i class="ri-door-lock-line"></i></n-icon>
            </template>
          </n-input>
        </n-form-item>
        
        <!-- <div class="form-actions">
          <n-checkbox v-model:checked="formData.remember" class="remember-checkbox">
            记住我
          </n-checkbox>
          <n-button text class="forgot-btn">忘记密码？</n-button>
        </div> -->
        
        <n-button
          type="primary"
          size="large"
          :loading="loading"
          @click="handleLogin"
          class="login-btn"
          :disabled="loading"
        >
          登录
        </n-button>
      </n-form>
      
      <div class="other-actions">
        <span>还没有账号？</span>
        <n-button text @click="handleRegister">立即注册</n-button>
      </div>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { NForm, NFormItem, NInput, NButton, NCheckbox, NIcon, NCard } from 'naive-ui';
import { loginApi } from '@/api/user';
import { useUserStore } from '@/stores/counter';

const userStore = useUserStore();

const router = useRouter();
const formRef = ref<InstanceType<typeof NForm>>();
const loading = ref(false);

const formData = reactive({
  username: '',
  password: '',
  remember: false
});

const rules = {
  username: [
    {
      required: true,
      message: '请输入用户名',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    },
    {
      min: 6,
      message: '密码长度至少为6位',
      trigger: 'blur'
    }
  ]
};

const handleLogin = async () => {
  if (!formRef.value) return;
  
  try {
    await formRef.value.validate();
    loading.value = true;
    
    loginApi(formData).then(r => {
      if (r.code === 0) {
        userStore.setUserInfo(r.data);
        router.push('/');
      }
    }).finally(() => {
      loading.value = false;
    })
  } catch (error) {
    console.error('登录验证失败:', error);
  }
};

const handleRegister = () => {
  // 注册逻辑
  console.log('跳转到注册页面');
};
</script>

<style scoped>
.login-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  /* 使用全局主题的背景色（由 NGlobalStyle 控制） */
}

.login-form-wrapper {
  width: 100%;
  max-width: 420px;
  padding: 24px;
}

.logo-area {
  text-align: center;
  margin-bottom: 30px;
}

.logo-icon { margin-bottom: 16px; }

.logo-title { font-size: 24px; font-weight: 600; margin: 0 0 8px 0; }

.logo-subtitle { font-size: 14px; margin: 0; }

.login-form {
  width: 100%;
}

.form-item {
  margin-bottom: 20px;
}

.input-item {
  width: 100%;
  height: 40px;
  line-height: 40px;
}

.form-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.remember-checkbox {
  font-size: 14px;
}

.forgot-btn {
  font-size: 14px;
}

.login-btn {
  width: 100%;
  height: 44px;
  font-size: 16px;
}

.other-actions { text-align: center; margin-top: 20px; font-size: 14px; }

@media (max-width: 480px) {
  .login-form-wrapper {
    padding: 30px 20px;
    margin: 0 20px;
  }
  
  .logo-title {
    font-size: 20px;
  }
}
</style>