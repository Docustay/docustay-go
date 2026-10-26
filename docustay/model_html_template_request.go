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

// checks if the HtmlTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HtmlTemplateRequest{}

// HtmlTemplateRequest struct for HtmlTemplateRequest
type HtmlTemplateRequest struct {
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	Name string `json:"name"`
	// HTML with text tags such as {{Tenant signature;type=signature;role=Tenant}}. Scripts, styles, frames, forms and remote addresses are removed; pictures must be data: addresses. Up to 2 MB.
	Html string `json:"html"`
	Paper *string `json:"paper,omitempty"`
}

type _HtmlTemplateRequest HtmlTemplateRequest

// NewHtmlTemplateRequest instantiates a new HtmlTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHtmlTemplateRequest(name string, html string) *HtmlTemplateRequest {
	this := HtmlTemplateRequest{}
	this.Name = name
	this.Html = html
	return &this
}

// NewHtmlTemplateRequestWithDefaults instantiates a new HtmlTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHtmlTemplateRequestWithDefaults() *HtmlTemplateRequest {
	this := HtmlTemplateRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HtmlTemplateRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HtmlTemplateRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HtmlTemplateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *HtmlTemplateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *HtmlTemplateRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HtmlTemplateRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *HtmlTemplateRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *HtmlTemplateRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetName returns the Name field value
func (o *HtmlTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HtmlTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HtmlTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetHtml returns the Html field value
func (o *HtmlTemplateRequest) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
func (o *HtmlTemplateRequest) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *HtmlTemplateRequest) SetHtml(v string) {
	o.Html = v
}

// GetPaper returns the Paper field value if set, zero value otherwise.
func (o *HtmlTemplateRequest) GetPaper() string {
	if o == nil || IsNil(o.Paper) {
		var ret string
		return ret
	}
	return *o.Paper
}

// GetPaperOk returns a tuple with the Paper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HtmlTemplateRequest) GetPaperOk() (*string, bool) {
	if o == nil || IsNil(o.Paper) {
		return nil, false
	}
	return o.Paper, true
}

// HasPaper returns a boolean if a field has been set.
func (o *HtmlTemplateRequest) HasPaper() bool {
	if o != nil && !IsNil(o.Paper) {
		return true
	}

	return false
}

// SetPaper gets a reference to the given string and assigns it to the Paper field.
func (o *HtmlTemplateRequest) SetPaper(v string) {
	o.Paper = &v
}

func (o HtmlTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HtmlTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["name"] = o.Name
	toSerialize["html"] = o.Html
	if !IsNil(o.Paper) {
		toSerialize["paper"] = o.Paper
	}
	return toSerialize, nil
}

func (o *HtmlTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"html",
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

	varHtmlTemplateRequest := _HtmlTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHtmlTemplateRequest)

	if err != nil {
		return err
	}

	*o = HtmlTemplateRequest(varHtmlTemplateRequest)

	return err
}

type NullableHtmlTemplateRequest struct {
	value *HtmlTemplateRequest
	isSet bool
}

func (v NullableHtmlTemplateRequest) Get() *HtmlTemplateRequest {
	return v.value
}

func (v *NullableHtmlTemplateRequest) Set(val *HtmlTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHtmlTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHtmlTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHtmlTemplateRequest(val *HtmlTemplateRequest) *NullableHtmlTemplateRequest {
	return &NullableHtmlTemplateRequest{value: val, isSet: true}
}

func (v NullableHtmlTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHtmlTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


