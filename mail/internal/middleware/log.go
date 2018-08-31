package middleware

import (
	"github.com/micro/go-micro/client"
	"context"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/metadata"
	"github.com/rs/zerolog/log"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption)	error {
	logRequest(ctx, req, nil)
	return l.Client.Call(ctx, req, rsp)
}

// Implements the server.HandlerWrapper
func LogHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		logRequest(ctx, req, nil)

		return fn(ctx, req, rsp)
	}
}

func LogSubscriberWrapper(fn server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		logRequest(ctx, nil, msg)
		return fn(ctx, msg)
	}
}

func (l *logWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error  {
	logRequest(ctx, nil, p)
	return l.Client.Publish(ctx, p, opts...)
}

func logRequest(ctx context.Context, req server.Request, msg server.Message) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}

	userId := md["X-User-Id"]
	traceId := md["x-b3-traceid"]
	spanId := md["x-b3-spanid"]
	from := md["X-Micro-From-Service"]

	method := ""
	if req != nil {
		method = req.Method()
	}

	if msg != nil {
		method = msg.Topic()
	}

	log.Debug().
		Str("trace", traceId).
		Str("user", userId).
		Str("span", spanId).
		Str("from_service", from).
		Msg(method)
}
