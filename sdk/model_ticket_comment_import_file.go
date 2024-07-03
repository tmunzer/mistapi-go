/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"os"
)

// checks if the TicketCommentImportFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TicketCommentImportFile{}

// TicketCommentImportFile struct for TicketCommentImportFile
type TicketCommentImportFile struct {
	Comment *string `json:"comment,omitempty"`
	File **os.File `json:"file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TicketCommentImportFile TicketCommentImportFile

// NewTicketCommentImportFile instantiates a new TicketCommentImportFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketCommentImportFile() *TicketCommentImportFile {
	this := TicketCommentImportFile{}
	return &this
}

// NewTicketCommentImportFileWithDefaults instantiates a new TicketCommentImportFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketCommentImportFileWithDefaults() *TicketCommentImportFile {
	this := TicketCommentImportFile{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TicketCommentImportFile) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentImportFile) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TicketCommentImportFile) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TicketCommentImportFile) SetComment(v string) {
	o.Comment = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *TicketCommentImportFile) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentImportFile) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *TicketCommentImportFile) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *TicketCommentImportFile) SetFile(v *os.File) {
	o.File = &v
}

func (o TicketCommentImportFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketCommentImportFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TicketCommentImportFile) UnmarshalJSON(data []byte) (err error) {
	varTicketCommentImportFile := _TicketCommentImportFile{}

	err = json.Unmarshal(data, &varTicketCommentImportFile)

	if err != nil {
		return err
	}

	*o = TicketCommentImportFile(varTicketCommentImportFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicketCommentImportFile struct {
	value *TicketCommentImportFile
	isSet bool
}

func (v NullableTicketCommentImportFile) Get() *TicketCommentImportFile {
	return v.value
}

func (v *NullableTicketCommentImportFile) Set(val *TicketCommentImportFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketCommentImportFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketCommentImportFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketCommentImportFile(val *TicketCommentImportFile) *NullableTicketCommentImportFile {
	return &NullableTicketCommentImportFile{value: val, isSet: true}
}

func (v NullableTicketCommentImportFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketCommentImportFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


