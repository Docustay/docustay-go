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

// checks if the ListTeams200ResponseTeamsInnerMembersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTeams200ResponseTeamsInnerMembersInner{}

// ListTeams200ResponseTeamsInnerMembersInner struct for ListTeams200ResponseTeamsInnerMembersInner
type ListTeams200ResponseTeamsInnerMembersInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ListTeams200ResponseTeamsInnerMembersInner ListTeams200ResponseTeamsInnerMembersInner

// NewListTeams200ResponseTeamsInnerMembersInner instantiates a new ListTeams200ResponseTeamsInnerMembersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTeams200ResponseTeamsInnerMembersInner(id string, name string) *ListTeams200ResponseTeamsInnerMembersInner {
	this := ListTeams200ResponseTeamsInnerMembersInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewListTeams200ResponseTeamsInnerMembersInnerWithDefaults instantiates a new ListTeams200ResponseTeamsInnerMembersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTeams200ResponseTeamsInnerMembersInnerWithDefaults() *ListTeams200ResponseTeamsInnerMembersInner {
	this := ListTeams200ResponseTeamsInnerMembersInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListTeams200ResponseTeamsInnerMembersInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInnerMembersInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListTeams200ResponseTeamsInnerMembersInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListTeams200ResponseTeamsInnerMembersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListTeams200ResponseTeamsInnerMembersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListTeams200ResponseTeamsInnerMembersInner) SetName(v string) {
	o.Name = v
}

func (o ListTeams200ResponseTeamsInnerMembersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTeams200ResponseTeamsInnerMembersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListTeams200ResponseTeamsInnerMembersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varListTeams200ResponseTeamsInnerMembersInner := _ListTeams200ResponseTeamsInnerMembersInner{}

	err = json.Unmarshal(data, &varListTeams200ResponseTeamsInnerMembersInner)

	if err != nil {
		return err
	}

	*o = ListTeams200ResponseTeamsInnerMembersInner(varListTeams200ResponseTeamsInnerMembersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListTeams200ResponseTeamsInnerMembersInner struct {
	value *ListTeams200ResponseTeamsInnerMembersInner
	isSet bool
}

func (v NullableListTeams200ResponseTeamsInnerMembersInner) Get() *ListTeams200ResponseTeamsInnerMembersInner {
	return v.value
}

func (v *NullableListTeams200ResponseTeamsInnerMembersInner) Set(val *ListTeams200ResponseTeamsInnerMembersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTeams200ResponseTeamsInnerMembersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTeams200ResponseTeamsInnerMembersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTeams200ResponseTeamsInnerMembersInner(val *ListTeams200ResponseTeamsInnerMembersInner) *NullableListTeams200ResponseTeamsInnerMembersInner {
	return &NullableListTeams200ResponseTeamsInnerMembersInner{value: val, isSet: true}
}

func (v NullableListTeams200ResponseTeamsInnerMembersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTeams200ResponseTeamsInnerMembersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


