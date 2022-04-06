package payload

type ResponseSuccessBody struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
} //@name ResponseSuccess

type ResponseSuccessWithoutDataBody struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"success"`
} //@name ResponseSuccessWithoutData

type ResponseFailedBody struct {
	Status  string `json:"status" example:"failed"`
	Message string `json:"message" example:"failed"`
	Data    string `json:"data"`
} //@name ResponseFailed

type HealthCheckResponseSuccessBody struct {
	ResponseSuccessWithoutDataBody
	Data HealthCheckModel `json:"data"`
} //@name HealthCheckResponseSuccess

type AuthResponseSuccessBody struct {
	ResponseSuccessWithoutDataBody
	Data AuthModel `json:"data"`
} //@name AuthResponseSuccess
