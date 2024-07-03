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

// checks if the UtilsShowSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsShowSession{}

// UtilsShowSession struct for UtilsShowSession
type UtilsShowSession struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// The exact service name for which to display the active sessions
	ServiceName *string `json:"service_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsShowSession UtilsShowSession

// NewUtilsShowSession instantiates a new UtilsShowSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsShowSession() *UtilsShowSession {
	this := UtilsShowSession{}
	return &this
}

// NewUtilsShowSessionWithDefaults instantiates a new UtilsShowSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsShowSessionWithDefaults() *UtilsShowSession {
	this := UtilsShowSession{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *UtilsShowSession) GetNode() HaClusterNodeEnum {
	if o == nil || IsNil(o.Node) {
		var ret HaClusterNodeEnum
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowSession) GetNodeOk() (*HaClusterNodeEnum, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *UtilsShowSession) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given HaClusterNodeEnum and assigns it to the Node field.
func (o *UtilsShowSession) SetNode(v HaClusterNodeEnum) {
	o.Node = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *UtilsShowSession) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowSession) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *UtilsShowSession) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *UtilsShowSession) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o UtilsShowSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsShowSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsShowSession) UnmarshalJSON(data []byte) (err error) {
	varUtilsShowSession := _UtilsShowSession{}

	err = json.Unmarshal(data, &varUtilsShowSession)

	if err != nil {
		return err
	}

	*o = UtilsShowSession(varUtilsShowSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "node")
		delete(additionalProperties, "service_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsShowSession struct {
	value *UtilsShowSession
	isSet bool
}

func (v NullableUtilsShowSession) Get() *UtilsShowSession {
	return v.value
}

func (v *NullableUtilsShowSession) Set(val *UtilsShowSession) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsShowSession) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsShowSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsShowSession(val *UtilsShowSession) *NullableUtilsShowSession {
	return &NullableUtilsShowSession{value: val, isSet: true}
}

func (v NullableUtilsShowSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsShowSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


