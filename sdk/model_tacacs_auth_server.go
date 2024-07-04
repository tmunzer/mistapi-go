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

// checks if the TacacsAuthServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TacacsAuthServer{}

// TacacsAuthServer struct for TacacsAuthServer
type TacacsAuthServer struct {
	Host *string `json:"host,omitempty"`
	Port *string `json:"port,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Timeout *int32 `json:"timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TacacsAuthServer TacacsAuthServer

// NewTacacsAuthServer instantiates a new TacacsAuthServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTacacsAuthServer() *TacacsAuthServer {
	this := TacacsAuthServer{}
	var timeout int32 = 10
	this.Timeout = &timeout
	return &this
}

// NewTacacsAuthServerWithDefaults instantiates a new TacacsAuthServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTacacsAuthServerWithDefaults() *TacacsAuthServer {
	this := TacacsAuthServer{}
	var timeout int32 = 10
	this.Timeout = &timeout
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *TacacsAuthServer) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAuthServer) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *TacacsAuthServer) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *TacacsAuthServer) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TacacsAuthServer) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAuthServer) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TacacsAuthServer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *TacacsAuthServer) SetPort(v string) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TacacsAuthServer) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAuthServer) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TacacsAuthServer) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TacacsAuthServer) SetSecret(v string) {
	o.Secret = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *TacacsAuthServer) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAuthServer) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *TacacsAuthServer) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *TacacsAuthServer) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o TacacsAuthServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TacacsAuthServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TacacsAuthServer) UnmarshalJSON(data []byte) (err error) {
	varTacacsAuthServer := _TacacsAuthServer{}

	err = json.Unmarshal(data, &varTacacsAuthServer)

	if err != nil {
		return err
	}

	*o = TacacsAuthServer(varTacacsAuthServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTacacsAuthServer struct {
	value *TacacsAuthServer
	isSet bool
}

func (v NullableTacacsAuthServer) Get() *TacacsAuthServer {
	return v.value
}

func (v *NullableTacacsAuthServer) Set(val *TacacsAuthServer) {
	v.value = val
	v.isSet = true
}

func (v NullableTacacsAuthServer) IsSet() bool {
	return v.isSet
}

func (v *NullableTacacsAuthServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacacsAuthServer(val *TacacsAuthServer) *NullableTacacsAuthServer {
	return &NullableTacacsAuthServer{value: val, isSet: true}
}

func (v NullableTacacsAuthServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTacacsAuthServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

