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

// checks if the ResponseDeviceConfigCmd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeviceConfigCmd{}

// ResponseDeviceConfigCmd struct for ResponseDeviceConfigCmd
type ResponseDeviceConfigCmd struct {
	Cmd string `json:"cmd"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDeviceConfigCmd ResponseDeviceConfigCmd

// NewResponseDeviceConfigCmd instantiates a new ResponseDeviceConfigCmd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeviceConfigCmd(cmd string) *ResponseDeviceConfigCmd {
	this := ResponseDeviceConfigCmd{}
	this.Cmd = cmd
	return &this
}

// NewResponseDeviceConfigCmdWithDefaults instantiates a new ResponseDeviceConfigCmd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeviceConfigCmdWithDefaults() *ResponseDeviceConfigCmd {
	this := ResponseDeviceConfigCmd{}
	return &this
}

// GetCmd returns the Cmd field value
func (o *ResponseDeviceConfigCmd) GetCmd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceConfigCmd) GetCmdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmd, true
}

// SetCmd sets field value
func (o *ResponseDeviceConfigCmd) SetCmd(v string) {
	o.Cmd = v
}

func (o ResponseDeviceConfigCmd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeviceConfigCmd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmd"] = o.Cmd

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDeviceConfigCmd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cmd",
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

	varResponseDeviceConfigCmd := _ResponseDeviceConfigCmd{}

	err = json.Unmarshal(data, &varResponseDeviceConfigCmd)

	if err != nil {
		return err
	}

	*o = ResponseDeviceConfigCmd(varResponseDeviceConfigCmd)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cmd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDeviceConfigCmd struct {
	value *ResponseDeviceConfigCmd
	isSet bool
}

func (v NullableResponseDeviceConfigCmd) Get() *ResponseDeviceConfigCmd {
	return v.value
}

func (v *NullableResponseDeviceConfigCmd) Set(val *ResponseDeviceConfigCmd) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeviceConfigCmd) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeviceConfigCmd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeviceConfigCmd(val *ResponseDeviceConfigCmd) *NullableResponseDeviceConfigCmd {
	return &NullableResponseDeviceConfigCmd{value: val, isSet: true}
}

func (v NullableResponseDeviceConfigCmd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeviceConfigCmd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


