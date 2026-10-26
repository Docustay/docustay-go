# CreateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailChoice**](EmailChoice.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Title** | **string** |  | 
**Source** | **map[string]interface{}** | A written document: TipTap JSON &#x60;{type:\&quot;doc\&quot;,content:[…]}&#x60;. Field tags (&#x60;fieldTag&#x60; with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (&#x60;variableChip&#x60;; &#x60;custom.&lt;key&gt;&#x60; chips are the template&#39;s questions), and headings with &#x60;audience: \&quot;sender\&quot;&#x60; (Sender only: in nothing a signer sees, signs or receives). | 
**People** | [**[]Person**](Person.md) |  | 
**Facts** | Pointer to **map[string]string** | Merge text: &#x60;org.name&#x60;, &#x60;sender.name&#x60;, &#x60;custom.&lt;key&gt;&#x60; … | [optional] 
**VariableValues** | Pointer to **map[string]string** |  | [optional] 
**Paper** | Pointer to **string** |  | [optional] 
**Send** | Pointer to [**SendChoices**](SendChoices.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SendNow** | Pointer to **bool** | Also send it. The same checks as the screen&#39;s Send apply (a payment field needs Stripe connected). | [optional] 

## Methods

### NewCreateDocumentRequest

`func NewCreateDocumentRequest(title string, source map[string]*interface{}, people []Person, ) *CreateDocumentRequest`

NewCreateDocumentRequest instantiates a new CreateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentRequestWithDefaults

`func NewCreateDocumentRequestWithDefaults() *CreateDocumentRequest`

NewCreateDocumentRequestWithDefaults instantiates a new CreateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateDocumentRequest) GetEmail() EmailChoice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateDocumentRequest) GetEmailOk() (*EmailChoice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateDocumentRequest) SetEmail(v EmailChoice)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateDocumentRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDocumentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDocumentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDocumentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDocumentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateDocumentRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateDocumentRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateDocumentRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateDocumentRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateDocumentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateDocumentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateDocumentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSource

`func (o *CreateDocumentRequest) GetSource() map[string]*interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateDocumentRequest) GetSourceOk() (*map[string]*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateDocumentRequest) SetSource(v map[string]*interface{})`

SetSource sets Source field to given value.


### GetPeople

`func (o *CreateDocumentRequest) GetPeople() []Person`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *CreateDocumentRequest) GetPeopleOk() (*[]Person, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *CreateDocumentRequest) SetPeople(v []Person)`

SetPeople sets People field to given value.


### GetFacts

`func (o *CreateDocumentRequest) GetFacts() map[string]string`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CreateDocumentRequest) GetFactsOk() (*map[string]string, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CreateDocumentRequest) SetFacts(v map[string]string)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *CreateDocumentRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetVariableValues

`func (o *CreateDocumentRequest) GetVariableValues() map[string]string`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *CreateDocumentRequest) GetVariableValuesOk() (*map[string]string, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *CreateDocumentRequest) SetVariableValues(v map[string]string)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *CreateDocumentRequest) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetPaper

`func (o *CreateDocumentRequest) GetPaper() string`

GetPaper returns the Paper field if non-nil, zero value otherwise.

### GetPaperOk

`func (o *CreateDocumentRequest) GetPaperOk() (*string, bool)`

GetPaperOk returns a tuple with the Paper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaper

`func (o *CreateDocumentRequest) SetPaper(v string)`

SetPaper sets Paper field to given value.

### HasPaper

`func (o *CreateDocumentRequest) HasPaper() bool`

HasPaper returns a boolean if a field has been set.

### GetSend

`func (o *CreateDocumentRequest) GetSend() SendChoices`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *CreateDocumentRequest) GetSendOk() (*SendChoices, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *CreateDocumentRequest) SetSend(v SendChoices)`

SetSend sets Send field to given value.

### HasSend

`func (o *CreateDocumentRequest) HasSend() bool`

HasSend returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDocumentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDocumentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDocumentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDocumentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSendNow

`func (o *CreateDocumentRequest) GetSendNow() bool`

GetSendNow returns the SendNow field if non-nil, zero value otherwise.

### GetSendNowOk

`func (o *CreateDocumentRequest) GetSendNowOk() (*bool, bool)`

GetSendNowOk returns a tuple with the SendNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNow

`func (o *CreateDocumentRequest) SetSendNow(v bool)`

SetSendNow sets SendNow field to given value.

### HasSendNow

`func (o *CreateDocumentRequest) HasSendNow() bool`

HasSendNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


