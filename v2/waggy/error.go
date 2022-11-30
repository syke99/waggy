package waggy

type WaggyError struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Status   int    `json:"status"`
	Instance string `json:"instance"`
	Field    string `json:"field"`
}
