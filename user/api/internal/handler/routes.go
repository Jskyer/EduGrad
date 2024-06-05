// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"edu-grad/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/condpage",
				Handler: getUserCondPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/downloadFile",
				Handler: downloadFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: getUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info/update",
				Handler: updateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/page",
				Handler: getUserPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: registerUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/uploadFile",
				Handler: uploadFileHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: addInstructHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listTeacher",
				Handler: getListByTeacherIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oneStu",
				Handler: getInstructByStuIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stulistForTeacher",
				Handler: stulistForTeacherHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: updateInstructHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateState",
				Handler: updateStateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateTitle",
				Handler: updateTitleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/instruct"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: addPhaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/termlist",
				Handler: listTermSortHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateProcess",
				Handler: updateProcessHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/phase"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: addPhaseRelationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/allowedtist",
				Handler: PhaseAllowedtistHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comittedtist",
				Handler: PhaseCommittedtistHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/end",
				Handler: endPhaseRelationHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/identitylist",
				Handler: listPhaseRelationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/inphase",
				Handler: InphaseRelationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/phaseRelation"),
	)
}
