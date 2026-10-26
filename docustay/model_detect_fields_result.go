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

// checks if the DetectFieldsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectFieldsResult{}

// DetectFieldsResult struct for DetectFieldsResult
type DetectFieldsResult struct {
	Fields []DetectedField `json:"fields"`
	Signers []TemplateSeatsInner `json:"signers"`
	Pages []DetectFieldsResultPagesInner `json:"pages"`
	Documents []DetectFieldsResultDocumentsInner `json:"documents"`
	Summary DetectFieldsResultSummary `json:"summary"`
	Notes []string `json:"notes"`
	AdditionalProperties map[string]interface{}
}

type _DetectFieldsResult DetectFieldsResult

// NewDetectFieldsResult instantiates a new DetectFieldsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectFieldsResult(fields []DetectedField, signers []TemplateSeatsInner, pages []DetectFieldsResultPagesInner, documents []DetectFieldsResultDocumentsInner, summary DetectFieldsResultSummary, notes []string) *DetectFieldsResult {
	this := DetectFieldsResult{}
	this.Fields = fields
	this.Signers = signers
	this.Pages = pages
	this.Documents = documents
	this.Summary = summary
	this.Notes = notes
	return &this
}

// NewDetectFieldsResultWithDefaults instantiates a new DetectFieldsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectFieldsResultWithDefaults() *DetectFieldsResult {
	this := DetectFieldsResult{}
	return &this
}

// GetFields returns the Fields field value
func (o *DetectFieldsResult) GetFields() []DetectedField {
	if o == nil {
		var ret []DetectedField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResult) GetFieldsOk() ([]DetectedField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *DetectFieldsResult) SetFields(v []DetectedField) {
	o.Fields = v
}

// GetSigners returns the Signers field value
func (o *DetectFieldsResult) GetSigners() []TemplateSeatsInner {
	if o == nil {
		var ret []TemplateSeatsInner
		return ret
	}

	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResult) GetSignersOk() ([]TemplateSeatsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signers, true
}

// SetSigners sets field value
func (o *DetectFieldsResult) SetSigners(v []TemplateSeatsInner) {
	o.Signers = v
}

// GetPages returns the Pages field value
func (o *DetectFieldsResult) GetPages() []DetectFieldsResultPagesInner {
	if o == nil {
		var ret []DetectFieldsResultPagesInner
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResult) GetPagesOk() ([]DetectFieldsResultPagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *DetectFieldsResult) SetPages(v []DetectFieldsResultPagesInner) {
	o.Pages = v
}

// GetDocuments returns the Documents field value
func (o *DetectFieldsResult) GetDocuments() []DetectFieldsResultDocumentsInner {
	if o == nil {
		var ret []DetectFieldsResultDocumentsInner
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResult) GetDocumentsOk() ([]DetectFieldsResultDocumentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents, true
}

// SetDocuments sets field value
func (o *DetectFieldsResult) SetDocuments(v []DetectFieldsResultDocumentsInner) {
	o.Documents = v
}

// GetSummary returns the Summary field value
func (o *DetectFieldsResult) GetSummary() DetectFieldsResultSummary {
	if o == nil {
		var ret DetectFieldsResultSummary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResult) GetSummaryOk() (*DetectFieldsResultSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *DetectFieldsResult) SetSummary(v DetectFieldsResultSummary) {
	o.Summary = v
}

// GetNotes returns the Notes field value
func (o *DetectFieldsResult) GetNotes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResult) GetNotesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *DetectFieldsResult) SetNotes(v []string) {
	o.Notes = v
}

func (o DetectFieldsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectFieldsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	toSerialize["signers"] = o.Signers
	toSerialize["pages"] = o.Pages
	toSerialize["documents"] = o.Documents
	toSerialize["summary"] = o.Summary
	toSerialize["notes"] = o.Notes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetectFieldsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
		"signers",
		"pages",
		"documents",
		"summary",
		"notes",
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

	varDetectFieldsResult := _DetectFieldsResult{}

	err = json.Unmarshal(data, &varDetectFieldsResult)

	if err != nil {
		return err
	}

	*o = DetectFieldsResult(varDetectFieldsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		delete(additionalProperties, "signers")
		delete(additionalProperties, "pages")
		delete(additionalProperties, "documents")
		delete(additionalProperties, "summary")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetectFieldsResult struct {
	value *DetectFieldsResult
	isSet bool
}

func (v NullableDetectFieldsResult) Get() *DetectFieldsResult {
	return v.value
}

func (v *NullableDetectFieldsResult) Set(val *DetectFieldsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectFieldsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectFieldsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectFieldsResult(val *DetectFieldsResult) *NullableDetectFieldsResult {
	return &NullableDetectFieldsResult{value: val, isSet: true}
}

func (v NullableDetectFieldsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectFieldsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


