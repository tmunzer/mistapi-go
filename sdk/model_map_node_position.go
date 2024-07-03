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

// checks if the MapNodePosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapNodePosition{}

// MapNodePosition struct for MapNodePosition
type MapNodePosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	AdditionalProperties map[string]interface{}
}

type _MapNodePosition MapNodePosition

// NewMapNodePosition instantiates a new MapNodePosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapNodePosition(x float32, y float32) *MapNodePosition {
	this := MapNodePosition{}
	this.X = x
	this.Y = y
	return &this
}

// NewMapNodePositionWithDefaults instantiates a new MapNodePosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapNodePositionWithDefaults() *MapNodePosition {
	this := MapNodePosition{}
	return &this
}

// GetX returns the X field value
func (o *MapNodePosition) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *MapNodePosition) GetXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *MapNodePosition) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *MapNodePosition) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *MapNodePosition) GetYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *MapNodePosition) SetY(v float32) {
	o.Y = v
}

func (o MapNodePosition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapNodePosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapNodePosition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varMapNodePosition := _MapNodePosition{}

	err = json.Unmarshal(data, &varMapNodePosition)

	if err != nil {
		return err
	}

	*o = MapNodePosition(varMapNodePosition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapNodePosition struct {
	value *MapNodePosition
	isSet bool
}

func (v NullableMapNodePosition) Get() *MapNodePosition {
	return v.value
}

func (v *NullableMapNodePosition) Set(val *MapNodePosition) {
	v.value = val
	v.isSet = true
}

func (v NullableMapNodePosition) IsSet() bool {
	return v.isSet
}

func (v *NullableMapNodePosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapNodePosition(val *MapNodePosition) *NullableMapNodePosition {
	return &NullableMapNodePosition{value: val, isSet: true}
}

func (v NullableMapNodePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapNodePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


