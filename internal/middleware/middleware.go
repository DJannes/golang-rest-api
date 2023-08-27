package middleware

import (
	"net/http"
	"strings"

	"gitlab.com/janneseffendi/rest-api/internal/security"
)

func TokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Write([]byte("Authorization Required"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		bearerPrefix := "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			w.Write([]byte("Bearer Authrization Required"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		bearerToken, _ := strings.CutPrefix(authHeader, bearerPrefix)
		pasetoGen := security.NewMockPasetoGen()
		payload, err := pasetoGen.VerifyToken(bearerToken)
		if err != nil {
			w.Write([]byte("token invalid" + err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if payload.IsExpired() {
			w.Write([]byte("token expired"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = setPayloadToReqCtx(r, payload)
		next.ServeHTTP(w, r)
	})
}
