package payload

func ResponseSuccess(message string, data interface{}) ResponseSuccessBody {
	return ResponseSuccessBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func ResponseSuccessWithoutData(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
	}
}

func ResponseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
}

func ResponseFailedWithData(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": message,
		"data":    data,
	}
}
