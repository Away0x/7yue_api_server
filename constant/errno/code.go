package errno

var (
	// 100x 通用类型
	OK                  = &Errno{Code: 0, Message: "成功"}
	ParamsError         = &Errno{Code: 1000, Message: "输入参数错误"}
	JsonError           = &Errno{Code: 1001, Message: "输入的 json 格式不正确"}
	NoResourceError     = &Errno{Code: 1002, Message: "找不到资源"}
	UnknownError        = &Errno{Code: 1003, Message: "未知错误"}
	AuthError           = &Errno{Code: 1004, Message: "禁止访问"}
	DevKeyError         = &Errno{Code: 1005, Message: "不正确的开发者 key"}
	InternalServerError = &Errno{Code: 1006, Message: "服务器内部错误"}
	ErrDatabase         = &Errno{Code: 1007, Message: "数据库错误"}
	RouterError         = &Errno{Code: 1008, Message: "路由错误"}

	// 200x 点赞类型
	FavoredError        = &Errno{Code: 2000, Message: "你已经点过赞了"}
	NoFavorError        = &Errno{Code: 2001, Message: "你还没点过赞"}

	// 300x 期刊类型
	NoClassicError     = &Errno{Code: 3000, Message: "该期内容不存在"}

	// 400x user
	UserCreateError    = &Errno{Code: 4000, Message: "用户创建失败"}
)