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

// checks if the IdpProfileMatching type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpProfileMatching{}

// IdpProfileMatching struct for IdpProfileMatching
type IdpProfileMatching struct {
	AttackName []string `json:"attack_name,omitempty"`
	DstSubnet []string `json:"dst_subnet,omitempty"`
	Severity []IdpProfileMatchingSeverityValue `json:"severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpProfileMatching IdpProfileMatching

// NewIdpProfileMatching instantiates a new IdpProfileMatching object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpProfileMatching() *IdpProfileMatching {
	this := IdpProfileMatching{}
	return &this
}

// NewIdpProfileMatchingWithDefaults instantiates a new IdpProfileMatching object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpProfileMatchingWithDefaults() *IdpProfileMatching {
	this := IdpProfileMatching{}
	return &this
}

// GetAttackName returns the AttackName field value if set, zero value otherwise.
func (o *IdpProfileMatching) GetAttackName() []string {
	if o == nil || IsNil(o.AttackName) {
		var ret []string
		return ret
	}
	return o.AttackName
}

// GetAttackNameOk returns a tuple with the AttackName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpProfileMatching) GetAttackNameOk() ([]string, bool) {
	if o == nil || IsNil(o.AttackName) {
		return nil, false
	}
	return o.AttackName, true
}

// HasAttackName returns a boolean if a field has been set.
func (o *IdpProfileMatching) HasAttackName() bool {
	if o != nil && !IsNil(o.AttackName) {
		return true
	}

	return false
}

// SetAttackName gets a reference to the given []string and assigns it to the AttackName field.
func (o *IdpProfileMatching) SetAttackName(v []string) {
	o.AttackName = v
}

// GetDstSubnet returns the DstSubnet field value if set, zero value otherwise.
func (o *IdpProfileMatching) GetDstSubnet() []string {
	if o == nil || IsNil(o.DstSubnet) {
		var ret []string
		return ret
	}
	return o.DstSubnet
}

// GetDstSubnetOk returns a tuple with the DstSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpProfileMatching) GetDstSubnetOk() ([]string, bool) {
	if o == nil || IsNil(o.DstSubnet) {
		return nil, false
	}
	return o.DstSubnet, true
}

// HasDstSubnet returns a boolean if a field has been set.
func (o *IdpProfileMatching) HasDstSubnet() bool {
	if o != nil && !IsNil(o.DstSubnet) {
		return true
	}

	return false
}

// SetDstSubnet gets a reference to the given []string and assigns it to the DstSubnet field.
func (o *IdpProfileMatching) SetDstSubnet(v []string) {
	o.DstSubnet = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IdpProfileMatching) GetSeverity() []IdpProfileMatchingSeverityValue {
	if o == nil || IsNil(o.Severity) {
		var ret []IdpProfileMatchingSeverityValue
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpProfileMatching) GetSeverityOk() ([]IdpProfileMatchingSeverityValue, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IdpProfileMatching) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given []IdpProfileMatchingSeverityValue and assigns it to the Severity field.
func (o *IdpProfileMatching) SetSeverity(v []IdpProfileMatchingSeverityValue) {
	o.Severity = v
}

func (o IdpProfileMatching) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpProfileMatching) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttackName) {
		toSerialize["attack_name"] = o.AttackName
	}
	if !IsNil(o.DstSubnet) {
		toSerialize["dst_subnet"] = o.DstSubnet
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdpProfileMatching) UnmarshalJSON(data []byte) (err error) {
	varIdpProfileMatching := _IdpProfileMatching{}

	err = json.Unmarshal(data, &varIdpProfileMatching)

	if err != nil {
		return err
	}

	*o = IdpProfileMatching(varIdpProfileMatching)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attack_name")
		delete(additionalProperties, "dst_subnet")
		delete(additionalProperties, "severity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdpProfileMatching struct {
	value *IdpProfileMatching
	isSet bool
}

func (v NullableIdpProfileMatching) Get() *IdpProfileMatching {
	return v.value
}

func (v *NullableIdpProfileMatching) Set(val *IdpProfileMatching) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpProfileMatching) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpProfileMatching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpProfileMatching(val *IdpProfileMatching) *NullableIdpProfileMatching {
	return &NullableIdpProfileMatching{value: val, isSet: true}
}

func (v NullableIdpProfileMatching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpProfileMatching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

