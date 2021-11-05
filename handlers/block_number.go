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
  var messageErr string

  client, err := service.EthConnect(c)
  if err != nil {
    messageErr = fmt.Sprintf("Error Connecting To Client: %d", err)
  }

  blockNumber, err := client.BlockNumber(context.Background())
  if err != nil {
    messageErr = fmt.Sprintf("Error Getting Block Number: %d", err)
  }

  response := renderings.Response{
    BlockNumber: blockNumber,
    Error: messageErr,
  }

  return c.JSON(http.StatusOK, response)
}