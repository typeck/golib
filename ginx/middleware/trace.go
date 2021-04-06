package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/typeck/golib/trace"
)

// Trace 链路追踪
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sp opentracing.Span
		opName := ctx.Request.URL.Path
		// Attempt to join a trace by getting trace context from the headers.
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			// If for whatever reason we can't join, go ahead an start a new root span.
			sp = opentracing.StartSpan(opName)
		} else {
			sp = opentracing.StartSpan(opName,
				opentracing.ChildOf(wireContext),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindConsumer,
			)
		}
		ctxt := opentracing.ContextWithSpan(ctx, sp)
		trace.SetCtx(ctxt, ctx)
		//spanCtx := opentracing.ContextWithSpan(ctx.Request.Context(), sp)
		//ctx.Request.WithContext(spanCtx)
		ctx.Next()
		opentracing.GlobalTracer().Inject(sp.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		sp.Finish()
	}
}


