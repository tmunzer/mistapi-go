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

// checks if the ModuleStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStat{}

// ModuleStat struct for ModuleStat
type ModuleStat struct {
	BackupVersion *string `json:"backup_version,omitempty"`
	BiosVersion *string `json:"bios_version,omitempty"`
	// used to report all error states the device node is running into. An error should always have `type` and `since` fields, and could have some other fields specific to that type.
	Errors []ModuleStatErrorsItems `json:"errors,omitempty"`
	Fans []ModuleStatFansItems `json:"fans,omitempty"`
	Model *string `json:"model,omitempty"`
	Poe *ModuleStatPoe `json:"poe,omitempty"`
	Psus []ModuleStatPsusItems `json:"psus,omitempty"`
	Serial *string `json:"serial,omitempty"`
	Temperatures []ModuleStatTemperaturesItems `json:"temperatures,omitempty"`
	VcLinks []ModuleStatVcLinksItems `json:"vc_links,omitempty"`
	VcMode *string `json:"vc_mode,omitempty"`
	// master / backup / linecard
	VcRole *string `json:"vc_role,omitempty"`
	VcState *string `json:"vc_state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleStat ModuleStat

// NewModuleStat instantiates a new ModuleStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStat() *ModuleStat {
	this := ModuleStat{}
	return &this
}

// NewModuleStatWithDefaults instantiates a new ModuleStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatWithDefaults() *ModuleStat {
	this := ModuleStat{}
	return &this
}

// GetBackupVersion returns the BackupVersion field value if set, zero value otherwise.
func (o *ModuleStat) GetBackupVersion() string {
	if o == nil || IsNil(o.BackupVersion) {
		var ret string
		return ret
	}
	return *o.BackupVersion
}

// GetBackupVersionOk returns a tuple with the BackupVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetBackupVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BackupVersion) {
		return nil, false
	}
	return o.BackupVersion, true
}

// HasBackupVersion returns a boolean if a field has been set.
func (o *ModuleStat) HasBackupVersion() bool {
	if o != nil && !IsNil(o.BackupVersion) {
		return true
	}

	return false
}

// SetBackupVersion gets a reference to the given string and assigns it to the BackupVersion field.
func (o *ModuleStat) SetBackupVersion(v string) {
	o.BackupVersion = &v
}

// GetBiosVersion returns the BiosVersion field value if set, zero value otherwise.
func (o *ModuleStat) GetBiosVersion() string {
	if o == nil || IsNil(o.BiosVersion) {
		var ret string
		return ret
	}
	return *o.BiosVersion
}

// GetBiosVersionOk returns a tuple with the BiosVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetBiosVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BiosVersion) {
		return nil, false
	}
	return o.BiosVersion, true
}

// HasBiosVersion returns a boolean if a field has been set.
func (o *ModuleStat) HasBiosVersion() bool {
	if o != nil && !IsNil(o.BiosVersion) {
		return true
	}

	return false
}

// SetBiosVersion gets a reference to the given string and assigns it to the BiosVersion field.
func (o *ModuleStat) SetBiosVersion(v string) {
	o.BiosVersion = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ModuleStat) GetErrors() []ModuleStatErrorsItems {
	if o == nil || IsNil(o.Errors) {
		var ret []ModuleStatErrorsItems
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetErrorsOk() ([]ModuleStatErrorsItems, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ModuleStat) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ModuleStatErrorsItems and assigns it to the Errors field.
func (o *ModuleStat) SetErrors(v []ModuleStatErrorsItems) {
	o.Errors = v
}

// GetFans returns the Fans field value if set, zero value otherwise.
func (o *ModuleStat) GetFans() []ModuleStatFansItems {
	if o == nil || IsNil(o.Fans) {
		var ret []ModuleStatFansItems
		return ret
	}
	return o.Fans
}

// GetFansOk returns a tuple with the Fans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetFansOk() ([]ModuleStatFansItems, bool) {
	if o == nil || IsNil(o.Fans) {
		return nil, false
	}
	return o.Fans, true
}

// HasFans returns a boolean if a field has been set.
func (o *ModuleStat) HasFans() bool {
	if o != nil && !IsNil(o.Fans) {
		return true
	}

	return false
}

// SetFans gets a reference to the given []ModuleStatFansItems and assigns it to the Fans field.
func (o *ModuleStat) SetFans(v []ModuleStatFansItems) {
	o.Fans = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ModuleStat) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ModuleStat) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ModuleStat) SetModel(v string) {
	o.Model = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *ModuleStat) GetPoe() ModuleStatPoe {
	if o == nil || IsNil(o.Poe) {
		var ret ModuleStatPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetPoeOk() (*ModuleStatPoe, bool) {
	if o == nil || IsNil(o.Poe) {
		return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *ModuleStat) HasPoe() bool {
	if o != nil && !IsNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given ModuleStatPoe and assigns it to the Poe field.
func (o *ModuleStat) SetPoe(v ModuleStatPoe) {
	o.Poe = &v
}

// GetPsus returns the Psus field value if set, zero value otherwise.
func (o *ModuleStat) GetPsus() []ModuleStatPsusItems {
	if o == nil || IsNil(o.Psus) {
		var ret []ModuleStatPsusItems
		return ret
	}
	return o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetPsusOk() ([]ModuleStatPsusItems, bool) {
	if o == nil || IsNil(o.Psus) {
		return nil, false
	}
	return o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *ModuleStat) HasPsus() bool {
	if o != nil && !IsNil(o.Psus) {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []ModuleStatPsusItems and assigns it to the Psus field.
func (o *ModuleStat) SetPsus(v []ModuleStatPsusItems) {
	o.Psus = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *ModuleStat) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *ModuleStat) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *ModuleStat) SetSerial(v string) {
	o.Serial = &v
}

// GetTemperatures returns the Temperatures field value if set, zero value otherwise.
func (o *ModuleStat) GetTemperatures() []ModuleStatTemperaturesItems {
	if o == nil || IsNil(o.Temperatures) {
		var ret []ModuleStatTemperaturesItems
		return ret
	}
	return o.Temperatures
}

// GetTemperaturesOk returns a tuple with the Temperatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetTemperaturesOk() ([]ModuleStatTemperaturesItems, bool) {
	if o == nil || IsNil(o.Temperatures) {
		return nil, false
	}
	return o.Temperatures, true
}

// HasTemperatures returns a boolean if a field has been set.
func (o *ModuleStat) HasTemperatures() bool {
	if o != nil && !IsNil(o.Temperatures) {
		return true
	}

	return false
}

// SetTemperatures gets a reference to the given []ModuleStatTemperaturesItems and assigns it to the Temperatures field.
func (o *ModuleStat) SetTemperatures(v []ModuleStatTemperaturesItems) {
	o.Temperatures = v
}

// GetVcLinks returns the VcLinks field value if set, zero value otherwise.
func (o *ModuleStat) GetVcLinks() []ModuleStatVcLinksItems {
	if o == nil || IsNil(o.VcLinks) {
		var ret []ModuleStatVcLinksItems
		return ret
	}
	return o.VcLinks
}

// GetVcLinksOk returns a tuple with the VcLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetVcLinksOk() ([]ModuleStatVcLinksItems, bool) {
	if o == nil || IsNil(o.VcLinks) {
		return nil, false
	}
	return o.VcLinks, true
}

// HasVcLinks returns a boolean if a field has been set.
func (o *ModuleStat) HasVcLinks() bool {
	if o != nil && !IsNil(o.VcLinks) {
		return true
	}

	return false
}

// SetVcLinks gets a reference to the given []ModuleStatVcLinksItems and assigns it to the VcLinks field.
func (o *ModuleStat) SetVcLinks(v []ModuleStatVcLinksItems) {
	o.VcLinks = v
}

// GetVcMode returns the VcMode field value if set, zero value otherwise.
func (o *ModuleStat) GetVcMode() string {
	if o == nil || IsNil(o.VcMode) {
		var ret string
		return ret
	}
	return *o.VcMode
}

// GetVcModeOk returns a tuple with the VcMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetVcModeOk() (*string, bool) {
	if o == nil || IsNil(o.VcMode) {
		return nil, false
	}
	return o.VcMode, true
}

// HasVcMode returns a boolean if a field has been set.
func (o *ModuleStat) HasVcMode() bool {
	if o != nil && !IsNil(o.VcMode) {
		return true
	}

	return false
}

// SetVcMode gets a reference to the given string and assigns it to the VcMode field.
func (o *ModuleStat) SetVcMode(v string) {
	o.VcMode = &v
}

// GetVcRole returns the VcRole field value if set, zero value otherwise.
func (o *ModuleStat) GetVcRole() string {
	if o == nil || IsNil(o.VcRole) {
		var ret string
		return ret
	}
	return *o.VcRole
}

// GetVcRoleOk returns a tuple with the VcRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetVcRoleOk() (*string, bool) {
	if o == nil || IsNil(o.VcRole) {
		return nil, false
	}
	return o.VcRole, true
}

// HasVcRole returns a boolean if a field has been set.
func (o *ModuleStat) HasVcRole() bool {
	if o != nil && !IsNil(o.VcRole) {
		return true
	}

	return false
}

// SetVcRole gets a reference to the given string and assigns it to the VcRole field.
func (o *ModuleStat) SetVcRole(v string) {
	o.VcRole = &v
}

// GetVcState returns the VcState field value if set, zero value otherwise.
func (o *ModuleStat) GetVcState() string {
	if o == nil || IsNil(o.VcState) {
		var ret string
		return ret
	}
	return *o.VcState
}

// GetVcStateOk returns a tuple with the VcState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStat) GetVcStateOk() (*string, bool) {
	if o == nil || IsNil(o.VcState) {
		return nil, false
	}
	return o.VcState, true
}

// HasVcState returns a boolean if a field has been set.
func (o *ModuleStat) HasVcState() bool {
	if o != nil && !IsNil(o.VcState) {
		return true
	}

	return false
}

// SetVcState gets a reference to the given string and assigns it to the VcState field.
func (o *ModuleStat) SetVcState(v string) {
	o.VcState = &v
}

func (o ModuleStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupVersion) {
		toSerialize["backup_version"] = o.BackupVersion
	}
	if !IsNil(o.BiosVersion) {
		toSerialize["bios_version"] = o.BiosVersion
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Fans) {
		toSerialize["fans"] = o.Fans
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	if !IsNil(o.Psus) {
		toSerialize["psus"] = o.Psus
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Temperatures) {
		toSerialize["temperatures"] = o.Temperatures
	}
	if !IsNil(o.VcLinks) {
		toSerialize["vc_links"] = o.VcLinks
	}
	if !IsNil(o.VcMode) {
		toSerialize["vc_mode"] = o.VcMode
	}
	if !IsNil(o.VcRole) {
		toSerialize["vc_role"] = o.VcRole
	}
	if !IsNil(o.VcState) {
		toSerialize["vc_state"] = o.VcState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleStat) UnmarshalJSON(data []byte) (err error) {
	varModuleStat := _ModuleStat{}

	err = json.Unmarshal(data, &varModuleStat)

	if err != nil {
		return err
	}

	*o = ModuleStat(varModuleStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "backup_version")
		delete(additionalProperties, "bios_version")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "fans")
		delete(additionalProperties, "model")
		delete(additionalProperties, "poe")
		delete(additionalProperties, "psus")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "temperatures")
		delete(additionalProperties, "vc_links")
		delete(additionalProperties, "vc_mode")
		delete(additionalProperties, "vc_role")
		delete(additionalProperties, "vc_state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleStat struct {
	value *ModuleStat
	isSet bool
}

func (v NullableModuleStat) Get() *ModuleStat {
	return v.value
}

func (v *NullableModuleStat) Set(val *ModuleStat) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStat) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStat(val *ModuleStat) *NullableModuleStat {
	return &NullableModuleStat{value: val, isSet: true}
}

func (v NullableModuleStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


