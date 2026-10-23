# ListEmails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]map[string]interface{}** |  | 
**NextCursor** | **string** |  | 

## Methods

### NewListEmails200Response

`func NewListEmails200Response(emails []map[string]interface{}, nextCursor string, ) *ListEmails200Response`

NewListEmails200Response instantiates a new ListEmails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEmails200ResponseWithDefaults

`func NewListEmails200ResponseWithDefaults() *ListEmails200Response`

NewListEmails200ResponseWithDefaults instantiates a new ListEmails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *ListEmails200Response) GetEmails() []map[string]interface{}`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ListEmails200Response) GetEmailsOk() (*[]map[string]interface{}, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ListEmails200Response) SetEmails(v []map[string]interface{})`

SetEmails sets Emails field to given value.


### GetNextCursor

`func (o *ListEmails200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListEmails200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListEmails200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


