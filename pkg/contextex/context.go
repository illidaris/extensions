package contextex

import "context"

type Key int32

const (
	KeyClientIP Key = iota
	GameTag
	KeyUser
	Biz
)

func (k Key) Save(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, k, value)
}

func (k Key) Get(ctx context.Context) string {
	v := ctx.Value(k)
	if v == nil {
		return ""
	}
	if result, ok := v.(string); ok {
		return result
	}
	return ""
}
