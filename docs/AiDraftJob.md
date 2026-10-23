# AiDraftJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Patches** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAiDraftJob

`func NewAiDraftJob(id string, state string, ) *AiDraftJob`

NewAiDraftJob instantiates a new AiDraftJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDraftJobWithDefaults

`func NewAiDraftJobWithDefaults() *AiDraftJob`

NewAiDraftJobWithDefaults instantiates a new AiDraftJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiDraftJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiDraftJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiDraftJob) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *AiDraftJob) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiDraftJob) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiDraftJob) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *AiDraftJob) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiDraftJob) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiDraftJob) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AiDraftJob) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSource

`func (o *AiDraftJob) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AiDraftJob) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AiDraftJob) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *AiDraftJob) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPatches

`func (o *AiDraftJob) GetPatches() map[string]interface{}`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *AiDraftJob) GetPatchesOk() (*map[string]interface{}, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *AiDraftJob) SetPatches(v map[string]interface{})`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *AiDraftJob) HasPatches() bool`

HasPatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


