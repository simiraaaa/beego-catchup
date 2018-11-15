package basicauth

import (
	"net/http"
)

// BasicAuth ... ベーシック認証機能を提供するミドルウェア
type BasicAuth struct {
	Config BasicAuthConfig
}

// Handle ... ハンドラ
func (a *BasicAuth) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "basic auth required.", http.StatusUnauthorized)
			return
		}
		if a.Config.Accounts[user] != password {
			http.Error(w, "basic auth error.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewBasicAuth ... BasicAuthを作成する
func NewBasicAuth(cfg BasicAuthConfig) *BasicAuth {
	return &BasicAuth{
		Config: cfg,
	}
}
