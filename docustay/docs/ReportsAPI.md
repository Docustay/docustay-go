# \ReportsAPI

All URIs are relative to *https://app.docustay.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Report**](ReportsAPI.md#Report) | **Get** /api/v1/reports | Counts of what was sent and completed



## Report

> Report Report(ctx).From(from).To(to).Team(team).Execute()

Counts of what was sent and completed



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
	from := "2026-10-01" // string |  (optional)
	to := "2026-10-31" // string |  (optional)
	team := "6f1c1f0e-7a0c-4d1d-9b5e-2f0a7b1c3d4e" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.Report(context.Background()).From(from).To(to).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.Report``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Report`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.Report`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** |  | 
 **to** | **string** |  | 
 **team** | **string** |  | 

### Return type

[**Report**](Report.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

