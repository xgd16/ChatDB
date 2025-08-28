import req, { type Response } from '@/utils/request';
import type { ChatRequest, DatabaseConfig, AIModelItem } from '@/types/ai';
import type { PaginatedResponse, PaginationParams } from '@/types/api';
import { useUserStore } from '@/stores/counter';

// 聊天API - 返回流式响应
export const chatApi = (data: ChatRequest): Promise<globalThis.Response> => {
    const userStore = useUserStore();
    return fetch('/api/v1/chats', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${userStore.token}`
        },
        body: JSON.stringify(data)
    });
};

// 获取数据库配置列表
export const getDatabaseConfigList = (params: PaginationParams = {}) => {
    return req<Response<PaginatedResponse<DatabaseConfig>>>({
        url: '/api/v1/config/db/listConfig',
        method: 'GET',
        params
    });
};

// 设置数据库配置
export const setDatabaseConfig = (data: DatabaseConfig) => {
    return req<Response<any>>({
        url: '/api/v1/config/db/setConfig',
        method: 'POST',
        data
    });
};

// 更新数据库配置
export const updateDatabaseConfig = (data: DatabaseConfig & { databaseId: number }) => {
    return req<Response<any>>({
        url: '/api/v1/config/db/updateConfig',
        method: 'PUT',
        data
    });
};

// 删除数据库配置
export const deleteDatabaseConfig = (databaseId: number) => {
    return req<Response<any>>({
        url: '/api/v1/config/db/deleteConfig',
        method: 'DELETE',
        params: { databaseId }
    });
};

// 获取AI模型列表
export const getAiModelList = () => {
    return req<Response<PaginatedResponse<AIModelItem>>>({
        url: '/api/v1/config/ai/model/list',
        method: 'GET'
    });
};
