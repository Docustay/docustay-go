# ReviseDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **map[string]interface{}** | A written document: TipTap JSON &#x60;{type:\&quot;doc\&quot;,content:[…]}&#x60;. Field tags (&#x60;fieldTag&#x60; with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (&#x60;variableChip&#x60;; &#x60;custom.&lt;key&gt;&#x60; chips are the template&#39;s questions), and headings with &#x60;audience: \&quot;sender\&quot;&#x60; (Sender only: in nothing a signer sees, signs or receives). | 
**Facts** | Pointer to **map[string]string** |  | [optional] 
**VariableValues** | Pointer to **map[string]string** |  | [optional] 
**RoleMap** | Pointer to **map[string]string** | Old role word → new role word when you renamed a role. | [optional] 

## Methods

### NewReviseDocumentRequest

`func NewReviseDocumentRequest(source map[string]interface{}, ) *ReviseDocumentRequest`

NewReviseDocumentRequest instantiates a new ReviseDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviseDocumentRequestWithDefaults

`func NewReviseDocumentRequestWithDefaults() *ReviseDocumentRequest`

NewReviseDocumentRequestWithDefaults instantiates a new ReviseDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ReviseDocumentRequest) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReviseDocumentRequest) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReviseDocumentRequest) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.


### GetFacts

`func (o *ReviseDocumentRequest) GetFacts() map[string]string`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *ReviseDocumentRequest) GetFactsOk() (*map[string]string, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *ReviseDocumentRequest) SetFacts(v map[string]string)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *ReviseDocumentRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetVariableValues

`func (o *ReviseDocumentRequest) GetVariableValues() map[string]string`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *ReviseDocumentRequest) GetVariableValuesOk() (*map[string]string, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *ReviseDocumentRequest) SetVariableValues(v map[string]string)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *ReviseDocumentRequest) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetRoleMap

`func (o *ReviseDocumentRequest) GetRoleMap() map[string]string`

GetRoleMap returns the RoleMap field if non-nil, zero value otherwise.

### GetRoleMapOk

`func (o *ReviseDocumentRequest) GetRoleMapOk() (*map[string]string, bool)`

GetRoleMapOk returns a tuple with the RoleMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMap

`func (o *ReviseDocumentRequest) SetRoleMap(v map[string]string)`

SetRoleMap sets RoleMap field to given value.

### HasRoleMap

`func (o *ReviseDocumentRequest) HasRoleMap() bool`

HasRoleMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


