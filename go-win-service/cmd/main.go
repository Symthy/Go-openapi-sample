package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kardianos/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/natefinch/lumberjack"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

var logger service.Logger

type exarvice struct {
	exit chan struct{}
}

func (s *exarvice) run(l io.Writer) {
	log.Print("[lumberjack] Exarvice Start !!!")
	logger.Info("[Service] Exarvice Start !!!")

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetOutput(l)
	// Routes
	e.GET("/", helloHandler)
	defer e.Close()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	logger.Info("[Service] echo Start !!!")
	// ticker := time.NewTicker(5 * time.Second)
	// for {
	// 	select {
	// 	case tm := <-ticker.C:
	// 		log.Printf("[lumberjack] Still running at %v", tm)
	// 		logger.Infof("[Service] Still running at %v", tm)
	// 	case <-e.exit:
	// 		ticker.Stop()
	// 		log.Print("[lumberjack] Exarvice Stop ...")
	// 		logger.Info("[Service] Exarvice Stop ...")
	// 		return
	// 	}
	// }
}

func (s *exarvice) Start(se service.Service) error {
	l := &lumberjack.Logger{
		Filename: "C:\\work\\win_service_lumberjack.log",
	}
	log.SetOutput(l)

	if service.Interactive() {
		log.Print("[lumberjack] Running in terminal.")
		logger.Info("[Service] Running in terminal.")
	} else {
		log.Print("[lumberjack] Running under service manager.")
		logger.Info("[Service] Running under service manager.")
	}
	s.exit = make(chan struct{})

	go s.run(l)
	return nil
}

func (e *exarvice) Stop(s service.Service) error {
	close(e.exit)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "Exarvice",
		DisplayName: "Exarvice (Go Service Example)",
		Description: "This is an example Go service.",
	}

	// Create Exarvice service
	program := &exarvice{}
	s, err := service.New(program, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the logger
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal()
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Failed (%s) : %s\n", os.Args[1], err)
			return
		}
		fmt.Printf("Succeeded (%s)\n", os.Args[1])
		return
	}

	// run in terminal
	s.Run()
}
