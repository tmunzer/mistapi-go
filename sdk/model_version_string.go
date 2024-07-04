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

// checks if the VersionString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionString{}

// VersionString struct for VersionString
type VersionString struct {
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VersionString VersionString

// NewVersionString instantiates a new VersionString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionString() *VersionString {
	this := VersionString{}
	return &this
}

// NewVersionStringWithDefaults instantiates a new VersionString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionStringWithDefaults() *VersionString {
	this := VersionString{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VersionString) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionString) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionString) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VersionString) SetVersion(v string) {
	o.Version = &v
}

func (o VersionString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VersionString) UnmarshalJSON(data []byte) (err error) {
	varVersionString := _VersionString{}

	err = json.Unmarshal(data, &varVersionString)

	if err != nil {
		return err
	}

	*o = VersionString(varVersionString)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVersionString struct {
	value *VersionString
	isSet bool
}

func (v NullableVersionString) Get() *VersionString {
	return v.value
}

func (v *NullableVersionString) Set(val *VersionString) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionString) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionString(val *VersionString) *NullableVersionString {
	return &NullableVersionString{value: val, isSet: true}
}

func (v NullableVersionString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


