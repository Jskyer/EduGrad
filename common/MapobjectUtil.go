package common

import (
	pamodel "edu-grad/paper/model"
	"edu-grad/paper/rpc/types/proto/paper"
	"edu-grad/user/model"
	"edu-grad/user/rpc/types/proto/user"
)

func ConvertUser(usermodel *model.User) *user.UserInfo {
	return &user.UserInfo{
		Id:       usermodel.ID.Hex(),
		Username: usermodel.Username,
		Identity: usermodel.Identity,
		Grade:    usermodel.Grade,
	}
}

func ConvertInstruct(instruct *model.InstructRelation) *user.InstructRelation {
	return &user.InstructRelation{
		Id:          instruct.ID.Hex(),
		StuId:       instruct.StuID,
		StuName:     instruct.StuName,
		TeacherId:   instruct.TeacherID,
		TeacherName: instruct.TeacherName,
		Permit:      instruct.Permit,
		Title:       instruct.Title,
		State:       instruct.State,
	}
}

func ConvertPhase(phase *model.Phase) *user.Phase {
	return &user.Phase{
		Id:      phase.ID.Hex(),
		Term:    phase.Term,
		Process: phase.Process,
	}
}

func ConvertPaperFormal(paperFormal *pamodel.PaperFormal) *paper.PaperFormal {
	return &paper.PaperFormal{
		Id:         paperFormal.ID.Hex(),
		InstructId: paperFormal.InstructId,
		Content:    paperFormal.Content,
		Comment:    paperFormal.Comment,
		CreateAt:   paperFormal.CreateAt.String(),
	}
}

func ConvertProposalDraft(proposalDraft *pamodel.ProposalDraft) *paper.ProposalDraft {
	return &paper.ProposalDraft{
		Id:         proposalDraft.ID.Hex(),
		InstructId: proposalDraft.InstructId,
		Content:    proposalDraft.Content,
		Comment:    proposalDraft.Comment,
		CreateAt:   proposalDraft.CreateAt.String(),
	}
}

func ConvertProposalFormal(proposalFormal *pamodel.ProposalFormal) *paper.ProposalFormal {
	return &paper.ProposalFormal{
		Id:         proposalFormal.ID.Hex(),
		InstructId: proposalFormal.InstructId,
		Content:    proposalFormal.Content,
		Comment:    proposalFormal.Comment,
		CreateAt:   proposalFormal.CreateAt.String(),
	}
}

func ConvertPaperDraft(paperDraft *pamodel.PaperDraft) *paper.PaperDraft {
	return &paper.PaperDraft{
		Id:         paperDraft.ID.Hex(),
		InstructId: paperDraft.InstructId,
		Content:    paperDraft.Content,
		Comment:    paperDraft.Comment,
		CreateAt:   paperDraft.CreateAt.String(),
	}
}

func ConvertPaperFinal(paperFinal *pamodel.PaperFinal) *paper.PaperFinal {
	return &paper.PaperFinal{
		Id:          paperFinal.ID.Hex(),
		InstructId:  paperFinal.InstructId,
		Content:     paperFinal.Content,
		Comment:     paperFinal.Comment,
		CheckResult: paperFinal.CheckResult,
		CreateAt:    paperFinal.CreateAt.String(),
	}
}
