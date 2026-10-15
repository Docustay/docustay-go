# ListDocuments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | [**[]DocumentSummary**](DocumentSummary.md) |  | 
**Page** | [**ListDocuments200ResponsePage**](ListDocuments200ResponsePage.md) |  | 

## Methods

### NewListDocuments200Response

`func NewListDocuments200Response(documents []DocumentSummary, page ListDocuments200ResponsePage, ) *ListDocuments200Response`

NewListDocuments200Response instantiates a new ListDocuments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDocuments200ResponseWithDefaults

`func NewListDocuments200ResponseWithDefaults() *ListDocuments200Response`

NewListDocuments200ResponseWithDefaults instantiates a new ListDocuments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *ListDocuments200Response) GetDocuments() []DocumentSummary`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ListDocuments200Response) GetDocumentsOk() (*[]DocumentSummary, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ListDocuments200Response) SetDocuments(v []DocumentSummary)`

SetDocuments sets Documents field to given value.


### GetPage

`func (o *ListDocuments200Response) GetPage() ListDocuments200ResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListDocuments200Response) GetPageOk() (*ListDocuments200ResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListDocuments200Response) SetPage(v ListDocuments200ResponsePage)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


