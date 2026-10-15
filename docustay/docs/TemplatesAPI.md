# \TemplatesAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateFromFile**](TemplatesAPI.md#CreateTemplateFromFile) | **Post** /api/v1/templates/file | Make a template from a file with text tags
[**CreateTemplateFromHtml**](TemplatesAPI.md#CreateTemplateFromHtml) | **Post** /api/v1/templates/html | Make a template from HTML with text tags
[**CreateTemplateFromSource**](TemplatesAPI.md#CreateTemplateFromSource) | **Post** /api/v1/templates/source | Make a template from a document source
[**CustomizeLibraryTemplate**](TemplatesAPI.md#CustomizeLibraryTemplate) | **Post** /api/v1/templates/library/{slug}/customize | Customize a built-in template
[**DraftFromLibraryTemplate**](TemplatesAPI.md#DraftFromLibraryTemplate) | **Post** /api/v1/templates/library/{slug}/documents | Make a draft from a built-in template
[**GetLibraryTemplate**](TemplatesAPI.md#GetLibraryTemplate) | **Get** /api/v1/templates/library/{slug} | One built-in template
[**GetTemplate**](TemplatesAPI.md#GetTemplate) | **Get** /api/v1/templates/{id} | One template and its seats
[**ListLibraryTemplates**](TemplatesAPI.md#ListLibraryTemplates) | **Get** /api/v1/templates/library | List the built-in templates
[**SetTemplateDefaults**](TemplatesAPI.md#SetTemplateDefaults) | **Post** /api/v1/templates/{id}/defaults | Set what a document made from a template starts with
[**UseTemplate**](TemplatesAPI.md#UseTemplate) | **Post** /api/v1/templates/{id}/documents | Use a written template for people



## CreateTemplateFromFile

> TemplateMade CreateTemplateFromFile(ctx).FileTemplateRequest(fileTemplateRequest).Execute()

Make a template from a file with text tags



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
	fileTemplateRequest := *openapiclient.NewFileTemplateRequest("Name_example", "FileName_example", "File_example") // FileTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplateFromFile(context.Background()).FileTemplateRequest(fileTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplateFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateFromFile`: TemplateMade
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplateFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileTemplateRequest** | [**FileTemplateRequest**](FileTemplateRequest.md) |  | 

### Return type

[**TemplateMade**](TemplateMade.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplateFromHtml

> TemplateMade CreateTemplateFromHtml(ctx).HtmlTemplateRequest(htmlTemplateRequest).Execute()

Make a template from HTML with text tags



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
	htmlTemplateRequest := *openapiclient.NewHtmlTemplateRequest("Name_example", "Html_example") // HtmlTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplateFromHtml(context.Background()).HtmlTemplateRequest(htmlTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplateFromHtml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateFromHtml`: TemplateMade
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplateFromHtml`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateFromHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **htmlTemplateRequest** | [**HtmlTemplateRequest**](HtmlTemplateRequest.md) |  | 

### Return type

[**TemplateMade**](TemplateMade.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplateFromSource

> CreateTemplateFromSource201Response CreateTemplateFromSource(ctx).CreateTemplateFromSourceRequest(createTemplateFromSourceRequest).Execute()

Make a template from a document source



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
	createTemplateFromSourceRequest := *openapiclient.NewCreateTemplateFromSourceRequest("Name_example", map[string]interface{}{"key": interface{}(123)}) // CreateTemplateFromSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplateFromSource(context.Background()).CreateTemplateFromSourceRequest(createTemplateFromSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplateFromSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateFromSource`: CreateTemplateFromSource201Response
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplateFromSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateFromSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTemplateFromSourceRequest** | [**CreateTemplateFromSourceRequest**](CreateTemplateFromSourceRequest.md) |  | 

### Return type

[**CreateTemplateFromSource201Response**](CreateTemplateFromSource201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomizeLibraryTemplate

> interface{} CustomizeLibraryTemplate(ctx, slug).CustomizeLibraryTemplateRequest(customizeLibraryTemplateRequest).Execute()

Customize a built-in template



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
	slug := "slug_example" // string | 
	customizeLibraryTemplateRequest := *openapiclient.NewCustomizeLibraryTemplateRequest() // CustomizeLibraryTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CustomizeLibraryTemplate(context.Background(), slug).CustomizeLibraryTemplateRequest(customizeLibraryTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CustomizeLibraryTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomizeLibraryTemplate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CustomizeLibraryTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomizeLibraryTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customizeLibraryTemplateRequest** | [**CustomizeLibraryTemplateRequest**](CustomizeLibraryTemplateRequest.md) |  | 

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


## DraftFromLibraryTemplate

> interface{} DraftFromLibraryTemplate(ctx, slug).DraftFromLibraryTemplateRequest(draftFromLibraryTemplateRequest).Execute()

Make a draft from a built-in template



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
	slug := "slug_example" // string | 
	draftFromLibraryTemplateRequest := *openapiclient.NewDraftFromLibraryTemplateRequest() // DraftFromLibraryTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.DraftFromLibraryTemplate(context.Background(), slug).DraftFromLibraryTemplateRequest(draftFromLibraryTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DraftFromLibraryTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DraftFromLibraryTemplate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.DraftFromLibraryTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDraftFromLibraryTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **draftFromLibraryTemplateRequest** | [**DraftFromLibraryTemplateRequest**](DraftFromLibraryTemplateRequest.md) |  | 

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


## GetLibraryTemplate

> interface{} GetLibraryTemplate(ctx, slug).Execute()

One built-in template



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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetLibraryTemplate(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetLibraryTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryTemplate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetLibraryTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryTemplateRequest struct via the builder pattern


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


## GetTemplate

> Template GetTemplate(ctx, id).Execute()

One template and its seats

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
	resp, r, err := apiClient.TemplatesAPI.GetTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Template**](Template.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLibraryTemplates

> interface{} ListLibraryTemplates(ctx).Kind(kind).Execute()

List the built-in templates



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.ListLibraryTemplates(context.Background()).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.ListLibraryTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLibraryTemplates`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.ListLibraryTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLibraryTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 

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


## SetTemplateDefaults

> SetTemplateDefaults200Response SetTemplateDefaults(ctx, id).SetTemplateDefaultsRequest(setTemplateDefaultsRequest).Execute()

Set what a document made from a template starts with

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
	setTemplateDefaultsRequest := *openapiclient.NewSetTemplateDefaultsRequest() // SetTemplateDefaultsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.SetTemplateDefaults(context.Background(), id).SetTemplateDefaultsRequest(setTemplateDefaultsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.SetTemplateDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTemplateDefaults`: SetTemplateDefaults200Response
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.SetTemplateDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTemplateDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setTemplateDefaultsRequest** | [**SetTemplateDefaultsRequest**](SetTemplateDefaultsRequest.md) |  | 

### Return type

[**SetTemplateDefaults200Response**](SetTemplateDefaults200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseTemplate

> DocumentMade UseTemplate(ctx, id).UseTemplateRequest(useTemplateRequest).Execute()

Use a written template for people



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
	useTemplateRequest := *openapiclient.NewUseTemplateRequest([]openapiclient.Person{*openapiclient.NewPerson("Name_example", "Email_example")}) // UseTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.UseTemplate(context.Background(), id).UseTemplateRequest(useTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.UseTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseTemplate`: DocumentMade
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.UseTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useTemplateRequest** | [**UseTemplateRequest**](UseTemplateRequest.md) |  | 

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

