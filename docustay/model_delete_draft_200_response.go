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

// checks if the DeleteDraft200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDraft200Response{}

// DeleteDraft200Response struct for DeleteDraft200Response
type DeleteDraft200Response struct {
	Id string `json:"id"`
	Deleted bool `json:"deleted"`
	AdditionalProperties map[string]interface{}
}

type _DeleteDraft200Response DeleteDraft200Response

// NewDeleteDraft200Response instantiates a new DeleteDraft200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDraft200Response(id string, deleted bool) *DeleteDraft200Response {
	this := DeleteDraft200Response{}
	this.Id = id
	this.Deleted = deleted
	return &this
}

// NewDeleteDraft200ResponseWithDefaults instantiates a new DeleteDraft200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDraft200ResponseWithDefaults() *DeleteDraft200Response {
	this := DeleteDraft200Response{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteDraft200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteDraft200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteDraft200Response) SetId(v string) {
	o.Id = v
}

// GetDeleted returns the Deleted field value
func (o *DeleteDraft200Response) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *DeleteDraft200Response) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteDraft200Response) SetDeleted(v bool) {
	o.Deleted = v
}

func (o DeleteDraft200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDraft200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["deleted"] = o.Deleted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteDraft200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"deleted",
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

	varDeleteDraft200Response := _DeleteDraft200Response{}

	err = json.Unmarshal(data, &varDeleteDraft200Response)

	if err != nil {
		return err
	}

	*o = DeleteDraft200Response(varDeleteDraft200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "deleted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteDraft200Response struct {
	value *DeleteDraft200Response
	isSet bool
}

func (v NullableDeleteDraft200Response) Get() *DeleteDraft200Response {
	return v.value
}

func (v *NullableDeleteDraft200Response) Set(val *DeleteDraft200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDraft200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDraft200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDraft200Response(val *DeleteDraft200Response) *NullableDeleteDraft200Response {
	return &NullableDeleteDraft200Response{value: val, isSet: true}
}

func (v NullableDeleteDraft200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDraft200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


