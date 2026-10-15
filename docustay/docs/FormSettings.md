# FormSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifyEmail** | Pointer to **bool** |  | [optional] 
**CopyTo** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableString** |  | [optional] 
**ResponseLimit** | Pointer to **NullableInt32** |  | [optional] 
**Intro** | Pointer to **NullableString** |  | [optional] 
**ThankYou** | Pointer to **NullableString** |  | [optional] 
**FillAgain** | Pointer to **bool** |  | [optional] 
**EmbedSites** | Pointer to **[]string** |  | [optional] 
**NotifyOwner** | Pointer to **bool** |  | [optional] 
**ConfirmRespondent** | Pointer to **bool** |  | [optional] 
**CounterSign** | Pointer to [**FormSettingsCounterSign**](FormSettingsCounterSign.md) |  | [optional] 

## Methods

### NewFormSettings

`func NewFormSettings() *FormSettings`

NewFormSettings instantiates a new FormSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSettingsWithDefaults

`func NewFormSettingsWithDefaults() *FormSettings`

NewFormSettingsWithDefaults instantiates a new FormSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyEmail

`func (o *FormSettings) GetVerifyEmail() bool`

GetVerifyEmail returns the VerifyEmail field if non-nil, zero value otherwise.

### GetVerifyEmailOk

`func (o *FormSettings) GetVerifyEmailOk() (*bool, bool)`

GetVerifyEmailOk returns a tuple with the VerifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyEmail

`func (o *FormSettings) SetVerifyEmail(v bool)`

SetVerifyEmail sets VerifyEmail field to given value.

### HasVerifyEmail

`func (o *FormSettings) HasVerifyEmail() bool`

HasVerifyEmail returns a boolean if a field has been set.

### GetCopyTo

`func (o *FormSettings) GetCopyTo() string`

GetCopyTo returns the CopyTo field if non-nil, zero value otherwise.

### GetCopyToOk

`func (o *FormSettings) GetCopyToOk() (*string, bool)`

GetCopyToOk returns a tuple with the CopyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTo

`func (o *FormSettings) SetCopyTo(v string)`

SetCopyTo sets CopyTo field to given value.

### HasCopyTo

`func (o *FormSettings) HasCopyTo() bool`

HasCopyTo returns a boolean if a field has been set.

### SetCopyToNil

`func (o *FormSettings) SetCopyToNil(b bool)`

 SetCopyToNil sets the value for CopyTo to be an explicit nil

### UnsetCopyTo
`func (o *FormSettings) UnsetCopyTo()`

UnsetCopyTo ensures that no value is present for CopyTo, not even an explicit nil
### GetExpiresAt

`func (o *FormSettings) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FormSettings) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FormSettings) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FormSettings) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *FormSettings) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *FormSettings) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetResponseLimit

`func (o *FormSettings) GetResponseLimit() int32`

GetResponseLimit returns the ResponseLimit field if non-nil, zero value otherwise.

### GetResponseLimitOk

`func (o *FormSettings) GetResponseLimitOk() (*int32, bool)`

GetResponseLimitOk returns a tuple with the ResponseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLimit

`func (o *FormSettings) SetResponseLimit(v int32)`

SetResponseLimit sets ResponseLimit field to given value.

### HasResponseLimit

`func (o *FormSettings) HasResponseLimit() bool`

HasResponseLimit returns a boolean if a field has been set.

### SetResponseLimitNil

`func (o *FormSettings) SetResponseLimitNil(b bool)`

 SetResponseLimitNil sets the value for ResponseLimit to be an explicit nil

### UnsetResponseLimit
`func (o *FormSettings) UnsetResponseLimit()`

UnsetResponseLimit ensures that no value is present for ResponseLimit, not even an explicit nil
### GetIntro

`func (o *FormSettings) GetIntro() string`

GetIntro returns the Intro field if non-nil, zero value otherwise.

### GetIntroOk

`func (o *FormSettings) GetIntroOk() (*string, bool)`

GetIntroOk returns a tuple with the Intro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntro

`func (o *FormSettings) SetIntro(v string)`

SetIntro sets Intro field to given value.

### HasIntro

`func (o *FormSettings) HasIntro() bool`

HasIntro returns a boolean if a field has been set.

### SetIntroNil

`func (o *FormSettings) SetIntroNil(b bool)`

 SetIntroNil sets the value for Intro to be an explicit nil

### UnsetIntro
`func (o *FormSettings) UnsetIntro()`

UnsetIntro ensures that no value is present for Intro, not even an explicit nil
### GetThankYou

`func (o *FormSettings) GetThankYou() string`

GetThankYou returns the ThankYou field if non-nil, zero value otherwise.

### GetThankYouOk

`func (o *FormSettings) GetThankYouOk() (*string, bool)`

GetThankYouOk returns a tuple with the ThankYou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYou

`func (o *FormSettings) SetThankYou(v string)`

SetThankYou sets ThankYou field to given value.

### HasThankYou

`func (o *FormSettings) HasThankYou() bool`

HasThankYou returns a boolean if a field has been set.

### SetThankYouNil

`func (o *FormSettings) SetThankYouNil(b bool)`

 SetThankYouNil sets the value for ThankYou to be an explicit nil

### UnsetThankYou
`func (o *FormSettings) UnsetThankYou()`

UnsetThankYou ensures that no value is present for ThankYou, not even an explicit nil
### GetFillAgain

`func (o *FormSettings) GetFillAgain() bool`

GetFillAgain returns the FillAgain field if non-nil, zero value otherwise.

### GetFillAgainOk

`func (o *FormSettings) GetFillAgainOk() (*bool, bool)`

GetFillAgainOk returns a tuple with the FillAgain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillAgain

`func (o *FormSettings) SetFillAgain(v bool)`

SetFillAgain sets FillAgain field to given value.

### HasFillAgain

`func (o *FormSettings) HasFillAgain() bool`

HasFillAgain returns a boolean if a field has been set.

### GetEmbedSites

`func (o *FormSettings) GetEmbedSites() []string`

GetEmbedSites returns the EmbedSites field if non-nil, zero value otherwise.

### GetEmbedSitesOk

`func (o *FormSettings) GetEmbedSitesOk() (*[]string, bool)`

GetEmbedSitesOk returns a tuple with the EmbedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedSites

`func (o *FormSettings) SetEmbedSites(v []string)`

SetEmbedSites sets EmbedSites field to given value.

### HasEmbedSites

`func (o *FormSettings) HasEmbedSites() bool`

HasEmbedSites returns a boolean if a field has been set.

### GetNotifyOwner

`func (o *FormSettings) GetNotifyOwner() bool`

GetNotifyOwner returns the NotifyOwner field if non-nil, zero value otherwise.

### GetNotifyOwnerOk

`func (o *FormSettings) GetNotifyOwnerOk() (*bool, bool)`

GetNotifyOwnerOk returns a tuple with the NotifyOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOwner

`func (o *FormSettings) SetNotifyOwner(v bool)`

SetNotifyOwner sets NotifyOwner field to given value.

### HasNotifyOwner

`func (o *FormSettings) HasNotifyOwner() bool`

HasNotifyOwner returns a boolean if a field has been set.

### GetConfirmRespondent

`func (o *FormSettings) GetConfirmRespondent() bool`

GetConfirmRespondent returns the ConfirmRespondent field if non-nil, zero value otherwise.

### GetConfirmRespondentOk

`func (o *FormSettings) GetConfirmRespondentOk() (*bool, bool)`

GetConfirmRespondentOk returns a tuple with the ConfirmRespondent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmRespondent

`func (o *FormSettings) SetConfirmRespondent(v bool)`

SetConfirmRespondent sets ConfirmRespondent field to given value.

### HasConfirmRespondent

`func (o *FormSettings) HasConfirmRespondent() bool`

HasConfirmRespondent returns a boolean if a field has been set.

### GetCounterSign

`func (o *FormSettings) GetCounterSign() FormSettingsCounterSign`

GetCounterSign returns the CounterSign field if non-nil, zero value otherwise.

### GetCounterSignOk

`func (o *FormSettings) GetCounterSignOk() (*FormSettingsCounterSign, bool)`

GetCounterSignOk returns a tuple with the CounterSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterSign

`func (o *FormSettings) SetCounterSign(v FormSettingsCounterSign)`

SetCounterSign sets CounterSign field to given value.

### HasCounterSign

`func (o *FormSettings) HasCounterSign() bool`

HasCounterSign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


