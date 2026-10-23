# \EmbeddingAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuilderSession**](EmbeddingAPI.md#BuilderSession) | **Post** /api/v1/embedded/builder-session | Start an embedded template builder session
[**ScribeSession**](EmbeddingAPI.md#ScribeSession) | **Post** /api/v1/embedded/scribe-session | Start an embedded Scribe session



## BuilderSession

> BuilderSession BuilderSession(ctx).BuilderSessionRequest(builderSessionRequest).Execute()

Start an embedded template builder session



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
	builderSessionRequest := *openapiclient.NewBuilderSessionRequest("https://app.example.com") // BuilderSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingAPI.BuilderSession(context.Background()).BuilderSessionRequest(builderSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingAPI.BuilderSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuilderSession`: BuilderSession
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingAPI.BuilderSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuilderSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **builderSessionRequest** | [**BuilderSessionRequest**](BuilderSessionRequest.md) |  | 

### Return type

[**BuilderSession**](BuilderSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScribeSession

> BuilderSession ScribeSession(ctx).ScribeSessionRequest(scribeSessionRequest).Execute()

Start an embedded Scribe session



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
	scribeSessionRequest := *openapiclient.NewScribeSessionRequest("https://app.example.com") // ScribeSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingAPI.ScribeSession(context.Background()).ScribeSessionRequest(scribeSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingAPI.ScribeSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScribeSession`: BuilderSession
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingAPI.ScribeSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScribeSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scribeSessionRequest** | [**ScribeSessionRequest**](ScribeSessionRequest.md) |  | 

### Return type

[**BuilderSession**](BuilderSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

