// 通用API响应格式
export interface ApiResponse<T = any> {
    code: number;
    message: string;
    data: T;
    serviceTime?: number;
}

// 分页响应格式
export interface PaginatedResponse<T = any> {
    list: T[];
    total: number;
}

// 分页请求参数
export interface PaginationParams {
    page?: number;
    size?: number;
}
