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

// checks if the SiteSettingGatewayMgmtAutoSignatureUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingGatewayMgmtAutoSignatureUpdate{}

// SiteSettingGatewayMgmtAutoSignatureUpdate struct for SiteSettingGatewayMgmtAutoSignatureUpdate
type SiteSettingGatewayMgmtAutoSignatureUpdate struct {
	DayOfWeek *DayOfWeek `json:"day_of_week,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	// optional, Mist will decide the timing
	TimeOfDay *string `json:"time_of_day,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingGatewayMgmtAutoSignatureUpdate SiteSettingGatewayMgmtAutoSignatureUpdate

// NewSiteSettingGatewayMgmtAutoSignatureUpdate instantiates a new SiteSettingGatewayMgmtAutoSignatureUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingGatewayMgmtAutoSignatureUpdate() *SiteSettingGatewayMgmtAutoSignatureUpdate {
	this := SiteSettingGatewayMgmtAutoSignatureUpdate{}
	var enable bool = true
	this.Enable = &enable
	return &this
}

// NewSiteSettingGatewayMgmtAutoSignatureUpdateWithDefaults instantiates a new SiteSettingGatewayMgmtAutoSignatureUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingGatewayMgmtAutoSignatureUpdateWithDefaults() *SiteSettingGatewayMgmtAutoSignatureUpdate {
	this := SiteSettingGatewayMgmtAutoSignatureUpdate{}
	var enable bool = true
	this.Enable = &enable
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetDayOfWeek() DayOfWeek {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret DayOfWeek
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetDayOfWeekOk() (*DayOfWeek, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given DayOfWeek and assigns it to the DayOfWeek field.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) SetDayOfWeek(v DayOfWeek) {
	o.DayOfWeek = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) SetEnable(v bool) {
	o.Enable = &v
}

// GetTimeOfDay returns the TimeOfDay field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetTimeOfDay() string {
	if o == nil || IsNil(o.TimeOfDay) {
		var ret string
		return ret
	}
	return *o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetTimeOfDayOk() (*string, bool) {
	if o == nil || IsNil(o.TimeOfDay) {
		return nil, false
	}
	return o.TimeOfDay, true
}

// HasTimeOfDay returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) HasTimeOfDay() bool {
	if o != nil && !IsNil(o.TimeOfDay) {
		return true
	}

	return false
}

// SetTimeOfDay gets a reference to the given string and assigns it to the TimeOfDay field.
func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) SetTimeOfDay(v string) {
	o.TimeOfDay = &v
}

func (o SiteSettingGatewayMgmtAutoSignatureUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingGatewayMgmtAutoSignatureUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DayOfWeek) {
		toSerialize["day_of_week"] = o.DayOfWeek
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.TimeOfDay) {
		toSerialize["time_of_day"] = o.TimeOfDay
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingGatewayMgmtAutoSignatureUpdate := _SiteSettingGatewayMgmtAutoSignatureUpdate{}

	err = json.Unmarshal(data, &varSiteSettingGatewayMgmtAutoSignatureUpdate)

	if err != nil {
		return err
	}

	*o = SiteSettingGatewayMgmtAutoSignatureUpdate(varSiteSettingGatewayMgmtAutoSignatureUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "day_of_week")
		delete(additionalProperties, "enable")
		delete(additionalProperties, "time_of_day")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingGatewayMgmtAutoSignatureUpdate struct {
	value *SiteSettingGatewayMgmtAutoSignatureUpdate
	isSet bool
}

func (v NullableSiteSettingGatewayMgmtAutoSignatureUpdate) Get() *SiteSettingGatewayMgmtAutoSignatureUpdate {
	return v.value
}

func (v *NullableSiteSettingGatewayMgmtAutoSignatureUpdate) Set(val *SiteSettingGatewayMgmtAutoSignatureUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingGatewayMgmtAutoSignatureUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingGatewayMgmtAutoSignatureUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingGatewayMgmtAutoSignatureUpdate(val *SiteSettingGatewayMgmtAutoSignatureUpdate) *NullableSiteSettingGatewayMgmtAutoSignatureUpdate {
	return &NullableSiteSettingGatewayMgmtAutoSignatureUpdate{value: val, isSet: true}
}

func (v NullableSiteSettingGatewayMgmtAutoSignatureUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingGatewayMgmtAutoSignatureUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

