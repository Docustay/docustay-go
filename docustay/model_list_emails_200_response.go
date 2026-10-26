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

// checks if the ListEmails200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEmails200Response{}

// ListEmails200Response struct for ListEmails200Response
type ListEmails200Response struct {
	Emails []map[string]*interface{} `json:"emails"`
	NextCursor NullableString `json:"nextCursor"`
	AdditionalProperties map[string]interface{}
}

type _ListEmails200Response ListEmails200Response

// NewListEmails200Response instantiates a new ListEmails200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEmails200Response(emails []map[string]*interface{}, nextCursor NullableString) *ListEmails200Response {
	this := ListEmails200Response{}
	this.Emails = emails
	this.NextCursor = nextCursor
	return &this
}

// NewListEmails200ResponseWithDefaults instantiates a new ListEmails200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEmails200ResponseWithDefaults() *ListEmails200Response {
	this := ListEmails200Response{}
	return &this
}

// GetEmails returns the Emails field value
func (o *ListEmails200Response) GetEmails() []map[string]*interface{} {
	if o == nil {
		var ret []map[string]*interface{}
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ListEmails200Response) GetEmailsOk() ([]map[string]*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *ListEmails200Response) SetEmails(v []map[string]*interface{}) {
	o.Emails = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListEmails200Response) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListEmails200Response) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *ListEmails200Response) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

func (o ListEmails200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEmails200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emails"] = o.Emails
	toSerialize["nextCursor"] = o.NextCursor.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListEmails200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emails",
		"nextCursor",
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

	varListEmails200Response := _ListEmails200Response{}

	err = json.Unmarshal(data, &varListEmails200Response)

	if err != nil {
		return err
	}

	*o = ListEmails200Response(varListEmails200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emails")
		delete(additionalProperties, "nextCursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListEmails200Response struct {
	value *ListEmails200Response
	isSet bool
}

func (v NullableListEmails200Response) Get() *ListEmails200Response {
	return v.value
}

func (v *NullableListEmails200Response) Set(val *ListEmails200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListEmails200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListEmails200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEmails200Response(val *ListEmails200Response) *NullableListEmails200Response {
	return &NullableListEmails200Response{value: val, isSet: true}
}

func (v NullableListEmails200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEmails200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


