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

// checks if the StartDraftRequestPeopleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartDraftRequestPeopleInner{}

// StartDraftRequestPeopleInner struct for StartDraftRequestPeopleInner
type StartDraftRequestPeopleInner struct {
	Name string `json:"name"`
	Email *string `json:"email,omitempty"`
	Role *string `json:"role,omitempty"`
	Action *string `json:"action,omitempty"`
}

type _StartDraftRequestPeopleInner StartDraftRequestPeopleInner

// NewStartDraftRequestPeopleInner instantiates a new StartDraftRequestPeopleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartDraftRequestPeopleInner(name string) *StartDraftRequestPeopleInner {
	this := StartDraftRequestPeopleInner{}
	this.Name = name
	return &this
}

// NewStartDraftRequestPeopleInnerWithDefaults instantiates a new StartDraftRequestPeopleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartDraftRequestPeopleInnerWithDefaults() *StartDraftRequestPeopleInner {
	this := StartDraftRequestPeopleInner{}
	return &this
}

// GetName returns the Name field value
func (o *StartDraftRequestPeopleInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StartDraftRequestPeopleInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StartDraftRequestPeopleInner) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *StartDraftRequestPeopleInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequestPeopleInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *StartDraftRequestPeopleInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *StartDraftRequestPeopleInner) SetEmail(v string) {
	o.Email = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *StartDraftRequestPeopleInner) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequestPeopleInner) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *StartDraftRequestPeopleInner) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *StartDraftRequestPeopleInner) SetRole(v string) {
	o.Role = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *StartDraftRequestPeopleInner) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequestPeopleInner) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *StartDraftRequestPeopleInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *StartDraftRequestPeopleInner) SetAction(v string) {
	o.Action = &v
}

func (o StartDraftRequestPeopleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartDraftRequestPeopleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

func (o *StartDraftRequestPeopleInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varStartDraftRequestPeopleInner := _StartDraftRequestPeopleInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStartDraftRequestPeopleInner)

	if err != nil {
		return err
	}

	*o = StartDraftRequestPeopleInner(varStartDraftRequestPeopleInner)

	return err
}

type NullableStartDraftRequestPeopleInner struct {
	value *StartDraftRequestPeopleInner
	isSet bool
}

func (v NullableStartDraftRequestPeopleInner) Get() *StartDraftRequestPeopleInner {
	return v.value
}

func (v *NullableStartDraftRequestPeopleInner) Set(val *StartDraftRequestPeopleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStartDraftRequestPeopleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStartDraftRequestPeopleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartDraftRequestPeopleInner(val *StartDraftRequestPeopleInner) *NullableStartDraftRequestPeopleInner {
	return &NullableStartDraftRequestPeopleInner{value: val, isSet: true}
}

func (v NullableStartDraftRequestPeopleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartDraftRequestPeopleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


