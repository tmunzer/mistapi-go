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

// checks if the OrgSettingGatewayMgmtHostOutPolicies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingGatewayMgmtHostOutPolicies{}

// OrgSettingGatewayMgmtHostOutPolicies optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp
type OrgSettingGatewayMgmtHostOutPolicies struct {
	Dns *OrgSettingGatewayMgmtHostOutPoliciesDns `json:"dns,omitempty"`
	Mist *OrgSettingGatewayMgmtHostOutPoliciesMist `json:"mist,omitempty"`
	Ntp *OrgSettingGatewayMgmtHostOutPoliciesNtp `json:"ntp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingGatewayMgmtHostOutPolicies OrgSettingGatewayMgmtHostOutPolicies

// NewOrgSettingGatewayMgmtHostOutPolicies instantiates a new OrgSettingGatewayMgmtHostOutPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingGatewayMgmtHostOutPolicies() *OrgSettingGatewayMgmtHostOutPolicies {
	this := OrgSettingGatewayMgmtHostOutPolicies{}
	return &this
}

// NewOrgSettingGatewayMgmtHostOutPoliciesWithDefaults instantiates a new OrgSettingGatewayMgmtHostOutPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingGatewayMgmtHostOutPoliciesWithDefaults() *OrgSettingGatewayMgmtHostOutPolicies {
	this := OrgSettingGatewayMgmtHostOutPolicies{}
	return &this
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmtHostOutPolicies) GetDns() OrgSettingGatewayMgmtHostOutPoliciesDns {
	if o == nil || IsNil(o.Dns) {
		var ret OrgSettingGatewayMgmtHostOutPoliciesDns
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmtHostOutPolicies) GetDnsOk() (*OrgSettingGatewayMgmtHostOutPoliciesDns, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmtHostOutPolicies) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given OrgSettingGatewayMgmtHostOutPoliciesDns and assigns it to the Dns field.
func (o *OrgSettingGatewayMgmtHostOutPolicies) SetDns(v OrgSettingGatewayMgmtHostOutPoliciesDns) {
	o.Dns = &v
}

// GetMist returns the Mist field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmtHostOutPolicies) GetMist() OrgSettingGatewayMgmtHostOutPoliciesMist {
	if o == nil || IsNil(o.Mist) {
		var ret OrgSettingGatewayMgmtHostOutPoliciesMist
		return ret
	}
	return *o.Mist
}

// GetMistOk returns a tuple with the Mist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmtHostOutPolicies) GetMistOk() (*OrgSettingGatewayMgmtHostOutPoliciesMist, bool) {
	if o == nil || IsNil(o.Mist) {
		return nil, false
	}
	return o.Mist, true
}

// HasMist returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmtHostOutPolicies) HasMist() bool {
	if o != nil && !IsNil(o.Mist) {
		return true
	}

	return false
}

// SetMist gets a reference to the given OrgSettingGatewayMgmtHostOutPoliciesMist and assigns it to the Mist field.
func (o *OrgSettingGatewayMgmtHostOutPolicies) SetMist(v OrgSettingGatewayMgmtHostOutPoliciesMist) {
	o.Mist = &v
}

// GetNtp returns the Ntp field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmtHostOutPolicies) GetNtp() OrgSettingGatewayMgmtHostOutPoliciesNtp {
	if o == nil || IsNil(o.Ntp) {
		var ret OrgSettingGatewayMgmtHostOutPoliciesNtp
		return ret
	}
	return *o.Ntp
}

// GetNtpOk returns a tuple with the Ntp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmtHostOutPolicies) GetNtpOk() (*OrgSettingGatewayMgmtHostOutPoliciesNtp, bool) {
	if o == nil || IsNil(o.Ntp) {
		return nil, false
	}
	return o.Ntp, true
}

// HasNtp returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmtHostOutPolicies) HasNtp() bool {
	if o != nil && !IsNil(o.Ntp) {
		return true
	}

	return false
}

// SetNtp gets a reference to the given OrgSettingGatewayMgmtHostOutPoliciesNtp and assigns it to the Ntp field.
func (o *OrgSettingGatewayMgmtHostOutPolicies) SetNtp(v OrgSettingGatewayMgmtHostOutPoliciesNtp) {
	o.Ntp = &v
}

func (o OrgSettingGatewayMgmtHostOutPolicies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingGatewayMgmtHostOutPolicies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Mist) {
		toSerialize["mist"] = o.Mist
	}
	if !IsNil(o.Ntp) {
		toSerialize["ntp"] = o.Ntp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingGatewayMgmtHostOutPolicies) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingGatewayMgmtHostOutPolicies := _OrgSettingGatewayMgmtHostOutPolicies{}

	err = json.Unmarshal(data, &varOrgSettingGatewayMgmtHostOutPolicies)

	if err != nil {
		return err
	}

	*o = OrgSettingGatewayMgmtHostOutPolicies(varOrgSettingGatewayMgmtHostOutPolicies)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dns")
		delete(additionalProperties, "mist")
		delete(additionalProperties, "ntp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingGatewayMgmtHostOutPolicies struct {
	value *OrgSettingGatewayMgmtHostOutPolicies
	isSet bool
}

func (v NullableOrgSettingGatewayMgmtHostOutPolicies) Get() *OrgSettingGatewayMgmtHostOutPolicies {
	return v.value
}

func (v *NullableOrgSettingGatewayMgmtHostOutPolicies) Set(val *OrgSettingGatewayMgmtHostOutPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingGatewayMgmtHostOutPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingGatewayMgmtHostOutPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingGatewayMgmtHostOutPolicies(val *OrgSettingGatewayMgmtHostOutPolicies) *NullableOrgSettingGatewayMgmtHostOutPolicies {
	return &NullableOrgSettingGatewayMgmtHostOutPolicies{value: val, isSet: true}
}

func (v NullableOrgSettingGatewayMgmtHostOutPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingGatewayMgmtHostOutPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


