# PublishFormRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** |  | 
**Slug** | **string** |  | 
**Settings** | Pointer to [**FormSettings**](FormSettings.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPublishFormRequest

`func NewPublishFormRequest(templateId string, slug string, ) *PublishFormRequest`

NewPublishFormRequest instantiates a new PublishFormRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishFormRequestWithDefaults

`func NewPublishFormRequestWithDefaults() *PublishFormRequest`

NewPublishFormRequestWithDefaults instantiates a new PublishFormRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *PublishFormRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PublishFormRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PublishFormRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetSlug

`func (o *PublishFormRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PublishFormRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PublishFormRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSettings

`func (o *PublishFormRequest) GetSettings() FormSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PublishFormRequest) GetSettingsOk() (*FormSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PublishFormRequest) SetSettings(v FormSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PublishFormRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetEnabled

`func (o *PublishFormRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PublishFormRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PublishFormRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PublishFormRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


