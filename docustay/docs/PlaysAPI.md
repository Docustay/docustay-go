# \PlaysAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerDraft**](PlaysAPI.md#AnswerDraft) | **Post** /api/v1/drafts/{draftRef}/answers | Answer a template&#39;s questions
[**ArchivePlay**](PlaysAPI.md#ArchivePlay) | **Delete** /api/v1/plays/{name} | Archive a play
[**DuplicatePlay**](PlaysAPI.md#DuplicatePlay) | **Post** /api/v1/plays/{name}/duplicate | Copy a play under a new name
[**GetPlay**](PlaysAPI.md#GetPlay) | **Get** /api/v1/plays/{name} | One play
[**ListPlaybooks**](PlaysAPI.md#ListPlaybooks) | **Get** /api/v1/playbooks | List playbooks (packs)
[**ListPlays**](PlaysAPI.md#ListPlays) | **Get** /api/v1/plays | List plays
[**PublishPlay**](PlaysAPI.md#PublishPlay) | **Post** /api/v1/plays/{name}/publish | Publish a play version
[**SavePlay**](PlaysAPI.md#SavePlay) | **Post** /api/v1/plays | Save a play (a new version)
[**SavePlaybook**](PlaysAPI.md#SavePlaybook) | **Post** /api/v1/playbooks | Save a playbook (pack)
[**SetPlaybookOn**](PlaysAPI.md#SetPlaybookOn) | **Post** /api/v1/playbooks/{name}/enable | Switch a pack on or off
[**StartDraft**](PlaysAPI.md#StartDraft) | **Post** /api/v1/drafts | Start a draft from a template
[**TestPlay**](PlaysAPI.md#TestPlay) | **Post** /api/v1/plays/test | Check a play without saving it



## AnswerDraft

> interface{} AnswerDraft(ctx, draftRef).AnswerDraftRequest(answerDraftRequest).Execute()

Answer a template's questions

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
	draftRef := "draftRef_example" // string | 
	answerDraftRequest := *openapiclient.NewAnswerDraftRequest(map[string]StartDraftRequestAnswersValue{"key": *openapiclient.NewStartDraftRequestAnswersValue()}) // AnswerDraftRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.AnswerDraft(context.Background(), draftRef).AnswerDraftRequest(answerDraftRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.AnswerDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnswerDraft`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.AnswerDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **answerDraftRequest** | [**AnswerDraftRequest**](AnswerDraftRequest.md) |  | 

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


## ArchivePlay

> interface{} ArchivePlay(ctx, name).Execute()

Archive a play

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.ArchivePlay(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.ArchivePlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchivePlay`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.ArchivePlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivePlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DuplicatePlay

> interface{} DuplicatePlay(ctx, name).DuplicatePlayRequest(duplicatePlayRequest).Execute()

Copy a play under a new name

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
	name := "name_example" // string | 
	duplicatePlayRequest := *openapiclient.NewDuplicatePlayRequest() // DuplicatePlayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.DuplicatePlay(context.Background(), name).DuplicatePlayRequest(duplicatePlayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.DuplicatePlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicatePlay`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.DuplicatePlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicatePlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicatePlayRequest** | [**DuplicatePlayRequest**](DuplicatePlayRequest.md) |  | 

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


## GetPlay

> interface{} GetPlay(ctx, name).Execute()

One play

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.GetPlay(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.GetPlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlay`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.GetPlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListPlaybooks

> interface{} ListPlaybooks(ctx).Execute()

List playbooks (packs)



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
	resp, r, err := apiClient.PlaysAPI.ListPlaybooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.ListPlaybooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlaybooks`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.ListPlaybooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPlaybooksRequest struct via the builder pattern


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


## ListPlays

> interface{} ListPlays(ctx).Kind(kind).Category(category).Execute()

List plays



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
	kind := "kind_example" // string |  (optional)
	category := "category_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.ListPlays(context.Background()).Kind(kind).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.ListPlays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlays`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.ListPlays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 
 **category** | **string** |  | 

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


## PublishPlay

> interface{} PublishPlay(ctx, name).PublishPlayRequest(publishPlayRequest).Execute()

Publish a play version

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
	name := "name_example" // string | 
	publishPlayRequest := *openapiclient.NewPublishPlayRequest() // PublishPlayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.PublishPlay(context.Background(), name).PublishPlayRequest(publishPlayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.PublishPlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishPlay`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.PublishPlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publishPlayRequest** | [**PublishPlayRequest**](PublishPlayRequest.md) |  | 

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


## SavePlay

> interface{} SavePlay(ctx).SavePlayRequest(savePlayRequest).Execute()

Save a play (a new version)



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
	savePlayRequest := *openapiclient.NewSavePlayRequest("Source_example") // SavePlayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.SavePlay(context.Background()).SavePlayRequest(savePlayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.SavePlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SavePlay`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.SavePlay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **savePlayRequest** | [**SavePlayRequest**](SavePlayRequest.md) |  | 

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


## SavePlaybook

> interface{} SavePlaybook(ctx).SavePlaybookRequest(savePlaybookRequest).Execute()

Save a playbook (pack)



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
	savePlaybookRequest := *openapiclient.NewSavePlaybookRequest("Name_example", "Title_example") // SavePlaybookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.SavePlaybook(context.Background()).SavePlaybookRequest(savePlaybookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.SavePlaybook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SavePlaybook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.SavePlaybook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePlaybookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **savePlaybookRequest** | [**SavePlaybookRequest**](SavePlaybookRequest.md) |  | 

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


## SetPlaybookOn

> interface{} SetPlaybookOn(ctx, name).SetPlaybookOnRequest(setPlaybookOnRequest).Execute()

Switch a pack on or off

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
	name := "name_example" // string | 
	setPlaybookOnRequest := *openapiclient.NewSetPlaybookOnRequest(false) // SetPlaybookOnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.SetPlaybookOn(context.Background(), name).SetPlaybookOnRequest(setPlaybookOnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.SetPlaybookOn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPlaybookOn`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.SetPlaybookOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPlaybookOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setPlaybookOnRequest** | [**SetPlaybookOnRequest**](SetPlaybookOnRequest.md) |  | 

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


## StartDraft

> interface{} StartDraft(ctx).StartDraftRequest(startDraftRequest).Execute()

Start a draft from a template



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
	startDraftRequest := *openapiclient.NewStartDraftRequest("Template_example", []openapiclient.StartDraftRequestPeopleInner{*openapiclient.NewStartDraftRequestPeopleInner("Name_example")}) // StartDraftRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.StartDraft(context.Background()).StartDraftRequest(startDraftRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.StartDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartDraft`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.StartDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDraftRequest** | [**StartDraftRequest**](StartDraftRequest.md) |  | 

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


## TestPlay

> interface{} TestPlay(ctx).SavePlayRequest(savePlayRequest).Execute()

Check a play without saving it



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
	savePlayRequest := *openapiclient.NewSavePlayRequest("Source_example") // SavePlayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaysAPI.TestPlay(context.Background()).SavePlayRequest(savePlayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaysAPI.TestPlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestPlay`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaysAPI.TestPlay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **savePlayRequest** | [**SavePlayRequest**](SavePlayRequest.md) |  | 

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

