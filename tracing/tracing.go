package tracing

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// New registers Jaeger as the OpenTracing implementation.
// If jaegerAgentHost is an empty string, tracing is disabled.
func New(jaegerAgentHost, serviceName, samplerType string, samplerParam float64) io.Closer {
	if jaegerAgentHost != "" {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				SamplingServerURL: fmt.Sprintf("http://%s:5778/sampling", jaegerAgentHost),
				Type:              samplerType,
				Param:             samplerParam,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LocalAgentHostPort: fmt.Sprintf("%s:6831", jaegerAgentHost),
			},
		}

		closer, err := cfg.InitGlobalTracer(serviceName)
		if err != nil {
			fmt.Printf("Could not initialize jaeger tracer: %s\n", err.Error())
			os.Exit(1)
		}
		return closer
	}
	return ioutil.NopCloser(nil)
}
