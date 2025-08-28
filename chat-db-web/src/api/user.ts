import req, { type Response } from "@/utils/request";
import type { LoginReq, LoginRes } from "@/types/user";

export const loginApi = (data: LoginReq) => {
    return req<Response<LoginRes>>({
        url: '/api/v1/user/login',
        method: 'post',
        data
    })
}