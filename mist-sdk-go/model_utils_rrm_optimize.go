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

// checks if the UtilsRrmOptimize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsRrmOptimize{}

// UtilsRrmOptimize struct for UtilsRrmOptimize
type UtilsRrmOptimize struct {
	// list of bands
	Bands []string `json:"bands"`
	// targeting AP (neighbor APs may get changed, too), default is empty for ALL APs
	Macs []string `json:"macs,omitempty"`
	// only changng TX Power (will not disconnect clients)
	TxpowerOnly *bool `json:"txpower_only,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsRrmOptimize UtilsRrmOptimize

// NewUtilsRrmOptimize instantiates a new UtilsRrmOptimize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsRrmOptimize(bands []string) *UtilsRrmOptimize {
	this := UtilsRrmOptimize{}
	this.Bands = bands
	var txpowerOnly bool = false
	this.TxpowerOnly = &txpowerOnly
	return &this
}

// NewUtilsRrmOptimizeWithDefaults instantiates a new UtilsRrmOptimize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsRrmOptimizeWithDefaults() *UtilsRrmOptimize {
	this := UtilsRrmOptimize{}
	var txpowerOnly bool = false
	this.TxpowerOnly = &txpowerOnly
	return &this
}

// GetBands returns the Bands field value
func (o *UtilsRrmOptimize) GetBands() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Bands
}

// GetBandsOk returns a tuple with the Bands field value
// and a boolean to check if the value has been set.
func (o *UtilsRrmOptimize) GetBandsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bands, true
}

// SetBands sets field value
func (o *UtilsRrmOptimize) SetBands(v []string) {
	o.Bands = v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *UtilsRrmOptimize) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsRrmOptimize) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *UtilsRrmOptimize) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *UtilsRrmOptimize) SetMacs(v []string) {
	o.Macs = v
}

// GetTxpowerOnly returns the TxpowerOnly field value if set, zero value otherwise.
func (o *UtilsRrmOptimize) GetTxpowerOnly() bool {
	if o == nil || IsNil(o.TxpowerOnly) {
		var ret bool
		return ret
	}
	return *o.TxpowerOnly
}

// GetTxpowerOnlyOk returns a tuple with the TxpowerOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsRrmOptimize) GetTxpowerOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.TxpowerOnly) {
		return nil, false
	}
	return o.TxpowerOnly, true
}

// HasTxpowerOnly returns a boolean if a field has been set.
func (o *UtilsRrmOptimize) HasTxpowerOnly() bool {
	if o != nil && !IsNil(o.TxpowerOnly) {
		return true
	}

	return false
}

// SetTxpowerOnly gets a reference to the given bool and assigns it to the TxpowerOnly field.
func (o *UtilsRrmOptimize) SetTxpowerOnly(v bool) {
	o.TxpowerOnly = &v
}

func (o UtilsRrmOptimize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsRrmOptimize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bands"] = o.Bands
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}
	if !IsNil(o.TxpowerOnly) {
		toSerialize["txpower_only"] = o.TxpowerOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsRrmOptimize) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bands",
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

	varUtilsRrmOptimize := _UtilsRrmOptimize{}

	err = json.Unmarshal(data, &varUtilsRrmOptimize)

	if err != nil {
		return err
	}

	*o = UtilsRrmOptimize(varUtilsRrmOptimize)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bands")
		delete(additionalProperties, "macs")
		delete(additionalProperties, "txpower_only")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsRrmOptimize struct {
	value *UtilsRrmOptimize
	isSet bool
}

func (v NullableUtilsRrmOptimize) Get() *UtilsRrmOptimize {
	return v.value
}

func (v *NullableUtilsRrmOptimize) Set(val *UtilsRrmOptimize) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsRrmOptimize) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsRrmOptimize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsRrmOptimize(val *UtilsRrmOptimize) *NullableUtilsRrmOptimize {
	return &NullableUtilsRrmOptimize{value: val, isSet: true}
}

func (v NullableUtilsRrmOptimize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsRrmOptimize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


