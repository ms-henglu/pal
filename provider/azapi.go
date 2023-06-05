package provider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ms-henglu/pal/types"
	"github.com/ms-henglu/pal/utils"
)

var _ types.Provider = AzAPIProvider{}

type AzAPIProvider struct {
}

func (a AzAPIProvider) IsRequestTrace(l types.RawLog) bool {
	return l.Level == "DEBUG" && strings.Contains(l.Message, "Request: ==> OUTGOING REQUEST")
}

func (a AzAPIProvider) IsResponseTrace(l types.RawLog) bool {
	return l.Level == "DEBUG" && strings.Contains(l.Message, "Response: ==> REQUEST/RESPONSE")
}

func (a AzAPIProvider) ParseRequest(l types.RawLog) (*types.RequestTrace, error) {
	method := ""
	host := ""
	uriPath := ""
	body := ""
	headers := make(map[string]string)
	for _, line := range strings.Split(l.Message, "\n") {
		line = strings.Trim(line, " ")
		switch {
		case line == "" || strings.Contains(line, "==>") ||
			strings.Contains(line, "Request contained no body") ||
			strings.Contains(line, "-----"):
			continue
		case strings.Contains(line, ": "):
			key, value, err := utils.ParseHeader(line)
			if err != nil {
				return nil, err
			}
			utils.AppendHeader(headers, key, value)
		case utils.IsJson(line):
			body = line
		default:
			if parts := strings.Split(line, " "); len(parts) == 2 {
				method = parts[0]
				parsedUrl, err := url.Parse(parts[1])
				if err == nil {
					host = parsedUrl.Host
					uriPath = fmt.Sprintf("%s?%s", parsedUrl.Path, parsedUrl.RawQuery)
				}
			}
		}
	}
	return &types.RequestTrace{
		TimeStamp: l.TimeStamp,
		Method:    method,
		Host:      host,
		Url:       uriPath,
		Provider:  "azapi",
		Request: &types.HttpRequest{
			Headers: headers,
			Body:    body,
		},
	}, nil
}

func (a AzAPIProvider) ParseResponse(l types.RawLog) (*types.RequestTrace, error) {
	method := ""
	host := ""
	uriPath := ""
	body := ""
	headers := make(map[string]string)

	lines := strings.Split(l.Message, strings.Repeat("-", 80))
	message := l.Message
	if len(lines) == 4 {
		body = lines[2]
		message = lines[0] + lines[1]
	}

	for _, line := range strings.Split(message, "\n") {
		line = strings.Trim(line, " ")
		switch {
		case line == "" || strings.Contains(line, "==>") ||
			strings.Contains(line, "contained no body") ||
			strings.Contains(line, "-----"):
			continue
		case strings.Contains(line, ": "):
			key, value, err := utils.ParseHeader(line)
			if err != nil {
				return nil, err
			}
			utils.AppendHeader(headers, key, value)
		case utils.IsJson(line):
			body = line
		default:
			if parts := strings.Split(line, " "); len(parts) == 2 {
				method = parts[0]
				parsedUrl, err := url.Parse(parts[1])
				if err == nil {
					host = parsedUrl.Host
					uriPath = fmt.Sprintf("%s?%s", parsedUrl.Path, parsedUrl.RawQuery)
				}
			}
		}
	}

	statusCode := 0
	if v := headers["RESPONSE Status"]; v != "" {
		fmt.Sscanf(v, "%d", &statusCode)
	}
	return &types.RequestTrace{
		TimeStamp:  l.TimeStamp,
		Method:     method,
		Host:       host,
		Url:        uriPath,
		StatusCode: statusCode,
		Provider:   "azapi",
		Response: &types.HttpResponse{
			Headers: headers,
			Body:    body,
		},
	}, nil
}
