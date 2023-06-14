package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/ms-henglu/pal/types"
)

var _ Formatter = OavTrafficFormatter{}

type OavTrafficFormatter struct {
}

type OavTraffic struct {
	LiveRequest  LiveRequest  `json:"liveRequest"`
	LiveResponse LiveResponse `json:"liveResponse"`
}

type LiveRequest struct {
	Headers map[string]string `json:"headers"`
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Body    interface{}       `json:"body"`
}

type LiveResponse struct {
	StatusCode string            `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
}

func (o OavTrafficFormatter) Format(r types.RequestTrace) string {
	var requestBody interface{}
	err := json.Unmarshal([]byte(r.Request.Body), &requestBody)
	if err != nil {
		requestBody = nil
	}
	var responseBody interface{}
	err = json.Unmarshal([]byte(r.Response.Body), &responseBody)
	if err != nil {
		responseBody = nil
	}

	out := OavTraffic{
		LiveRequest: LiveRequest{
			Headers: r.Request.Headers,
			Method:  r.Method,
			Url:     r.Url,
			Body:    requestBody,
		},
		LiveResponse: LiveResponse{
			StatusCode: fmt.Sprintf("%d", r.StatusCode),
			Headers:    r.Response.Headers,
			Body:       responseBody,
		},
	}

	content, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return ""
	}
	return string(content)
}
