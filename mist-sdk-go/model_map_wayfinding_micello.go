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

// checks if the MapWayfindingMicello type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapWayfindingMicello{}

// MapWayfindingMicello struct for MapWayfindingMicello
type MapWayfindingMicello struct {
	AccountKey *string `json:"account_key,omitempty"`
	DefaultLevelId *int32 `json:"default_level_id,omitempty"`
	MapId *string `json:"map_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MapWayfindingMicello MapWayfindingMicello

// NewMapWayfindingMicello instantiates a new MapWayfindingMicello object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapWayfindingMicello() *MapWayfindingMicello {
	this := MapWayfindingMicello{}
	return &this
}

// NewMapWayfindingMicelloWithDefaults instantiates a new MapWayfindingMicello object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapWayfindingMicelloWithDefaults() *MapWayfindingMicello {
	this := MapWayfindingMicello{}
	return &this
}

// GetAccountKey returns the AccountKey field value if set, zero value otherwise.
func (o *MapWayfindingMicello) GetAccountKey() string {
	if o == nil || IsNil(o.AccountKey) {
		var ret string
		return ret
	}
	return *o.AccountKey
}

// GetAccountKeyOk returns a tuple with the AccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWayfindingMicello) GetAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccountKey) {
		return nil, false
	}
	return o.AccountKey, true
}

// HasAccountKey returns a boolean if a field has been set.
func (o *MapWayfindingMicello) HasAccountKey() bool {
	if o != nil && !IsNil(o.AccountKey) {
		return true
	}

	return false
}

// SetAccountKey gets a reference to the given string and assigns it to the AccountKey field.
func (o *MapWayfindingMicello) SetAccountKey(v string) {
	o.AccountKey = &v
}

// GetDefaultLevelId returns the DefaultLevelId field value if set, zero value otherwise.
func (o *MapWayfindingMicello) GetDefaultLevelId() int32 {
	if o == nil || IsNil(o.DefaultLevelId) {
		var ret int32
		return ret
	}
	return *o.DefaultLevelId
}

// GetDefaultLevelIdOk returns a tuple with the DefaultLevelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWayfindingMicello) GetDefaultLevelIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultLevelId) {
		return nil, false
	}
	return o.DefaultLevelId, true
}

// HasDefaultLevelId returns a boolean if a field has been set.
func (o *MapWayfindingMicello) HasDefaultLevelId() bool {
	if o != nil && !IsNil(o.DefaultLevelId) {
		return true
	}

	return false
}

// SetDefaultLevelId gets a reference to the given int32 and assigns it to the DefaultLevelId field.
func (o *MapWayfindingMicello) SetDefaultLevelId(v int32) {
	o.DefaultLevelId = &v
}

// GetMapId returns the MapId field value if set, zero value otherwise.
func (o *MapWayfindingMicello) GetMapId() string {
	if o == nil || IsNil(o.MapId) {
		var ret string
		return ret
	}
	return *o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWayfindingMicello) GetMapIdOk() (*string, bool) {
	if o == nil || IsNil(o.MapId) {
		return nil, false
	}
	return o.MapId, true
}

// HasMapId returns a boolean if a field has been set.
func (o *MapWayfindingMicello) HasMapId() bool {
	if o != nil && !IsNil(o.MapId) {
		return true
	}

	return false
}

// SetMapId gets a reference to the given string and assigns it to the MapId field.
func (o *MapWayfindingMicello) SetMapId(v string) {
	o.MapId = &v
}

func (o MapWayfindingMicello) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapWayfindingMicello) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountKey) {
		toSerialize["account_key"] = o.AccountKey
	}
	if !IsNil(o.DefaultLevelId) {
		toSerialize["default_level_id"] = o.DefaultLevelId
	}
	if !IsNil(o.MapId) {
		toSerialize["map_id"] = o.MapId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapWayfindingMicello) UnmarshalJSON(data []byte) (err error) {
	varMapWayfindingMicello := _MapWayfindingMicello{}

	err = json.Unmarshal(data, &varMapWayfindingMicello)

	if err != nil {
		return err
	}

	*o = MapWayfindingMicello(varMapWayfindingMicello)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_key")
		delete(additionalProperties, "default_level_id")
		delete(additionalProperties, "map_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapWayfindingMicello struct {
	value *MapWayfindingMicello
	isSet bool
}

func (v NullableMapWayfindingMicello) Get() *MapWayfindingMicello {
	return v.value
}

func (v *NullableMapWayfindingMicello) Set(val *MapWayfindingMicello) {
	v.value = val
	v.isSet = true
}

func (v NullableMapWayfindingMicello) IsSet() bool {
	return v.isSet
}

func (v *NullableMapWayfindingMicello) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapWayfindingMicello(val *MapWayfindingMicello) *NullableMapWayfindingMicello {
	return &NullableMapWayfindingMicello{value: val, isSet: true}
}

func (v NullableMapWayfindingMicello) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapWayfindingMicello) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

