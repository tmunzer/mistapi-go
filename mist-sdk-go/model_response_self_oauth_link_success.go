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

// checks if the ResponseSelfOauthLinkSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSelfOauthLinkSuccess{}

// ResponseSelfOauthLinkSuccess struct for ResponseSelfOauthLinkSuccess
type ResponseSelfOauthLinkSuccess struct {
	Action string `json:"action"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSelfOauthLinkSuccess ResponseSelfOauthLinkSuccess

// NewResponseSelfOauthLinkSuccess instantiates a new ResponseSelfOauthLinkSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSelfOauthLinkSuccess(action string, id string) *ResponseSelfOauthLinkSuccess {
	this := ResponseSelfOauthLinkSuccess{}
	this.Action = action
	this.Id = id
	return &this
}

// NewResponseSelfOauthLinkSuccessWithDefaults instantiates a new ResponseSelfOauthLinkSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSelfOauthLinkSuccessWithDefaults() *ResponseSelfOauthLinkSuccess {
	this := ResponseSelfOauthLinkSuccess{}
	return &this
}

// GetAction returns the Action field value
func (o *ResponseSelfOauthLinkSuccess) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ResponseSelfOauthLinkSuccess) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ResponseSelfOauthLinkSuccess) SetAction(v string) {
	o.Action = v
}

// GetId returns the Id field value
func (o *ResponseSelfOauthLinkSuccess) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseSelfOauthLinkSuccess) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseSelfOauthLinkSuccess) SetId(v string) {
	o.Id = v
}

func (o ResponseSelfOauthLinkSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSelfOauthLinkSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSelfOauthLinkSuccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"id",
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

	varResponseSelfOauthLinkSuccess := _ResponseSelfOauthLinkSuccess{}

	err = json.Unmarshal(data, &varResponseSelfOauthLinkSuccess)

	if err != nil {
		return err
	}

	*o = ResponseSelfOauthLinkSuccess(varResponseSelfOauthLinkSuccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSelfOauthLinkSuccess struct {
	value *ResponseSelfOauthLinkSuccess
	isSet bool
}

func (v NullableResponseSelfOauthLinkSuccess) Get() *ResponseSelfOauthLinkSuccess {
	return v.value
}

func (v *NullableResponseSelfOauthLinkSuccess) Set(val *ResponseSelfOauthLinkSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSelfOauthLinkSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSelfOauthLinkSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSelfOauthLinkSuccess(val *ResponseSelfOauthLinkSuccess) *NullableResponseSelfOauthLinkSuccess {
	return &NullableResponseSelfOauthLinkSuccess{value: val, isSet: true}
}

func (v NullableResponseSelfOauthLinkSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSelfOauthLinkSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


