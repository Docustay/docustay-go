# CheckDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **map[string]interface{}** | A written document: TipTap JSON &#x60;{type:\&quot;doc\&quot;,content:[…]}&#x60;. Field tags (&#x60;fieldTag&#x60; with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (&#x60;variableChip&#x60;; &#x60;custom.&lt;key&gt;&#x60; chips are the template&#39;s questions), and headings with &#x60;audience: \&quot;sender\&quot;&#x60; (Sender only: in nothing a signer sees, signs or receives). | 
**Facts** | Pointer to **map[string]string** |  | [optional] 
**People** | Pointer to [**[]Person**](Person.md) | With people, each person&#39;s own merge text is known, as when the document is made. | [optional] 
**VariableValues** | Pointer to **map[string]string** |  | [optional] 
**RequireSignature** | Pointer to **bool** |  | [optional] 
**ForTemplate** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheckDocumentRequest

`func NewCheckDocumentRequest(source map[string]interface{}, ) *CheckDocumentRequest`

NewCheckDocumentRequest instantiates a new CheckDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDocumentRequestWithDefaults

`func NewCheckDocumentRequestWithDefaults() *CheckDocumentRequest`

NewCheckDocumentRequestWithDefaults instantiates a new CheckDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *CheckDocumentRequest) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CheckDocumentRequest) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CheckDocumentRequest) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.


### GetFacts

`func (o *CheckDocumentRequest) GetFacts() map[string]string`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CheckDocumentRequest) GetFactsOk() (*map[string]string, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CheckDocumentRequest) SetFacts(v map[string]string)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *CheckDocumentRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetPeople

`func (o *CheckDocumentRequest) GetPeople() []Person`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *CheckDocumentRequest) GetPeopleOk() (*[]Person, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *CheckDocumentRequest) SetPeople(v []Person)`

SetPeople sets People field to given value.

### HasPeople

`func (o *CheckDocumentRequest) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetVariableValues

`func (o *CheckDocumentRequest) GetVariableValues() map[string]string`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *CheckDocumentRequest) GetVariableValuesOk() (*map[string]string, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *CheckDocumentRequest) SetVariableValues(v map[string]string)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *CheckDocumentRequest) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetRequireSignature

`func (o *CheckDocumentRequest) GetRequireSignature() bool`

GetRequireSignature returns the RequireSignature field if non-nil, zero value otherwise.

### GetRequireSignatureOk

`func (o *CheckDocumentRequest) GetRequireSignatureOk() (*bool, bool)`

GetRequireSignatureOk returns a tuple with the RequireSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignature

`func (o *CheckDocumentRequest) SetRequireSignature(v bool)`

SetRequireSignature sets RequireSignature field to given value.

### HasRequireSignature

`func (o *CheckDocumentRequest) HasRequireSignature() bool`

HasRequireSignature returns a boolean if a field has been set.

### GetForTemplate

`func (o *CheckDocumentRequest) GetForTemplate() bool`

GetForTemplate returns the ForTemplate field if non-nil, zero value otherwise.

### GetForTemplateOk

`func (o *CheckDocumentRequest) GetForTemplateOk() (*bool, bool)`

GetForTemplateOk returns a tuple with the ForTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForTemplate

`func (o *CheckDocumentRequest) SetForTemplate(v bool)`

SetForTemplate sets ForTemplate field to given value.

### HasForTemplate

`func (o *CheckDocumentRequest) HasForTemplate() bool`

HasForTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


