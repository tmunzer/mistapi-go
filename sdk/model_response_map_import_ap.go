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

// checks if the ResponseMapImportAp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMapImportAp{}

// ResponseMapImportAp struct for ResponseMapImportAp
type ResponseMapImportAp struct {
	Action ResponseMapImportApAction `json:"action"`
	FloorplanId string `json:"floorplan_id"`
	Height *float32 `json:"height,omitempty"`
	Mac string `json:"mac"`
	MapId string `json:"map_id"`
	Orientation float32 `json:"orientation"`
	Reason *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseMapImportAp ResponseMapImportAp

// NewResponseMapImportAp instantiates a new ResponseMapImportAp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMapImportAp(action ResponseMapImportApAction, floorplanId string, mac string, mapId string, orientation float32) *ResponseMapImportAp {
	this := ResponseMapImportAp{}
	this.Action = action
	this.FloorplanId = floorplanId
	this.Mac = mac
	this.MapId = mapId
	this.Orientation = orientation
	return &this
}

// NewResponseMapImportApWithDefaults instantiates a new ResponseMapImportAp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMapImportApWithDefaults() *ResponseMapImportAp {
	this := ResponseMapImportAp{}
	return &this
}

// GetAction returns the Action field value
func (o *ResponseMapImportAp) GetAction() ResponseMapImportApAction {
	if o == nil {
		var ret ResponseMapImportApAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetActionOk() (*ResponseMapImportApAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ResponseMapImportAp) SetAction(v ResponseMapImportApAction) {
	o.Action = v
}

// GetFloorplanId returns the FloorplanId field value
func (o *ResponseMapImportAp) GetFloorplanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FloorplanId
}

// GetFloorplanIdOk returns a tuple with the FloorplanId field value
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetFloorplanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FloorplanId, true
}

// SetFloorplanId sets field value
func (o *ResponseMapImportAp) SetFloorplanId(v string) {
	o.FloorplanId = v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponseMapImportAp) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponseMapImportAp) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *ResponseMapImportAp) SetHeight(v float32) {
	o.Height = &v
}

// GetMac returns the Mac field value
func (o *ResponseMapImportAp) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *ResponseMapImportAp) SetMac(v string) {
	o.Mac = v
}

// GetMapId returns the MapId field value
func (o *ResponseMapImportAp) GetMapId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *ResponseMapImportAp) SetMapId(v string) {
	o.MapId = v
}

// GetOrientation returns the Orientation field value
func (o *ResponseMapImportAp) GetOrientation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetOrientationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orientation, true
}

// SetOrientation sets field value
func (o *ResponseMapImportAp) SetOrientation(v float32) {
	o.Orientation = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ResponseMapImportAp) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMapImportAp) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ResponseMapImportAp) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ResponseMapImportAp) SetReason(v string) {
	o.Reason = &v
}

func (o ResponseMapImportAp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMapImportAp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["floorplan_id"] = o.FloorplanId
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	toSerialize["mac"] = o.Mac
	toSerialize["map_id"] = o.MapId
	toSerialize["orientation"] = o.Orientation
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseMapImportAp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"floorplan_id",
		"mac",
		"map_id",
		"orientation",
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

	varResponseMapImportAp := _ResponseMapImportAp{}

	err = json.Unmarshal(data, &varResponseMapImportAp)

	if err != nil {
		return err
	}

	*o = ResponseMapImportAp(varResponseMapImportAp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "floorplan_id")
		delete(additionalProperties, "height")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "orientation")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseMapImportAp struct {
	value *ResponseMapImportAp
	isSet bool
}

func (v NullableResponseMapImportAp) Get() *ResponseMapImportAp {
	return v.value
}

func (v *NullableResponseMapImportAp) Set(val *ResponseMapImportAp) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMapImportAp) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMapImportAp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMapImportAp(val *ResponseMapImportAp) *NullableResponseMapImportAp {
	return &NullableResponseMapImportAp{value: val, isSet: true}
}

func (v NullableResponseMapImportAp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMapImportAp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


