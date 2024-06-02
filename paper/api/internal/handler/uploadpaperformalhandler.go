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
	"edu-grad/paper/api/internal/logic"
	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func uploadPaperFormalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		_, err := primitive.ObjectIDFromHex(req.InstructId)
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		// 同一个instructId最多2份文件
		listLogic := logic.NewListPaperFormalLogic(r.Context(), svcCtx)
		listResp, err := listLogic.ListPaperFormal(&types.ListPaperFormalReq{InstructId: req.InstructId})
		if err != nil {
			common.HttpResponse(r, w, nil, err)
			return
		}

		if len(listResp.PaperFormals) >= 2 {
			common.HttpResponse(r, w, nil, errors.New("同一个instructId最多2份文件"))
			return
		}

		// 执行文件上传逻辑

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

		// prefix := "C:/Myfiles/tmp/"
		prefix := "/app/goFile/"

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

		// 设置content
		req.Content = filename

		l := logic.NewUploadPaperFormalLogic(r.Context(), svcCtx)
		resp, err := l.UploadPaperFormal(&req)
		// if err != nil {
		// 	common.HttpResponse(r, w, nil, err)
		// 	return
		// }

		// resp := &types.UploadPaperFormalResp{
		// 	Filename: filename,
		// }

		common.HttpResponse(r, w, resp, err)

		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
