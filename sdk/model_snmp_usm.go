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

// checks if the SnmpUsm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpUsm{}

// SnmpUsm struct for SnmpUsm
type SnmpUsm struct {
	// required only if `engine_type`==`remote_engine`
	EngineId *string `json:"engine-id,omitempty"`
	EngineType *SnmpUsmEngineType `json:"engine_type,omitempty"`
	Users []SnmpUsmpUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpUsm SnmpUsm

// NewSnmpUsm instantiates a new SnmpUsm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpUsm() *SnmpUsm {
	this := SnmpUsm{}
	return &this
}

// NewSnmpUsmWithDefaults instantiates a new SnmpUsm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpUsmWithDefaults() *SnmpUsm {
	this := SnmpUsm{}
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *SnmpUsm) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsm) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *SnmpUsm) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *SnmpUsm) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineType returns the EngineType field value if set, zero value otherwise.
func (o *SnmpUsm) GetEngineType() SnmpUsmEngineType {
	if o == nil || IsNil(o.EngineType) {
		var ret SnmpUsmEngineType
		return ret
	}
	return *o.EngineType
}

// GetEngineTypeOk returns a tuple with the EngineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsm) GetEngineTypeOk() (*SnmpUsmEngineType, bool) {
	if o == nil || IsNil(o.EngineType) {
		return nil, false
	}
	return o.EngineType, true
}

// HasEngineType returns a boolean if a field has been set.
func (o *SnmpUsm) HasEngineType() bool {
	if o != nil && !IsNil(o.EngineType) {
		return true
	}

	return false
}

// SetEngineType gets a reference to the given SnmpUsmEngineType and assigns it to the EngineType field.
func (o *SnmpUsm) SetEngineType(v SnmpUsmEngineType) {
	o.EngineType = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *SnmpUsm) GetUsers() []SnmpUsmpUser {
	if o == nil || IsNil(o.Users) {
		var ret []SnmpUsmpUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsm) GetUsersOk() ([]SnmpUsmpUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *SnmpUsm) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []SnmpUsmpUser and assigns it to the Users field.
func (o *SnmpUsm) SetUsers(v []SnmpUsmpUser) {
	o.Users = v
}

func (o SnmpUsm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpUsm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineId) {
		toSerialize["engine-id"] = o.EngineId
	}
	if !IsNil(o.EngineType) {
		toSerialize["engine_type"] = o.EngineType
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnmpUsm) UnmarshalJSON(data []byte) (err error) {
	varSnmpUsm := _SnmpUsm{}

	err = json.Unmarshal(data, &varSnmpUsm)

	if err != nil {
		return err
	}

	*o = SnmpUsm(varSnmpUsm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "engine-id")
		delete(additionalProperties, "engine_type")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpUsm struct {
	value *SnmpUsm
	isSet bool
}

func (v NullableSnmpUsm) Get() *SnmpUsm {
	return v.value
}

func (v *NullableSnmpUsm) Set(val *SnmpUsm) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpUsm) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpUsm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpUsm(val *SnmpUsm) *NullableSnmpUsm {
	return &NullableSnmpUsm{value: val, isSet: true}
}

func (v NullableSnmpUsm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpUsm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

