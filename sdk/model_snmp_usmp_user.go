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

// checks if the SnmpUsmpUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpUsmpUser{}

// SnmpUsmpUser struct for SnmpUsmpUser
type SnmpUsmpUser struct {
	// Not required if `authentication_type`==`authentication_none` include alphabetic, numeric, and special characters, but it cannot include control characters.
	AuthenticationPassword *string `json:"authentication_password,omitempty"`
	AuthenticationType *SnmpUsmpUserAuthenticationType `json:"authentication_type,omitempty"`
	// Not required if `encryption_type`==`privacy-none` include alphabetic, numeric, and special characters, but it cannot include control characters
	EncryptionPassword *string `json:"encryption_password,omitempty"`
	EncryptionType *SnmpUsmpUserEncryptionType `json:"encryption_type,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpUsmpUser SnmpUsmpUser

// NewSnmpUsmpUser instantiates a new SnmpUsmpUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpUsmpUser() *SnmpUsmpUser {
	this := SnmpUsmpUser{}
	return &this
}

// NewSnmpUsmpUserWithDefaults instantiates a new SnmpUsmpUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpUsmpUserWithDefaults() *SnmpUsmpUser {
	this := SnmpUsmpUser{}
	return &this
}

// GetAuthenticationPassword returns the AuthenticationPassword field value if set, zero value otherwise.
func (o *SnmpUsmpUser) GetAuthenticationPassword() string {
	if o == nil || IsNil(o.AuthenticationPassword) {
		var ret string
		return ret
	}
	return *o.AuthenticationPassword
}

// GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsmpUser) GetAuthenticationPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationPassword) {
		return nil, false
	}
	return o.AuthenticationPassword, true
}

// HasAuthenticationPassword returns a boolean if a field has been set.
func (o *SnmpUsmpUser) HasAuthenticationPassword() bool {
	if o != nil && !IsNil(o.AuthenticationPassword) {
		return true
	}

	return false
}

// SetAuthenticationPassword gets a reference to the given string and assigns it to the AuthenticationPassword field.
func (o *SnmpUsmpUser) SetAuthenticationPassword(v string) {
	o.AuthenticationPassword = &v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *SnmpUsmpUser) GetAuthenticationType() SnmpUsmpUserAuthenticationType {
	if o == nil || IsNil(o.AuthenticationType) {
		var ret SnmpUsmpUserAuthenticationType
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsmpUser) GetAuthenticationTypeOk() (*SnmpUsmpUserAuthenticationType, bool) {
	if o == nil || IsNil(o.AuthenticationType) {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *SnmpUsmpUser) HasAuthenticationType() bool {
	if o != nil && !IsNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given SnmpUsmpUserAuthenticationType and assigns it to the AuthenticationType field.
func (o *SnmpUsmpUser) SetAuthenticationType(v SnmpUsmpUserAuthenticationType) {
	o.AuthenticationType = &v
}

// GetEncryptionPassword returns the EncryptionPassword field value if set, zero value otherwise.
func (o *SnmpUsmpUser) GetEncryptionPassword() string {
	if o == nil || IsNil(o.EncryptionPassword) {
		var ret string
		return ret
	}
	return *o.EncryptionPassword
}

// GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsmpUser) GetEncryptionPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionPassword) {
		return nil, false
	}
	return o.EncryptionPassword, true
}

// HasEncryptionPassword returns a boolean if a field has been set.
func (o *SnmpUsmpUser) HasEncryptionPassword() bool {
	if o != nil && !IsNil(o.EncryptionPassword) {
		return true
	}

	return false
}

// SetEncryptionPassword gets a reference to the given string and assigns it to the EncryptionPassword field.
func (o *SnmpUsmpUser) SetEncryptionPassword(v string) {
	o.EncryptionPassword = &v
}

// GetEncryptionType returns the EncryptionType field value if set, zero value otherwise.
func (o *SnmpUsmpUser) GetEncryptionType() SnmpUsmpUserEncryptionType {
	if o == nil || IsNil(o.EncryptionType) {
		var ret SnmpUsmpUserEncryptionType
		return ret
	}
	return *o.EncryptionType
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsmpUser) GetEncryptionTypeOk() (*SnmpUsmpUserEncryptionType, bool) {
	if o == nil || IsNil(o.EncryptionType) {
		return nil, false
	}
	return o.EncryptionType, true
}

// HasEncryptionType returns a boolean if a field has been set.
func (o *SnmpUsmpUser) HasEncryptionType() bool {
	if o != nil && !IsNil(o.EncryptionType) {
		return true
	}

	return false
}

// SetEncryptionType gets a reference to the given SnmpUsmpUserEncryptionType and assigns it to the EncryptionType field.
func (o *SnmpUsmpUser) SetEncryptionType(v SnmpUsmpUserEncryptionType) {
	o.EncryptionType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnmpUsmpUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUsmpUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SnmpUsmpUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnmpUsmpUser) SetName(v string) {
	o.Name = &v
}

func (o SnmpUsmpUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpUsmpUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationPassword) {
		toSerialize["authentication_password"] = o.AuthenticationPassword
	}
	if !IsNil(o.AuthenticationType) {
		toSerialize["authentication_type"] = o.AuthenticationType
	}
	if !IsNil(o.EncryptionPassword) {
		toSerialize["encryption_password"] = o.EncryptionPassword
	}
	if !IsNil(o.EncryptionType) {
		toSerialize["encryption_type"] = o.EncryptionType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnmpUsmpUser) UnmarshalJSON(data []byte) (err error) {
	varSnmpUsmpUser := _SnmpUsmpUser{}

	err = json.Unmarshal(data, &varSnmpUsmpUser)

	if err != nil {
		return err
	}

	*o = SnmpUsmpUser(varSnmpUsmpUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authentication_password")
		delete(additionalProperties, "authentication_type")
		delete(additionalProperties, "encryption_password")
		delete(additionalProperties, "encryption_type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpUsmpUser struct {
	value *SnmpUsmpUser
	isSet bool
}

func (v NullableSnmpUsmpUser) Get() *SnmpUsmpUser {
	return v.value
}

func (v *NullableSnmpUsmpUser) Set(val *SnmpUsmpUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpUsmpUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpUsmpUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpUsmpUser(val *SnmpUsmpUser) *NullableSnmpUsmpUser {
	return &NullableSnmpUsmpUser{value: val, isSet: true}
}

func (v NullableSnmpUsmpUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpUsmpUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

