package sysipc

type response struct {
    CallId uint64   `json:"call_id"`
    Ret interface{} `json:"return"`
    Err string      `json:"error"`
}

