package response

type ErrorResponse struct {
    Error string `json:"error"`
}

// 共通エラーレスポンス
func Error(err error) ErrorResponse {
    return ErrorResponse{Error: err.Error()}
}