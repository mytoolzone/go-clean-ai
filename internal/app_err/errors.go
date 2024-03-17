package app_err

import (
	"fmt"
	"runtime"
)

const (
	// 通用错误类型
	ErrInternalError = "INTERNAL_ERROR" // 内部错误
	ErrNotFound      = "NOT_FOUND"      // 找不到资源
	ErrUnauthorized  = "UNAUTHORIZED"   // 未授权
	ErrForbidden     = "FORBIDDEN"      // 禁止访问
	ErrBadRequest    = "BAD_REQUEST"    // 错误的请求
	ErrConflict      = "CONFLICT"       // 冲突

	// 数据库相关错误类型
	ErrDBConnectionFailed = "DB_CONNECTION_FAILED" // 数据库连接失败
	ErrDBQueryFailed      = "DB_QUERY_FAILED"      // 数据库查询失败
	ErrDBInsertFailed     = "DB_INSERT_FAILED"     // 数据库插入失败
	ErrDBUpdateFailed     = "DB_UPDATE_FAILED"     // 数据库更新失败
	ErrDBDeleteFailed     = "DB_DELETE_FAILED"     // 数据库删除失败

	// 认证和授权相关错误类型
	ErrInvalidCredentials = "INVALID_CREDENTIALS" // 无效的凭证
	ErrInvalidToken       = "INVALID_TOKEN"       // 无效的凭证
	ErrExpiredToken       = "EXPIRED_TOKEN"       // 令牌已过期
	ErrPermissionDenied   = "PERMISSION_DENIED"   // 权限被拒绝

	// 文件操作相关错误类型
	ErrFileNotFound    = "FILE_NOT_FOUND"    // 文件不存在
	ErrFileReadFailed  = "FILE_READ_FAILED"  // 文件读取失败
	ErrFileWriteFailed = "FILE_WRITE_FAILED" // 文件写入失败

)

const (
	MsgInitAppError     = "初始化程序异常"
	MsgDbQueryErr       = "查找数据异常"
	MsgDbUpdateErr      = "更新数据异常"
	MsgPermissionDenied = "非法访问"
	MsgExpiredToken     = "访问令牌过期"
)

// AppError 是应用程序中的自定义错误类型
type AppError struct {
	Code    string // 错误代码
	Message string // 错误消息
	File    string // 出错文件名
	Line    int    // 出错行号
	Err     error  // 原始错误（可选）
}

// NewError 创建一个新的 AppError 实例
func NewError(code string, message string, err ...error) *AppError {
	_, file, line, _ := runtime.Caller(1)
	e := &AppError{
		Code:    code,
		Message: message,
		File:    file,
		Line:    line,
	}
	if len(err) > 0 {
		e.Err = err[0]
	}
	return e
}

// Error 实现了 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s:%d] %s (Code: %s) - %s", e.File, e.Line, e.Message, e.Code, e.Err.Error())
	}
	return fmt.Sprintf("[%s:%d] %s (Code: %s)", e.File, e.Line, e.Message, e.Code)
}

// Unwrap 实现了 Wrapper 接口
func (e *AppError) Unwrap() error {
	return e.Err
}
