package filters

import (
	"github.com/spidernet-io/spiderdoctor/pkg/apiserver/pkg/request"
	"net/http"
)

func WithRequestQuery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req = req.WithContext(request.WithRequestQuery(req.Context(), req.URL.Query()))
		handler.ServeHTTP(w, req)
	})
}
