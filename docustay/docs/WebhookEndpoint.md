# WebhookEndpoint

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

## Methods

### NewWebhookEndpoint

`func NewWebhookEndpoint(id string, name string, url NullableString, relay bool, eventTypes []string, status string, disabledReason NullableString, consecutiveFailures int32, fingerprint string, createdAt string, ) *WebhookEndpoint`

NewWebhookEndpoint instantiates a new WebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointWithDefaults

`func NewWebhookEndpointWithDefaults() *WebhookEndpoint`

NewWebhookEndpointWithDefaults instantiates a new WebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebhookEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *WebhookEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *WebhookEndpoint) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebhookEndpoint) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRelay

`func (o *WebhookEndpoint) GetRelay() bool`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *WebhookEndpoint) GetRelayOk() (*bool, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *WebhookEndpoint) SetRelay(v bool)`

SetRelay sets Relay field to given value.


### GetEventTypes

`func (o *WebhookEndpoint) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookEndpoint) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookEndpoint) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetStatus

`func (o *WebhookEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDisabledReason

`func (o *WebhookEndpoint) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *WebhookEndpoint) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *WebhookEndpoint) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.


### SetDisabledReasonNil

`func (o *WebhookEndpoint) SetDisabledReasonNil(b bool)`

 SetDisabledReasonNil sets the value for DisabledReason to be an explicit nil

### UnsetDisabledReason
`func (o *WebhookEndpoint) UnsetDisabledReason()`

UnsetDisabledReason ensures that no value is present for DisabledReason, not even an explicit nil
### GetConsecutiveFailures

`func (o *WebhookEndpoint) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *WebhookEndpoint) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *WebhookEndpoint) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetFingerprint

`func (o *WebhookEndpoint) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *WebhookEndpoint) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *WebhookEndpoint) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetCreatedAt

`func (o *WebhookEndpoint) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEndpoint) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEndpoint) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


