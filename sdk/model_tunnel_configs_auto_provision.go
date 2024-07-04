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

// checks if the TunnelConfigsAutoProvision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelConfigsAutoProvision{}

// TunnelConfigsAutoProvision struct for TunnelConfigsAutoProvision
type TunnelConfigsAutoProvision struct {
	Enable *bool `json:"enable,omitempty"`
	Latlng *LatLng `json:"latlng,omitempty"`
	Primary *TunnelConfigsAutoProvisionNode `json:"primary,omitempty"`
	Region *TunnelConfigsAutoProvisionRegion `json:"region,omitempty"`
	Secondary *TunnelConfigsAutoProvisionNode `json:"secondary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelConfigsAutoProvision TunnelConfigsAutoProvision

// NewTunnelConfigsAutoProvision instantiates a new TunnelConfigsAutoProvision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelConfigsAutoProvision() *TunnelConfigsAutoProvision {
	this := TunnelConfigsAutoProvision{}
	var region TunnelConfigsAutoProvisionRegion = TUNNELCONFIGSAUTOPROVISIONREGION_AUTO
	this.Region = &region
	return &this
}

// NewTunnelConfigsAutoProvisionWithDefaults instantiates a new TunnelConfigsAutoProvision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelConfigsAutoProvisionWithDefaults() *TunnelConfigsAutoProvision {
	this := TunnelConfigsAutoProvision{}
	var region TunnelConfigsAutoProvisionRegion = TUNNELCONFIGSAUTOPROVISIONREGION_AUTO
	this.Region = &region
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvision) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvision) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvision) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *TunnelConfigsAutoProvision) SetEnable(v bool) {
	o.Enable = &v
}

// GetLatlng returns the Latlng field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvision) GetLatlng() LatLng {
	if o == nil || IsNil(o.Latlng) {
		var ret LatLng
		return ret
	}
	return *o.Latlng
}

// GetLatlngOk returns a tuple with the Latlng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvision) GetLatlngOk() (*LatLng, bool) {
	if o == nil || IsNil(o.Latlng) {
		return nil, false
	}
	return o.Latlng, true
}

// HasLatlng returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvision) HasLatlng() bool {
	if o != nil && !IsNil(o.Latlng) {
		return true
	}

	return false
}

// SetLatlng gets a reference to the given LatLng and assigns it to the Latlng field.
func (o *TunnelConfigsAutoProvision) SetLatlng(v LatLng) {
	o.Latlng = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvision) GetPrimary() TunnelConfigsAutoProvisionNode {
	if o == nil || IsNil(o.Primary) {
		var ret TunnelConfigsAutoProvisionNode
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvision) GetPrimaryOk() (*TunnelConfigsAutoProvisionNode, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvision) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given TunnelConfigsAutoProvisionNode and assigns it to the Primary field.
func (o *TunnelConfigsAutoProvision) SetPrimary(v TunnelConfigsAutoProvisionNode) {
	o.Primary = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvision) GetRegion() TunnelConfigsAutoProvisionRegion {
	if o == nil || IsNil(o.Region) {
		var ret TunnelConfigsAutoProvisionRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvision) GetRegionOk() (*TunnelConfigsAutoProvisionRegion, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvision) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given TunnelConfigsAutoProvisionRegion and assigns it to the Region field.
func (o *TunnelConfigsAutoProvision) SetRegion(v TunnelConfigsAutoProvisionRegion) {
	o.Region = &v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvision) GetSecondary() TunnelConfigsAutoProvisionNode {
	if o == nil || IsNil(o.Secondary) {
		var ret TunnelConfigsAutoProvisionNode
		return ret
	}
	return *o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvision) GetSecondaryOk() (*TunnelConfigsAutoProvisionNode, bool) {
	if o == nil || IsNil(o.Secondary) {
		return nil, false
	}
	return o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvision) HasSecondary() bool {
	if o != nil && !IsNil(o.Secondary) {
		return true
	}

	return false
}

// SetSecondary gets a reference to the given TunnelConfigsAutoProvisionNode and assigns it to the Secondary field.
func (o *TunnelConfigsAutoProvision) SetSecondary(v TunnelConfigsAutoProvisionNode) {
	o.Secondary = &v
}

func (o TunnelConfigsAutoProvision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelConfigsAutoProvision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Latlng) {
		toSerialize["latlng"] = o.Latlng
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Secondary) {
		toSerialize["secondary"] = o.Secondary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelConfigsAutoProvision) UnmarshalJSON(data []byte) (err error) {
	varTunnelConfigsAutoProvision := _TunnelConfigsAutoProvision{}

	err = json.Unmarshal(data, &varTunnelConfigsAutoProvision)

	if err != nil {
		return err
	}

	*o = TunnelConfigsAutoProvision(varTunnelConfigsAutoProvision)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable")
		delete(additionalProperties, "latlng")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "region")
		delete(additionalProperties, "secondary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelConfigsAutoProvision struct {
	value *TunnelConfigsAutoProvision
	isSet bool
}

func (v NullableTunnelConfigsAutoProvision) Get() *TunnelConfigsAutoProvision {
	return v.value
}

func (v *NullableTunnelConfigsAutoProvision) Set(val *TunnelConfigsAutoProvision) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelConfigsAutoProvision) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelConfigsAutoProvision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelConfigsAutoProvision(val *TunnelConfigsAutoProvision) *NullableTunnelConfigsAutoProvision {
	return &NullableTunnelConfigsAutoProvision{value: val, isSet: true}
}

func (v NullableTunnelConfigsAutoProvision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelConfigsAutoProvision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


