package handler

import (
	"net/http"

	"edu-grad/common"
	"edu-grad/paper/api/internal/logic"
	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func commentPaperFormalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentPaperFormalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentPaperFormalLogic(r.Context(), svcCtx)
		resp, err := l.CommentPaperFormal(&req)
		common.HttpResponse(r, w, resp, err)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
