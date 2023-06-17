package global

import (
	file "hdfs/kitex_gen/file/fileservice"
	user "hdfs/kitex_gen/user/userservice"
)

var (
	UserClient user.Client
	FileClient file.Client
)
