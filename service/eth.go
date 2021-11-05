package service

import (
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/labstack/echo"
)

func EthConnect(c echo.Context) (*ethclient.Client, error){
  return ethclient.Dial("https://mainnet.infura.io/v3/9d0db43eb51e46d59967ac5a90b95b4b")
}
