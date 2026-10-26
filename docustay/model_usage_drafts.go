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

// checks if the UsageDrafts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageDrafts{}

// UsageDrafts AI drafts used this period against the plan's included amount (hosted plans).
type UsageDrafts struct {
	Used NullableFloat32 `json:"used"`
	Included NullableFloat32 `json:"included"`
}

type _UsageDrafts UsageDrafts

// NewUsageDrafts instantiates a new UsageDrafts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageDrafts(used NullableFloat32, included NullableFloat32) *UsageDrafts {
	this := UsageDrafts{}
	this.Used = used
	this.Included = included
	return &this
}

// NewUsageDraftsWithDefaults instantiates a new UsageDrafts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageDraftsWithDefaults() *UsageDrafts {
	this := UsageDrafts{}
	return &this
}

// GetUsed returns the Used field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *UsageDrafts) GetUsed() float32 {
	if o == nil || o.Used.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Used.Get()
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageDrafts) GetUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Used.Get(), o.Used.IsSet()
}

// SetUsed sets field value
func (o *UsageDrafts) SetUsed(v float32) {
	o.Used.Set(&v)
}

// GetIncluded returns the Included field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *UsageDrafts) GetIncluded() float32 {
	if o == nil || o.Included.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Included.Get()
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageDrafts) GetIncludedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Included.Get(), o.Included.IsSet()
}

// SetIncluded sets field value
func (o *UsageDrafts) SetIncluded(v float32) {
	o.Included.Set(&v)
}

func (o UsageDrafts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageDrafts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["used"] = o.Used.Get()
	toSerialize["included"] = o.Included.Get()
	return toSerialize, nil
}

func (o *UsageDrafts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"used",
		"included",
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

	varUsageDrafts := _UsageDrafts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageDrafts)

	if err != nil {
		return err
	}

	*o = UsageDrafts(varUsageDrafts)

	return err
}

type NullableUsageDrafts struct {
	value *UsageDrafts
	isSet bool
}

func (v NullableUsageDrafts) Get() *UsageDrafts {
	return v.value
}

func (v *NullableUsageDrafts) Set(val *UsageDrafts) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageDrafts) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageDrafts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageDrafts(val *UsageDrafts) *NullableUsageDrafts {
	return &NullableUsageDrafts{value: val, isSet: true}
}

func (v NullableUsageDrafts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageDrafts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


