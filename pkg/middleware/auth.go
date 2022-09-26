package middleware

import (
	"net/http"

	"github.com/nhanpham699/demo/config"
	"github.com/nhanpham699/demo/pkg/auth"
	"github.com/nhanpham699/demo/pkg/respond"
)

func Auth(h http.HandlerFunc, em *config.ErrorMessage) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenpath := auth.ExtractToken(r)
		if tokenpath == "" {
			respond.JSON(w, http.StatusUnauthorized, &em.InvalidValue.FailedAuthentication)
			return
		}
		_, err := auth.IsAuthorized(tokenpath)

		if err != nil {
			respond.JSON(w, http.StatusUnauthorized, &em.InvalidValue.FailedAuthentication)
			return
		}

		h.ServeHTTP(w, r)
	})
}
