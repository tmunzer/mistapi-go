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
	"fmt"
)

// checks if the JunosPortConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JunosPortConfig{}

// JunosPortConfig Switch port config
type JunosPortConfig struct {
	// To disable LACP support for the AE interface
	AeDisableLacp *bool `json:"ae_disable_lacp,omitempty"`
	// Users could force to use the designated AE name
	AeIdx *int32 `json:"ae_idx,omitempty"`
	// to use fast timeout
	AeLacpSlow *bool `json:"ae_lacp_slow,omitempty"`
	Aggregated *bool `json:"aggregated,omitempty"`
	// if want to generate port up/down alarm
	Critical *bool `json:"critical,omitempty"`
	Description *string `json:"description,omitempty"`
	// if `speed` and `duplex` are specified, whether to disable autonegotiation
	DisableAutoneg *bool `json:"disable_autoneg,omitempty"`
	Duplex *JunosPortConfigDuplex `json:"duplex,omitempty"`
	// Enable dynamic usage for this port. Set to `dynamic` to enable.
	DynamicUsage NullableString `json:"dynamic_usage,omitempty"`
	Esilag *bool `json:"esilag,omitempty"`
	// media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation
	Mtu *int32 `json:"mtu,omitempty"`
	// prevent helpdesk to override the port config
	NoLocalOverwrite *bool `json:"no_local_overwrite,omitempty"`
	PoeDisabled *bool `json:"poe_disabled,omitempty"`
	Speed *JunosPortConfigSpeed `json:"speed,omitempty"`
	// port usage name.   If EVPN is used, use `evpn_uplink`or `evpn_downlink`
	Usage string `json:"usage"`
	AdditionalProperties map[string]interface{}
}

type _JunosPortConfig JunosPortConfig

// NewJunosPortConfig instantiates a new JunosPortConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJunosPortConfig(usage string) *JunosPortConfig {
	this := JunosPortConfig{}
	var aeLacpSlow bool = true
	this.AeLacpSlow = &aeLacpSlow
	var aggregated bool = false
	this.Aggregated = &aggregated
	var disableAutoneg bool = false
	this.DisableAutoneg = &disableAutoneg
	var duplex JunosPortConfigDuplex = JUNOSPORTCONFIGDUPLEX_AUTO
	this.Duplex = &duplex
	var mtu int32 = 1514
	this.Mtu = &mtu
	var poeDisabled bool = false
	this.PoeDisabled = &poeDisabled
	var speed JunosPortConfigSpeed = JUNOSPORTCONFIGSPEED_AUTO
	this.Speed = &speed
	this.Usage = usage
	return &this
}

// NewJunosPortConfigWithDefaults instantiates a new JunosPortConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJunosPortConfigWithDefaults() *JunosPortConfig {
	this := JunosPortConfig{}
	var aeLacpSlow bool = true
	this.AeLacpSlow = &aeLacpSlow
	var aggregated bool = false
	this.Aggregated = &aggregated
	var disableAutoneg bool = false
	this.DisableAutoneg = &disableAutoneg
	var duplex JunosPortConfigDuplex = JUNOSPORTCONFIGDUPLEX_AUTO
	this.Duplex = &duplex
	var mtu int32 = 1514
	this.Mtu = &mtu
	var poeDisabled bool = false
	this.PoeDisabled = &poeDisabled
	var speed JunosPortConfigSpeed = JUNOSPORTCONFIGSPEED_AUTO
	this.Speed = &speed
	return &this
}

// GetAeDisableLacp returns the AeDisableLacp field value if set, zero value otherwise.
func (o *JunosPortConfig) GetAeDisableLacp() bool {
	if o == nil || IsNil(o.AeDisableLacp) {
		var ret bool
		return ret
	}
	return *o.AeDisableLacp
}

// GetAeDisableLacpOk returns a tuple with the AeDisableLacp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetAeDisableLacpOk() (*bool, bool) {
	if o == nil || IsNil(o.AeDisableLacp) {
		return nil, false
	}
	return o.AeDisableLacp, true
}

// HasAeDisableLacp returns a boolean if a field has been set.
func (o *JunosPortConfig) HasAeDisableLacp() bool {
	if o != nil && !IsNil(o.AeDisableLacp) {
		return true
	}

	return false
}

// SetAeDisableLacp gets a reference to the given bool and assigns it to the AeDisableLacp field.
func (o *JunosPortConfig) SetAeDisableLacp(v bool) {
	o.AeDisableLacp = &v
}

// GetAeIdx returns the AeIdx field value if set, zero value otherwise.
func (o *JunosPortConfig) GetAeIdx() int32 {
	if o == nil || IsNil(o.AeIdx) {
		var ret int32
		return ret
	}
	return *o.AeIdx
}

// GetAeIdxOk returns a tuple with the AeIdx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetAeIdxOk() (*int32, bool) {
	if o == nil || IsNil(o.AeIdx) {
		return nil, false
	}
	return o.AeIdx, true
}

// HasAeIdx returns a boolean if a field has been set.
func (o *JunosPortConfig) HasAeIdx() bool {
	if o != nil && !IsNil(o.AeIdx) {
		return true
	}

	return false
}

// SetAeIdx gets a reference to the given int32 and assigns it to the AeIdx field.
func (o *JunosPortConfig) SetAeIdx(v int32) {
	o.AeIdx = &v
}

// GetAeLacpSlow returns the AeLacpSlow field value if set, zero value otherwise.
func (o *JunosPortConfig) GetAeLacpSlow() bool {
	if o == nil || IsNil(o.AeLacpSlow) {
		var ret bool
		return ret
	}
	return *o.AeLacpSlow
}

// GetAeLacpSlowOk returns a tuple with the AeLacpSlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetAeLacpSlowOk() (*bool, bool) {
	if o == nil || IsNil(o.AeLacpSlow) {
		return nil, false
	}
	return o.AeLacpSlow, true
}

// HasAeLacpSlow returns a boolean if a field has been set.
func (o *JunosPortConfig) HasAeLacpSlow() bool {
	if o != nil && !IsNil(o.AeLacpSlow) {
		return true
	}

	return false
}

// SetAeLacpSlow gets a reference to the given bool and assigns it to the AeLacpSlow field.
func (o *JunosPortConfig) SetAeLacpSlow(v bool) {
	o.AeLacpSlow = &v
}

// GetAggregated returns the Aggregated field value if set, zero value otherwise.
func (o *JunosPortConfig) GetAggregated() bool {
	if o == nil || IsNil(o.Aggregated) {
		var ret bool
		return ret
	}
	return *o.Aggregated
}

// GetAggregatedOk returns a tuple with the Aggregated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetAggregatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Aggregated) {
		return nil, false
	}
	return o.Aggregated, true
}

// HasAggregated returns a boolean if a field has been set.
func (o *JunosPortConfig) HasAggregated() bool {
	if o != nil && !IsNil(o.Aggregated) {
		return true
	}

	return false
}

// SetAggregated gets a reference to the given bool and assigns it to the Aggregated field.
func (o *JunosPortConfig) SetAggregated(v bool) {
	o.Aggregated = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *JunosPortConfig) GetCritical() bool {
	if o == nil || IsNil(o.Critical) {
		var ret bool
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetCriticalOk() (*bool, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *JunosPortConfig) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given bool and assigns it to the Critical field.
func (o *JunosPortConfig) SetCritical(v bool) {
	o.Critical = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JunosPortConfig) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JunosPortConfig) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JunosPortConfig) SetDescription(v string) {
	o.Description = &v
}

// GetDisableAutoneg returns the DisableAutoneg field value if set, zero value otherwise.
func (o *JunosPortConfig) GetDisableAutoneg() bool {
	if o == nil || IsNil(o.DisableAutoneg) {
		var ret bool
		return ret
	}
	return *o.DisableAutoneg
}

// GetDisableAutonegOk returns a tuple with the DisableAutoneg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetDisableAutonegOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAutoneg) {
		return nil, false
	}
	return o.DisableAutoneg, true
}

// HasDisableAutoneg returns a boolean if a field has been set.
func (o *JunosPortConfig) HasDisableAutoneg() bool {
	if o != nil && !IsNil(o.DisableAutoneg) {
		return true
	}

	return false
}

// SetDisableAutoneg gets a reference to the given bool and assigns it to the DisableAutoneg field.
func (o *JunosPortConfig) SetDisableAutoneg(v bool) {
	o.DisableAutoneg = &v
}

// GetDuplex returns the Duplex field value if set, zero value otherwise.
func (o *JunosPortConfig) GetDuplex() JunosPortConfigDuplex {
	if o == nil || IsNil(o.Duplex) {
		var ret JunosPortConfigDuplex
		return ret
	}
	return *o.Duplex
}

// GetDuplexOk returns a tuple with the Duplex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetDuplexOk() (*JunosPortConfigDuplex, bool) {
	if o == nil || IsNil(o.Duplex) {
		return nil, false
	}
	return o.Duplex, true
}

// HasDuplex returns a boolean if a field has been set.
func (o *JunosPortConfig) HasDuplex() bool {
	if o != nil && !IsNil(o.Duplex) {
		return true
	}

	return false
}

// SetDuplex gets a reference to the given JunosPortConfigDuplex and assigns it to the Duplex field.
func (o *JunosPortConfig) SetDuplex(v JunosPortConfigDuplex) {
	o.Duplex = &v
}

// GetDynamicUsage returns the DynamicUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JunosPortConfig) GetDynamicUsage() string {
	if o == nil || IsNil(o.DynamicUsage.Get()) {
		var ret string
		return ret
	}
	return *o.DynamicUsage.Get()
}

// GetDynamicUsageOk returns a tuple with the DynamicUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JunosPortConfig) GetDynamicUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DynamicUsage.Get(), o.DynamicUsage.IsSet()
}

// HasDynamicUsage returns a boolean if a field has been set.
func (o *JunosPortConfig) HasDynamicUsage() bool {
	if o != nil && o.DynamicUsage.IsSet() {
		return true
	}

	return false
}

// SetDynamicUsage gets a reference to the given NullableString and assigns it to the DynamicUsage field.
func (o *JunosPortConfig) SetDynamicUsage(v string) {
	o.DynamicUsage.Set(&v)
}
// SetDynamicUsageNil sets the value for DynamicUsage to be an explicit nil
func (o *JunosPortConfig) SetDynamicUsageNil() {
	o.DynamicUsage.Set(nil)
}

// UnsetDynamicUsage ensures that no value is present for DynamicUsage, not even an explicit nil
func (o *JunosPortConfig) UnsetDynamicUsage() {
	o.DynamicUsage.Unset()
}

// GetEsilag returns the Esilag field value if set, zero value otherwise.
func (o *JunosPortConfig) GetEsilag() bool {
	if o == nil || IsNil(o.Esilag) {
		var ret bool
		return ret
	}
	return *o.Esilag
}

// GetEsilagOk returns a tuple with the Esilag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetEsilagOk() (*bool, bool) {
	if o == nil || IsNil(o.Esilag) {
		return nil, false
	}
	return o.Esilag, true
}

// HasEsilag returns a boolean if a field has been set.
func (o *JunosPortConfig) HasEsilag() bool {
	if o != nil && !IsNil(o.Esilag) {
		return true
	}

	return false
}

// SetEsilag gets a reference to the given bool and assigns it to the Esilag field.
func (o *JunosPortConfig) SetEsilag(v bool) {
	o.Esilag = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *JunosPortConfig) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu) {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetMtuOk() (*int32, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *JunosPortConfig) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *JunosPortConfig) SetMtu(v int32) {
	o.Mtu = &v
}

// GetNoLocalOverwrite returns the NoLocalOverwrite field value if set, zero value otherwise.
func (o *JunosPortConfig) GetNoLocalOverwrite() bool {
	if o == nil || IsNil(o.NoLocalOverwrite) {
		var ret bool
		return ret
	}
	return *o.NoLocalOverwrite
}

// GetNoLocalOverwriteOk returns a tuple with the NoLocalOverwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetNoLocalOverwriteOk() (*bool, bool) {
	if o == nil || IsNil(o.NoLocalOverwrite) {
		return nil, false
	}
	return o.NoLocalOverwrite, true
}

// HasNoLocalOverwrite returns a boolean if a field has been set.
func (o *JunosPortConfig) HasNoLocalOverwrite() bool {
	if o != nil && !IsNil(o.NoLocalOverwrite) {
		return true
	}

	return false
}

// SetNoLocalOverwrite gets a reference to the given bool and assigns it to the NoLocalOverwrite field.
func (o *JunosPortConfig) SetNoLocalOverwrite(v bool) {
	o.NoLocalOverwrite = &v
}

// GetPoeDisabled returns the PoeDisabled field value if set, zero value otherwise.
func (o *JunosPortConfig) GetPoeDisabled() bool {
	if o == nil || IsNil(o.PoeDisabled) {
		var ret bool
		return ret
	}
	return *o.PoeDisabled
}

// GetPoeDisabledOk returns a tuple with the PoeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetPoeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PoeDisabled) {
		return nil, false
	}
	return o.PoeDisabled, true
}

// HasPoeDisabled returns a boolean if a field has been set.
func (o *JunosPortConfig) HasPoeDisabled() bool {
	if o != nil && !IsNil(o.PoeDisabled) {
		return true
	}

	return false
}

// SetPoeDisabled gets a reference to the given bool and assigns it to the PoeDisabled field.
func (o *JunosPortConfig) SetPoeDisabled(v bool) {
	o.PoeDisabled = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *JunosPortConfig) GetSpeed() JunosPortConfigSpeed {
	if o == nil || IsNil(o.Speed) {
		var ret JunosPortConfigSpeed
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetSpeedOk() (*JunosPortConfigSpeed, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *JunosPortConfig) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given JunosPortConfigSpeed and assigns it to the Speed field.
func (o *JunosPortConfig) SetSpeed(v JunosPortConfigSpeed) {
	o.Speed = &v
}

// GetUsage returns the Usage field value
func (o *JunosPortConfig) GetUsage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *JunosPortConfig) GetUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *JunosPortConfig) SetUsage(v string) {
	o.Usage = v
}

func (o JunosPortConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JunosPortConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AeDisableLacp) {
		toSerialize["ae_disable_lacp"] = o.AeDisableLacp
	}
	if !IsNil(o.AeIdx) {
		toSerialize["ae_idx"] = o.AeIdx
	}
	if !IsNil(o.AeLacpSlow) {
		toSerialize["ae_lacp_slow"] = o.AeLacpSlow
	}
	if !IsNil(o.Aggregated) {
		toSerialize["aggregated"] = o.Aggregated
	}
	if !IsNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisableAutoneg) {
		toSerialize["disable_autoneg"] = o.DisableAutoneg
	}
	if !IsNil(o.Duplex) {
		toSerialize["duplex"] = o.Duplex
	}
	if o.DynamicUsage.IsSet() {
		toSerialize["dynamic_usage"] = o.DynamicUsage.Get()
	}
	if !IsNil(o.Esilag) {
		toSerialize["esilag"] = o.Esilag
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !IsNil(o.NoLocalOverwrite) {
		toSerialize["no_local_overwrite"] = o.NoLocalOverwrite
	}
	if !IsNil(o.PoeDisabled) {
		toSerialize["poe_disabled"] = o.PoeDisabled
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	toSerialize["usage"] = o.Usage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JunosPortConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"usage",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varJunosPortConfig := _JunosPortConfig{}

	err = json.Unmarshal(data, &varJunosPortConfig)

	if err != nil {
		return err
	}

	*o = JunosPortConfig(varJunosPortConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ae_disable_lacp")
		delete(additionalProperties, "ae_idx")
		delete(additionalProperties, "ae_lacp_slow")
		delete(additionalProperties, "aggregated")
		delete(additionalProperties, "critical")
		delete(additionalProperties, "description")
		delete(additionalProperties, "disable_autoneg")
		delete(additionalProperties, "duplex")
		delete(additionalProperties, "dynamic_usage")
		delete(additionalProperties, "esilag")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "no_local_overwrite")
		delete(additionalProperties, "poe_disabled")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJunosPortConfig struct {
	value *JunosPortConfig
	isSet bool
}

func (v NullableJunosPortConfig) Get() *JunosPortConfig {
	return v.value
}

func (v *NullableJunosPortConfig) Set(val *JunosPortConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJunosPortConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJunosPortConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJunosPortConfig(val *JunosPortConfig) *NullableJunosPortConfig {
	return &NullableJunosPortConfig{value: val, isSet: true}
}

func (v NullableJunosPortConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJunosPortConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


