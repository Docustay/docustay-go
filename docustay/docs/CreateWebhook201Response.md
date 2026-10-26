# CreateWebhook201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Url** | **NullableString** | null for a relay endpoint | 
**Relay** | **bool** |  | 
**EventTypes** | **[]string** |  | 
**Status** | **string** |  | 
**DisabledReason** | **NullableString** |  | 
**ConsecutiveFailures** | **int32** |  | 
**Fingerprint** | **string** |  | 
**CreatedAt** | **string** |  | 
**Secret** | **string** |  | 

## Methods

### NewCreateWebhook201Response

`func NewCreateWebhook201Response(id string, name string, url NullableString, relay bool, eventTypes []string, status string, disabledReason NullableString, consecutiveFailures int32, fingerprint string, createdAt string, secret string, ) *CreateWebhook201Response`

NewCreateWebhook201Response instantiates a new CreateWebhook201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhook201ResponseWithDefaults

`func NewCreateWebhook201ResponseWithDefaults() *CreateWebhook201Response`

NewCreateWebhook201ResponseWithDefaults instantiates a new CreateWebhook201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateWebhook201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateWebhook201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateWebhook201Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateWebhook201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhook201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhook201Response) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateWebhook201Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhook201Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhook201Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *CreateWebhook201Response) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateWebhook201Response) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRelay

`func (o *CreateWebhook201Response) GetRelay() bool`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *CreateWebhook201Response) GetRelayOk() (*bool, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *CreateWebhook201Response) SetRelay(v bool)`

SetRelay sets Relay field to given value.


### GetEventTypes

`func (o *CreateWebhook201Response) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *CreateWebhook201Response) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *CreateWebhook201Response) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetStatus

`func (o *CreateWebhook201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateWebhook201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateWebhook201Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDisabledReason

`func (o *CreateWebhook201Response) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *CreateWebhook201Response) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *CreateWebhook201Response) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.


### SetDisabledReasonNil

`func (o *CreateWebhook201Response) SetDisabledReasonNil(b bool)`

 SetDisabledReasonNil sets the value for DisabledReason to be an explicit nil

### UnsetDisabledReason
`func (o *CreateWebhook201Response) UnsetDisabledReason()`

UnsetDisabledReason ensures that no value is present for DisabledReason, not even an explicit nil
### GetConsecutiveFailures

`func (o *CreateWebhook201Response) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *CreateWebhook201Response) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *CreateWebhook201Response) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetFingerprint

`func (o *CreateWebhook201Response) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CreateWebhook201Response) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CreateWebhook201Response) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetCreatedAt

`func (o *CreateWebhook201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateWebhook201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateWebhook201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetSecret

`func (o *CreateWebhook201Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateWebhook201Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateWebhook201Response) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


