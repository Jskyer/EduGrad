package common

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func HttpResponse(r *http.Request, w http.ResponseWriter, resp any, err error) {
	if err != nil {
		body := Body{
			Code: -100,
			Msg:  err.Error(),
			Data: nil,
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}

	body := Body{
		Code: 200,
		Msg:  "success",
		Data: resp,
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
