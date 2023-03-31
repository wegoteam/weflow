namespace go api

struct ApiReq {
    1: string Name (api.query="api"); // 添加 api 注解为方便进行参数绑定
}

struct ApiResp {
    1: string RespBody;
}

service ApiService {
    ApiResp HelloMethod(1: ApiReq request) (api.get="/api");
}