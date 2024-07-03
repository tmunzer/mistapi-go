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

// checks if the UiSettingsTileTimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiSettingsTileTimeRange{}

// UiSettingsTileTimeRange struct for UiSettingsTileTimeRange
type UiSettingsTileTimeRange struct {
	End *float32 `json:"end,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	Start *int32 `json:"start,omitempty"`
	UsePreset *bool `json:"usePreset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiSettingsTileTimeRange UiSettingsTileTimeRange

// NewUiSettingsTileTimeRange instantiates a new UiSettingsTileTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiSettingsTileTimeRange() *UiSettingsTileTimeRange {
	this := UiSettingsTileTimeRange{}
	return &this
}

// NewUiSettingsTileTimeRangeWithDefaults instantiates a new UiSettingsTileTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiSettingsTileTimeRangeWithDefaults() *UiSettingsTileTimeRange {
	this := UiSettingsTileTimeRange{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *UiSettingsTileTimeRange) SetEnd(v float32) {
	o.End = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *UiSettingsTileTimeRange) SetEndDate(v string) {
	o.EndDate = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *UiSettingsTileTimeRange) SetInterval(v string) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UiSettingsTileTimeRange) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *UiSettingsTileTimeRange) SetShortName(v string) {
	o.ShortName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *UiSettingsTileTimeRange) SetStart(v int32) {
	o.Start = &v
}

// GetUsePreset returns the UsePreset field value if set, zero value otherwise.
func (o *UiSettingsTileTimeRange) GetUsePreset() bool {
	if o == nil || IsNil(o.UsePreset) {
		var ret bool
		return ret
	}
	return *o.UsePreset
}

// GetUsePresetOk returns a tuple with the UsePreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTileTimeRange) GetUsePresetOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePreset) {
		return nil, false
	}
	return o.UsePreset, true
}

// HasUsePreset returns a boolean if a field has been set.
func (o *UiSettingsTileTimeRange) HasUsePreset() bool {
	if o != nil && !IsNil(o.UsePreset) {
		return true
	}

	return false
}

// SetUsePreset gets a reference to the given bool and assigns it to the UsePreset field.
func (o *UiSettingsTileTimeRange) SetUsePreset(v bool) {
	o.UsePreset = &v
}

func (o UiSettingsTileTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiSettingsTileTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.UsePreset) {
		toSerialize["usePreset"] = o.UsePreset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UiSettingsTileTimeRange) UnmarshalJSON(data []byte) (err error) {
	varUiSettingsTileTimeRange := _UiSettingsTileTimeRange{}

	err = json.Unmarshal(data, &varUiSettingsTileTimeRange)

	if err != nil {
		return err
	}

	*o = UiSettingsTileTimeRange(varUiSettingsTileTimeRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "name")
		delete(additionalProperties, "shortName")
		delete(additionalProperties, "start")
		delete(additionalProperties, "usePreset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiSettingsTileTimeRange struct {
	value *UiSettingsTileTimeRange
	isSet bool
}

func (v NullableUiSettingsTileTimeRange) Get() *UiSettingsTileTimeRange {
	return v.value
}

func (v *NullableUiSettingsTileTimeRange) Set(val *UiSettingsTileTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableUiSettingsTileTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableUiSettingsTileTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiSettingsTileTimeRange(val *UiSettingsTileTimeRange) *NullableUiSettingsTileTimeRange {
	return &NullableUiSettingsTileTimeRange{value: val, isSet: true}
}

func (v NullableUiSettingsTileTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiSettingsTileTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


