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

// checks if the ProtectRe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtectRe{}

// ProtectRe restrict inbound-traffic to host when enabled, all traffic that is not essential to our operation will be dropped  e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
type ProtectRe struct {
	// optionally, services we'll allow
	AllowedServices []string `json:"allowed_services,omitempty"`
	Custom []ProtectReCustom `json:"custom,omitempty"`
	// when enabled, all traffic that is not essential to our operation will be dropped e.g. ntp / dns / traffic to mist will be allowed by default      if dhcpd is enabled, we'll make sure it works
	Enabled *bool `json:"enabled,omitempty"`
	// host/subnets we'll allow traffic to/from
	TrustedHosts []string `json:"trusted_hosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtectRe ProtectRe

// NewProtectRe instantiates a new ProtectRe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectRe() *ProtectRe {
	this := ProtectRe{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewProtectReWithDefaults instantiates a new ProtectRe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectReWithDefaults() *ProtectRe {
	this := ProtectRe{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetAllowedServices returns the AllowedServices field value if set, zero value otherwise.
func (o *ProtectRe) GetAllowedServices() []string {
	if o == nil || IsNil(o.AllowedServices) {
		var ret []string
		return ret
	}
	return o.AllowedServices
}

// GetAllowedServicesOk returns a tuple with the AllowedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectRe) GetAllowedServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedServices) {
		return nil, false
	}
	return o.AllowedServices, true
}

// HasAllowedServices returns a boolean if a field has been set.
func (o *ProtectRe) HasAllowedServices() bool {
	if o != nil && !IsNil(o.AllowedServices) {
		return true
	}

	return false
}

// SetAllowedServices gets a reference to the given []string and assigns it to the AllowedServices field.
func (o *ProtectRe) SetAllowedServices(v []string) {
	o.AllowedServices = v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *ProtectRe) GetCustom() []ProtectReCustom {
	if o == nil || IsNil(o.Custom) {
		var ret []ProtectReCustom
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectRe) GetCustomOk() ([]ProtectReCustom, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *ProtectRe) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given []ProtectReCustom and assigns it to the Custom field.
func (o *ProtectRe) SetCustom(v []ProtectReCustom) {
	o.Custom = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProtectRe) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectRe) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProtectRe) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProtectRe) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTrustedHosts returns the TrustedHosts field value if set, zero value otherwise.
func (o *ProtectRe) GetTrustedHosts() []string {
	if o == nil || IsNil(o.TrustedHosts) {
		var ret []string
		return ret
	}
	return o.TrustedHosts
}

// GetTrustedHostsOk returns a tuple with the TrustedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectRe) GetTrustedHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedHosts) {
		return nil, false
	}
	return o.TrustedHosts, true
}

// HasTrustedHosts returns a boolean if a field has been set.
func (o *ProtectRe) HasTrustedHosts() bool {
	if o != nil && !IsNil(o.TrustedHosts) {
		return true
	}

	return false
}

// SetTrustedHosts gets a reference to the given []string and assigns it to the TrustedHosts field.
func (o *ProtectRe) SetTrustedHosts(v []string) {
	o.TrustedHosts = v
}

func (o ProtectRe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtectRe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedServices) {
		toSerialize["allowed_services"] = o.AllowedServices
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.TrustedHosts) {
		toSerialize["trusted_hosts"] = o.TrustedHosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtectRe) UnmarshalJSON(data []byte) (err error) {
	varProtectRe := _ProtectRe{}

	err = json.Unmarshal(data, &varProtectRe)

	if err != nil {
		return err
	}

	*o = ProtectRe(varProtectRe)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowed_services")
		delete(additionalProperties, "custom")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "trusted_hosts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtectRe struct {
	value *ProtectRe
	isSet bool
}

func (v NullableProtectRe) Get() *ProtectRe {
	return v.value
}

func (v *NullableProtectRe) Set(val *ProtectRe) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectRe) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectRe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectRe(val *ProtectRe) *NullableProtectRe {
	return &NullableProtectRe{value: val, isSet: true}
}

func (v NullableProtectRe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectRe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


