# SetTemplateDefaultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**ClientDescription** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Send** | Pointer to [**SendChoices**](SendChoices.md) |  | [optional] 

## Methods

### NewSetTemplateDefaultsRequest

`func NewSetTemplateDefaultsRequest() *SetTemplateDefaultsRequest`

NewSetTemplateDefaultsRequest instantiates a new SetTemplateDefaultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetTemplateDefaultsRequestWithDefaults

`func NewSetTemplateDefaultsRequestWithDefaults() *SetTemplateDefaultsRequest`

NewSetTemplateDefaultsRequestWithDefaults instantiates a new SetTemplateDefaultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SetTemplateDefaultsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetTemplateDefaultsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetTemplateDefaultsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetTemplateDefaultsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SetTemplateDefaultsRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SetTemplateDefaultsRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetClientDescription

`func (o *SetTemplateDefaultsRequest) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *SetTemplateDefaultsRequest) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *SetTemplateDefaultsRequest) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.

### HasClientDescription

`func (o *SetTemplateDefaultsRequest) HasClientDescription() bool`

HasClientDescription returns a boolean if a field has been set.

### SetClientDescriptionNil

`func (o *SetTemplateDefaultsRequest) SetClientDescriptionNil(b bool)`

 SetClientDescriptionNil sets the value for ClientDescription to be an explicit nil

### UnsetClientDescription
`func (o *SetTemplateDefaultsRequest) UnsetClientDescription()`

UnsetClientDescription ensures that no value is present for ClientDescription, not even an explicit nil
### GetMessage

`func (o *SetTemplateDefaultsRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SetTemplateDefaultsRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SetTemplateDefaultsRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SetTemplateDefaultsRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SetTemplateDefaultsRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SetTemplateDefaultsRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSend

`func (o *SetTemplateDefaultsRequest) GetSend() SendChoices`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *SetTemplateDefaultsRequest) GetSendOk() (*SendChoices, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *SetTemplateDefaultsRequest) SetSend(v SendChoices)`

SetSend sets Send field to given value.

### HasSend

`func (o *SetTemplateDefaultsRequest) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


