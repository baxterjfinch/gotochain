package renderings

type Response struct {
  Error         string      `json:"error"`
  BlockNumber   uint64      `json:"block_number"`
  HealthCheck   string      `json:"health_check"`
}