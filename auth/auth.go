package auth

import (
	"metric_service/config"
	constant "metric_service/utils"
	"net/http"
)

func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username, password, _ := r.BasicAuth()
        if username != config.AppConfig.Auth.Username || password != config.AppConfig.Auth.Password {
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, constant.MsgInvalidAuth, http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}
