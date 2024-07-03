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

// checks if the RrmBandMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RrmBandMetric{}

// RrmBandMetric struct for RrmBandMetric
type RrmBandMetric struct {
	// average number of co-channel neighbors
	CochannelNeighbors float32 `json:"cochannel_neighbors"`
	// defined by how APs can hear from one and another, 0 - 1 (everyone can hear everyone)
	Density float32 `json:"density"`
	// Property key is the channel number
	Interferences *map[string]RrmBandMetricInterference `json:"interferences,omitempty"`
	// average number of neighbors
	Neighbors float32 `json:"neighbors"`
	// average noise in dBm
	Noise float32 `json:"noise"`
	AdditionalProperties map[string]interface{}
}

type _RrmBandMetric RrmBandMetric

// NewRrmBandMetric instantiates a new RrmBandMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRrmBandMetric(cochannelNeighbors float32, density float32, neighbors float32, noise float32) *RrmBandMetric {
	this := RrmBandMetric{}
	this.CochannelNeighbors = cochannelNeighbors
	this.Density = density
	this.Neighbors = neighbors
	this.Noise = noise
	return &this
}

// NewRrmBandMetricWithDefaults instantiates a new RrmBandMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRrmBandMetricWithDefaults() *RrmBandMetric {
	this := RrmBandMetric{}
	return &this
}

// GetCochannelNeighbors returns the CochannelNeighbors field value
func (o *RrmBandMetric) GetCochannelNeighbors() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CochannelNeighbors
}

// GetCochannelNeighborsOk returns a tuple with the CochannelNeighbors field value
// and a boolean to check if the value has been set.
func (o *RrmBandMetric) GetCochannelNeighborsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CochannelNeighbors, true
}

// SetCochannelNeighbors sets field value
func (o *RrmBandMetric) SetCochannelNeighbors(v float32) {
	o.CochannelNeighbors = v
}

// GetDensity returns the Density field value
func (o *RrmBandMetric) GetDensity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Density
}

// GetDensityOk returns a tuple with the Density field value
// and a boolean to check if the value has been set.
func (o *RrmBandMetric) GetDensityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Density, true
}

// SetDensity sets field value
func (o *RrmBandMetric) SetDensity(v float32) {
	o.Density = v
}

// GetInterferences returns the Interferences field value if set, zero value otherwise.
func (o *RrmBandMetric) GetInterferences() map[string]RrmBandMetricInterference {
	if o == nil || IsNil(o.Interferences) {
		var ret map[string]RrmBandMetricInterference
		return ret
	}
	return *o.Interferences
}

// GetInterferencesOk returns a tuple with the Interferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmBandMetric) GetInterferencesOk() (*map[string]RrmBandMetricInterference, bool) {
	if o == nil || IsNil(o.Interferences) {
		return nil, false
	}
	return o.Interferences, true
}

// HasInterferences returns a boolean if a field has been set.
func (o *RrmBandMetric) HasInterferences() bool {
	if o != nil && !IsNil(o.Interferences) {
		return true
	}

	return false
}

// SetInterferences gets a reference to the given map[string]RrmBandMetricInterference and assigns it to the Interferences field.
func (o *RrmBandMetric) SetInterferences(v map[string]RrmBandMetricInterference) {
	o.Interferences = &v
}

// GetNeighbors returns the Neighbors field value
func (o *RrmBandMetric) GetNeighbors() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value
// and a boolean to check if the value has been set.
func (o *RrmBandMetric) GetNeighborsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Neighbors, true
}

// SetNeighbors sets field value
func (o *RrmBandMetric) SetNeighbors(v float32) {
	o.Neighbors = v
}

// GetNoise returns the Noise field value
func (o *RrmBandMetric) GetNoise() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value
// and a boolean to check if the value has been set.
func (o *RrmBandMetric) GetNoiseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Noise, true
}

// SetNoise sets field value
func (o *RrmBandMetric) SetNoise(v float32) {
	o.Noise = v
}

func (o RrmBandMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RrmBandMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cochannel_neighbors"] = o.CochannelNeighbors
	toSerialize["density"] = o.Density
	if !IsNil(o.Interferences) {
		toSerialize["interferences"] = o.Interferences
	}
	toSerialize["neighbors"] = o.Neighbors
	toSerialize["noise"] = o.Noise

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RrmBandMetric) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cochannel_neighbors",
		"density",
		"neighbors",
		"noise",
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

	varRrmBandMetric := _RrmBandMetric{}

	err = json.Unmarshal(data, &varRrmBandMetric)

	if err != nil {
		return err
	}

	*o = RrmBandMetric(varRrmBandMetric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cochannel_neighbors")
		delete(additionalProperties, "density")
		delete(additionalProperties, "interferences")
		delete(additionalProperties, "neighbors")
		delete(additionalProperties, "noise")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRrmBandMetric struct {
	value *RrmBandMetric
	isSet bool
}

func (v NullableRrmBandMetric) Get() *RrmBandMetric {
	return v.value
}

func (v *NullableRrmBandMetric) Set(val *RrmBandMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableRrmBandMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableRrmBandMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRrmBandMetric(val *RrmBandMetric) *NullableRrmBandMetric {
	return &NullableRrmBandMetric{value: val, isSet: true}
}

func (v NullableRrmBandMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRrmBandMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


