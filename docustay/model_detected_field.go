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

// checks if the DetectedField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectedField{}

// DetectedField struct for DetectedField
type DetectedField struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Page int32 `json:"page"`
	X float32 `json:"x"`
	Y float32 `json:"y"`
	W float32 `json:"w"`
	H float32 `json:"h"`
	Confidence float32 `json:"confidence"`
	// `native` (already a form field), `line`, `text`, `box`, `table`, `checkbox` or `highlight`.
	Source string `json:"source"`
	Label *string `json:"label,omitempty"`
	// The seat this field belongs to: one of `signers[].key`.
	Role *string `json:"role,omitempty"`
	Sensitive *bool `json:"sensitive,omitempty"`
	Format *string `json:"format,omitempty"`
	// Tick boxes that answer one question share a group: pick one.
	Group *string `json:"group,omitempty"`
	Uncertain bool `json:"uncertain"`
	AdditionalProperties map[string]interface{}
}

type _DetectedField DetectedField

// NewDetectedField instantiates a new DetectedField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectedField(id string, type_ string, page int32, x float32, y float32, w float32, h float32, confidence float32, source string, uncertain bool) *DetectedField {
	this := DetectedField{}
	this.Id = id
	this.Type = type_
	this.Page = page
	this.X = x
	this.Y = y
	this.W = w
	this.H = h
	this.Confidence = confidence
	this.Source = source
	this.Uncertain = uncertain
	return &this
}

// NewDetectedFieldWithDefaults instantiates a new DetectedField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectedFieldWithDefaults() *DetectedField {
	this := DetectedField{}
	return &this
}

// GetId returns the Id field value
func (o *DetectedField) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DetectedField) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DetectedField) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DetectedField) SetType(v string) {
	o.Type = v
}

// GetPage returns the Page field value
func (o *DetectedField) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DetectedField) SetPage(v int32) {
	o.Page = v
}

// GetX returns the X field value
func (o *DetectedField) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *DetectedField) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *DetectedField) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *DetectedField) SetY(v float32) {
	o.Y = v
}

// GetW returns the W field value
func (o *DetectedField) GetW() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.W
}

// GetWOk returns a tuple with the W field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetWOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.W, true
}

// SetW sets field value
func (o *DetectedField) SetW(v float32) {
	o.W = v
}

// GetH returns the H field value
func (o *DetectedField) GetH() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.H
}

// GetHOk returns a tuple with the H field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetHOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.H, true
}

// SetH sets field value
func (o *DetectedField) SetH(v float32) {
	o.H = v
}

// GetConfidence returns the Confidence field value
func (o *DetectedField) GetConfidence() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetConfidenceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *DetectedField) SetConfidence(v float32) {
	o.Confidence = v
}

// GetSource returns the Source field value
func (o *DetectedField) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DetectedField) SetSource(v string) {
	o.Source = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DetectedField) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectedField) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DetectedField) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DetectedField) SetLabel(v string) {
	o.Label = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *DetectedField) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectedField) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *DetectedField) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *DetectedField) SetRole(v string) {
	o.Role = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *DetectedField) GetSensitive() bool {
	if o == nil || IsNil(o.Sensitive) {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectedField) GetSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Sensitive) {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *DetectedField) HasSensitive() bool {
	if o != nil && !IsNil(o.Sensitive) {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *DetectedField) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DetectedField) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectedField) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DetectedField) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DetectedField) SetFormat(v string) {
	o.Format = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *DetectedField) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectedField) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *DetectedField) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *DetectedField) SetGroup(v string) {
	o.Group = &v
}

// GetUncertain returns the Uncertain field value
func (o *DetectedField) GetUncertain() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Uncertain
}

// GetUncertainOk returns a tuple with the Uncertain field value
// and a boolean to check if the value has been set.
func (o *DetectedField) GetUncertainOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uncertain, true
}

// SetUncertain sets field value
func (o *DetectedField) SetUncertain(v bool) {
	o.Uncertain = v
}

func (o DetectedField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectedField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["page"] = o.Page
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["w"] = o.W
	toSerialize["h"] = o.H
	toSerialize["confidence"] = o.Confidence
	toSerialize["source"] = o.Source
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Sensitive) {
		toSerialize["sensitive"] = o.Sensitive
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	toSerialize["uncertain"] = o.Uncertain

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetectedField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"page",
		"x",
		"y",
		"w",
		"h",
		"confidence",
		"source",
		"uncertain",
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

	varDetectedField := _DetectedField{}

	err = json.Unmarshal(data, &varDetectedField)

	if err != nil {
		return err
	}

	*o = DetectedField(varDetectedField)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "page")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "w")
		delete(additionalProperties, "h")
		delete(additionalProperties, "confidence")
		delete(additionalProperties, "source")
		delete(additionalProperties, "label")
		delete(additionalProperties, "role")
		delete(additionalProperties, "sensitive")
		delete(additionalProperties, "format")
		delete(additionalProperties, "group")
		delete(additionalProperties, "uncertain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetectedField struct {
	value *DetectedField
	isSet bool
}

func (v NullableDetectedField) Get() *DetectedField {
	return v.value
}

func (v *NullableDetectedField) Set(val *DetectedField) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectedField) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectedField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectedField(val *DetectedField) *NullableDetectedField {
	return &NullableDetectedField{value: val, isSet: true}
}

func (v NullableDetectedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectedField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


