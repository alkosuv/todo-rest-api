package middleware

type ctxKay int8

const (
	// CtxKeyUser ...
	CtxKeyUser ctxKay = iota

	// CtxKeyRequestID ...
	CtxKeyRequestID ctxKay = iota
)
