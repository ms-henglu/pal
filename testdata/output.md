<!--
Tips:

1. Use Markdown preview mode to get a better reading experience.
2. If you want to select some of the request traces, in VSCode, use shortcut "Ctrl + K, 0" to fold all blocks.

-->


##### <!--- GET /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 404 -->
<details>
  <summary>
    13:12:48 GET management.azure.com /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 404
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| provider.terraform-provider-azurerm | AzureRM Request: |
| Host | management.azure.com |
| User-Agent | Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 404 Not Found**

| Header | Value |
| ----- | ----- |
| Content-Length | 102 |
| Cache-Control | no-cache |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |
| Content-Type | application/json; charset=utf-8 |
| Date | Fri, 28 Apr 2023 05:12:48 GMT |
| Expires | -1 |
| Pragma | no-cache |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Content-Type-Options | nosniff |
| X-Ms-Failure-Cause | gateway |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11999 |
| X-Ms-Request-Id | d609c356-986d-4b3d-969a-c5465c223aa5 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051248Z:d609c356-986d-4b3d-969a-c5465c223aa5 |


Response Body:
```json
{
    "error": {
        "code": "ResourceGroupNotFound",
        "message": "Resource group 'henglu1114' could not be found."
    }
}
```

</details>
</blockquote>
</details>

-----


##### <!--- PUT /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 201 -->
<details>
  <summary>
    13:12:48 PUT management.azure.com /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 201
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| User-Agent | Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Content-Length | 35 |
| Content-Type | application/json; charset=utf-8 |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |
| Accept-Encoding | gzip |
| provider.terraform-provider-azurerm | AzureRM Request: |
| Host | management.azure.com |


Request Body:
```json
{
    "location": "westeurope",
    "tags": {}
}
```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 201 Created**

| Header | Value |
| ----- | ----- |
| X-Ms-Request-Id | 216d8744-1a9a-46a6-b27c-2fb8613f7d40 |
| Content-Length | 231 |
| Content-Type | application/json; charset=utf-8 |
| Pragma | no-cache |
| X-Content-Type-Options | nosniff |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |
| X-Ms-Ratelimit-Remaining-Subscription-Writes | 1199 |
| Cache-Control | no-cache |
| Date | Fri, 28 Apr 2023 05:12:50 GMT |
| Expires | -1 |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051251Z:216d8744-1a9a-46a6-b27c-2fb8613f7d40 |


Response Body:
```json
{
    "id": "/subscriptions/******/resourceGroups/henglu1114",
    "location": "westeurope",
    "name": "henglu1114",
    "properties": {
        "provisioningState": "Succeeded"
    },
    "tags": {},
    "type": "Microsoft.Resources/resourceGroups"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 200 -->
<details>
  <summary>
    13:12:51 GET management.azure.com /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Host | management.azure.com |
| User-Agent | Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |
| provider.terraform-provider-azurerm | AzureRM Request: |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Pragma | no-cache |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11998 |
| X-Ms-Request-Id | 11b559b8-fee7-449a-981b-7a96a9ffc1f3 |
| Cache-Control | no-cache |
| Date | Fri, 28 Apr 2023 05:12:50 GMT |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Vary | Accept-Encoding |
| X-Content-Type-Options | nosniff |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051251Z:11b559b8-fee7-449a-981b-7a96a9ffc1f3 |
| Content-Type | application/json; charset=utf-8 |
| Expires | -1 |


Response Body:
```json
{
    "id": "/subscriptions/******/resourceGroups/henglu1114",
    "location": "westeurope",
    "name": "henglu1114",
    "properties": {
        "provisioningState": "Succeeded"
    },
    "tags": {},
    "type": "Microsoft.Resources/resourceGroups"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 200 -->
<details>
  <summary>
    13:12:51 GET management.azure.com /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| provider.terraform-provider-azurerm | AzureRM Request: |
| Host | management.azure.com |
| User-Agent | Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| X-Ms-Request-Id | 7f2aac02-fcf5-440b-9f63-24f4765ae7e4 |
| Cache-Control | no-cache |
| Content-Type | application/json; charset=utf-8 |
| Date | Fri, 28 Apr 2023 05:12:50 GMT |
| Expires | -1 |
| Pragma | no-cache |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Vary | Accept-Encoding |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051251Z:7f2aac02-fcf5-440b-9f63-24f4765ae7e4 |
| X-Content-Type-Options | nosniff |
| X-Ms-Correlation-Request-Id | 154f5cd8-95c5-b518-8c4e-b5e31f02b064 |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11997 |


Response Body:
```json
{
    "id": "/subscriptions/******/resourceGroups/henglu1114",
    "location": "westeurope",
    "name": "henglu1114",
    "properties": {
        "provisioningState": "Succeeded"
    },
    "tags": {},
    "type": "Microsoft.Resources/resourceGroups"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=REDACTED 0 -->
<details>
  <summary>
    13:12:51 GET 169.254.169.254 /metadata/identity/oauth2/token?api-version=2018-02-01&resource=REDACTED 0
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Metadata | REDACTED |
| User-Agent | azsdk-go-azidentity/v1.2.0 (go1.19.3; darwin) |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 0 **

| Header | Value |
| ----- | ----- |
| Get "http | //169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.core.windows.net%2F": context deadline exceeded |
| 	/Users/luheng/go/src/github.com/Azure/terraform-provider-azapi/vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported/request.go | 84 |
| Metadata | REDACTED |
| User-Agent | azsdk-go-azidentity/v1.2.0 (go1.19.3; darwin) |


Response Body:
```json

```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 404 -->
<details>
  <summary>
    13:12:52 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 404
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Accept | application/json |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 404 Not Found**

| Header | Value |
| ----- | ----- |
| X-Ms-Failure-Cause | REDACTED |
| RESPONSE Status | 404 Not Found |
| Content-Length | 229 |
| Date | Fri, 28 Apr 2023 05:12:53 GMT |
| Expires | -1 |
| X-Ms-Request-Id | e718dd71-7703-4ca0-a461-9cf510071ede |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Authorization | REDACTED |
| Pragma | no-cache |
| Content-Type | application/json; charset=utf-8 |
| X-Content-Type-Options | nosniff |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051254Z:e718dd71-7703-4ca0-a461-9cf510071ede |
| Accept | application/json |
| Cache-Control | no-cache |


Response Body:
```json
{
    "error": {
        "code": "ResourceNotFound",
        "message": "The Resource 'Microsoft.Automation/automationAccounts/henglu1' under resource group 'henglu1114' was not found. For more details please go to https://aka.ms/ARMResourceNotFoundFix"
    }
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 404 -->
<details>
  <summary>
    13:12:52 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 404
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Accept | application/json |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 404 Not Found**

| Header | Value |
| ----- | ----- |
| Content-Length | 229 |
| Date | Fri, 28 Apr 2023 05:12:53 GMT |
| Expires | -1 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Authorization | REDACTED |
| RESPONSE Status | 404 Not Found |
| X-Content-Type-Options | nosniff |
| X-Ms-Failure-Cause | REDACTED |
| X-Ms-Request-Id | c288ab59-527c-4609-adbe-0f7b940aa2ce |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051254Z:c288ab59-527c-4609-adbe-0f7b940aa2ce |
| Accept | application/json |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Cache-Control | no-cache |
| Content-Type | application/json; charset=utf-8 |
| Pragma | no-cache |


Response Body:
```json
{
    "error": {
        "code": "ResourceNotFound",
        "message": "The Resource 'Microsoft.Automation/automationAccounts/henglu0' under resource group 'henglu1114' was not found. For more details please go to https://aka.ms/ARMResourceNotFoundFix"
    }
}
```

</details>
</blockquote>
</details>

-----


##### <!--- PUT /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 201 -->
<details>
  <summary>
    13:12:54 PUT management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 201
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |
| Accept | application/json |
| Authorization | REDACTED |
| Content-Length | 80 |
| Content-Type | application/json |


Request Body:
```json
{
    "location": "westeurope",
    "name": "henglu1",
    "properties": {
        "sku": {
            "name": "Basic"
        }
    }
}
```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 201 Created**

| Header | Value |
| ----- | ----- |
| Authorization | REDACTED |
| Content-Type | application/json, application/json; charset=utf-8 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051300Z:8cca4a82-b200-4546-9114-96bd6a3b9dfa |
| Accept | application/json |
| Content-Length | 80, 812 |
| RESPONSE Status | 201 Created |
| Date | Fri, 28 Apr 2023 05:13:00 GMT |
| Expires | -1 |
| X-Content-Type-Options | nosniff |
| Location | REDACTED |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Ms-Ratelimit-Remaining-Subscription-Writes | 1198 |
| X-Ms-Request-Id | 5e488609-4d2d-4c51-9efe-ee293d97ddd5 |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Cache-Control | no-cache |
| Pragma | no-cache |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1",
    "location": "westeurope",
    "name": "henglu1",
    "properties": {
        "RegistrationUrl": "https://ff3a2b15-73eb-4d82-af7b-b399b3560822.agentsvc.we.azure-automation.net/accounts/ff3a2b15-73eb-4d82-af7b-b399b3560822",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.247+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.247+00:00",
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- PUT /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 201 -->
<details>
  <summary>
    13:12:54 PUT management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 201
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Accept | application/json |
| Authorization | REDACTED |
| Content-Length | 80 |
| Content-Type | application/json |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |


Request Body:
```json
{
    "location": "westeurope",
    "name": "henglu0",
    "properties": {
        "sku": {
            "name": "Basic"
        }
    }
}
```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 201 Created**

| Header | Value |
| ----- | ----- |
| X-Content-Type-Options | nosniff |
| X-Ms-Ratelimit-Remaining-Subscription-Writes | 1198 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| Date | Fri, 28 Apr 2023 05:13:00 GMT |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051300Z:556e9b08-9786-4a83-bc9f-4fc1470f6b54 |
| Content-Length | 80, 810 |
| RESPONSE Status | 201 Created |
| Cache-Control | no-cache |
| Authorization | REDACTED |
| Content-Type | application/json, application/json; charset=utf-8 |
| X-Ms-Request-Id | 44b6b6af-bc1b-4fa7-92f7-ea78359bbe5f |
| Location | REDACTED |
| Pragma | no-cache |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Accept | application/json |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Expires | -1 |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0",
    "location": "westeurope",
    "name": "henglu0",
    "properties": {
        "RegistrationUrl": "https://166bd667-7a8a-45e2-8a28-b73a65d6b4ae.agentsvc.we.azure-automation.net/accounts/166bd667-7a8a-45e2-8a28-b73a65d6b4ae",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.47+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.47+00:00",
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 200 -->
<details>
  <summary>
    13:13:00 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051301Z:f1f6d44c-90fe-4a4d-940d-93af6bcda3b9 |
| Authorization | REDACTED |
| Content-Type | application/json; charset=utf-8 |
| Ocp-Automation-Accountid | ff3a2b15-73eb-4d82-af7b-b399b3560822 |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Content-Type-Options | nosniff |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11996 |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Pragma | no-cache |
| Vary | Accept-Encoding |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| Date | Fri, 28 Apr 2023 05:13:01 GMT |
| RESPONSE Status | 200 OK |
| Cache-Control | no-cache |
| Expires | -1 |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Request-Id | 13140c2d-30e2-4e5d-a5d2-530ff08558eb |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1",
    "location": "westeurope",
    "name": "henglu1",
    "properties": {
        "RegistrationUrl": "https://ff3a2b15-73eb-4d82-af7b-b399b3560822.agentsvc.we.azure-automation.net/accounts/ff3a2b15-73eb-4d82-af7b-b399b3560822",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.247+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.247+00:00",
        "privateEndpointConnections": [],
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 200 -->
<details>
  <summary>
    13:13:00 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Ocp-Automation-Accountid | 166bd667-7a8a-45e2-8a28-b73a65d6b4ae |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Content-Type-Options | nosniff |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11995 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| RESPONSE Status | 200 OK |
| Content-Type | application/json; charset=utf-8 |
| Date | Fri, 28 Apr 2023 05:13:01 GMT |
| Expires | -1 |
| Vary | Accept-Encoding |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Pragma | no-cache |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051301Z:bdb5d7a5-497d-4720-83bb-599ca597675b |
| Cache-Control | no-cache |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Request-Id | a113d9a0-b3a8-4b4e-8872-99a87624f4b0 |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0",
    "location": "westeurope",
    "name": "henglu0",
    "properties": {
        "RegistrationUrl": "https://166bd667-7a8a-45e2-8a28-b73a65d6b4ae.agentsvc.we.azure-automation.net/accounts/166bd667-7a8a-45e2-8a28-b73a65d6b4ae",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.47+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.47+00:00",
        "privateEndpointConnections": [],
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 200 -->
<details>
  <summary>
    13:13:01 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Accept | application/json |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Ocp-Automation-Accountid | ff3a2b15-73eb-4d82-af7b-b399b3560822 |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Content-Type-Options | nosniff |
| Authorization | REDACTED |
| Content-Type | application/json; charset=utf-8 |
| Date | Fri, 28 Apr 2023 05:13:02 GMT |
| X-Ms-Request-Id | 8e405a38-defd-4d78-a9b5-c723ef2c2293 |
| Accept | application/json |
| Pragma | no-cache |
| Vary | Accept-Encoding |
| Expires | -1 |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11994 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051302Z:387ff158-35ac-412e-80ad-c88ee0d14e4a |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| RESPONSE Status | 200 OK |
| Cache-Control | no-cache |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1",
    "location": "westeurope",
    "name": "henglu1",
    "properties": {
        "RegistrationUrl": "https://ff3a2b15-73eb-4d82-af7b-b399b3560822.agentsvc.we.azure-automation.net/accounts/ff3a2b15-73eb-4d82-af7b-b399b3560822",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.247+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.247+00:00",
        "privateEndpointConnections": [],
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 200 -->
<details>
  <summary>
    13:13:01 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb |
| Accept | application/json |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Authorization | REDACTED |
| RESPONSE Status | 200 OK |
| Expires | -1 |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Vary | Accept-Encoding |
| X-Content-Type-Options | nosniff |
| X-Ms-Correlation-Request-Id | 8817767b-435f-9298-42f8-534407f68afb, 8817767b-435f-9298-42f8-534407f68afb |
| Content-Type | application/json; charset=utf-8 |
| Pragma | no-cache |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11994 |
| X-Ms-Request-Id | 230635c5-4cdf-4349-9d83-0b3fe0fb8a03 |
| Accept | application/json |
| Cache-Control | no-cache |
| Date | Fri, 28 Apr 2023 05:13:02 GMT |
| Ocp-Automation-Accountid | 166bd667-7a8a-45e2-8a28-b73a65d6b4ae |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051302Z:fcb6b959-5c87-4d1d-821e-fd71208420ca |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0",
    "location": "westeurope",
    "name": "henglu0",
    "properties": {
        "RegistrationUrl": "https://166bd667-7a8a-45e2-8a28-b73a65d6b4ae.agentsvc.we.azure-automation.net/accounts/166bd667-7a8a-45e2-8a28-b73a65d6b4ae",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.47+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.47+00:00",
        "privateEndpointConnections": [],
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 200 -->
<details>
  <summary>
    13:13:13 GET management.azure.com /subscriptions/******/resourcegroups/henglu1114?api-version=2020-06-01 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| X-Ms-Correlation-Request-Id | dc2f5658-4cda-512e-baee-03a7aedfa080 |
| provider.terraform-provider-azurerm | AzureRM Request: |
| Host | management.azure.com |
| User-Agent | Go/go1.19.3 (arm64-darwin) go-autorest/v14.2.1 Azure-SDK-For-Go/v66.0.0 resources/2020-06-01 HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.10.1 terraform-provider-azurerm/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Cache-Control | no-cache |
| Date | Fri, 28 Apr 2023 05:13:13 GMT |
| Expires | -1 |
| Vary | Accept-Encoding |
| X-Content-Type-Options | nosniff |
| X-Ms-Correlation-Request-Id | dc2f5658-4cda-512e-baee-03a7aedfa080 |
| Content-Type | application/json; charset=utf-8 |
| Pragma | no-cache |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11999 |
| X-Ms-Request-Id | 5f98c4a6-3701-489e-9565-4f319f109438 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051314Z:5f98c4a6-3701-489e-9565-4f319f109438 |


Response Body:
```json
{
    "id": "/subscriptions/******/resourceGroups/henglu1114",
    "location": "westeurope",
    "name": "henglu1114",
    "properties": {
        "provisioningState": "Succeeded"
    },
    "tags": {},
    "type": "Microsoft.Resources/resourceGroups"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=REDACTED 0 -->
<details>
  <summary>
    13:13:14 GET 169.254.169.254 /metadata/identity/oauth2/token?api-version=2018-02-01&resource=REDACTED 0
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Metadata | REDACTED |
| User-Agent | azsdk-go-azidentity/v1.2.0 (go1.19.3; darwin) |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 0 **

| Header | Value |
| ----- | ----- |
| Metadata | REDACTED |
| User-Agent | azsdk-go-azidentity/v1.2.0 (go1.19.3; darwin) |
| Get "http | //169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.core.windows.net%2F": context deadline exceeded |
| 	/Users/luheng/go/src/github.com/Azure/terraform-provider-azapi/vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported/request.go | 84 |


Response Body:
```json

```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 200 -->
<details>
  <summary>
    13:13:16 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0?api-version=2020-01-13-preview 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| X-Ms-Correlation-Request-Id | d2bfcff9-655a-65de-a871-97fac0990b09 |
| Accept | application/json |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Expires | -1 |
| X-Content-Type-Options | nosniff |
| Accept | application/json |
| Content-Type | application/json; charset=utf-8 |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11998 |
| X-Ms-Request-Id | 9340f464-ef5f-4332-a417-a769906770a6 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051318Z:50ed28d9-6d5a-4fd1-bb17-abe6ddc0b1d2 |
| Authorization | REDACTED |
| Server | Microsoft-HTTPAPI/2.0 |
| Ocp-Automation-Accountid | 166bd667-7a8a-45e2-8a28-b73a65d6b4ae |
| Pragma | no-cache |
| X-Ms-Correlation-Request-Id | d2bfcff9-655a-65de-a871-97fac0990b09, d2bfcff9-655a-65de-a871-97fac0990b09 |
| RESPONSE Status | 200 OK |
| Date | Fri, 28 Apr 2023 05:13:18 GMT |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Vary | Accept-Encoding |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Cache-Control | no-cache |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu0",
    "location": "westeurope",
    "name": "henglu0",
    "properties": {
        "RegistrationUrl": "https://166bd667-7a8a-45e2-8a28-b73a65d6b4ae.agentsvc.we.azure-automation.net/accounts/166bd667-7a8a-45e2-8a28-b73a65d6b4ae",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.47+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.47+00:00",
        "privateEndpointConnections": [],
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----


##### <!--- GET /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 200 -->
<details>
  <summary>
    13:13:16 GET management.azure.com /subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1?api-version=2020-01-13-preview 200
  </summary>
<blockquote>
<details>
  <summary><strong>Request</strong></summary>

| Header | Value |
| ----- | ----- |
| Accept | application/json |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| X-Ms-Correlation-Request-Id | d2bfcff9-655a-65de-a871-97fac0990b09 |


Request Body:
```json

```

</details>
<details>
  <summary><strong>Response</strong></summary>

**Response Status: 200 OK**

| Header | Value |
| ----- | ----- |
| Date | Fri, 28 Apr 2023 05:13:18 GMT |
| Strict-Transport-Security | max-age=31536000; includeSubDomains |
| Vary | Accept-Encoding |
| X-Ms-Request-Id | 4f49197d-d5b6-44ef-9e9b-e72e7ed1d939 |
| Ocp-Automation-Accountid | ff3a2b15-73eb-4d82-af7b-b399b3560822 |
| Server | Microsoft-HTTPAPI/2.0 |
| X-Ms-Routing-Request-Id | SOUTHEASTASIA:20230428T051318Z:05c47f4a-7f6d-4fcd-ab55-fdc2ef45277e |
| Authorization | REDACTED |
| User-Agent | HashiCorp Terraform/1.4.5 (+https://www.terraform.io) Terraform Plugin SDK/2.8.0 terraform-provider-azapi/dev pid-222c6c49-1b0a-5959-a213-6608f9eb8820 |
| Cache-Control | no-cache |
| X-Content-Type-Options | nosniff |
| Expires | -1 |
| Pragma | no-cache |
| X-Ms-Ratelimit-Remaining-Subscription-Reads | 11998 |
| Accept | application/json |
| X-Ms-Correlation-Request-Id | d2bfcff9-655a-65de-a871-97fac0990b09, d2bfcff9-655a-65de-a871-97fac0990b09 |
| RESPONSE Status | 200 OK |
| Content-Type | application/json; charset=utf-8 |


Response Body:
```json
{
    "etag": null,
    "id": "/subscriptions/******/resourceGroups/henglu1114/providers/Microsoft.Automation/automationAccounts/henglu1",
    "location": "westeurope",
    "name": "henglu1",
    "properties": {
        "RegistrationUrl": "https://ff3a2b15-73eb-4d82-af7b-b399b3560822.agentsvc.we.azure-automation.net/accounts/ff3a2b15-73eb-4d82-af7b-b399b3560822",
        "RuntimeConfiguration": {
            "powershell": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            },
            "powershell7": {
                "builtinModules": {
                    "Az": "8.0.0"
                }
            }
        },
        "creationTime": "2023-04-28T05:12:58.247+00:00",
        "encryption": {
            "identity": {
                "userAssignedIdentity": null
            },
            "keySource": "Microsoft.Automation"
        },
        "lastModifiedBy": null,
        "lastModifiedTime": "2023-04-28T05:12:58.247+00:00",
        "privateEndpointConnections": [],
        "sku": {
            "capacity": null,
            "family": null,
            "name": "Basic"
        },
        "state": "Ok"
    },
    "tags": {},
    "type": "Microsoft.Automation/AutomationAccounts"
}
```

</details>
</blockquote>
</details>

-----

