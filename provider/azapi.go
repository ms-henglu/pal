package provider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ms-henglu/pal/rawlog"
	"github.com/ms-henglu/pal/types"
	"github.com/ms-henglu/pal/utils"
)

var _ Provider = AzAPIProvider{}

type AzAPIProvider struct {
}

func (a AzAPIProvider) IsRequestTrace(l rawlog.RawLog) bool {
	return l.Level == "DEBUG" && strings.Contains(l.Message, "Request: ==> OUTGOING REQUEST")
}

func (a AzAPIProvider) IsResponseTrace(l rawlog.RawLog) bool {
	return l.Level == "DEBUG" && strings.Contains(l.Message, "Response: ==> REQUEST/RESPONSE")
}

func (a AzAPIProvider) ParseRequest(l rawlog.RawLog) (*types.RequestTrace, error) {
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
			headers[key] = value
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
		Url:       utils.NormalizeUrlPath(uriPath),
		Provider:  "azapi",
		Request: &types.HttpRequest{
			Headers: headers,
			Body:    body,
		},
	}, nil
}

func (a AzAPIProvider) ParseResponse(l rawlog.RawLog) (*types.RequestTrace, error) {
	method := ""
	host := ""
	uriPath := ""
	body := ""
	headers := make(map[string]string)

	sections := strings.Split(l.Message, strings.Repeat("-", 80))
	message := l.Message
	if len(sections) == 4 {
		body = sections[2]
		message = utils.LineAt(sections[0], 1) + sections[1]
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
			headers[key] = value
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
		delete(headers, "RESPONSE Status")
		fmt.Sscanf(v, "%d", &statusCode)
	}
	return &types.RequestTrace{
		TimeStamp:  l.TimeStamp,
		Method:     method,
		Host:       host,
		Url:        utils.NormalizeUrlPath(uriPath),
		StatusCode: statusCode,
		Provider:   "azapi",
		Response: &types.HttpResponse{
			Headers: headers,
			Body:    body,
		},
	}, nil
}
