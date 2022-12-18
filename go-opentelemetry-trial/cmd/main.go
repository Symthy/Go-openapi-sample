package main

import (
	"context"
	"go-opentelemetry-trial/constants"
	"go-opentelemetry-trial/internal/service"
	"log"
	"os"
	"os/signal"

	l_trace "go-opentelemetry-trial/internal/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	// otel 追加 start
	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	exp, err := l_trace.NewExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(l_trace.NewResource(constants.ServiceName, "v0.1.0")),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)
	// otel 追加 end

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	errCh := make(chan error)
	app := service.NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
