package httperr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "Server is out of service, try again later"
	message[RequestParamError] = "Bad request parameters"
	message[TokenExpireError] = "The token is invalid, please log in again"
	message[TokenGenerateError] = "Failed to generate token"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "Server is out of service, try again later"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
