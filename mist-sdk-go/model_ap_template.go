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

// checks if the ApTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApTemplate{}

// ApTemplate struct for ApTemplate
type ApTemplate struct {
	ApMatching ApTemplateMatching `json:"ap_matching"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	Wifi *ApTemplateWifi `json:"wifi,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApTemplate ApTemplate

// NewApTemplate instantiates a new ApTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApTemplate(apMatching ApTemplateMatching) *ApTemplate {
	this := ApTemplate{}
	this.ApMatching = apMatching
	return &this
}

// NewApTemplateWithDefaults instantiates a new ApTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApTemplateWithDefaults() *ApTemplate {
	this := ApTemplate{}
	return &this
}

// GetApMatching returns the ApMatching field value
func (o *ApTemplate) GetApMatching() ApTemplateMatching {
	if o == nil {
		var ret ApTemplateMatching
		return ret
	}

	return o.ApMatching
}

// GetApMatchingOk returns a tuple with the ApMatching field value
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetApMatchingOk() (*ApTemplateMatching, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApMatching, true
}

// SetApMatching sets field value
func (o *ApTemplate) SetApMatching(v ApTemplateMatching) {
	o.ApMatching = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ApTemplate) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ApTemplate) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *ApTemplate) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *ApTemplate) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *ApTemplate) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *ApTemplate) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApTemplate) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *ApTemplate) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *ApTemplate) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *ApTemplate) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ApTemplate) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ApTemplate) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ApTemplate) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ApTemplate) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ApTemplate) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *ApTemplate) SetSiteId(v string) {
	o.SiteId = &v
}

// GetWifi returns the Wifi field value if set, zero value otherwise.
func (o *ApTemplate) GetWifi() ApTemplateWifi {
	if o == nil || IsNil(o.Wifi) {
		var ret ApTemplateWifi
		return ret
	}
	return *o.Wifi
}

// GetWifiOk returns a tuple with the Wifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplate) GetWifiOk() (*ApTemplateWifi, bool) {
	if o == nil || IsNil(o.Wifi) {
		return nil, false
	}
	return o.Wifi, true
}

// HasWifi returns a boolean if a field has been set.
func (o *ApTemplate) HasWifi() bool {
	if o != nil && !IsNil(o.Wifi) {
		return true
	}

	return false
}

// SetWifi gets a reference to the given ApTemplateWifi and assigns it to the Wifi field.
func (o *ApTemplate) SetWifi(v ApTemplateWifi) {
	o.Wifi = &v
}

func (o ApTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ap_matching"] = o.ApMatching
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Wifi) {
		toSerialize["wifi"] = o.Wifi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ap_matching",
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

	varApTemplate := _ApTemplate{}

	err = json.Unmarshal(data, &varApTemplate)

	if err != nil {
		return err
	}

	*o = ApTemplate(varApTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_matching")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "wifi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApTemplate struct {
	value *ApTemplate
	isSet bool
}

func (v NullableApTemplate) Get() *ApTemplate {
	return v.value
}

func (v *NullableApTemplate) Set(val *ApTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableApTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableApTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApTemplate(val *ApTemplate) *NullableApTemplate {
	return &NullableApTemplate{value: val, isSet: true}
}

func (v NullableApTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


