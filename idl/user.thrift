namespace go user

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: string username
}

struct RegisterRequest {
    1: string username
    2: string password
}

struct CommonResponse {
    1: string message
}

service UserService {
    CommonResponse Register(1: RegisterRequest req)
    LoginResponse Login(1: LoginRequest req)
}