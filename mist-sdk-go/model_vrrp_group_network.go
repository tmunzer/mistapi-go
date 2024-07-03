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

// checks if the VrrpGroupNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrrpGroupNetwork{}

// VrrpGroupNetwork struct for VrrpGroupNetwork
type VrrpGroupNetwork struct {
	Ip *string `json:"ip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrrpGroupNetwork VrrpGroupNetwork

// NewVrrpGroupNetwork instantiates a new VrrpGroupNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrrpGroupNetwork() *VrrpGroupNetwork {
	this := VrrpGroupNetwork{}
	return &this
}

// NewVrrpGroupNetworkWithDefaults instantiates a new VrrpGroupNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrrpGroupNetworkWithDefaults() *VrrpGroupNetwork {
	this := VrrpGroupNetwork{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *VrrpGroupNetwork) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrrpGroupNetwork) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *VrrpGroupNetwork) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *VrrpGroupNetwork) SetIp(v string) {
	o.Ip = &v
}

func (o VrrpGroupNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrrpGroupNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrrpGroupNetwork) UnmarshalJSON(data []byte) (err error) {
	varVrrpGroupNetwork := _VrrpGroupNetwork{}

	err = json.Unmarshal(data, &varVrrpGroupNetwork)

	if err != nil {
		return err
	}

	*o = VrrpGroupNetwork(varVrrpGroupNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrrpGroupNetwork struct {
	value *VrrpGroupNetwork
	isSet bool
}

func (v NullableVrrpGroupNetwork) Get() *VrrpGroupNetwork {
	return v.value
}

func (v *NullableVrrpGroupNetwork) Set(val *VrrpGroupNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVrrpGroupNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVrrpGroupNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrrpGroupNetwork(val *VrrpGroupNetwork) *NullableVrrpGroupNetwork {
	return &NullableVrrpGroupNetwork{value: val, isSet: true}
}

func (v NullableVrrpGroupNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrrpGroupNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


