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

// checks if the BuilderSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuilderSessionRequest{}

// BuilderSessionRequest struct for BuilderSessionRequest
type BuilderSessionRequest struct {
	// The website that will show the builder; must be on your allowed list.
	Origin string `json:"origin"`
	// A name to start the template with.
	Name *string `json:"name,omitempty"`
}

type _BuilderSessionRequest BuilderSessionRequest

// NewBuilderSessionRequest instantiates a new BuilderSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuilderSessionRequest(origin string) *BuilderSessionRequest {
	this := BuilderSessionRequest{}
	this.Origin = origin
	return &this
}

// NewBuilderSessionRequestWithDefaults instantiates a new BuilderSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuilderSessionRequestWithDefaults() *BuilderSessionRequest {
	this := BuilderSessionRequest{}
	return &this
}

// GetOrigin returns the Origin field value
func (o *BuilderSessionRequest) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *BuilderSessionRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *BuilderSessionRequest) SetOrigin(v string) {
	o.Origin = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BuilderSessionRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuilderSessionRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BuilderSessionRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BuilderSessionRequest) SetName(v string) {
	o.Name = &v
}

func (o BuilderSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuilderSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origin"] = o.Origin
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *BuilderSessionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBuilderSessionRequest := _BuilderSessionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuilderSessionRequest)

	if err != nil {
		return err
	}

	*o = BuilderSessionRequest(varBuilderSessionRequest)

	return err
}

type NullableBuilderSessionRequest struct {
	value *BuilderSessionRequest
	isSet bool
}

func (v NullableBuilderSessionRequest) Get() *BuilderSessionRequest {
	return v.value
}

func (v *NullableBuilderSessionRequest) Set(val *BuilderSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuilderSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuilderSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuilderSessionRequest(val *BuilderSessionRequest) *NullableBuilderSessionRequest {
	return &NullableBuilderSessionRequest{value: val, isSet: true}
}

func (v NullableBuilderSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuilderSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


