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

// checks if the WlanAppLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanAppLimit{}

// WlanAppLimit bandwidth limiting for apps (applies to up/down)
type WlanAppLimit struct {
	// Map from app key to bandwidth in kbps.  Property key is the app key, defined in Get Application List
	Apps *map[string]int32 `json:"apps,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// Map from wxtag_id of Hostname Wxlan Tags to bandwidth in kbps Property key is the wxtag id
	WxtagIds *map[string]int32 `json:"wxtag_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanAppLimit WlanAppLimit

// NewWlanAppLimit instantiates a new WlanAppLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanAppLimit() *WlanAppLimit {
	this := WlanAppLimit{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewWlanAppLimitWithDefaults instantiates a new WlanAppLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanAppLimitWithDefaults() *WlanAppLimit {
	this := WlanAppLimit{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *WlanAppLimit) GetApps() map[string]int32 {
	if o == nil || IsNil(o.Apps) {
		var ret map[string]int32
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppLimit) GetAppsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *WlanAppLimit) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given map[string]int32 and assigns it to the Apps field.
func (o *WlanAppLimit) SetApps(v map[string]int32) {
	o.Apps = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanAppLimit) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppLimit) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanAppLimit) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanAppLimit) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetWxtagIds returns the WxtagIds field value if set, zero value otherwise.
func (o *WlanAppLimit) GetWxtagIds() map[string]int32 {
	if o == nil || IsNil(o.WxtagIds) {
		var ret map[string]int32
		return ret
	}
	return *o.WxtagIds
}

// GetWxtagIdsOk returns a tuple with the WxtagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppLimit) GetWxtagIdsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.WxtagIds) {
		return nil, false
	}
	return o.WxtagIds, true
}

// HasWxtagIds returns a boolean if a field has been set.
func (o *WlanAppLimit) HasWxtagIds() bool {
	if o != nil && !IsNil(o.WxtagIds) {
		return true
	}

	return false
}

// SetWxtagIds gets a reference to the given map[string]int32 and assigns it to the WxtagIds field.
func (o *WlanAppLimit) SetWxtagIds(v map[string]int32) {
	o.WxtagIds = &v
}

func (o WlanAppLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanAppLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.WxtagIds) {
		toSerialize["wxtag_ids"] = o.WxtagIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanAppLimit) UnmarshalJSON(data []byte) (err error) {
	varWlanAppLimit := _WlanAppLimit{}

	err = json.Unmarshal(data, &varWlanAppLimit)

	if err != nil {
		return err
	}

	*o = WlanAppLimit(varWlanAppLimit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "wxtag_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanAppLimit struct {
	value *WlanAppLimit
	isSet bool
}

func (v NullableWlanAppLimit) Get() *WlanAppLimit {
	return v.value
}

func (v *NullableWlanAppLimit) Set(val *WlanAppLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanAppLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanAppLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanAppLimit(val *WlanAppLimit) *NullableWlanAppLimit {
	return &NullableWlanAppLimit{value: val, isSet: true}
}

func (v NullableWlanAppLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanAppLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


