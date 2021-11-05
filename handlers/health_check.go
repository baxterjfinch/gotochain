package handlers

import (
  "net/http"
  "github.com/baxterjfinch/gotochain/renderings"
  "github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
  resp := renderings.HealthCheckResponse{
    Message: "Everything is good!",
  }
  return c.JSON(http.StatusOK, resp)
}