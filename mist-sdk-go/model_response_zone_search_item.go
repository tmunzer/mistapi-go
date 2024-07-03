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

// checks if the ResponseZoneSearchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseZoneSearchItem{}

// ResponseZoneSearchItem struct for ResponseZoneSearchItem
type ResponseZoneSearchItem struct {
	Enter *int32 `json:"enter,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	User *string `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseZoneSearchItem ResponseZoneSearchItem

// NewResponseZoneSearchItem instantiates a new ResponseZoneSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseZoneSearchItem() *ResponseZoneSearchItem {
	this := ResponseZoneSearchItem{}
	return &this
}

// NewResponseZoneSearchItemWithDefaults instantiates a new ResponseZoneSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseZoneSearchItemWithDefaults() *ResponseZoneSearchItem {
	this := ResponseZoneSearchItem{}
	return &this
}

// GetEnter returns the Enter field value if set, zero value otherwise.
func (o *ResponseZoneSearchItem) GetEnter() int32 {
	if o == nil || IsNil(o.Enter) {
		var ret int32
		return ret
	}
	return *o.Enter
}

// GetEnterOk returns a tuple with the Enter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseZoneSearchItem) GetEnterOk() (*int32, bool) {
	if o == nil || IsNil(o.Enter) {
		return nil, false
	}
	return o.Enter, true
}

// HasEnter returns a boolean if a field has been set.
func (o *ResponseZoneSearchItem) HasEnter() bool {
	if o != nil && !IsNil(o.Enter) {
		return true
	}

	return false
}

// SetEnter gets a reference to the given int32 and assigns it to the Enter field.
func (o *ResponseZoneSearchItem) SetEnter(v int32) {
	o.Enter = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ResponseZoneSearchItem) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseZoneSearchItem) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ResponseZoneSearchItem) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ResponseZoneSearchItem) SetScope(v string) {
	o.Scope = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ResponseZoneSearchItem) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseZoneSearchItem) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ResponseZoneSearchItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *ResponseZoneSearchItem) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ResponseZoneSearchItem) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseZoneSearchItem) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ResponseZoneSearchItem) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ResponseZoneSearchItem) SetUser(v string) {
	o.User = &v
}

func (o ResponseZoneSearchItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseZoneSearchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enter) {
		toSerialize["enter"] = o.Enter
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseZoneSearchItem) UnmarshalJSON(data []byte) (err error) {
	varResponseZoneSearchItem := _ResponseZoneSearchItem{}

	err = json.Unmarshal(data, &varResponseZoneSearchItem)

	if err != nil {
		return err
	}

	*o = ResponseZoneSearchItem(varResponseZoneSearchItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enter")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseZoneSearchItem struct {
	value *ResponseZoneSearchItem
	isSet bool
}

func (v NullableResponseZoneSearchItem) Get() *ResponseZoneSearchItem {
	return v.value
}

func (v *NullableResponseZoneSearchItem) Set(val *ResponseZoneSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseZoneSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseZoneSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseZoneSearchItem(val *ResponseZoneSearchItem) *NullableResponseZoneSearchItem {
	return &NullableResponseZoneSearchItem{value: val, isSet: true}
}

func (v NullableResponseZoneSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseZoneSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


