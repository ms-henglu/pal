package provider_test

import (
	"encoding/json"
	"reflect"
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
			name: "azurerm PUT request trace",
			log: `2023/07/28 14:03:39 [DEBUG] AzureRM Request: 
PUT /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acctestSecRG-230728140330751068?api-version=2020-06-01 HTTP/1.1
Host: management.azure.com
User-Agent: Go/go1.20.6 (amd64-windows) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.5.3 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820
Content-Length: 35
Content-Type: application/json; charset=utf-8
X-Ms-Correlation-Request-Id: 61f38736-e811-d8b4-befd-c3f5a39c3f35
Accept-Encoding: gzip

{"location":"westeurope","tags":{}}`,
			want: types.RequestTrace{
				Provider: "azurerm",
				Method:   "PUT",
				Host:     "management.azure.com",
				Url:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acctestSecRG-230728140330751068?api-version=2020-06-01",
				Request: &types.HttpRequest{
					Headers: map[string]string{
						"Host":                        "management.azure.com",
						"User-Agent":                  "Go/go1.20.6 (amd64-windows) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.5.3 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820",
						"Content-Length":              "35",
						"Content-Type":                "application/json; charset=utf-8",
						"X-Ms-Correlation-Request-Id": "61f38736-e811-d8b4-befd-c3f5a39c3f35",
						"Accept-Encoding":             "gzip",
					},
					Body: `{"location":"westeurope","tags":{}}`,
				},
			},
		},
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

				if len(tc.want.Request.Body) == 0 {
					if len(got.Request.Body) != 0 {
						t.Errorf("ParseRequest() request body = %v, want %v", got.Request.Body, tc.want.Request.Body)
					}
					return
				}

				var gotBody, wantBody interface{}
				err = json.Unmarshal([]byte(got.Request.Body), &gotBody)
				if err != nil {
					t.Errorf("ParseRequest() request body unmarshal error = %v", err)
				}
				err = json.Unmarshal([]byte(tc.want.Request.Body), &wantBody)
				if err != nil {
					t.Errorf("ParseRequest() request body unmarshal error = %v", err)
				}

				if !reflect.DeepEqual(gotBody, wantBody) {
					t.Errorf("ParseRequest() request body = %v, want %v", gotBody, wantBody)
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
			log: `2023/07/28 14:03:36 [DEBUG] AzureRM Response for https://management.azure.com//providers/Microsoft.Authorization/roleDefinitions/fbdf93bf-df7d-467e-a4d2-9458aa1360c8?api-version=2018-01-01-preview: 
HTTP/2.0 200 OK
Cache-Control: no-cache
Content-Type: application/json; charset=utf-8
Date: Fri, 28 Jul 2023 06:03:36 GMT
Expires: -1
Pragma: no-cache
Set-Cookie: x-ms-gateway-slice=Production; path=/; secure; samesite=none; httponly
Strict-Transport-Security: max-age=31536000; includeSubDomains
Vary: Accept-Encoding
X-Content-Type-Options: nosniff
X-Ms-Correlation-Request-Id: 61f38736-e811-d8b4-befd-c3f5a39c3f35
X-Ms-Ratelimit-Remaining-Tenant-Reads: 11999
X-Ms-Request-Id: 1bd7e400-947e-482b-9ad5-cc7f7265a512
X-Ms-Routing-Request-Id: JAPANEAST:20230728T060336Z:040e3cf8-3c01-461c-9ce0-c27f46037f58

{"properties":{"roleName":"Cosmos DB Account Reader Role","type":"BuiltInRole","description":"Can read Azure Cosmos DB Accounts data","assignableScopes":["/"],"permissions":[{"actions":["Microsoft.Authorization/*/read","Microsoft.DocumentDB/*/read","Microsoft.DocumentDB/databaseAccounts/readonlykeys/action","Microsoft.Insights/MetricDefinitions/read","Microsoft.Insights/Metrics/read","Microsoft.Resources/subscriptions/resourceGroups/read","Microsoft.Support/*"],"notActions":[],"dataActions":[],"notDataActions":[]}],"createdOn":"2017-10-30T17:53:54.6005577Z","updatedOn":"2021-11-11T20:13:28.7911765Z","createdBy":null,"updatedBy":null},"id":"/providers/Microsoft.Authorization/roleDefinitions/fbdf93bf-df7d-467e-a4d2-9458aa1360c8","type":"Microsoft.Authorization/roleDefinitions","name":"fbdf93bf-df7d-467e-a4d2-9458aa1360c8"}`,
			want: types.RequestTrace{
				Provider:   "azurerm",
				Method:     "",
				StatusCode: 200,
				Host:       "management.azure.com",
				Url:        "/providers/Microsoft.Authorization/roleDefinitions/fbdf93bf-df7d-467e-a4d2-9458aa1360c8?api-version=2018-01-01-preview",
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Cache-Control":                         "no-cache",
						"Content-Type":                          "application/json; charset=utf-8",
						"Date":                                  "Fri, 28 Jul 2023 06:03:36 GMT",
						"Expires":                               "-1",
						"Pragma":                                "no-cache",
						"Set-Cookie":                            "x-ms-gateway-slice=Production; path=/; secure; samesite=none; httponly",
						"Strict-Transport-Security":             "max-age=31536000; includeSubDomains",
						"Vary":                                  "Accept-Encoding",
						"X-Content-Type-Options":                "nosniff",
						"X-Ms-Correlation-Request-Id":           "61f38736-e811-d8b4-befd-c3f5a39c3f35",
						"X-Ms-Ratelimit-Remaining-Tenant-Reads": "11999",
						"X-Ms-Request-Id":                       "1bd7e400-947e-482b-9ad5-cc7f7265a512",
						"X-Ms-Routing-Request-Id":               "JAPANEAST:20230728T060336Z:040e3cf8-3c01-461c-9ce0-c27f46037f58",
					},
					Body: `{"properties":{"roleName":"Cosmos DB Account Reader Role","type":"BuiltInRole","description":"Can read Azure Cosmos DB Accounts data","assignableScopes":["/"],"permissions":[{"actions":["Microsoft.Authorization/*/read","Microsoft.DocumentDB/*/read","Microsoft.DocumentDB/databaseAccounts/readonlykeys/action","Microsoft.Insights/MetricDefinitions/read","Microsoft.Insights/Metrics/read","Microsoft.Resources/subscriptions/resourceGroups/read","Microsoft.Support/*"],"notActions":[],"dataActions":[],"notDataActions":[]}],"createdOn":"2017-10-30T17:53:54.6005577Z","updatedOn":"2021-11-11T20:13:28.7911765Z","createdBy":null,"updatedBy":null},"id":"/providers/Microsoft.Authorization/roleDefinitions/fbdf93bf-df7d-467e-a4d2-9458aa1360c8","type":"Microsoft.Authorization/roleDefinitions","name":"fbdf93bf-df7d-467e-a4d2-9458aa1360c8"}`,
				},
			},
		},
		{
			name: "azurerm PUT response trace",
			log: `2023-07-11T16:24:16.674+0800 [DEBUG] provider.terraform-provider-azurerm: AzureRM Response for https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu6421/providers/Microsoft.Insights/components/henglu9128?api-version=2020-02-02: 
HTTP/2.0 200 OK
Access-Control-Expose-Headers: Request-Context
Cache-Control: no-cache
Content-Type: application/json; charset=utf-8
Date: Tue, 11 Jul 2023 08:24:16 GMT
Expires: -1
Pragma: no-cache
Request-Context: appId=cid-v1:60b64f55-e716-48c8-8b96-83eb9c6a7a9b
Server: Microsoft-IIS/10.0
Strict-Transport-Security: max-age=31536000; includeSubDomains
Vary: Accept-Encoding
X-Content-Type-Options: nosniff
X-Ms-Correlation-Request-Id: 14e77fe7-e1a1-b8a1-65d5-40363cac85f5
X-Ms-Ratelimit-Remaining-Subscription-Writes: 1199
X-Ms-Request-Id: aa9250b0-9955-4821-ad66-ff354a68d834
X-Ms-Routing-Request-Id: SOUTHEASTASIA:20230711T082416Z:aa9250b0-9955-4821-ad66-ff354a68d834
X-Powered-By: ASP.NET

{
  "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu6421/providers/microsoft.insights/components/henglu9128",
  "name": "henglu9128",
  "type": "microsoft.insights/components",
  "location": "westeurope",
  "tags": {},
  "kind": "web",
  "etag": "\"920312e9-0000-0200-0000-64ad11ae0000\"",
  "properties": {
    "ApplicationId": "henglu9128",
    "AppId": "76d8a56f-5069-42ea-bae2-645248170408",
    "Application_Type": "web",
    "Flow_Type": null,
    "Request_Source": null,
    "InstrumentationKey": "2d727e28-a681-48af-95aa-677daba464d6",
    "ConnectionString": "InstrumentationKey=2d727e28-a681-48af-95aa-677daba464d6;IngestionEndpoint=https://westeurope-4.in.applicationinsights.azure.com/;LiveEndpoint=https://westeurope.livediagnostics.monitor.azure.com/",
    "Name": "henglu9128",
    "CreationDate": "2023-07-11T08:24:14.7161403+00:00",
    "TenantId": "00000000-0000-0000-0000-000000000000",
    "provisioningState": "Succeeded",
    "SamplingPercentage": 100.0,
    "RetentionInDays": 90,
    "Retention": "P90D",
    "DisableIpMasking": false,
    "IngestionMode": "ApplicationInsights",
    "publicNetworkAccessForIngestion": "Enabled",
    "publicNetworkAccessForQuery": "Enabled",
    "DisableLocalAuth": false,
    "ForceCustomerStorageForProfiler": false,
    "Ver": "v2"
  }
}: timestamp=2023-07-11T16:24:16.674+0800`,
			want: types.RequestTrace{
				Provider:   "azurerm",
				Method:     "",
				StatusCode: 200,
				Host:       "management.azure.com",
				Url:        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu6421/providers/Microsoft.Insights/components/henglu9128?api-version=2020-02-02",
				Response: &types.HttpResponse{
					Headers: map[string]string{
						"Access-Control-Expose-Headers": "Request-Context",
						"Cache-Control":                 "no-cache",
						"Content-Type":                  "application/json; charset=utf-8",
						"Date":                          "Tue, 11 Jul 2023 08:24:16 GMT",
						"Expires":                       "-1",
						"Pragma":                        "no-cache",
						"Request-Context":               "appId=cid-v1:60b64f55-e716-48c8-8b96-83eb9c6a7a9b",
						"Server":                        "Microsoft-IIS/10.0",
						"Strict-Transport-Security":     "max-age=31536000; includeSubDomains",
						"Vary":                          "Accept-Encoding",
						"X-Content-Type-Options":        "nosniff",
						"X-Ms-Correlation-Request-Id":   "14e77fe7-e1a1-b8a1-65d5-40363cac85f5",
						"X-Ms-Ratelimit-Remaining-Subscription-Writes": "1199",
						"X-Ms-Request-Id":         "aa9250b0-9955-4821-ad66-ff354a68d834",
						"X-Ms-Routing-Request-Id": "SOUTHEASTASIA:20230711T082416Z:aa9250b0-9955-4821-ad66-ff354a68d834",
						"X-Powered-By":            "ASP.NET",
					},
					Body: `
{
  "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu6421/providers/microsoft.insights/components/henglu9128",
  "name": "henglu9128",
  "type": "microsoft.insights/components",
  "location": "westeurope",
  "tags": {},
  "kind": "web",
  "etag": "\"920312e9-0000-0200-0000-64ad11ae0000\"",
  "properties": {
    "ApplicationId": "henglu9128",
    "AppId": "76d8a56f-5069-42ea-bae2-645248170408",
    "Application_Type": "web",
    "Flow_Type": null,
    "Request_Source": null,
    "InstrumentationKey": "2d727e28-a681-48af-95aa-677daba464d6",
    "ConnectionString": "InstrumentationKey=2d727e28-a681-48af-95aa-677daba464d6;IngestionEndpoint=https://westeurope-4.in.applicationinsights.azure.com/;LiveEndpoint=https://westeurope.livediagnostics.monitor.azure.com/",
    "Name": "henglu9128",
    "CreationDate": "2023-07-11T08:24:14.7161403+00:00",
    "TenantId": "00000000-0000-0000-0000-000000000000",
    "provisioningState": "Succeeded",
    "SamplingPercentage": 100.0,
    "RetentionInDays": 90,
    "Retention": "P90D",
    "DisableIpMasking": false,
    "IngestionMode": "ApplicationInsights",
    "publicNetworkAccessForIngestion": "Enabled",
    "publicNetworkAccessForQuery": "Enabled",
    "DisableLocalAuth": false,
    "ForceCustomerStorageForProfiler": false,
    "Ver": "v2"
  }
}`,
				},
			},
		},
		{
			name: "azurerm GET response trace",
			log: `2023-06-15T15:10:42.112+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Response for https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/henglu615aa?api-version=2020-06-01: 
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

{"id":"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}: timestamp=2023-06-15T15:10:42.112+0800`,
			want: types.RequestTrace{
				Provider:   "azurerm",
				Method:     "",
				StatusCode: 200,
				Host:       "management.azure.com",
				Url:        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/henglu615aa?api-version=2020-06-01",
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
					Body: `{"id":"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}`,
				},
			},
		},
		{
			name: "azapi PUT response trace",
			log: `2023-06-15T15:12:41.772+0800 [DEBUG] provider.terraform-provider-azurerm_v3.61.0_x5: AzureRM Response for https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/henglu615aa?api-version=2020-06-01: 
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

{"id":"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}: timestamp=2023-06-15T15:12:41.772+0800
`,
			want: types.RequestTrace{
				Provider:   "azurerm",
				StatusCode: 201,
				Host:       "management.azure.com",
				Url:        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/henglu615aa?api-version=2020-06-01",
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
					Body: `{"id":"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/henglu615aa","name":"henglu615aa","type":"Microsoft.Resources/resourceGroups","location":"westeurope","tags":{},"properties":{"provisioningState":"Succeeded"}}`,
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

				var gotBody, wantBody interface{}
				err = json.Unmarshal([]byte(got.Response.Body), &gotBody)
				if err != nil {
					t.Errorf("ParseResponse() response body unmarshal error = %v", err)
				}
				err = json.Unmarshal([]byte(tc.want.Response.Body), &wantBody)
				if err != nil {
					t.Errorf("ParseResponse() response body unmarshal error = %v", err)
				}

				if !reflect.DeepEqual(gotBody, wantBody) {
					t.Errorf("ParseResponse() response body = %v, want %v", gotBody, wantBody)
				}
			}

		})
	}
}
