package _const

const (
	SuccNo       = 0
	ParamErrorNo = 1
	DBQueryError = 2
	DBExecError  = 3
	AuthError    = 4
	GPTErrorNo   = 5
)

const (
	SuccMsg         = "success"
	ParamErrorMsg   = "param error"
	DBQueryErrorMsg = "db query error"
	DBExecErrorMsg  = "db exec error"
	AuthErrorMsg    = "auth error"
	GPTErrorMsg     = "gpt error"
)

func GetErrorMsg(no int) string {
	switch no {
	case SuccNo:
		return SuccMsg
	case ParamErrorNo:
		return ParamErrorMsg
	case DBQueryError:
		return DBQueryErrorMsg
	case DBExecError:
		return DBExecErrorMsg
	case AuthError:
		return AuthErrorMsg
	}

	return SuccMsg
}
