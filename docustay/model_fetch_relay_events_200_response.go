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

// checks if the FetchRelayEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchRelayEvents200Response{}

// FetchRelayEvents200Response struct for FetchRelayEvents200Response
type FetchRelayEvents200Response struct {
	Events []FetchRelayEvents200ResponseEventsInner `json:"events"`
	Cursor NullableString `json:"cursor"`
	AdditionalProperties map[string]interface{}
}

type _FetchRelayEvents200Response FetchRelayEvents200Response

// NewFetchRelayEvents200Response instantiates a new FetchRelayEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchRelayEvents200Response(events []FetchRelayEvents200ResponseEventsInner, cursor NullableString) *FetchRelayEvents200Response {
	this := FetchRelayEvents200Response{}
	this.Events = events
	this.Cursor = cursor
	return &this
}

// NewFetchRelayEvents200ResponseWithDefaults instantiates a new FetchRelayEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchRelayEvents200ResponseWithDefaults() *FetchRelayEvents200Response {
	this := FetchRelayEvents200Response{}
	return &this
}

// GetEvents returns the Events field value
func (o *FetchRelayEvents200Response) GetEvents() []FetchRelayEvents200ResponseEventsInner {
	if o == nil {
		var ret []FetchRelayEvents200ResponseEventsInner
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *FetchRelayEvents200Response) GetEventsOk() ([]FetchRelayEvents200ResponseEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *FetchRelayEvents200Response) SetEvents(v []FetchRelayEvents200ResponseEventsInner) {
	o.Events = v
}

// GetCursor returns the Cursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FetchRelayEvents200Response) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchRelayEvents200Response) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// SetCursor sets field value
func (o *FetchRelayEvents200Response) SetCursor(v string) {
	o.Cursor.Set(&v)
}

func (o FetchRelayEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchRelayEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	toSerialize["cursor"] = o.Cursor.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchRelayEvents200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
		"cursor",
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

	varFetchRelayEvents200Response := _FetchRelayEvents200Response{}

	err = json.Unmarshal(data, &varFetchRelayEvents200Response)

	if err != nil {
		return err
	}

	*o = FetchRelayEvents200Response(varFetchRelayEvents200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		delete(additionalProperties, "cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchRelayEvents200Response struct {
	value *FetchRelayEvents200Response
	isSet bool
}

func (v NullableFetchRelayEvents200Response) Get() *FetchRelayEvents200Response {
	return v.value
}

func (v *NullableFetchRelayEvents200Response) Set(val *FetchRelayEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchRelayEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchRelayEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchRelayEvents200Response(val *FetchRelayEvents200Response) *NullableFetchRelayEvents200Response {
	return &NullableFetchRelayEvents200Response{value: val, isSet: true}
}

func (v NullableFetchRelayEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchRelayEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


