package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"edu-grad/common"
	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func uploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		defer file.Close()

		// 上传文件到指定路径，返回文件全名

		fmt.Println("file: " + header.Filename)

		idx := strings.LastIndex(header.Filename, ".")
		if idx == -1 {
			common.HttpResponse(r, w, nil, errors.New("文件格式不合法"))
			return
		}

		fmt.Println(header.Filename[idx:])

		// 非pdf格式的文件不合法
		if header.Filename[idx:] != ".pdf" {
			common.HttpResponse(r, w, nil, errors.New("文件格式不合法"))
			return
		}

		prefix := "C:/Myfiles/tmp/"
		// prefix := "/app/goFile/"

		filename := uuid.New().String() + "_" + time.Now().Format("2006-01-02_15-04-05") + header.Filename[idx:]

		dstPath := prefix + filename
		dst, err := os.Create(dstPath)
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		defer dst.Close()

		// fmt.Println(dst)

		_, err = io.Copy(dst, file)
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		// l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		// resp, err := l.UploadFile(&req)

		resp := &types.UploadFileResp{
			Filename: filename,
		}
		common.HttpResponse(r, w, resp, nil)

		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
