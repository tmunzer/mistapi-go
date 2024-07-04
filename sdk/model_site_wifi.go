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

// checks if the SiteWifi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteWifi{}

// SiteWifi Wi-Fi site settings
type SiteWifi struct {
	CiscoEnabled *bool `json:"cisco_enabled,omitempty"`
	// whether to disable 11k
	Disable11k *bool `json:"disable_11k,omitempty"`
	DisableRadiosWhenPowerConstrained *bool `json:"disable_radios_when_power_constrained,omitempty"`
	// when proxy_arp is enabled, check for arp spoofing.
	EnableArpSpoofCheck *bool `json:"enable_arp_spoof_check,omitempty"`
	EnableSharedRadioScanning *bool `json:"enable_shared_radio_scanning,omitempty"`
	// enable WIFI feature (using SUB-MAN license)
	Enabled *bool `json:"enabled,omitempty"`
	// whether to locate connected clients
	LocateConnected *bool `json:"locate_connected,omitempty"`
	// whether to locate unconnected clients
	LocateUnconnected *bool `json:"locate_unconnected,omitempty"`
	// whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will distrupt the connection. If roaming is desired, keep it disabled.
	MeshAllowDfs *bool `json:"mesh_allow_dfs,omitempty"`
	// used to enable/disable CRM
	MeshEnableCrm *bool `json:"mesh_enable_crm,omitempty"`
	// whether to enable Mesh feature for the site
	MeshEnabled *bool `json:"mesh_enabled,omitempty"`
	// optional passphrase of mesh networking, default is generated randomly
	MeshPsk NullableString `json:"mesh_psk,omitempty"`
	// optional ssid of mesh networking, default is based on site_id
	MeshSsid NullableString `json:"mesh_ssid,omitempty"`
	ProxyArp NullableSiteWifiProxyArp `json:"proxy_arp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteWifi SiteWifi

// NewSiteWifi instantiates a new SiteWifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteWifi() *SiteWifi {
	this := SiteWifi{}
	var ciscoEnabled bool = true
	this.CiscoEnabled = &ciscoEnabled
	var disable11k bool = false
	this.Disable11k = &disable11k
	var disableRadiosWhenPowerConstrained bool = false
	this.DisableRadiosWhenPowerConstrained = &disableRadiosWhenPowerConstrained
	var enableArpSpoofCheck bool = false
	this.EnableArpSpoofCheck = &enableArpSpoofCheck
	var enableSharedRadioScanning bool = true
	this.EnableSharedRadioScanning = &enableSharedRadioScanning
	var enabled bool = true
	this.Enabled = &enabled
	var locateConnected bool = true
	this.LocateConnected = &locateConnected
	var locateUnconnected bool = false
	this.LocateUnconnected = &locateUnconnected
	var meshAllowDfs bool = false
	this.MeshAllowDfs = &meshAllowDfs
	var meshEnableCrm bool = false
	this.MeshEnableCrm = &meshEnableCrm
	var meshEnabled bool = false
	this.MeshEnabled = &meshEnabled
	return &this
}

// NewSiteWifiWithDefaults instantiates a new SiteWifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWifiWithDefaults() *SiteWifi {
	this := SiteWifi{}
	var ciscoEnabled bool = true
	this.CiscoEnabled = &ciscoEnabled
	var disable11k bool = false
	this.Disable11k = &disable11k
	var disableRadiosWhenPowerConstrained bool = false
	this.DisableRadiosWhenPowerConstrained = &disableRadiosWhenPowerConstrained
	var enableArpSpoofCheck bool = false
	this.EnableArpSpoofCheck = &enableArpSpoofCheck
	var enableSharedRadioScanning bool = true
	this.EnableSharedRadioScanning = &enableSharedRadioScanning
	var enabled bool = true
	this.Enabled = &enabled
	var locateConnected bool = true
	this.LocateConnected = &locateConnected
	var locateUnconnected bool = false
	this.LocateUnconnected = &locateUnconnected
	var meshAllowDfs bool = false
	this.MeshAllowDfs = &meshAllowDfs
	var meshEnableCrm bool = false
	this.MeshEnableCrm = &meshEnableCrm
	var meshEnabled bool = false
	this.MeshEnabled = &meshEnabled
	return &this
}

// GetCiscoEnabled returns the CiscoEnabled field value if set, zero value otherwise.
func (o *SiteWifi) GetCiscoEnabled() bool {
	if o == nil || IsNil(o.CiscoEnabled) {
		var ret bool
		return ret
	}
	return *o.CiscoEnabled
}

// GetCiscoEnabledOk returns a tuple with the CiscoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetCiscoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CiscoEnabled) {
		return nil, false
	}
	return o.CiscoEnabled, true
}

// HasCiscoEnabled returns a boolean if a field has been set.
func (o *SiteWifi) HasCiscoEnabled() bool {
	if o != nil && !IsNil(o.CiscoEnabled) {
		return true
	}

	return false
}

// SetCiscoEnabled gets a reference to the given bool and assigns it to the CiscoEnabled field.
func (o *SiteWifi) SetCiscoEnabled(v bool) {
	o.CiscoEnabled = &v
}

// GetDisable11k returns the Disable11k field value if set, zero value otherwise.
func (o *SiteWifi) GetDisable11k() bool {
	if o == nil || IsNil(o.Disable11k) {
		var ret bool
		return ret
	}
	return *o.Disable11k
}

// GetDisable11kOk returns a tuple with the Disable11k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetDisable11kOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable11k) {
		return nil, false
	}
	return o.Disable11k, true
}

// HasDisable11k returns a boolean if a field has been set.
func (o *SiteWifi) HasDisable11k() bool {
	if o != nil && !IsNil(o.Disable11k) {
		return true
	}

	return false
}

// SetDisable11k gets a reference to the given bool and assigns it to the Disable11k field.
func (o *SiteWifi) SetDisable11k(v bool) {
	o.Disable11k = &v
}

// GetDisableRadiosWhenPowerConstrained returns the DisableRadiosWhenPowerConstrained field value if set, zero value otherwise.
func (o *SiteWifi) GetDisableRadiosWhenPowerConstrained() bool {
	if o == nil || IsNil(o.DisableRadiosWhenPowerConstrained) {
		var ret bool
		return ret
	}
	return *o.DisableRadiosWhenPowerConstrained
}

// GetDisableRadiosWhenPowerConstrainedOk returns a tuple with the DisableRadiosWhenPowerConstrained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetDisableRadiosWhenPowerConstrainedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableRadiosWhenPowerConstrained) {
		return nil, false
	}
	return o.DisableRadiosWhenPowerConstrained, true
}

// HasDisableRadiosWhenPowerConstrained returns a boolean if a field has been set.
func (o *SiteWifi) HasDisableRadiosWhenPowerConstrained() bool {
	if o != nil && !IsNil(o.DisableRadiosWhenPowerConstrained) {
		return true
	}

	return false
}

// SetDisableRadiosWhenPowerConstrained gets a reference to the given bool and assigns it to the DisableRadiosWhenPowerConstrained field.
func (o *SiteWifi) SetDisableRadiosWhenPowerConstrained(v bool) {
	o.DisableRadiosWhenPowerConstrained = &v
}

// GetEnableArpSpoofCheck returns the EnableArpSpoofCheck field value if set, zero value otherwise.
func (o *SiteWifi) GetEnableArpSpoofCheck() bool {
	if o == nil || IsNil(o.EnableArpSpoofCheck) {
		var ret bool
		return ret
	}
	return *o.EnableArpSpoofCheck
}

// GetEnableArpSpoofCheckOk returns a tuple with the EnableArpSpoofCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetEnableArpSpoofCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableArpSpoofCheck) {
		return nil, false
	}
	return o.EnableArpSpoofCheck, true
}

// HasEnableArpSpoofCheck returns a boolean if a field has been set.
func (o *SiteWifi) HasEnableArpSpoofCheck() bool {
	if o != nil && !IsNil(o.EnableArpSpoofCheck) {
		return true
	}

	return false
}

// SetEnableArpSpoofCheck gets a reference to the given bool and assigns it to the EnableArpSpoofCheck field.
func (o *SiteWifi) SetEnableArpSpoofCheck(v bool) {
	o.EnableArpSpoofCheck = &v
}

// GetEnableSharedRadioScanning returns the EnableSharedRadioScanning field value if set, zero value otherwise.
func (o *SiteWifi) GetEnableSharedRadioScanning() bool {
	if o == nil || IsNil(o.EnableSharedRadioScanning) {
		var ret bool
		return ret
	}
	return *o.EnableSharedRadioScanning
}

// GetEnableSharedRadioScanningOk returns a tuple with the EnableSharedRadioScanning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetEnableSharedRadioScanningOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSharedRadioScanning) {
		return nil, false
	}
	return o.EnableSharedRadioScanning, true
}

// HasEnableSharedRadioScanning returns a boolean if a field has been set.
func (o *SiteWifi) HasEnableSharedRadioScanning() bool {
	if o != nil && !IsNil(o.EnableSharedRadioScanning) {
		return true
	}

	return false
}

// SetEnableSharedRadioScanning gets a reference to the given bool and assigns it to the EnableSharedRadioScanning field.
func (o *SiteWifi) SetEnableSharedRadioScanning(v bool) {
	o.EnableSharedRadioScanning = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SiteWifi) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SiteWifi) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SiteWifi) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLocateConnected returns the LocateConnected field value if set, zero value otherwise.
func (o *SiteWifi) GetLocateConnected() bool {
	if o == nil || IsNil(o.LocateConnected) {
		var ret bool
		return ret
	}
	return *o.LocateConnected
}

// GetLocateConnectedOk returns a tuple with the LocateConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetLocateConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.LocateConnected) {
		return nil, false
	}
	return o.LocateConnected, true
}

// HasLocateConnected returns a boolean if a field has been set.
func (o *SiteWifi) HasLocateConnected() bool {
	if o != nil && !IsNil(o.LocateConnected) {
		return true
	}

	return false
}

// SetLocateConnected gets a reference to the given bool and assigns it to the LocateConnected field.
func (o *SiteWifi) SetLocateConnected(v bool) {
	o.LocateConnected = &v
}

// GetLocateUnconnected returns the LocateUnconnected field value if set, zero value otherwise.
func (o *SiteWifi) GetLocateUnconnected() bool {
	if o == nil || IsNil(o.LocateUnconnected) {
		var ret bool
		return ret
	}
	return *o.LocateUnconnected
}

// GetLocateUnconnectedOk returns a tuple with the LocateUnconnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetLocateUnconnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.LocateUnconnected) {
		return nil, false
	}
	return o.LocateUnconnected, true
}

// HasLocateUnconnected returns a boolean if a field has been set.
func (o *SiteWifi) HasLocateUnconnected() bool {
	if o != nil && !IsNil(o.LocateUnconnected) {
		return true
	}

	return false
}

// SetLocateUnconnected gets a reference to the given bool and assigns it to the LocateUnconnected field.
func (o *SiteWifi) SetLocateUnconnected(v bool) {
	o.LocateUnconnected = &v
}

// GetMeshAllowDfs returns the MeshAllowDfs field value if set, zero value otherwise.
func (o *SiteWifi) GetMeshAllowDfs() bool {
	if o == nil || IsNil(o.MeshAllowDfs) {
		var ret bool
		return ret
	}
	return *o.MeshAllowDfs
}

// GetMeshAllowDfsOk returns a tuple with the MeshAllowDfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetMeshAllowDfsOk() (*bool, bool) {
	if o == nil || IsNil(o.MeshAllowDfs) {
		return nil, false
	}
	return o.MeshAllowDfs, true
}

// HasMeshAllowDfs returns a boolean if a field has been set.
func (o *SiteWifi) HasMeshAllowDfs() bool {
	if o != nil && !IsNil(o.MeshAllowDfs) {
		return true
	}

	return false
}

// SetMeshAllowDfs gets a reference to the given bool and assigns it to the MeshAllowDfs field.
func (o *SiteWifi) SetMeshAllowDfs(v bool) {
	o.MeshAllowDfs = &v
}

// GetMeshEnableCrm returns the MeshEnableCrm field value if set, zero value otherwise.
func (o *SiteWifi) GetMeshEnableCrm() bool {
	if o == nil || IsNil(o.MeshEnableCrm) {
		var ret bool
		return ret
	}
	return *o.MeshEnableCrm
}

// GetMeshEnableCrmOk returns a tuple with the MeshEnableCrm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetMeshEnableCrmOk() (*bool, bool) {
	if o == nil || IsNil(o.MeshEnableCrm) {
		return nil, false
	}
	return o.MeshEnableCrm, true
}

// HasMeshEnableCrm returns a boolean if a field has been set.
func (o *SiteWifi) HasMeshEnableCrm() bool {
	if o != nil && !IsNil(o.MeshEnableCrm) {
		return true
	}

	return false
}

// SetMeshEnableCrm gets a reference to the given bool and assigns it to the MeshEnableCrm field.
func (o *SiteWifi) SetMeshEnableCrm(v bool) {
	o.MeshEnableCrm = &v
}

// GetMeshEnabled returns the MeshEnabled field value if set, zero value otherwise.
func (o *SiteWifi) GetMeshEnabled() bool {
	if o == nil || IsNil(o.MeshEnabled) {
		var ret bool
		return ret
	}
	return *o.MeshEnabled
}

// GetMeshEnabledOk returns a tuple with the MeshEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWifi) GetMeshEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MeshEnabled) {
		return nil, false
	}
	return o.MeshEnabled, true
}

// HasMeshEnabled returns a boolean if a field has been set.
func (o *SiteWifi) HasMeshEnabled() bool {
	if o != nil && !IsNil(o.MeshEnabled) {
		return true
	}

	return false
}

// SetMeshEnabled gets a reference to the given bool and assigns it to the MeshEnabled field.
func (o *SiteWifi) SetMeshEnabled(v bool) {
	o.MeshEnabled = &v
}

// GetMeshPsk returns the MeshPsk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteWifi) GetMeshPsk() string {
	if o == nil || IsNil(o.MeshPsk.Get()) {
		var ret string
		return ret
	}
	return *o.MeshPsk.Get()
}

// GetMeshPskOk returns a tuple with the MeshPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteWifi) GetMeshPskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeshPsk.Get(), o.MeshPsk.IsSet()
}

// HasMeshPsk returns a boolean if a field has been set.
func (o *SiteWifi) HasMeshPsk() bool {
	if o != nil && o.MeshPsk.IsSet() {
		return true
	}

	return false
}

// SetMeshPsk gets a reference to the given NullableString and assigns it to the MeshPsk field.
func (o *SiteWifi) SetMeshPsk(v string) {
	o.MeshPsk.Set(&v)
}
// SetMeshPskNil sets the value for MeshPsk to be an explicit nil
func (o *SiteWifi) SetMeshPskNil() {
	o.MeshPsk.Set(nil)
}

// UnsetMeshPsk ensures that no value is present for MeshPsk, not even an explicit nil
func (o *SiteWifi) UnsetMeshPsk() {
	o.MeshPsk.Unset()
}

// GetMeshSsid returns the MeshSsid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteWifi) GetMeshSsid() string {
	if o == nil || IsNil(o.MeshSsid.Get()) {
		var ret string
		return ret
	}
	return *o.MeshSsid.Get()
}

// GetMeshSsidOk returns a tuple with the MeshSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteWifi) GetMeshSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeshSsid.Get(), o.MeshSsid.IsSet()
}

// HasMeshSsid returns a boolean if a field has been set.
func (o *SiteWifi) HasMeshSsid() bool {
	if o != nil && o.MeshSsid.IsSet() {
		return true
	}

	return false
}

// SetMeshSsid gets a reference to the given NullableString and assigns it to the MeshSsid field.
func (o *SiteWifi) SetMeshSsid(v string) {
	o.MeshSsid.Set(&v)
}
// SetMeshSsidNil sets the value for MeshSsid to be an explicit nil
func (o *SiteWifi) SetMeshSsidNil() {
	o.MeshSsid.Set(nil)
}

// UnsetMeshSsid ensures that no value is present for MeshSsid, not even an explicit nil
func (o *SiteWifi) UnsetMeshSsid() {
	o.MeshSsid.Unset()
}

// GetProxyArp returns the ProxyArp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteWifi) GetProxyArp() SiteWifiProxyArp {
	if o == nil || IsNil(o.ProxyArp.Get()) {
		var ret SiteWifiProxyArp
		return ret
	}
	return *o.ProxyArp.Get()
}

// GetProxyArpOk returns a tuple with the ProxyArp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteWifi) GetProxyArpOk() (*SiteWifiProxyArp, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyArp.Get(), o.ProxyArp.IsSet()
}

// HasProxyArp returns a boolean if a field has been set.
func (o *SiteWifi) HasProxyArp() bool {
	if o != nil && o.ProxyArp.IsSet() {
		return true
	}

	return false
}

// SetProxyArp gets a reference to the given NullableSiteWifiProxyArp and assigns it to the ProxyArp field.
func (o *SiteWifi) SetProxyArp(v SiteWifiProxyArp) {
	o.ProxyArp.Set(&v)
}
// SetProxyArpNil sets the value for ProxyArp to be an explicit nil
func (o *SiteWifi) SetProxyArpNil() {
	o.ProxyArp.Set(nil)
}

// UnsetProxyArp ensures that no value is present for ProxyArp, not even an explicit nil
func (o *SiteWifi) UnsetProxyArp() {
	o.ProxyArp.Unset()
}

func (o SiteWifi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteWifi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CiscoEnabled) {
		toSerialize["cisco_enabled"] = o.CiscoEnabled
	}
	if !IsNil(o.Disable11k) {
		toSerialize["disable_11k"] = o.Disable11k
	}
	if !IsNil(o.DisableRadiosWhenPowerConstrained) {
		toSerialize["disable_radios_when_power_constrained"] = o.DisableRadiosWhenPowerConstrained
	}
	if !IsNil(o.EnableArpSpoofCheck) {
		toSerialize["enable_arp_spoof_check"] = o.EnableArpSpoofCheck
	}
	if !IsNil(o.EnableSharedRadioScanning) {
		toSerialize["enable_shared_radio_scanning"] = o.EnableSharedRadioScanning
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.LocateConnected) {
		toSerialize["locate_connected"] = o.LocateConnected
	}
	if !IsNil(o.LocateUnconnected) {
		toSerialize["locate_unconnected"] = o.LocateUnconnected
	}
	if !IsNil(o.MeshAllowDfs) {
		toSerialize["mesh_allow_dfs"] = o.MeshAllowDfs
	}
	if !IsNil(o.MeshEnableCrm) {
		toSerialize["mesh_enable_crm"] = o.MeshEnableCrm
	}
	if !IsNil(o.MeshEnabled) {
		toSerialize["mesh_enabled"] = o.MeshEnabled
	}
	if o.MeshPsk.IsSet() {
		toSerialize["mesh_psk"] = o.MeshPsk.Get()
	}
	if o.MeshSsid.IsSet() {
		toSerialize["mesh_ssid"] = o.MeshSsid.Get()
	}
	if o.ProxyArp.IsSet() {
		toSerialize["proxy_arp"] = o.ProxyArp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteWifi) UnmarshalJSON(data []byte) (err error) {
	varSiteWifi := _SiteWifi{}

	err = json.Unmarshal(data, &varSiteWifi)

	if err != nil {
		return err
	}

	*o = SiteWifi(varSiteWifi)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cisco_enabled")
		delete(additionalProperties, "disable_11k")
		delete(additionalProperties, "disable_radios_when_power_constrained")
		delete(additionalProperties, "enable_arp_spoof_check")
		delete(additionalProperties, "enable_shared_radio_scanning")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "locate_connected")
		delete(additionalProperties, "locate_unconnected")
		delete(additionalProperties, "mesh_allow_dfs")
		delete(additionalProperties, "mesh_enable_crm")
		delete(additionalProperties, "mesh_enabled")
		delete(additionalProperties, "mesh_psk")
		delete(additionalProperties, "mesh_ssid")
		delete(additionalProperties, "proxy_arp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteWifi struct {
	value *SiteWifi
	isSet bool
}

func (v NullableSiteWifi) Get() *SiteWifi {
	return v.value
}

func (v *NullableSiteWifi) Set(val *SiteWifi) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteWifi(val *SiteWifi) *NullableSiteWifi {
	return &NullableSiteWifi{value: val, isSet: true}
}

func (v NullableSiteWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


