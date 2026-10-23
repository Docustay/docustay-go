# \EmailAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmails**](EmailAPI.md#ListEmails) | **Get** /api/v1/emails | The log of email about your documents



## ListEmails

> ListEmails200Response ListEmails(ctx).DocumentId(documentId).To(to).Status(status).Kind(kind).Limit(limit).Cursor(cursor).Execute()

The log of email about your documents



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
	documentId := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string |  (optional)
	to := "to_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	kind := "kind_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.ListEmails(context.Background()).DocumentId(documentId).To(to).Status(status).Kind(kind).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.ListEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmails`: ListEmails200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.ListEmails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentId** | **string** |  | 
 **to** | **string** |  | 
 **status** | **string** |  | 
 **kind** | **string** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**ListEmails200Response**](ListEmails200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

