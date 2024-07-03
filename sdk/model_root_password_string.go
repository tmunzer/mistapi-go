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
	"fmt"
)

// checks if the RootPasswordString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootPasswordString{}

// RootPasswordString struct for RootPasswordString
type RootPasswordString struct {
	RootPassword string `json:"root_password"`
	AdditionalProperties map[string]interface{}
}

type _RootPasswordString RootPasswordString

// NewRootPasswordString instantiates a new RootPasswordString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootPasswordString(rootPassword string) *RootPasswordString {
	this := RootPasswordString{}
	this.RootPassword = rootPassword
	return &this
}

// NewRootPasswordStringWithDefaults instantiates a new RootPasswordString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootPasswordStringWithDefaults() *RootPasswordString {
	this := RootPasswordString{}
	return &this
}

// GetRootPassword returns the RootPassword field value
func (o *RootPasswordString) GetRootPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value
// and a boolean to check if the value has been set.
func (o *RootPasswordString) GetRootPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootPassword, true
}

// SetRootPassword sets field value
func (o *RootPasswordString) SetRootPassword(v string) {
	o.RootPassword = v
}

func (o RootPasswordString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootPasswordString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["root_password"] = o.RootPassword

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RootPasswordString) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"root_password",
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

	varRootPasswordString := _RootPasswordString{}

	err = json.Unmarshal(data, &varRootPasswordString)

	if err != nil {
		return err
	}

	*o = RootPasswordString(varRootPasswordString)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "root_password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRootPasswordString struct {
	value *RootPasswordString
	isSet bool
}

func (v NullableRootPasswordString) Get() *RootPasswordString {
	return v.value
}

func (v *NullableRootPasswordString) Set(val *RootPasswordString) {
	v.value = val
	v.isSet = true
}

func (v NullableRootPasswordString) IsSet() bool {
	return v.isSet
}

func (v *NullableRootPasswordString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootPasswordString(val *RootPasswordString) *NullableRootPasswordString {
	return &NullableRootPasswordString{value: val, isSet: true}
}

func (v NullableRootPasswordString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootPasswordString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


