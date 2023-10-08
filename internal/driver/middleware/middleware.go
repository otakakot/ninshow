package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CORS(hdl http.Handler) http.Handler {
	return cors.AllowAll().Handler(hdl)
}
