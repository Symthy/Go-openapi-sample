package service

import (
	"context"
	"fmt"
	"go-opentelemetry-trial/constants"
	"io"
	"log"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// App is a Fibonacci computation application.
type App struct {
	reader io.Reader
	logger *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{reader: r, logger: l}
}

// Run starts polling users for Fibonacci number requests and writes results.
func (a *App) Run(ctx context.Context) error {
	for {
		// ループを実行するたびに、新しい「ルート スパンとコンテキストを取得する必要がある
		newCtx, span := otel.Tracer(constants.ServiceName).Start(ctx, "Run") // Otel追加

		n, err := a.Poll(newCtx)
		if err != nil {
			return err
		}

		a.Write(ctx, n)
		span.End() // Otel追加
	}
}

// Poll asks a user for input and returns the request.
func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := otel.Tracer(constants.ServiceName).Start(ctx, "Poll") // Otel追加
	defer span.End()                                                 // Otel追加

	a.logger.Print("What Fibonacci number would you like to know: ")

	var n uint
	_, err := fmt.Fscanf(a.reader, "%d\n", &n)

	// int 64をオーバーフローさせないように、nを文字列として格納
	nStr := strconv.FormatUint(uint64(n), 10)               // Otel追加
	span.SetAttributes(attribute.String("request.n", nStr)) // Otel追加

	return n, err
}

// Write writes the n-th Fibonacci number back to the user.
func (a *App) Write(ctx context.Context, n uint) {
	var span trace.Span
	ctx, span = otel.Tracer(constants.ServiceName).Start(ctx, "Write") // Otel追加
	defer span.End()                                                   // Otel追加

	// f, err := Fibonacci(n)
	f, err := func(ctx context.Context) (uint64, error) { // Otel追加
		_, span := otel.Tracer(constants.ServiceName).Start(ctx, "Fibonacci")
		defer span.End()
		return Fibonacci(n)
	}(ctx)

	if err != nil {
		a.logger.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		a.logger.Printf("Fibonacci(%d) = %d\n", n, f)
	}
}
