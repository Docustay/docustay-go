# Party

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | **string** |  | 
**ExternalId** | **NullableString** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | 
**Metadata** | **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | 
**Name** | **string** |  | 
**Email** | **NullableString** |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 
**ViewedAt** | **NullableString** |  | 
**CompletedAt** | **NullableString** |  | 
**DeclineReason** | **NullableString** |  | 

## Methods

### NewParty

`func NewParty(seat string, externalId NullableString, metadata map[string]string, name string, email NullableString, role string, status string, viewedAt NullableString, completedAt NullableString, declineReason NullableString, ) *Party`

NewParty instantiates a new Party object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyWithDefaults

`func NewPartyWithDefaults() *Party`

NewPartyWithDefaults instantiates a new Party object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *Party) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *Party) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *Party) SetSeat(v string)`

SetSeat sets Seat field to given value.


### GetExternalId

`func (o *Party) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Party) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Party) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *Party) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Party) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMetadata

`func (o *Party) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Party) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Party) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *Party) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Party) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Party) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *Party) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Party) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Party) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *Party) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Party) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRole

`func (o *Party) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Party) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Party) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *Party) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Party) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Party) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetViewedAt

`func (o *Party) GetViewedAt() string`

GetViewedAt returns the ViewedAt field if non-nil, zero value otherwise.

### GetViewedAtOk

`func (o *Party) GetViewedAtOk() (*string, bool)`

GetViewedAtOk returns a tuple with the ViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedAt

`func (o *Party) SetViewedAt(v string)`

SetViewedAt sets ViewedAt field to given value.


### SetViewedAtNil

`func (o *Party) SetViewedAtNil(b bool)`

 SetViewedAtNil sets the value for ViewedAt to be an explicit nil

### UnsetViewedAt
`func (o *Party) UnsetViewedAt()`

UnsetViewedAt ensures that no value is present for ViewedAt, not even an explicit nil
### GetCompletedAt

`func (o *Party) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Party) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Party) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *Party) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Party) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDeclineReason

`func (o *Party) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *Party) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *Party) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.


### SetDeclineReasonNil

`func (o *Party) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *Party) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


