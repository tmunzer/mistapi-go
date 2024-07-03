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

// checks if the ResponseVerifyTokenSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseVerifyTokenSuccess{}

// ResponseVerifyTokenSuccess struct for ResponseVerifyTokenSuccess
type ResponseVerifyTokenSuccess struct {
	Detail *string `json:"detail,omitempty"`
	InviteNotApplied *bool `json:"invite_not_applied,omitempty"`
	MinLength *int32 `json:"min_length,omitempty"`
	ReturnTo *string `json:"return_to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseVerifyTokenSuccess ResponseVerifyTokenSuccess

// NewResponseVerifyTokenSuccess instantiates a new ResponseVerifyTokenSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseVerifyTokenSuccess() *ResponseVerifyTokenSuccess {
	this := ResponseVerifyTokenSuccess{}
	return &this
}

// NewResponseVerifyTokenSuccessWithDefaults instantiates a new ResponseVerifyTokenSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseVerifyTokenSuccessWithDefaults() *ResponseVerifyTokenSuccess {
	this := ResponseVerifyTokenSuccess{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseVerifyTokenSuccess) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseVerifyTokenSuccess) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseVerifyTokenSuccess) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseVerifyTokenSuccess) SetDetail(v string) {
	o.Detail = &v
}

// GetInviteNotApplied returns the InviteNotApplied field value if set, zero value otherwise.
func (o *ResponseVerifyTokenSuccess) GetInviteNotApplied() bool {
	if o == nil || IsNil(o.InviteNotApplied) {
		var ret bool
		return ret
	}
	return *o.InviteNotApplied
}

// GetInviteNotAppliedOk returns a tuple with the InviteNotApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseVerifyTokenSuccess) GetInviteNotAppliedOk() (*bool, bool) {
	if o == nil || IsNil(o.InviteNotApplied) {
		return nil, false
	}
	return o.InviteNotApplied, true
}

// HasInviteNotApplied returns a boolean if a field has been set.
func (o *ResponseVerifyTokenSuccess) HasInviteNotApplied() bool {
	if o != nil && !IsNil(o.InviteNotApplied) {
		return true
	}

	return false
}

// SetInviteNotApplied gets a reference to the given bool and assigns it to the InviteNotApplied field.
func (o *ResponseVerifyTokenSuccess) SetInviteNotApplied(v bool) {
	o.InviteNotApplied = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *ResponseVerifyTokenSuccess) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseVerifyTokenSuccess) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *ResponseVerifyTokenSuccess) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *ResponseVerifyTokenSuccess) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *ResponseVerifyTokenSuccess) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseVerifyTokenSuccess) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *ResponseVerifyTokenSuccess) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *ResponseVerifyTokenSuccess) SetReturnTo(v string) {
	o.ReturnTo = &v
}

func (o ResponseVerifyTokenSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseVerifyTokenSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.InviteNotApplied) {
		toSerialize["invite_not_applied"] = o.InviteNotApplied
	}
	if !IsNil(o.MinLength) {
		toSerialize["min_length"] = o.MinLength
	}
	if !IsNil(o.ReturnTo) {
		toSerialize["return_to"] = o.ReturnTo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseVerifyTokenSuccess) UnmarshalJSON(data []byte) (err error) {
	varResponseVerifyTokenSuccess := _ResponseVerifyTokenSuccess{}

	err = json.Unmarshal(data, &varResponseVerifyTokenSuccess)

	if err != nil {
		return err
	}

	*o = ResponseVerifyTokenSuccess(varResponseVerifyTokenSuccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "invite_not_applied")
		delete(additionalProperties, "min_length")
		delete(additionalProperties, "return_to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseVerifyTokenSuccess struct {
	value *ResponseVerifyTokenSuccess
	isSet bool
}

func (v NullableResponseVerifyTokenSuccess) Get() *ResponseVerifyTokenSuccess {
	return v.value
}

func (v *NullableResponseVerifyTokenSuccess) Set(val *ResponseVerifyTokenSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseVerifyTokenSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseVerifyTokenSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseVerifyTokenSuccess(val *ResponseVerifyTokenSuccess) *NullableResponseVerifyTokenSuccess {
	return &NullableResponseVerifyTokenSuccess{value: val, isSet: true}
}

func (v NullableResponseVerifyTokenSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseVerifyTokenSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


