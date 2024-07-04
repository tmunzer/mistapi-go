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

// checks if the ConfigSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSwitch{}

// ConfigSwitch Switch settings
type ConfigSwitch struct {
	// ap_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used.
	ApAffinityThreshold *int32 `json:"ap_affinity_threshold,omitempty"`
	// Set Banners for switches. Allows markup formatting
	CliBanner *string `json:"cli_banner,omitempty"`
	// Sets timeout for switches
	CliIdleTimeout *int32 `json:"cli_idle_timeout,omitempty"`
	// the rollback timer for commit confirmed
	ConfigRevertTimer *int32 `json:"config_revert_timer,omitempty"`
	// Enable to provide the FQDN with DHCP option 81
	DhcpOptionFqdn *bool `json:"dhcp_option_fqdn,omitempty"`
	// Property key is the user name. For Local user authentication
	LocalAccounts *map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
	MxedgeProxyHost *string `json:"mxedge_proxy_host,omitempty"`
	MxedgeProxyPort *int32 `json:"mxedge_proxy_port,omitempty"`
	ProtectRe *ProtectRe `json:"protect_re,omitempty"`
	Radius *ConfigSwitchRadius `json:"radius,omitempty"`
	RootPassword *string `json:"root_password,omitempty"`
	Tacacs *Tacacs `json:"tacacs,omitempty"`
	// to use mxedge as proxy
	UseMxedgeProxy *bool `json:"use_mxedge_proxy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigSwitch ConfigSwitch

// NewConfigSwitch instantiates a new ConfigSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSwitch() *ConfigSwitch {
	this := ConfigSwitch{}
	var apAffinityThreshold int32 = 10
	this.ApAffinityThreshold = &apAffinityThreshold
	var configRevertTimer int32 = 10
	this.ConfigRevertTimer = &configRevertTimer
	var dhcpOptionFqdn bool = false
	this.DhcpOptionFqdn = &dhcpOptionFqdn
	var mxedgeProxyPort int32 = 2222
	this.MxedgeProxyPort = &mxedgeProxyPort
	return &this
}

// NewConfigSwitchWithDefaults instantiates a new ConfigSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSwitchWithDefaults() *ConfigSwitch {
	this := ConfigSwitch{}
	var apAffinityThreshold int32 = 10
	this.ApAffinityThreshold = &apAffinityThreshold
	var configRevertTimer int32 = 10
	this.ConfigRevertTimer = &configRevertTimer
	var dhcpOptionFqdn bool = false
	this.DhcpOptionFqdn = &dhcpOptionFqdn
	var mxedgeProxyPort int32 = 2222
	this.MxedgeProxyPort = &mxedgeProxyPort
	return &this
}

// GetApAffinityThreshold returns the ApAffinityThreshold field value if set, zero value otherwise.
func (o *ConfigSwitch) GetApAffinityThreshold() int32 {
	if o == nil || IsNil(o.ApAffinityThreshold) {
		var ret int32
		return ret
	}
	return *o.ApAffinityThreshold
}

// GetApAffinityThresholdOk returns a tuple with the ApAffinityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetApAffinityThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApAffinityThreshold) {
		return nil, false
	}
	return o.ApAffinityThreshold, true
}

// HasApAffinityThreshold returns a boolean if a field has been set.
func (o *ConfigSwitch) HasApAffinityThreshold() bool {
	if o != nil && !IsNil(o.ApAffinityThreshold) {
		return true
	}

	return false
}

// SetApAffinityThreshold gets a reference to the given int32 and assigns it to the ApAffinityThreshold field.
func (o *ConfigSwitch) SetApAffinityThreshold(v int32) {
	o.ApAffinityThreshold = &v
}

// GetCliBanner returns the CliBanner field value if set, zero value otherwise.
func (o *ConfigSwitch) GetCliBanner() string {
	if o == nil || IsNil(o.CliBanner) {
		var ret string
		return ret
	}
	return *o.CliBanner
}

// GetCliBannerOk returns a tuple with the CliBanner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetCliBannerOk() (*string, bool) {
	if o == nil || IsNil(o.CliBanner) {
		return nil, false
	}
	return o.CliBanner, true
}

// HasCliBanner returns a boolean if a field has been set.
func (o *ConfigSwitch) HasCliBanner() bool {
	if o != nil && !IsNil(o.CliBanner) {
		return true
	}

	return false
}

// SetCliBanner gets a reference to the given string and assigns it to the CliBanner field.
func (o *ConfigSwitch) SetCliBanner(v string) {
	o.CliBanner = &v
}

// GetCliIdleTimeout returns the CliIdleTimeout field value if set, zero value otherwise.
func (o *ConfigSwitch) GetCliIdleTimeout() int32 {
	if o == nil || IsNil(o.CliIdleTimeout) {
		var ret int32
		return ret
	}
	return *o.CliIdleTimeout
}

// GetCliIdleTimeoutOk returns a tuple with the CliIdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetCliIdleTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.CliIdleTimeout) {
		return nil, false
	}
	return o.CliIdleTimeout, true
}

// HasCliIdleTimeout returns a boolean if a field has been set.
func (o *ConfigSwitch) HasCliIdleTimeout() bool {
	if o != nil && !IsNil(o.CliIdleTimeout) {
		return true
	}

	return false
}

// SetCliIdleTimeout gets a reference to the given int32 and assigns it to the CliIdleTimeout field.
func (o *ConfigSwitch) SetCliIdleTimeout(v int32) {
	o.CliIdleTimeout = &v
}

// GetConfigRevertTimer returns the ConfigRevertTimer field value if set, zero value otherwise.
func (o *ConfigSwitch) GetConfigRevertTimer() int32 {
	if o == nil || IsNil(o.ConfigRevertTimer) {
		var ret int32
		return ret
	}
	return *o.ConfigRevertTimer
}

// GetConfigRevertTimerOk returns a tuple with the ConfigRevertTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetConfigRevertTimerOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfigRevertTimer) {
		return nil, false
	}
	return o.ConfigRevertTimer, true
}

// HasConfigRevertTimer returns a boolean if a field has been set.
func (o *ConfigSwitch) HasConfigRevertTimer() bool {
	if o != nil && !IsNil(o.ConfigRevertTimer) {
		return true
	}

	return false
}

// SetConfigRevertTimer gets a reference to the given int32 and assigns it to the ConfigRevertTimer field.
func (o *ConfigSwitch) SetConfigRevertTimer(v int32) {
	o.ConfigRevertTimer = &v
}

// GetDhcpOptionFqdn returns the DhcpOptionFqdn field value if set, zero value otherwise.
func (o *ConfigSwitch) GetDhcpOptionFqdn() bool {
	if o == nil || IsNil(o.DhcpOptionFqdn) {
		var ret bool
		return ret
	}
	return *o.DhcpOptionFqdn
}

// GetDhcpOptionFqdnOk returns a tuple with the DhcpOptionFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetDhcpOptionFqdnOk() (*bool, bool) {
	if o == nil || IsNil(o.DhcpOptionFqdn) {
		return nil, false
	}
	return o.DhcpOptionFqdn, true
}

// HasDhcpOptionFqdn returns a boolean if a field has been set.
func (o *ConfigSwitch) HasDhcpOptionFqdn() bool {
	if o != nil && !IsNil(o.DhcpOptionFqdn) {
		return true
	}

	return false
}

// SetDhcpOptionFqdn gets a reference to the given bool and assigns it to the DhcpOptionFqdn field.
func (o *ConfigSwitch) SetDhcpOptionFqdn(v bool) {
	o.DhcpOptionFqdn = &v
}

// GetLocalAccounts returns the LocalAccounts field value if set, zero value otherwise.
func (o *ConfigSwitch) GetLocalAccounts() map[string]ConfigSwitchLocalAccountsUser {
	if o == nil || IsNil(o.LocalAccounts) {
		var ret map[string]ConfigSwitchLocalAccountsUser
		return ret
	}
	return *o.LocalAccounts
}

// GetLocalAccountsOk returns a tuple with the LocalAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetLocalAccountsOk() (*map[string]ConfigSwitchLocalAccountsUser, bool) {
	if o == nil || IsNil(o.LocalAccounts) {
		return nil, false
	}
	return o.LocalAccounts, true
}

// HasLocalAccounts returns a boolean if a field has been set.
func (o *ConfigSwitch) HasLocalAccounts() bool {
	if o != nil && !IsNil(o.LocalAccounts) {
		return true
	}

	return false
}

// SetLocalAccounts gets a reference to the given map[string]ConfigSwitchLocalAccountsUser and assigns it to the LocalAccounts field.
func (o *ConfigSwitch) SetLocalAccounts(v map[string]ConfigSwitchLocalAccountsUser) {
	o.LocalAccounts = &v
}

// GetMxedgeProxyHost returns the MxedgeProxyHost field value if set, zero value otherwise.
func (o *ConfigSwitch) GetMxedgeProxyHost() string {
	if o == nil || IsNil(o.MxedgeProxyHost) {
		var ret string
		return ret
	}
	return *o.MxedgeProxyHost
}

// GetMxedgeProxyHostOk returns a tuple with the MxedgeProxyHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetMxedgeProxyHostOk() (*string, bool) {
	if o == nil || IsNil(o.MxedgeProxyHost) {
		return nil, false
	}
	return o.MxedgeProxyHost, true
}

// HasMxedgeProxyHost returns a boolean if a field has been set.
func (o *ConfigSwitch) HasMxedgeProxyHost() bool {
	if o != nil && !IsNil(o.MxedgeProxyHost) {
		return true
	}

	return false
}

// SetMxedgeProxyHost gets a reference to the given string and assigns it to the MxedgeProxyHost field.
func (o *ConfigSwitch) SetMxedgeProxyHost(v string) {
	o.MxedgeProxyHost = &v
}

// GetMxedgeProxyPort returns the MxedgeProxyPort field value if set, zero value otherwise.
func (o *ConfigSwitch) GetMxedgeProxyPort() int32 {
	if o == nil || IsNil(o.MxedgeProxyPort) {
		var ret int32
		return ret
	}
	return *o.MxedgeProxyPort
}

// GetMxedgeProxyPortOk returns a tuple with the MxedgeProxyPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetMxedgeProxyPortOk() (*int32, bool) {
	if o == nil || IsNil(o.MxedgeProxyPort) {
		return nil, false
	}
	return o.MxedgeProxyPort, true
}

// HasMxedgeProxyPort returns a boolean if a field has been set.
func (o *ConfigSwitch) HasMxedgeProxyPort() bool {
	if o != nil && !IsNil(o.MxedgeProxyPort) {
		return true
	}

	return false
}

// SetMxedgeProxyPort gets a reference to the given int32 and assigns it to the MxedgeProxyPort field.
func (o *ConfigSwitch) SetMxedgeProxyPort(v int32) {
	o.MxedgeProxyPort = &v
}

// GetProtectRe returns the ProtectRe field value if set, zero value otherwise.
func (o *ConfigSwitch) GetProtectRe() ProtectRe {
	if o == nil || IsNil(o.ProtectRe) {
		var ret ProtectRe
		return ret
	}
	return *o.ProtectRe
}

// GetProtectReOk returns a tuple with the ProtectRe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetProtectReOk() (*ProtectRe, bool) {
	if o == nil || IsNil(o.ProtectRe) {
		return nil, false
	}
	return o.ProtectRe, true
}

// HasProtectRe returns a boolean if a field has been set.
func (o *ConfigSwitch) HasProtectRe() bool {
	if o != nil && !IsNil(o.ProtectRe) {
		return true
	}

	return false
}

// SetProtectRe gets a reference to the given ProtectRe and assigns it to the ProtectRe field.
func (o *ConfigSwitch) SetProtectRe(v ProtectRe) {
	o.ProtectRe = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *ConfigSwitch) GetRadius() ConfigSwitchRadius {
	if o == nil || IsNil(o.Radius) {
		var ret ConfigSwitchRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetRadiusOk() (*ConfigSwitchRadius, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *ConfigSwitch) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given ConfigSwitchRadius and assigns it to the Radius field.
func (o *ConfigSwitch) SetRadius(v ConfigSwitchRadius) {
	o.Radius = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *ConfigSwitch) GetRootPassword() string {
	if o == nil || IsNil(o.RootPassword) {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetRootPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RootPassword) {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *ConfigSwitch) HasRootPassword() bool {
	if o != nil && !IsNil(o.RootPassword) {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *ConfigSwitch) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetTacacs returns the Tacacs field value if set, zero value otherwise.
func (o *ConfigSwitch) GetTacacs() Tacacs {
	if o == nil || IsNil(o.Tacacs) {
		var ret Tacacs
		return ret
	}
	return *o.Tacacs
}

// GetTacacsOk returns a tuple with the Tacacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetTacacsOk() (*Tacacs, bool) {
	if o == nil || IsNil(o.Tacacs) {
		return nil, false
	}
	return o.Tacacs, true
}

// HasTacacs returns a boolean if a field has been set.
func (o *ConfigSwitch) HasTacacs() bool {
	if o != nil && !IsNil(o.Tacacs) {
		return true
	}

	return false
}

// SetTacacs gets a reference to the given Tacacs and assigns it to the Tacacs field.
func (o *ConfigSwitch) SetTacacs(v Tacacs) {
	o.Tacacs = &v
}

// GetUseMxedgeProxy returns the UseMxedgeProxy field value if set, zero value otherwise.
func (o *ConfigSwitch) GetUseMxedgeProxy() bool {
	if o == nil || IsNil(o.UseMxedgeProxy) {
		var ret bool
		return ret
	}
	return *o.UseMxedgeProxy
}

// GetUseMxedgeProxyOk returns a tuple with the UseMxedgeProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSwitch) GetUseMxedgeProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMxedgeProxy) {
		return nil, false
	}
	return o.UseMxedgeProxy, true
}

// HasUseMxedgeProxy returns a boolean if a field has been set.
func (o *ConfigSwitch) HasUseMxedgeProxy() bool {
	if o != nil && !IsNil(o.UseMxedgeProxy) {
		return true
	}

	return false
}

// SetUseMxedgeProxy gets a reference to the given bool and assigns it to the UseMxedgeProxy field.
func (o *ConfigSwitch) SetUseMxedgeProxy(v bool) {
	o.UseMxedgeProxy = &v
}

func (o ConfigSwitch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApAffinityThreshold) {
		toSerialize["ap_affinity_threshold"] = o.ApAffinityThreshold
	}
	if !IsNil(o.CliBanner) {
		toSerialize["cli_banner"] = o.CliBanner
	}
	if !IsNil(o.CliIdleTimeout) {
		toSerialize["cli_idle_timeout"] = o.CliIdleTimeout
	}
	if !IsNil(o.ConfigRevertTimer) {
		toSerialize["config_revert_timer"] = o.ConfigRevertTimer
	}
	if !IsNil(o.DhcpOptionFqdn) {
		toSerialize["dhcp_option_fqdn"] = o.DhcpOptionFqdn
	}
	if !IsNil(o.LocalAccounts) {
		toSerialize["local_accounts"] = o.LocalAccounts
	}
	if !IsNil(o.MxedgeProxyHost) {
		toSerialize["mxedge_proxy_host"] = o.MxedgeProxyHost
	}
	if !IsNil(o.MxedgeProxyPort) {
		toSerialize["mxedge_proxy_port"] = o.MxedgeProxyPort
	}
	if !IsNil(o.ProtectRe) {
		toSerialize["protect_re"] = o.ProtectRe
	}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !IsNil(o.RootPassword) {
		toSerialize["root_password"] = o.RootPassword
	}
	if !IsNil(o.Tacacs) {
		toSerialize["tacacs"] = o.Tacacs
	}
	if !IsNil(o.UseMxedgeProxy) {
		toSerialize["use_mxedge_proxy"] = o.UseMxedgeProxy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigSwitch) UnmarshalJSON(data []byte) (err error) {
	varConfigSwitch := _ConfigSwitch{}

	err = json.Unmarshal(data, &varConfigSwitch)

	if err != nil {
		return err
	}

	*o = ConfigSwitch(varConfigSwitch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_affinity_threshold")
		delete(additionalProperties, "cli_banner")
		delete(additionalProperties, "cli_idle_timeout")
		delete(additionalProperties, "config_revert_timer")
		delete(additionalProperties, "dhcp_option_fqdn")
		delete(additionalProperties, "local_accounts")
		delete(additionalProperties, "mxedge_proxy_host")
		delete(additionalProperties, "mxedge_proxy_port")
		delete(additionalProperties, "protect_re")
		delete(additionalProperties, "radius")
		delete(additionalProperties, "root_password")
		delete(additionalProperties, "tacacs")
		delete(additionalProperties, "use_mxedge_proxy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigSwitch struct {
	value *ConfigSwitch
	isSet bool
}

func (v NullableConfigSwitch) Get() *ConfigSwitch {
	return v.value
}

func (v *NullableConfigSwitch) Set(val *ConfigSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSwitch(val *ConfigSwitch) *NullableConfigSwitch {
	return &NullableConfigSwitch{value: val, isSet: true}
}

func (v NullableConfigSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


