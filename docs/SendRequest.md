# SendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** |  | 
**Email** | Pointer to [**EmailChoice**](EmailChoice.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**ExpiryDays** | Pointer to **int32** |  | [optional] 
**Embedded** | Pointer to **[]string** | Seats signed inside your own app instead of by emailed link (for example [\&quot;signer_1\&quot;]). Nobody is emailed for these seats: you ask for a session with POST /api/v1/documents/{id}/embed-session and show it on a website you listed under Settings → Developers → Embedding. | [optional] 
**Parties** | [**[]SendRequestPartiesInner**](SendRequestPartiesInner.md) | One entry for each seat of the template (see GET /api/v1/documents/templates). | 

## Methods

### NewSendRequest

`func NewSendRequest(templateId string, parties []SendRequestPartiesInner, ) *SendRequest`

NewSendRequest instantiates a new SendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendRequestWithDefaults

`func NewSendRequestWithDefaults() *SendRequest`

NewSendRequestWithDefaults instantiates a new SendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *SendRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetEmail

`func (o *SendRequest) GetEmail() EmailChoice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendRequest) GetEmailOk() (*EmailChoice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendRequest) SetEmail(v EmailChoice)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SendRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *SendRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SendRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SendRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SendRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *SendRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SendRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SendRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SendRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTitle

`func (o *SendRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SendRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessage

`func (o *SendRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SendRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SendRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSubject

`func (o *SendRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetExpiryDays

`func (o *SendRequest) GetExpiryDays() int32`

GetExpiryDays returns the ExpiryDays field if non-nil, zero value otherwise.

### GetExpiryDaysOk

`func (o *SendRequest) GetExpiryDaysOk() (*int32, bool)`

GetExpiryDaysOk returns a tuple with the ExpiryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDays

`func (o *SendRequest) SetExpiryDays(v int32)`

SetExpiryDays sets ExpiryDays field to given value.

### HasExpiryDays

`func (o *SendRequest) HasExpiryDays() bool`

HasExpiryDays returns a boolean if a field has been set.

### GetEmbedded

`func (o *SendRequest) GetEmbedded() []string`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SendRequest) GetEmbeddedOk() (*[]string, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SendRequest) SetEmbedded(v []string)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SendRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetParties

`func (o *SendRequest) GetParties() []SendRequestPartiesInner`

GetParties returns the Parties field if non-nil, zero value otherwise.

### GetPartiesOk

`func (o *SendRequest) GetPartiesOk() (*[]SendRequestPartiesInner, bool)`

GetPartiesOk returns a tuple with the Parties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParties

`func (o *SendRequest) SetParties(v []SendRequestPartiesInner)`

SetParties sets Parties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


