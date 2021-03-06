package httphoneypot

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Proto   string            `json:"version"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	IP      string            `json:"ip"`
}

type RequestRepository interface {
	Create(req *Request) error
}
