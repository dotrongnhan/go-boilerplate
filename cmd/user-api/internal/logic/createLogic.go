package logic

import (
	"context"
	"git.savvycom.vn/go-services/go-boilerplate/internal/app/models"

	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/svc"
	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	createUserOutput, err := l.svcCtx.UserUseCase.CreateUser(l.ctx, &models.CreateUserInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateUserResp{Id: createUserOutput.UserID}, nil
}
