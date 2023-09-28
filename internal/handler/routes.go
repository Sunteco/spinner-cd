package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"go-project-template/internal/service"
	"net/http"
)

func RegisterHandles(engine *rest.Server, serverCtx *service.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/test",
				Handler: TestHandler(serverCtx),
			},
		},
	)
}
