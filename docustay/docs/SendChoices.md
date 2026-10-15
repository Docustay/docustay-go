# SendChoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyTo** | Pointer to **string** |  | [optional] 
**SenderName** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ExpiryDays** | Pointer to **NullableInt32** |  | [optional] 
**ReminderDays** | Pointer to **[]int32** |  | [optional] 
**SigningOrder** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**InternalNote** | Pointer to **string** | A note only the sender sees; never in anything a signer gets. | [optional] 
**RequireDeclineReason** | Pointer to **bool** |  | [optional] 

## Methods

### NewSendChoices

`func NewSendChoices() *SendChoices`

NewSendChoices instantiates a new SendChoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendChoicesWithDefaults

`func NewSendChoicesWithDefaults() *SendChoices`

NewSendChoicesWithDefaults instantiates a new SendChoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyTo

`func (o *SendChoices) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *SendChoices) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *SendChoices) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *SendChoices) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSenderName

`func (o *SendChoices) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SendChoices) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SendChoices) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *SendChoices) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetSubject

`func (o *SendChoices) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendChoices) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendChoices) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendChoices) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *SendChoices) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendChoices) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendChoices) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendChoices) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExpiryDays

`func (o *SendChoices) GetExpiryDays() int32`

GetExpiryDays returns the ExpiryDays field if non-nil, zero value otherwise.

### GetExpiryDaysOk

`func (o *SendChoices) GetExpiryDaysOk() (*int32, bool)`

GetExpiryDaysOk returns a tuple with the ExpiryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDays

`func (o *SendChoices) SetExpiryDays(v int32)`

SetExpiryDays sets ExpiryDays field to given value.

### HasExpiryDays

`func (o *SendChoices) HasExpiryDays() bool`

HasExpiryDays returns a boolean if a field has been set.

### SetExpiryDaysNil

`func (o *SendChoices) SetExpiryDaysNil(b bool)`

 SetExpiryDaysNil sets the value for ExpiryDays to be an explicit nil

### UnsetExpiryDays
`func (o *SendChoices) UnsetExpiryDays()`

UnsetExpiryDays ensures that no value is present for ExpiryDays, not even an explicit nil
### GetReminderDays

`func (o *SendChoices) GetReminderDays() []int32`

GetReminderDays returns the ReminderDays field if non-nil, zero value otherwise.

### GetReminderDaysOk

`func (o *SendChoices) GetReminderDaysOk() (*[]int32, bool)`

GetReminderDaysOk returns a tuple with the ReminderDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDays

`func (o *SendChoices) SetReminderDays(v []int32)`

SetReminderDays sets ReminderDays field to given value.

### HasReminderDays

`func (o *SendChoices) HasReminderDays() bool`

HasReminderDays returns a boolean if a field has been set.

### GetSigningOrder

`func (o *SendChoices) GetSigningOrder() string`

GetSigningOrder returns the SigningOrder field if non-nil, zero value otherwise.

### GetSigningOrderOk

`func (o *SendChoices) GetSigningOrderOk() (*string, bool)`

GetSigningOrderOk returns a tuple with the SigningOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningOrder

`func (o *SendChoices) SetSigningOrder(v string)`

SetSigningOrder sets SigningOrder field to given value.

### HasSigningOrder

`func (o *SendChoices) HasSigningOrder() bool`

HasSigningOrder returns a boolean if a field has been set.

### GetLanguage

`func (o *SendChoices) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SendChoices) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SendChoices) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SendChoices) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetInternalNote

`func (o *SendChoices) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *SendChoices) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *SendChoices) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.

### HasInternalNote

`func (o *SendChoices) HasInternalNote() bool`

HasInternalNote returns a boolean if a field has been set.

### GetRequireDeclineReason

`func (o *SendChoices) GetRequireDeclineReason() bool`

GetRequireDeclineReason returns the RequireDeclineReason field if non-nil, zero value otherwise.

### GetRequireDeclineReasonOk

`func (o *SendChoices) GetRequireDeclineReasonOk() (*bool, bool)`

GetRequireDeclineReasonOk returns a tuple with the RequireDeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireDeclineReason

`func (o *SendChoices) SetRequireDeclineReason(v bool)`

SetRequireDeclineReason sets RequireDeclineReason field to given value.

### HasRequireDeclineReason

`func (o *SendChoices) HasRequireDeclineReason() bool`

HasRequireDeclineReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


