package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"edu-grad/common"
	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func downloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if !strings.Contains(req.Filepath, ".pdf") {
			common.HttpResponse(r, w, nil, errors.New("请求文件格式不合法"))
			return
		}

		prefix := "C:/Myfiles/tmp/"
		// prefix := "/app/goFile/"

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

		// resp := &types.DownloadResp{
		// 	Ok: 1,
		// }
		// common.HttpResponse(r, w, resp, nil)

		// l := logic.NewDownloadFileLogic(r.Context(), svcCtx)
		// resp, err := l.DownloadFile(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
