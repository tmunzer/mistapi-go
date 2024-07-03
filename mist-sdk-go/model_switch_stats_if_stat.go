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

// checks if the SwitchStatsIfStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchStatsIfStat{}

// SwitchStatsIfStat struct for SwitchStatsIfStat
type SwitchStatsIfStat struct {
	Ips []string `json:"ips,omitempty"`
	PortId *string `json:"port_id,omitempty"`
	RxBytes *int32 `json:"rx_bytes,omitempty"`
	RxPkts *int32 `json:"rx_pkts,omitempty"`
	TxBytes *int32 `json:"tx_bytes,omitempty"`
	TxPkts *int32 `json:"tx_pkts,omitempty"`
	Up *bool `json:"up,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchStatsIfStat SwitchStatsIfStat

// NewSwitchStatsIfStat instantiates a new SwitchStatsIfStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchStatsIfStat() *SwitchStatsIfStat {
	this := SwitchStatsIfStat{}
	return &this
}

// NewSwitchStatsIfStatWithDefaults instantiates a new SwitchStatsIfStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchStatsIfStatWithDefaults() *SwitchStatsIfStat {
	this := SwitchStatsIfStat{}
	return &this
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *SwitchStatsIfStat) SetIps(v []string) {
	o.Ips = v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *SwitchStatsIfStat) SetPortId(v string) {
	o.PortId = &v
}

// GetRxBytes returns the RxBytes field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetRxBytes() int32 {
	if o == nil || IsNil(o.RxBytes) {
		var ret int32
		return ret
	}
	return *o.RxBytes
}

// GetRxBytesOk returns a tuple with the RxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetRxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.RxBytes) {
		return nil, false
	}
	return o.RxBytes, true
}

// HasRxBytes returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasRxBytes() bool {
	if o != nil && !IsNil(o.RxBytes) {
		return true
	}

	return false
}

// SetRxBytes gets a reference to the given int32 and assigns it to the RxBytes field.
func (o *SwitchStatsIfStat) SetRxBytes(v int32) {
	o.RxBytes = &v
}

// GetRxPkts returns the RxPkts field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetRxPkts() int32 {
	if o == nil || IsNil(o.RxPkts) {
		var ret int32
		return ret
	}
	return *o.RxPkts
}

// GetRxPktsOk returns a tuple with the RxPkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetRxPktsOk() (*int32, bool) {
	if o == nil || IsNil(o.RxPkts) {
		return nil, false
	}
	return o.RxPkts, true
}

// HasRxPkts returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasRxPkts() bool {
	if o != nil && !IsNil(o.RxPkts) {
		return true
	}

	return false
}

// SetRxPkts gets a reference to the given int32 and assigns it to the RxPkts field.
func (o *SwitchStatsIfStat) SetRxPkts(v int32) {
	o.RxPkts = &v
}

// GetTxBytes returns the TxBytes field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetTxBytes() int32 {
	if o == nil || IsNil(o.TxBytes) {
		var ret int32
		return ret
	}
	return *o.TxBytes
}

// GetTxBytesOk returns a tuple with the TxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetTxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.TxBytes) {
		return nil, false
	}
	return o.TxBytes, true
}

// HasTxBytes returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasTxBytes() bool {
	if o != nil && !IsNil(o.TxBytes) {
		return true
	}

	return false
}

// SetTxBytes gets a reference to the given int32 and assigns it to the TxBytes field.
func (o *SwitchStatsIfStat) SetTxBytes(v int32) {
	o.TxBytes = &v
}

// GetTxPkts returns the TxPkts field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetTxPkts() int32 {
	if o == nil || IsNil(o.TxPkts) {
		var ret int32
		return ret
	}
	return *o.TxPkts
}

// GetTxPktsOk returns a tuple with the TxPkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetTxPktsOk() (*int32, bool) {
	if o == nil || IsNil(o.TxPkts) {
		return nil, false
	}
	return o.TxPkts, true
}

// HasTxPkts returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasTxPkts() bool {
	if o != nil && !IsNil(o.TxPkts) {
		return true
	}

	return false
}

// SetTxPkts gets a reference to the given int32 and assigns it to the TxPkts field.
func (o *SwitchStatsIfStat) SetTxPkts(v int32) {
	o.TxPkts = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *SwitchStatsIfStat) GetUp() bool {
	if o == nil || IsNil(o.Up) {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchStatsIfStat) GetUpOk() (*bool, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *SwitchStatsIfStat) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *SwitchStatsIfStat) SetUp(v bool) {
	o.Up = &v
}

func (o SwitchStatsIfStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchStatsIfStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	if !IsNil(o.RxBytes) {
		toSerialize["rx_bytes"] = o.RxBytes
	}
	if !IsNil(o.RxPkts) {
		toSerialize["rx_pkts"] = o.RxPkts
	}
	if !IsNil(o.TxBytes) {
		toSerialize["tx_bytes"] = o.TxBytes
	}
	if !IsNil(o.TxPkts) {
		toSerialize["tx_pkts"] = o.TxPkts
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchStatsIfStat) UnmarshalJSON(data []byte) (err error) {
	varSwitchStatsIfStat := _SwitchStatsIfStat{}

	err = json.Unmarshal(data, &varSwitchStatsIfStat)

	if err != nil {
		return err
	}

	*o = SwitchStatsIfStat(varSwitchStatsIfStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ips")
		delete(additionalProperties, "port_id")
		delete(additionalProperties, "rx_bytes")
		delete(additionalProperties, "rx_pkts")
		delete(additionalProperties, "tx_bytes")
		delete(additionalProperties, "tx_pkts")
		delete(additionalProperties, "up")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchStatsIfStat struct {
	value *SwitchStatsIfStat
	isSet bool
}

func (v NullableSwitchStatsIfStat) Get() *SwitchStatsIfStat {
	return v.value
}

func (v *NullableSwitchStatsIfStat) Set(val *SwitchStatsIfStat) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchStatsIfStat) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchStatsIfStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchStatsIfStat(val *SwitchStatsIfStat) *NullableSwitchStatsIfStat {
	return &NullableSwitchStatsIfStat{value: val, isSet: true}
}

func (v NullableSwitchStatsIfStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchStatsIfStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


