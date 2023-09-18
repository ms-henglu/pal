package provider_test

import (
	"testing"

	"github.com/ms-henglu/pal/provider"
	"github.com/ms-henglu/pal/rawlog"
	"github.com/ms-henglu/pal/types"
)

func TestAzAPIProvider_IsTrafficTrace(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want bool
	}{
		{
			name: "azapi GET request trace",
			log:  "2023-09-15T13:40:31.447+0800 [DEBUG]  provider.terraform-provider-azapi: Live traffic: {}: timestamp=2023-09-15T13:40:31.447+0800",
			want: true,
		},
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
			got := provider.AzAPIProvider{}.IsTrafficTrace(*l)
			if got != tc.want {
				t.Errorf("IsRequestTrace() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAzAPIProvider_ParseTraffic(t *testing.T) {
	testcases := []struct {
		name string
		log  string
		want types.RequestTrace
	}{
		{
			name: "azapi GET request trace",
			log:  `2023-09-15T13:40:31.447+0800 [DEBUG]  provider.terraform-provider-azapi: Live traffic: {"request":{"headers":{"Accept":"application/json","User-Agent":"HashiCorp Terraform/1.5.2 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820","X-Ms-Correlation-Request-Id":"ea90befd-41e6-0f04-a1e5-3be4d56632cc"},"method":"GET","url":"https://management.azure.com/subscriptions/******/resourceGroups/henglu915?api-version=2023-07-01","body":""},"response":{"statusCode":200,"headers":{"Cache-Control":"no-cache","Content-Type":"application/json; charset=utf-8","Date":"Fri, 15 Sep 2023 05:40:30 GMT","Expires":"-1","Pragma":"no-cache","Strict-Transport-Security":"max-age=31536000; includeSubDomains","Vary":"Accept-Encoding","X-Content-Type-Options":"nosniff","X-Ms-Correlation-Request-Id":"ea90befd-41e6-0f04-a1e5-3be4d56632cc","X-Ms-Ratelimit-Remaining-Subscription-Reads":"11999","X-Ms-Request-Id":"04ae0566-5067-4114-816c-1dace25a7b65","X-Ms-Routing-Request-Id":"SOUTHEASTASIA:20230915T054031Z:04ae0566-5067-4114-816c-1dace25a7b65"},"body":"{\"id\":\"/subscriptions/******/resourceGroups/henglu915\",\"name\":\"henglu915\",\"type\":\"Microsoft.Resources/resourceGroups\",\"location\":\"westus\",\"properties\":{\"provisioningState\":\"Succeeded\"}}"}}: timestamp=2023-09-15T13:40:31.447+0800`,
			want: types.RequestTrace{
				Provider:   "azapi",
				Method:     "GET",
				Host:       "management.azure.com",
				Url:        "/subscriptions/******/resourceGroups/henglu915?api-version=2023-07-01",
				StatusCode: 200,
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Accept":                      "application/json",
						"User-Agent":                  "HashiCorp Terraform/1.5.2 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"X-Ms-Correlation-Request-Id": "ea90befd-41e6-0f04-a1e5-3be4d56632cc",
					},
					Body: "",
				},
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Cache-Control":               "no-cache",
						"Content-Type":                "application/json; charset=utf-8",
						"Date":                        "Fri, 15 Sep 2023 05:40:30 GMT",
						"Expires":                     "-1",
						"Pragma":                      "no-cache",
						"Strict-Transport-Security":   "max-age=31536000; includeSubDomains",
						"Vary":                        "Accept-Encoding",
						"X-Content-Type-Options":      "nosniff",
						"X-Ms-Correlation-Request-Id": "ea90befd-41e6-0f04-a1e5-3be4d56632cc",
						"X-Ms-Ratelimit-Remaining-Subscription-Reads": "11999",
						"X-Ms-Request-Id":         "04ae0566-5067-4114-816c-1dace25a7b65",
						"X-Ms-Routing-Request-Id": "SOUTHEASTASIA:20230915T054031Z:04ae0566-5067-4114-816c-1dace25a7b65",
					},
					Body: "{\"id\":\"/subscriptions/******/resourceGroups/henglu915\",\"name\":\"henglu915\",\"type\":\"Microsoft.Resources/resourceGroups\",\"location\":\"westus\",\"properties\":{\"provisioningState\":\"Succeeded\"}}",
				},
			},
		},
		{
			name: "azapi PUT request trace",
			log:  `2023-09-15T14:33:49.909+0800 [DEBUG] provider.terraform-provider-azapi: Live traffic: {"request":{"headers":{"Accept":"application/json","Content-Length":"40","Content-Type":"application/json","User-Agent":"HashiCorp Terraform/1.5.2 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820","X-Ms-Correlation-Request-Id":"6b6ed2c7-055a-3986-5d5f-4c04a4b78193"},"method":"PUT","url":"https://management.azure.com/subscriptions/******/resourceGroups/henglu915?api-version=2023-07-01","body":"{\"location\":\"westus\",\"name\":\"henglu915\"}"},"response":{"statusCode":201,"headers":{"Cache-Control":"no-cache","Content-Length":"215","Content-Type":"application/json; charset=utf-8","Date":"Fri, 15 Sep 2023 06:33:49 GMT","Expires":"-1","Pragma":"no-cache","Strict-Transport-Security":"max-age=31536000; includeSubDomains","X-Content-Type-Options":"nosniff","X-Ms-Correlation-Request-Id":"6b6ed2c7-055a-3986-5d5f-4c04a4b78193","X-Ms-Ratelimit-Remaining-Subscription-Writes":"1199","X-Ms-Request-Id":"e1eef6e2-b814-4f69-82b6-4c117d33cc13","X-Ms-Routing-Request-Id":"SOUTHEASTASIA:20230915T063350Z:e1eef6e2-b814-4f69-82b6-4c117d33cc13"},"body":"{\"id\":\"/subscriptions/******/resourceGroups/henglu915\",\"name\":\"henglu915\",\"type\":\"Microsoft.Resources/resourceGroups\",\"location\":\"westus\",\"properties\":{\"provisioningState\":\"Succeeded\"}}"}}: timestamp=2023-09-15T14:33:49.908+0800`,
			want: types.RequestTrace{
				Provider:   "azapi",
				Method:     "PUT",
				Host:       "management.azure.com",
				Url:        "/subscriptions/******/resourceGroups/henglu915?api-version=2023-07-01",
				StatusCode: 201,
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Accept":                      "application/json",
						"Content-Length":              "40",
						"Content-Type":                "application/json",
						"User-Agent":                  "HashiCorp Terraform/1.5.2 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"X-Ms-Correlation-Request-Id": "6b6ed2c7-055a-3986-5d5f-4c04a4b78193",
					},
					Body: "{\"location\":\"westus\",\"name\":\"henglu915\"}",
				},
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Cache-Control":               "no-cache",
						"Content-Length":              "215",
						"Content-Type":                "application/json; charset=utf-8",
						"Date":                        "Fri, 15 Sep 2023 06:33:49 GMT",
						"Expires":                     "-1",
						"Pragma":                      "no-cache",
						"Strict-Transport-Security":   "max-age=31536000; includeSubDomains",
						"X-Content-Type-Options":      "nosniff",
						"X-Ms-Correlation-Request-Id": "6b6ed2c7-055a-3986-5d5f-4c04a4b78193",
						"X-Ms-Ratelimit-Remaining-Subscription-Writes": "1199",
						"X-Ms-Request-Id":         "e1eef6e2-b814-4f69-82b6-4c117d33cc13",
						"X-Ms-Routing-Request-Id": "SOUTHEASTASIA:20230915T063350Z:e1eef6e2-b814-4f69-82b6-4c117d33cc13",
					},
					Body: "{\"id\":\"/subscriptions/******/resourceGroups/henglu915\",\"name\":\"henglu915\",\"type\":\"Microsoft.Resources/resourceGroups\",\"location\":\"westus\",\"properties\":{\"provisioningState\":\"Succeeded\"}}",
				},
			},
		},
	}

	for _, tc := range testcases {
		l, err := rawlog.NewRawLog(tc.log)
		if err != nil {
			t.Fatalf("failed to parse log: %v", err)
		}
		got, err := provider.AzAPIProvider{}.ParseTraffic(*l)
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
		if got.StatusCode != tc.want.StatusCode {
			t.Errorf("ParseRequest() status code = %v, want %v", got.StatusCode, tc.want.StatusCode)
		}
		if got.Request == nil {
			t.Errorf("ParseRequest() request is nil")
			continue
		}
		if got.Response == nil {
			t.Errorf("ParseRequest() response is nil")
			continue
		}
		if got.Request.Body != tc.want.Request.Body {
			t.Errorf("ParseRequest() request body = %v, want %v", got.Request.Body, tc.want.Request.Body)
		}
		if got.Response.Body != tc.want.Response.Body {
			t.Errorf("ParseRequest() response body = %v, want %v", got.Response.Body, tc.want.Response.Body)
		}
		if len(got.Request.Headers) != len(tc.want.Request.Headers) {
			t.Errorf("ParseRequest() request headers length = %v, want %v", len(got.Request.Headers), len(tc.want.Request.Headers))
			continue
		}
		for k, v := range got.Request.Headers {
			if tc.want.Request.Headers[k] != v {
				t.Errorf("ParseRequest() request header %v = %v, want %v", k, v, tc.want.Request.Headers[k])
			}
		}
		if len(got.Response.Headers) != len(tc.want.Response.Headers) {
			t.Errorf("ParseRequest() response headers length = %v, want %v", len(got.Response.Headers), len(tc.want.Response.Headers))
			continue
		}
		for k, v := range got.Response.Headers {
			if tc.want.Response.Headers[k] != v {
				t.Errorf("ParseRequest() response header %v = %v, want %v", k, v, tc.want.Response.Headers[k])
			}
		}
	}
}
