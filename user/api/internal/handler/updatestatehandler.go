package handler

import (
	"net/http"

	"edu-grad/common"
	"edu-grad/user/api/internal/logic"
	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateStateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateStateLogic(r.Context(), svcCtx)
		resp, err := l.UpdateState(&req)
		common.HttpResponse(r, w, resp, err)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
