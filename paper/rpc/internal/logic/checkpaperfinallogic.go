package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CheckPaperFinalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPaperFinalLogic {
	return &CheckPaperFinalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckPaperFinalLogic) CheckPaperFinal(in *paper.CheckPaperFinalReq) (*paper.CheckPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.PaperFinalModel.Update(l.ctx, &model.PaperFinal{
		ID:          oid,
		CheckResult: in.CheckResult,
	})
	if err != nil {
		return nil, err
	}

	return &paper.CheckPaperFinalResp{
		Ok: updateResult.ModifiedCount,
	}, nil
}
