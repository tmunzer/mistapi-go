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

// checks if the SiteSettingConfigPushPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingConfigPushPolicy{}

// SiteSettingConfigPushPolicy mist also uses some heuristic rules to prevent destructive configs from being pushed
type SiteSettingConfigPushPolicy struct {
	// stop any new config from being pushed to the device
	NoPush *bool `json:"no_push,omitempty"`
	PushWindow *SiteSettingConfigPushPolicyPushWindow `json:"push_window,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingConfigPushPolicy SiteSettingConfigPushPolicy

// NewSiteSettingConfigPushPolicy instantiates a new SiteSettingConfigPushPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingConfigPushPolicy() *SiteSettingConfigPushPolicy {
	this := SiteSettingConfigPushPolicy{}
	var noPush bool = false
	this.NoPush = &noPush
	return &this
}

// NewSiteSettingConfigPushPolicyWithDefaults instantiates a new SiteSettingConfigPushPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingConfigPushPolicyWithDefaults() *SiteSettingConfigPushPolicy {
	this := SiteSettingConfigPushPolicy{}
	var noPush bool = false
	this.NoPush = &noPush
	return &this
}

// GetNoPush returns the NoPush field value if set, zero value otherwise.
func (o *SiteSettingConfigPushPolicy) GetNoPush() bool {
	if o == nil || IsNil(o.NoPush) {
		var ret bool
		return ret
	}
	return *o.NoPush
}

// GetNoPushOk returns a tuple with the NoPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingConfigPushPolicy) GetNoPushOk() (*bool, bool) {
	if o == nil || IsNil(o.NoPush) {
		return nil, false
	}
	return o.NoPush, true
}

// HasNoPush returns a boolean if a field has been set.
func (o *SiteSettingConfigPushPolicy) HasNoPush() bool {
	if o != nil && !IsNil(o.NoPush) {
		return true
	}

	return false
}

// SetNoPush gets a reference to the given bool and assigns it to the NoPush field.
func (o *SiteSettingConfigPushPolicy) SetNoPush(v bool) {
	o.NoPush = &v
}

// GetPushWindow returns the PushWindow field value if set, zero value otherwise.
func (o *SiteSettingConfigPushPolicy) GetPushWindow() SiteSettingConfigPushPolicyPushWindow {
	if o == nil || IsNil(o.PushWindow) {
		var ret SiteSettingConfigPushPolicyPushWindow
		return ret
	}
	return *o.PushWindow
}

// GetPushWindowOk returns a tuple with the PushWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingConfigPushPolicy) GetPushWindowOk() (*SiteSettingConfigPushPolicyPushWindow, bool) {
	if o == nil || IsNil(o.PushWindow) {
		return nil, false
	}
	return o.PushWindow, true
}

// HasPushWindow returns a boolean if a field has been set.
func (o *SiteSettingConfigPushPolicy) HasPushWindow() bool {
	if o != nil && !IsNil(o.PushWindow) {
		return true
	}

	return false
}

// SetPushWindow gets a reference to the given SiteSettingConfigPushPolicyPushWindow and assigns it to the PushWindow field.
func (o *SiteSettingConfigPushPolicy) SetPushWindow(v SiteSettingConfigPushPolicyPushWindow) {
	o.PushWindow = &v
}

func (o SiteSettingConfigPushPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingConfigPushPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoPush) {
		toSerialize["no_push"] = o.NoPush
	}
	if !IsNil(o.PushWindow) {
		toSerialize["push_window"] = o.PushWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingConfigPushPolicy) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingConfigPushPolicy := _SiteSettingConfigPushPolicy{}

	err = json.Unmarshal(data, &varSiteSettingConfigPushPolicy)

	if err != nil {
		return err
	}

	*o = SiteSettingConfigPushPolicy(varSiteSettingConfigPushPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "no_push")
		delete(additionalProperties, "push_window")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingConfigPushPolicy struct {
	value *SiteSettingConfigPushPolicy
	isSet bool
}

func (v NullableSiteSettingConfigPushPolicy) Get() *SiteSettingConfigPushPolicy {
	return v.value
}

func (v *NullableSiteSettingConfigPushPolicy) Set(val *SiteSettingConfigPushPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingConfigPushPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingConfigPushPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingConfigPushPolicy(val *SiteSettingConfigPushPolicy) *NullableSiteSettingConfigPushPolicy {
	return &NullableSiteSettingConfigPushPolicy{value: val, isSet: true}
}

func (v NullableSiteSettingConfigPushPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingConfigPushPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

