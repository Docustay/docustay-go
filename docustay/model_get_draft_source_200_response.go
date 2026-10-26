/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"fmt"
)

// checks if the GetDraftSource200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDraftSource200Response{}

// GetDraftSource200Response struct for GetDraftSource200Response
type GetDraftSource200Response struct {
	Id string `json:"id"`
	State string `json:"state"`
	Title string `json:"title"`
	// A written document: TipTap JSON `{type:\"doc\",content:[…]}`. Field tags (`fieldTag` with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (`variableChip`; `custom.<key>` chips are the template's questions), and headings with `audience: \"sender\"` (Sender only: in nothing a signer sees, signs or receives).
	Source map[string]*interface{} `json:"source"`
	Facts map[string]*interface{} `json:"facts"`
	VariableValues map[string]string `json:"variableValues"`
	Paper string `json:"paper"`
	Roles []interface{} `json:"roles"`
	People []interface{} `json:"people"`
	SendOptions map[string]*interface{} `json:"sendOptions"`
	AdditionalProperties map[string]interface{}
}

type _GetDraftSource200Response GetDraftSource200Response

// NewGetDraftSource200Response instantiates a new GetDraftSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDraftSource200Response(id string, state string, title string, source map[string]*interface{}, facts map[string]*interface{}, variableValues map[string]string, paper string, roles []interface{}, people []interface{}, sendOptions map[string]*interface{}) *GetDraftSource200Response {
	this := GetDraftSource200Response{}
	this.Id = id
	this.State = state
	this.Title = title
	this.Source = source
	this.Facts = facts
	this.VariableValues = variableValues
	this.Paper = paper
	this.Roles = roles
	this.People = people
	this.SendOptions = sendOptions
	return &this
}

// NewGetDraftSource200ResponseWithDefaults instantiates a new GetDraftSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDraftSource200ResponseWithDefaults() *GetDraftSource200Response {
	this := GetDraftSource200Response{}
	return &this
}

// GetId returns the Id field value
func (o *GetDraftSource200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetDraftSource200Response) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *GetDraftSource200Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *GetDraftSource200Response) SetState(v string) {
	o.State = v
}

// GetTitle returns the Title field value
func (o *GetDraftSource200Response) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GetDraftSource200Response) SetTitle(v string) {
	o.Title = v
}

// GetSource returns the Source field value
func (o *GetDraftSource200Response) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// SetSource sets field value
func (o *GetDraftSource200Response) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetFacts returns the Facts field value
func (o *GetDraftSource200Response) GetFacts() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Facts
}

// GetFactsOk returns a tuple with the Facts field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetFactsOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Facts, true
}

// SetFacts sets field value
func (o *GetDraftSource200Response) SetFacts(v map[string]*interface{}) {
	o.Facts = v
}

// GetVariableValues returns the VariableValues field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *GetDraftSource200Response) GetVariableValues() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.VariableValues
}

// GetVariableValuesOk returns a tuple with the VariableValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetDraftSource200Response) GetVariableValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.VariableValues) {
		return map[string]string{}, false
	}
	return o.VariableValues, true
}

// SetVariableValues sets field value
func (o *GetDraftSource200Response) SetVariableValues(v map[string]string) {
	o.VariableValues = v
}

// GetPaper returns the Paper field value
func (o *GetDraftSource200Response) GetPaper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Paper
}

// GetPaperOk returns a tuple with the Paper field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetPaperOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paper, true
}

// SetPaper sets field value
func (o *GetDraftSource200Response) SetPaper(v string) {
	o.Paper = v
}

// GetRoles returns the Roles field value
func (o *GetDraftSource200Response) GetRoles() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetRolesOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *GetDraftSource200Response) SetRoles(v []interface{}) {
	o.Roles = v
}

// GetPeople returns the People field value
func (o *GetDraftSource200Response) GetPeople() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.People
}

// GetPeopleOk returns a tuple with the People field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetPeopleOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.People, true
}

// SetPeople sets field value
func (o *GetDraftSource200Response) SetPeople(v []interface{}) {
	o.People = v
}

// GetSendOptions returns the SendOptions field value
func (o *GetDraftSource200Response) GetSendOptions() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.SendOptions
}

// GetSendOptionsOk returns a tuple with the SendOptions field value
// and a boolean to check if the value has been set.
func (o *GetDraftSource200Response) GetSendOptionsOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.SendOptions, true
}

// SetSendOptions sets field value
func (o *GetDraftSource200Response) SetSendOptions(v map[string]*interface{}) {
	o.SendOptions = v
}

func (o GetDraftSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDraftSource200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["title"] = o.Title
	toSerialize["source"] = o.Source
	toSerialize["facts"] = o.Facts
	if o.VariableValues != nil {
		toSerialize["variableValues"] = o.VariableValues
	}
	toSerialize["paper"] = o.Paper
	toSerialize["roles"] = o.Roles
	toSerialize["people"] = o.People
	toSerialize["sendOptions"] = o.SendOptions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDraftSource200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"title",
		"source",
		"facts",
		"variableValues",
		"paper",
		"roles",
		"people",
		"sendOptions",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetDraftSource200Response := _GetDraftSource200Response{}

	err = json.Unmarshal(data, &varGetDraftSource200Response)

	if err != nil {
		return err
	}

	*o = GetDraftSource200Response(varGetDraftSource200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "title")
		delete(additionalProperties, "source")
		delete(additionalProperties, "facts")
		delete(additionalProperties, "variableValues")
		delete(additionalProperties, "paper")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "people")
		delete(additionalProperties, "sendOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDraftSource200Response struct {
	value *GetDraftSource200Response
	isSet bool
}

func (v NullableGetDraftSource200Response) Get() *GetDraftSource200Response {
	return v.value
}

func (v *NullableGetDraftSource200Response) Set(val *GetDraftSource200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDraftSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDraftSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDraftSource200Response(val *GetDraftSource200Response) *NullableGetDraftSource200Response {
	return &NullableGetDraftSource200Response{value: val, isSet: true}
}

func (v NullableGetDraftSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDraftSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


