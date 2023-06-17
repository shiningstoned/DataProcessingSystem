namespace go file

struct GetFilesRequest {
    1: string username
}

struct GetFilesResponse {
    1: list<string> files
}

struct  RemoveRepeatRequest {
    1: string filename
}

struct RemoveRepeatResponse {
    1: string filename
}

struct SortByNumRequest {
    1: string filename
}

struct SortByNumResponse {
    1: string filename
}

struct SortByTimeRequest {
    1: string filename
}

struct SortByTimeResponse {
    1: string filename
}

service FileService {
    GetFilesResponse GetFiles(1: GetFilesRequest req)
    RemoveRepeatResponse RemoveRepeat(1: RemoveRepeatRequest req)
    SortByNumResponse SortByNum(1: SortByNumRequest req)
    SortByTimeResponse SortByTime(1: SortByNumRequest req)
}