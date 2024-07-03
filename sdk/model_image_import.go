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
	"fmt"
)

// checks if the ImageImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageImport{}

// ImageImport struct for ImageImport
type ImageImport struct {
	// binary file
	File *os.File `json:"file"`
	// JSON string describing your upload
	Json *string `json:"json,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImageImport ImageImport

// NewImageImport instantiates a new ImageImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImport(file *os.File) *ImageImport {
	this := ImageImport{}
	this.File = file
	return &this
}

// NewImageImportWithDefaults instantiates a new ImageImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportWithDefaults() *ImageImport {
	this := ImageImport{}
	return &this
}

// GetFile returns the File field value
func (o *ImageImport) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ImageImport) GetFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ImageImport) SetFile(v *os.File) {
	o.File = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ImageImport) GetJson() string {
	if o == nil || IsNil(o.Json) {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImport) GetJsonOk() (*string, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ImageImport) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *ImageImport) SetJson(v string) {
	o.Json = &v
}

func (o ImageImport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageImport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
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

	varImageImport := _ImageImport{}

	err = json.Unmarshal(data, &varImageImport)

	if err != nil {
		return err
	}

	*o = ImageImport(varImageImport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "json")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageImport struct {
	value *ImageImport
	isSet bool
}

func (v NullableImageImport) Get() *ImageImport {
	return v.value
}

func (v *NullableImageImport) Set(val *ImageImport) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImport) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImport(val *ImageImport) *NullableImageImport {
	return &NullableImageImport{value: val, isSet: true}
}

func (v NullableImageImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


