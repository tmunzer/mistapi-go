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
	"fmt"
)

// checks if the ResponseDswitchesMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDswitchesMetrics{}

// ResponseDswitchesMetrics struct for ResponseDswitchesMetrics
type ResponseDswitchesMetrics struct {
	InactiveWiredVlans DswitchesMetricsInactiveWiredVlans `json:"inactive_wired_vlans"`
	PoeCompliance DswitchesMetricsPoeCompliance `json:"poe_compliance"`
	SwitchApAffinity DswitchesMetricsSwitchApAffinity `json:"switch_ap_affinity"`
	VersionCompliance DswitchesMetricsVersionCompliance `json:"version_compliance"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDswitchesMetrics ResponseDswitchesMetrics

// NewResponseDswitchesMetrics instantiates a new ResponseDswitchesMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDswitchesMetrics(inactiveWiredVlans DswitchesMetricsInactiveWiredVlans, poeCompliance DswitchesMetricsPoeCompliance, switchApAffinity DswitchesMetricsSwitchApAffinity, versionCompliance DswitchesMetricsVersionCompliance) *ResponseDswitchesMetrics {
	this := ResponseDswitchesMetrics{}
	this.InactiveWiredVlans = inactiveWiredVlans
	this.PoeCompliance = poeCompliance
	this.SwitchApAffinity = switchApAffinity
	this.VersionCompliance = versionCompliance
	return &this
}

// NewResponseDswitchesMetricsWithDefaults instantiates a new ResponseDswitchesMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDswitchesMetricsWithDefaults() *ResponseDswitchesMetrics {
	this := ResponseDswitchesMetrics{}
	return &this
}

// GetInactiveWiredVlans returns the InactiveWiredVlans field value
func (o *ResponseDswitchesMetrics) GetInactiveWiredVlans() DswitchesMetricsInactiveWiredVlans {
	if o == nil {
		var ret DswitchesMetricsInactiveWiredVlans
		return ret
	}

	return o.InactiveWiredVlans
}

// GetInactiveWiredVlansOk returns a tuple with the InactiveWiredVlans field value
// and a boolean to check if the value has been set.
func (o *ResponseDswitchesMetrics) GetInactiveWiredVlansOk() (*DswitchesMetricsInactiveWiredVlans, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InactiveWiredVlans, true
}

// SetInactiveWiredVlans sets field value
func (o *ResponseDswitchesMetrics) SetInactiveWiredVlans(v DswitchesMetricsInactiveWiredVlans) {
	o.InactiveWiredVlans = v
}

// GetPoeCompliance returns the PoeCompliance field value
func (o *ResponseDswitchesMetrics) GetPoeCompliance() DswitchesMetricsPoeCompliance {
	if o == nil {
		var ret DswitchesMetricsPoeCompliance
		return ret
	}

	return o.PoeCompliance
}

// GetPoeComplianceOk returns a tuple with the PoeCompliance field value
// and a boolean to check if the value has been set.
func (o *ResponseDswitchesMetrics) GetPoeComplianceOk() (*DswitchesMetricsPoeCompliance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoeCompliance, true
}

// SetPoeCompliance sets field value
func (o *ResponseDswitchesMetrics) SetPoeCompliance(v DswitchesMetricsPoeCompliance) {
	o.PoeCompliance = v
}

// GetSwitchApAffinity returns the SwitchApAffinity field value
func (o *ResponseDswitchesMetrics) GetSwitchApAffinity() DswitchesMetricsSwitchApAffinity {
	if o == nil {
		var ret DswitchesMetricsSwitchApAffinity
		return ret
	}

	return o.SwitchApAffinity
}

// GetSwitchApAffinityOk returns a tuple with the SwitchApAffinity field value
// and a boolean to check if the value has been set.
func (o *ResponseDswitchesMetrics) GetSwitchApAffinityOk() (*DswitchesMetricsSwitchApAffinity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwitchApAffinity, true
}

// SetSwitchApAffinity sets field value
func (o *ResponseDswitchesMetrics) SetSwitchApAffinity(v DswitchesMetricsSwitchApAffinity) {
	o.SwitchApAffinity = v
}

// GetVersionCompliance returns the VersionCompliance field value
func (o *ResponseDswitchesMetrics) GetVersionCompliance() DswitchesMetricsVersionCompliance {
	if o == nil {
		var ret DswitchesMetricsVersionCompliance
		return ret
	}

	return o.VersionCompliance
}

// GetVersionComplianceOk returns a tuple with the VersionCompliance field value
// and a boolean to check if the value has been set.
func (o *ResponseDswitchesMetrics) GetVersionComplianceOk() (*DswitchesMetricsVersionCompliance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionCompliance, true
}

// SetVersionCompliance sets field value
func (o *ResponseDswitchesMetrics) SetVersionCompliance(v DswitchesMetricsVersionCompliance) {
	o.VersionCompliance = v
}

func (o ResponseDswitchesMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDswitchesMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inactive_wired_vlans"] = o.InactiveWiredVlans
	toSerialize["poe_compliance"] = o.PoeCompliance
	toSerialize["switch_ap_affinity"] = o.SwitchApAffinity
	toSerialize["version_compliance"] = o.VersionCompliance

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDswitchesMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inactive_wired_vlans",
		"poe_compliance",
		"switch_ap_affinity",
		"version_compliance",
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

	varResponseDswitchesMetrics := _ResponseDswitchesMetrics{}

	err = json.Unmarshal(data, &varResponseDswitchesMetrics)

	if err != nil {
		return err
	}

	*o = ResponseDswitchesMetrics(varResponseDswitchesMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "inactive_wired_vlans")
		delete(additionalProperties, "poe_compliance")
		delete(additionalProperties, "switch_ap_affinity")
		delete(additionalProperties, "version_compliance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDswitchesMetrics struct {
	value *ResponseDswitchesMetrics
	isSet bool
}

func (v NullableResponseDswitchesMetrics) Get() *ResponseDswitchesMetrics {
	return v.value
}

func (v *NullableResponseDswitchesMetrics) Set(val *ResponseDswitchesMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDswitchesMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDswitchesMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDswitchesMetrics(val *ResponseDswitchesMetrics) *NullableResponseDswitchesMetrics {
	return &NullableResponseDswitchesMetrics{value: val, isSet: true}
}

func (v NullableResponseDswitchesMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDswitchesMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


