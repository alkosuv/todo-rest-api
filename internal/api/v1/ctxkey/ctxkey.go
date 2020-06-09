package ctxkey

type ctxKey int8

const (
	// CtxKeyUser ключ для контехата User
	CtxKeyUser ctxKey = iota

	// CtxKeyRequestID ключ для RequestID
	CtxKeyRequestID ctxKey = iota
)
