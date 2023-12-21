package helpers

func SuccessResponse(message string) map[string]interface{} {
	return map[string]any{
		"status" : "success",
		"message" : message,
	}
}

func SuccessWithDataResponse(message string, data any) map[string]interface{} {
	return map[string]any{
		"status" : "success",
		"message" : message,
		"data" : data,
	}
}

func FailedResponse(message string) map[string]interface{} {
	return map[string]any{
		"status" : "failed",
		"message" : message,
	}
}