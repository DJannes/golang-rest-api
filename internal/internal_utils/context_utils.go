package internal_utils

import (
	"context"
	"errors"
	"net/http"
)

type internalKey struct {
	key string
}

var (
	payloadCtxKey = &internalKey{key: "payload"}
)

func setPayloadToReqCtx(r *http.Request, payload *Payload) *http.Request {
	newCtx := context.WithValue(r.Context(), payloadCtxKey, payload)
	return r.WithContext(newCtx)
}

func MustGetPayloadFromReqCtx(ctx context.Context) *Payload {
	payload, err := GetPayloadFromReqCtx(ctx)
	if err != nil {
		panic(err)
	}
	return payload
}

func GetPayloadFromReqCtx(ctx context.Context) (*Payload, error) {
	ctxValue := ctx.Value(payloadCtxKey)
	payload, ok := ctxValue.(*Payload)
	if !ok {
		return nil, errors.New("Payload must be set to context")
	}

	return payload, nil
}
