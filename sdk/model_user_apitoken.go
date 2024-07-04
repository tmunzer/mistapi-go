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

// checks if the UserApitoken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserApitoken{}

// UserApitoken User API Token
type UserApitoken struct {
	CreatedTime *int32 `json:"created_time,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	LastUsed NullableInt32 `json:"last_used,omitempty"`
	// name of the token
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserApitoken UserApitoken

// NewUserApitoken instantiates a new UserApitoken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserApitoken() *UserApitoken {
	this := UserApitoken{}
	return &this
}

// NewUserApitokenWithDefaults instantiates a new UserApitoken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserApitokenWithDefaults() *UserApitoken {
	this := UserApitoken{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *UserApitoken) GetCreatedTime() int32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret int32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApitoken) GetCreatedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *UserApitoken) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int32 and assigns it to the CreatedTime field.
func (o *UserApitoken) SetCreatedTime(v int32) {
	o.CreatedTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserApitoken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApitoken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserApitoken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserApitoken) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UserApitoken) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApitoken) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UserApitoken) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UserApitoken) SetKey(v string) {
	o.Key = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserApitoken) GetLastUsed() int32 {
	if o == nil || IsNil(o.LastUsed.Get()) {
		var ret int32
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserApitoken) GetLastUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// HasLastUsed returns a boolean if a field has been set.
func (o *UserApitoken) HasLastUsed() bool {
	if o != nil && o.LastUsed.IsSet() {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given NullableInt32 and assigns it to the LastUsed field.
func (o *UserApitoken) SetLastUsed(v int32) {
	o.LastUsed.Set(&v)
}
// SetLastUsedNil sets the value for LastUsed to be an explicit nil
func (o *UserApitoken) SetLastUsedNil() {
	o.LastUsed.Set(nil)
}

// UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
func (o *UserApitoken) UnsetLastUsed() {
	o.LastUsed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserApitoken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserApitoken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserApitoken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserApitoken) SetName(v string) {
	o.Name = &v
}

func (o UserApitoken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserApitoken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.LastUsed.IsSet() {
		toSerialize["last_used"] = o.LastUsed.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserApitoken) UnmarshalJSON(data []byte) (err error) {
	varUserApitoken := _UserApitoken{}

	err = json.Unmarshal(data, &varUserApitoken)

	if err != nil {
		return err
	}

	*o = UserApitoken(varUserApitoken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "last_used")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserApitoken struct {
	value *UserApitoken
	isSet bool
}

func (v NullableUserApitoken) Get() *UserApitoken {
	return v.value
}

func (v *NullableUserApitoken) Set(val *UserApitoken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserApitoken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserApitoken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserApitoken(val *UserApitoken) *NullableUserApitoken {
	return &NullableUserApitoken{value: val, isSet: true}
}

func (v NullableUserApitoken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserApitoken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


