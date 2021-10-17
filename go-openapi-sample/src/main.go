package main

import (
	"flag"
	"fmt"

	"github.com/Symthy/golang-practices/go-openapi-sample/autogen/server"
	"github.com/Symthy/golang-practices/go-openapi-sample/handler"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func main() {
	var port = flag.Int("port", 3030, "Port for test proxy server")
	flag.Parse()

	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Log Request/Response body
	e.Use(echomiddleware.BodyDump(bodyDumpHandler))

	proxyHandler := handler.NewQiitaProxyHandler()

	server.RegisterHandlers(e, proxyHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
