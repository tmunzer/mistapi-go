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

// checks if the SwitchStatsNumClientsTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchStatsNumClientsTotal{}

// SwitchStatsNumClientsTotal struct for SwitchStatsNumClientsTotal
type SwitchStatsNumClientsTotal struct {
	NumAps *string `json:"num_aps,omitempty"`
	NumClients *string `json:"num_clients,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchStatsNumClientsTotal SwitchStatsNumClientsTotal

// NewSwitchStatsNumClientsTotal instantiates a new SwitchStatsNumClientsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchStatsNumClientsTotal() *SwitchStatsNumClientsTotal {
	this := SwitchStatsNumClientsTotal{}
	return &this
}

// NewSwitchStatsNumClientsTotalWithDefaults instantiates a new SwitchStatsNumClientsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchStatsNumClientsTotalWithDefaults() *SwitchStatsNumClientsTotal {
	this := SwitchStatsNumClientsTotal{}
	return &this
}

// GetNumAps returns the NumAps field value if set, zero value otherwise.
func (o *SwitchStatsNumClientsTotal) GetNumAps() string {
	if o == nil || IsNil(o.NumAps) {
		var ret string
		return ret
	}
	return *o.NumAps
}

// GetNumApsOk returns a tuple with the NumAps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsNumClientsTotal) GetNumApsOk() (*string, bool) {
	if o == nil || IsNil(o.NumAps) {
		return nil, false
	}
	return o.NumAps, true
}

// HasNumAps returns a boolean if a field has been set.
func (o *SwitchStatsNumClientsTotal) HasNumAps() bool {
	if o != nil && !IsNil(o.NumAps) {
		return true
	}

	return false
}

// SetNumAps gets a reference to the given string and assigns it to the NumAps field.
func (o *SwitchStatsNumClientsTotal) SetNumAps(v string) {
	o.NumAps = &v
}

// GetNumClients returns the NumClients field value if set, zero value otherwise.
func (o *SwitchStatsNumClientsTotal) GetNumClients() string {
	if o == nil || IsNil(o.NumClients) {
		var ret string
		return ret
	}
	return *o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsNumClientsTotal) GetNumClientsOk() (*string, bool) {
	if o == nil || IsNil(o.NumClients) {
		return nil, false
	}
	return o.NumClients, true
}

// HasNumClients returns a boolean if a field has been set.
func (o *SwitchStatsNumClientsTotal) HasNumClients() bool {
	if o != nil && !IsNil(o.NumClients) {
		return true
	}

	return false
}

// SetNumClients gets a reference to the given string and assigns it to the NumClients field.
func (o *SwitchStatsNumClientsTotal) SetNumClients(v string) {
	o.NumClients = &v
}

func (o SwitchStatsNumClientsTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchStatsNumClientsTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumAps) {
		toSerialize["num_aps"] = o.NumAps
	}
	if !IsNil(o.NumClients) {
		toSerialize["num_clients"] = o.NumClients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchStatsNumClientsTotal) UnmarshalJSON(data []byte) (err error) {
	varSwitchStatsNumClientsTotal := _SwitchStatsNumClientsTotal{}

	err = json.Unmarshal(data, &varSwitchStatsNumClientsTotal)

	if err != nil {
		return err
	}

	*o = SwitchStatsNumClientsTotal(varSwitchStatsNumClientsTotal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "num_aps")
		delete(additionalProperties, "num_clients")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchStatsNumClientsTotal struct {
	value *SwitchStatsNumClientsTotal
	isSet bool
}

func (v NullableSwitchStatsNumClientsTotal) Get() *SwitchStatsNumClientsTotal {
	return v.value
}

func (v *NullableSwitchStatsNumClientsTotal) Set(val *SwitchStatsNumClientsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchStatsNumClientsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchStatsNumClientsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchStatsNumClientsTotal(val *SwitchStatsNumClientsTotal) *NullableSwitchStatsNumClientsTotal {
	return &NullableSwitchStatsNumClientsTotal{value: val, isSet: true}
}

func (v NullableSwitchStatsNumClientsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchStatsNumClientsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


