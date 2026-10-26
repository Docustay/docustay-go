# SendDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **map[string]string** | The boxes the sender fills in (\&quot;Filled in by: Sender\&quot;) and the template&#39;s send-time amounts, by key: a box&#39;s label in lower case with underscores (\&quot;Effective date\&quot; → &#x60;effective_date&#x60;). A send that leaves one empty answers 400 naming it. | [optional] 

## Methods

### NewSendDraftRequest

`func NewSendDraftRequest() *SendDraftRequest`

NewSendDraftRequest instantiates a new SendDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDraftRequestWithDefaults

`func NewSendDraftRequestWithDefaults() *SendDraftRequest`

NewSendDraftRequestWithDefaults instantiates a new SendDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *SendDraftRequest) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SendDraftRequest) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SendDraftRequest) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SendDraftRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


