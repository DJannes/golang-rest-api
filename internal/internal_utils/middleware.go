package internal_utils

import (
	"fmt"
	"net/http"
	"strings"
)

func TokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(HTTP_AUTHORIZATION_HEADER)
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
		pasetoGen := NewMockPasetoGen()
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

		fmt.Println(payload)
		next.ServeHTTP(w, r)
	})
}
