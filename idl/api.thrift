namespace go api

struct LoginRequest {
    1: string username
    2: string password
}

struct RegisterRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: string token
}

struct FileUploadRequest {
}

struct FileDownloadRequest {
    1: string filename
}

struct CommonResponse {
    1: string message
}

struct GetDirRequest {
    1: string filename
}

struct GetDirResponse {
    1: list<string> dir
}

struct RemoveRepeatRequest {
    1: string filename
}

struct RemoveRepeatResponse {
    1: string path
}

struct SortByNumRequest {
    1: string filename
}

struct SortByNumResponse {
    1: string path
}

service ApiService {
    LoginResponse Login(1: LoginRequest req)(api.post="/user/login")
    CommonResponse Regiter(1: RegisterRequest req)(api.post="/user/register")

    CommonResponse UploadFile(1: FileUploadRequest req)(api.post="/file/upload")
    CommonResponse DownloadFile(1: FileDownloadRequest req)(api.get="/file/download")
    GetDirResponse GetDir(1: GetDirRequest req)(api.get="/file/gitdir")
    RemoveRepeatResponse RemoveRepeat(1: RemoveRepeatRequest req)(api.post="/file/rmrepeat")
    SortByNumResponse SortByNum(1: SortByNumRequest req)(api.post="/file/sort")
}
