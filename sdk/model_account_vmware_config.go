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

// checks if the AccountVmwareConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountVmwareConfig{}

// AccountVmwareConfig struct for AccountVmwareConfig
type AccountVmwareConfig struct {
	// customer account Client ID
	ClientId string `json:"client_id"`
	// customer account Client Secret
	ClientSecret string `json:"client_secret"`
	// customer account VMware instance URL
	InstanceUrl string `json:"instance_url"`
	AdditionalProperties map[string]interface{}
}

type _AccountVmwareConfig AccountVmwareConfig

// NewAccountVmwareConfig instantiates a new AccountVmwareConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountVmwareConfig(clientId string, clientSecret string, instanceUrl string) *AccountVmwareConfig {
	this := AccountVmwareConfig{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.InstanceUrl = instanceUrl
	return &this
}

// NewAccountVmwareConfigWithDefaults instantiates a new AccountVmwareConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountVmwareConfigWithDefaults() *AccountVmwareConfig {
	this := AccountVmwareConfig{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *AccountVmwareConfig) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AccountVmwareConfig) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AccountVmwareConfig) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *AccountVmwareConfig) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AccountVmwareConfig) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *AccountVmwareConfig) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetInstanceUrl returns the InstanceUrl field value
func (o *AccountVmwareConfig) GetInstanceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value
// and a boolean to check if the value has been set.
func (o *AccountVmwareConfig) GetInstanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceUrl, true
}

// SetInstanceUrl sets field value
func (o *AccountVmwareConfig) SetInstanceUrl(v string) {
	o.InstanceUrl = v
}

func (o AccountVmwareConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountVmwareConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["instance_url"] = o.InstanceUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountVmwareConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
		"client_secret",
		"instance_url",
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

	varAccountVmwareConfig := _AccountVmwareConfig{}

	err = json.Unmarshal(data, &varAccountVmwareConfig)

	if err != nil {
		return err
	}

	*o = AccountVmwareConfig(varAccountVmwareConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "instance_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountVmwareConfig struct {
	value *AccountVmwareConfig
	isSet bool
}

func (v NullableAccountVmwareConfig) Get() *AccountVmwareConfig {
	return v.value
}

func (v *NullableAccountVmwareConfig) Set(val *AccountVmwareConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountVmwareConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountVmwareConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountVmwareConfig(val *AccountVmwareConfig) *NullableAccountVmwareConfig {
	return &NullableAccountVmwareConfig{value: val, isSet: true}
}

func (v NullableAccountVmwareConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountVmwareConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

