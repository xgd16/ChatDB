

export interface LoginReq {
    username: string;
    password: string;
}


export interface LoginRes {
    token:        string;
    expire:       number;
    userId:       number;
    username:     string;
    ruleLevel:    number;
    lastLoginTme: number;
    createTime:   number;
}
