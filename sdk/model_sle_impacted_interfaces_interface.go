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

// checks if the SleImpactedInterfacesInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleImpactedInterfacesInterface{}

// SleImpactedInterfacesInterface struct for SleImpactedInterfacesInterface
type SleImpactedInterfacesInterface struct {
	Degraded *float32 `json:"degraded,omitempty"`
	Duration *float32 `json:"duration,omitempty"`
	InterfaceName *string `json:"interface_name,omitempty"`
	SwitchMac *string `json:"switch_mac,omitempty"`
	SwitchName *string `json:"switch_name,omitempty"`
	Total *float32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SleImpactedInterfacesInterface SleImpactedInterfacesInterface

// NewSleImpactedInterfacesInterface instantiates a new SleImpactedInterfacesInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleImpactedInterfacesInterface() *SleImpactedInterfacesInterface {
	this := SleImpactedInterfacesInterface{}
	return &this
}

// NewSleImpactedInterfacesInterfaceWithDefaults instantiates a new SleImpactedInterfacesInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleImpactedInterfacesInterfaceWithDefaults() *SleImpactedInterfacesInterface {
	this := SleImpactedInterfacesInterface{}
	return &this
}

// GetDegraded returns the Degraded field value if set, zero value otherwise.
func (o *SleImpactedInterfacesInterface) GetDegraded() float32 {
	if o == nil || IsNil(o.Degraded) {
		var ret float32
		return ret
	}
	return *o.Degraded
}

// GetDegradedOk returns a tuple with the Degraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedInterfacesInterface) GetDegradedOk() (*float32, bool) {
	if o == nil || IsNil(o.Degraded) {
		return nil, false
	}
	return o.Degraded, true
}

// HasDegraded returns a boolean if a field has been set.
func (o *SleImpactedInterfacesInterface) HasDegraded() bool {
	if o != nil && !IsNil(o.Degraded) {
		return true
	}

	return false
}

// SetDegraded gets a reference to the given float32 and assigns it to the Degraded field.
func (o *SleImpactedInterfacesInterface) SetDegraded(v float32) {
	o.Degraded = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SleImpactedInterfacesInterface) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedInterfacesInterface) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SleImpactedInterfacesInterface) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *SleImpactedInterfacesInterface) SetDuration(v float32) {
	o.Duration = &v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *SleImpactedInterfacesInterface) GetInterfaceName() string {
	if o == nil || IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedInterfacesInterface) GetInterfaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *SleImpactedInterfacesInterface) HasInterfaceName() bool {
	if o != nil && !IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *SleImpactedInterfacesInterface) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetSwitchMac returns the SwitchMac field value if set, zero value otherwise.
func (o *SleImpactedInterfacesInterface) GetSwitchMac() string {
	if o == nil || IsNil(o.SwitchMac) {
		var ret string
		return ret
	}
	return *o.SwitchMac
}

// GetSwitchMacOk returns a tuple with the SwitchMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedInterfacesInterface) GetSwitchMacOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchMac) {
		return nil, false
	}
	return o.SwitchMac, true
}

// HasSwitchMac returns a boolean if a field has been set.
func (o *SleImpactedInterfacesInterface) HasSwitchMac() bool {
	if o != nil && !IsNil(o.SwitchMac) {
		return true
	}

	return false
}

// SetSwitchMac gets a reference to the given string and assigns it to the SwitchMac field.
func (o *SleImpactedInterfacesInterface) SetSwitchMac(v string) {
	o.SwitchMac = &v
}

// GetSwitchName returns the SwitchName field value if set, zero value otherwise.
func (o *SleImpactedInterfacesInterface) GetSwitchName() string {
	if o == nil || IsNil(o.SwitchName) {
		var ret string
		return ret
	}
	return *o.SwitchName
}

// GetSwitchNameOk returns a tuple with the SwitchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedInterfacesInterface) GetSwitchNameOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchName) {
		return nil, false
	}
	return o.SwitchName, true
}

// HasSwitchName returns a boolean if a field has been set.
func (o *SleImpactedInterfacesInterface) HasSwitchName() bool {
	if o != nil && !IsNil(o.SwitchName) {
		return true
	}

	return false
}

// SetSwitchName gets a reference to the given string and assigns it to the SwitchName field.
func (o *SleImpactedInterfacesInterface) SetSwitchName(v string) {
	o.SwitchName = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SleImpactedInterfacesInterface) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedInterfacesInterface) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SleImpactedInterfacesInterface) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *SleImpactedInterfacesInterface) SetTotal(v float32) {
	o.Total = &v
}

func (o SleImpactedInterfacesInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleImpactedInterfacesInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Degraded) {
		toSerialize["degraded"] = o.Degraded
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.InterfaceName) {
		toSerialize["interface_name"] = o.InterfaceName
	}
	if !IsNil(o.SwitchMac) {
		toSerialize["switch_mac"] = o.SwitchMac
	}
	if !IsNil(o.SwitchName) {
		toSerialize["switch_name"] = o.SwitchName
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleImpactedInterfacesInterface) UnmarshalJSON(data []byte) (err error) {
	varSleImpactedInterfacesInterface := _SleImpactedInterfacesInterface{}

	err = json.Unmarshal(data, &varSleImpactedInterfacesInterface)

	if err != nil {
		return err
	}

	*o = SleImpactedInterfacesInterface(varSleImpactedInterfacesInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "degraded")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "interface_name")
		delete(additionalProperties, "switch_mac")
		delete(additionalProperties, "switch_name")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleImpactedInterfacesInterface struct {
	value *SleImpactedInterfacesInterface
	isSet bool
}

func (v NullableSleImpactedInterfacesInterface) Get() *SleImpactedInterfacesInterface {
	return v.value
}

func (v *NullableSleImpactedInterfacesInterface) Set(val *SleImpactedInterfacesInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSleImpactedInterfacesInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSleImpactedInterfacesInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleImpactedInterfacesInterface(val *SleImpactedInterfacesInterface) *NullableSleImpactedInterfacesInterface {
	return &NullableSleImpactedInterfacesInterface{value: val, isSet: true}
}

func (v NullableSleImpactedInterfacesInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleImpactedInterfacesInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


