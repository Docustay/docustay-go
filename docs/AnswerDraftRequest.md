# AnswerDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | [**map[string]StartDraftRequestAnswersValue**](StartDraftRequestAnswersValue.md) |  | 
**UseDefaults** | Pointer to **bool** |  | [optional] 

## Methods

### NewAnswerDraftRequest

`func NewAnswerDraftRequest(answers map[string]StartDraftRequestAnswersValue, ) *AnswerDraftRequest`

NewAnswerDraftRequest instantiates a new AnswerDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerDraftRequestWithDefaults

`func NewAnswerDraftRequestWithDefaults() *AnswerDraftRequest`

NewAnswerDraftRequestWithDefaults instantiates a new AnswerDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *AnswerDraftRequest) GetAnswers() map[string]StartDraftRequestAnswersValue`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *AnswerDraftRequest) GetAnswersOk() (*map[string]StartDraftRequestAnswersValue, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *AnswerDraftRequest) SetAnswers(v map[string]StartDraftRequestAnswersValue)`

SetAnswers sets Answers field to given value.


### GetUseDefaults

`func (o *AnswerDraftRequest) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *AnswerDraftRequest) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *AnswerDraftRequest) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *AnswerDraftRequest) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


