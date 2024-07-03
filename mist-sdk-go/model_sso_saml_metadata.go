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

// checks if the SsoSamlMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSamlMetadata{}

// SsoSamlMetadata struct for SsoSamlMetadata
type SsoSamlMetadata struct {
	AcsUrl string `json:"acs_url"`
	EntityId string `json:"entity_id"`
	LogoutUrl string `json:"logout_url"`
	MetadataXml string `json:"metadata_xml"`
	AdditionalProperties map[string]interface{}
}

type _SsoSamlMetadata SsoSamlMetadata

// NewSsoSamlMetadata instantiates a new SsoSamlMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSamlMetadata(acsUrl string, entityId string, logoutUrl string, metadataXml string) *SsoSamlMetadata {
	this := SsoSamlMetadata{}
	this.AcsUrl = acsUrl
	this.EntityId = entityId
	this.LogoutUrl = logoutUrl
	this.MetadataXml = metadataXml
	return &this
}

// NewSsoSamlMetadataWithDefaults instantiates a new SsoSamlMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSamlMetadataWithDefaults() *SsoSamlMetadata {
	this := SsoSamlMetadata{}
	return &this
}

// GetAcsUrl returns the AcsUrl field value
func (o *SsoSamlMetadata) GetAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value
// and a boolean to check if the value has been set.
func (o *SsoSamlMetadata) GetAcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcsUrl, true
}

// SetAcsUrl sets field value
func (o *SsoSamlMetadata) SetAcsUrl(v string) {
	o.AcsUrl = v
}

// GetEntityId returns the EntityId field value
func (o *SsoSamlMetadata) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *SsoSamlMetadata) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *SsoSamlMetadata) SetEntityId(v string) {
	o.EntityId = v
}

// GetLogoutUrl returns the LogoutUrl field value
func (o *SsoSamlMetadata) GetLogoutUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value
// and a boolean to check if the value has been set.
func (o *SsoSamlMetadata) GetLogoutUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoutUrl, true
}

// SetLogoutUrl sets field value
func (o *SsoSamlMetadata) SetLogoutUrl(v string) {
	o.LogoutUrl = v
}

// GetMetadataXml returns the MetadataXml field value
func (o *SsoSamlMetadata) GetMetadataXml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataXml
}

// GetMetadataXmlOk returns a tuple with the MetadataXml field value
// and a boolean to check if the value has been set.
func (o *SsoSamlMetadata) GetMetadataXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataXml, true
}

// SetMetadataXml sets field value
func (o *SsoSamlMetadata) SetMetadataXml(v string) {
	o.MetadataXml = v
}

func (o SsoSamlMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSamlMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acs_url"] = o.AcsUrl
	toSerialize["entity_id"] = o.EntityId
	toSerialize["logout_url"] = o.LogoutUrl
	toSerialize["metadata_xml"] = o.MetadataXml

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsoSamlMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"acs_url",
		"entity_id",
		"logout_url",
		"metadata_xml",
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

	varSsoSamlMetadata := _SsoSamlMetadata{}

	err = json.Unmarshal(data, &varSsoSamlMetadata)

	if err != nil {
		return err
	}

	*o = SsoSamlMetadata(varSsoSamlMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acs_url")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "logout_url")
		delete(additionalProperties, "metadata_xml")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsoSamlMetadata struct {
	value *SsoSamlMetadata
	isSet bool
}

func (v NullableSsoSamlMetadata) Get() *SsoSamlMetadata {
	return v.value
}

func (v *NullableSsoSamlMetadata) Set(val *SsoSamlMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSamlMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSamlMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSamlMetadata(val *SsoSamlMetadata) *NullableSsoSamlMetadata {
	return &NullableSsoSamlMetadata{value: val, isSet: true}
}

func (v NullableSsoSamlMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSamlMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


