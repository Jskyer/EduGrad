package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateInstructLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInstructLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInstructLogic {
	return &UpdateInstructLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInstructLogic) UpdateInstruct(in *user.UpdateInstructReq) (*user.UpdateInstructResp, error) {
	// todo: add your logic here and delete this line
	instruct, err := l.svcCtx.InstructModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// permit合法性
	if in.Permit != "1" && in.Permit != "2" {
		return nil, errors.New("permit应设置1或2")
	}

	if instruct.Permit != "0" {
		return nil, errors.New("申请已设置permit")
	}

	// 设置通过需检验该学生是否已有导师,同时该导师已有的instruct不能>= 6
	if in.Permit == "1" {
		permited, err := l.svcCtx.InstructModel.FindPermitedByStuId(l.ctx, instruct.StuID)
		if permited != nil {
			// 删除
			delNum, err := l.svcCtx.InstructModel.Delete(l.ctx, instruct.ID.Hex())
			if err != nil {
				return nil, err
			}

			logx.Info(fmt.Sprintf("delNum: %d", delNum))

			return nil, errors.New("该学生已有导师，申请失效")
		}

		if err != mon.ErrNotFound {
			return nil, err
		}

		cnt, err := l.svcCtx.InstructModel.CountPermitedByTeacherId(l.ctx, instruct.TeacherID)
		if err != nil {
			return nil, err
		}

		if cnt >= 6 {
			// 删除
			// delNum, err := l.svcCtx.InstructModel.Delete(l.ctx, instruct.ID.Hex())
			// if err != nil {
			// 	return nil, err
			// }

			// logx.Info(fmt.Sprintf("delNum: %d", delNum))

			return nil, errors.New("导师的指导学生已满")
		}
	}

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	newInstruct := model.InstructRelation{
		ID:     oid,
		Permit: in.Permit,
	}

	if in.Permit == "1" {
		newInstruct.State = "1"
	}

	// 更新permit为1或2
	updateResult, err := l.svcCtx.InstructModel.Update(l.ctx, &newInstruct)
	if err != nil {
		return nil, err
	}

	return &user.UpdateInstructResp{
		Updated: updateResult.ModifiedCount,
	}, nil
}
