package errs

import "errors"

var (
	ERR_NONE            = errors.New("没有错误")
	ERR_UNKNOWN         = errors.New("未知错误")
	ERR_NOT_IMPLEMENTED = errors.New("暂未实现")
	ERR_OK              = errors.New("执行成功")

	ERR_BAD_REQUEST              = errors.New("错误的请求")
	ERR_NOT_FOUND                = errors.New("内容未找到")
	ERR_DUPLICATED               = errors.New("内容已经存在")
	ERR_INVALID_PARAMETERS       = errors.New("无效的参数")
	ERR_INVALID_VALUES           = errors.New("数据库操作失败")
	ERR_DATABASE_OPERATION       = errors.New("数据库操作失败")
	ERR_DATABASE_NOT_INITIALIZED = errors.New("数据库未初始化")
	ERR_SERVICE_UNAVAILABLE      = errors.New("服务不可用")
	ERR_UNAUTHORIZED             = errors.New("访问未授权")
)
