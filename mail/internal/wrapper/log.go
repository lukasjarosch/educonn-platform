package wrapper

import (
	"context"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/rs/zerolog/log"
)

func NewLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		md, _ := metadata.FromContext(ctx)
		log.Logger = log.Logger.With().
			Str("request_id", md["x-request-id"]).
			Str("trace_id", md["x-b3-traceid"]).
			Logger()
		err := fn(ctx, req, rsp)
		return err
	}
}
