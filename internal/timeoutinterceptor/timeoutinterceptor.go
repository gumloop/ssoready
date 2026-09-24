package timeoutinterceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
)

// New returns an interceptor that bounds each RPC's context with the given timeout.
func New(timeout time.Duration) connect.UnaryInterceptorFunc {
	return func(f connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return f(ctx, req)
		}
	}
}
