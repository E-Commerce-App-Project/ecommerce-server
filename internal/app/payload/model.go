package payload

type HealthCheckModel struct {
	Status  string `json:"status" example:"success"`
	Runtime string `json:"runtime" example:"go1.13.4"`
	UpTime  string `json:"uptime" example:"1h"`
	Version string `json:"version" example:"1.0.0"`
} //@name HealthCheckModel
