package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"edu-grad/common"
	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func downloadProposalFormalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if !strings.Contains(req.Filepath, ".pdf") {
			common.HttpResponse(r, w, nil, errors.New("请求文件格式不合法"))
			return
		}

		// prefix := "C:/Myfiles/tmp/"
		prefix := "/app/goFile/"

		file, err := os.Open(prefix + req.Filepath)
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		// 设置http响应头
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		_, err = io.Copy(w, file)
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		logx.Info("complete file download")

		// l := logic.NewDownloadProposalFormalLogic(r.Context(), svcCtx)
		// resp, err := l.DownloadProposalFormal(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
