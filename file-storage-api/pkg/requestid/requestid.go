package requestid

import "context"

type reqIdx int

const (
	reqKey = iota + 1
)

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, reqKey, requestId)
}

func GetRequestId(ctx context.Context) string {
	val := ctx.Value(reqKey)

	if val == nil {
		return ""
	}

	valStr, ok := val.(string)

	if !ok {
		return ""
	}

	return valStr
}
