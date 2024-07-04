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

// checks if the ConstDeviceSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceSwitch{}

// ConstDeviceSwitch struct for ConstDeviceSwitch
type ConstDeviceSwitch struct {
	Alias *string `json:"alias,omitempty"`
	Defaults *ConstDeviceSwitchDefault `json:"defaults,omitempty"`
	Description *string `json:"description,omitempty"`
	Display *string `json:"display,omitempty"`
	EvolvedOs *bool `json:"evolved_os,omitempty"`
	EvpnRiType *string `json:"evpn_ri_type,omitempty"`
	Experimental *bool `json:"experimental,omitempty"`
	FansPluggable *bool `json:"fans_pluggable,omitempty"`
	HasBgp *bool `json:"has_bgp,omitempty"`
	HasEts *bool `json:"has_ets,omitempty"`
	HasEvpn *bool `json:"has_evpn,omitempty"`
	HasIrb *bool `json:"has_irb,omitempty"`
	HasPoeOut *bool `json:"has_poe_out,omitempty"`
	HasSnapshot *bool `json:"has_snapshot,omitempty"`
	HasVc *bool `json:"has_vc,omitempty"`
	Model *string `json:"model,omitempty"`
	Modular *bool `json:"modular,omitempty"`
	NoShapingRate *bool `json:"no_shaping_rate,omitempty"`
	NumberFans *int32 `json:"number_fans,omitempty"`
	OcDevice *bool `json:"oc_device,omitempty"`
	OobInterface *string `json:"oob_interface,omitempty"`
	PacketActionDropOnly *bool `json:"packet_action_drop_only,omitempty"`
	// Object Key is the PIC number
	Pic *map[string]string `json:"pic,omitempty"`
	SubRequired *string `json:"sub_required,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceSwitch ConstDeviceSwitch

// NewConstDeviceSwitch instantiates a new ConstDeviceSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceSwitch() *ConstDeviceSwitch {
	this := ConstDeviceSwitch{}
	var evolvedOs bool = false
	this.EvolvedOs = &evolvedOs
	var experimental bool = false
	this.Experimental = &experimental
	var fansPluggable bool = false
	this.FansPluggable = &fansPluggable
	var hasBgp bool = false
	this.HasBgp = &hasBgp
	var hasEts bool = false
	this.HasEts = &hasEts
	var hasEvpn bool = false
	this.HasEvpn = &hasEvpn
	var hasIrb bool = false
	this.HasIrb = &hasIrb
	var hasPoeOut bool = false
	this.HasPoeOut = &hasPoeOut
	var hasSnapshot bool = true
	this.HasSnapshot = &hasSnapshot
	var hasVc bool = true
	this.HasVc = &hasVc
	var modular bool = false
	this.Modular = &modular
	var noShapingRate bool = false
	this.NoShapingRate = &noShapingRate
	var ocDevice bool = false
	this.OcDevice = &ocDevice
	var packetActionDropOnly bool = false
	this.PacketActionDropOnly = &packetActionDropOnly
	var type_ string = "switch"
	this.Type = &type_
	return &this
}

// NewConstDeviceSwitchWithDefaults instantiates a new ConstDeviceSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceSwitchWithDefaults() *ConstDeviceSwitch {
	this := ConstDeviceSwitch{}
	var evolvedOs bool = false
	this.EvolvedOs = &evolvedOs
	var experimental bool = false
	this.Experimental = &experimental
	var fansPluggable bool = false
	this.FansPluggable = &fansPluggable
	var hasBgp bool = false
	this.HasBgp = &hasBgp
	var hasEts bool = false
	this.HasEts = &hasEts
	var hasEvpn bool = false
	this.HasEvpn = &hasEvpn
	var hasIrb bool = false
	this.HasIrb = &hasIrb
	var hasPoeOut bool = false
	this.HasPoeOut = &hasPoeOut
	var hasSnapshot bool = true
	this.HasSnapshot = &hasSnapshot
	var hasVc bool = true
	this.HasVc = &hasVc
	var modular bool = false
	this.Modular = &modular
	var noShapingRate bool = false
	this.NoShapingRate = &noShapingRate
	var ocDevice bool = false
	this.OcDevice = &ocDevice
	var packetActionDropOnly bool = false
	this.PacketActionDropOnly = &packetActionDropOnly
	var type_ string = "switch"
	this.Type = &type_
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ConstDeviceSwitch) SetAlias(v string) {
	o.Alias = &v
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetDefaults() ConstDeviceSwitchDefault {
	if o == nil || IsNil(o.Defaults) {
		var ret ConstDeviceSwitchDefault
		return ret
	}
	return *o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetDefaultsOk() (*ConstDeviceSwitchDefault, bool) {
	if o == nil || IsNil(o.Defaults) {
		return nil, false
	}
	return o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasDefaults() bool {
	if o != nil && !IsNil(o.Defaults) {
		return true
	}

	return false
}

// SetDefaults gets a reference to the given ConstDeviceSwitchDefault and assigns it to the Defaults field.
func (o *ConstDeviceSwitch) SetDefaults(v ConstDeviceSwitchDefault) {
	o.Defaults = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConstDeviceSwitch) SetDescription(v string) {
	o.Description = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ConstDeviceSwitch) SetDisplay(v string) {
	o.Display = &v
}

// GetEvolvedOs returns the EvolvedOs field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetEvolvedOs() bool {
	if o == nil || IsNil(o.EvolvedOs) {
		var ret bool
		return ret
	}
	return *o.EvolvedOs
}

// GetEvolvedOsOk returns a tuple with the EvolvedOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetEvolvedOsOk() (*bool, bool) {
	if o == nil || IsNil(o.EvolvedOs) {
		return nil, false
	}
	return o.EvolvedOs, true
}

// HasEvolvedOs returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasEvolvedOs() bool {
	if o != nil && !IsNil(o.EvolvedOs) {
		return true
	}

	return false
}

// SetEvolvedOs gets a reference to the given bool and assigns it to the EvolvedOs field.
func (o *ConstDeviceSwitch) SetEvolvedOs(v bool) {
	o.EvolvedOs = &v
}

// GetEvpnRiType returns the EvpnRiType field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetEvpnRiType() string {
	if o == nil || IsNil(o.EvpnRiType) {
		var ret string
		return ret
	}
	return *o.EvpnRiType
}

// GetEvpnRiTypeOk returns a tuple with the EvpnRiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetEvpnRiTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EvpnRiType) {
		return nil, false
	}
	return o.EvpnRiType, true
}

// HasEvpnRiType returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasEvpnRiType() bool {
	if o != nil && !IsNil(o.EvpnRiType) {
		return true
	}

	return false
}

// SetEvpnRiType gets a reference to the given string and assigns it to the EvpnRiType field.
func (o *ConstDeviceSwitch) SetEvpnRiType(v string) {
	o.EvpnRiType = &v
}

// GetExperimental returns the Experimental field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetExperimental() bool {
	if o == nil || IsNil(o.Experimental) {
		var ret bool
		return ret
	}
	return *o.Experimental
}

// GetExperimentalOk returns a tuple with the Experimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetExperimentalOk() (*bool, bool) {
	if o == nil || IsNil(o.Experimental) {
		return nil, false
	}
	return o.Experimental, true
}

// HasExperimental returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasExperimental() bool {
	if o != nil && !IsNil(o.Experimental) {
		return true
	}

	return false
}

// SetExperimental gets a reference to the given bool and assigns it to the Experimental field.
func (o *ConstDeviceSwitch) SetExperimental(v bool) {
	o.Experimental = &v
}

// GetFansPluggable returns the FansPluggable field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetFansPluggable() bool {
	if o == nil || IsNil(o.FansPluggable) {
		var ret bool
		return ret
	}
	return *o.FansPluggable
}

// GetFansPluggableOk returns a tuple with the FansPluggable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetFansPluggableOk() (*bool, bool) {
	if o == nil || IsNil(o.FansPluggable) {
		return nil, false
	}
	return o.FansPluggable, true
}

// HasFansPluggable returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasFansPluggable() bool {
	if o != nil && !IsNil(o.FansPluggable) {
		return true
	}

	return false
}

// SetFansPluggable gets a reference to the given bool and assigns it to the FansPluggable field.
func (o *ConstDeviceSwitch) SetFansPluggable(v bool) {
	o.FansPluggable = &v
}

// GetHasBgp returns the HasBgp field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasBgp() bool {
	if o == nil || IsNil(o.HasBgp) {
		var ret bool
		return ret
	}
	return *o.HasBgp
}

// GetHasBgpOk returns a tuple with the HasBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasBgpOk() (*bool, bool) {
	if o == nil || IsNil(o.HasBgp) {
		return nil, false
	}
	return o.HasBgp, true
}

// HasHasBgp returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasBgp() bool {
	if o != nil && !IsNil(o.HasBgp) {
		return true
	}

	return false
}

// SetHasBgp gets a reference to the given bool and assigns it to the HasBgp field.
func (o *ConstDeviceSwitch) SetHasBgp(v bool) {
	o.HasBgp = &v
}

// GetHasEts returns the HasEts field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasEts() bool {
	if o == nil || IsNil(o.HasEts) {
		var ret bool
		return ret
	}
	return *o.HasEts
}

// GetHasEtsOk returns a tuple with the HasEts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasEtsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEts) {
		return nil, false
	}
	return o.HasEts, true
}

// HasHasEts returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasEts() bool {
	if o != nil && !IsNil(o.HasEts) {
		return true
	}

	return false
}

// SetHasEts gets a reference to the given bool and assigns it to the HasEts field.
func (o *ConstDeviceSwitch) SetHasEts(v bool) {
	o.HasEts = &v
}

// GetHasEvpn returns the HasEvpn field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasEvpn() bool {
	if o == nil || IsNil(o.HasEvpn) {
		var ret bool
		return ret
	}
	return *o.HasEvpn
}

// GetHasEvpnOk returns a tuple with the HasEvpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasEvpnOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEvpn) {
		return nil, false
	}
	return o.HasEvpn, true
}

// HasHasEvpn returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasEvpn() bool {
	if o != nil && !IsNil(o.HasEvpn) {
		return true
	}

	return false
}

// SetHasEvpn gets a reference to the given bool and assigns it to the HasEvpn field.
func (o *ConstDeviceSwitch) SetHasEvpn(v bool) {
	o.HasEvpn = &v
}

// GetHasIrb returns the HasIrb field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasIrb() bool {
	if o == nil || IsNil(o.HasIrb) {
		var ret bool
		return ret
	}
	return *o.HasIrb
}

// GetHasIrbOk returns a tuple with the HasIrb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasIrbOk() (*bool, bool) {
	if o == nil || IsNil(o.HasIrb) {
		return nil, false
	}
	return o.HasIrb, true
}

// HasHasIrb returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasIrb() bool {
	if o != nil && !IsNil(o.HasIrb) {
		return true
	}

	return false
}

// SetHasIrb gets a reference to the given bool and assigns it to the HasIrb field.
func (o *ConstDeviceSwitch) SetHasIrb(v bool) {
	o.HasIrb = &v
}

// GetHasPoeOut returns the HasPoeOut field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasPoeOut() bool {
	if o == nil || IsNil(o.HasPoeOut) {
		var ret bool
		return ret
	}
	return *o.HasPoeOut
}

// GetHasPoeOutOk returns a tuple with the HasPoeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasPoeOutOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPoeOut) {
		return nil, false
	}
	return o.HasPoeOut, true
}

// HasHasPoeOut returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasPoeOut() bool {
	if o != nil && !IsNil(o.HasPoeOut) {
		return true
	}

	return false
}

// SetHasPoeOut gets a reference to the given bool and assigns it to the HasPoeOut field.
func (o *ConstDeviceSwitch) SetHasPoeOut(v bool) {
	o.HasPoeOut = &v
}

// GetHasSnapshot returns the HasSnapshot field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasSnapshot() bool {
	if o == nil || IsNil(o.HasSnapshot) {
		var ret bool
		return ret
	}
	return *o.HasSnapshot
}

// GetHasSnapshotOk returns a tuple with the HasSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSnapshot) {
		return nil, false
	}
	return o.HasSnapshot, true
}

// HasHasSnapshot returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasSnapshot() bool {
	if o != nil && !IsNil(o.HasSnapshot) {
		return true
	}

	return false
}

// SetHasSnapshot gets a reference to the given bool and assigns it to the HasSnapshot field.
func (o *ConstDeviceSwitch) SetHasSnapshot(v bool) {
	o.HasSnapshot = &v
}

// GetHasVc returns the HasVc field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetHasVc() bool {
	if o == nil || IsNil(o.HasVc) {
		var ret bool
		return ret
	}
	return *o.HasVc
}

// GetHasVcOk returns a tuple with the HasVc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetHasVcOk() (*bool, bool) {
	if o == nil || IsNil(o.HasVc) {
		return nil, false
	}
	return o.HasVc, true
}

// HasHasVc returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasHasVc() bool {
	if o != nil && !IsNil(o.HasVc) {
		return true
	}

	return false
}

// SetHasVc gets a reference to the given bool and assigns it to the HasVc field.
func (o *ConstDeviceSwitch) SetHasVc(v bool) {
	o.HasVc = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConstDeviceSwitch) SetModel(v string) {
	o.Model = &v
}

// GetModular returns the Modular field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetModular() bool {
	if o == nil || IsNil(o.Modular) {
		var ret bool
		return ret
	}
	return *o.Modular
}

// GetModularOk returns a tuple with the Modular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetModularOk() (*bool, bool) {
	if o == nil || IsNil(o.Modular) {
		return nil, false
	}
	return o.Modular, true
}

// HasModular returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasModular() bool {
	if o != nil && !IsNil(o.Modular) {
		return true
	}

	return false
}

// SetModular gets a reference to the given bool and assigns it to the Modular field.
func (o *ConstDeviceSwitch) SetModular(v bool) {
	o.Modular = &v
}

// GetNoShapingRate returns the NoShapingRate field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetNoShapingRate() bool {
	if o == nil || IsNil(o.NoShapingRate) {
		var ret bool
		return ret
	}
	return *o.NoShapingRate
}

// GetNoShapingRateOk returns a tuple with the NoShapingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetNoShapingRateOk() (*bool, bool) {
	if o == nil || IsNil(o.NoShapingRate) {
		return nil, false
	}
	return o.NoShapingRate, true
}

// HasNoShapingRate returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasNoShapingRate() bool {
	if o != nil && !IsNil(o.NoShapingRate) {
		return true
	}

	return false
}

// SetNoShapingRate gets a reference to the given bool and assigns it to the NoShapingRate field.
func (o *ConstDeviceSwitch) SetNoShapingRate(v bool) {
	o.NoShapingRate = &v
}

// GetNumberFans returns the NumberFans field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetNumberFans() int32 {
	if o == nil || IsNil(o.NumberFans) {
		var ret int32
		return ret
	}
	return *o.NumberFans
}

// GetNumberFansOk returns a tuple with the NumberFans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetNumberFansOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberFans) {
		return nil, false
	}
	return o.NumberFans, true
}

// HasNumberFans returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasNumberFans() bool {
	if o != nil && !IsNil(o.NumberFans) {
		return true
	}

	return false
}

// SetNumberFans gets a reference to the given int32 and assigns it to the NumberFans field.
func (o *ConstDeviceSwitch) SetNumberFans(v int32) {
	o.NumberFans = &v
}

// GetOcDevice returns the OcDevice field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetOcDevice() bool {
	if o == nil || IsNil(o.OcDevice) {
		var ret bool
		return ret
	}
	return *o.OcDevice
}

// GetOcDeviceOk returns a tuple with the OcDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetOcDeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.OcDevice) {
		return nil, false
	}
	return o.OcDevice, true
}

// HasOcDevice returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasOcDevice() bool {
	if o != nil && !IsNil(o.OcDevice) {
		return true
	}

	return false
}

// SetOcDevice gets a reference to the given bool and assigns it to the OcDevice field.
func (o *ConstDeviceSwitch) SetOcDevice(v bool) {
	o.OcDevice = &v
}

// GetOobInterface returns the OobInterface field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetOobInterface() string {
	if o == nil || IsNil(o.OobInterface) {
		var ret string
		return ret
	}
	return *o.OobInterface
}

// GetOobInterfaceOk returns a tuple with the OobInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetOobInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.OobInterface) {
		return nil, false
	}
	return o.OobInterface, true
}

// HasOobInterface returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasOobInterface() bool {
	if o != nil && !IsNil(o.OobInterface) {
		return true
	}

	return false
}

// SetOobInterface gets a reference to the given string and assigns it to the OobInterface field.
func (o *ConstDeviceSwitch) SetOobInterface(v string) {
	o.OobInterface = &v
}

// GetPacketActionDropOnly returns the PacketActionDropOnly field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetPacketActionDropOnly() bool {
	if o == nil || IsNil(o.PacketActionDropOnly) {
		var ret bool
		return ret
	}
	return *o.PacketActionDropOnly
}

// GetPacketActionDropOnlyOk returns a tuple with the PacketActionDropOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetPacketActionDropOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PacketActionDropOnly) {
		return nil, false
	}
	return o.PacketActionDropOnly, true
}

// HasPacketActionDropOnly returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasPacketActionDropOnly() bool {
	if o != nil && !IsNil(o.PacketActionDropOnly) {
		return true
	}

	return false
}

// SetPacketActionDropOnly gets a reference to the given bool and assigns it to the PacketActionDropOnly field.
func (o *ConstDeviceSwitch) SetPacketActionDropOnly(v bool) {
	o.PacketActionDropOnly = &v
}

// GetPic returns the Pic field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetPic() map[string]string {
	if o == nil || IsNil(o.Pic) {
		var ret map[string]string
		return ret
	}
	return *o.Pic
}

// GetPicOk returns a tuple with the Pic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetPicOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Pic) {
		return nil, false
	}
	return o.Pic, true
}

// HasPic returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasPic() bool {
	if o != nil && !IsNil(o.Pic) {
		return true
	}

	return false
}

// SetPic gets a reference to the given map[string]string and assigns it to the Pic field.
func (o *ConstDeviceSwitch) SetPic(v map[string]string) {
	o.Pic = &v
}

// GetSubRequired returns the SubRequired field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetSubRequired() string {
	if o == nil || IsNil(o.SubRequired) {
		var ret string
		return ret
	}
	return *o.SubRequired
}

// GetSubRequiredOk returns a tuple with the SubRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetSubRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.SubRequired) {
		return nil, false
	}
	return o.SubRequired, true
}

// HasSubRequired returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasSubRequired() bool {
	if o != nil && !IsNil(o.SubRequired) {
		return true
	}

	return false
}

// SetSubRequired gets a reference to the given string and assigns it to the SubRequired field.
func (o *ConstDeviceSwitch) SetSubRequired(v string) {
	o.SubRequired = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConstDeviceSwitch) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceSwitch) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConstDeviceSwitch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConstDeviceSwitch) SetType(v string) {
	o.Type = &v
}

func (o ConstDeviceSwitch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Defaults) {
		toSerialize["defaults"] = o.Defaults
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.EvolvedOs) {
		toSerialize["evolved_os"] = o.EvolvedOs
	}
	if !IsNil(o.EvpnRiType) {
		toSerialize["evpn_ri_type"] = o.EvpnRiType
	}
	if !IsNil(o.Experimental) {
		toSerialize["experimental"] = o.Experimental
	}
	if !IsNil(o.FansPluggable) {
		toSerialize["fans_pluggable"] = o.FansPluggable
	}
	if !IsNil(o.HasBgp) {
		toSerialize["has_bgp"] = o.HasBgp
	}
	if !IsNil(o.HasEts) {
		toSerialize["has_ets"] = o.HasEts
	}
	if !IsNil(o.HasEvpn) {
		toSerialize["has_evpn"] = o.HasEvpn
	}
	if !IsNil(o.HasIrb) {
		toSerialize["has_irb"] = o.HasIrb
	}
	if !IsNil(o.HasPoeOut) {
		toSerialize["has_poe_out"] = o.HasPoeOut
	}
	if !IsNil(o.HasSnapshot) {
		toSerialize["has_snapshot"] = o.HasSnapshot
	}
	if !IsNil(o.HasVc) {
		toSerialize["has_vc"] = o.HasVc
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Modular) {
		toSerialize["modular"] = o.Modular
	}
	if !IsNil(o.NoShapingRate) {
		toSerialize["no_shaping_rate"] = o.NoShapingRate
	}
	if !IsNil(o.NumberFans) {
		toSerialize["number_fans"] = o.NumberFans
	}
	if !IsNil(o.OcDevice) {
		toSerialize["oc_device"] = o.OcDevice
	}
	if !IsNil(o.OobInterface) {
		toSerialize["oob_interface"] = o.OobInterface
	}
	if !IsNil(o.PacketActionDropOnly) {
		toSerialize["packet_action_drop_only"] = o.PacketActionDropOnly
	}
	if !IsNil(o.Pic) {
		toSerialize["pic"] = o.Pic
	}
	if !IsNil(o.SubRequired) {
		toSerialize["sub_required"] = o.SubRequired
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceSwitch) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceSwitch := _ConstDeviceSwitch{}

	err = json.Unmarshal(data, &varConstDeviceSwitch)

	if err != nil {
		return err
	}

	*o = ConstDeviceSwitch(varConstDeviceSwitch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alias")
		delete(additionalProperties, "defaults")
		delete(additionalProperties, "description")
		delete(additionalProperties, "display")
		delete(additionalProperties, "evolved_os")
		delete(additionalProperties, "evpn_ri_type")
		delete(additionalProperties, "experimental")
		delete(additionalProperties, "fans_pluggable")
		delete(additionalProperties, "has_bgp")
		delete(additionalProperties, "has_ets")
		delete(additionalProperties, "has_evpn")
		delete(additionalProperties, "has_irb")
		delete(additionalProperties, "has_poe_out")
		delete(additionalProperties, "has_snapshot")
		delete(additionalProperties, "has_vc")
		delete(additionalProperties, "model")
		delete(additionalProperties, "modular")
		delete(additionalProperties, "no_shaping_rate")
		delete(additionalProperties, "number_fans")
		delete(additionalProperties, "oc_device")
		delete(additionalProperties, "oob_interface")
		delete(additionalProperties, "packet_action_drop_only")
		delete(additionalProperties, "pic")
		delete(additionalProperties, "sub_required")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceSwitch struct {
	value *ConstDeviceSwitch
	isSet bool
}

func (v NullableConstDeviceSwitch) Get() *ConstDeviceSwitch {
	return v.value
}

func (v *NullableConstDeviceSwitch) Set(val *ConstDeviceSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceSwitch(val *ConstDeviceSwitch) *NullableConstDeviceSwitch {
	return &NullableConstDeviceSwitch{value: val, isSet: true}
}

func (v NullableConstDeviceSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


