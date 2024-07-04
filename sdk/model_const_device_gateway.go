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

// checks if the ConstDeviceGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceGateway{}

// ConstDeviceGateway struct for ConstDeviceGateway
type ConstDeviceGateway struct {
	// Object Key is the interface type name (e.g. \"lan_ports\", \"wan_ports\", ...)
	Defaults *map[string]string `json:"defaults,omitempty"`
	Description *string `json:"description,omitempty"`
	Experimental *bool `json:"experimental,omitempty"`
	FansPluggable *bool `json:"fans_pluggable,omitempty"`
	HaNode0Fpc *int32 `json:"ha_node0_fpc,omitempty"`
	HaNode1Fpc *int32 `json:"ha_node1_fpc,omitempty"`
	HasBgp *bool `json:"has_bgp,omitempty"`
	HasFxp0 *bool `json:"has_fxp0,omitempty"`
	HasHaControl *bool `json:"has_ha_control,omitempty"`
	HasHaData *bool `json:"has_ha_data,omitempty"`
	HasIrb *bool `json:"has_irb,omitempty"`
	HasPoeOut *bool `json:"has_poe_out,omitempty"`
	HasSnapshot *bool `json:"has_snapshot,omitempty"`
	IrbDisabledByDefault *bool `json:"irb_disabled_by_default,omitempty"`
	Model *string `json:"model,omitempty"`
	NumberFans *int32 `json:"number_fans,omitempty"`
	OcDevice *bool `json:"oc_device,omitempty"`
	// Object Key is the PIC number
	Pic *map[string]string `json:"pic,omitempty"`
	Ports *ConstDeviceGatewayPorts `json:"ports,omitempty"`
	SubRequired *string `json:"sub_required,omitempty"`
	T128Device *bool `json:"t128_device,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceGateway ConstDeviceGateway

// NewConstDeviceGateway instantiates a new ConstDeviceGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceGateway() *ConstDeviceGateway {
	this := ConstDeviceGateway{}
	var experimental bool = false
	this.Experimental = &experimental
	var fansPluggable bool = true
	this.FansPluggable = &fansPluggable
	var hasBgp bool = false
	this.HasBgp = &hasBgp
	var hasFxp0 bool = true
	this.HasFxp0 = &hasFxp0
	var hasHaControl bool = false
	this.HasHaControl = &hasHaControl
	var hasHaData bool = false
	this.HasHaData = &hasHaData
	var hasIrb bool = false
	this.HasIrb = &hasIrb
	var hasPoeOut bool = true
	this.HasPoeOut = &hasPoeOut
	var hasSnapshot bool = true
	this.HasSnapshot = &hasSnapshot
	var irbDisabledByDefault bool = false
	this.IrbDisabledByDefault = &irbDisabledByDefault
	var ocDevice bool = false
	this.OcDevice = &ocDevice
	var t128Device bool = false
	this.T128Device = &t128Device
	var type_ string = "gateway"
	this.Type = &type_
	return &this
}

// NewConstDeviceGatewayWithDefaults instantiates a new ConstDeviceGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceGatewayWithDefaults() *ConstDeviceGateway {
	this := ConstDeviceGateway{}
	var experimental bool = false
	this.Experimental = &experimental
	var fansPluggable bool = true
	this.FansPluggable = &fansPluggable
	var hasBgp bool = false
	this.HasBgp = &hasBgp
	var hasFxp0 bool = true
	this.HasFxp0 = &hasFxp0
	var hasHaControl bool = false
	this.HasHaControl = &hasHaControl
	var hasHaData bool = false
	this.HasHaData = &hasHaData
	var hasIrb bool = false
	this.HasIrb = &hasIrb
	var hasPoeOut bool = true
	this.HasPoeOut = &hasPoeOut
	var hasSnapshot bool = true
	this.HasSnapshot = &hasSnapshot
	var irbDisabledByDefault bool = false
	this.IrbDisabledByDefault = &irbDisabledByDefault
	var ocDevice bool = false
	this.OcDevice = &ocDevice
	var t128Device bool = false
	this.T128Device = &t128Device
	var type_ string = "gateway"
	this.Type = &type_
	return &this
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetDefaults() map[string]string {
	if o == nil || IsNil(o.Defaults) {
		var ret map[string]string
		return ret
	}
	return *o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetDefaultsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Defaults) {
		return nil, false
	}
	return o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasDefaults() bool {
	if o != nil && !IsNil(o.Defaults) {
		return true
	}

	return false
}

// SetDefaults gets a reference to the given map[string]string and assigns it to the Defaults field.
func (o *ConstDeviceGateway) SetDefaults(v map[string]string) {
	o.Defaults = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConstDeviceGateway) SetDescription(v string) {
	o.Description = &v
}

// GetExperimental returns the Experimental field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetExperimental() bool {
	if o == nil || IsNil(o.Experimental) {
		var ret bool
		return ret
	}
	return *o.Experimental
}

// GetExperimentalOk returns a tuple with the Experimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetExperimentalOk() (*bool, bool) {
	if o == nil || IsNil(o.Experimental) {
		return nil, false
	}
	return o.Experimental, true
}

// HasExperimental returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasExperimental() bool {
	if o != nil && !IsNil(o.Experimental) {
		return true
	}

	return false
}

// SetExperimental gets a reference to the given bool and assigns it to the Experimental field.
func (o *ConstDeviceGateway) SetExperimental(v bool) {
	o.Experimental = &v
}

// GetFansPluggable returns the FansPluggable field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetFansPluggable() bool {
	if o == nil || IsNil(o.FansPluggable) {
		var ret bool
		return ret
	}
	return *o.FansPluggable
}

// GetFansPluggableOk returns a tuple with the FansPluggable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetFansPluggableOk() (*bool, bool) {
	if o == nil || IsNil(o.FansPluggable) {
		return nil, false
	}
	return o.FansPluggable, true
}

// HasFansPluggable returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasFansPluggable() bool {
	if o != nil && !IsNil(o.FansPluggable) {
		return true
	}

	return false
}

// SetFansPluggable gets a reference to the given bool and assigns it to the FansPluggable field.
func (o *ConstDeviceGateway) SetFansPluggable(v bool) {
	o.FansPluggable = &v
}

// GetHaNode0Fpc returns the HaNode0Fpc field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHaNode0Fpc() int32 {
	if o == nil || IsNil(o.HaNode0Fpc) {
		var ret int32
		return ret
	}
	return *o.HaNode0Fpc
}

// GetHaNode0FpcOk returns a tuple with the HaNode0Fpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHaNode0FpcOk() (*int32, bool) {
	if o == nil || IsNil(o.HaNode0Fpc) {
		return nil, false
	}
	return o.HaNode0Fpc, true
}

// HasHaNode0Fpc returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHaNode0Fpc() bool {
	if o != nil && !IsNil(o.HaNode0Fpc) {
		return true
	}

	return false
}

// SetHaNode0Fpc gets a reference to the given int32 and assigns it to the HaNode0Fpc field.
func (o *ConstDeviceGateway) SetHaNode0Fpc(v int32) {
	o.HaNode0Fpc = &v
}

// GetHaNode1Fpc returns the HaNode1Fpc field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHaNode1Fpc() int32 {
	if o == nil || IsNil(o.HaNode1Fpc) {
		var ret int32
		return ret
	}
	return *o.HaNode1Fpc
}

// GetHaNode1FpcOk returns a tuple with the HaNode1Fpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHaNode1FpcOk() (*int32, bool) {
	if o == nil || IsNil(o.HaNode1Fpc) {
		return nil, false
	}
	return o.HaNode1Fpc, true
}

// HasHaNode1Fpc returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHaNode1Fpc() bool {
	if o != nil && !IsNil(o.HaNode1Fpc) {
		return true
	}

	return false
}

// SetHaNode1Fpc gets a reference to the given int32 and assigns it to the HaNode1Fpc field.
func (o *ConstDeviceGateway) SetHaNode1Fpc(v int32) {
	o.HaNode1Fpc = &v
}

// GetHasBgp returns the HasBgp field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasBgp() bool {
	if o == nil || IsNil(o.HasBgp) {
		var ret bool
		return ret
	}
	return *o.HasBgp
}

// GetHasBgpOk returns a tuple with the HasBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasBgpOk() (*bool, bool) {
	if o == nil || IsNil(o.HasBgp) {
		return nil, false
	}
	return o.HasBgp, true
}

// HasHasBgp returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasBgp() bool {
	if o != nil && !IsNil(o.HasBgp) {
		return true
	}

	return false
}

// SetHasBgp gets a reference to the given bool and assigns it to the HasBgp field.
func (o *ConstDeviceGateway) SetHasBgp(v bool) {
	o.HasBgp = &v
}

// GetHasFxp0 returns the HasFxp0 field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasFxp0() bool {
	if o == nil || IsNil(o.HasFxp0) {
		var ret bool
		return ret
	}
	return *o.HasFxp0
}

// GetHasFxp0Ok returns a tuple with the HasFxp0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasFxp0Ok() (*bool, bool) {
	if o == nil || IsNil(o.HasFxp0) {
		return nil, false
	}
	return o.HasFxp0, true
}

// HasHasFxp0 returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasFxp0() bool {
	if o != nil && !IsNil(o.HasFxp0) {
		return true
	}

	return false
}

// SetHasFxp0 gets a reference to the given bool and assigns it to the HasFxp0 field.
func (o *ConstDeviceGateway) SetHasFxp0(v bool) {
	o.HasFxp0 = &v
}

// GetHasHaControl returns the HasHaControl field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasHaControl() bool {
	if o == nil || IsNil(o.HasHaControl) {
		var ret bool
		return ret
	}
	return *o.HasHaControl
}

// GetHasHaControlOk returns a tuple with the HasHaControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasHaControlOk() (*bool, bool) {
	if o == nil || IsNil(o.HasHaControl) {
		return nil, false
	}
	return o.HasHaControl, true
}

// HasHasHaControl returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasHaControl() bool {
	if o != nil && !IsNil(o.HasHaControl) {
		return true
	}

	return false
}

// SetHasHaControl gets a reference to the given bool and assigns it to the HasHaControl field.
func (o *ConstDeviceGateway) SetHasHaControl(v bool) {
	o.HasHaControl = &v
}

// GetHasHaData returns the HasHaData field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasHaData() bool {
	if o == nil || IsNil(o.HasHaData) {
		var ret bool
		return ret
	}
	return *o.HasHaData
}

// GetHasHaDataOk returns a tuple with the HasHaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasHaDataOk() (*bool, bool) {
	if o == nil || IsNil(o.HasHaData) {
		return nil, false
	}
	return o.HasHaData, true
}

// HasHasHaData returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasHaData() bool {
	if o != nil && !IsNil(o.HasHaData) {
		return true
	}

	return false
}

// SetHasHaData gets a reference to the given bool and assigns it to the HasHaData field.
func (o *ConstDeviceGateway) SetHasHaData(v bool) {
	o.HasHaData = &v
}

// GetHasIrb returns the HasIrb field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasIrb() bool {
	if o == nil || IsNil(o.HasIrb) {
		var ret bool
		return ret
	}
	return *o.HasIrb
}

// GetHasIrbOk returns a tuple with the HasIrb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasIrbOk() (*bool, bool) {
	if o == nil || IsNil(o.HasIrb) {
		return nil, false
	}
	return o.HasIrb, true
}

// HasHasIrb returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasIrb() bool {
	if o != nil && !IsNil(o.HasIrb) {
		return true
	}

	return false
}

// SetHasIrb gets a reference to the given bool and assigns it to the HasIrb field.
func (o *ConstDeviceGateway) SetHasIrb(v bool) {
	o.HasIrb = &v
}

// GetHasPoeOut returns the HasPoeOut field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasPoeOut() bool {
	if o == nil || IsNil(o.HasPoeOut) {
		var ret bool
		return ret
	}
	return *o.HasPoeOut
}

// GetHasPoeOutOk returns a tuple with the HasPoeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasPoeOutOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPoeOut) {
		return nil, false
	}
	return o.HasPoeOut, true
}

// HasHasPoeOut returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasPoeOut() bool {
	if o != nil && !IsNil(o.HasPoeOut) {
		return true
	}

	return false
}

// SetHasPoeOut gets a reference to the given bool and assigns it to the HasPoeOut field.
func (o *ConstDeviceGateway) SetHasPoeOut(v bool) {
	o.HasPoeOut = &v
}

// GetHasSnapshot returns the HasSnapshot field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetHasSnapshot() bool {
	if o == nil || IsNil(o.HasSnapshot) {
		var ret bool
		return ret
	}
	return *o.HasSnapshot
}

// GetHasSnapshotOk returns a tuple with the HasSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetHasSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSnapshot) {
		return nil, false
	}
	return o.HasSnapshot, true
}

// HasHasSnapshot returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasHasSnapshot() bool {
	if o != nil && !IsNil(o.HasSnapshot) {
		return true
	}

	return false
}

// SetHasSnapshot gets a reference to the given bool and assigns it to the HasSnapshot field.
func (o *ConstDeviceGateway) SetHasSnapshot(v bool) {
	o.HasSnapshot = &v
}

// GetIrbDisabledByDefault returns the IrbDisabledByDefault field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetIrbDisabledByDefault() bool {
	if o == nil || IsNil(o.IrbDisabledByDefault) {
		var ret bool
		return ret
	}
	return *o.IrbDisabledByDefault
}

// GetIrbDisabledByDefaultOk returns a tuple with the IrbDisabledByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetIrbDisabledByDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IrbDisabledByDefault) {
		return nil, false
	}
	return o.IrbDisabledByDefault, true
}

// HasIrbDisabledByDefault returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasIrbDisabledByDefault() bool {
	if o != nil && !IsNil(o.IrbDisabledByDefault) {
		return true
	}

	return false
}

// SetIrbDisabledByDefault gets a reference to the given bool and assigns it to the IrbDisabledByDefault field.
func (o *ConstDeviceGateway) SetIrbDisabledByDefault(v bool) {
	o.IrbDisabledByDefault = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConstDeviceGateway) SetModel(v string) {
	o.Model = &v
}

// GetNumberFans returns the NumberFans field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetNumberFans() int32 {
	if o == nil || IsNil(o.NumberFans) {
		var ret int32
		return ret
	}
	return *o.NumberFans
}

// GetNumberFansOk returns a tuple with the NumberFans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetNumberFansOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberFans) {
		return nil, false
	}
	return o.NumberFans, true
}

// HasNumberFans returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasNumberFans() bool {
	if o != nil && !IsNil(o.NumberFans) {
		return true
	}

	return false
}

// SetNumberFans gets a reference to the given int32 and assigns it to the NumberFans field.
func (o *ConstDeviceGateway) SetNumberFans(v int32) {
	o.NumberFans = &v
}

// GetOcDevice returns the OcDevice field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetOcDevice() bool {
	if o == nil || IsNil(o.OcDevice) {
		var ret bool
		return ret
	}
	return *o.OcDevice
}

// GetOcDeviceOk returns a tuple with the OcDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetOcDeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.OcDevice) {
		return nil, false
	}
	return o.OcDevice, true
}

// HasOcDevice returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasOcDevice() bool {
	if o != nil && !IsNil(o.OcDevice) {
		return true
	}

	return false
}

// SetOcDevice gets a reference to the given bool and assigns it to the OcDevice field.
func (o *ConstDeviceGateway) SetOcDevice(v bool) {
	o.OcDevice = &v
}

// GetPic returns the Pic field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetPic() map[string]string {
	if o == nil || IsNil(o.Pic) {
		var ret map[string]string
		return ret
	}
	return *o.Pic
}

// GetPicOk returns a tuple with the Pic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetPicOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Pic) {
		return nil, false
	}
	return o.Pic, true
}

// HasPic returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasPic() bool {
	if o != nil && !IsNil(o.Pic) {
		return true
	}

	return false
}

// SetPic gets a reference to the given map[string]string and assigns it to the Pic field.
func (o *ConstDeviceGateway) SetPic(v map[string]string) {
	o.Pic = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetPorts() ConstDeviceGatewayPorts {
	if o == nil || IsNil(o.Ports) {
		var ret ConstDeviceGatewayPorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetPortsOk() (*ConstDeviceGatewayPorts, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given ConstDeviceGatewayPorts and assigns it to the Ports field.
func (o *ConstDeviceGateway) SetPorts(v ConstDeviceGatewayPorts) {
	o.Ports = &v
}

// GetSubRequired returns the SubRequired field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetSubRequired() string {
	if o == nil || IsNil(o.SubRequired) {
		var ret string
		return ret
	}
	return *o.SubRequired
}

// GetSubRequiredOk returns a tuple with the SubRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetSubRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.SubRequired) {
		return nil, false
	}
	return o.SubRequired, true
}

// HasSubRequired returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasSubRequired() bool {
	if o != nil && !IsNil(o.SubRequired) {
		return true
	}

	return false
}

// SetSubRequired gets a reference to the given string and assigns it to the SubRequired field.
func (o *ConstDeviceGateway) SetSubRequired(v string) {
	o.SubRequired = &v
}

// GetT128Device returns the T128Device field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetT128Device() bool {
	if o == nil || IsNil(o.T128Device) {
		var ret bool
		return ret
	}
	return *o.T128Device
}

// GetT128DeviceOk returns a tuple with the T128Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetT128DeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.T128Device) {
		return nil, false
	}
	return o.T128Device, true
}

// HasT128Device returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasT128Device() bool {
	if o != nil && !IsNil(o.T128Device) {
		return true
	}

	return false
}

// SetT128Device gets a reference to the given bool and assigns it to the T128Device field.
func (o *ConstDeviceGateway) SetT128Device(v bool) {
	o.T128Device = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConstDeviceGateway) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGateway) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConstDeviceGateway) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConstDeviceGateway) SetType(v string) {
	o.Type = &v
}

func (o ConstDeviceGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Defaults) {
		toSerialize["defaults"] = o.Defaults
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Experimental) {
		toSerialize["experimental"] = o.Experimental
	}
	if !IsNil(o.FansPluggable) {
		toSerialize["fans_pluggable"] = o.FansPluggable
	}
	if !IsNil(o.HaNode0Fpc) {
		toSerialize["ha_node0_fpc"] = o.HaNode0Fpc
	}
	if !IsNil(o.HaNode1Fpc) {
		toSerialize["ha_node1_fpc"] = o.HaNode1Fpc
	}
	if !IsNil(o.HasBgp) {
		toSerialize["has_bgp"] = o.HasBgp
	}
	if !IsNil(o.HasFxp0) {
		toSerialize["has_fxp0"] = o.HasFxp0
	}
	if !IsNil(o.HasHaControl) {
		toSerialize["has_ha_control"] = o.HasHaControl
	}
	if !IsNil(o.HasHaData) {
		toSerialize["has_ha_data"] = o.HasHaData
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
	if !IsNil(o.IrbDisabledByDefault) {
		toSerialize["irb_disabled_by_default"] = o.IrbDisabledByDefault
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.NumberFans) {
		toSerialize["number_fans"] = o.NumberFans
	}
	if !IsNil(o.OcDevice) {
		toSerialize["oc_device"] = o.OcDevice
	}
	if !IsNil(o.Pic) {
		toSerialize["pic"] = o.Pic
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.SubRequired) {
		toSerialize["sub_required"] = o.SubRequired
	}
	if !IsNil(o.T128Device) {
		toSerialize["t128_device"] = o.T128Device
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceGateway) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceGateway := _ConstDeviceGateway{}

	err = json.Unmarshal(data, &varConstDeviceGateway)

	if err != nil {
		return err
	}

	*o = ConstDeviceGateway(varConstDeviceGateway)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaults")
		delete(additionalProperties, "description")
		delete(additionalProperties, "experimental")
		delete(additionalProperties, "fans_pluggable")
		delete(additionalProperties, "ha_node0_fpc")
		delete(additionalProperties, "ha_node1_fpc")
		delete(additionalProperties, "has_bgp")
		delete(additionalProperties, "has_fxp0")
		delete(additionalProperties, "has_ha_control")
		delete(additionalProperties, "has_ha_data")
		delete(additionalProperties, "has_irb")
		delete(additionalProperties, "has_poe_out")
		delete(additionalProperties, "has_snapshot")
		delete(additionalProperties, "irb_disabled_by_default")
		delete(additionalProperties, "model")
		delete(additionalProperties, "number_fans")
		delete(additionalProperties, "oc_device")
		delete(additionalProperties, "pic")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "sub_required")
		delete(additionalProperties, "t128_device")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceGateway struct {
	value *ConstDeviceGateway
	isSet bool
}

func (v NullableConstDeviceGateway) Get() *ConstDeviceGateway {
	return v.value
}

func (v *NullableConstDeviceGateway) Set(val *ConstDeviceGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceGateway(val *ConstDeviceGateway) *NullableConstDeviceGateway {
	return &NullableConstDeviceGateway{value: val, isSet: true}
}

func (v NullableConstDeviceGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

