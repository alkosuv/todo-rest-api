package ctxkey

type ctxKey int8

const (
	// CtxKeyUser ...
	CtxKeyUser ctxKey = iota

	// CtxKeyRequestID ...
	CtxKeyRequestID ctxKey = iota
)
