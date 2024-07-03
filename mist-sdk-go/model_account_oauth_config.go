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

// checks if the AccountOauthConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountOauthConfig{}

// AccountOauthConfig OAuth linked apps (zoom/teams/intune) account details
type AccountOauthConfig struct {
	// Linked app(zoom/teams/intune) account id
	AccountId string `json:"account_id"`
	// Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/
	MaxDailyApiRequests *int32 `json:"max_daily_api_requests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountOauthConfig AccountOauthConfig

// NewAccountOauthConfig instantiates a new AccountOauthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOauthConfig(accountId string) *AccountOauthConfig {
	this := AccountOauthConfig{}
	this.AccountId = accountId
	return &this
}

// NewAccountOauthConfigWithDefaults instantiates a new AccountOauthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOauthConfigWithDefaults() *AccountOauthConfig {
	this := AccountOauthConfig{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountOauthConfig) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountOauthConfig) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountOauthConfig) SetAccountId(v string) {
	o.AccountId = v
}

// GetMaxDailyApiRequests returns the MaxDailyApiRequests field value if set, zero value otherwise.
func (o *AccountOauthConfig) GetMaxDailyApiRequests() int32 {
	if o == nil || IsNil(o.MaxDailyApiRequests) {
		var ret int32
		return ret
	}
	return *o.MaxDailyApiRequests
}

// GetMaxDailyApiRequestsOk returns a tuple with the MaxDailyApiRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthConfig) GetMaxDailyApiRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDailyApiRequests) {
		return nil, false
	}
	return o.MaxDailyApiRequests, true
}

// HasMaxDailyApiRequests returns a boolean if a field has been set.
func (o *AccountOauthConfig) HasMaxDailyApiRequests() bool {
	if o != nil && !IsNil(o.MaxDailyApiRequests) {
		return true
	}

	return false
}

// SetMaxDailyApiRequests gets a reference to the given int32 and assigns it to the MaxDailyApiRequests field.
func (o *AccountOauthConfig) SetMaxDailyApiRequests(v int32) {
	o.MaxDailyApiRequests = &v
}

func (o AccountOauthConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountOauthConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	if !IsNil(o.MaxDailyApiRequests) {
		toSerialize["max_daily_api_requests"] = o.MaxDailyApiRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountOauthConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
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

	varAccountOauthConfig := _AccountOauthConfig{}

	err = json.Unmarshal(data, &varAccountOauthConfig)

	if err != nil {
		return err
	}

	*o = AccountOauthConfig(varAccountOauthConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "max_daily_api_requests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountOauthConfig struct {
	value *AccountOauthConfig
	isSet bool
}

func (v NullableAccountOauthConfig) Get() *AccountOauthConfig {
	return v.value
}

func (v *NullableAccountOauthConfig) Set(val *AccountOauthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOauthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOauthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOauthConfig(val *AccountOauthConfig) *NullableAccountOauthConfig {
	return &NullableAccountOauthConfig{value: val, isSet: true}
}

func (v NullableAccountOauthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOauthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

