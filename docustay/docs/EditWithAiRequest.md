# EditWithAiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | **string** |  | 
**Source** | **map[string]interface{}** | A written document: TipTap JSON &#x60;{type:\&quot;doc\&quot;,content:[…]}&#x60;. Field tags (&#x60;fieldTag&#x60; with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (&#x60;variableChip&#x60;; &#x60;custom.&lt;key&gt;&#x60; chips are the template&#39;s questions), and headings with &#x60;audience: \&quot;sender\&quot;&#x60; (Sender only: in nothing a signer sees, signs or receives). | 
**Mode** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**[]EditWithAiRequestVariablesInner**](EditWithAiRequestVariablesInner.md) |  | [optional] 

## Methods

### NewEditWithAiRequest

`func NewEditWithAiRequest(request string, source map[string]*interface{}, ) *EditWithAiRequest`

NewEditWithAiRequest instantiates a new EditWithAiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditWithAiRequestWithDefaults

`func NewEditWithAiRequestWithDefaults() *EditWithAiRequest`

NewEditWithAiRequestWithDefaults instantiates a new EditWithAiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *EditWithAiRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *EditWithAiRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *EditWithAiRequest) SetRequest(v string)`

SetRequest sets Request field to given value.


### GetSource

`func (o *EditWithAiRequest) GetSource() map[string]*interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EditWithAiRequest) GetSourceOk() (*map[string]*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EditWithAiRequest) SetSource(v map[string]*interface{})`

SetSource sets Source field to given value.


### GetMode

`func (o *EditWithAiRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EditWithAiRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EditWithAiRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EditWithAiRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetVariables

`func (o *EditWithAiRequest) GetVariables() []EditWithAiRequestVariablesInner`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *EditWithAiRequest) GetVariablesOk() (*[]EditWithAiRequestVariablesInner, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *EditWithAiRequest) SetVariables(v []EditWithAiRequestVariablesInner)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *EditWithAiRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


