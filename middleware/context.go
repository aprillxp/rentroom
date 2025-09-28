package middleware

import (
	"errors"
	"net/http"
)

type ctxKey int

const (
	CtxUserID  ctxKey = iota
	CtxAdminID ctxKey = iota
	CtxRole
)

func MustUserID(r *http.Request) (uint, error) {
	id, ok := r.Context().Value(CtxUserID).(int)
	if !ok {
		return 0, errors.New("unauthorized")
	}
	return uint(id), nil
}

func MustAdminID(r *http.Request) error {
	_, ok := r.Context().Value(CtxAdminID).(int)
	if !ok {
		return errors.New("unauthorized")
	}
	return nil
}
