# UseTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailChoice**](EmailChoice.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**People** | [**[]Person**](Person.md) |  | 
**VariableValues** | Pointer to **map[string]string** |  | [optional] 
**Facts** | Pointer to **map[string]string** |  | [optional] 
**Send** | Pointer to [**SendChoices**](SendChoices.md) |  | [optional] 
**SendNow** | Pointer to **bool** |  | [optional] 

## Methods

### NewUseTemplateRequest

`func NewUseTemplateRequest(people []Person, ) *UseTemplateRequest`

NewUseTemplateRequest instantiates a new UseTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUseTemplateRequestWithDefaults

`func NewUseTemplateRequestWithDefaults() *UseTemplateRequest`

NewUseTemplateRequestWithDefaults instantiates a new UseTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UseTemplateRequest) GetEmail() EmailChoice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UseTemplateRequest) GetEmailOk() (*EmailChoice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UseTemplateRequest) SetEmail(v EmailChoice)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UseTemplateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *UseTemplateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UseTemplateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UseTemplateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UseTemplateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *UseTemplateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UseTemplateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UseTemplateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UseTemplateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTitle

`func (o *UseTemplateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UseTemplateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UseTemplateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UseTemplateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPeople

`func (o *UseTemplateRequest) GetPeople() []Person`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *UseTemplateRequest) GetPeopleOk() (*[]Person, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *UseTemplateRequest) SetPeople(v []Person)`

SetPeople sets People field to given value.


### GetVariableValues

`func (o *UseTemplateRequest) GetVariableValues() map[string]string`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *UseTemplateRequest) GetVariableValuesOk() (*map[string]string, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *UseTemplateRequest) SetVariableValues(v map[string]string)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *UseTemplateRequest) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetFacts

`func (o *UseTemplateRequest) GetFacts() map[string]string`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *UseTemplateRequest) GetFactsOk() (*map[string]string, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *UseTemplateRequest) SetFacts(v map[string]string)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *UseTemplateRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetSend

`func (o *UseTemplateRequest) GetSend() SendChoices`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *UseTemplateRequest) GetSendOk() (*SendChoices, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *UseTemplateRequest) SetSend(v SendChoices)`

SetSend sets Send field to given value.

### HasSend

`func (o *UseTemplateRequest) HasSend() bool`

HasSend returns a boolean if a field has been set.

### GetSendNow

`func (o *UseTemplateRequest) GetSendNow() bool`

GetSendNow returns the SendNow field if non-nil, zero value otherwise.

### GetSendNowOk

`func (o *UseTemplateRequest) GetSendNowOk() (*bool, bool)`

GetSendNowOk returns a tuple with the SendNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNow

`func (o *UseTemplateRequest) SetSendNow(v bool)`

SetSendNow sets SendNow field to given value.

### HasSendNow

`func (o *UseTemplateRequest) HasSendNow() bool`

HasSendNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


