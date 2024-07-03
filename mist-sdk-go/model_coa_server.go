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

// checks if the CoaServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoaServer{}

// CoaServer CoA Server
type CoaServer struct {
	// whether to disable Event-Timestamp Check
	DisableEventTimestampCheck *bool `json:"disable_event_timestamp_check,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Ip string `json:"ip"`
	Port *int32 `json:"port,omitempty"`
	Secret string `json:"secret"`
	AdditionalProperties map[string]interface{}
}

type _CoaServer CoaServer

// NewCoaServer instantiates a new CoaServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoaServer(ip string, secret string) *CoaServer {
	this := CoaServer{}
	var disableEventTimestampCheck bool = false
	this.DisableEventTimestampCheck = &disableEventTimestampCheck
	var enabled bool = false
	this.Enabled = &enabled
	this.Ip = ip
	var port int32 = 3799
	this.Port = &port
	this.Secret = secret
	return &this
}

// NewCoaServerWithDefaults instantiates a new CoaServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoaServerWithDefaults() *CoaServer {
	this := CoaServer{}
	var disableEventTimestampCheck bool = false
	this.DisableEventTimestampCheck = &disableEventTimestampCheck
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 3799
	this.Port = &port
	return &this
}

// GetDisableEventTimestampCheck returns the DisableEventTimestampCheck field value if set, zero value otherwise.
func (o *CoaServer) GetDisableEventTimestampCheck() bool {
	if o == nil || IsNil(o.DisableEventTimestampCheck) {
		var ret bool
		return ret
	}
	return *o.DisableEventTimestampCheck
}

// GetDisableEventTimestampCheckOk returns a tuple with the DisableEventTimestampCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoaServer) GetDisableEventTimestampCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableEventTimestampCheck) {
		return nil, false
	}
	return o.DisableEventTimestampCheck, true
}

// HasDisableEventTimestampCheck returns a boolean if a field has been set.
func (o *CoaServer) HasDisableEventTimestampCheck() bool {
	if o != nil && !IsNil(o.DisableEventTimestampCheck) {
		return true
	}

	return false
}

// SetDisableEventTimestampCheck gets a reference to the given bool and assigns it to the DisableEventTimestampCheck field.
func (o *CoaServer) SetDisableEventTimestampCheck(v bool) {
	o.DisableEventTimestampCheck = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CoaServer) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoaServer) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CoaServer) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CoaServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIp returns the Ip field value
func (o *CoaServer) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *CoaServer) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *CoaServer) SetIp(v string) {
	o.Ip = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CoaServer) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoaServer) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CoaServer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *CoaServer) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value
func (o *CoaServer) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CoaServer) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CoaServer) SetSecret(v string) {
	o.Secret = v
}

func (o CoaServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoaServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisableEventTimestampCheck) {
		toSerialize["disable_event_timestamp_check"] = o.DisableEventTimestampCheck
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["ip"] = o.Ip
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoaServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"secret",
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

	varCoaServer := _CoaServer{}

	err = json.Unmarshal(data, &varCoaServer)

	if err != nil {
		return err
	}

	*o = CoaServer(varCoaServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disable_event_timestamp_check")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "port")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoaServer struct {
	value *CoaServer
	isSet bool
}

func (v NullableCoaServer) Get() *CoaServer {
	return v.value
}

func (v *NullableCoaServer) Set(val *CoaServer) {
	v.value = val
	v.isSet = true
}

func (v NullableCoaServer) IsSet() bool {
	return v.isSet
}

func (v *NullableCoaServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoaServer(val *CoaServer) *NullableCoaServer {
	return &NullableCoaServer{value: val, isSet: true}
}

func (v NullableCoaServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoaServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


