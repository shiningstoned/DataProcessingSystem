// Code generated by hertz generator.

package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"hdfs/api/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Cors()}
}

func _fileMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _downloadfileMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _uploadfileMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _regiterMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getdirMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _removerepeatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _sortbynumMw() []app.HandlerFunc {
	// your code...
	return nil
}
