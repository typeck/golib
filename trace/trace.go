package trace

import (
	"context"
	"github.com/uber/jaeger-client-go"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)


var closer io.Closer

func JeagerTracer() {
	cfg := config.Configuration{
		ServiceName: "sands-api-server",
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint: "http://172.27.122.2:14268/api/traces",
		},
	}
	var tracer opentracing.Tracer
	var err error
	tracer, closer, err = cfg.NewTracer()
	if err !=nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	//defer closer.Close()
	//someFunction()
}

func JeagerClose() {
	closer.Close()
}

func someFunction() {
	parent := opentracing.GlobalTracer().StartSpan("hello")
	defer parent.Finish()
	child := opentracing.GlobalTracer().StartSpan(
		"world", opentracing.ChildOf(parent.Context()))
	defer child.Finish()
}

//SetCtx set span ctx into gin context
func SetCtx(ctx context.Context, c *gin.Context ) {
	c.Set("Span-Context", ctx)
}

//GetCtx .
func GetCtx(c *gin.Context) context.Context {
	v, _ := c.Get("Span-Context")
	if ctx, ok := v.(context.Context); ok {
		return ctx
	}
	return c
}
