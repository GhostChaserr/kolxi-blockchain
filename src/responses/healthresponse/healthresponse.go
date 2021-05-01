package healthresponse

type HealthResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
