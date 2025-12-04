<template>
  <div class="login-container">
    <el-card class="login-form-wrapper" shadow="always">
      <div class="logo-area">
        <h1 class="logo-title">问库</h1>
        <p class="logo-subtitle">让数据查询变得简单高效</p>
      </div>
      
      <el-form ref="formRef" :model="formData" :rules="rules" class="login-form" label-position="top">
        <el-form-item label="用户名" prop="username" class="form-item">
          <el-input
            v-model="formData.username"
            placeholder="请输入用户名"
            clearable
            class="input-item"
          >
            <template #prefix>
              <i class="ri-user-line"></i>
            </template>
          </el-input>
        </el-form-item>
        
        <el-form-item label="密码" prop="password" class="form-item">
          <el-input
            v-model="formData.password"
            type="password"
            placeholder="请输入密码"
            clearable
            class="input-item"
          >
            <template #prefix>
              <i class="ri-door-lock-line"></i>
            </template>
          </el-input>
        </el-form-item>
        
        <el-button
          type="primary"
          size="large"
          :loading="loading"
          @click="handleLogin"
          class="login-btn"
          :disabled="loading"
        >
          登录
        </el-button>
      </el-form>
      
      <div class="other-actions">
        <span>还没有账号？</span>
        <el-button link @click="showRegisterModal = true">立即注册</el-button>
      </div>
    </el-card>

    <!-- 注册模态框 -->
    <el-dialog v-model="showRegisterModal" title="注册账号" width="420px">
      <el-form ref="registerFormRef" :model="registerData" :rules="registerRules" class="register-form" label-position="top">
        <el-form-item label="注册密钥" prop="key" class="form-item">
          <el-input
            v-model="registerData.key"
            placeholder="请输入注册密钥"
            clearable
            class="input-item"
          >
            <template #prefix>
              <i class="ri-key-line"></i>
            </template>
          </el-input>
        </el-form-item>
        
        <el-form-item label="用户名" prop="username" class="form-item">
          <el-input
            v-model="registerData.username"
            placeholder="请输入用户名"
            clearable
            class="input-item"
          >
            <template #prefix>
              <i class="ri-user-line"></i>
            </template>
          </el-input>
        </el-form-item>
        
        <el-form-item label="密码" prop="password" class="form-item">
          <el-input
            v-model="registerData.password"
            type="password"
            placeholder="请输入密码（至少6位）"
            clearable
            class="input-item"
          >
            <template #prefix>
              <i class="ri-door-lock-line"></i>
            </template>
          </el-input>
        </el-form-item>
        
        <el-form-item label="确认密码" prop="confirmPassword" class="form-item">
          <el-input
            v-model="registerData.confirmPassword"
            type="password"
            placeholder="请再次输入密码"
            clearable
            class="input-item"
          >
            <template #prefix>
              <i class="ri-door-lock-line"></i>
            </template>
          </el-input>
        </el-form-item>
        
        <el-button
          type="primary"
          size="large"
          :loading="registerLoading"
          @click="handleRegister"
          class="register-btn"
          :disabled="registerLoading"
          style="width: 100%;"
        >
          注册
        </el-button>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { ElForm, ElFormItem, ElInput, ElButton, ElCard, ElDialog, ElMessage } from 'element-plus';
import { loginApi, registerApi } from '@/api/user';
import { useUserStore } from '@/stores/counter';

const userStore = useUserStore();
const message = ElMessage;

const router = useRouter();
const formRef = ref<InstanceType<typeof ElForm>>();
const registerFormRef = ref<InstanceType<typeof ElForm>>();
const loading = ref(false);
const registerLoading = ref(false);
const showRegisterModal = ref(false);

const formData = reactive({
  username: '',
  password: ''
});

const registerData = reactive({
  key: '',
  username: '',
  password: '',
  confirmPassword: ''
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

const validatePassword = (rule: any, value: string) => {
  if (!value) {
    return new Error('请输入密码');
  }
  if (value.length < 6) {
    return new Error('密码长度至少为6位');
  }
  return true;
};

const validateConfirmPassword = (rule: any, value: string) => {
  if (!value) {
    return new Error('请再次输入密码');
  }
  if (value !== registerData.password) {
    return new Error('两次输入的密码不一致');
  }
  return true;
};

const registerRules = {
  key: [
    {
      required: true,
      message: '请输入注册密钥',
      trigger: 'blur'
    }
  ],
  username: [
    {
      required: true,
      message: '请输入用户名',
      trigger: 'blur'
    }
  ],
  password: [
    {
      validator: validatePassword,
      trigger: 'blur'
    }
  ],
  confirmPassword: [
    {
      validator: validateConfirmPassword,
      trigger: 'blur'
    }
  ]
};

const handleLogin = async () => {
  if (!formRef.value) return;
  
  try {
    await formRef.value.validate();
    loading.value = true;
    
    const res = await loginApi(formData);
    if (res.code === 0) {
      userStore.setUserInfo(res.data);
      message.success('登录成功');
      router.push('/');
    } else {
      message.error(res.message || '登录失败');
    }
  } catch (error: any) {
    console.error('登录验证失败:', error);
    message.error(error.message || '登录失败，请重试');
  } finally {
    loading.value = false;
  }
};

const handleRegister = async () => {
  if (!registerFormRef.value) return;
  
  try {
    await registerFormRef.value.validate();
    registerLoading.value = true;
    
    const res = await registerApi({
      key: registerData.key,
      username: registerData.username,
      password: registerData.password
    });
    
    if (res.code === 0) {
      message.success('注册成功');
      // 注册成功后自动登录
      userStore.setUserInfo({
        token: res.data.token,
        expire: res.data.expire,
        userId: 0,
        username: registerData.username,
        ruleLevel: 0,
        lastLoginTme: 0,
        createTime: 0
      });
      showRegisterModal.value = false;
      router.push('/');
    } else {
      message.error(res.message || '注册失败');
    }
  } catch (error: any) {
    console.error('注册失败:', error);
    message.error(error.message || '注册失败，请重试');
  } finally {
    registerLoading.value = false;
  }
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