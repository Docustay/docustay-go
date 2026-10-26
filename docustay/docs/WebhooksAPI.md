# \WebhooksAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDocumentsAiAllowanceLow**](WebhooksAPI.md#WebhookDocumentsAiAllowanceLow) | **Post** /documents.ai.allowance_low | documents.ai.allowance_low
[**WebhookDocumentsApprovalDecided**](WebhooksAPI.md#WebhookDocumentsApprovalDecided) | **Post** /documents.approval.decided | documents.approval.decided
[**WebhookDocumentsApprovalRequested**](WebhooksAPI.md#WebhookDocumentsApprovalRequested) | **Post** /documents.approval.requested | documents.approval.requested
[**WebhookDocumentsDocumentBounced**](WebhooksAPI.md#WebhookDocumentsDocumentBounced) | **Post** /documents.document.bounced | documents.document.bounced
[**WebhookDocumentsDocumentChangesRequested**](WebhooksAPI.md#WebhookDocumentsDocumentChangesRequested) | **Post** /documents.document.changes_requested | documents.document.changes_requested
[**WebhookDocumentsDocumentCreated**](WebhooksAPI.md#WebhookDocumentsDocumentCreated) | **Post** /documents.document.created | documents.document.created
[**WebhookDocumentsDocumentDeclined**](WebhooksAPI.md#WebhookDocumentsDocumentDeclined) | **Post** /documents.document.declined | documents.document.declined
[**WebhookDocumentsDocumentDelegated**](WebhooksAPI.md#WebhookDocumentsDocumentDelegated) | **Post** /documents.document.delegated | documents.document.delegated
[**WebhookDocumentsDocumentDownloaded**](WebhooksAPI.md#WebhookDocumentsDocumentDownloaded) | **Post** /documents.document.downloaded | documents.document.downloaded
[**WebhookDocumentsDocumentExecuted**](WebhooksAPI.md#WebhookDocumentsDocumentExecuted) | **Post** /documents.document.executed | documents.document.executed
[**WebhookDocumentsDocumentExpired**](WebhooksAPI.md#WebhookDocumentsDocumentExpired) | **Post** /documents.document.expired | documents.document.expired
[**WebhookDocumentsDocumentOpened**](WebhooksAPI.md#WebhookDocumentsDocumentOpened) | **Post** /documents.document.opened | documents.document.opened
[**WebhookDocumentsDocumentReminderDue**](WebhooksAPI.md#WebhookDocumentsDocumentReminderDue) | **Post** /documents.document.reminder_due | documents.document.reminder_due
[**WebhookDocumentsDocumentSent**](WebhooksAPI.md#WebhookDocumentsDocumentSent) | **Post** /documents.document.sent | documents.document.sent
[**WebhookDocumentsDocumentSigned**](WebhooksAPI.md#WebhookDocumentsDocumentSigned) | **Post** /documents.document.signed | documents.document.signed
[**WebhookDocumentsDocumentVoided**](WebhooksAPI.md#WebhookDocumentsDocumentVoided) | **Post** /documents.document.voided | documents.document.voided
[**WebhookDocumentsEmailBounced**](WebhooksAPI.md#WebhookDocumentsEmailBounced) | **Post** /documents.email.bounced | documents.email.bounced
[**WebhookDocumentsEmailComplained**](WebhooksAPI.md#WebhookDocumentsEmailComplained) | **Post** /documents.email.complained | documents.email.complained
[**WebhookDocumentsEmailDelivered**](WebhooksAPI.md#WebhookDocumentsEmailDelivered) | **Post** /documents.email.delivered | documents.email.delivered
[**WebhookDocumentsEmailPaused**](WebhooksAPI.md#WebhookDocumentsEmailPaused) | **Post** /documents.email.paused | documents.email.paused
[**WebhookDocumentsPaymentFailed**](WebhooksAPI.md#WebhookDocumentsPaymentFailed) | **Post** /documents.payment.failed | documents.payment.failed
[**WebhookDocumentsPaymentHoldLapsed**](WebhooksAPI.md#WebhookDocumentsPaymentHoldLapsed) | **Post** /documents.payment.hold_lapsed | documents.payment.hold_lapsed
[**WebhookDocumentsPaymentReceived**](WebhooksAPI.md#WebhookDocumentsPaymentReceived) | **Post** /documents.payment.received | documents.payment.received
[**WebhookDocumentsTeamCreated**](WebhooksAPI.md#WebhookDocumentsTeamCreated) | **Post** /documents.team.created | documents.team.created
[**WebhookDocumentsTeamDeleted**](WebhooksAPI.md#WebhookDocumentsTeamDeleted) | **Post** /documents.team.deleted | documents.team.deleted
[**WebhookDocumentsTeamUpdated**](WebhooksAPI.md#WebhookDocumentsTeamUpdated) | **Post** /documents.team.updated | documents.team.updated
[**WebhookDocumentsTemplateCreated**](WebhooksAPI.md#WebhookDocumentsTemplateCreated) | **Post** /documents.template.created | documents.template.created
[**WebhookDocumentsTemplateFieldsDetected**](WebhooksAPI.md#WebhookDocumentsTemplateFieldsDetected) | **Post** /documents.template.fields_detected | documents.template.fields_detected
[**WebhookDocumentsUsageThreshold**](WebhooksAPI.md#WebhookDocumentsUsageThreshold) | **Post** /documents.usage.threshold | documents.usage.threshold



## WebhookDocumentsAiAllowanceLow

> WebhookDocumentsAiAllowanceLow(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.ai.allowance_low



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsAiAllowanceLow(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsAiAllowanceLow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsAiAllowanceLowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsApprovalDecided

> WebhookDocumentsApprovalDecided(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.approval.decided



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsApprovalDecided(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsApprovalDecided``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsApprovalDecidedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsApprovalRequested

> WebhookDocumentsApprovalRequested(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.approval.requested



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsApprovalRequested(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsApprovalRequested``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsApprovalRequestedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentBounced

> WebhookDocumentsDocumentBounced(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.bounced



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentBounced(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentBounced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentBouncedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentChangesRequested

> WebhookDocumentsDocumentChangesRequested(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.changes_requested



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentChangesRequested(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentChangesRequested``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentChangesRequestedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentCreated

> WebhookDocumentsDocumentCreated(ctx).WebhookDocumentsDocumentCreatedRequest(webhookDocumentsDocumentCreatedRequest).Execute()

documents.document.created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentCreatedRequest :=  // WebhookDocumentsDocumentCreatedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentCreated(context.Background()).WebhookDocumentsDocumentCreatedRequest(webhookDocumentsDocumentCreatedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentCreatedRequest** | [**WebhookDocumentsDocumentCreatedRequest**](WebhookDocumentsDocumentCreatedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentDeclined

> WebhookDocumentsDocumentDeclined(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.declined



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentDeclined(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentDeclined``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentDeclinedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentDelegated

> WebhookDocumentsDocumentDelegated(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.delegated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentDelegated(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentDelegated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentDelegatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentDownloaded

> WebhookDocumentsDocumentDownloaded(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.downloaded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentDownloaded(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentDownloaded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentDownloadedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentExecuted

> WebhookDocumentsDocumentExecuted(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.executed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentExecuted(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentExecuted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentExecutedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentExpired

> WebhookDocumentsDocumentExpired(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.expired



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentExpired(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentExpired``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentExpiredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentOpened

> WebhookDocumentsDocumentOpened(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.opened



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentOpened(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentOpened``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentOpenedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentReminderDue

> WebhookDocumentsDocumentReminderDue(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.reminder_due



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentReminderDue(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentReminderDue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentReminderDueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentSent

> WebhookDocumentsDocumentSent(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.sent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentSent(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentSent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentSentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentSigned

> WebhookDocumentsDocumentSigned(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.signed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentSigned(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentSigned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentSignedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsDocumentVoided

> WebhookDocumentsDocumentVoided(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.document.voided



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsDocumentVoided(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsDocumentVoided``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsDocumentVoidedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsEmailBounced

> WebhookDocumentsEmailBounced(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.email.bounced



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsEmailBounced(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsEmailBounced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsEmailBouncedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsEmailComplained

> WebhookDocumentsEmailComplained(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.email.complained



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsEmailComplained(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsEmailComplained``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsEmailComplainedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsEmailDelivered

> WebhookDocumentsEmailDelivered(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.email.delivered



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsEmailDelivered(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsEmailDelivered``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsEmailDeliveredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsEmailPaused

> WebhookDocumentsEmailPaused(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.email.paused



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsEmailPaused(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsEmailPaused``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsEmailPausedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsPaymentFailed

> WebhookDocumentsPaymentFailed(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.payment.failed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsPaymentFailed(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsPaymentFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsPaymentFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsPaymentHoldLapsed

> WebhookDocumentsPaymentHoldLapsed(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.payment.hold_lapsed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsPaymentHoldLapsed(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsPaymentHoldLapsed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsPaymentHoldLapsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsPaymentReceived

> WebhookDocumentsPaymentReceived(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.payment.received



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsPaymentReceived(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsPaymentReceived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsPaymentReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsTeamCreated

> WebhookDocumentsTeamCreated(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.team.created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsTeamCreated(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsTeamCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsTeamCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsTeamDeleted

> WebhookDocumentsTeamDeleted(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.team.deleted



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsTeamDeleted(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsTeamDeleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsTeamDeletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsTeamUpdated

> WebhookDocumentsTeamUpdated(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.team.updated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsTeamUpdated(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsTeamUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsTeamUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsTemplateCreated

> WebhookDocumentsTemplateCreated(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.template.created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsTemplateCreated(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsTemplateCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsTemplateCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsTemplateFieldsDetected

> WebhookDocumentsTemplateFieldsDetected(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.template.fields_detected



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsTemplateFieldsDetected(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsTemplateFieldsDetected``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsTemplateFieldsDetectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDocumentsUsageThreshold

> WebhookDocumentsUsageThreshold(ctx).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()

documents.usage.threshold



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	webhookDocumentsDocumentSentRequest :=  // WebhookDocumentsDocumentSentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhookDocumentsUsageThreshold(context.Background()).WebhookDocumentsDocumentSentRequest(webhookDocumentsDocumentSentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhookDocumentsUsageThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDocumentsUsageThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDocumentsDocumentSentRequest** | [**WebhookDocumentsDocumentSentRequest**](WebhookDocumentsDocumentSentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

