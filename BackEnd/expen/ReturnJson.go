package expen

const (
	MissingParameters = -1 // 缺少参数
	InternalError     = -2 // 内部错误
	ParameterError    = -3 // 参数错误
	UnknownError      = -4 // 未知错误
	NotOwned          = -5 // 参数在内部未拥有或者不存在
)

// Return 返回值结果
type Return struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(data interface{}, msg string) Return {
	return Return{
		Code: 200,
		Data: data,
		Msg:  msg,
	}
}

func MissingParametersFun(msg string) Return {
	return Return{
		Code: MissingParameters,
		Data: nil,
		Msg:  msg,
	}
}

func InternalErrorFun(msg string) Return {
	return Return{
		Code: InternalError,
		Data: nil,
		Msg:  msg,
	}
}

func ParameterErrorFun(msg string) Return {
	return Return{
		Code: ParameterError,
		Data: nil,
		Msg:  msg,
	}
}

func UnknownErrorFun(msg string) Return {
	return Return{
		Code: UnknownError,
		Data: nil,
		Msg:  msg,
	}
}

func NotOwnedFun(msg string) Return {
	return Return{
		Code: NotOwned,
		Data: nil,
		Msg:  msg,
	}
}
