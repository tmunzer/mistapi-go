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

// checks if the TroubleshootCallItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TroubleshootCallItem{}

// TroubleshootCallItem struct for TroubleshootCallItem
type TroubleshootCallItem struct {
	AudioIn *CallTroubleshootData `json:"audio_in,omitempty"`
	AudioOut *CallTroubleshootData `json:"audio_out,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	VideoIn *CallTroubleshootData `json:"video_in,omitempty"`
	VideoOut *CallTroubleshootData `json:"video_out,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TroubleshootCallItem TroubleshootCallItem

// NewTroubleshootCallItem instantiates a new TroubleshootCallItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTroubleshootCallItem() *TroubleshootCallItem {
	this := TroubleshootCallItem{}
	return &this
}

// NewTroubleshootCallItemWithDefaults instantiates a new TroubleshootCallItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTroubleshootCallItemWithDefaults() *TroubleshootCallItem {
	this := TroubleshootCallItem{}
	return &this
}

// GetAudioIn returns the AudioIn field value if set, zero value otherwise.
func (o *TroubleshootCallItem) GetAudioIn() CallTroubleshootData {
	if o == nil || IsNil(o.AudioIn) {
		var ret CallTroubleshootData
		return ret
	}
	return *o.AudioIn
}

// GetAudioInOk returns a tuple with the AudioIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootCallItem) GetAudioInOk() (*CallTroubleshootData, bool) {
	if o == nil || IsNil(o.AudioIn) {
		return nil, false
	}
	return o.AudioIn, true
}

// HasAudioIn returns a boolean if a field has been set.
func (o *TroubleshootCallItem) HasAudioIn() bool {
	if o != nil && !IsNil(o.AudioIn) {
		return true
	}

	return false
}

// SetAudioIn gets a reference to the given CallTroubleshootData and assigns it to the AudioIn field.
func (o *TroubleshootCallItem) SetAudioIn(v CallTroubleshootData) {
	o.AudioIn = &v
}

// GetAudioOut returns the AudioOut field value if set, zero value otherwise.
func (o *TroubleshootCallItem) GetAudioOut() CallTroubleshootData {
	if o == nil || IsNil(o.AudioOut) {
		var ret CallTroubleshootData
		return ret
	}
	return *o.AudioOut
}

// GetAudioOutOk returns a tuple with the AudioOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootCallItem) GetAudioOutOk() (*CallTroubleshootData, bool) {
	if o == nil || IsNil(o.AudioOut) {
		return nil, false
	}
	return o.AudioOut, true
}

// HasAudioOut returns a boolean if a field has been set.
func (o *TroubleshootCallItem) HasAudioOut() bool {
	if o != nil && !IsNil(o.AudioOut) {
		return true
	}

	return false
}

// SetAudioOut gets a reference to the given CallTroubleshootData and assigns it to the AudioOut field.
func (o *TroubleshootCallItem) SetAudioOut(v CallTroubleshootData) {
	o.AudioOut = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TroubleshootCallItem) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootCallItem) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TroubleshootCallItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *TroubleshootCallItem) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetVideoIn returns the VideoIn field value if set, zero value otherwise.
func (o *TroubleshootCallItem) GetVideoIn() CallTroubleshootData {
	if o == nil || IsNil(o.VideoIn) {
		var ret CallTroubleshootData
		return ret
	}
	return *o.VideoIn
}

// GetVideoInOk returns a tuple with the VideoIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootCallItem) GetVideoInOk() (*CallTroubleshootData, bool) {
	if o == nil || IsNil(o.VideoIn) {
		return nil, false
	}
	return o.VideoIn, true
}

// HasVideoIn returns a boolean if a field has been set.
func (o *TroubleshootCallItem) HasVideoIn() bool {
	if o != nil && !IsNil(o.VideoIn) {
		return true
	}

	return false
}

// SetVideoIn gets a reference to the given CallTroubleshootData and assigns it to the VideoIn field.
func (o *TroubleshootCallItem) SetVideoIn(v CallTroubleshootData) {
	o.VideoIn = &v
}

// GetVideoOut returns the VideoOut field value if set, zero value otherwise.
func (o *TroubleshootCallItem) GetVideoOut() CallTroubleshootData {
	if o == nil || IsNil(o.VideoOut) {
		var ret CallTroubleshootData
		return ret
	}
	return *o.VideoOut
}

// GetVideoOutOk returns a tuple with the VideoOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootCallItem) GetVideoOutOk() (*CallTroubleshootData, bool) {
	if o == nil || IsNil(o.VideoOut) {
		return nil, false
	}
	return o.VideoOut, true
}

// HasVideoOut returns a boolean if a field has been set.
func (o *TroubleshootCallItem) HasVideoOut() bool {
	if o != nil && !IsNil(o.VideoOut) {
		return true
	}

	return false
}

// SetVideoOut gets a reference to the given CallTroubleshootData and assigns it to the VideoOut field.
func (o *TroubleshootCallItem) SetVideoOut(v CallTroubleshootData) {
	o.VideoOut = &v
}

func (o TroubleshootCallItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TroubleshootCallItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioIn) {
		toSerialize["audio_in"] = o.AudioIn
	}
	if !IsNil(o.AudioOut) {
		toSerialize["audio_out"] = o.AudioOut
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.VideoIn) {
		toSerialize["video_in"] = o.VideoIn
	}
	if !IsNil(o.VideoOut) {
		toSerialize["video_out"] = o.VideoOut
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TroubleshootCallItem) UnmarshalJSON(data []byte) (err error) {
	varTroubleshootCallItem := _TroubleshootCallItem{}

	err = json.Unmarshal(data, &varTroubleshootCallItem)

	if err != nil {
		return err
	}

	*o = TroubleshootCallItem(varTroubleshootCallItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "audio_in")
		delete(additionalProperties, "audio_out")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "video_in")
		delete(additionalProperties, "video_out")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTroubleshootCallItem struct {
	value *TroubleshootCallItem
	isSet bool
}

func (v NullableTroubleshootCallItem) Get() *TroubleshootCallItem {
	return v.value
}

func (v *NullableTroubleshootCallItem) Set(val *TroubleshootCallItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTroubleshootCallItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTroubleshootCallItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTroubleshootCallItem(val *TroubleshootCallItem) *NullableTroubleshootCallItem {
	return &NullableTroubleshootCallItem{value: val, isSet: true}
}

func (v NullableTroubleshootCallItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTroubleshootCallItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


