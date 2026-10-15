# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **NullableString** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | 
**Metadata** | **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | 
**Test** | **bool** |  | 
**Id** | **string** |  | 
**Title** | **string** |  | 
**State** | **string** |  | 
**SentAt** | **NullableString** |  | 
**ExpiresAt** | **NullableString** |  | 
**ExecutedAt** | **NullableString** |  | 
**VoidedAt** | **NullableString** |  | 
**VoidReason** | **NullableString** |  | 
**Progress** | [**TestSign200ResponseProgress**](TestSign200ResponseProgress.md) |  | 
**SignedCopyAvailable** | **bool** |  | 
**Parties** | [**[]Party**](Party.md) |  | 
**Payments** | [**[]Payment**](Payment.md) | Payments made through payment fields on this document (empty when it has none). A payment field&#39;s amount is set by ONE of: a fixed amount, a formula over the signer&#39;s own fields, a Stripe price, or a Stripe payment link — see the text tag settings &#x60;amount&#x60;, &#x60;formula&#x60;, &#x60;price&#x60; and &#x60;link&#x60;. | 
**Events** | [**[]DocumentEventsInner**](DocumentEventsInner.md) |  | 

## Methods

### NewDocument

`func NewDocument(externalId NullableString, metadata map[string]string, test bool, id string, title string, state string, sentAt NullableString, expiresAt NullableString, executedAt NullableString, voidedAt NullableString, voidReason NullableString, progress TestSign200ResponseProgress, signedCopyAvailable bool, parties []Party, payments []Payment, events []DocumentEventsInner, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *Document) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Document) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Document) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *Document) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Document) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMetadata

`func (o *Document) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Document) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Document) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetTest

`func (o *Document) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Document) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Document) SetTest(v bool)`

SetTest sets Test field to given value.


### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Document) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Document) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Document) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetState

`func (o *Document) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Document) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Document) SetState(v string)`

SetState sets State field to given value.


### GetSentAt

`func (o *Document) GetSentAt() string`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *Document) GetSentAtOk() (*string, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *Document) SetSentAt(v string)`

SetSentAt sets SentAt field to given value.


### SetSentAtNil

`func (o *Document) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *Document) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetExpiresAt

`func (o *Document) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Document) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Document) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Document) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Document) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetExecutedAt

`func (o *Document) GetExecutedAt() string`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *Document) GetExecutedAtOk() (*string, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *Document) SetExecutedAt(v string)`

SetExecutedAt sets ExecutedAt field to given value.


### SetExecutedAtNil

`func (o *Document) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *Document) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil
### GetVoidedAt

`func (o *Document) GetVoidedAt() string`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *Document) GetVoidedAtOk() (*string, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *Document) SetVoidedAt(v string)`

SetVoidedAt sets VoidedAt field to given value.


### SetVoidedAtNil

`func (o *Document) SetVoidedAtNil(b bool)`

 SetVoidedAtNil sets the value for VoidedAt to be an explicit nil

### UnsetVoidedAt
`func (o *Document) UnsetVoidedAt()`

UnsetVoidedAt ensures that no value is present for VoidedAt, not even an explicit nil
### GetVoidReason

`func (o *Document) GetVoidReason() string`

GetVoidReason returns the VoidReason field if non-nil, zero value otherwise.

### GetVoidReasonOk

`func (o *Document) GetVoidReasonOk() (*string, bool)`

GetVoidReasonOk returns a tuple with the VoidReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidReason

`func (o *Document) SetVoidReason(v string)`

SetVoidReason sets VoidReason field to given value.


### SetVoidReasonNil

`func (o *Document) SetVoidReasonNil(b bool)`

 SetVoidReasonNil sets the value for VoidReason to be an explicit nil

### UnsetVoidReason
`func (o *Document) UnsetVoidReason()`

UnsetVoidReason ensures that no value is present for VoidReason, not even an explicit nil
### GetProgress

`func (o *Document) GetProgress() TestSign200ResponseProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Document) GetProgressOk() (*TestSign200ResponseProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Document) SetProgress(v TestSign200ResponseProgress)`

SetProgress sets Progress field to given value.


### GetSignedCopyAvailable

`func (o *Document) GetSignedCopyAvailable() bool`

GetSignedCopyAvailable returns the SignedCopyAvailable field if non-nil, zero value otherwise.

### GetSignedCopyAvailableOk

`func (o *Document) GetSignedCopyAvailableOk() (*bool, bool)`

GetSignedCopyAvailableOk returns a tuple with the SignedCopyAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedCopyAvailable

`func (o *Document) SetSignedCopyAvailable(v bool)`

SetSignedCopyAvailable sets SignedCopyAvailable field to given value.


### GetParties

`func (o *Document) GetParties() []Party`

GetParties returns the Parties field if non-nil, zero value otherwise.

### GetPartiesOk

`func (o *Document) GetPartiesOk() (*[]Party, bool)`

GetPartiesOk returns a tuple with the Parties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParties

`func (o *Document) SetParties(v []Party)`

SetParties sets Parties field to given value.


### GetPayments

`func (o *Document) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Document) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Document) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.


### GetEvents

`func (o *Document) GetEvents() []DocumentEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Document) GetEventsOk() (*[]DocumentEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Document) SetEvents(v []DocumentEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


