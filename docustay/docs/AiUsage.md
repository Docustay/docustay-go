# AiUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Plan** | **NullableString** |  | 
**Allowance** | **NullableFloat32** |  | 
**Used** | **NullableFloat32** |  | 
**Remaining** | **NullableFloat32** |  | 
**Percent** | **NullableFloat32** |  | 
**State** | **string** |  | 
**ResetsAt** | **NullableString** |  | 
**ByFeature** | [**map[string]AiUsageByFeatureValue**](AiUsageByFeatureValue.md) |  | 
**ThinkHarder** | [**AiUsageThinkHarder**](AiUsageThinkHarder.md) |  | 
**Paused** | **bool** |  | 

## Methods

### NewAiUsage

`func NewAiUsage(mode string, plan NullableString, allowance NullableFloat32, used NullableFloat32, remaining NullableFloat32, percent NullableFloat32, state string, resetsAt NullableString, byFeature map[string]AiUsageByFeatureValue, thinkHarder AiUsageThinkHarder, paused bool, ) *AiUsage`

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


### SetPlanNil

`func (o *AiUsage) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *AiUsage) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
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


### SetAllowanceNil

`func (o *AiUsage) SetAllowanceNil(b bool)`

 SetAllowanceNil sets the value for Allowance to be an explicit nil

### UnsetAllowance
`func (o *AiUsage) UnsetAllowance()`

UnsetAllowance ensures that no value is present for Allowance, not even an explicit nil
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


### SetUsedNil

`func (o *AiUsage) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *AiUsage) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil
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


### SetRemainingNil

`func (o *AiUsage) SetRemainingNil(b bool)`

 SetRemainingNil sets the value for Remaining to be an explicit nil

### UnsetRemaining
`func (o *AiUsage) UnsetRemaining()`

UnsetRemaining ensures that no value is present for Remaining, not even an explicit nil
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


### SetPercentNil

`func (o *AiUsage) SetPercentNil(b bool)`

 SetPercentNil sets the value for Percent to be an explicit nil

### UnsetPercent
`func (o *AiUsage) UnsetPercent()`

UnsetPercent ensures that no value is present for Percent, not even an explicit nil
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


### SetResetsAtNil

`func (o *AiUsage) SetResetsAtNil(b bool)`

 SetResetsAtNil sets the value for ResetsAt to be an explicit nil

### UnsetResetsAt
`func (o *AiUsage) UnsetResetsAt()`

UnsetResetsAt ensures that no value is present for ResetsAt, not even an explicit nil
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


