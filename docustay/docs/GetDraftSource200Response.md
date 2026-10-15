# GetDraftSource200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** |  | 
**Title** | **string** |  | 
**Source** | **map[string]interface{}** | A written document: TipTap JSON &#x60;{type:\&quot;doc\&quot;,content:[…]}&#x60;. Field tags (&#x60;fieldTag&#x60; with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (&#x60;variableChip&#x60;; &#x60;custom.&lt;key&gt;&#x60; chips are the template&#39;s questions), and headings with &#x60;audience: \&quot;sender\&quot;&#x60; (Sender only: in nothing a signer sees, signs or receives). | 
**Facts** | **map[string]interface{}** |  | 
**VariableValues** | **map[string]string** |  | 
**Paper** | **string** |  | 
**Roles** | **[]interface{}** |  | 
**People** | **[]interface{}** |  | 
**SendOptions** | **map[string]interface{}** |  | 

## Methods

### NewGetDraftSource200Response

`func NewGetDraftSource200Response(id string, state string, title string, source map[string]interface{}, facts map[string]interface{}, variableValues map[string]string, paper string, roles []interface{}, people []interface{}, sendOptions map[string]interface{}, ) *GetDraftSource200Response`

NewGetDraftSource200Response instantiates a new GetDraftSource200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDraftSource200ResponseWithDefaults

`func NewGetDraftSource200ResponseWithDefaults() *GetDraftSource200Response`

NewGetDraftSource200ResponseWithDefaults instantiates a new GetDraftSource200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDraftSource200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDraftSource200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDraftSource200Response) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *GetDraftSource200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetDraftSource200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetDraftSource200Response) SetState(v string)`

SetState sets State field to given value.


### GetTitle

`func (o *GetDraftSource200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetDraftSource200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetDraftSource200Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSource

`func (o *GetDraftSource200Response) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetDraftSource200Response) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetDraftSource200Response) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.


### GetFacts

`func (o *GetDraftSource200Response) GetFacts() map[string]interface{}`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *GetDraftSource200Response) GetFactsOk() (*map[string]interface{}, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *GetDraftSource200Response) SetFacts(v map[string]interface{})`

SetFacts sets Facts field to given value.


### GetVariableValues

`func (o *GetDraftSource200Response) GetVariableValues() map[string]string`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *GetDraftSource200Response) GetVariableValuesOk() (*map[string]string, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *GetDraftSource200Response) SetVariableValues(v map[string]string)`

SetVariableValues sets VariableValues field to given value.


### GetPaper

`func (o *GetDraftSource200Response) GetPaper() string`

GetPaper returns the Paper field if non-nil, zero value otherwise.

### GetPaperOk

`func (o *GetDraftSource200Response) GetPaperOk() (*string, bool)`

GetPaperOk returns a tuple with the Paper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaper

`func (o *GetDraftSource200Response) SetPaper(v string)`

SetPaper sets Paper field to given value.


### GetRoles

`func (o *GetDraftSource200Response) GetRoles() []interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetDraftSource200Response) GetRolesOk() (*[]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetDraftSource200Response) SetRoles(v []interface{})`

SetRoles sets Roles field to given value.


### GetPeople

`func (o *GetDraftSource200Response) GetPeople() []interface{}`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *GetDraftSource200Response) GetPeopleOk() (*[]interface{}, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *GetDraftSource200Response) SetPeople(v []interface{})`

SetPeople sets People field to given value.


### GetSendOptions

`func (o *GetDraftSource200Response) GetSendOptions() map[string]interface{}`

GetSendOptions returns the SendOptions field if non-nil, zero value otherwise.

### GetSendOptionsOk

`func (o *GetDraftSource200Response) GetSendOptionsOk() (*map[string]interface{}, bool)`

GetSendOptionsOk returns a tuple with the SendOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOptions

`func (o *GetDraftSource200Response) SetSendOptions(v map[string]interface{})`

SetSendOptions sets SendOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


