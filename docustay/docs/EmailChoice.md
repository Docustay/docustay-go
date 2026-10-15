# EmailChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to **string** | Where replies to this document&#39;s emails go. | [optional] 
**SenderName** | Pointer to **string** | The name the emails say they are from, for this document. | [optional] 
**Locale** | Pointer to **string** | The language of the emails (en, es, fr, …), or auto to match each person. | [optional] 

## Methods

### NewEmailChoice

`func NewEmailChoice() *EmailChoice`

NewEmailChoice instantiates a new EmailChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailChoiceWithDefaults

`func NewEmailChoiceWithDefaults() *EmailChoice`

NewEmailChoiceWithDefaults instantiates a new EmailChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailChoice) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailChoice) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailChoice) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailChoice) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *EmailChoice) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailChoice) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailChoice) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EmailChoice) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailChoice) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailChoice) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailChoice) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailChoice) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSenderName

`func (o *EmailChoice) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *EmailChoice) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *EmailChoice) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *EmailChoice) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetLocale

`func (o *EmailChoice) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EmailChoice) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EmailChoice) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *EmailChoice) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


