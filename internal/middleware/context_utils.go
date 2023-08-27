package middleware

import (
	"context"
	"errors"
	"net/http"

	"gitlab.com/janneseffendi/rest-api/internal/security"
)

type internalKey struct {
	key string
}

var (
	payloadCtxKey = &internalKey{key: "payload"}
)

func setPayloadToReqCtx(r *http.Request, payload *security.Payload) *http.Request {
	newCtx := context.WithValue(r.Context(), payloadCtxKey, payload)
	return r.WithContext(newCtx)
}

func MustGetPayload(ctx context.Context) *security.Payload {
	payload, err := GetPayload(ctx)
	if err != nil {
		panic(err)
	}
	return payload
}

func GetPayload(ctx context.Context) (*security.Payload, error) {
	ctxValue := ctx.Value(payloadCtxKey)
	payload, ok := ctxValue.(*security.Payload)
	if !ok {
		return nil, errors.New("Payload must be set to context")
	}

	return payload, nil
}
