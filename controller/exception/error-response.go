package exception

type ErrorResponse struct {
	Message        string  `json:"message"`
	DetailedReason *string `json:"detailed_reason"`
}
