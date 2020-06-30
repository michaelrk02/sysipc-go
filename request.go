package sysipc

type request struct {
    CallId uint64               `json:"call_id"`
    Method string               `json:"method"`
    Args map[string]interface{} `json:"args"`
}

