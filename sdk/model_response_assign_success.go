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

// checks if the ResponseAssignSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAssignSuccess{}

// ResponseAssignSuccess struct for ResponseAssignSuccess
type ResponseAssignSuccess struct {
	Success []string `json:"success"`
	AdditionalProperties map[string]interface{}
}

type _ResponseAssignSuccess ResponseAssignSuccess

// NewResponseAssignSuccess instantiates a new ResponseAssignSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAssignSuccess(success []string) *ResponseAssignSuccess {
	this := ResponseAssignSuccess{}
	this.Success = success
	return &this
}

// NewResponseAssignSuccessWithDefaults instantiates a new ResponseAssignSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAssignSuccessWithDefaults() *ResponseAssignSuccess {
	this := ResponseAssignSuccess{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ResponseAssignSuccess) GetSuccess() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ResponseAssignSuccess) GetSuccessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success, true
}

// SetSuccess sets field value
func (o *ResponseAssignSuccess) SetSuccess(v []string) {
	o.Success = v
}

func (o ResponseAssignSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAssignSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseAssignSuccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varResponseAssignSuccess := _ResponseAssignSuccess{}

	err = json.Unmarshal(data, &varResponseAssignSuccess)

	if err != nil {
		return err
	}

	*o = ResponseAssignSuccess(varResponseAssignSuccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseAssignSuccess struct {
	value *ResponseAssignSuccess
	isSet bool
}

func (v NullableResponseAssignSuccess) Get() *ResponseAssignSuccess {
	return v.value
}

func (v *NullableResponseAssignSuccess) Set(val *ResponseAssignSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAssignSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAssignSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAssignSuccess(val *ResponseAssignSuccess) *NullableResponseAssignSuccess {
	return &NullableResponseAssignSuccess{value: val, isSet: true}
}

func (v NullableResponseAssignSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAssignSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

