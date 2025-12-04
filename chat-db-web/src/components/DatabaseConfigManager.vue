<template>
  <div class="database-config-manager">
    <div class="header">
      <h3>数据库配置管理</h3>
      <el-button type="primary" @click="handleAdd">
        <i class="ri-add-line"></i>
        添加数据库
      </el-button>
    </div>

    <!-- 数据库列表 -->
    <el-table :data="databaseList" v-loading="loading" border style="width: 100%">
      <el-table-column prop="databaseId" label="ID" width="80" />
      <el-table-column prop="dbType" label="数据库类型" width="120">
        <template #default="{ row }">
          <el-tag :type="getDbTypeTagType(row.dbType)">
            {{ row.dbType?.toUpperCase() || '未知' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="dbName" label="数据库名称" min-width="120" />
      <el-table-column prop="host" label="主机地址" min-width="150">
        <template #default="{ row }">
          {{ row.host }}:{{ row.port }}
        </template>
      </el-table-column>
      <el-table-column prop="userName" label="用户名" min-width="100" />
      <el-table-column label="操作" width="180" fixed="right">
        <template #default="{ row }">
          <el-button link type="primary" size="small" @click="handleEdit(row)">
            <i class="ri-edit-line"></i>
            编辑
          </el-button>
          <el-button link type="danger" size="small" @click="handleDelete(row)">
            <i class="ri-delete-bin-line"></i>
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination" v-if="total > 0">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.size"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="loadDatabaseList"
        @current-change="loadDatabaseList"
      />
    </div>

    <!-- 添加/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="handleDialogClose"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
      >
        <el-form-item label="数据库类型" prop="dbType">
          <el-select
            v-model="formData.dbType"
            placeholder="请选择数据库类型"
            style="width: 100%"
          >
            <el-option
              v-for="type in dbTypes"
              :key="type.value"
              :label="type.label"
              :value="type.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="数据库名称" prop="dbName">
          <el-input
            v-model="formData.dbName"
            placeholder="请输入数据库名称"
          />
        </el-form-item>

        <el-form-item label="主机地址" prop="host">
          <el-input
            v-model="formData.host"
            placeholder="请输入主机地址，如: localhost 或 192.168.1.1"
          />
        </el-form-item>

        <el-form-item label="端口号" prop="port">
          <el-input-number
            v-model="formData.port"
            :min="1"
            :max="65535"
            placeholder="请输入端口号"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="用户名" prop="userName">
          <el-input
            v-model="formData.userName"
            placeholder="请输入用户名"
          />
        </el-form-item>

        <el-form-item 
          label="密码" 
          prop="password"
          :rules="isEdit ? [] : formRules.password"
        >
          <el-input
            v-model="formData.password"
            type="password"
            show-password
            :placeholder="isEdit ? '留空则不修改密码，如需更新请输入新密码' : '请输入密码'"
          />
          <template v-if="isEdit" #error>
            <span></span>
          </template>
        </el-form-item>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit" :loading="submitting">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue';

interface Emits {
  (e: 'refresh'): void;
}

const emit = defineEmits<Emits>();
import {
  ElTable,
  ElTableColumn,
  ElButton,
  ElTag,
  ElDialog,
  ElForm,
  ElFormItem,
  ElInput,
  ElInputNumber,
  ElSelect,
  ElOption,
  ElPagination,
  ElMessage,
  ElMessageBox,
  type FormInstance,
  type FormRules
} from 'element-plus';
import {
  getDatabaseConfigList,
  setDatabaseConfig,
  updateDatabaseConfig,
  deleteDatabaseConfig
} from '@/api/ai';
import type { DatabaseConfig, DatabaseConfigItem } from '@/types/ai';
import { useSettingsStore } from '@/stores/settings';

const settingsStore = useSettingsStore();

const loading = ref(false);
const submitting = ref(false);
const databaseList = ref<DatabaseConfigItem[]>([]);
const total = ref(0);
const dialogVisible = ref(false);
const isEdit = ref(false);
const formRef = ref<FormInstance>();

const pagination = reactive({
  page: 1,
  size: 10
});

const dbTypes = [
  { label: 'MySQL', value: 'mysql' },
  { label: 'PostgreSQL', value: 'postgres' },
  { label: 'SQLite', value: 'sqlite' }
];

const formData = reactive<DatabaseConfig>({
  databaseId: undefined,
  dbName: '',
  userName: '',
  password: '',
  host: '',
  port: 3306,
  dbType: 'mysql'
});

const formRules: FormRules = {
  dbType: [{ required: true, message: '请选择数据库类型', trigger: 'change' }],
  dbName: [{ required: true, message: '请输入数据库名称', trigger: 'blur' }],
  host: [{ required: true, message: '请输入主机地址', trigger: 'blur' }],
  port: [{ required: true, message: '请输入端口号', trigger: 'blur' }],
  userName: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
};

const dialogTitle = computed(() => {
  return isEdit.value ? '编辑数据库配置' : '添加数据库配置';
});

const getDbTypeTagType = (dbType: string) => {
  const typeMap: Record<string, string> = {
    mysql: 'primary',
    postgres: 'success',
    sqlite: 'info'
  };
  return typeMap[dbType?.toLowerCase()] || '';
};

// 加载数据库列表
const loadDatabaseList = async () => {
  loading.value = true;
  try {
    const res = await getDatabaseConfigList({
      page: pagination.page,
      size: pagination.size
    });

    if (res && res.code === 0 && res.data) {
      databaseList.value = res.data.list || [];
      total.value = res.data.total || 0;
    } else {
      ElMessage.error(res?.message || '加载数据库配置列表失败');
    }
  } catch (error: any) {
    console.error('加载数据库配置列表失败:', error);
    ElMessage.error(error.message || '加载数据库配置列表失败');
  } finally {
    loading.value = false;
  }
};

// 刷新数据库选项（用于更新设置面板的下拉列表）
const refreshDatabaseOptions = async () => {
  try {
    const res = await getDatabaseConfigList({ page: 1, size: 100 });
    if (res && res.code === 0 && res.data) {
      const list = res.data.list || [];
      if (list.length > 0) {
        const options = list.map((db: DatabaseConfigItem) => ({
          label: `${db.dbName || '未命名'} (${db.dbType || 'unknown'})`,
          value: db.databaseId || 0
        }));
        settingsStore.updateDatabaseOptions(options);
      } else {
        settingsStore.updateDatabaseOptions([]);
      }
    }
  } catch (error) {
    console.error('刷新数据库选项失败:', error);
  }
};

// 添加
const handleAdd = () => {
  isEdit.value = false;
  resetForm();
  dialogVisible.value = true;
};

// 编辑
const handleEdit = (row: DatabaseConfigItem) => {
  isEdit.value = true;
  Object.assign(formData, {
    databaseId: row.databaseId,
    dbName: row.dbName,
    userName: row.userName,
    password: '', // 列表响应不包含密码，编辑时需要重新输入
    host: row.host,
    port: row.port,
    dbType: row.dbType
  });
  dialogVisible.value = true;
};

// 删除
const handleDelete = async (row: DatabaseConfigItem) => {
  if (!row.databaseId) {
    ElMessage.error('数据库配置ID不存在');
    return;
  }

  try {
    await ElMessageBox.confirm(
      `确定要删除数据库配置 "${row.dbName}" 吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    );

    const res = await deleteDatabaseConfig(row.databaseId);
        if (res && res.code === 0) {
          ElMessage.success('删除成功');
          await loadDatabaseList();
          await refreshDatabaseOptions();
          emit('refresh'); // 通知父组件刷新
        } else {
      ElMessage.error(res?.message || '删除失败');
    }
  } catch (error: any) {
    if (error !== 'cancel') {
      console.error('删除数据库配置失败:', error);
      ElMessage.error(error.message || '删除失败');
    }
  }
};

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return;

  await formRef.value.validate(async (valid) => {
    if (!valid) return;

    submitting.value = true;
    try {
        if (isEdit.value) {
        // 更新
        if (!formData.databaseId) {
          ElMessage.error('数据库配置ID不存在');
          return;
        }

        // 构建更新数据，如果密码为空则不包含密码字段（后端需要支持不传密码时不更新）
        const updateData: DatabaseConfig & { databaseId: number } = {
          databaseId: formData.databaseId,
          dbName: formData.dbName,
          userName: formData.userName,
          host: formData.host,
          port: formData.port,
          dbType: formData.dbType,
          password: formData.password || '' // 如果为空字符串，后端将不更新密码
        };

        const res = await updateDatabaseConfig(updateData);

        if (res && res.code === 0) {
          ElMessage.success('更新成功');
          dialogVisible.value = false;
          await loadDatabaseList();
          await refreshDatabaseOptions();
          emit('refresh'); // 通知父组件刷新
        } else {
          ElMessage.error(res?.message || '更新失败');
        }
      } else {
        // 新增
        const res = await setDatabaseConfig({
          dbName: formData.dbName,
          userName: formData.userName,
          password: formData.password,
          host: formData.host,
          port: formData.port,
          dbType: formData.dbType
        });

        if (res && res.code === 0) {
          ElMessage.success('添加成功');
          dialogVisible.value = false;
          await loadDatabaseList();
          await refreshDatabaseOptions();
          emit('refresh'); // 通知父组件刷新
        } else {
          ElMessage.error(res?.message || '添加失败');
        }
      }
    } catch (error: any) {
      console.error('操作失败:', error);
      ElMessage.error(error.message || '操作失败');
    } finally {
      submitting.value = false;
    }
  });
};

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    databaseId: undefined,
    dbName: '',
    userName: '',
    password: '',
    host: '',
    port: 3306,
    dbType: 'mysql'
  });
  formRef.value?.clearValidate();
};

// 对话框关闭
const handleDialogClose = () => {
  resetForm();
};

// 初始化
onMounted(() => {
  loadDatabaseList();
});

// 暴露方法供父组件调用
defineExpose({
  loadDatabaseList,
  refreshDatabaseOptions
});
</script>

<style scoped lang="scss">
.database-config-manager {
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 16px;

    h3 {
      margin: 0;
      font-size: 16px;
      font-weight: 600;
    }
  }

  .pagination {
    margin-top: 16px;
    display: flex;
    justify-content: flex-end;
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
