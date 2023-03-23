package responses

// 1. creating a reusable struct to describe our API’s response

type UserResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}
