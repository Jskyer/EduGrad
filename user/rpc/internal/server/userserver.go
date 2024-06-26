// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"edu-grad/user/rpc/internal/logic"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterReq) (*user.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) UpdateUserInfo(ctx context.Context, in *user.UpdateUserInfoReq) (*user.UpdateUserInfoResp, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UserServer) GenerateToken(ctx context.Context, in *user.GenerateTokenReq) (*user.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UserServer) GetUserPage(ctx context.Context, in *user.GetUserPageReq) (*user.GetUserPageResp, error) {
	l := logic.NewGetUserPageLogic(ctx, s.svcCtx)
	return l.GetUserPage(in)
}

func (s *UserServer) GetUserCondPage(ctx context.Context, in *user.GetUserCondPageReq) (*user.GetUserCondPageResp, error) {
	l := logic.NewGetUserCondPageLogic(ctx, s.svcCtx)
	return l.GetUserCondPage(in)
}

func (s *UserServer) AddInstruct(ctx context.Context, in *user.AddInstructReq) (*user.AddInstructResp, error) {
	l := logic.NewAddInstructLogic(ctx, s.svcCtx)
	return l.AddInstruct(in)
}

func (s *UserServer) UpdateInstruct(ctx context.Context, in *user.UpdateInstructReq) (*user.UpdateInstructResp, error) {
	l := logic.NewUpdateInstructLogic(ctx, s.svcCtx)
	return l.UpdateInstruct(in)
}

func (s *UserServer) GetInstructByStuId(ctx context.Context, in *user.GetInstructByStuIdReq) (*user.GetInstructByStuIdResp, error) {
	l := logic.NewGetInstructByStuIdLogic(ctx, s.svcCtx)
	return l.GetInstructByStuId(in)
}

func (s *UserServer) GetListByTeacherId(ctx context.Context, in *user.GetListByTeacherIdReq) (*user.GetListByTeacherIdResp, error) {
	l := logic.NewGetListByTeacherIdLogic(ctx, s.svcCtx)
	return l.GetListByTeacherId(in)
}

func (s *UserServer) UpdateTitle(ctx context.Context, in *user.UpdateTitleReq) (*user.UpdateTitleResp, error) {
	l := logic.NewUpdateTitleLogic(ctx, s.svcCtx)
	return l.UpdateTitle(in)
}

func (s *UserServer) UpdateState(ctx context.Context, in *user.UpdateStateReq) (*user.UpdateStateResp, error) {
	l := logic.NewUpdateStateLogic(ctx, s.svcCtx)
	return l.UpdateState(in)
}

func (s *UserServer) AddPhase(ctx context.Context, in *user.AddPhaseReq) (*user.AddPhaseResp, error) {
	l := logic.NewAddPhaseLogic(ctx, s.svcCtx)
	return l.AddPhase(in)
}

func (s *UserServer) UpdateProcess(ctx context.Context, in *user.UpdateProcessReq) (*user.UpdateProcessResp, error) {
	l := logic.NewUpdateProcessLogic(ctx, s.svcCtx)
	return l.UpdateProcess(in)
}

func (s *UserServer) ListTermSort(ctx context.Context, in *user.ListTermSortReq) (*user.ListTermSortResp, error) {
	l := logic.NewListTermSortLogic(ctx, s.svcCtx)
	return l.ListTermSort(in)
}

func (s *UserServer) GetPhaseinfo(ctx context.Context, in *user.PhaseInfoReq) (*user.PhaseInfoResp, error) {
	l := logic.NewGetPhaseinfoLogic(ctx, s.svcCtx)
	return l.GetPhaseinfo(in)
}

func (s *UserServer) AddPhaseRelation(ctx context.Context, in *user.AddPhaseRelationReq) (*user.AddPhaseRelationResp, error) {
	l := logic.NewAddPhaseRelationLogic(ctx, s.svcCtx)
	return l.AddPhaseRelation(in)
}

func (s *UserServer) EndPhaseRelation(ctx context.Context, in *user.EndPhaseRelationReq) (*user.EndPhaseRelationResp, error) {
	l := logic.NewEndPhaseRelationLogic(ctx, s.svcCtx)
	return l.EndPhaseRelation(in)
}

func (s *UserServer) ListPhaseRelation(ctx context.Context, in *user.ListPhaseRelationReq) (*user.ListPhaseRelationResp, error) {
	l := logic.NewListPhaseRelationLogic(ctx, s.svcCtx)
	return l.ListPhaseRelation(in)
}

func (s *UserServer) StulistForTeacher(ctx context.Context, in *user.StulistForTeacherReq) (*user.StulistForTeacherResp, error) {
	l := logic.NewStulistForTeacherLogic(ctx, s.svcCtx)
	return l.StulistForTeacher(in)
}

func (s *UserServer) PhaseAllowedtist(ctx context.Context, in *user.PhaseAllowedtistReq) (*user.PhaseAllowedtistResp, error) {
	l := logic.NewPhaseAllowedtistLogic(ctx, s.svcCtx)
	return l.PhaseAllowedtist(in)
}

func (s *UserServer) PhaseCommittedtist(ctx context.Context, in *user.PhaseCommittedtistReq) (*user.PhaseCommittedtistResp, error) {
	l := logic.NewPhaseCommittedtistLogic(ctx, s.svcCtx)
	return l.PhaseCommittedtist(in)
}

func (s *UserServer) InphaseRelation(ctx context.Context, in *user.InphaseRelationReq) (*user.InphaseRelationResp, error) {
	l := logic.NewInphaseRelationLogic(ctx, s.svcCtx)
	return l.InphaseRelation(in)
}
