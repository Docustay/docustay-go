# \DocumentsAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogCsv**](DocumentsAPI.md#AuditLogCsv) | **Get** /api/v1/documents/{id}/audit-log.csv | The document&#39;s audit log as CSV
[**CancelAiJob**](DocumentsAPI.md#CancelAiJob) | **Post** /api/v1/documents/ai/jobs/{id}/cancel | Cancel an AI job
[**CheckDocument**](DocumentsAPI.md#CheckDocument) | **Post** /api/v1/documents/check | Check a document&#39;s words before making it
[**CreateDocument**](DocumentsAPI.md#CreateDocument) | **Post** /api/v1/documents | Make a document from words and people
[**DeleteDraft**](DocumentsAPI.md#DeleteDraft) | **Delete** /api/v1/documents/{id} | Delete a draft that was never sent
[**DetectFields**](DocumentsAPI.md#DetectFields) | **Post** /api/v1/documents/detect-fields | Find the fields on a PDF
[**DownloadSignedPdf**](DocumentsAPI.md#DownloadSignedPdf) | **Get** /api/v1/documents/{id}/signed.pdf | Download the signed copy
[**DraftWithAi**](DocumentsAPI.md#DraftWithAi) | **Post** /api/v1/documents/ai/draft | Ask for an AI-written draft
[**EditWithAi**](DocumentsAPI.md#EditWithAi) | **Post** /api/v1/documents/ai/edit | Ask the assistant to change a document
[**EmbedSession**](DocumentsAPI.md#EmbedSession) | **Post** /api/v1/documents/{id}/embed-session | Start an embedded signing session for a seat
[**GetAiDraft**](DocumentsAPI.md#GetAiDraft) | **Get** /api/v1/documents/ai/jobs/{id} | Read an AI draft job
[**GetDocument**](DocumentsAPI.md#GetDocument) | **Get** /api/v1/documents/{id} | One document and where each person is
[**GetDraftSource**](DocumentsAPI.md#GetDraftSource) | **Get** /api/v1/documents/{id}/source | The words, people and choices of a draft
[**GetWizard**](DocumentsAPI.md#GetWizard) | **Get** /api/v1/ai/wizard | The wizard&#39;s questions
[**ListDocuments**](DocumentsAPI.md#ListDocuments) | **Get** /api/v1/documents | List documents
[**ListTemplates**](DocumentsAPI.md#ListTemplates) | **Get** /api/v1/documents/templates | Templates you can send
[**Remind**](DocumentsAPI.md#Remind) | **Post** /api/v1/documents/{id}/remind | Remind the people who still have to sign
[**ReviseDocument**](DocumentsAPI.md#ReviseDocument) | **Post** /api/v1/documents/{id}/revise | Change the words of a draft
[**Search**](DocumentsAPI.md#Search) | **Get** /api/v1/search | Search documents, templates and people
[**SendDocument**](DocumentsAPI.md#SendDocument) | **Post** /api/v1/documents/send | Send a template to people to sign
[**SendDraft**](DocumentsAPI.md#SendDraft) | **Post** /api/v1/documents/{id}/send | Send a draft
[**TestSign**](DocumentsAPI.md#TestSign) | **Post** /api/v1/documents/{id}/test-sign | Sign a test document (test mode only)
[**VoidDocument**](DocumentsAPI.md#VoidDocument) | **Post** /api/v1/documents/{id}/void | Void a document that is still out for signature



## AuditLogCsv

> string AuditLogCsv(ctx, id).Execute()

The document's audit log as CSV



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.AuditLogCsv(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.AuditLogCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditLogCsv`: string
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.AuditLogCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAiJob

> AiDraftJob CancelAiJob(ctx, id).Execute()

Cancel an AI job



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.CancelAiJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CancelAiJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAiJob`: AiDraftJob
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CancelAiJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAiJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiDraftJob**](AiDraftJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckDocument

> CheckDocument200Response CheckDocument(ctx).CheckDocumentRequest(checkDocumentRequest).Execute()

Check a document's words before making it



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
	checkDocumentRequest := *openapiclient.NewCheckDocumentRequest(map[string]*interface{}{"key": interface{}(123)}) // CheckDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.CheckDocument(context.Background()).CheckDocumentRequest(checkDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CheckDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDocument`: CheckDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CheckDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDocumentRequest** | [**CheckDocumentRequest**](CheckDocumentRequest.md) |  | 

### Return type

[**CheckDocument200Response**](CheckDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> DocumentMade CreateDocument(ctx).IdempotencyKey(idempotencyKey).CreateDocumentRequest(createDocumentRequest).Execute()

Make a document from words and people



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
	idempotencyKey := "idempotencyKey_example" // string | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice.
	createDocumentRequest := *openapiclient.NewCreateDocumentRequest("Title_example", map[string]*interface{}{"key": interface{}(123)}, []openapiclient.Person{*openapiclient.NewPerson("Name_example", "Email_example")}) // CreateDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.CreateDocument(context.Background()).IdempotencyKey(idempotencyKey).CreateDocumentRequest(createDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CreateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDocument`: DocumentMade
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CreateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice. | 
 **createDocumentRequest** | [**CreateDocumentRequest**](CreateDocumentRequest.md) |  | 

### Return type

[**DocumentMade**](DocumentMade.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDraft

> DeleteDraft200Response DeleteDraft(ctx, id).Execute()

Delete a draft that was never sent



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DeleteDraft(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DeleteDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDraft`: DeleteDraft200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DeleteDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDraft200Response**](DeleteDraft200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetectFields

> DetectFieldsResult DetectFields(ctx).DetectFieldsRequest(detectFieldsRequest).Execute()

Find the fields on a PDF



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
	detectFieldsRequest := *openapiclient.NewDetectFieldsRequest("File_example") // DetectFieldsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DetectFields(context.Background()).DetectFieldsRequest(detectFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DetectFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetectFields`: DetectFieldsResult
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DetectFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetectFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detectFieldsRequest** | [**DetectFieldsRequest**](DetectFieldsRequest.md) |  | 

### Return type

[**DetectFieldsResult**](DetectFieldsResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadSignedPdf

> *os.File DownloadSignedPdf(ctx, id).Execute()

Download the signed copy



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DownloadSignedPdf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DownloadSignedPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSignedPdf`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DownloadSignedPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSignedPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DraftWithAi

> DraftWithAi202Response DraftWithAi(ctx).DraftWithAiRequest(draftWithAiRequest).Execute()

Ask for an AI-written draft



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
	draftWithAiRequest := *openapiclient.NewDraftWithAiRequest() // DraftWithAiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DraftWithAi(context.Background()).DraftWithAiRequest(draftWithAiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DraftWithAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DraftWithAi`: DraftWithAi202Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DraftWithAi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDraftWithAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **draftWithAiRequest** | [**DraftWithAiRequest**](DraftWithAiRequest.md) |  | 

### Return type

[**DraftWithAi202Response**](DraftWithAi202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWithAi

> DraftWithAi202Response EditWithAi(ctx).EditWithAiRequest(editWithAiRequest).Execute()

Ask the assistant to change a document



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
	editWithAiRequest := *openapiclient.NewEditWithAiRequest("Request_example", map[string]*interface{}{"key": interface{}(123)}) // EditWithAiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.EditWithAi(context.Background()).EditWithAiRequest(editWithAiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.EditWithAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditWithAi`: DraftWithAi202Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.EditWithAi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditWithAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editWithAiRequest** | [**EditWithAiRequest**](EditWithAiRequest.md) |  | 

### Return type

[**DraftWithAi202Response**](DraftWithAi202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmbedSession

> EmbedSession200Response EmbedSession(ctx, id).EmbedSessionRequest(embedSessionRequest).Execute()

Start an embedded signing session for a seat



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 
	embedSessionRequest := *openapiclient.NewEmbedSessionRequest("signer_1", "https://app.example.com") // EmbedSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.EmbedSession(context.Background(), id).EmbedSessionRequest(embedSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.EmbedSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmbedSession`: EmbedSession200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.EmbedSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmbedSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embedSessionRequest** | [**EmbedSessionRequest**](EmbedSessionRequest.md) |  | 

### Return type

[**EmbedSession200Response**](EmbedSession200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDraft

> AiDraftJob GetAiDraft(ctx, id).Execute()

Read an AI draft job



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GetAiDraft(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetAiDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDraft`: AiDraftJob
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetAiDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiDraftJob**](AiDraftJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> Document GetDocument(ctx, id).Execute()

One document and where each person is



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GetDocument(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocument`: Document
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftSource

> GetDraftSource200Response GetDraftSource(ctx, id).Execute()

The words, people and choices of a draft



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GetDraftSource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDraftSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftSource`: GetDraftSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDraftSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDraftSource200Response**](GetDraftSource200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWizard

> map[string]*interface{} GetWizard(ctx).Execute()

The wizard's questions



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
	resp, r, err := apiClient.DocumentsAPI.GetWizard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetWizard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWizard`: map[string]*interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetWizard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWizardRequest struct via the builder pattern


### Return type

**map[string]*interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> ListDocuments200Response ListDocuments(ctx).State(state).Limit(limit).Cursor(cursor).ExternalId(externalId).Execute()

List documents



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
	state := "state_example" // string | Only documents in this state; any other value is refused. (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	externalId := "externalId_example" // string | Only documents with this `externalId`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.ListDocuments(context.Background()).State(state).Limit(limit).Cursor(cursor).ExternalId(externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.ListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDocuments`: ListDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Only documents in this state; any other value is refused. | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 
 **externalId** | **string** | Only documents with this &#x60;externalId&#x60;. | 

### Return type

[**ListDocuments200Response**](ListDocuments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> ListTemplates200Response ListTemplates(ctx).Execute()

Templates you can send



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
	resp, r, err := apiClient.DocumentsAPI.ListTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.ListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplates`: ListTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.ListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


### Return type

[**ListTemplates200Response**](ListTemplates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Remind

> Remind200Response Remind(ctx, id).RemindRequest(remindRequest).Execute()

Remind the people who still have to sign



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 
	remindRequest := *openapiclient.NewRemindRequest() // RemindRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.Remind(context.Background(), id).RemindRequest(remindRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.Remind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Remind`: Remind200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.Remind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remindRequest** | [**RemindRequest**](RemindRequest.md) |  | 

### Return type

[**Remind200Response**](Remind200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviseDocument

> ReviseDocument201Response ReviseDocument(ctx, id).IdempotencyKey(idempotencyKey).ReviseDocumentRequest(reviseDocumentRequest).Execute()

Change the words of a draft



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 
	idempotencyKey := "idempotencyKey_example" // string | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice.
	reviseDocumentRequest := *openapiclient.NewReviseDocumentRequest(map[string]*interface{}{"key": interface{}(123)}) // ReviseDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.ReviseDocument(context.Background(), id).IdempotencyKey(idempotencyKey).ReviseDocumentRequest(reviseDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.ReviseDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviseDocument`: ReviseDocument201Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.ReviseDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviseDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice. | 
 **reviseDocumentRequest** | [**ReviseDocumentRequest**](ReviseDocumentRequest.md) |  | 

### Return type

[**ReviseDocument201Response**](ReviseDocument201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> interface{} Search(ctx).Q(q).Per(per).Execute()

Search documents, templates and people



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
	q := "q_example" // string | 
	per := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.Search(context.Background()).Q(q).Per(per).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **per** | **int32** |  | 

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


## SendDocument

> SendDocument201Response SendDocument(ctx).IdempotencyKey(idempotencyKey).SendRequest(sendRequest).Execute()

Send a template to people to sign



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
	idempotencyKey := "idempotencyKey_example" // string | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice.
	sendRequest := *openapiclient.NewSendRequest("6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e", []openapiclient.SendRequestPartiesInner{*openapiclient.NewSendRequestPartiesInner("Seat_example", "Name_example", "Email_example")}) // SendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.SendDocument(context.Background()).IdempotencyKey(idempotencyKey).SendRequest(sendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SendDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendDocument`: SendDocument201Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SendDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice. | 
 **sendRequest** | [**SendRequest**](SendRequest.md) |  | 

### Return type

[**SendDocument201Response**](SendDocument201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDraft

> SendDraft201Response SendDraft(ctx, id).IdempotencyKey(idempotencyKey).SendDraftRequest(sendDraftRequest).Execute()

Send a draft



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 
	idempotencyKey := "idempotencyKey_example" // string | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice.
	sendDraftRequest := *openapiclient.NewSendDraftRequest() // SendDraftRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.SendDraft(context.Background(), id).IdempotencyKey(idempotencyKey).SendDraftRequest(sendDraftRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SendDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendDraft`: SendDraft201Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SendDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique value for this request (8–200 characters, for example a UUID). Repeating the request with the same key and body returns the first answer instead of doing it twice. | 
 **sendDraftRequest** | [**SendDraftRequest**](SendDraftRequest.md) |  | 

### Return type

[**SendDraft201Response**](SendDraft201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSign

> TestSign200Response TestSign(ctx, id).TestSignRequest(testSignRequest).Execute()

Sign a test document (test mode only)



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 
	testSignRequest := *openapiclient.NewTestSignRequest("signer_1") // TestSignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.TestSign(context.Background(), id).TestSignRequest(testSignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.TestSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSign`: TestSign200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.TestSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testSignRequest** | [**TestSignRequest**](TestSignRequest.md) |  | 

### Return type

[**TestSign200Response**](TestSign200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidDocument

> VoidDocument200Response VoidDocument(ctx, id).VoidDocumentRequest(voidDocumentRequest).Execute()

Void a document that is still out for signature



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
	id := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string | 
	voidDocumentRequest := *openapiclient.NewVoidDocumentRequest() // VoidDocumentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.VoidDocument(context.Background(), id).VoidDocumentRequest(voidDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.VoidDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidDocument`: VoidDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.VoidDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voidDocumentRequest** | [**VoidDocumentRequest**](VoidDocumentRequest.md) |  | 

### Return type

[**VoidDocument200Response**](VoidDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

