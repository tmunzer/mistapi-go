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
)

// checks if the UtilsClearBpdu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsClearBpdu{}

// UtilsClearBpdu struct for UtilsClearBpdu
type UtilsClearBpdu struct {
	// the port on which to clear the detected BPDU error, or `all` for all ports
	Port *string `json:"port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsClearBpdu UtilsClearBpdu

// NewUtilsClearBpdu instantiates a new UtilsClearBpdu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsClearBpdu() *UtilsClearBpdu {
	this := UtilsClearBpdu{}
	return &this
}

// NewUtilsClearBpduWithDefaults instantiates a new UtilsClearBpdu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsClearBpduWithDefaults() *UtilsClearBpdu {
	this := UtilsClearBpdu{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UtilsClearBpdu) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsClearBpdu) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UtilsClearBpdu) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *UtilsClearBpdu) SetPort(v string) {
	o.Port = &v
}

func (o UtilsClearBpdu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsClearBpdu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsClearBpdu) UnmarshalJSON(data []byte) (err error) {
	varUtilsClearBpdu := _UtilsClearBpdu{}

	err = json.Unmarshal(data, &varUtilsClearBpdu)

	if err != nil {
		return err
	}

	*o = UtilsClearBpdu(varUtilsClearBpdu)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsClearBpdu struct {
	value *UtilsClearBpdu
	isSet bool
}

func (v NullableUtilsClearBpdu) Get() *UtilsClearBpdu {
	return v.value
}

func (v *NullableUtilsClearBpdu) Set(val *UtilsClearBpdu) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsClearBpdu) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsClearBpdu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsClearBpdu(val *UtilsClearBpdu) *NullableUtilsClearBpdu {
	return &NullableUtilsClearBpdu{value: val, isSet: true}
}

func (v NullableUtilsClearBpdu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsClearBpdu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

