# TemplateMade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**TemplateId** | **string** |  | 
**Fields** | **int32** |  | 
**Warnings** | [**[]TemplateMadeWarningsInner**](TemplateMadeWarningsInner.md) |  | 
**Detected** | Pointer to [**TemplateMadeDetected**](TemplateMadeDetected.md) |  | [optional] 

## Methods

### NewTemplateMade

`func NewTemplateMade(templateId string, fields int32, warnings []TemplateMadeWarningsInner, ) *TemplateMade`

NewTemplateMade instantiates a new TemplateMade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateMadeWithDefaults

`func NewTemplateMadeWithDefaults() *TemplateMade`

NewTemplateMadeWithDefaults instantiates a new TemplateMade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *TemplateMade) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TemplateMade) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TemplateMade) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TemplateMade) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMetadata

`func (o *TemplateMade) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TemplateMade) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TemplateMade) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TemplateMade) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplateMade) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateMade) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateMade) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetFields

`func (o *TemplateMade) GetFields() int32`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TemplateMade) GetFieldsOk() (*int32, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TemplateMade) SetFields(v int32)`

SetFields sets Fields field to given value.


### GetWarnings

`func (o *TemplateMade) GetWarnings() []TemplateMadeWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TemplateMade) GetWarningsOk() (*[]TemplateMadeWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TemplateMade) SetWarnings(v []TemplateMadeWarningsInner)`

SetWarnings sets Warnings field to given value.


### GetDetected

`func (o *TemplateMade) GetDetected() TemplateMadeDetected`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *TemplateMade) GetDetectedOk() (*TemplateMadeDetected, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *TemplateMade) SetDetected(v TemplateMadeDetected)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *TemplateMade) HasDetected() bool`

HasDetected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


