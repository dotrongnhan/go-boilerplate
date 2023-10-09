package httpresult

import (
	"fmt"
	"net/http"

	"git.savvycom.vn/go-services/go-boilerplate/internal/presentation/pkg/httperr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func HTTPResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// return successfully
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// return Error
		errCode := httperr.ServerCommonError
		errMsg := err.Error()

		causeErr := errors.Cause(err)                   // get error cause
		if e, ok := causeErr.(*httperr.CodeError); ok { // if it is custom error type
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err
				grpcCode := uint32(gStatus.Code())
				if httperr.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMsg = gStatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}

func AuthHTTPResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errCode := httperr.ServerCommonError
		errMsg := "Server is out of service, try again later"

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*httperr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				grpcCode := uint32(gstatus.Code())
				if httperr.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errCode, errMsg))
	}
}

func ParamErrorResult(_ *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", httperr.MapErrMsg(httperr.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(httperr.RequestParamError, errMsg))
}
