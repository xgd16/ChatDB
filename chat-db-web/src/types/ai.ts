// 聊天请求参数
export interface ChatRequest {
    ai: string;
    model: string;
    prompt?: string;
    message: string;
    databaseId: number;
}

// 数据库配置项（列表响应，不包含密码）
export interface DatabaseConfigItem {
    databaseId: number;
    dbName: string;
    userName: string;
    host: string;
    port: number;
    dbType: string;
    createTime?: number;
    updateTime?: number;
}

// 数据库配置（创建/更新请求，包含密码）
export interface DatabaseConfig {
    databaseId?: number;
    dbName: string;
    userName: string;
    password: string;
    host: string;
    port: number;
    dbType: string;
}

// AI模型项
export interface AIModelItem {
    id: string;
    object: string;
    created: number;
    owned_by: string;
    permission: any[];
    root: string;
    parent: any;
}

// 聊天消息
export interface ChatMessage {
    role: 'user' | 'assistant';
    content: string;
    timestamp: number;
}

// 聊天流式响应事件
export interface ChatStreamEvent {
    event: 'start' | 'message' | 'end' | 'tool_call' | 'ping';
    traceId?: string;
    role?: string;
    content?: string;
    data?: any;
    createTime?: number;
}
