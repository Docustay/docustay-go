# AiUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Plan** | **string** |  | 
**Allowance** | **float32** |  | 
**Used** | **float32** |  | 
**Remaining** | **float32** |  | 
**Percent** | **float32** |  | 
**State** | **string** |  | 
**ResetsAt** | **string** |  | 
**ByFeature** | [**map[string]AiUsageByFeatureValue**](AiUsageByFeatureValue.md) |  | 
**ThinkHarder** | [**AiUsageThinkHarder**](AiUsageThinkHarder.md) |  | 
**Paused** | **bool** |  | 

## Methods

### NewAiUsage

`func NewAiUsage(mode string, plan string, allowance float32, used float32, remaining float32, percent float32, state string, resetsAt string, byFeature map[string]AiUsageByFeatureValue, thinkHarder AiUsageThinkHarder, paused bool, ) *AiUsage`

NewAiUsage instantiates a new AiUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiUsageWithDefaults

`func NewAiUsageWithDefaults() *AiUsage`

NewAiUsageWithDefaults instantiates a new AiUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AiUsage) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AiUsage) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AiUsage) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPlan

`func (o *AiUsage) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AiUsage) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AiUsage) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetAllowance

`func (o *AiUsage) GetAllowance() float32`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *AiUsage) GetAllowanceOk() (*float32, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *AiUsage) SetAllowance(v float32)`

SetAllowance sets Allowance field to given value.


### GetUsed

`func (o *AiUsage) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AiUsage) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AiUsage) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetRemaining

`func (o *AiUsage) GetRemaining() float32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *AiUsage) GetRemainingOk() (*float32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *AiUsage) SetRemaining(v float32)`

SetRemaining sets Remaining field to given value.


### GetPercent

`func (o *AiUsage) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *AiUsage) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *AiUsage) SetPercent(v float32)`

SetPercent sets Percent field to given value.


### GetState

`func (o *AiUsage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiUsage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiUsage) SetState(v string)`

SetState sets State field to given value.


### GetResetsAt

`func (o *AiUsage) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *AiUsage) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *AiUsage) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.


### GetByFeature

`func (o *AiUsage) GetByFeature() map[string]AiUsageByFeatureValue`

GetByFeature returns the ByFeature field if non-nil, zero value otherwise.

### GetByFeatureOk

`func (o *AiUsage) GetByFeatureOk() (*map[string]AiUsageByFeatureValue, bool)`

GetByFeatureOk returns a tuple with the ByFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByFeature

`func (o *AiUsage) SetByFeature(v map[string]AiUsageByFeatureValue)`

SetByFeature sets ByFeature field to given value.


### GetThinkHarder

`func (o *AiUsage) GetThinkHarder() AiUsageThinkHarder`

GetThinkHarder returns the ThinkHarder field if non-nil, zero value otherwise.

### GetThinkHarderOk

`func (o *AiUsage) GetThinkHarderOk() (*AiUsageThinkHarder, bool)`

GetThinkHarderOk returns a tuple with the ThinkHarder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinkHarder

`func (o *AiUsage) SetThinkHarder(v AiUsageThinkHarder)`

SetThinkHarder sets ThinkHarder field to given value.


### GetPaused

`func (o *AiUsage) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *AiUsage) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *AiUsage) SetPaused(v bool)`

SetPaused sets Paused field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


