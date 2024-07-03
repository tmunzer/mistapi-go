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

// checks if the SiteSettingPaloaltoNetworkGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingPaloaltoNetworkGateway{}

// SiteSettingPaloaltoNetworkGateway struct for SiteSettingPaloaltoNetworkGateway
type SiteSettingPaloaltoNetworkGateway struct {
	ApiKey *string `json:"api_key,omitempty"`
	ApiUrl *string `json:"api_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingPaloaltoNetworkGateway SiteSettingPaloaltoNetworkGateway

// NewSiteSettingPaloaltoNetworkGateway instantiates a new SiteSettingPaloaltoNetworkGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingPaloaltoNetworkGateway() *SiteSettingPaloaltoNetworkGateway {
	this := SiteSettingPaloaltoNetworkGateway{}
	return &this
}

// NewSiteSettingPaloaltoNetworkGatewayWithDefaults instantiates a new SiteSettingPaloaltoNetworkGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingPaloaltoNetworkGatewayWithDefaults() *SiteSettingPaloaltoNetworkGateway {
	this := SiteSettingPaloaltoNetworkGateway{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SiteSettingPaloaltoNetworkGateway) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingPaloaltoNetworkGateway) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SiteSettingPaloaltoNetworkGateway) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SiteSettingPaloaltoNetworkGateway) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *SiteSettingPaloaltoNetworkGateway) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl) {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingPaloaltoNetworkGateway) GetApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApiUrl) {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *SiteSettingPaloaltoNetworkGateway) HasApiUrl() bool {
	if o != nil && !IsNil(o.ApiUrl) {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *SiteSettingPaloaltoNetworkGateway) SetApiUrl(v string) {
	o.ApiUrl = &v
}

func (o SiteSettingPaloaltoNetworkGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingPaloaltoNetworkGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.ApiUrl) {
		toSerialize["api_url"] = o.ApiUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingPaloaltoNetworkGateway) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingPaloaltoNetworkGateway := _SiteSettingPaloaltoNetworkGateway{}

	err = json.Unmarshal(data, &varSiteSettingPaloaltoNetworkGateway)

	if err != nil {
		return err
	}

	*o = SiteSettingPaloaltoNetworkGateway(varSiteSettingPaloaltoNetworkGateway)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "api_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingPaloaltoNetworkGateway struct {
	value *SiteSettingPaloaltoNetworkGateway
	isSet bool
}

func (v NullableSiteSettingPaloaltoNetworkGateway) Get() *SiteSettingPaloaltoNetworkGateway {
	return v.value
}

func (v *NullableSiteSettingPaloaltoNetworkGateway) Set(val *SiteSettingPaloaltoNetworkGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingPaloaltoNetworkGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingPaloaltoNetworkGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingPaloaltoNetworkGateway(val *SiteSettingPaloaltoNetworkGateway) *NullableSiteSettingPaloaltoNetworkGateway {
	return &NullableSiteSettingPaloaltoNetworkGateway{value: val, isSet: true}
}

func (v NullableSiteSettingPaloaltoNetworkGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingPaloaltoNetworkGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


