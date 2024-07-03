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

// checks if the OtherDeviceUpdateMulti type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtherDeviceUpdateMulti{}

// OtherDeviceUpdateMulti struct for OtherDeviceUpdateMulti
type OtherDeviceUpdateMulti struct {
	// The mac address of the peer device.
	Macs []string `json:"macs,omitempty"`
	Op OtherDeviceUpdateOperation `json:"op"`
	SiteId *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OtherDeviceUpdateMulti OtherDeviceUpdateMulti

// NewOtherDeviceUpdateMulti instantiates a new OtherDeviceUpdateMulti object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherDeviceUpdateMulti(op OtherDeviceUpdateOperation) *OtherDeviceUpdateMulti {
	this := OtherDeviceUpdateMulti{}
	this.Op = op
	return &this
}

// NewOtherDeviceUpdateMultiWithDefaults instantiates a new OtherDeviceUpdateMulti object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherDeviceUpdateMultiWithDefaults() *OtherDeviceUpdateMulti {
	this := OtherDeviceUpdateMulti{}
	return &this
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *OtherDeviceUpdateMulti) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherDeviceUpdateMulti) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *OtherDeviceUpdateMulti) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *OtherDeviceUpdateMulti) SetMacs(v []string) {
	o.Macs = v
}

// GetOp returns the Op field value
func (o *OtherDeviceUpdateMulti) GetOp() OtherDeviceUpdateOperation {
	if o == nil {
		var ret OtherDeviceUpdateOperation
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *OtherDeviceUpdateMulti) GetOpOk() (*OtherDeviceUpdateOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *OtherDeviceUpdateMulti) SetOp(v OtherDeviceUpdateOperation) {
	o.Op = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *OtherDeviceUpdateMulti) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherDeviceUpdateMulti) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *OtherDeviceUpdateMulti) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *OtherDeviceUpdateMulti) SetSiteId(v string) {
	o.SiteId = &v
}

func (o OtherDeviceUpdateMulti) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtherDeviceUpdateMulti) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}
	toSerialize["op"] = o.Op
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OtherDeviceUpdateMulti) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
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

	varOtherDeviceUpdateMulti := _OtherDeviceUpdateMulti{}

	err = json.Unmarshal(data, &varOtherDeviceUpdateMulti)

	if err != nil {
		return err
	}

	*o = OtherDeviceUpdateMulti(varOtherDeviceUpdateMulti)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "macs")
		delete(additionalProperties, "op")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOtherDeviceUpdateMulti struct {
	value *OtherDeviceUpdateMulti
	isSet bool
}

func (v NullableOtherDeviceUpdateMulti) Get() *OtherDeviceUpdateMulti {
	return v.value
}

func (v *NullableOtherDeviceUpdateMulti) Set(val *OtherDeviceUpdateMulti) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherDeviceUpdateMulti) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherDeviceUpdateMulti) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherDeviceUpdateMulti(val *OtherDeviceUpdateMulti) *NullableOtherDeviceUpdateMulti {
	return &NullableOtherDeviceUpdateMulti{value: val, isSet: true}
}

func (v NullableOtherDeviceUpdateMulti) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherDeviceUpdateMulti) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


