/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ScribeSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScribeSessionRequest{}

// ScribeSessionRequest struct for ScribeSessionRequest
type ScribeSessionRequest struct {
	// The website that will show Scribe; must be on your allowed list.
	Origin string `json:"origin"`
	Name *string `json:"name,omitempty"`
}

type _ScribeSessionRequest ScribeSessionRequest

// NewScribeSessionRequest instantiates a new ScribeSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScribeSessionRequest(origin string) *ScribeSessionRequest {
	this := ScribeSessionRequest{}
	this.Origin = origin
	return &this
}

// NewScribeSessionRequestWithDefaults instantiates a new ScribeSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScribeSessionRequestWithDefaults() *ScribeSessionRequest {
	this := ScribeSessionRequest{}
	return &this
}

// GetOrigin returns the Origin field value
func (o *ScribeSessionRequest) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *ScribeSessionRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *ScribeSessionRequest) SetOrigin(v string) {
	o.Origin = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScribeSessionRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScribeSessionRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScribeSessionRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScribeSessionRequest) SetName(v string) {
	o.Name = &v
}

func (o ScribeSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScribeSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origin"] = o.Origin
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *ScribeSessionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"origin",
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

	varScribeSessionRequest := _ScribeSessionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScribeSessionRequest)

	if err != nil {
		return err
	}

	*o = ScribeSessionRequest(varScribeSessionRequest)

	return err
}

type NullableScribeSessionRequest struct {
	value *ScribeSessionRequest
	isSet bool
}

func (v NullableScribeSessionRequest) Get() *ScribeSessionRequest {
	return v.value
}

func (v *NullableScribeSessionRequest) Set(val *ScribeSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScribeSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScribeSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScribeSessionRequest(val *ScribeSessionRequest) *NullableScribeSessionRequest {
	return &NullableScribeSessionRequest{value: val, isSet: true}
}

func (v NullableScribeSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScribeSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


