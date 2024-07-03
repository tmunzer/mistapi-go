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

// checks if the OrgSettingPasswordPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingPasswordPolicy{}

// OrgSettingPasswordPolicy password policy
type OrgSettingPasswordPolicy struct {
	// whether the policy is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// days, required if password policy is enabled
	Freshness *int32 `json:"freshness,omitempty"`
	// required password length
	MinLength *int32 `json:"min_length,omitempty"`
	// whether to require special character
	RequiresSpecialChar *bool `json:"requires_special_char,omitempty"`
	// whether to require two-factor auth
	RequiresTwoFactorAuth *bool `json:"requires_two_factor_auth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingPasswordPolicy OrgSettingPasswordPolicy

// NewOrgSettingPasswordPolicy instantiates a new OrgSettingPasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingPasswordPolicy() *OrgSettingPasswordPolicy {
	this := OrgSettingPasswordPolicy{}
	var enabled bool = false
	this.Enabled = &enabled
	var minLength int32 = 8
	this.MinLength = &minLength
	var requiresSpecialChar bool = false
	this.RequiresSpecialChar = &requiresSpecialChar
	var requiresTwoFactorAuth bool = false
	this.RequiresTwoFactorAuth = &requiresTwoFactorAuth
	return &this
}

// NewOrgSettingPasswordPolicyWithDefaults instantiates a new OrgSettingPasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingPasswordPolicyWithDefaults() *OrgSettingPasswordPolicy {
	this := OrgSettingPasswordPolicy{}
	var enabled bool = false
	this.Enabled = &enabled
	var minLength int32 = 8
	this.MinLength = &minLength
	var requiresSpecialChar bool = false
	this.RequiresSpecialChar = &requiresSpecialChar
	var requiresTwoFactorAuth bool = false
	this.RequiresTwoFactorAuth = &requiresTwoFactorAuth
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrgSettingPasswordPolicy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPasswordPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrgSettingPasswordPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrgSettingPasswordPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFreshness returns the Freshness field value if set, zero value otherwise.
func (o *OrgSettingPasswordPolicy) GetFreshness() int32 {
	if o == nil || IsNil(o.Freshness) {
		var ret int32
		return ret
	}
	return *o.Freshness
}

// GetFreshnessOk returns a tuple with the Freshness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPasswordPolicy) GetFreshnessOk() (*int32, bool) {
	if o == nil || IsNil(o.Freshness) {
		return nil, false
	}
	return o.Freshness, true
}

// HasFreshness returns a boolean if a field has been set.
func (o *OrgSettingPasswordPolicy) HasFreshness() bool {
	if o != nil && !IsNil(o.Freshness) {
		return true
	}

	return false
}

// SetFreshness gets a reference to the given int32 and assigns it to the Freshness field.
func (o *OrgSettingPasswordPolicy) SetFreshness(v int32) {
	o.Freshness = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *OrgSettingPasswordPolicy) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPasswordPolicy) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *OrgSettingPasswordPolicy) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *OrgSettingPasswordPolicy) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetRequiresSpecialChar returns the RequiresSpecialChar field value if set, zero value otherwise.
func (o *OrgSettingPasswordPolicy) GetRequiresSpecialChar() bool {
	if o == nil || IsNil(o.RequiresSpecialChar) {
		var ret bool
		return ret
	}
	return *o.RequiresSpecialChar
}

// GetRequiresSpecialCharOk returns a tuple with the RequiresSpecialChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPasswordPolicy) GetRequiresSpecialCharOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresSpecialChar) {
		return nil, false
	}
	return o.RequiresSpecialChar, true
}

// HasRequiresSpecialChar returns a boolean if a field has been set.
func (o *OrgSettingPasswordPolicy) HasRequiresSpecialChar() bool {
	if o != nil && !IsNil(o.RequiresSpecialChar) {
		return true
	}

	return false
}

// SetRequiresSpecialChar gets a reference to the given bool and assigns it to the RequiresSpecialChar field.
func (o *OrgSettingPasswordPolicy) SetRequiresSpecialChar(v bool) {
	o.RequiresSpecialChar = &v
}

// GetRequiresTwoFactorAuth returns the RequiresTwoFactorAuth field value if set, zero value otherwise.
func (o *OrgSettingPasswordPolicy) GetRequiresTwoFactorAuth() bool {
	if o == nil || IsNil(o.RequiresTwoFactorAuth) {
		var ret bool
		return ret
	}
	return *o.RequiresTwoFactorAuth
}

// GetRequiresTwoFactorAuthOk returns a tuple with the RequiresTwoFactorAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPasswordPolicy) GetRequiresTwoFactorAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresTwoFactorAuth) {
		return nil, false
	}
	return o.RequiresTwoFactorAuth, true
}

// HasRequiresTwoFactorAuth returns a boolean if a field has been set.
func (o *OrgSettingPasswordPolicy) HasRequiresTwoFactorAuth() bool {
	if o != nil && !IsNil(o.RequiresTwoFactorAuth) {
		return true
	}

	return false
}

// SetRequiresTwoFactorAuth gets a reference to the given bool and assigns it to the RequiresTwoFactorAuth field.
func (o *OrgSettingPasswordPolicy) SetRequiresTwoFactorAuth(v bool) {
	o.RequiresTwoFactorAuth = &v
}

func (o OrgSettingPasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingPasswordPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Freshness) {
		toSerialize["freshness"] = o.Freshness
	}
	if !IsNil(o.MinLength) {
		toSerialize["min_length"] = o.MinLength
	}
	if !IsNil(o.RequiresSpecialChar) {
		toSerialize["requires_special_char"] = o.RequiresSpecialChar
	}
	if !IsNil(o.RequiresTwoFactorAuth) {
		toSerialize["requires_two_factor_auth"] = o.RequiresTwoFactorAuth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingPasswordPolicy) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingPasswordPolicy := _OrgSettingPasswordPolicy{}

	err = json.Unmarshal(data, &varOrgSettingPasswordPolicy)

	if err != nil {
		return err
	}

	*o = OrgSettingPasswordPolicy(varOrgSettingPasswordPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "freshness")
		delete(additionalProperties, "min_length")
		delete(additionalProperties, "requires_special_char")
		delete(additionalProperties, "requires_two_factor_auth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingPasswordPolicy struct {
	value *OrgSettingPasswordPolicy
	isSet bool
}

func (v NullableOrgSettingPasswordPolicy) Get() *OrgSettingPasswordPolicy {
	return v.value
}

func (v *NullableOrgSettingPasswordPolicy) Set(val *OrgSettingPasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingPasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingPasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingPasswordPolicy(val *OrgSettingPasswordPolicy) *NullableOrgSettingPasswordPolicy {
	return &NullableOrgSettingPasswordPolicy{value: val, isSet: true}
}

func (v NullableOrgSettingPasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingPasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


