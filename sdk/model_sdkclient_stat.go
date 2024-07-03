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

// checks if the SdkclientStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdkclientStat{}

// SdkclientStat SDK Client statistics
type SdkclientStat struct {
	Id string `json:"id"`
	// last seen timestamp
	LastSeen float32 `json:"last_seen"`
	// map_id of the sdk client (if known), or null
	MapId NullableString `json:"map_id,omitempty"`
	// name of the sdk client (if provided)
	Name *string `json:"name,omitempty"`
	NetworkConnection SdkclientStatNetworkConnection `json:"network_connection"`
	// uuid of the sdk client
	Uuid string `json:"uuid"`
	// x (in pixels) of user location on the map (if known)
	X *float32 `json:"x,omitempty"`
	// y (in pixels) of user location on the map (if known)
	Y *float32 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdkclientStat SdkclientStat

// NewSdkclientStat instantiates a new SdkclientStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkclientStat(id string, lastSeen float32, networkConnection SdkclientStatNetworkConnection, uuid string) *SdkclientStat {
	this := SdkclientStat{}
	this.Id = id
	this.LastSeen = lastSeen
	this.NetworkConnection = networkConnection
	this.Uuid = uuid
	return &this
}

// NewSdkclientStatWithDefaults instantiates a new SdkclientStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkclientStatWithDefaults() *SdkclientStat {
	this := SdkclientStat{}
	return &this
}

// GetId returns the Id field value
func (o *SdkclientStat) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SdkclientStat) SetId(v string) {
	o.Id = v
}

// GetLastSeen returns the LastSeen field value
func (o *SdkclientStat) GetLastSeen() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetLastSeenOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *SdkclientStat) SetLastSeen(v float32) {
	o.LastSeen = v
}

// GetMapId returns the MapId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdkclientStat) GetMapId() string {
	if o == nil || IsNil(o.MapId.Get()) {
		var ret string
		return ret
	}
	return *o.MapId.Get()
}

// GetMapIdOk returns a tuple with the MapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdkclientStat) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MapId.Get(), o.MapId.IsSet()
}

// HasMapId returns a boolean if a field has been set.
func (o *SdkclientStat) HasMapId() bool {
	if o != nil && o.MapId.IsSet() {
		return true
	}

	return false
}

// SetMapId gets a reference to the given NullableString and assigns it to the MapId field.
func (o *SdkclientStat) SetMapId(v string) {
	o.MapId.Set(&v)
}
// SetMapIdNil sets the value for MapId to be an explicit nil
func (o *SdkclientStat) SetMapIdNil() {
	o.MapId.Set(nil)
}

// UnsetMapId ensures that no value is present for MapId, not even an explicit nil
func (o *SdkclientStat) UnsetMapId() {
	o.MapId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SdkclientStat) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SdkclientStat) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SdkclientStat) SetName(v string) {
	o.Name = &v
}

// GetNetworkConnection returns the NetworkConnection field value
func (o *SdkclientStat) GetNetworkConnection() SdkclientStatNetworkConnection {
	if o == nil {
		var ret SdkclientStatNetworkConnection
		return ret
	}

	return o.NetworkConnection
}

// GetNetworkConnectionOk returns a tuple with the NetworkConnection field value
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetNetworkConnectionOk() (*SdkclientStatNetworkConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkConnection, true
}

// SetNetworkConnection sets field value
func (o *SdkclientStat) SetNetworkConnection(v SdkclientStatNetworkConnection) {
	o.NetworkConnection = v
}

// GetUuid returns the Uuid field value
func (o *SdkclientStat) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SdkclientStat) SetUuid(v string) {
	o.Uuid = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SdkclientStat) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SdkclientStat) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *SdkclientStat) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SdkclientStat) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkclientStat) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SdkclientStat) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *SdkclientStat) SetY(v float32) {
	o.Y = &v
}

func (o SdkclientStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdkclientStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["last_seen"] = o.LastSeen
	if o.MapId.IsSet() {
		toSerialize["map_id"] = o.MapId.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["network_connection"] = o.NetworkConnection
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SdkclientStat) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"last_seen",
		"network_connection",
		"uuid",
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

	varSdkclientStat := _SdkclientStat{}

	err = json.Unmarshal(data, &varSdkclientStat)

	if err != nil {
		return err
	}

	*o = SdkclientStat(varSdkclientStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network_connection")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdkclientStat struct {
	value *SdkclientStat
	isSet bool
}

func (v NullableSdkclientStat) Get() *SdkclientStat {
	return v.value
}

func (v *NullableSdkclientStat) Set(val *SdkclientStat) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkclientStat) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkclientStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkclientStat(val *SdkclientStat) *NullableSdkclientStat {
	return &NullableSdkclientStat{value: val, isSet: true}
}

func (v NullableSdkclientStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkclientStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


