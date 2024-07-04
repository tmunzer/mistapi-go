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

// checks if the ResponseSsrUpgradeStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSsrUpgradeStatus{}

// ResponseSsrUpgradeStatus struct for ResponseSsrUpgradeStatus
type ResponseSsrUpgradeStatus struct {
	Channel string `json:"channel"`
	DeviceType *string `json:"device_type,omitempty"`
	Id string `json:"id"`
	Status string `json:"status"`
	Targets ResponseSsrUpgradeStatusTargets `json:"targets"`
	Versions map[string]interface{} `json:"versions"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSsrUpgradeStatus ResponseSsrUpgradeStatus

// NewResponseSsrUpgradeStatus instantiates a new ResponseSsrUpgradeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSsrUpgradeStatus(channel string, id string, status string, targets ResponseSsrUpgradeStatusTargets, versions map[string]interface{}) *ResponseSsrUpgradeStatus {
	this := ResponseSsrUpgradeStatus{}
	this.Channel = channel
	this.Id = id
	this.Status = status
	this.Targets = targets
	this.Versions = versions
	return &this
}

// NewResponseSsrUpgradeStatusWithDefaults instantiates a new ResponseSsrUpgradeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSsrUpgradeStatusWithDefaults() *ResponseSsrUpgradeStatus {
	this := ResponseSsrUpgradeStatus{}
	return &this
}

// GetChannel returns the Channel field value
func (o *ResponseSsrUpgradeStatus) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatus) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ResponseSsrUpgradeStatus) SetChannel(v string) {
	o.Channel = v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *ResponseSsrUpgradeStatus) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatus) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ResponseSsrUpgradeStatus) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *ResponseSsrUpgradeStatus) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetId returns the Id field value
func (o *ResponseSsrUpgradeStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseSsrUpgradeStatus) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *ResponseSsrUpgradeStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResponseSsrUpgradeStatus) SetStatus(v string) {
	o.Status = v
}

// GetTargets returns the Targets field value
func (o *ResponseSsrUpgradeStatus) GetTargets() ResponseSsrUpgradeStatusTargets {
	if o == nil {
		var ret ResponseSsrUpgradeStatusTargets
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatus) GetTargetsOk() (*ResponseSsrUpgradeStatusTargets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value
func (o *ResponseSsrUpgradeStatus) SetTargets(v ResponseSsrUpgradeStatusTargets) {
	o.Targets = v
}

// GetVersions returns the Versions field value
func (o *ResponseSsrUpgradeStatus) GetVersions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatus) GetVersionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *ResponseSsrUpgradeStatus) SetVersions(v map[string]interface{}) {
	o.Versions = v
}

func (o ResponseSsrUpgradeStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSsrUpgradeStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["targets"] = o.Targets
	toSerialize["versions"] = o.Versions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSsrUpgradeStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
		"id",
		"status",
		"targets",
		"versions",
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

	varResponseSsrUpgradeStatus := _ResponseSsrUpgradeStatus{}

	err = json.Unmarshal(data, &varResponseSsrUpgradeStatus)

	if err != nil {
		return err
	}

	*o = ResponseSsrUpgradeStatus(varResponseSsrUpgradeStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "targets")
		delete(additionalProperties, "versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSsrUpgradeStatus struct {
	value *ResponseSsrUpgradeStatus
	isSet bool
}

func (v NullableResponseSsrUpgradeStatus) Get() *ResponseSsrUpgradeStatus {
	return v.value
}

func (v *NullableResponseSsrUpgradeStatus) Set(val *ResponseSsrUpgradeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSsrUpgradeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSsrUpgradeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSsrUpgradeStatus(val *ResponseSsrUpgradeStatus) *NullableResponseSsrUpgradeStatus {
	return &NullableResponseSsrUpgradeStatus{value: val, isSet: true}
}

func (v NullableResponseSsrUpgradeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSsrUpgradeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


