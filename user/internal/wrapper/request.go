package wrapper

import (
	"context"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/rs/xid"
)

// RequestIdHandlerWrapper simply adds a X-Request-ID header if none is given.
// Usually the request ID is calculated on the client and send as header.
func RequestIdWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		md, _ := metadata.FromContext(ctx)
		rid := md["x-request-id"]
		if rid == "" {
			md["x-request-id"] = xid.New().String()
			ctx = metadata.NewContext(ctx, md)
		}
		err := fn(ctx, req, rsp)
		return err
	}
}
