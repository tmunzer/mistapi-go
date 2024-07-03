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

// checks if the ConstLicenseType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstLicenseType{}

// ConstLicenseType struct for ConstLicenseType
type ConstLicenseType struct {
	Description *string `json:"description,omitempty"`
	Includes []string `json:"includes,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstLicenseType ConstLicenseType

// NewConstLicenseType instantiates a new ConstLicenseType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstLicenseType() *ConstLicenseType {
	this := ConstLicenseType{}
	return &this
}

// NewConstLicenseTypeWithDefaults instantiates a new ConstLicenseType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstLicenseTypeWithDefaults() *ConstLicenseType {
	this := ConstLicenseType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConstLicenseType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstLicenseType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConstLicenseType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConstLicenseType) SetDescription(v string) {
	o.Description = &v
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *ConstLicenseType) GetIncludes() []string {
	if o == nil || IsNil(o.Includes) {
		var ret []string
		return ret
	}
	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstLicenseType) GetIncludesOk() ([]string, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *ConstLicenseType) HasIncludes() bool {
	if o != nil && !IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given []string and assigns it to the Includes field.
func (o *ConstLicenseType) SetIncludes(v []string) {
	o.Includes = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ConstLicenseType) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstLicenseType) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ConstLicenseType) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ConstLicenseType) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConstLicenseType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstLicenseType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConstLicenseType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConstLicenseType) SetName(v string) {
	o.Name = &v
}

func (o ConstLicenseType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstLicenseType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Includes) {
		toSerialize["includes"] = o.Includes
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstLicenseType) UnmarshalJSON(data []byte) (err error) {
	varConstLicenseType := _ConstLicenseType{}

	err = json.Unmarshal(data, &varConstLicenseType)

	if err != nil {
		return err
	}

	*o = ConstLicenseType(varConstLicenseType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "includes")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstLicenseType struct {
	value *ConstLicenseType
	isSet bool
}

func (v NullableConstLicenseType) Get() *ConstLicenseType {
	return v.value
}

func (v *NullableConstLicenseType) Set(val *ConstLicenseType) {
	v.value = val
	v.isSet = true
}

func (v NullableConstLicenseType) IsSet() bool {
	return v.isSet
}

func (v *NullableConstLicenseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstLicenseType(val *ConstLicenseType) *NullableConstLicenseType {
	return &NullableConstLicenseType{value: val, isSet: true}
}

func (v NullableConstLicenseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstLicenseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


