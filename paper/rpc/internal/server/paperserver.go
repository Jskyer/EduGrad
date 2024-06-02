// Code generated by goctl. DO NOT EDIT.
// Source: paper.proto

package server

import (
	"context"

	"edu-grad/paper/rpc/internal/logic"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"
)

type PaperServer struct {
	svcCtx *svc.ServiceContext
	paper.UnimplementedPaperServer
}

func NewPaperServer(svcCtx *svc.ServiceContext) *PaperServer {
	return &PaperServer{
		svcCtx: svcCtx,
	}
}

func (s *PaperServer) UploadPaperFormal(ctx context.Context, in *paper.UploadFileReq) (*paper.UploadPaperFormalResp, error) {
	l := logic.NewUploadPaperFormalLogic(ctx, s.svcCtx)
	return l.UploadPaperFormal(in)
}

// rpc downloadPaperFormal(DownloadFileReq) returns (DownloadPaperFormalResp);
func (s *PaperServer) ListPaperFormal(ctx context.Context, in *paper.ListPaperFormalReq) (*paper.ListPaperFormalResp, error) {
	l := logic.NewListPaperFormalLogic(ctx, s.svcCtx)
	return l.ListPaperFormal(in)
}

func (s *PaperServer) CommentPaperFormal(ctx context.Context, in *paper.CommentPaperFormalReq) (*paper.CommentPaperFormalResp, error) {
	l := logic.NewCommentPaperFormalLogic(ctx, s.svcCtx)
	return l.CommentPaperFormal(in)
}

func (s *PaperServer) UploadProposalDraft(ctx context.Context, in *paper.UploadFileReq) (*paper.UploadProposalDraftResp, error) {
	l := logic.NewUploadProposalDraftLogic(ctx, s.svcCtx)
	return l.UploadProposalDraft(in)
}

func (s *PaperServer) ListProposalDraft(ctx context.Context, in *paper.ListProposalDraftReq) (*paper.ListProposalDraftResp, error) {
	l := logic.NewListProposalDraftLogic(ctx, s.svcCtx)
	return l.ListProposalDraft(in)
}

func (s *PaperServer) CommentProposalDraft(ctx context.Context, in *paper.CommentProposalDraftReq) (*paper.CommentProposalDraftResp, error) {
	l := logic.NewCommentProposalDraftLogic(ctx, s.svcCtx)
	return l.CommentProposalDraft(in)
}

func (s *PaperServer) UploadProposalFormal(ctx context.Context, in *paper.UploadFileReq) (*paper.UploadProposalFormalResp, error) {
	l := logic.NewUploadProposalFormalLogic(ctx, s.svcCtx)
	return l.UploadProposalFormal(in)
}

func (s *PaperServer) ListProposalFormal(ctx context.Context, in *paper.ListProposalFormalReq) (*paper.ListProposalFormalResp, error) {
	l := logic.NewListProposalFormalLogic(ctx, s.svcCtx)
	return l.ListProposalFormal(in)
}

func (s *PaperServer) CommentProposalFormal(ctx context.Context, in *paper.CommentProposalFormalReq) (*paper.CommentProposalFormalResp, error) {
	l := logic.NewCommentProposalFormalLogic(ctx, s.svcCtx)
	return l.CommentProposalFormal(in)
}

func (s *PaperServer) UploadPaperDraft(ctx context.Context, in *paper.UploadFileReq) (*paper.UploadPaperDraftResp, error) {
	l := logic.NewUploadPaperDraftLogic(ctx, s.svcCtx)
	return l.UploadPaperDraft(in)
}

func (s *PaperServer) ListPaperDraft(ctx context.Context, in *paper.ListPaperDraftReq) (*paper.ListPaperDraftResp, error) {
	l := logic.NewListPaperDraftLogic(ctx, s.svcCtx)
	return l.ListPaperDraft(in)
}

func (s *PaperServer) CommentPaperDraft(ctx context.Context, in *paper.CommentPaperDraftReq) (*paper.CommentPaperDraftResp, error) {
	l := logic.NewCommentPaperDraftLogic(ctx, s.svcCtx)
	return l.CommentPaperDraft(in)
}

func (s *PaperServer) UploadPaperFinal(ctx context.Context, in *paper.UploadFileReq) (*paper.UploadPaperFinalResp, error) {
	l := logic.NewUploadPaperFinalLogic(ctx, s.svcCtx)
	return l.UploadPaperFinal(in)
}

func (s *PaperServer) ListPaperFinal(ctx context.Context, in *paper.ListPaperFinalReq) (*paper.ListPaperFinalResp, error) {
	l := logic.NewListPaperFinalLogic(ctx, s.svcCtx)
	return l.ListPaperFinal(in)
}

func (s *PaperServer) CommentPaperFinal(ctx context.Context, in *paper.CommentPaperFinalReq) (*paper.CommentPaperFinalResp, error) {
	l := logic.NewCommentPaperFinalLogic(ctx, s.svcCtx)
	return l.CommentPaperFinal(in)
}

func (s *PaperServer) CheckPaperFinal(ctx context.Context, in *paper.CheckPaperFinalReq) (*paper.CheckPaperFinalResp, error) {
	l := logic.NewCheckPaperFinalLogic(ctx, s.svcCtx)
	return l.CheckPaperFinal(in)
}
