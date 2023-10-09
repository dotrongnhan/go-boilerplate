package httperr

import (
	"github.com/cockroachdb/errors"
)

var ErrorNotFoundData = errors.New("sql: no rows in result set")

const OK uint32 = 200

// global error code

const (
	ServerCommonError         uint32 = 100001
	RequestParamError         uint32 = 100002
	TokenExpireError          uint32 = 100003
	TokenGenerateError        uint32 = 100004
	DBError                   uint32 = 100005
	DBUpdateAffectedZeroError uint32 = 100006
)

// specific module error code (e.g User module)

// user module
