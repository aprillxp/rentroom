package middleware

import (
	"errors"
	"net/http"
)

type ctxKey int

const (
	CtxUserID ctxKey = iota
	CtxRole
)

func MustUserID(r *http.Request) (int, error) {
	id, ok := r.Context().Value(CtxUserID).(int)
	if !ok {
		return 0, errors.New("unauthorized")
	}
	return id, nil
}
