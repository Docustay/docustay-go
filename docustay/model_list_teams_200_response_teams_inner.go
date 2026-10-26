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

// checks if the ListTeams200ResponseTeamsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTeams200ResponseTeamsInner{}

// ListTeams200ResponseTeamsInner struct for ListTeams200ResponseTeamsInner
type ListTeams200ResponseTeamsInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Members []ListTeams200ResponseTeamsInnerMembersInner `json:"members"`
	Documents int32 `json:"documents"`
	Templates int32 `json:"templates"`
	AdditionalProperties map[string]interface{}
}

type _ListTeams200ResponseTeamsInner ListTeams200ResponseTeamsInner

// NewListTeams200ResponseTeamsInner instantiates a new ListTeams200ResponseTeamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTeams200ResponseTeamsInner(id string, name string, members []ListTeams200ResponseTeamsInnerMembersInner, documents int32, templates int32) *ListTeams200ResponseTeamsInner {
	this := ListTeams200ResponseTeamsInner{}
	this.Id = id
	this.Name = name
	this.Members = members
	this.Documents = documents
	this.Templates = templates
	return &this
}

// NewListTeams200ResponseTeamsInnerWithDefaults instantiates a new ListTeams200ResponseTeamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTeams200ResponseTeamsInnerWithDefaults() *ListTeams200ResponseTeamsInner {
	this := ListTeams200ResponseTeamsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListTeams200ResponseTeamsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListTeams200ResponseTeamsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListTeams200ResponseTeamsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListTeams200ResponseTeamsInner) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value
func (o *ListTeams200ResponseTeamsInner) GetMembers() []ListTeams200ResponseTeamsInnerMembersInner {
	if o == nil {
		var ret []ListTeams200ResponseTeamsInnerMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInner) GetMembersOk() ([]ListTeams200ResponseTeamsInnerMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *ListTeams200ResponseTeamsInner) SetMembers(v []ListTeams200ResponseTeamsInnerMembersInner) {
	o.Members = v
}

// GetDocuments returns the Documents field value
func (o *ListTeams200ResponseTeamsInner) GetDocuments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInner) GetDocumentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Documents, true
}

// SetDocuments sets field value
func (o *ListTeams200ResponseTeamsInner) SetDocuments(v int32) {
	o.Documents = v
}

// GetTemplates returns the Templates field value
func (o *ListTeams200ResponseTeamsInner) GetTemplates() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInner) GetTemplatesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Templates, true
}

// SetTemplates sets field value
func (o *ListTeams200ResponseTeamsInner) SetTemplates(v int32) {
	o.Templates = v
}

func (o ListTeams200ResponseTeamsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTeams200ResponseTeamsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["members"] = o.Members
	toSerialize["documents"] = o.Documents
	toSerialize["templates"] = o.Templates

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListTeams200ResponseTeamsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"members",
		"documents",
		"templates",
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

	varListTeams200ResponseTeamsInner := _ListTeams200ResponseTeamsInner{}

	err = json.Unmarshal(data, &varListTeams200ResponseTeamsInner)

	if err != nil {
		return err
	}

	*o = ListTeams200ResponseTeamsInner(varListTeams200ResponseTeamsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "members")
		delete(additionalProperties, "documents")
		delete(additionalProperties, "templates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListTeams200ResponseTeamsInner struct {
	value *ListTeams200ResponseTeamsInner
	isSet bool
}

func (v NullableListTeams200ResponseTeamsInner) Get() *ListTeams200ResponseTeamsInner {
	return v.value
}

func (v *NullableListTeams200ResponseTeamsInner) Set(val *ListTeams200ResponseTeamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTeams200ResponseTeamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTeams200ResponseTeamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTeams200ResponseTeamsInner(val *ListTeams200ResponseTeamsInner) *NullableListTeams200ResponseTeamsInner {
	return &NullableListTeams200ResponseTeamsInner{value: val, isSet: true}
}

func (v NullableListTeams200ResponseTeamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTeams200ResponseTeamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


