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

// checks if the ApIotOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApIotOutput{}

// ApIotOutput IoT output AP settings
type ApIotOutput struct {
	// whether to enable a pin
	Enabled *bool `json:"enabled,omitempty"`
	// optional; descriptive pin name
	Name *string `json:"name,omitempty"`
	// whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing
	Output *bool `json:"output,omitempty"`
	Pullup *ApIotOutputPullup `json:"pullup,omitempty"`
	// output pin signal level, default 0
	Value *int32 `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApIotOutput ApIotOutput

// NewApIotOutput instantiates a new ApIotOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApIotOutput() *ApIotOutput {
	this := ApIotOutput{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApIotOutputWithDefaults instantiates a new ApIotOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApIotOutputWithDefaults() *ApIotOutput {
	this := ApIotOutput{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApIotOutput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIotOutput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApIotOutput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApIotOutput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApIotOutput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIotOutput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApIotOutput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApIotOutput) SetName(v string) {
	o.Name = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *ApIotOutput) GetOutput() bool {
	if o == nil || IsNil(o.Output) {
		var ret bool
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIotOutput) GetOutputOk() (*bool, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *ApIotOutput) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given bool and assigns it to the Output field.
func (o *ApIotOutput) SetOutput(v bool) {
	o.Output = &v
}

// GetPullup returns the Pullup field value if set, zero value otherwise.
func (o *ApIotOutput) GetPullup() ApIotOutputPullup {
	if o == nil || IsNil(o.Pullup) {
		var ret ApIotOutputPullup
		return ret
	}
	return *o.Pullup
}

// GetPullupOk returns a tuple with the Pullup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIotOutput) GetPullupOk() (*ApIotOutputPullup, bool) {
	if o == nil || IsNil(o.Pullup) {
		return nil, false
	}
	return o.Pullup, true
}

// HasPullup returns a boolean if a field has been set.
func (o *ApIotOutput) HasPullup() bool {
	if o != nil && !IsNil(o.Pullup) {
		return true
	}

	return false
}

// SetPullup gets a reference to the given ApIotOutputPullup and assigns it to the Pullup field.
func (o *ApIotOutput) SetPullup(v ApIotOutputPullup) {
	o.Pullup = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApIotOutput) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApIotOutput) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApIotOutput) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ApIotOutput) SetValue(v int32) {
	o.Value = &v
}

func (o ApIotOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApIotOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Pullup) {
		toSerialize["pullup"] = o.Pullup
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApIotOutput) UnmarshalJSON(data []byte) (err error) {
	varApIotOutput := _ApIotOutput{}

	err = json.Unmarshal(data, &varApIotOutput)

	if err != nil {
		return err
	}

	*o = ApIotOutput(varApIotOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "output")
		delete(additionalProperties, "pullup")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApIotOutput struct {
	value *ApIotOutput
	isSet bool
}

func (v NullableApIotOutput) Get() *ApIotOutput {
	return v.value
}

func (v *NullableApIotOutput) Set(val *ApIotOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApIotOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApIotOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApIotOutput(val *ApIotOutput) *NullableApIotOutput {
	return &NullableApIotOutput{value: val, isSet: true}
}

func (v NullableApIotOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApIotOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

