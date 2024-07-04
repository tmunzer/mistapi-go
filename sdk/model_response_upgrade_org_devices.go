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

// checks if the ResponseUpgradeOrgDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseUpgradeOrgDevices{}

// ResponseUpgradeOrgDevices struct for ResponseUpgradeOrgDevices
type ResponseUpgradeOrgDevices struct {
	// whether to allow local AP-to-AP FW upgrade
	EnableP2p *bool `json:"enable_p2p,omitempty"`
	// whether to force upgrade when requested version is same as running version
	Force *bool `json:"force,omitempty"`
	Id *string `json:"id,omitempty"`
	Strategy *DeviceUpgradeStrategy `json:"strategy,omitempty"`
	// version to upgrade to
	TargetVersion *string `json:"target_version,omitempty"`
	Upgrades []ResponseUpgradeOrgDevice `json:"upgrades,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseUpgradeOrgDevices ResponseUpgradeOrgDevices

// NewResponseUpgradeOrgDevices instantiates a new ResponseUpgradeOrgDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseUpgradeOrgDevices() *ResponseUpgradeOrgDevices {
	this := ResponseUpgradeOrgDevices{}
	var strategy DeviceUpgradeStrategy = DEVICEUPGRADESTRATEGY_BIG_BANG
	this.Strategy = &strategy
	return &this
}

// NewResponseUpgradeOrgDevicesWithDefaults instantiates a new ResponseUpgradeOrgDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseUpgradeOrgDevicesWithDefaults() *ResponseUpgradeOrgDevices {
	this := ResponseUpgradeOrgDevices{}
	var strategy DeviceUpgradeStrategy = DEVICEUPGRADESTRATEGY_BIG_BANG
	this.Strategy = &strategy
	return &this
}

// GetEnableP2p returns the EnableP2p field value if set, zero value otherwise.
func (o *ResponseUpgradeOrgDevices) GetEnableP2p() bool {
	if o == nil || IsNil(o.EnableP2p) {
		var ret bool
		return ret
	}
	return *o.EnableP2p
}

// GetEnableP2pOk returns a tuple with the EnableP2p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeOrgDevices) GetEnableP2pOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableP2p) {
		return nil, false
	}
	return o.EnableP2p, true
}

// HasEnableP2p returns a boolean if a field has been set.
func (o *ResponseUpgradeOrgDevices) HasEnableP2p() bool {
	if o != nil && !IsNil(o.EnableP2p) {
		return true
	}

	return false
}

// SetEnableP2p gets a reference to the given bool and assigns it to the EnableP2p field.
func (o *ResponseUpgradeOrgDevices) SetEnableP2p(v bool) {
	o.EnableP2p = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *ResponseUpgradeOrgDevices) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeOrgDevices) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *ResponseUpgradeOrgDevices) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *ResponseUpgradeOrgDevices) SetForce(v bool) {
	o.Force = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponseUpgradeOrgDevices) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeOrgDevices) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponseUpgradeOrgDevices) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResponseUpgradeOrgDevices) SetId(v string) {
	o.Id = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ResponseUpgradeOrgDevices) GetStrategy() DeviceUpgradeStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret DeviceUpgradeStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeOrgDevices) GetStrategyOk() (*DeviceUpgradeStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ResponseUpgradeOrgDevices) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given DeviceUpgradeStrategy and assigns it to the Strategy field.
func (o *ResponseUpgradeOrgDevices) SetStrategy(v DeviceUpgradeStrategy) {
	o.Strategy = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *ResponseUpgradeOrgDevices) GetTargetVersion() string {
	if o == nil || IsNil(o.TargetVersion) {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeOrgDevices) GetTargetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetVersion) {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *ResponseUpgradeOrgDevices) HasTargetVersion() bool {
	if o != nil && !IsNil(o.TargetVersion) {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *ResponseUpgradeOrgDevices) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// GetUpgrades returns the Upgrades field value if set, zero value otherwise.
func (o *ResponseUpgradeOrgDevices) GetUpgrades() []ResponseUpgradeOrgDevice {
	if o == nil || IsNil(o.Upgrades) {
		var ret []ResponseUpgradeOrgDevice
		return ret
	}
	return o.Upgrades
}

// GetUpgradesOk returns a tuple with the Upgrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeOrgDevices) GetUpgradesOk() ([]ResponseUpgradeOrgDevice, bool) {
	if o == nil || IsNil(o.Upgrades) {
		return nil, false
	}
	return o.Upgrades, true
}

// HasUpgrades returns a boolean if a field has been set.
func (o *ResponseUpgradeOrgDevices) HasUpgrades() bool {
	if o != nil && !IsNil(o.Upgrades) {
		return true
	}

	return false
}

// SetUpgrades gets a reference to the given []ResponseUpgradeOrgDevice and assigns it to the Upgrades field.
func (o *ResponseUpgradeOrgDevices) SetUpgrades(v []ResponseUpgradeOrgDevice) {
	o.Upgrades = v
}

func (o ResponseUpgradeOrgDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseUpgradeOrgDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableP2p) {
		toSerialize["enable_p2p"] = o.EnableP2p
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.TargetVersion) {
		toSerialize["target_version"] = o.TargetVersion
	}
	if !IsNil(o.Upgrades) {
		toSerialize["upgrades"] = o.Upgrades
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseUpgradeOrgDevices) UnmarshalJSON(data []byte) (err error) {
	varResponseUpgradeOrgDevices := _ResponseUpgradeOrgDevices{}

	err = json.Unmarshal(data, &varResponseUpgradeOrgDevices)

	if err != nil {
		return err
	}

	*o = ResponseUpgradeOrgDevices(varResponseUpgradeOrgDevices)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_p2p")
		delete(additionalProperties, "force")
		delete(additionalProperties, "id")
		delete(additionalProperties, "strategy")
		delete(additionalProperties, "target_version")
		delete(additionalProperties, "upgrades")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseUpgradeOrgDevices struct {
	value *ResponseUpgradeOrgDevices
	isSet bool
}

func (v NullableResponseUpgradeOrgDevices) Get() *ResponseUpgradeOrgDevices {
	return v.value
}

func (v *NullableResponseUpgradeOrgDevices) Set(val *ResponseUpgradeOrgDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseUpgradeOrgDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseUpgradeOrgDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseUpgradeOrgDevices(val *ResponseUpgradeOrgDevices) *NullableResponseUpgradeOrgDevices {
	return &NullableResponseUpgradeOrgDevices{value: val, isSet: true}
}

func (v NullableResponseUpgradeOrgDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseUpgradeOrgDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

