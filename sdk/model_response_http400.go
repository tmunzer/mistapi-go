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

// checks if the ResponseHttp400 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseHttp400{}

// ResponseHttp400 struct for ResponseHttp400
type ResponseHttp400 struct {
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseHttp400 ResponseHttp400

// NewResponseHttp400 instantiates a new ResponseHttp400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseHttp400() *ResponseHttp400 {
	this := ResponseHttp400{}
	return &this
}

// NewResponseHttp400WithDefaults instantiates a new ResponseHttp400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseHttp400WithDefaults() *ResponseHttp400 {
	this := ResponseHttp400{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseHttp400) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseHttp400) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseHttp400) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseHttp400) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseHttp400) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseHttp400) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseHttp400) UnmarshalJSON(data []byte) (err error) {
	varResponseHttp400 := _ResponseHttp400{}

	err = json.Unmarshal(data, &varResponseHttp400)

	if err != nil {
		return err
	}

	*o = ResponseHttp400(varResponseHttp400)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseHttp400 struct {
	value *ResponseHttp400
	isSet bool
}

func (v NullableResponseHttp400) Get() *ResponseHttp400 {
	return v.value
}

func (v *NullableResponseHttp400) Set(val *ResponseHttp400) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseHttp400) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseHttp400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseHttp400(val *ResponseHttp400) *NullableResponseHttp400 {
	return &NullableResponseHttp400{value: val, isSet: true}
}

func (v NullableResponseHttp400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseHttp400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


