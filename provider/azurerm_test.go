package provider_test

import (
	"strings"
	"testing"

	"github.com/ms-henglu/pal/provider"
	"github.com/ms-henglu/pal/rawlog"
	"github.com/ms-henglu/pal/types"
)

func TestAzureRMProvider_IsRequestTrace(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want bool
	}{
		{
			name: "azapi request trace",
			log:  "2023-04-28T13:13:16.092+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:13:16.092051 Request: ==> OUTGOING REQUEST (Try=1)",
			want: false,
		},
		{
			name: "azapi response trace",
			log:  "2023-04-28T13:13:18.304+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:13:18.304139 Response: ==> REQUEST/RESPONSE (Try=1/2.211800791s, OpTime=2.211881666s) -- RESPONSE RECEIVED",
			want: false,
		},
		{
			name: "azurerm request trace",
			log:  "2023-04-28T13:12:48.330+0800 [DEBUG] provider.terraform-provider-azurerm: AzureRM Request:",
			want: true,
		},
		{
			name: "azurerm request trace",
			log:  "2023-04-28T13:12:48.330+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Request:",
			want: true,
		},
		{
			name: "azurerm response trace",
			log:  "2023-04-28T13:12:48.908+0800 [DEBUG] provider.terraform-provider-azurerm: AzureRM Response for https://management.azure.com/subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01:",
			want: false,
		},
		{
			name: "azuread request trace",
			log:  "2023-04-14T15:14:53.530+0800 [INFO]  provider.terraform-provider-azuread: 2023/04/14 15:14:53 [DEBUG] ============================ Begin AzureAD Request ============================",
			want: false,
		},
		{
			name: "azuread response trace",
			log:  "2023-04-14T15:14:54.084+0800 [INFO]  provider.terraform-provider-azuread: 2023/04/14 15:14:54 [DEBUG] ============================ Begin AzureAD Response ===========================",
			want: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := rawlog.NewRawLog(tc.log)
			if err != nil {
				t.Fatalf("failed to parse log: %v", err)
			}
			got := provider.AzureRMProvider{}.IsRequestTrace(*l)
			if got != tc.want {
				t.Errorf("IsRequestTrace() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAzureRMProvider_IsResponseTrace(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want bool
	}{
		{
			name: "azapi request trace",
			log:  "2023-04-28T13:13:16.092+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:13:16.092051 Request: ==> OUTGOING REQUEST (Try=1)",
			want: false,
		},
		{
			name: "azapi response trace",
			log:  "2023-04-28T13:13:18.304+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:13:18.304139 Response: ==> REQUEST/RESPONSE (Try=1/2.211800791s, OpTime=2.211881666s) -- RESPONSE RECEIVED",
			want: false,
		},
		{
			name: "azurerm request trace",
			log:  "2023-04-28T13:12:48.330+0800 [DEBUG] provider.terraform-provider-azurerm: AzureRM Request:",
			want: false,
		},
		{
			name: "azurerm response trace",
			log:  "2023-04-28T13:12:48.908+0800 [DEBUG] provider.terraform-provider-azurerm: AzureRM Response for https://management.azure.com/subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01:",
			want: true,
		},
		{
			name: "azurerm response trace",
			log:  "2023-04-28T13:12:48.908+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Response for https://management.azure.com/subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01:",
			want: true,
		},
		{
			name: "azuread request trace",
			log:  "2023-04-14T15:14:53.530+0800 [INFO]  provider.terraform-provider-azuread: 2023/04/14 15:14:53 [DEBUG] ============================ Begin AzureAD Request ============================",
			want: false,
		},
		{
			name: "azuread response trace",
			log:  "2023-04-14T15:14:54.084+0800 [INFO]  provider.terraform-provider-azuread: 2023/04/14 15:14:54 [DEBUG] ============================ Begin AzureAD Response ===========================",
			want: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := rawlog.NewRawLog(tc.log)
			if err != nil {
				t.Fatalf("failed to parse log: %v", err)
			}
			got := provider.AzureRMProvider{}.IsResponseTrace(*l)
			if got != tc.want {
				t.Errorf("IsResponseTrace() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAzureRMProvider_ParseRequest(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want types.RequestTrace
	}{
		{
			name: "azurerm GET request trace",
			log: `2023-06-15T14:34:07.601+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Request: 
GET /subscriptions/******/resourcegroups/henglu615aa?api-version=2020-06-01 HTTP/1.1
Host: management.azure.com
User-Agent: Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
X-Ms-Correlation-Request-Id: 50a65a30-0b6a-0c7e-bd32-388facfa615d
Accept-Encoding: gzip: timestamp=2023-06-15T14:34:07.600+0800`,
			want: types.RequestTrace{
				Provider: "azurerm",
				Method:   "GET",
				Host:     "management.azure.com",
				Url:      "/subscriptions/******/resourcegroups/henglu615aa?api-version=2020-06-01",
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Host":                        "management.azure.com",
						"User-Agent":                  "Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"X-Ms-Correlation-Request-Id": "50a65a30-0b6a-0c7e-bd32-388facfa615d",
						"Accept-Encoding":             "gzip",
					},
				},
			},
		},
		{
			name: "azurerm PUT request trace",
			log: `2023-06-15T14:34:08.031+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Request: 
PUT /subscriptions/******/resourcegroups/henglu615aa?api-version=2020-06-01 HTTP/1.1
Host: management.azure.com
User-Agent: Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
Content-Length: 35
Content-Type: application/json; charset=utf-8
X-Ms-Correlation-Request-Id: 50a65a30-0b6a-0c7e-bd32-388facfa615d
Accept-Encoding: gzip

{"location":"westeurope","tags":{}}: timestamp=2023-06-15T14:34:08.031+0800`,
			want: types.RequestTrace{
				Provider: "azurerm",
				Method:   "PUT",
				Host:     "management.azure.com",
				Url:      "/subscriptions/******/resourcegroups/henglu615aa?api-version=2020-06-01",
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Host":                        "management.azure.com",
						"User-Agent":                  "Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"Content-Length":              "35",
						"Content-Type":                "application/json; charset=utf-8",
						"X-Ms-Correlation-Request-Id": "50a65a30-0b6a-0c7e-bd32-388facfa615d",
						"Accept-Encoding":             "gzip",
					},
					Body: `{"location":"westeurope","tags":{}}`,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := rawlog.NewRawLog(tc.log)
			if err != nil {
				t.Fatalf("failed to parse log: %v", err)
			}
			got, err := provider.AzureRMProvider{}.ParseRequest(*l)
			if err != nil {
				t.Fatalf("failed to parse request: %v", err)
			}
			if got.Host != tc.want.Host {
				t.Errorf("ParseRequest() host = %v, want %v", got.Host, tc.want.Host)
			}
			if got.Method != tc.want.Method {
				t.Errorf("ParseRequest() method = %v, want %v", got.Method, tc.want.Method)
			}
			if got.Url != tc.want.Url {
				t.Errorf("ParseRequest() url = %v, want %v", got.Url, tc.want.Url)
			}
			if got.Request == nil {
				t.Errorf("ParseRequest() request = nil, want not nil")
			} else {
				if len(got.Request.Headers) != len(tc.want.Request.Headers) {
					t.Errorf("ParseRequest() request headers length = %v, want %v", len(got.Request.Headers), len(tc.want.Request.Headers))
				}
				for k, v := range got.Request.Headers {
					if v != tc.want.Request.Headers[k] {
						t.Errorf("ParseRequest() request header %v = %v, want %v", k, v, tc.want.Request.Headers[k])
					}
				}
				if got.Request.Body != tc.want.Request.Body {
					t.Errorf("ParseRequest() request body = %v, want %v", got.Request.Body, tc.want.Request.Body)
				}
			}

		})
	}
}

func TestAzureRMProvider_ParseResponse(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want types.RequestTrace
	}{
		{
			name: "azurerm GET response trace",
			log: `2023-06-15T15:10:42.112+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Response for https://management.azure.com/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourcegroups/henglu615aa?api-version=2020-06-01: 
HTTP/2.0 200 OK
Cache-Control: no-cache
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Jun 2023 07:10:42 GMT
Expires: -1
Pragma: no-cache
Strict-Transport-Security: max-age=31536000; includeSubDomains
Vary: Accept-Encoding
X-Content-Type-Options: nosniff
X-Ms-Correlation-Request-Id: 64f52fb8-5580-58c0-a428-6220bd7811df
X-Ms-Ratelimit-Remaining-Subscription-Reads: 11999
X-Ms-Request-Id: c11a4994-287b-4ea7-ac13-ccf635682936
X-Ms-Routing-Request-Id: SOUTHEASTASIA:20230615T071042Z:c11a4994-287b-4ea7-ac13-ccf635682936

{"id":"/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}: timestamp=2023-06-15T15:10:42.112+0800`,
			want: types.RequestTrace{
				Provider:   "azurerm",
				Method:     "",
				StatusCode: 200,
				Host:       "management.azure.com",
				Url:        "/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourcegroups/henglu615aa?api-version=2020-06-01",
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Cache-Control":               "no-cache",
						"Content-Type":                "application/json; charset=utf-8",
						"Date":                        "Thu, 15 Jun 2023 07:10:42 GMT",
						"Expires":                     "-1",
						"Pragma":                      "no-cache",
						"Strict-Transport-Security":   "max-age=31536000; includeSubDomains",
						"Vary":                        "Accept-Encoding",
						"X-Content-Type-Options":      "nosniff",
						"X-Ms-Correlation-Request-Id": "64f52fb8-5580-58c0-a428-6220bd7811df",
						"X-Ms-Ratelimit-Remaining-Subscription-Reads": "11999",
						"X-Ms-Request-Id":         "c11a4994-287b-4ea7-ac13-ccf635682936",
						"X-Ms-Routing-Request-Id": "SOUTHEASTASIA:20230615T071042Z:c11a4994-287b-4ea7-ac13-ccf635682936",
					},
					Body: `{"id":"/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}`,
				},
			},
		},
		{
			name: "azapi PUT response trace",
			log: `2023-06-15T15:12:41.772+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Response for https://management.azure.com/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourcegroups/henglu615aa?api-version=2020-06-01: 
HTTP/2.0 201 Created
Content-Length: 233
Cache-Control: no-cache
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Jun 2023 07:12:40 GMT
Expires: -1
Pragma: no-cache
Strict-Transport-Security: max-age=31536000; includeSubDomains
X-Content-Type-Options: nosniff
X-Ms-Correlation-Request-Id: 67bba740-9525-5e2b-1431-0f50ea61b961
X-Ms-Ratelimit-Remaining-Subscription-Writes: 1199
X-Ms-Request-Id: 99b0be99-5635-4044-9961-06772eb84bdc
X-Ms-Routing-Request-Id: SOUTHEASTASIA:20230615T071241Z:99b0be99-5635-4044-9961-06772eb84bdc

{"id":"/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}: timestamp=2023-06-15T15:12:41.772+0800
`,
			want: types.RequestTrace{
				Provider:   "azurerm",
				StatusCode: 201,
				Host:       "management.azure.com",
				Url:        "/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourcegroups/henglu615aa?api-version=2020-06-01",
				Method:     "",
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Content-Length":              "233",
						"Cache-Control":               "no-cache",
						"Content-Type":                "application/json; charset=utf-8",
						"Date":                        "Thu, 15 Jun 2023 07:12:40 GMT",
						"Expires":                     "-1",
						"Pragma":                      "no-cache",
						"Strict-Transport-Security":   "max-age=31536000; includeSubDomains",
						"X-Content-Type-Options":      "nosniff",
						"X-Ms-Correlation-Request-Id": "67bba740-9525-5e2b-1431-0f50ea61b961",
						"X-Ms-Ratelimit-Remaining-Subscription-Writes": "1199",
						"X-Ms-Request-Id":         "99b0be99-5635-4044-9961-06772eb84bdc",
						"X-Ms-Routing-Request-Id": "SOUTHEASTASIA:20230615T071241Z:99b0be99-5635-4044-9961-06772eb84bdc",
					},
					Body: `{"id":"/subscriptions/85b3dbca-5974-4067-9669-67a141095a76/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}`,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := rawlog.NewRawLog(tc.log)
			if err != nil {
				t.Fatalf("failed to parse log: %v", err)
			}
			got, err := provider.AzureRMProvider{}.ParseResponse(*l)
			if err != nil {
				t.Fatalf("ParseResponse() error = %v", err)
			}
			if got.Host != tc.want.Host {
				t.Errorf("ParseResponse() host = %v, want %v", got.Host, tc.want.Host)
			}
			if got.Method != tc.want.Method {
				t.Errorf("ParseResponse() method = %v, want %v", got.Method, tc.want.Method)
			}
			if got.Url != tc.want.Url {
				t.Errorf("ParseResponse() url = %v, want %v", got.Url, tc.want.Url)
			}
			if got.Response == nil {
				t.Errorf("ParseResponse() response = nil, want not nil")
			} else {
				if len(got.Response.Headers) != len(tc.want.Response.Headers) {
					t.Errorf("ParseResponse() response headers length = %v, want %v", len(got.Response.Headers), len(tc.want.Response.Headers))
				}
				for k, v := range got.Response.Headers {
					if v != tc.want.Response.Headers[k] {
						t.Errorf("ParseResponse() response header %v = %v, want %v", k, v, tc.want.Response.Headers[k])
					}
				}
				if strings.Trim(got.Response.Body, " \n") != tc.want.Response.Body {
					t.Errorf("ParseResponse() response body = %v, want %v", got.Response.Body, tc.want.Response.Body)
				}
			}

		})
	}
}
