package main

import (
  "github.com/baxterjfinch/gotochain/handlers"
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()

  e.GET("/health-check", handlers.HealthCheck)
  e.GET("/block-number", handlers.BlockNumber)

  e.Logger.Fatal(e.Start(":9090"))
}
