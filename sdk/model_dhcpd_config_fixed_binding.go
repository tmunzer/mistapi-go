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

// checks if the DhcpdConfigFixedBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpdConfigFixedBinding{}

// DhcpdConfigFixedBinding struct for DhcpdConfigFixedBinding
type DhcpdConfigFixedBinding struct {
	Ip *string `json:"ip,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DhcpdConfigFixedBinding DhcpdConfigFixedBinding

// NewDhcpdConfigFixedBinding instantiates a new DhcpdConfigFixedBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpdConfigFixedBinding() *DhcpdConfigFixedBinding {
	this := DhcpdConfigFixedBinding{}
	return &this
}

// NewDhcpdConfigFixedBindingWithDefaults instantiates a new DhcpdConfigFixedBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpdConfigFixedBindingWithDefaults() *DhcpdConfigFixedBinding {
	this := DhcpdConfigFixedBinding{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *DhcpdConfigFixedBinding) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfigFixedBinding) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *DhcpdConfigFixedBinding) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *DhcpdConfigFixedBinding) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DhcpdConfigFixedBinding) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfigFixedBinding) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DhcpdConfigFixedBinding) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DhcpdConfigFixedBinding) SetName(v string) {
	o.Name = &v
}

func (o DhcpdConfigFixedBinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpdConfigFixedBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DhcpdConfigFixedBinding) UnmarshalJSON(data []byte) (err error) {
	varDhcpdConfigFixedBinding := _DhcpdConfigFixedBinding{}

	err = json.Unmarshal(data, &varDhcpdConfigFixedBinding)

	if err != nil {
		return err
	}

	*o = DhcpdConfigFixedBinding(varDhcpdConfigFixedBinding)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDhcpdConfigFixedBinding struct {
	value *DhcpdConfigFixedBinding
	isSet bool
}

func (v NullableDhcpdConfigFixedBinding) Get() *DhcpdConfigFixedBinding {
	return v.value
}

func (v *NullableDhcpdConfigFixedBinding) Set(val *DhcpdConfigFixedBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpdConfigFixedBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpdConfigFixedBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpdConfigFixedBinding(val *DhcpdConfigFixedBinding) *NullableDhcpdConfigFixedBinding {
	return &NullableDhcpdConfigFixedBinding{value: val, isSet: true}
}

func (v NullableDhcpdConfigFixedBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpdConfigFixedBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

