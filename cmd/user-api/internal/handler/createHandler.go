package handler

import (
	"git.savvycom.vn/go-services/go-boilerplate/internal/presentation/pkg/httpresult"
	"net/http"

	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/logic"
	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/svc"
	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresult.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		httpresult.HTTPResult(r, w, resp, err)
	}
}
