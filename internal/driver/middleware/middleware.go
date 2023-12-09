package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ogen-go/ogen/middleware"
	"github.com/rs/cors"
)

func CORS(hdl http.Handler) http.Handler {
	return cors.AllowAll().Handler(hdl)
}

func Logging() middleware.Middleware {
	return func(
		req middleware.Request,
		next middleware.Next,
	) (middleware.Response, error) {
		slog.Info(fmt.Sprintf("%s %s", req.Raw.Method, req.Raw.URL))

		res, err := next(req)
		if err != nil {
			slog.Error(fmt.Sprintf("Error %v", err))
		} else {
			slog.Info(fmt.Sprintf("Response %T", res.Type))
		}

		return res, err
	}
}
