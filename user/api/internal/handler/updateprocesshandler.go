package handler

import (
	"net/http"

	"edu-grad/common"
	"edu-grad/user/api/internal/logic"
	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateProcessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProcessReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateProcessLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProcess(&req)
		common.HttpResponse(r, w, resp, err)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
