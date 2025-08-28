// 聊天请求参数
export interface ChatRequest {
    ai: string;
    model: string;
    prompt?: string;
    message: string;
    databaseId: number;
}

// 数据库配置
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
