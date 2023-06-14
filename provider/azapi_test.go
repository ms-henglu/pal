package provider_test

import (
	"strings"
	"testing"

	"github.com/ms-henglu/pal/provider"
	"github.com/ms-henglu/pal/rawlog"
	"github.com/ms-henglu/pal/types"
)

func TestAzAPIProvider_IsRequestTrace(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want bool
	}{
		{
			name: "azapi request trace",
			log:  "2023-04-28T13:13:16.092+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:13:16.092051 Request: ==> OUTGOING REQUEST (Try=1)",
			want: true,
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
			got := provider.AzAPIProvider{}.IsRequestTrace(*l)
			if got != tc.want {
				t.Errorf("IsRequestTrace() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAzAPIProvider_IsResponseTrace(t *testing.T) {
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
			want: true,
		},
		{
			name: "azurerm request trace",
			log:  "2023-04-28T13:12:48.330+0800 [DEBUG] provider.terraform-provider-azurerm: AzureRM Request:",
			want: false,
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
			got := provider.AzAPIProvider{}.IsResponseTrace(*l)
			if got != tc.want {
				t.Errorf("IsResponseTrace() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAzAPIProvider_ParseRequest(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want types.RequestTrace
	}{
		{
			name: "azapi GET request trace",
			log: `2023-04-28T13:12:52.862+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:12:52.862113 Request: ==> OUTGOING REQUEST (Try=1)
   GET https://management.azure.com/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview
   Accept: application/json
   Authorization: REDACTED
   User-Agent: HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
   X-Ms-Correlation-Request-Id: 8817767b-435f-9298-42f8-534407f68afb
   Request contained no body: timestamp=2023-04-28T13:12:52.862+0800`,
			want: types.RequestTrace{
				Provider: "azapi",
				Method:   "GET",
				Host:     "management.azure.com",
				Url:      "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview",
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Accept":                      "application/json",
						"Authorization":               "REDACTED",
						"User-Agent":                  "HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"X-Ms-Correlation-Request-Id": "8817767b-435f-9298-42f8-534407f68afb",
					},
				},
			},
		},
		{
			name: "azapi PUT request trace",
			log: `2023-04-28T13:12:54.170+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:12:54.169864 Request: ==> OUTGOING REQUEST (Try=1)
   PUT https://management.azure.com/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview
   Accept: application/json
   Authorization: REDACTED
   Content-Length: 80
   Content-Type: application/json
   User-Agent: HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
   X-Ms-Correlation-Request-Id: 8817767b-435f-9298-42f8-534407f68afb
   --------------------------------------------------------------------------------
{"location":"westeurope","name":"henglu1","properties":{"sku":{"name":"Basic"}}}
   --------------------------------------------------------------------------------: timestamp=2023-04-28T13:12:54.169+0800`,
			want: types.RequestTrace{
				Provider: "azapi",
				Method:   "PUT",
				Host:     "management.azure.com",
				Url:      "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview",
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Accept":                      "application/json",
						"Authorization":               "REDACTED",
						"Content-Length":              "80",
						"Content-Type":                "application/json",
						"User-Agent":                  "HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"X-Ms-Correlation-Request-Id": "8817767b-435f-9298-42f8-534407f68afb",
					},
					Body: `{"location":"westeurope","name":"henglu1","properties":{"sku":{"name":"Basic"}}}`,
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
			got, err := provider.AzAPIProvider{}.ParseRequest(*l)
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

func TestAzAPIProvider_ParseResponse(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want types.RequestTrace
	}{
		{
			name: "azapi GET response trace",
			log: `2023-04-28T13:12:54.167+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:12:54.166603 Response: ==> REQUEST/RESPONSE (Try=1/1.304370833s, OpTime=1.304446958s) -- RESPONSE RECEIVED
   GET https://management.azure.com/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview
   Accept: application/json
   Authorization: REDACTED
   User-Agent: HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
   X-Ms-Correlation-Request-Id: 8817767b-435f-9298-42f8-534407f68afb
   --------------------------------------------------------------------------------
   RESPONSE Status: 404 Not Found
   Cache-Control: no-cache
   Content-Length: 229
   Content-Type: application/json; charset=utf-8
   Date: Fri, 28 Apr 2023 05:12:53 GMT
   Expires: -1
   Pragma: no-cache
   Strict-Transport-Security: max-age=31536000; includeSubDomains
   X-Content-Type-Options: nosniff
   X-Ms-Correlation-Request-Id: 8817767b-435f-9298-42f8-534407f68afb
   X-Ms-Failure-Cause: REDACTED
   X-Ms-Request-Id: e718dd71-7703-4ca0-a461-9cf510071ede
   X-Ms-Routing-Request-Id: SOUTHEASTASIA:20230428T051254Z:e718dd71-7703-4ca0-a461-9cf510071ede
   --------------------------------------------------------------------------------
{"error":{"code":"ResourceNotFound","message":"The Resource 'Microsoft.Automation/automationAccounts/henglu1' under resource group 'henglu1114' was not found. For more details please go to https://aka.ms/ARMResourceNotFoundFix"}}
   --------------------------------------------------------------------------------: timestamp=2023-04-28T13:12:54.166+0800
`,
			want: types.RequestTrace{
				Provider:   "azapi",
				Method:     "GET",
				StatusCode: 404,
				Host:       "management.azure.com",
				Url:        "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview",
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Cache-Control":               "no-cache",
						"Content-Length":              "229",
						"Content-Type":                "application/json; charset=utf-8",
						"Date":                        "Fri, 28 Apr 2023 05:12:53 GMT",
						"Expires":                     "-1",
						"Pragma":                      "no-cache",
						"Strict-Transport-Security":   "max-age=31536000; includeSubDomains",
						"X-Content-Type-Options":      "nosniff",
						"X-Ms-Correlation-Request-Id": "8817767b-435f-9298-42f8-534407f68afb",
						"X-Ms-Failure-Cause":          "REDACTED",
						"X-Ms-Request-Id":             "e718dd71-7703-4ca0-a461-9cf510071ede",
						"X-Ms-Routing-Request-Id":     "SOUTHEASTASIA:20230428T051254Z:e718dd71-7703-4ca0-a461-9cf510071ede",
					},
					Body: `{"error":{"code":"ResourceNotFound","message":"The Resource 'Microsoft.Automation/automationAccounts/henglu1' under resource group 'henglu1114' was not found. For more details please go to https://aka.ms/ARMResourceNotFoundFix"}}`,
				},
			},
		},
		{
			name: "azapi PUT response trace",
			log: `2023-04-28T13:13:00.563+0800 [DEBUG] provider.terraform-provider-azapi: Apr 28 13:13:00.563408 Response: ==> REQUEST/RESPONSE (Try=1/6.393429583s, OpTime=6.393500375s) -- RESPONSE RECEIVED
   PUT https://management.azure.com/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview
   Accept: application/json
   Authorization: REDACTED
   Content-Length: 80
   Content-Type: application/json
   User-Agent: HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
   X-Ms-Correlation-Request-Id: 8817767b-435f-9298-42f8-534407f68afb
   --------------------------------------------------------------------------------
   RESPONSE Status: 201 Created
   Cache-Control: no-cache
   Content-Length: 812
   Content-Type: application/json; charset=utf-8
   Date: Fri, 28 Apr 2023 05:13:00 GMT
   Expires: -1
   Location: REDACTED
   Pragma: no-cache
   Server: Microsoft-HTTPAPI/2.0
   Strict-Transport-Security: max-age=31536000; includeSubDomains
   X-Content-Type-Options: nosniff
   X-Ms-Correlation-Request-Id: 8817767b-435f-9298-42f8-534407f68afb
   X-Ms-Ratelimit-Remaining-Subscription-Writes: 1198
   X-Ms-Request-Id: 5e488609-4d2d-4c51-9efe-ee293d97ddd5
   X-Ms-Routing-Request-Id: SOUTHEASTASIA:20230428T051300Z:8cca4a82-b200-4546-9114-96bd6a3b9dfa
   --------------------------------------------------------------------------------
{"name":"henglu1","id":"/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1","type":"Microsoft.Automation/AutomationAccounts","location":"westeurope","tags":{},"etag":null,"properties":{"sku":{"name":"Basic","family":null,"capacity":null},"state":"Ok","RegistrationUrl":"https://ff3a2b15-73eb-4d82-af7b-b399b3560822.agentsvc.we.azure-automation.net/accounts/ff3a2b15-73eb-4d82-af7b-b399b3560822","encryption":{"keySource":"Microsoft.Automation","identity":{"userAssignedIdentity":null}},"RuntimeConfiguration":{"powershell":{"builtinModules":{"Az":"8.0.0"}},"powershell7":{"builtinModules":{"Az":"8.0.0"}}},"creationTime":"2023-04-28T05:12:58.247+00:00","lastModifiedBy":null,"lastModifiedTime":"2023-04-28T05:12:58.247+00:00"}}
   --------------------------------------------------------------------------------: timestamp=2023-04-28T13:13:00.563+0800
`,
			want: types.RequestTrace{
				Provider:   "azapi",
				StatusCode: 201,
				Host:       "management.azure.com",
				Url:        "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview",
				Method:     "PUT",
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Cache-Control":               "no-cache",
						"Content-Length":              "812",
						"Content-Type":                "application/json; charset=utf-8",
						"Date":                        "Fri, 28 Apr 2023 05:13:00 GMT",
						"Expires":                     "-1",
						"Location":                    "REDACTED",
						"Pragma":                      "no-cache",
						"Server":                      "Microsoft-HTTPAPI/2.0",
						"Strict-Transport-Security":   "max-age=31536000; includeSubDomains",
						"X-Content-Type-Options":      "nosniff",
						"X-Ms-Correlation-Request-Id": "8817767b-435f-9298-42f8-534407f68afb",
						"X-Ms-Ratelimit-Remaining-Subscription-Writes": "1198",
						"X-Ms-Request-Id":         "5e488609-4d2d-4c51-9efe-ee293d97ddd5",
						"X-Ms-Routing-Request-Id": "SOUTHEASTASIA:20230428T051300Z:8cca4a82-b200-4546-9114-96bd6a3b9dfa",
					},
					Body: `{"name":"henglu1","id":"/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1","type":"Microsoft.Automation/AutomationAccounts","location":"westeurope","tags":{},"etag":null,"properties":{"sku":{"name":"Basic","family":null,"capacity":null},"state":"Ok","RegistrationUrl":"https://ff3a2b15-73eb-4d82-af7b-b399b3560822.agentsvc.we.azure-automation.net/accounts/ff3a2b15-73eb-4d82-af7b-b399b3560822","encryption":{"keySource":"Microsoft.Automation","identity":{"userAssignedIdentity":null}},"RuntimeConfiguration":{"powershell":{"builtinModules":{"Az":"8.0.0"}},"powershell7":{"builtinModules":{"Az":"8.0.0"}}},"creationTime":"2023-04-28T05:12:58.247+00:00","lastModifiedBy":null,"lastModifiedTime":"2023-04-28T05:12:58.247+00:00"}}`,
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
			got, err := provider.AzAPIProvider{}.ParseResponse(*l)
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
