package request

type Request interface {
	GetURL() string
	GetMethod() string
	GetVersion() string
	GetHeaders() map[string]string
}

type BaseRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Version string
}

func (r BaseRequest) GetURL() string {
	return r.URL
}

func (r BaseRequest) GetMethod() string {
	return r.Method
}

func (r BaseRequest) GetVersion() string {
	return r.Version
}

func (r BaseRequest) GetHeaders() map[string]string {
	return r.Headers
}

func (r *BaseRequest) AddHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[key] = value
}
