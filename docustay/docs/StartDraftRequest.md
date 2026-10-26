# StartDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | **string** | A library template&#39;s slug (see &#x60;GET /api/v1/templates/library&#x60;). | 
**Title** | Pointer to **string** |  | [optional] 
**People** | [**[]StartDraftRequestPeopleInner**](StartDraftRequestPeopleInner.md) |  | 
**Answers** | Pointer to [**map[string]StartDraftRequestAnswersValue**](StartDraftRequestAnswersValue.md) |  | [optional] 
**Values** | Pointer to **map[string]string** | Send-time amounts the template&#39;s words need, by key (a &#x60;needs_input&#x60; answer lists them under &#x60;values&#x60;), e.g. &#x60;{ \&quot;deposit\&quot;: \&quot;$500\&quot; }&#x60;. | [optional] 
**UseDefaults** | Pointer to **bool** |  | [optional] 

## Methods

### NewStartDraftRequest

`func NewStartDraftRequest(template string, people []StartDraftRequestPeopleInner, ) *StartDraftRequest`

NewStartDraftRequest instantiates a new StartDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartDraftRequestWithDefaults

`func NewStartDraftRequestWithDefaults() *StartDraftRequest`

NewStartDraftRequestWithDefaults instantiates a new StartDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *StartDraftRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *StartDraftRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *StartDraftRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetTitle

`func (o *StartDraftRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StartDraftRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StartDraftRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StartDraftRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPeople

`func (o *StartDraftRequest) GetPeople() []StartDraftRequestPeopleInner`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *StartDraftRequest) GetPeopleOk() (*[]StartDraftRequestPeopleInner, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *StartDraftRequest) SetPeople(v []StartDraftRequestPeopleInner)`

SetPeople sets People field to given value.


### GetAnswers

`func (o *StartDraftRequest) GetAnswers() map[string]StartDraftRequestAnswersValue`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *StartDraftRequest) GetAnswersOk() (*map[string]StartDraftRequestAnswersValue, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *StartDraftRequest) SetAnswers(v map[string]StartDraftRequestAnswersValue)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *StartDraftRequest) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetValues

`func (o *StartDraftRequest) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StartDraftRequest) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StartDraftRequest) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StartDraftRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUseDefaults

`func (o *StartDraftRequest) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *StartDraftRequest) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *StartDraftRequest) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *StartDraftRequest) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


