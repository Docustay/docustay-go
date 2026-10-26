# \AccountAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiUsage**](AccountAPI.md#AiUsage) | **Get** /api/v1/ai/usage | AI drafts used against your plan
[**PaymentsStatus**](AccountAPI.md#PaymentsStatus) | **Get** /api/v1/payments/status | Is Stripe ready to take a payment?
[**Preflight**](AccountAPI.md#Preflight) | **Post** /api/v1/preflight | Ask what an action would do to your plan, before doing it
[**Usage**](AccountAPI.md#Usage) | **Get** /api/v1/usage | Documents, templates and emails used against your plan



## AiUsage

> AiUsage AiUsage(ctx).Execute()

AI drafts used against your plan



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AiUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AiUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiUsage`: AiUsage
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AiUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiUsageRequest struct via the builder pattern


### Return type

[**AiUsage**](AiUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsStatus

> PaymentsStatus200Response PaymentsStatus(ctx).Execute()

Is Stripe ready to take a payment?



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PaymentsStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PaymentsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsStatus`: PaymentsStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PaymentsStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsStatusRequest struct via the builder pattern


### Return type

[**PaymentsStatus200Response**](PaymentsStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Preflight

> PreflightAnswer Preflight(ctx).PreflightRequest(preflightRequest).Execute()

Ask what an action would do to your plan, before doing it



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
	preflightRequest := *openapiclient.NewPreflightRequest("Action_example") // PreflightRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.Preflight(context.Background()).PreflightRequest(preflightRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Preflight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Preflight`: PreflightAnswer
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Preflight`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreflightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preflightRequest** | [**PreflightRequest**](PreflightRequest.md) |  | 

### Return type

[**PreflightAnswer**](PreflightAnswer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Usage

> Usage Usage(ctx).Execute()

Documents, templates and emails used against your plan



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.Usage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Usage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Usage`: Usage
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Usage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsageRequest struct via the builder pattern


### Return type

[**Usage**](Usage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

