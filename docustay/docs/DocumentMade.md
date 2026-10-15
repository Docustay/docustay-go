# DocumentMade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to [**DocumentMadeApproval**](DocumentMadeApproval.md) |  | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**Id** | **string** |  | 
**State** | **string** |  | 
**Roles** | Pointer to [**[]TemplateSeatsInner**](TemplateSeatsInner.md) |  | [optional] 
**Emailed** | Pointer to **int32** |  | [optional] 
**NotEmailed** | Pointer to **[]interface{}** |  | [optional] 
**Warnings** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewDocumentMade

`func NewDocumentMade(id string, state string, ) *DocumentMade`

NewDocumentMade instantiates a new DocumentMade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentMadeWithDefaults

`func NewDocumentMadeWithDefaults() *DocumentMade`

NewDocumentMadeWithDefaults instantiates a new DocumentMade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *DocumentMade) GetApproval() DocumentMadeApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *DocumentMade) GetApprovalOk() (*DocumentMadeApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *DocumentMade) SetApproval(v DocumentMadeApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *DocumentMade) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetExternalId

`func (o *DocumentMade) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DocumentMade) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DocumentMade) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *DocumentMade) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMetadata

`func (o *DocumentMade) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentMade) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentMade) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DocumentMade) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *DocumentMade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentMade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentMade) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *DocumentMade) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DocumentMade) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DocumentMade) SetState(v string)`

SetState sets State field to given value.


### GetRoles

`func (o *DocumentMade) GetRoles() []TemplateSeatsInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DocumentMade) GetRolesOk() (*[]TemplateSeatsInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DocumentMade) SetRoles(v []TemplateSeatsInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DocumentMade) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetEmailed

`func (o *DocumentMade) GetEmailed() int32`

GetEmailed returns the Emailed field if non-nil, zero value otherwise.

### GetEmailedOk

`func (o *DocumentMade) GetEmailedOk() (*int32, bool)`

GetEmailedOk returns a tuple with the Emailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailed

`func (o *DocumentMade) SetEmailed(v int32)`

SetEmailed sets Emailed field to given value.

### HasEmailed

`func (o *DocumentMade) HasEmailed() bool`

HasEmailed returns a boolean if a field has been set.

### GetNotEmailed

`func (o *DocumentMade) GetNotEmailed() []interface{}`

GetNotEmailed returns the NotEmailed field if non-nil, zero value otherwise.

### GetNotEmailedOk

`func (o *DocumentMade) GetNotEmailedOk() (*[]interface{}, bool)`

GetNotEmailedOk returns a tuple with the NotEmailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEmailed

`func (o *DocumentMade) SetNotEmailed(v []interface{})`

SetNotEmailed sets NotEmailed field to given value.

### HasNotEmailed

`func (o *DocumentMade) HasNotEmailed() bool`

HasNotEmailed returns a boolean if a field has been set.

### GetWarnings

`func (o *DocumentMade) GetWarnings() []interface{}`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DocumentMade) GetWarningsOk() (*[]interface{}, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DocumentMade) SetWarnings(v []interface{})`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DocumentMade) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


