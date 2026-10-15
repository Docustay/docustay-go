# ListDocuments200ResponsePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**NextCursor** | **string** |  | 
**HasMore** | **bool** |  | 

## Methods

### NewListDocuments200ResponsePage

`func NewListDocuments200ResponsePage(limit int32, nextCursor string, hasMore bool, ) *ListDocuments200ResponsePage`

NewListDocuments200ResponsePage instantiates a new ListDocuments200ResponsePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDocuments200ResponsePageWithDefaults

`func NewListDocuments200ResponsePageWithDefaults() *ListDocuments200ResponsePage`

NewListDocuments200ResponsePageWithDefaults instantiates a new ListDocuments200ResponsePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListDocuments200ResponsePage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListDocuments200ResponsePage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListDocuments200ResponsePage) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNextCursor

`func (o *ListDocuments200ResponsePage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListDocuments200ResponsePage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListDocuments200ResponsePage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetHasMore

`func (o *ListDocuments200ResponsePage) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListDocuments200ResponsePage) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListDocuments200ResponsePage) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


