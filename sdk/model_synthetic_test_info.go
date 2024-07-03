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

// checks if the SyntheticTestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticTestInfo{}

// SyntheticTestInfo struct for SyntheticTestInfo
type SyntheticTestInfo struct {
	By *string `json:"by,omitempty"`
	DeviceType *SyntheticTestInfoDeviceType `json:"device_type,omitempty"`
	Failed *bool `json:"failed,omitempty"`
	Latency *int32 `json:"latency,omitempty"`
	Mac *string `json:"mac,omitempty"`
	PortId *string `json:"port_id,omitempty"`
	Reason *string `json:"reason,omitempty"`
	RxMbps *int32 `json:"rx_mbps,omitempty"`
	StartTime *int32 `json:"start_time,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *float32 `json:"timestamp,omitempty"`
	TxMbps *int32 `json:"tx_mbps,omitempty"`
	Type *SyntheticTestType `json:"type,omitempty"`
	VlanId *int32 `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyntheticTestInfo SyntheticTestInfo

// NewSyntheticTestInfo instantiates a new SyntheticTestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTestInfo() *SyntheticTestInfo {
	this := SyntheticTestInfo{}
	return &this
}

// NewSyntheticTestInfoWithDefaults instantiates a new SyntheticTestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestInfoWithDefaults() *SyntheticTestInfo {
	this := SyntheticTestInfo{}
	return &this
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetBy() string {
	if o == nil || IsNil(o.By) {
		var ret string
		return ret
	}
	return *o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetByOk() (*string, bool) {
	if o == nil || IsNil(o.By) {
		return nil, false
	}
	return o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasBy() bool {
	if o != nil && !IsNil(o.By) {
		return true
	}

	return false
}

// SetBy gets a reference to the given string and assigns it to the By field.
func (o *SyntheticTestInfo) SetBy(v string) {
	o.By = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetDeviceType() SyntheticTestInfoDeviceType {
	if o == nil || IsNil(o.DeviceType) {
		var ret SyntheticTestInfoDeviceType
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetDeviceTypeOk() (*SyntheticTestInfoDeviceType, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given SyntheticTestInfoDeviceType and assigns it to the DeviceType field.
func (o *SyntheticTestInfo) SetDeviceType(v SyntheticTestInfoDeviceType) {
	o.DeviceType = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *SyntheticTestInfo) SetFailed(v bool) {
	o.Failed = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetLatency() int32 {
	if o == nil || IsNil(o.Latency) {
		var ret int32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasLatency() bool {
	if o != nil && !IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int32 and assigns it to the Latency field.
func (o *SyntheticTestInfo) SetLatency(v int32) {
	o.Latency = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *SyntheticTestInfo) SetMac(v string) {
	o.Mac = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *SyntheticTestInfo) SetPortId(v string) {
	o.PortId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SyntheticTestInfo) SetReason(v string) {
	o.Reason = &v
}

// GetRxMbps returns the RxMbps field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetRxMbps() int32 {
	if o == nil || IsNil(o.RxMbps) {
		var ret int32
		return ret
	}
	return *o.RxMbps
}

// GetRxMbpsOk returns a tuple with the RxMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetRxMbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.RxMbps) {
		return nil, false
	}
	return o.RxMbps, true
}

// HasRxMbps returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasRxMbps() bool {
	if o != nil && !IsNil(o.RxMbps) {
		return true
	}

	return false
}

// SetRxMbps gets a reference to the given int32 and assigns it to the RxMbps field.
func (o *SyntheticTestInfo) SetRxMbps(v int32) {
	o.RxMbps = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetStartTime() int32 {
	if o == nil || IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *SyntheticTestInfo) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticTestInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetTimestamp() float32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret float32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float32 and assigns it to the Timestamp field.
func (o *SyntheticTestInfo) SetTimestamp(v float32) {
	o.Timestamp = &v
}

// GetTxMbps returns the TxMbps field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetTxMbps() int32 {
	if o == nil || IsNil(o.TxMbps) {
		var ret int32
		return ret
	}
	return *o.TxMbps
}

// GetTxMbpsOk returns a tuple with the TxMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetTxMbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.TxMbps) {
		return nil, false
	}
	return o.TxMbps, true
}

// HasTxMbps returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasTxMbps() bool {
	if o != nil && !IsNil(o.TxMbps) {
		return true
	}

	return false
}

// SetTxMbps gets a reference to the given int32 and assigns it to the TxMbps field.
func (o *SyntheticTestInfo) SetTxMbps(v int32) {
	o.TxMbps = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetType() SyntheticTestType {
	if o == nil || IsNil(o.Type) {
		var ret SyntheticTestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetTypeOk() (*SyntheticTestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SyntheticTestType and assigns it to the Type field.
func (o *SyntheticTestInfo) SetType(v SyntheticTestType) {
	o.Type = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *SyntheticTestInfo) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestInfo) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *SyntheticTestInfo) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *SyntheticTestInfo) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o SyntheticTestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticTestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.By) {
		toSerialize["by"] = o.By
	}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.RxMbps) {
		toSerialize["rx_mbps"] = o.RxMbps
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TxMbps) {
		toSerialize["tx_mbps"] = o.TxMbps
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SyntheticTestInfo) UnmarshalJSON(data []byte) (err error) {
	varSyntheticTestInfo := _SyntheticTestInfo{}

	err = json.Unmarshal(data, &varSyntheticTestInfo)

	if err != nil {
		return err
	}

	*o = SyntheticTestInfo(varSyntheticTestInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "by")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "latency")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "port_id")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "rx_mbps")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "tx_mbps")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyntheticTestInfo struct {
	value *SyntheticTestInfo
	isSet bool
}

func (v NullableSyntheticTestInfo) Get() *SyntheticTestInfo {
	return v.value
}

func (v *NullableSyntheticTestInfo) Set(val *SyntheticTestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTestInfo(val *SyntheticTestInfo) *NullableSyntheticTestInfo {
	return &NullableSyntheticTestInfo{value: val, isSet: true}
}

func (v NullableSyntheticTestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


