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

// checks if the ListTeams200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTeams200Response{}

// ListTeams200Response struct for ListTeams200Response
type ListTeams200Response struct {
	Teams []ListTeams200ResponseTeamsInner `json:"teams"`
	AdditionalProperties map[string]interface{}
}

type _ListTeams200Response ListTeams200Response

// NewListTeams200Response instantiates a new ListTeams200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTeams200Response(teams []ListTeams200ResponseTeamsInner) *ListTeams200Response {
	this := ListTeams200Response{}
	this.Teams = teams
	return &this
}

// NewListTeams200ResponseWithDefaults instantiates a new ListTeams200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTeams200ResponseWithDefaults() *ListTeams200Response {
	this := ListTeams200Response{}
	return &this
}

// GetTeams returns the Teams field value
func (o *ListTeams200Response) GetTeams() []ListTeams200ResponseTeamsInner {
	if o == nil {
		var ret []ListTeams200ResponseTeamsInner
		return ret
	}

	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *ListTeams200Response) GetTeamsOk() ([]ListTeams200ResponseTeamsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teams, true
}

// SetTeams sets field value
func (o *ListTeams200Response) SetTeams(v []ListTeams200ResponseTeamsInner) {
	o.Teams = v
}

func (o ListTeams200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTeams200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["teams"] = o.Teams

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListTeams200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"teams",
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

	varListTeams200Response := _ListTeams200Response{}

	err = json.Unmarshal(data, &varListTeams200Response)

	if err != nil {
		return err
	}

	*o = ListTeams200Response(varListTeams200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "teams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListTeams200Response struct {
	value *ListTeams200Response
	isSet bool
}

func (v NullableListTeams200Response) Get() *ListTeams200Response {
	return v.value
}

func (v *NullableListTeams200Response) Set(val *ListTeams200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTeams200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTeams200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTeams200Response(val *ListTeams200Response) *NullableListTeams200Response {
	return &NullableListTeams200Response{value: val, isSet: true}
}

func (v NullableListTeams200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTeams200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


