// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddInstructReq         = user.AddInstructReq
	AddInstructResp        = user.AddInstructResp
	AddPhaseRelationReq    = user.AddPhaseRelationReq
	AddPhaseRelationResp   = user.AddPhaseRelationResp
	AddPhaseReq            = user.AddPhaseReq
	AddPhaseResp           = user.AddPhaseResp
	EndPhaseRelationReq    = user.EndPhaseRelationReq
	EndPhaseRelationResp   = user.EndPhaseRelationResp
	GenerateTokenReq       = user.GenerateTokenReq
	GenerateTokenResp      = user.GenerateTokenResp
	GetInstructByStuIdReq  = user.GetInstructByStuIdReq
	GetInstructByStuIdResp = user.GetInstructByStuIdResp
	GetListByTeacherIdReq  = user.GetListByTeacherIdReq
	GetListByTeacherIdResp = user.GetListByTeacherIdResp
	GetUserCondPageReq     = user.GetUserCondPageReq
	GetUserCondPageResp    = user.GetUserCondPageResp
	GetUserInfoReq         = user.GetUserInfoReq
	GetUserInfoResp        = user.GetUserInfoResp
	GetUserPageReq         = user.GetUserPageReq
	GetUserPageResp        = user.GetUserPageResp
	InstructRelation       = user.InstructRelation
	ListPhaseRelationReq   = user.ListPhaseRelationReq
	ListPhaseRelationResp  = user.ListPhaseRelationResp
	ListTermSortReq        = user.ListTermSortReq
	ListTermSortResp       = user.ListTermSortResp
	LoginReq               = user.LoginReq
	LoginResp              = user.LoginResp
	Phase                  = user.Phase
	PhaseRelation          = user.PhaseRelation
	RegisterReq            = user.RegisterReq
	RegisterResp           = user.RegisterResp
	UpdateInstructReq      = user.UpdateInstructReq
	UpdateInstructResp     = user.UpdateInstructResp
	UpdateProcessReq       = user.UpdateProcessReq
	UpdateProcessResp      = user.UpdateProcessResp
	UpdateStateReq         = user.UpdateStateReq
	UpdateStateResp        = user.UpdateStateResp
	UpdateTitleReq         = user.UpdateTitleReq
	UpdateTitleResp        = user.UpdateTitleResp
	UpdateUserInfoReq      = user.UpdateUserInfoReq
	UpdateUserInfoResp     = user.UpdateUserInfoResp
	UserInfo               = user.UserInfo

	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		GetUserPage(ctx context.Context, in *GetUserPageReq, opts ...grpc.CallOption) (*GetUserPageResp, error)
		GetUserCondPage(ctx context.Context, in *GetUserCondPageReq, opts ...grpc.CallOption) (*GetUserCondPageResp, error)
		AddInstruct(ctx context.Context, in *AddInstructReq, opts ...grpc.CallOption) (*AddInstructResp, error)
		UpdateInstruct(ctx context.Context, in *UpdateInstructReq, opts ...grpc.CallOption) (*UpdateInstructResp, error)
		GetInstructByStuId(ctx context.Context, in *GetInstructByStuIdReq, opts ...grpc.CallOption) (*GetInstructByStuIdResp, error)
		GetListByTeacherId(ctx context.Context, in *GetListByTeacherIdReq, opts ...grpc.CallOption) (*GetListByTeacherIdResp, error)
		UpdateTitle(ctx context.Context, in *UpdateTitleReq, opts ...grpc.CallOption) (*UpdateTitleResp, error)
		UpdateState(ctx context.Context, in *UpdateStateReq, opts ...grpc.CallOption) (*UpdateStateResp, error)
		AddPhase(ctx context.Context, in *AddPhaseReq, opts ...grpc.CallOption) (*AddPhaseResp, error)
		UpdateProcess(ctx context.Context, in *UpdateProcessReq, opts ...grpc.CallOption) (*UpdateProcessResp, error)
		ListTermSort(ctx context.Context, in *ListTermSortReq, opts ...grpc.CallOption) (*ListTermSortResp, error)
		AddPhaseRelation(ctx context.Context, in *AddPhaseRelationReq, opts ...grpc.CallOption) (*AddPhaseRelationResp, error)
		EndPhaseRelation(ctx context.Context, in *EndPhaseRelationReq, opts ...grpc.CallOption) (*EndPhaseRelationResp, error)
		ListPhaseRelation(ctx context.Context, in *ListPhaseRelationReq, opts ...grpc.CallOption) (*ListPhaseRelationResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUser) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUser) GetUserPage(ctx context.Context, in *GetUserPageReq, opts ...grpc.CallOption) (*GetUserPageResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserPage(ctx, in, opts...)
}

func (m *defaultUser) GetUserCondPage(ctx context.Context, in *GetUserCondPageReq, opts ...grpc.CallOption) (*GetUserCondPageResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserCondPage(ctx, in, opts...)
}

func (m *defaultUser) AddInstruct(ctx context.Context, in *AddInstructReq, opts ...grpc.CallOption) (*AddInstructResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddInstruct(ctx, in, opts...)
}

func (m *defaultUser) UpdateInstruct(ctx context.Context, in *UpdateInstructReq, opts ...grpc.CallOption) (*UpdateInstructResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateInstruct(ctx, in, opts...)
}

func (m *defaultUser) GetInstructByStuId(ctx context.Context, in *GetInstructByStuIdReq, opts ...grpc.CallOption) (*GetInstructByStuIdResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetInstructByStuId(ctx, in, opts...)
}

func (m *defaultUser) GetListByTeacherId(ctx context.Context, in *GetListByTeacherIdReq, opts ...grpc.CallOption) (*GetListByTeacherIdResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetListByTeacherId(ctx, in, opts...)
}

func (m *defaultUser) UpdateTitle(ctx context.Context, in *UpdateTitleReq, opts ...grpc.CallOption) (*UpdateTitleResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateTitle(ctx, in, opts...)
}

func (m *defaultUser) UpdateState(ctx context.Context, in *UpdateStateReq, opts ...grpc.CallOption) (*UpdateStateResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateState(ctx, in, opts...)
}

func (m *defaultUser) AddPhase(ctx context.Context, in *AddPhaseReq, opts ...grpc.CallOption) (*AddPhaseResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddPhase(ctx, in, opts...)
}

func (m *defaultUser) UpdateProcess(ctx context.Context, in *UpdateProcessReq, opts ...grpc.CallOption) (*UpdateProcessResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateProcess(ctx, in, opts...)
}

func (m *defaultUser) ListTermSort(ctx context.Context, in *ListTermSortReq, opts ...grpc.CallOption) (*ListTermSortResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.ListTermSort(ctx, in, opts...)
}

func (m *defaultUser) AddPhaseRelation(ctx context.Context, in *AddPhaseRelationReq, opts ...grpc.CallOption) (*AddPhaseRelationResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddPhaseRelation(ctx, in, opts...)
}

func (m *defaultUser) EndPhaseRelation(ctx context.Context, in *EndPhaseRelationReq, opts ...grpc.CallOption) (*EndPhaseRelationResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EndPhaseRelation(ctx, in, opts...)
}

func (m *defaultUser) ListPhaseRelation(ctx context.Context, in *ListPhaseRelationReq, opts ...grpc.CallOption) (*ListPhaseRelationResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.ListPhaseRelation(ctx, in, opts...)
}