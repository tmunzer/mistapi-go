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

// checks if the DeviceOtherStatsVendorSpecific type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceOtherStatsVendorSpecific{}

// DeviceOtherStatsVendorSpecific when `vendor`==`cradlepoint`
type DeviceOtherStatsVendorSpecific struct {
	Ports *map[string]DeviceOtherStatsVendorSpecificPort `json:"ports,omitempty"`
	TargetVersion *string `json:"target_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceOtherStatsVendorSpecific DeviceOtherStatsVendorSpecific

// NewDeviceOtherStatsVendorSpecific instantiates a new DeviceOtherStatsVendorSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceOtherStatsVendorSpecific() *DeviceOtherStatsVendorSpecific {
	this := DeviceOtherStatsVendorSpecific{}
	return &this
}

// NewDeviceOtherStatsVendorSpecificWithDefaults instantiates a new DeviceOtherStatsVendorSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceOtherStatsVendorSpecificWithDefaults() *DeviceOtherStatsVendorSpecific {
	this := DeviceOtherStatsVendorSpecific{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *DeviceOtherStatsVendorSpecific) GetPorts() map[string]DeviceOtherStatsVendorSpecificPort {
	if o == nil || IsNil(o.Ports) {
		var ret map[string]DeviceOtherStatsVendorSpecificPort
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOtherStatsVendorSpecific) GetPortsOk() (*map[string]DeviceOtherStatsVendorSpecificPort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DeviceOtherStatsVendorSpecific) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given map[string]DeviceOtherStatsVendorSpecificPort and assigns it to the Ports field.
func (o *DeviceOtherStatsVendorSpecific) SetPorts(v map[string]DeviceOtherStatsVendorSpecificPort) {
	o.Ports = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *DeviceOtherStatsVendorSpecific) GetTargetVersion() string {
	if o == nil || IsNil(o.TargetVersion) {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOtherStatsVendorSpecific) GetTargetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetVersion) {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *DeviceOtherStatsVendorSpecific) HasTargetVersion() bool {
	if o != nil && !IsNil(o.TargetVersion) {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *DeviceOtherStatsVendorSpecific) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

func (o DeviceOtherStatsVendorSpecific) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceOtherStatsVendorSpecific) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.TargetVersion) {
		toSerialize["target_version"] = o.TargetVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceOtherStatsVendorSpecific) UnmarshalJSON(data []byte) (err error) {
	varDeviceOtherStatsVendorSpecific := _DeviceOtherStatsVendorSpecific{}

	err = json.Unmarshal(data, &varDeviceOtherStatsVendorSpecific)

	if err != nil {
		return err
	}

	*o = DeviceOtherStatsVendorSpecific(varDeviceOtherStatsVendorSpecific)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		delete(additionalProperties, "target_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceOtherStatsVendorSpecific struct {
	value *DeviceOtherStatsVendorSpecific
	isSet bool
}

func (v NullableDeviceOtherStatsVendorSpecific) Get() *DeviceOtherStatsVendorSpecific {
	return v.value
}

func (v *NullableDeviceOtherStatsVendorSpecific) Set(val *DeviceOtherStatsVendorSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceOtherStatsVendorSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceOtherStatsVendorSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceOtherStatsVendorSpecific(val *DeviceOtherStatsVendorSpecific) *NullableDeviceOtherStatsVendorSpecific {
	return &NullableDeviceOtherStatsVendorSpecific{value: val, isSet: true}
}

func (v NullableDeviceOtherStatsVendorSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceOtherStatsVendorSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


