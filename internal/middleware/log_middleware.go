package middleware

import (
	"bytes"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w = NewLogResponseWriter(w)
		next.ServeHTTP(w, r)
	})
}

type LogResponseWriter struct {
	http.ResponseWriter
	resBody *bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	return &LogResponseWriter{
		ResponseWriter: w,
		resBody:        new(bytes.Buffer),
	}
}

func (w *LogResponseWriter) Write(b []byte) (int, error) {
	w.resBody.Write(b)
	return w.ResponseWriter.Write(b)
}
