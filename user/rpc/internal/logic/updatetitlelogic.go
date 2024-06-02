package logic

import (
	"context"
	"errors"

	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTitleLogic {
	return &UpdateTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTitleLogic) UpdateTitle(in *user.UpdateTitleReq) (*user.UpdateTitleResp, error) {
	// todo: add your logic here and delete this line
	instruct, err := l.svcCtx.InstructModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if instruct.Title != "" {
		return nil, errors.New("已设置开题，无法更改")
	}

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.InstructModel.Update(l.ctx, &model.InstructRelation{
		ID:    oid,
		Title: in.Title,
	})
	if err != nil {
		return nil, err
	}

	return &user.UpdateTitleResp{
		Updated: updateResult.ModifiedCount,
	}, nil
}
