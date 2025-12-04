import req, { type Response } from "@/utils/request";
import type { LoginReq, LoginRes, RegisterReq, RegisterRes } from "@/types/user";

export const loginApi = (data: LoginReq) => {
    return req<Response<LoginRes>>({
        url: '/api/v1/user/login',
        method: 'post',
        data
    })
}

export const registerApi = (data: RegisterReq) => {
    return req<Response<RegisterRes>>({
        url: '/api/v1/user/register',
        method: 'post',
        data
    })
}