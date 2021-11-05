package handlers

import (
  "context"
  "net/http"
  "github.com/baxterjfinch/gotochain/service"
  "github.com/baxterjfinch/gotochain/renderings"
  "github.com/labstack/echo"
  "fmt"
)

func BlockNumber(c echo.Context) error {
  var message string

  client, err := service.EthConnect(c)
  if err != nil {
    message = fmt.Sprintf("Error Connecting To Client: %d", err)
  }

  blockNumber, err := client.BlockNumber(context.Background())
  if err != nil {
    message = fmt.Sprintf("Error Getting Block Number: %d", err)
  } else {
    message = fmt.Sprintf("%d", blockNumber)
  }

  response := renderings.BlockNumberResponse{
    BlockNumber: message,
  }

  return c.JSON(http.StatusOK, response)
}