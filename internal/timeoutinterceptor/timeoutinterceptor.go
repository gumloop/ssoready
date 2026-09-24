package timeoutinterceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
)

// New returns an interceptor that bounds every RPC with a context deadline.
// pgx aborts in-flight queries and discards their connections on context
// cancellation, so a dead database connection cannot hold a request open
// past the deadline.
func New(timeout time.Duration) connect.UnaryInterceptorFunc {
	return func(f connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return f(ctx, req)
		}
	}
}
