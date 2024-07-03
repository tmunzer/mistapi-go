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

// checks if the SiteSettingCriticalUrlMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingCriticalUrlMonitoring{}

// SiteSettingCriticalUrlMonitoring you can define some URLs that's critical to site operaitons the latency will be captured and considered for site health
type SiteSettingCriticalUrlMonitoring struct {
	Enabled *bool `json:"enabled,omitempty"`
	Monitors []SiteSettingCriticalUrlMonitoringMonitor `json:"monitors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingCriticalUrlMonitoring SiteSettingCriticalUrlMonitoring

// NewSiteSettingCriticalUrlMonitoring instantiates a new SiteSettingCriticalUrlMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingCriticalUrlMonitoring() *SiteSettingCriticalUrlMonitoring {
	this := SiteSettingCriticalUrlMonitoring{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewSiteSettingCriticalUrlMonitoringWithDefaults instantiates a new SiteSettingCriticalUrlMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingCriticalUrlMonitoringWithDefaults() *SiteSettingCriticalUrlMonitoring {
	this := SiteSettingCriticalUrlMonitoring{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SiteSettingCriticalUrlMonitoring) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingCriticalUrlMonitoring) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SiteSettingCriticalUrlMonitoring) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SiteSettingCriticalUrlMonitoring) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *SiteSettingCriticalUrlMonitoring) GetMonitors() []SiteSettingCriticalUrlMonitoringMonitor {
	if o == nil || IsNil(o.Monitors) {
		var ret []SiteSettingCriticalUrlMonitoringMonitor
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingCriticalUrlMonitoring) GetMonitorsOk() ([]SiteSettingCriticalUrlMonitoringMonitor, bool) {
	if o == nil || IsNil(o.Monitors) {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *SiteSettingCriticalUrlMonitoring) HasMonitors() bool {
	if o != nil && !IsNil(o.Monitors) {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []SiteSettingCriticalUrlMonitoringMonitor and assigns it to the Monitors field.
func (o *SiteSettingCriticalUrlMonitoring) SetMonitors(v []SiteSettingCriticalUrlMonitoringMonitor) {
	o.Monitors = v
}

func (o SiteSettingCriticalUrlMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingCriticalUrlMonitoring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Monitors) {
		toSerialize["monitors"] = o.Monitors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingCriticalUrlMonitoring) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingCriticalUrlMonitoring := _SiteSettingCriticalUrlMonitoring{}

	err = json.Unmarshal(data, &varSiteSettingCriticalUrlMonitoring)

	if err != nil {
		return err
	}

	*o = SiteSettingCriticalUrlMonitoring(varSiteSettingCriticalUrlMonitoring)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "monitors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingCriticalUrlMonitoring struct {
	value *SiteSettingCriticalUrlMonitoring
	isSet bool
}

func (v NullableSiteSettingCriticalUrlMonitoring) Get() *SiteSettingCriticalUrlMonitoring {
	return v.value
}

func (v *NullableSiteSettingCriticalUrlMonitoring) Set(val *SiteSettingCriticalUrlMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingCriticalUrlMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingCriticalUrlMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingCriticalUrlMonitoring(val *SiteSettingCriticalUrlMonitoring) *NullableSiteSettingCriticalUrlMonitoring {
	return &NullableSiteSettingCriticalUrlMonitoring{value: val, isSet: true}
}

func (v NullableSiteSettingCriticalUrlMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingCriticalUrlMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


