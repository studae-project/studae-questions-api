package exception

type BusinessException struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
