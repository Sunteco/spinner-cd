package handler

import (
	"go-project-template/internal/service"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestHandler(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, "Project with version v1.0.0")
	}
}
