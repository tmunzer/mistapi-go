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
	"os"
)

// checks if the AssetsImportFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsImportFile{}

// AssetsImportFile struct for AssetsImportFile
type AssetsImportFile struct {
	// CSV file
	File **os.File `json:"file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetsImportFile AssetsImportFile

// NewAssetsImportFile instantiates a new AssetsImportFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsImportFile() *AssetsImportFile {
	this := AssetsImportFile{}
	return &this
}

// NewAssetsImportFileWithDefaults instantiates a new AssetsImportFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsImportFileWithDefaults() *AssetsImportFile {
	this := AssetsImportFile{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *AssetsImportFile) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsImportFile) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *AssetsImportFile) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *AssetsImportFile) SetFile(v *os.File) {
	o.File = &v
}

func (o AssetsImportFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsImportFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetsImportFile) UnmarshalJSON(data []byte) (err error) {
	varAssetsImportFile := _AssetsImportFile{}

	err = json.Unmarshal(data, &varAssetsImportFile)

	if err != nil {
		return err
	}

	*o = AssetsImportFile(varAssetsImportFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetsImportFile struct {
	value *AssetsImportFile
	isSet bool
}

func (v NullableAssetsImportFile) Get() *AssetsImportFile {
	return v.value
}

func (v *NullableAssetsImportFile) Set(val *AssetsImportFile) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsImportFile) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsImportFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsImportFile(val *AssetsImportFile) *NullableAssetsImportFile {
	return &NullableAssetsImportFile{value: val, isSet: true}
}

func (v NullableAssetsImportFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsImportFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


