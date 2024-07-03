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

// checks if the ResponseOauthAppLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseOauthAppLink{}

// ResponseOauthAppLink struct for ResponseOauthAppLink
type ResponseOauthAppLink struct {
	// List of linked account details
	Accounts []ResponseOauthAppLinkItem `json:"accounts,omitempty"`
	// Basic Auth application linked status in mist portal enabled for VMware
	Linked *bool `json:"linked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseOauthAppLink ResponseOauthAppLink

// NewResponseOauthAppLink instantiates a new ResponseOauthAppLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseOauthAppLink() *ResponseOauthAppLink {
	this := ResponseOauthAppLink{}
	return &this
}

// NewResponseOauthAppLinkWithDefaults instantiates a new ResponseOauthAppLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseOauthAppLinkWithDefaults() *ResponseOauthAppLink {
	this := ResponseOauthAppLink{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ResponseOauthAppLink) GetAccounts() []ResponseOauthAppLinkItem {
	if o == nil || IsNil(o.Accounts) {
		var ret []ResponseOauthAppLinkItem
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOauthAppLink) GetAccountsOk() ([]ResponseOauthAppLinkItem, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ResponseOauthAppLink) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []ResponseOauthAppLinkItem and assigns it to the Accounts field.
func (o *ResponseOauthAppLink) SetAccounts(v []ResponseOauthAppLinkItem) {
	o.Accounts = v
}

// GetLinked returns the Linked field value if set, zero value otherwise.
func (o *ResponseOauthAppLink) GetLinked() bool {
	if o == nil || IsNil(o.Linked) {
		var ret bool
		return ret
	}
	return *o.Linked
}

// GetLinkedOk returns a tuple with the Linked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOauthAppLink) GetLinkedOk() (*bool, bool) {
	if o == nil || IsNil(o.Linked) {
		return nil, false
	}
	return o.Linked, true
}

// HasLinked returns a boolean if a field has been set.
func (o *ResponseOauthAppLink) HasLinked() bool {
	if o != nil && !IsNil(o.Linked) {
		return true
	}

	return false
}

// SetLinked gets a reference to the given bool and assigns it to the Linked field.
func (o *ResponseOauthAppLink) SetLinked(v bool) {
	o.Linked = &v
}

func (o ResponseOauthAppLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseOauthAppLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Linked) {
		toSerialize["linked"] = o.Linked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseOauthAppLink) UnmarshalJSON(data []byte) (err error) {
	varResponseOauthAppLink := _ResponseOauthAppLink{}

	err = json.Unmarshal(data, &varResponseOauthAppLink)

	if err != nil {
		return err
	}

	*o = ResponseOauthAppLink(varResponseOauthAppLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "linked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseOauthAppLink struct {
	value *ResponseOauthAppLink
	isSet bool
}

func (v NullableResponseOauthAppLink) Get() *ResponseOauthAppLink {
	return v.value
}

func (v *NullableResponseOauthAppLink) Set(val *ResponseOauthAppLink) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseOauthAppLink) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseOauthAppLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseOauthAppLink(val *ResponseOauthAppLink) *NullableResponseOauthAppLink {
	return &NullableResponseOauthAppLink{value: val, isSet: true}
}

func (v NullableResponseOauthAppLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseOauthAppLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

