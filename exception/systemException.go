package exception

const (
	SystemError = 1
	ServerMaintenance = 2
	ParamError = 3
)

type Exception struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SystemException(msg ...string) *Exception {
	exception := Exception{}
	exception.Code = SystemError
	exception.Msg = "System error"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return &exception
}

func ServerMaintainException(msg ...string) *Exception {
	exception := Exception{}
	exception.Code = ServerMaintenance
	exception.Msg = "Server maintenance"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return &exception
}

func ParamException(msg ...string) *Exception {
	exception := Exception{}
	exception.Code = ParamError
	exception.Msg = "Invaldated parameter"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return &exception
}
