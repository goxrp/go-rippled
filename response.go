package gorippled

const (
	StatusError   = "error"
	StatusSuccess = "success"
)

type JsonRpcResponseError struct {
	Result JsonRpcResponseErrorResult `json:"result"`
}

type JsonRpcResponseErrorResult struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}
