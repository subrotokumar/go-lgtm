package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subrotokumar/go-lgtm/api"
	"github.com/subrotokumar/go-lgtm/observability"
)

func main() {
	ctx := context.Background()

	telemConfig, err := observability.NewConfigFromEnv()
	if err != nil {
		fmt.Println("failed to load telemetry config")
		os.Exit(1)
	}

	// Initialize telemetry. If the exporter fails, fallback to nop.
	var telem observability.TelemetryProvider
	telem, err = observability.NewTelemetry(ctx, telemConfig)
	if err != nil {
		fmt.Println("failed to create telemetry, falling back to no-op telemetry")
		telem, _ = observability.NewNoopTelemetry(telemConfig)
	}
	defer telem.Shutdown(ctx)

	telem.LogInfo("telemetry initialized")

	r := gin.New()
	r.Use(telem.LogRequest())
	r.Use(telem.MeterRequestDuration())
	r.Use(telem.MeterRequestsInFlight())

	api := api.NewAPI(
		telem, &http.Server{
			Addr:    "0.0.0.0:3001",
			Handler: r,
		},
	)

	r.GET("/", api.GetSomething)
	api.Start()
}
