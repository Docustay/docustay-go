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

// checks if the Template type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Template{}

// Template struct for Template
type Template struct {
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId NullableString `json:"externalId,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	Pages NullableInt32 `json:"pages"`
	Seats []TemplateSeatsInner `json:"seats"`
	AdditionalProperties map[string]interface{}
}

type _Template Template

// NewTemplate instantiates a new Template object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplate(id string, name string, pages NullableInt32, seats []TemplateSeatsInner) *Template {
	this := Template{}
	this.Id = id
	this.Name = name
	this.Pages = pages
	this.Seats = seats
	return &this
}

// NewTemplateWithDefaults instantiates a new Template object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateWithDefaults() *Template {
	this := Template{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Template) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *Template) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *Template) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *Template) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *Template) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Template) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Template) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Template) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetId returns the Id field value
func (o *Template) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Template) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Template) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Template) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Template) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Template) SetName(v string) {
	o.Name = v
}

// GetPages returns the Pages field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Template) GetPages() int32 {
	if o == nil || o.Pages.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Pages.Get()
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Template) GetPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages.Get(), o.Pages.IsSet()
}

// SetPages sets field value
func (o *Template) SetPages(v int32) {
	o.Pages.Set(&v)
}

// GetSeats returns the Seats field value
func (o *Template) GetSeats() []TemplateSeatsInner {
	if o == nil {
		var ret []TemplateSeatsInner
		return ret
	}

	return o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value
// and a boolean to check if the value has been set.
func (o *Template) GetSeatsOk() ([]TemplateSeatsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seats, true
}

// SetSeats sets field value
func (o *Template) SetSeats(v []TemplateSeatsInner) {
	o.Seats = v
}

func (o Template) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Template) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["pages"] = o.Pages.Get()
	toSerialize["seats"] = o.Seats

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Template) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"pages",
		"seats",
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

	varTemplate := _Template{}

	err = json.Unmarshal(data, &varTemplate)

	if err != nil {
		return err
	}

	*o = Template(varTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pages")
		delete(additionalProperties, "seats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplate struct {
	value *Template
	isSet bool
}

func (v NullableTemplate) Get() *Template {
	return v.value
}

func (v *NullableTemplate) Set(val *Template) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplate(val *Template) *NullableTemplate {
	return &NullableTemplate{value: val, isSet: true}
}

func (v NullableTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


