// Code generated by goctl. DO NOT EDIT.
package types

type AddInstructReq struct {
	StuID     string `json:"stuId"`
	TeacherID string `json:"teacherId"`
	Permit    string `json:"permit"`
}

type AddInstructResp struct {
	Ok int64 `json:"ok"`
}

type AddPhaseRelationReq struct {
	Id    string   `json:"id"`
	Users []string `json:"users,optional"`
	Grade string   `json:"grade,optional"`
}

type AddPhaseRelationResp struct {
	Ok int64 `json:"ok"`
}

type AddPhaseReq struct {
	Term    string `json:"term"`
	Process int64  `json:"process"`
}

type AddPhaseResp struct {
	Ok      int64  `json:"ok"`
	Phaseid string `json:"phaseid"`
}

type DownloadReq struct {
	Filepath string `json:"filepath"`
}

type DownloadResp struct {
	Ok int64 `json:"ok"`
}

type EndPhaseRelationReq struct {
	Id string `json:"id"`
}

type EndPhaseRelationResp struct {
	Ok int64 `json:"ok"`
}

type GetInstructByStuIdReq struct {
	StuID string `form:"stuId"`
}

type GetInstructByStuIdResp struct {
	Instructs []InstructRelation `json:"instructs"`
}

type GetListByTeacherIdReq struct {
	TeacherID string `form:"teacherId"`
}

type GetListByTeacherIdResp struct {
	Instructs []InstructRelation `json:"instructs"`
}

type GetUserCondPageReq struct {
	Grade    string `form:"grade,optional"`
	Identity string `form:"identity,optional"`
	PageNum  int64  `form:"pageNum"`
	PageSize int64  `form:"pageSize"`
}

type GetUserCondPageResp struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
	Count int64  `json:"count"`
}

type GetUserInfoReq struct {
	Id string `form:"id"`
}

type GetUserInfoResp struct {
	User
}

type GetUserPageReq struct {
	PageNum  int64 `form:"pageNum"`
	PageSize int64 `form:"pageSize"`
}

type GetUserPageResp struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
	Count int64  `json:"count"`
}

type InstructRelation struct {
	Id          string `json:"id"`
	StuID       string `json:"stuId"`
	StuName     string `json:"stuName"`
	TeacherID   string `json:"teacherId"`
	TeacherName string `json:"teacherName"`
	Permit      string `json:"permit"`
	State       string `json:"state"`
	Title       string `json:"title"`
}

type ListPhaseRelationReq struct {
	Id       string `form:"id"`
	Identity string `form:"identity"`
	PageNum  int64  `form:"pageNum"`
	PageSize int64  `form:"pageSize"`
}

type ListPhaseRelationResp struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
	Count int64  `json:"count"`
}

type ListTermSortReq struct {
	PageNum  int64 `form:"pageNum"`
	PageSize int64 `form:"pageSize"`
}

type ListTermSortResp struct {
	Phases []Phase `json:"phases"`
	Total  int64   `json:"total"`
	Count  int64   `json:"count"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	ExpireAt string `json:"expireAt"`
}

type Phase struct {
	ID      string `json:"id"`
	Term    string `json:"term"`
	Process int64  `json:"process"`
}

type PhaseRelation struct {
	Id       string `json:"id"`
	PhaseId  string `json:"phaseId"`
	UserId   string `json:"userId"`
	Identity string `json:"identity"`
	Pass     int64  `json:"pass"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `json:"identity"`
	Grade    string `json:"grade"`
}

type RegisterResp struct {
	Ok int64 `json:"ok"`
}

type UpdateInstructReq struct {
	Id     string `json:"id"`
	Permit string `json:"permit"`
}

type UpdateInstructResp struct {
	Updated int64 `json:"updated"`
}

type UpdateProcessReq struct {
	Id      string `json:"id"`
	Process int64  `json:"process"`
}

type UpdateProcessResp struct {
	Updated int64 `json:"updated"`
}

type UpdateStateReq struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

type UpdateStateResp struct {
	Updated int64 `json:"updated"`
}

type UpdateTitleReq struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type UpdateTitleResp struct {
	Updated int64 `json:"updated"`
}

type UpdateUserInfoReq struct {
	ID       string `json:"id"`
	Password string `json:"password,optional"`
	Grade    string `json:"grade,optional"`
}

type UpdateUserInfoResp struct {
	Updated int64 `json:"updated"`
}

type UploadFileReq struct {
	UploadFile string `json:"upload,optional"`
}

type UploadFileResp struct {
	Filename string `json:"filename"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,optional"`
	Identity string `json:"identity,optional"`
	Grade    string `json:"grade,optional"`
}
