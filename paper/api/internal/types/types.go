// Code generated by goctl. DO NOT EDIT.
package types

type CheckPaperFinalReq struct {
	Id          string `json:"id"`
	CheckResult string `json:"checkResult"`
}

type CheckPaperFinalResp struct {
	Ok int64 `json:"ok"`
}

type CommentPaperDraftReq struct {
	Id      string `json:"id"`
	Comment string `json:"comment"`
}

type CommentPaperDraftResp struct {
	Ok int64 `json:"ok"`
}

type CommentPaperFinalReq struct {
	Id      string `json:"id"`
	Comment string `json:"comment"`
}

type CommentPaperFinalResp struct {
	Ok int64 `json:"ok"`
}

type CommentPaperFormalReq struct {
	Id      string `json:"id"`
	Comment string `json:"comment"`
}

type CommentPaperFormalResp struct {
	Ok int64 `json:"ok"`
}

type CommentProposalDraftReq struct {
	Id      string `json:"id"`
	Comment string `json:"comment"`
}

type CommentProposalDraftResp struct {
	Ok int64 `json:"ok"`
}

type CommentProposalFormalReq struct {
	Id      string `json:"id"`
	Comment string `json:"comment"`
}

type CommentProposalFormalResp struct {
	Ok int64 `json:"ok"`
}

type DownloadFileReq struct {
	Filepath string `json:"filepath"`
}

type DownloadPaperDraftResp struct {
}

type DownloadPaperFinalResp struct {
}

type DownloadPaperFormalResp struct {
}

type DownloadProposalDraftResp struct {
}

type DownloadProposalFormalResp struct {
}

type ListPaperDraftReq struct {
	InstructId string `form:"instructId"`
}

type ListPaperDraftResp struct {
	PaperDrafts []PaperDraft `json:"paperDrafts"`
}

type ListPaperFinalReq struct {
	InstructId string `form:"instructId"`
}

type ListPaperFinalResp struct {
	PaperFinals []PaperFinal `json:"paperFinals"`
}

type ListPaperFormalReq struct {
	InstructId string `form:"instructId"`
}

type ListPaperFormalResp struct {
	PaperFormals []PaperFormal `json:"paperFormals"`
}

type ListProposalDraftReq struct {
	InstructId string `form:"instructId"`
}

type ListProposalDraftResp struct {
	ProposalDrafts []ProposalDraft `json:"proposalDrafts"`
}

type ListProposalFormalReq struct {
	InstructId string `form:"instructId"`
}

type ListProposalFormalResp struct {
	ProposalFormals []ProposalFormal `json:"proposalFormals"`
}

type PaperDraft struct {
	ID         string `json:"id"`
	InstructId string `json:"instructId"`
	Content    string `json:"content"`
	Comment    string `json:"comment"`
	CreateAt   string `json:"createAt"`
}

type PaperFinal struct {
	ID          string `json:"id"`
	InstructId  string `json:"instructId"`
	Content     string `json:"content"`
	Comment     string `json:"comment"`
	CheckResult string `json:"checkResult"`
	CreateAt    string `json:"createAt"`
}

type PaperFormal struct {
	ID         string `json:"id"`
	InstructId string `json:"instructId"`
	Content    string `json:"content"`
	Comment    string `json:"comment"`
	CreateAt   string `json:"createAt"`
}

type ProposalDraft struct {
	ID         string `json:"id"`
	InstructId string `json:"instructId"`
	Content    string `json:"content"`
	Comment    string `json:"comment"`
	CreateAt   string `json:"createAt"`
}

type ProposalFormal struct {
	ID         string `json:"id"`
	InstructId string `json:"instructId"`
	Content    string `json:"content"`
	Comment    string `json:"comment"`
	CreateAt   string `json:"createAt"`
}

type UploadFileReq struct {
	InstructId string `form:"instructId"`
	Content    string `form:"content,optional"`
}

type UploadPaperDraftResp struct {
	Filename string `json:"filename"`
}

type UploadPaperFinalResp struct {
	Filename string `json:"filename"`
}

type UploadPaperFormalResp struct {
	Filename string `json:"filename"`
}

type UploadProposalDraftResp struct {
	Filename string `json:"filename"`
}

type UploadProposalFormalResp struct {
	Filename string `json:"filename"`
}
