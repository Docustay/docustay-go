# PreflightRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | What you are about to do: &#x60;send_document&#x60;, &#x60;bulk_send&#x60; (with &#x60;documents&#x60;), &#x60;invite_person&#x60; or &#x60;remove_person&#x60; (with &#x60;people&#x60;), &#x60;ai_draft&#x60;, &#x60;ai_edit&#x60;, &#x60;ai_read_file&#x60;, &#x60;create_template&#x60;, &#x60;upload&#x60; (with &#x60;bytes&#x60;), &#x60;use_feature&#x60; (with &#x60;item&#x60;) or &#x60;connect_payments&#x60;. | 
**Documents** | Pointer to **int32** |  | [optional] 
**People** | Pointer to **int32** |  | [optional] 
**Recipients** | Pointer to **[]string** | For &#x60;send_document&#x60;: the email addresses you are about to send to, so a Free workspace is told how many of them are new signers this month (anyone already sent to this month, and your own team, do not count). | [optional] 
**Bytes** | Pointer to **int32** |  | [optional] 
**Item** | Pointer to **string** | For &#x60;use_feature&#x60;: the item to ask about, one of &#x60;unlimited_documents&#x60;, &#x60;hide_footer&#x60;, &#x60;auto_reminders&#x60;, &#x60;bulk_send&#x60;, &#x60;team_roles&#x60;, &#x60;signing_groups&#x60;, &#x60;scribe_actions&#x60;, &#x60;ai_bulk&#x60;, &#x60;custom_plays&#x60; (your own plays and playbooks), &#x60;plays_api&#x60;. Any other name is a 400 &#x60;unknown_item&#x60;. | [optional] 

## Methods

### NewPreflightRequest

`func NewPreflightRequest(action string, ) *PreflightRequest`

NewPreflightRequest instantiates a new PreflightRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreflightRequestWithDefaults

`func NewPreflightRequestWithDefaults() *PreflightRequest`

NewPreflightRequestWithDefaults instantiates a new PreflightRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PreflightRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PreflightRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PreflightRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetDocuments

`func (o *PreflightRequest) GetDocuments() int32`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *PreflightRequest) GetDocumentsOk() (*int32, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *PreflightRequest) SetDocuments(v int32)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *PreflightRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetPeople

`func (o *PreflightRequest) GetPeople() int32`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PreflightRequest) GetPeopleOk() (*int32, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PreflightRequest) SetPeople(v int32)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PreflightRequest) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetRecipients

`func (o *PreflightRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PreflightRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PreflightRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *PreflightRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetBytes

`func (o *PreflightRequest) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *PreflightRequest) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *PreflightRequest) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *PreflightRequest) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetItem

`func (o *PreflightRequest) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PreflightRequest) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PreflightRequest) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *PreflightRequest) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


