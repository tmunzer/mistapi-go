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

// checks if the ApStatsLldpStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsLldpStat{}

// ApStatsLldpStat LLDP Stat (neighbor information, power negotiations)
type ApStatsLldpStat struct {
	ChassisId *string `json:"chassis_id,omitempty"`
	// whether it support LLDP-MED
	LldpMedSupported *bool `json:"lldp_med_supported,omitempty"`
	// switch’s management address (if advertised), can be IPv4, IPv6, or MAC
	MgmtAddr *string `json:"mgmt_addr,omitempty"`
	// port description, e.g. “2/20”, “Port 2 on Switch0”
	PortDesc *string `json:"port_desc,omitempty"`
	// in mW, provided/allocated by PSE
	PowerAllocated *float32 `json:"power_allocated,omitempty"`
	// in mW, total power needed by PD
	PowerDraw *float32 `json:"power_draw,omitempty"`
	// number of negotiations, if it keeps increasing, we don’t have a stable power
	PowerRequestCount *int32 `json:"power_request_count,omitempty"`
	// in mW, the current power requested by PD
	PowerRequested *float32 `json:"power_requested,omitempty"`
	// description provided by switch, e.g. “HP J9729A 2920-48G-POE+ Switch”
	SystemDesc *string `json:"system_desc,omitempty"`
	// name of the switch, e.g. “TC2-OWL-Stack-01”
	SystemName *string `json:"system_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsLldpStat ApStatsLldpStat

// NewApStatsLldpStat instantiates a new ApStatsLldpStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsLldpStat() *ApStatsLldpStat {
	this := ApStatsLldpStat{}
	return &this
}

// NewApStatsLldpStatWithDefaults instantiates a new ApStatsLldpStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsLldpStatWithDefaults() *ApStatsLldpStat {
	this := ApStatsLldpStat{}
	return &this
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetChassisId() string {
	if o == nil || IsNil(o.ChassisId) {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetChassisIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *ApStatsLldpStat) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetLldpMedSupported returns the LldpMedSupported field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetLldpMedSupported() bool {
	if o == nil || IsNil(o.LldpMedSupported) {
		var ret bool
		return ret
	}
	return *o.LldpMedSupported
}

// GetLldpMedSupportedOk returns a tuple with the LldpMedSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetLldpMedSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.LldpMedSupported) {
		return nil, false
	}
	return o.LldpMedSupported, true
}

// HasLldpMedSupported returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasLldpMedSupported() bool {
	if o != nil && !IsNil(o.LldpMedSupported) {
		return true
	}

	return false
}

// SetLldpMedSupported gets a reference to the given bool and assigns it to the LldpMedSupported field.
func (o *ApStatsLldpStat) SetLldpMedSupported(v bool) {
	o.LldpMedSupported = &v
}

// GetMgmtAddr returns the MgmtAddr field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetMgmtAddr() string {
	if o == nil || IsNil(o.MgmtAddr) {
		var ret string
		return ret
	}
	return *o.MgmtAddr
}

// GetMgmtAddrOk returns a tuple with the MgmtAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetMgmtAddrOk() (*string, bool) {
	if o == nil || IsNil(o.MgmtAddr) {
		return nil, false
	}
	return o.MgmtAddr, true
}

// HasMgmtAddr returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasMgmtAddr() bool {
	if o != nil && !IsNil(o.MgmtAddr) {
		return true
	}

	return false
}

// SetMgmtAddr gets a reference to the given string and assigns it to the MgmtAddr field.
func (o *ApStatsLldpStat) SetMgmtAddr(v string) {
	o.MgmtAddr = &v
}

// GetPortDesc returns the PortDesc field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetPortDesc() string {
	if o == nil || IsNil(o.PortDesc) {
		var ret string
		return ret
	}
	return *o.PortDesc
}

// GetPortDescOk returns a tuple with the PortDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetPortDescOk() (*string, bool) {
	if o == nil || IsNil(o.PortDesc) {
		return nil, false
	}
	return o.PortDesc, true
}

// HasPortDesc returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasPortDesc() bool {
	if o != nil && !IsNil(o.PortDesc) {
		return true
	}

	return false
}

// SetPortDesc gets a reference to the given string and assigns it to the PortDesc field.
func (o *ApStatsLldpStat) SetPortDesc(v string) {
	o.PortDesc = &v
}

// GetPowerAllocated returns the PowerAllocated field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetPowerAllocated() float32 {
	if o == nil || IsNil(o.PowerAllocated) {
		var ret float32
		return ret
	}
	return *o.PowerAllocated
}

// GetPowerAllocatedOk returns a tuple with the PowerAllocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetPowerAllocatedOk() (*float32, bool) {
	if o == nil || IsNil(o.PowerAllocated) {
		return nil, false
	}
	return o.PowerAllocated, true
}

// HasPowerAllocated returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasPowerAllocated() bool {
	if o != nil && !IsNil(o.PowerAllocated) {
		return true
	}

	return false
}

// SetPowerAllocated gets a reference to the given float32 and assigns it to the PowerAllocated field.
func (o *ApStatsLldpStat) SetPowerAllocated(v float32) {
	o.PowerAllocated = &v
}

// GetPowerDraw returns the PowerDraw field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetPowerDraw() float32 {
	if o == nil || IsNil(o.PowerDraw) {
		var ret float32
		return ret
	}
	return *o.PowerDraw
}

// GetPowerDrawOk returns a tuple with the PowerDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetPowerDrawOk() (*float32, bool) {
	if o == nil || IsNil(o.PowerDraw) {
		return nil, false
	}
	return o.PowerDraw, true
}

// HasPowerDraw returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasPowerDraw() bool {
	if o != nil && !IsNil(o.PowerDraw) {
		return true
	}

	return false
}

// SetPowerDraw gets a reference to the given float32 and assigns it to the PowerDraw field.
func (o *ApStatsLldpStat) SetPowerDraw(v float32) {
	o.PowerDraw = &v
}

// GetPowerRequestCount returns the PowerRequestCount field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetPowerRequestCount() int32 {
	if o == nil || IsNil(o.PowerRequestCount) {
		var ret int32
		return ret
	}
	return *o.PowerRequestCount
}

// GetPowerRequestCountOk returns a tuple with the PowerRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetPowerRequestCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PowerRequestCount) {
		return nil, false
	}
	return o.PowerRequestCount, true
}

// HasPowerRequestCount returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasPowerRequestCount() bool {
	if o != nil && !IsNil(o.PowerRequestCount) {
		return true
	}

	return false
}

// SetPowerRequestCount gets a reference to the given int32 and assigns it to the PowerRequestCount field.
func (o *ApStatsLldpStat) SetPowerRequestCount(v int32) {
	o.PowerRequestCount = &v
}

// GetPowerRequested returns the PowerRequested field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetPowerRequested() float32 {
	if o == nil || IsNil(o.PowerRequested) {
		var ret float32
		return ret
	}
	return *o.PowerRequested
}

// GetPowerRequestedOk returns a tuple with the PowerRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetPowerRequestedOk() (*float32, bool) {
	if o == nil || IsNil(o.PowerRequested) {
		return nil, false
	}
	return o.PowerRequested, true
}

// HasPowerRequested returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasPowerRequested() bool {
	if o != nil && !IsNil(o.PowerRequested) {
		return true
	}

	return false
}

// SetPowerRequested gets a reference to the given float32 and assigns it to the PowerRequested field.
func (o *ApStatsLldpStat) SetPowerRequested(v float32) {
	o.PowerRequested = &v
}

// GetSystemDesc returns the SystemDesc field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetSystemDesc() string {
	if o == nil || IsNil(o.SystemDesc) {
		var ret string
		return ret
	}
	return *o.SystemDesc
}

// GetSystemDescOk returns a tuple with the SystemDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetSystemDescOk() (*string, bool) {
	if o == nil || IsNil(o.SystemDesc) {
		return nil, false
	}
	return o.SystemDesc, true
}

// HasSystemDesc returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasSystemDesc() bool {
	if o != nil && !IsNil(o.SystemDesc) {
		return true
	}

	return false
}

// SetSystemDesc gets a reference to the given string and assigns it to the SystemDesc field.
func (o *ApStatsLldpStat) SetSystemDesc(v string) {
	o.SystemDesc = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *ApStatsLldpStat) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsLldpStat) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *ApStatsLldpStat) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *ApStatsLldpStat) SetSystemName(v string) {
	o.SystemName = &v
}

func (o ApStatsLldpStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsLldpStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChassisId) {
		toSerialize["chassis_id"] = o.ChassisId
	}
	if !IsNil(o.LldpMedSupported) {
		toSerialize["lldp_med_supported"] = o.LldpMedSupported
	}
	if !IsNil(o.MgmtAddr) {
		toSerialize["mgmt_addr"] = o.MgmtAddr
	}
	if !IsNil(o.PortDesc) {
		toSerialize["port_desc"] = o.PortDesc
	}
	if !IsNil(o.PowerAllocated) {
		toSerialize["power_allocated"] = o.PowerAllocated
	}
	if !IsNil(o.PowerDraw) {
		toSerialize["power_draw"] = o.PowerDraw
	}
	if !IsNil(o.PowerRequestCount) {
		toSerialize["power_request_count"] = o.PowerRequestCount
	}
	if !IsNil(o.PowerRequested) {
		toSerialize["power_requested"] = o.PowerRequested
	}
	if !IsNil(o.SystemDesc) {
		toSerialize["system_desc"] = o.SystemDesc
	}
	if !IsNil(o.SystemName) {
		toSerialize["system_name"] = o.SystemName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatsLldpStat) UnmarshalJSON(data []byte) (err error) {
	varApStatsLldpStat := _ApStatsLldpStat{}

	err = json.Unmarshal(data, &varApStatsLldpStat)

	if err != nil {
		return err
	}

	*o = ApStatsLldpStat(varApStatsLldpStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chassis_id")
		delete(additionalProperties, "lldp_med_supported")
		delete(additionalProperties, "mgmt_addr")
		delete(additionalProperties, "port_desc")
		delete(additionalProperties, "power_allocated")
		delete(additionalProperties, "power_draw")
		delete(additionalProperties, "power_request_count")
		delete(additionalProperties, "power_requested")
		delete(additionalProperties, "system_desc")
		delete(additionalProperties, "system_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsLldpStat struct {
	value *ApStatsLldpStat
	isSet bool
}

func (v NullableApStatsLldpStat) Get() *ApStatsLldpStat {
	return v.value
}

func (v *NullableApStatsLldpStat) Set(val *ApStatsLldpStat) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsLldpStat) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsLldpStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsLldpStat(val *ApStatsLldpStat) *NullableApStatsLldpStat {
	return &NullableApStatsLldpStat{value: val, isSet: true}
}

func (v NullableApStatsLldpStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsLldpStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


