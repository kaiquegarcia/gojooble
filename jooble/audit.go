package jooble

type Audit interface {
	PayloadSent() string
	ResponseReceived() string
	StatusCode() int
}

type audit struct {
	payload    []byte
	response   []byte
	statusCode int
}

func (a *audit) PayloadSent() string {
	if len(a.payload) == 0 {
		return ""
	}

	return string(a.payload)
}

func (a *audit) ResponseReceived() string {
	if len(a.response) == 0 {
		return ""
	}

	return string(a.response)
}

func (a *audit) StatusCode() int {
	return a.statusCode
}
