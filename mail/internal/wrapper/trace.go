package wrapper

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/opentracing/opentracing-go"
)

type otWrapper struct {
	ot opentracing.Tracer
	client.Client
}

func traceIntoContext(ctx context.Context, tracer opentracing.Tracer, name string) (context.Context, opentracing.Span, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	if err != nil {
		sp = tracer.StartSpan(name)
	} else {
		sp = tracer.StartSpan(name, opentracing.ChildOf(wireContext))
	}
	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
		return nil, nil, err
	}
	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	return ctx, sp, nil
}

func (o *otWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	name := fmt.Sprintf("%s.%s", req.Service(), req.Method())
	ctx, span, err := traceIntoContext(ctx, o.ot, name)
	if err != nil {
		return err
	}
	defer span.Finish()
	return o.Client.Call(ctx, req, rsp, opts...)
}

func (o *otWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	name := fmt.Sprintf("Pub to %s", p.Topic())
	ctx, span, err := traceIntoContext(ctx, o.ot, name)
	span.LogKV("event", "sub: "+p.Topic())
	if err != nil {
		return err
	}
	defer span.Finish()
	return o.Client.Publish(ctx, p, opts...)
}

// NewClientWrapper accepts an open tracing Trace and returns a Client Wrapper
func NewTraceClientWrapper(ot opentracing.Tracer) client.Wrapper {
	return func(c client.Client) client.Client {
		return &otWrapper{ot, c}
	}
}

// NewHandlerWrapper accepts an opentracing Tracer and returns a Call Wrapper
func NewTraceCallWrapper(ot opentracing.Tracer) client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Method())
			ctx, span, err := traceIntoContext(ctx, ot, name)
			span.SetOperationName(req.Method())
			span.LogKV("event", "call: "+name)
			if err != nil {
				return err
			}
			defer span.Finish()
			err = cf(ctx, addr, req, rsp, opts)
			span.LogKV("event", "return: "+name)

			return err
		}
	}
}

// NewHandlerWrapper accepts an opentracing Tracer and returns a Handler Wrapper
func NewTraceHandlerWrapper(ot opentracing.Tracer) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Method())
			ctx, span, err := traceIntoContext(ctx, ot, name)
			span.SetOperationName(req.Method())
			span.LogKV("event", "call: "+name)
			if err != nil {
				return err
			}
			defer span.Finish()
			err = h(ctx, req, rsp)
			span.LogKV("event", "return: "+name)

			return err
		}
	}
}

// NewSubscriberWrapper accepts an opentracing Tracer and returns a Subscriber Wrapper
func NewTraceSubscriberWrapper(ot opentracing.Tracer) server.SubscriberWrapper {
	return func(next server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, msg server.Message) error {
			name := "Sub to " + msg.Topic()
			ctx, span, err := traceIntoContext(ctx, ot, name)
			span.LogKV("event", "sub: "+msg.Topic())
			if err != nil {
				return err
			}
			defer span.Finish()
			err = next(ctx, msg)

			return err
		}
	}
}
