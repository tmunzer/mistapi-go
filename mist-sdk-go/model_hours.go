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

// checks if the Hours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hours{}

// Hours hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).   **Note**: If the dow is not defined then it’s treated as 00:00-23:59.
type Hours struct {
	Fri *string `json:"fri,omitempty"`
	Mon *string `json:"mon,omitempty"`
	Sat *string `json:"sat,omitempty"`
	Sun *string `json:"sun,omitempty"`
	Thu *string `json:"thu,omitempty"`
	Tue *string `json:"tue,omitempty"`
	Wed *string `json:"wed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Hours Hours

// NewHours instantiates a new Hours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHours() *Hours {
	this := Hours{}
	var fri string = ""
	this.Fri = &fri
	var mon string = ""
	this.Mon = &mon
	var sat string = ""
	this.Sat = &sat
	var sun string = ""
	this.Sun = &sun
	var thu string = ""
	this.Thu = &thu
	var tue string = ""
	this.Tue = &tue
	var wed string = ""
	this.Wed = &wed
	return &this
}

// NewHoursWithDefaults instantiates a new Hours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoursWithDefaults() *Hours {
	this := Hours{}
	var fri string = ""
	this.Fri = &fri
	var mon string = ""
	this.Mon = &mon
	var sat string = ""
	this.Sat = &sat
	var sun string = ""
	this.Sun = &sun
	var thu string = ""
	this.Thu = &thu
	var tue string = ""
	this.Tue = &tue
	var wed string = ""
	this.Wed = &wed
	return &this
}

// GetFri returns the Fri field value if set, zero value otherwise.
func (o *Hours) GetFri() string {
	if o == nil || IsNil(o.Fri) {
		var ret string
		return ret
	}
	return *o.Fri
}

// GetFriOk returns a tuple with the Fri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetFriOk() (*string, bool) {
	if o == nil || IsNil(o.Fri) {
		return nil, false
	}
	return o.Fri, true
}

// HasFri returns a boolean if a field has been set.
func (o *Hours) HasFri() bool {
	if o != nil && !IsNil(o.Fri) {
		return true
	}

	return false
}

// SetFri gets a reference to the given string and assigns it to the Fri field.
func (o *Hours) SetFri(v string) {
	o.Fri = &v
}

// GetMon returns the Mon field value if set, zero value otherwise.
func (o *Hours) GetMon() string {
	if o == nil || IsNil(o.Mon) {
		var ret string
		return ret
	}
	return *o.Mon
}

// GetMonOk returns a tuple with the Mon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetMonOk() (*string, bool) {
	if o == nil || IsNil(o.Mon) {
		return nil, false
	}
	return o.Mon, true
}

// HasMon returns a boolean if a field has been set.
func (o *Hours) HasMon() bool {
	if o != nil && !IsNil(o.Mon) {
		return true
	}

	return false
}

// SetMon gets a reference to the given string and assigns it to the Mon field.
func (o *Hours) SetMon(v string) {
	o.Mon = &v
}

// GetSat returns the Sat field value if set, zero value otherwise.
func (o *Hours) GetSat() string {
	if o == nil || IsNil(o.Sat) {
		var ret string
		return ret
	}
	return *o.Sat
}

// GetSatOk returns a tuple with the Sat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetSatOk() (*string, bool) {
	if o == nil || IsNil(o.Sat) {
		return nil, false
	}
	return o.Sat, true
}

// HasSat returns a boolean if a field has been set.
func (o *Hours) HasSat() bool {
	if o != nil && !IsNil(o.Sat) {
		return true
	}

	return false
}

// SetSat gets a reference to the given string and assigns it to the Sat field.
func (o *Hours) SetSat(v string) {
	o.Sat = &v
}

// GetSun returns the Sun field value if set, zero value otherwise.
func (o *Hours) GetSun() string {
	if o == nil || IsNil(o.Sun) {
		var ret string
		return ret
	}
	return *o.Sun
}

// GetSunOk returns a tuple with the Sun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetSunOk() (*string, bool) {
	if o == nil || IsNil(o.Sun) {
		return nil, false
	}
	return o.Sun, true
}

// HasSun returns a boolean if a field has been set.
func (o *Hours) HasSun() bool {
	if o != nil && !IsNil(o.Sun) {
		return true
	}

	return false
}

// SetSun gets a reference to the given string and assigns it to the Sun field.
func (o *Hours) SetSun(v string) {
	o.Sun = &v
}

// GetThu returns the Thu field value if set, zero value otherwise.
func (o *Hours) GetThu() string {
	if o == nil || IsNil(o.Thu) {
		var ret string
		return ret
	}
	return *o.Thu
}

// GetThuOk returns a tuple with the Thu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetThuOk() (*string, bool) {
	if o == nil || IsNil(o.Thu) {
		return nil, false
	}
	return o.Thu, true
}

// HasThu returns a boolean if a field has been set.
func (o *Hours) HasThu() bool {
	if o != nil && !IsNil(o.Thu) {
		return true
	}

	return false
}

// SetThu gets a reference to the given string and assigns it to the Thu field.
func (o *Hours) SetThu(v string) {
	o.Thu = &v
}

// GetTue returns the Tue field value if set, zero value otherwise.
func (o *Hours) GetTue() string {
	if o == nil || IsNil(o.Tue) {
		var ret string
		return ret
	}
	return *o.Tue
}

// GetTueOk returns a tuple with the Tue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetTueOk() (*string, bool) {
	if o == nil || IsNil(o.Tue) {
		return nil, false
	}
	return o.Tue, true
}

// HasTue returns a boolean if a field has been set.
func (o *Hours) HasTue() bool {
	if o != nil && !IsNil(o.Tue) {
		return true
	}

	return false
}

// SetTue gets a reference to the given string and assigns it to the Tue field.
func (o *Hours) SetTue(v string) {
	o.Tue = &v
}

// GetWed returns the Wed field value if set, zero value otherwise.
func (o *Hours) GetWed() string {
	if o == nil || IsNil(o.Wed) {
		var ret string
		return ret
	}
	return *o.Wed
}

// GetWedOk returns a tuple with the Wed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetWedOk() (*string, bool) {
	if o == nil || IsNil(o.Wed) {
		return nil, false
	}
	return o.Wed, true
}

// HasWed returns a boolean if a field has been set.
func (o *Hours) HasWed() bool {
	if o != nil && !IsNil(o.Wed) {
		return true
	}

	return false
}

// SetWed gets a reference to the given string and assigns it to the Wed field.
func (o *Hours) SetWed(v string) {
	o.Wed = &v
}

func (o Hours) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fri) {
		toSerialize["fri"] = o.Fri
	}
	if !IsNil(o.Mon) {
		toSerialize["mon"] = o.Mon
	}
	if !IsNil(o.Sat) {
		toSerialize["sat"] = o.Sat
	}
	if !IsNil(o.Sun) {
		toSerialize["sun"] = o.Sun
	}
	if !IsNil(o.Thu) {
		toSerialize["thu"] = o.Thu
	}
	if !IsNil(o.Tue) {
		toSerialize["tue"] = o.Tue
	}
	if !IsNil(o.Wed) {
		toSerialize["wed"] = o.Wed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Hours) UnmarshalJSON(data []byte) (err error) {
	varHours := _Hours{}

	err = json.Unmarshal(data, &varHours)

	if err != nil {
		return err
	}

	*o = Hours(varHours)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fri")
		delete(additionalProperties, "mon")
		delete(additionalProperties, "sat")
		delete(additionalProperties, "sun")
		delete(additionalProperties, "thu")
		delete(additionalProperties, "tue")
		delete(additionalProperties, "wed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHours struct {
	value *Hours
	isSet bool
}

func (v NullableHours) Get() *Hours {
	return v.value
}

func (v *NullableHours) Set(val *Hours) {
	v.value = val
	v.isSet = true
}

func (v NullableHours) IsSet() bool {
	return v.isSet
}

func (v *NullableHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHours(val *Hours) *NullableHours {
	return &NullableHours{value: val, isSet: true}
}

func (v NullableHours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


