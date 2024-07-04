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
)

// checks if the SnmpVacmSecurityToGroupContentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpVacmSecurityToGroupContentItem{}

// SnmpVacmSecurityToGroupContentItem struct for SnmpVacmSecurityToGroupContentItem
type SnmpVacmSecurityToGroupContentItem struct {
	// refer to group_name under access
	Group *string `json:"group,omitempty"`
	SecurityName *string `json:"security_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpVacmSecurityToGroupContentItem SnmpVacmSecurityToGroupContentItem

// NewSnmpVacmSecurityToGroupContentItem instantiates a new SnmpVacmSecurityToGroupContentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpVacmSecurityToGroupContentItem() *SnmpVacmSecurityToGroupContentItem {
	this := SnmpVacmSecurityToGroupContentItem{}
	return &this
}

// NewSnmpVacmSecurityToGroupContentItemWithDefaults instantiates a new SnmpVacmSecurityToGroupContentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpVacmSecurityToGroupContentItemWithDefaults() *SnmpVacmSecurityToGroupContentItem {
	this := SnmpVacmSecurityToGroupContentItem{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SnmpVacmSecurityToGroupContentItem) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpVacmSecurityToGroupContentItem) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SnmpVacmSecurityToGroupContentItem) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *SnmpVacmSecurityToGroupContentItem) SetGroup(v string) {
	o.Group = &v
}

// GetSecurityName returns the SecurityName field value if set, zero value otherwise.
func (o *SnmpVacmSecurityToGroupContentItem) GetSecurityName() string {
	if o == nil || IsNil(o.SecurityName) {
		var ret string
		return ret
	}
	return *o.SecurityName
}

// GetSecurityNameOk returns a tuple with the SecurityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpVacmSecurityToGroupContentItem) GetSecurityNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityName) {
		return nil, false
	}
	return o.SecurityName, true
}

// HasSecurityName returns a boolean if a field has been set.
func (o *SnmpVacmSecurityToGroupContentItem) HasSecurityName() bool {
	if o != nil && !IsNil(o.SecurityName) {
		return true
	}

	return false
}

// SetSecurityName gets a reference to the given string and assigns it to the SecurityName field.
func (o *SnmpVacmSecurityToGroupContentItem) SetSecurityName(v string) {
	o.SecurityName = &v
}

func (o SnmpVacmSecurityToGroupContentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpVacmSecurityToGroupContentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.SecurityName) {
		toSerialize["security_name"] = o.SecurityName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnmpVacmSecurityToGroupContentItem) UnmarshalJSON(data []byte) (err error) {
	varSnmpVacmSecurityToGroupContentItem := _SnmpVacmSecurityToGroupContentItem{}

	err = json.Unmarshal(data, &varSnmpVacmSecurityToGroupContentItem)

	if err != nil {
		return err
	}

	*o = SnmpVacmSecurityToGroupContentItem(varSnmpVacmSecurityToGroupContentItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "security_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpVacmSecurityToGroupContentItem struct {
	value *SnmpVacmSecurityToGroupContentItem
	isSet bool
}

func (v NullableSnmpVacmSecurityToGroupContentItem) Get() *SnmpVacmSecurityToGroupContentItem {
	return v.value
}

func (v *NullableSnmpVacmSecurityToGroupContentItem) Set(val *SnmpVacmSecurityToGroupContentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpVacmSecurityToGroupContentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpVacmSecurityToGroupContentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpVacmSecurityToGroupContentItem(val *SnmpVacmSecurityToGroupContentItem) *NullableSnmpVacmSecurityToGroupContentItem {
	return &NullableSnmpVacmSecurityToGroupContentItem{value: val, isSet: true}
}

func (v NullableSnmpVacmSecurityToGroupContentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpVacmSecurityToGroupContentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


