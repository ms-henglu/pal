package types

import (
	"time"
)

type RawLog struct {
	TimeStamp time.Time
	Level     string
	Message   string
}

type RequestTrace struct {
	Url        string
	Method     string
	Host       string
	StatusCode int
	Provider   string
	TimeStamp  time.Time
	Request    *HttpRequest
	Response   *HttpResponse
}

type HttpRequest struct {
	Headers map[string]string
	Body    string
}

type HttpResponse struct {
	Headers map[string]string
	Body    string
}
