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

// checks if the ListDocuments200ResponsePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDocuments200ResponsePage{}

// ListDocuments200ResponsePage struct for ListDocuments200ResponsePage
type ListDocuments200ResponsePage struct {
	Limit int32 `json:"limit"`
	NextCursor NullableString `json:"nextCursor"`
	HasMore bool `json:"hasMore"`
	AdditionalProperties map[string]interface{}
}

type _ListDocuments200ResponsePage ListDocuments200ResponsePage

// NewListDocuments200ResponsePage instantiates a new ListDocuments200ResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDocuments200ResponsePage(limit int32, nextCursor NullableString, hasMore bool) *ListDocuments200ResponsePage {
	this := ListDocuments200ResponsePage{}
	this.Limit = limit
	this.NextCursor = nextCursor
	this.HasMore = hasMore
	return &this
}

// NewListDocuments200ResponsePageWithDefaults instantiates a new ListDocuments200ResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDocuments200ResponsePageWithDefaults() *ListDocuments200ResponsePage {
	this := ListDocuments200ResponsePage{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListDocuments200ResponsePage) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListDocuments200ResponsePage) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListDocuments200ResponsePage) SetLimit(v int32) {
	o.Limit = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListDocuments200ResponsePage) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDocuments200ResponsePage) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *ListDocuments200ResponsePage) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetHasMore returns the HasMore field value
func (o *ListDocuments200ResponsePage) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ListDocuments200ResponsePage) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ListDocuments200ResponsePage) SetHasMore(v bool) {
	o.HasMore = v
}

func (o ListDocuments200ResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDocuments200ResponsePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["nextCursor"] = o.NextCursor.Get()
	toSerialize["hasMore"] = o.HasMore

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDocuments200ResponsePage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"nextCursor",
		"hasMore",
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

	varListDocuments200ResponsePage := _ListDocuments200ResponsePage{}

	err = json.Unmarshal(data, &varListDocuments200ResponsePage)

	if err != nil {
		return err
	}

	*o = ListDocuments200ResponsePage(varListDocuments200ResponsePage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "nextCursor")
		delete(additionalProperties, "hasMore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDocuments200ResponsePage struct {
	value *ListDocuments200ResponsePage
	isSet bool
}

func (v NullableListDocuments200ResponsePage) Get() *ListDocuments200ResponsePage {
	return v.value
}

func (v *NullableListDocuments200ResponsePage) Set(val *ListDocuments200ResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableListDocuments200ResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableListDocuments200ResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDocuments200ResponsePage(val *ListDocuments200ResponsePage) *NullableListDocuments200ResponsePage {
	return &NullableListDocuments200ResponsePage{value: val, isSet: true}
}

func (v NullableListDocuments200ResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDocuments200ResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


