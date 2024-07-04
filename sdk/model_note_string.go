/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the NoteString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteString{}

// NoteString struct for NoteString
type NoteString struct {
	// Some text note describing the intent
	Note *string `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NoteString NoteString

// NewNoteString instantiates a new NoteString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteString() *NoteString {
	this := NoteString{}
	return &this
}

// NewNoteStringWithDefaults instantiates a new NoteString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteStringWithDefaults() *NoteString {
	this := NoteString{}
	return &this
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *NoteString) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteString) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *NoteString) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *NoteString) SetNote(v string) {
	o.Note = &v
}

func (o NoteString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NoteString) UnmarshalJSON(data []byte) (err error) {
	varNoteString := _NoteString{}

	err = json.Unmarshal(data, &varNoteString)

	if err != nil {
		return err
	}

	*o = NoteString(varNoteString)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNoteString struct {
	value *NoteString
	isSet bool
}

func (v NullableNoteString) Get() *NoteString {
	return v.value
}

func (v *NullableNoteString) Set(val *NoteString) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteString) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteString(val *NoteString) *NullableNoteString {
	return &NullableNoteString{value: val, isSet: true}
}

func (v NullableNoteString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


