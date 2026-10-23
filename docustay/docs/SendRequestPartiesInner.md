# SendRequestPartiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 

## Methods

### NewSendRequestPartiesInner

`func NewSendRequestPartiesInner(seat string, name string, email string, ) *SendRequestPartiesInner`

NewSendRequestPartiesInner instantiates a new SendRequestPartiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendRequestPartiesInnerWithDefaults

`func NewSendRequestPartiesInnerWithDefaults() *SendRequestPartiesInner`

NewSendRequestPartiesInnerWithDefaults instantiates a new SendRequestPartiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *SendRequestPartiesInner) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *SendRequestPartiesInner) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *SendRequestPartiesInner) SetSeat(v string)`

SetSeat sets Seat field to given value.


### GetName

`func (o *SendRequestPartiesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendRequestPartiesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendRequestPartiesInner) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *SendRequestPartiesInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendRequestPartiesInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendRequestPartiesInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *SendRequestPartiesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SendRequestPartiesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SendRequestPartiesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SendRequestPartiesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMetadata

`func (o *SendRequestPartiesInner) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SendRequestPartiesInner) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SendRequestPartiesInner) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SendRequestPartiesInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


