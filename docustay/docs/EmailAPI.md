# \EmailAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmailTemplates**](EmailAPI.md#ListEmailTemplates) | **Get** /api/v1/email-templates | The workspace&#39;s email wording
[**ListEmails**](EmailAPI.md#ListEmails) | **Get** /api/v1/emails | The log of email about your documents
[**PreviewEmailTemplate**](EmailAPI.md#PreviewEmailTemplate) | **Post** /api/v1/email-templates/{kind}/preview | Preview an email with made-up names
[**ResetEmailTemplate**](EmailAPI.md#ResetEmailTemplate) | **Delete** /api/v1/email-templates/{kind} | Back to the standard wording
[**SaveEmailTemplate**](EmailAPI.md#SaveEmailTemplate) | **Put** /api/v1/email-templates/{kind} | Set the wording of one email



## ListEmailTemplates

> ListEmailTemplates200Response ListEmailTemplates(ctx).Locale(locale).Execute()

The workspace's email wording

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
	locale := "locale_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.ListEmailTemplates(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.ListEmailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmailTemplates`: ListEmailTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.ListEmailTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** |  | 

### Return type

[**ListEmailTemplates200Response**](ListEmailTemplates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## PreviewEmailTemplate

> map[string]interface{} PreviewEmailTemplate(ctx, kind).EmailWording(emailWording).Execute()

Preview an email with made-up names

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
	kind := "invitation" // string | invitation, reminder, completed, declined, voided, expired, code or receipt
	emailWording := *openapiclient.NewEmailWording() // EmailWording |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.PreviewEmailTemplate(context.Background(), kind).EmailWording(emailWording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.PreviewEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewEmailTemplate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.PreviewEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | invitation, reminder, completed, declined, voided, expired, code or receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailWording** | [**EmailWording**](EmailWording.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetEmailTemplate

> map[string]interface{} ResetEmailTemplate(ctx, kind).Locale(locale).Execute()

Back to the standard wording

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
	kind := "invitation" // string | invitation, reminder, completed, declined, voided, expired, code or receipt
	locale := "locale_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.ResetEmailTemplate(context.Background(), kind).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.ResetEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetEmailTemplate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.ResetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | invitation, reminder, completed, declined, voided, expired, code or receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveEmailTemplate

> map[string]interface{} SaveEmailTemplate(ctx, kind).EmailWording(emailWording).Execute()

Set the wording of one email



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
	kind := "invitation" // string | invitation, reminder, completed, declined, voided, expired, code or receipt
	emailWording := *openapiclient.NewEmailWording() // EmailWording | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.SaveEmailTemplate(context.Background(), kind).EmailWording(emailWording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.SaveEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveEmailTemplate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.SaveEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | invitation, reminder, completed, declined, voided, expired, code or receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailWording** | [**EmailWording**](EmailWording.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

