package handler

import (
	"net/http"

	"go-zero-devops/api/internal/logic"
	"go-zero-devops/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetAllUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUsers()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
