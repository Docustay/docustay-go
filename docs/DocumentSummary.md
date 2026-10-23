# DocumentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **NullableString** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | 
**Metadata** | **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | 
**Test** | **bool** | True for a document made in test mode (a test API key or the workspace&#39;s test switch). | 
**Id** | **string** |  | 
**Title** | **string** |  | 
**State** | **string** |  | 
**Parties** | **int32** |  | 
**CreatedAt** | **NullableString** |  | 
**SentAt** | **NullableString** |  | 
**ExecutedAt** | **NullableString** |  | 
**ExpiresAt** | **NullableString** |  | 

## Methods

### NewDocumentSummary

`func NewDocumentSummary(externalId NullableString, metadata map[string]string, test bool, id string, title string, state string, parties int32, createdAt NullableString, sentAt NullableString, executedAt NullableString, expiresAt NullableString, ) *DocumentSummary`

NewDocumentSummary instantiates a new DocumentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSummaryWithDefaults

`func NewDocumentSummaryWithDefaults() *DocumentSummary`

NewDocumentSummaryWithDefaults instantiates a new DocumentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *DocumentSummary) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DocumentSummary) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DocumentSummary) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *DocumentSummary) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *DocumentSummary) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMetadata

`func (o *DocumentSummary) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentSummary) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentSummary) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetTest

`func (o *DocumentSummary) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *DocumentSummary) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *DocumentSummary) SetTest(v bool)`

SetTest sets Test field to given value.


### GetId

`func (o *DocumentSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentSummary) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *DocumentSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentSummary) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetState

`func (o *DocumentSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DocumentSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DocumentSummary) SetState(v string)`

SetState sets State field to given value.


### GetParties

`func (o *DocumentSummary) GetParties() int32`

GetParties returns the Parties field if non-nil, zero value otherwise.

### GetPartiesOk

`func (o *DocumentSummary) GetPartiesOk() (*int32, bool)`

GetPartiesOk returns a tuple with the Parties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParties

`func (o *DocumentSummary) SetParties(v int32)`

SetParties sets Parties field to given value.


### GetCreatedAt

`func (o *DocumentSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DocumentSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DocumentSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetSentAt

`func (o *DocumentSummary) GetSentAt() string`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *DocumentSummary) GetSentAtOk() (*string, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *DocumentSummary) SetSentAt(v string)`

SetSentAt sets SentAt field to given value.


### SetSentAtNil

`func (o *DocumentSummary) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *DocumentSummary) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetExecutedAt

`func (o *DocumentSummary) GetExecutedAt() string`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *DocumentSummary) GetExecutedAtOk() (*string, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *DocumentSummary) SetExecutedAt(v string)`

SetExecutedAt sets ExecutedAt field to given value.


### SetExecutedAtNil

`func (o *DocumentSummary) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *DocumentSummary) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil
### GetExpiresAt

`func (o *DocumentSummary) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DocumentSummary) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DocumentSummary) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *DocumentSummary) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *DocumentSummary) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


