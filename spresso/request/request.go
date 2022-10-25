package request

import (
	"io"
	"net/http"
	"time"
)

type Request struct {
	Config spresso.Config

	AttemptTime            time.Time
	Time                   time.Time
	Operation              *Operation
	HTTPRequest            *http.Request
	HTTPResponse           *http.Response
	Body                   io.ReadSeeker
	streamingBody          io.ReadCloser
	BodyStart              int64
	Params                 interface{}
	Error                  error
	Data                   interface{}
	RequestID              string
	RetryCount             int
	Retryable              *bool
	RetryDelay             time.Duration
	NotHoist               bool
	SignedHeaderVals       http.Header
	LastSignedAt           time.Time
	DisableFollowRedirects bool

	ExpireTime time.Duration

	context spresso.Context

	built bool

	safeBody *offsetReader
}
