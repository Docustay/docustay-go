# \PaymentsAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPayments**](PaymentsAPI.md#ListPayments) | **Get** /api/v1/payments | List payments taken through documents
[**RefundPayment**](PaymentsAPI.md#RefundPayment) | **Post** /api/v1/payments/{id}/refund | Refund a payment, in full or in part



## ListPayments

> interface{} ListPayments(ctx).Before(before).Limit(limit).Execute()

List payments taken through documents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/docustay/docustay-go/docustay"
)

func main() {
	before := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ListPayments(context.Background()).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ListPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPayments`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ListPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **time.Time** |  | 
 **limit** | **int32** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundPayment

> interface{} RefundPayment(ctx, id).IdempotencyKey(idempotencyKey).RefundPaymentRequest(refundPaymentRequest).Execute()

Refund a payment, in full or in part



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice.
	refundPaymentRequest := *openapiclient.NewRefundPaymentRequest("Reason_example") // RefundPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.RefundPayment(context.Background(), id).IdempotencyKey(idempotencyKey).RefundPaymentRequest(refundPaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.RefundPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundPayment`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.RefundPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice. | 
 **refundPaymentRequest** | [**RefundPaymentRequest**](RefundPaymentRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

