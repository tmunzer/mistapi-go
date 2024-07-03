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

// checks if the SyntheticTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticTest{}

// SyntheticTest struct for SyntheticTest
type SyntheticTest struct {
	Email *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyntheticTest SyntheticTest

// NewSyntheticTest instantiates a new SyntheticTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTest() *SyntheticTest {
	this := SyntheticTest{}
	return &this
}

// NewSyntheticTestWithDefaults instantiates a new SyntheticTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestWithDefaults() *SyntheticTest {
	this := SyntheticTest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SyntheticTest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SyntheticTest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SyntheticTest) SetEmail(v string) {
	o.Email = &v
}

func (o SyntheticTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SyntheticTest) UnmarshalJSON(data []byte) (err error) {
	varSyntheticTest := _SyntheticTest{}

	err = json.Unmarshal(data, &varSyntheticTest)

	if err != nil {
		return err
	}

	*o = SyntheticTest(varSyntheticTest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyntheticTest struct {
	value *SyntheticTest
	isSet bool
}

func (v NullableSyntheticTest) Get() *SyntheticTest {
	return v.value
}

func (v *NullableSyntheticTest) Set(val *SyntheticTest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTest(val *SyntheticTest) *NullableSyntheticTest {
	return &NullableSyntheticTest{value: val, isSet: true}
}

func (v NullableSyntheticTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

