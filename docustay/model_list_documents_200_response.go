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

// checks if the ListDocuments200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDocuments200Response{}

// ListDocuments200Response struct for ListDocuments200Response
type ListDocuments200Response struct {
	Documents []DocumentSummary `json:"documents"`
	Page ListDocuments200ResponsePage `json:"page"`
	AdditionalProperties map[string]interface{}
}

type _ListDocuments200Response ListDocuments200Response

// NewListDocuments200Response instantiates a new ListDocuments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDocuments200Response(documents []DocumentSummary, page ListDocuments200ResponsePage) *ListDocuments200Response {
	this := ListDocuments200Response{}
	this.Documents = documents
	this.Page = page
	return &this
}

// NewListDocuments200ResponseWithDefaults instantiates a new ListDocuments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDocuments200ResponseWithDefaults() *ListDocuments200Response {
	this := ListDocuments200Response{}
	return &this
}

// GetDocuments returns the Documents field value
func (o *ListDocuments200Response) GetDocuments() []DocumentSummary {
	if o == nil {
		var ret []DocumentSummary
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *ListDocuments200Response) GetDocumentsOk() ([]DocumentSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents, true
}

// SetDocuments sets field value
func (o *ListDocuments200Response) SetDocuments(v []DocumentSummary) {
	o.Documents = v
}

// GetPage returns the Page field value
func (o *ListDocuments200Response) GetPage() ListDocuments200ResponsePage {
	if o == nil {
		var ret ListDocuments200ResponsePage
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ListDocuments200Response) GetPageOk() (*ListDocuments200ResponsePage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ListDocuments200Response) SetPage(v ListDocuments200ResponsePage) {
	o.Page = v
}

func (o ListDocuments200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDocuments200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documents"] = o.Documents
	toSerialize["page"] = o.Page

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDocuments200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"documents",
		"page",
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

	varListDocuments200Response := _ListDocuments200Response{}

	err = json.Unmarshal(data, &varListDocuments200Response)

	if err != nil {
		return err
	}

	*o = ListDocuments200Response(varListDocuments200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "documents")
		delete(additionalProperties, "page")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDocuments200Response struct {
	value *ListDocuments200Response
	isSet bool
}

func (v NullableListDocuments200Response) Get() *ListDocuments200Response {
	return v.value
}

func (v *NullableListDocuments200Response) Set(val *ListDocuments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListDocuments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListDocuments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDocuments200Response(val *ListDocuments200Response) *NullableListDocuments200Response {
	return &NullableListDocuments200Response{value: val, isSet: true}
}

func (v NullableListDocuments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDocuments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


