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

// checks if the BeaconStatsItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeaconStatsItems{}

// BeaconStatsItems struct for BeaconStatsItems
type BeaconStatsItems struct {
	// battery voltage, in mV
	BatteryVoltage *float32 `json:"battery_voltage,omitempty"`
	EddystoneInstance *string `json:"eddystone_instance,omitempty"`
	EddystoneNamespace *string `json:"eddystone_namespace,omitempty"`
	LastSeen float32 `json:"last_seen"`
	Mac string `json:"mac"`
	MapId string `json:"map_id"`
	Name string `json:"name"`
	Power int32 `json:"power"`
	Type string `json:"type"`
	X float32 `json:"x"`
	Y float32 `json:"y"`
	AdditionalProperties map[string]interface{}
}

type _BeaconStatsItems BeaconStatsItems

// NewBeaconStatsItems instantiates a new BeaconStatsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconStatsItems(lastSeen float32, mac string, mapId string, name string, power int32, type_ string, x float32, y float32) *BeaconStatsItems {
	this := BeaconStatsItems{}
	this.LastSeen = lastSeen
	this.Mac = mac
	this.MapId = mapId
	this.Name = name
	this.Power = power
	this.Type = type_
	this.X = x
	this.Y = y
	return &this
}

// NewBeaconStatsItemsWithDefaults instantiates a new BeaconStatsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconStatsItemsWithDefaults() *BeaconStatsItems {
	this := BeaconStatsItems{}
	return &this
}

// GetBatteryVoltage returns the BatteryVoltage field value if set, zero value otherwise.
func (o *BeaconStatsItems) GetBatteryVoltage() float32 {
	if o == nil || IsNil(o.BatteryVoltage) {
		var ret float32
		return ret
	}
	return *o.BatteryVoltage
}

// GetBatteryVoltageOk returns a tuple with the BatteryVoltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetBatteryVoltageOk() (*float32, bool) {
	if o == nil || IsNil(o.BatteryVoltage) {
		return nil, false
	}
	return o.BatteryVoltage, true
}

// HasBatteryVoltage returns a boolean if a field has been set.
func (o *BeaconStatsItems) HasBatteryVoltage() bool {
	if o != nil && !IsNil(o.BatteryVoltage) {
		return true
	}

	return false
}

// SetBatteryVoltage gets a reference to the given float32 and assigns it to the BatteryVoltage field.
func (o *BeaconStatsItems) SetBatteryVoltage(v float32) {
	o.BatteryVoltage = &v
}

// GetEddystoneInstance returns the EddystoneInstance field value if set, zero value otherwise.
func (o *BeaconStatsItems) GetEddystoneInstance() string {
	if o == nil || IsNil(o.EddystoneInstance) {
		var ret string
		return ret
	}
	return *o.EddystoneInstance
}

// GetEddystoneInstanceOk returns a tuple with the EddystoneInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetEddystoneInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneInstance) {
		return nil, false
	}
	return o.EddystoneInstance, true
}

// HasEddystoneInstance returns a boolean if a field has been set.
func (o *BeaconStatsItems) HasEddystoneInstance() bool {
	if o != nil && !IsNil(o.EddystoneInstance) {
		return true
	}

	return false
}

// SetEddystoneInstance gets a reference to the given string and assigns it to the EddystoneInstance field.
func (o *BeaconStatsItems) SetEddystoneInstance(v string) {
	o.EddystoneInstance = &v
}

// GetEddystoneNamespace returns the EddystoneNamespace field value if set, zero value otherwise.
func (o *BeaconStatsItems) GetEddystoneNamespace() string {
	if o == nil || IsNil(o.EddystoneNamespace) {
		var ret string
		return ret
	}
	return *o.EddystoneNamespace
}

// GetEddystoneNamespaceOk returns a tuple with the EddystoneNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetEddystoneNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneNamespace) {
		return nil, false
	}
	return o.EddystoneNamespace, true
}

// HasEddystoneNamespace returns a boolean if a field has been set.
func (o *BeaconStatsItems) HasEddystoneNamespace() bool {
	if o != nil && !IsNil(o.EddystoneNamespace) {
		return true
	}

	return false
}

// SetEddystoneNamespace gets a reference to the given string and assigns it to the EddystoneNamespace field.
func (o *BeaconStatsItems) SetEddystoneNamespace(v string) {
	o.EddystoneNamespace = &v
}

// GetLastSeen returns the LastSeen field value
func (o *BeaconStatsItems) GetLastSeen() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetLastSeenOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *BeaconStatsItems) SetLastSeen(v float32) {
	o.LastSeen = v
}

// GetMac returns the Mac field value
func (o *BeaconStatsItems) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *BeaconStatsItems) SetMac(v string) {
	o.Mac = v
}

// GetMapId returns the MapId field value
func (o *BeaconStatsItems) GetMapId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *BeaconStatsItems) SetMapId(v string) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *BeaconStatsItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BeaconStatsItems) SetName(v string) {
	o.Name = v
}

// GetPower returns the Power field value
func (o *BeaconStatsItems) GetPower() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Power
}

// GetPowerOk returns a tuple with the Power field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetPowerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Power, true
}

// SetPower sets field value
func (o *BeaconStatsItems) SetPower(v int32) {
	o.Power = v
}

// GetType returns the Type field value
func (o *BeaconStatsItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BeaconStatsItems) SetType(v string) {
	o.Type = v
}

// GetX returns the X field value
func (o *BeaconStatsItems) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *BeaconStatsItems) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *BeaconStatsItems) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *BeaconStatsItems) GetYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *BeaconStatsItems) SetY(v float32) {
	o.Y = v
}

func (o BeaconStatsItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeaconStatsItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatteryVoltage) {
		toSerialize["battery_voltage"] = o.BatteryVoltage
	}
	if !IsNil(o.EddystoneInstance) {
		toSerialize["eddystone_instance"] = o.EddystoneInstance
	}
	if !IsNil(o.EddystoneNamespace) {
		toSerialize["eddystone_namespace"] = o.EddystoneNamespace
	}
	toSerialize["last_seen"] = o.LastSeen
	toSerialize["mac"] = o.Mac
	toSerialize["map_id"] = o.MapId
	toSerialize["name"] = o.Name
	toSerialize["power"] = o.Power
	toSerialize["type"] = o.Type
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BeaconStatsItems) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_seen",
		"mac",
		"map_id",
		"name",
		"power",
		"type",
		"x",
		"y",
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

	varBeaconStatsItems := _BeaconStatsItems{}

	err = json.Unmarshal(data, &varBeaconStatsItems)

	if err != nil {
		return err
	}

	*o = BeaconStatsItems(varBeaconStatsItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "battery_voltage")
		delete(additionalProperties, "eddystone_instance")
		delete(additionalProperties, "eddystone_namespace")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "power")
		delete(additionalProperties, "type")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeaconStatsItems struct {
	value *BeaconStatsItems
	isSet bool
}

func (v NullableBeaconStatsItems) Get() *BeaconStatsItems {
	return v.value
}

func (v *NullableBeaconStatsItems) Set(val *BeaconStatsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconStatsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconStatsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconStatsItems(val *BeaconStatsItems) *NullableBeaconStatsItems {
	return &NullableBeaconStatsItems{value: val, isSet: true}
}

func (v NullableBeaconStatsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconStatsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

