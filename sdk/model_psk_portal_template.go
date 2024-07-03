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

// checks if the PskPortalTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PskPortalTemplate{}

// PskPortalTemplate struct for PskPortalTemplate
type PskPortalTemplate struct {
	Alignment *PskPortalTemplateAlignment `json:"alignment,omitempty"`
	Color *string `json:"color,omitempty"`
	// custom logo.  default null, uses Juniper Mist Logo
	Logo NullableString `json:"logo,omitempty"`
	// whether to hide \"Powered by Juniper Mist\" and email footers
	PoweredBy *bool `json:"poweredBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PskPortalTemplate PskPortalTemplate

// NewPskPortalTemplate instantiates a new PskPortalTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPskPortalTemplate() *PskPortalTemplate {
	this := PskPortalTemplate{}
	var alignment PskPortalTemplateAlignment = PSKPORTALTEMPLATEALIGNMENT_CENTER
	this.Alignment = &alignment
	var color string = "#1074bc"
	this.Color = &color
	var poweredBy bool = false
	this.PoweredBy = &poweredBy
	return &this
}

// NewPskPortalTemplateWithDefaults instantiates a new PskPortalTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPskPortalTemplateWithDefaults() *PskPortalTemplate {
	this := PskPortalTemplate{}
	var alignment PskPortalTemplateAlignment = PSKPORTALTEMPLATEALIGNMENT_CENTER
	this.Alignment = &alignment
	var color string = "#1074bc"
	this.Color = &color
	var poweredBy bool = false
	this.PoweredBy = &poweredBy
	return &this
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *PskPortalTemplate) GetAlignment() PskPortalTemplateAlignment {
	if o == nil || IsNil(o.Alignment) {
		var ret PskPortalTemplateAlignment
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PskPortalTemplate) GetAlignmentOk() (*PskPortalTemplateAlignment, bool) {
	if o == nil || IsNil(o.Alignment) {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *PskPortalTemplate) HasAlignment() bool {
	if o != nil && !IsNil(o.Alignment) {
		return true
	}

	return false
}

// SetAlignment gets a reference to the given PskPortalTemplateAlignment and assigns it to the Alignment field.
func (o *PskPortalTemplate) SetAlignment(v PskPortalTemplateAlignment) {
	o.Alignment = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PskPortalTemplate) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PskPortalTemplate) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PskPortalTemplate) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PskPortalTemplate) SetColor(v string) {
	o.Color = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PskPortalTemplate) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PskPortalTemplate) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *PskPortalTemplate) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *PskPortalTemplate) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *PskPortalTemplate) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *PskPortalTemplate) UnsetLogo() {
	o.Logo.Unset()
}

// GetPoweredBy returns the PoweredBy field value if set, zero value otherwise.
func (o *PskPortalTemplate) GetPoweredBy() bool {
	if o == nil || IsNil(o.PoweredBy) {
		var ret bool
		return ret
	}
	return *o.PoweredBy
}

// GetPoweredByOk returns a tuple with the PoweredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PskPortalTemplate) GetPoweredByOk() (*bool, bool) {
	if o == nil || IsNil(o.PoweredBy) {
		return nil, false
	}
	return o.PoweredBy, true
}

// HasPoweredBy returns a boolean if a field has been set.
func (o *PskPortalTemplate) HasPoweredBy() bool {
	if o != nil && !IsNil(o.PoweredBy) {
		return true
	}

	return false
}

// SetPoweredBy gets a reference to the given bool and assigns it to the PoweredBy field.
func (o *PskPortalTemplate) SetPoweredBy(v bool) {
	o.PoweredBy = &v
}

func (o PskPortalTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PskPortalTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alignment) {
		toSerialize["alignment"] = o.Alignment
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if !IsNil(o.PoweredBy) {
		toSerialize["poweredBy"] = o.PoweredBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PskPortalTemplate) UnmarshalJSON(data []byte) (err error) {
	varPskPortalTemplate := _PskPortalTemplate{}

	err = json.Unmarshal(data, &varPskPortalTemplate)

	if err != nil {
		return err
	}

	*o = PskPortalTemplate(varPskPortalTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alignment")
		delete(additionalProperties, "color")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "poweredBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePskPortalTemplate struct {
	value *PskPortalTemplate
	isSet bool
}

func (v NullablePskPortalTemplate) Get() *PskPortalTemplate {
	return v.value
}

func (v *NullablePskPortalTemplate) Set(val *PskPortalTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePskPortalTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePskPortalTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePskPortalTemplate(val *PskPortalTemplate) *NullablePskPortalTemplate {
	return &NullablePskPortalTemplate{value: val, isSet: true}
}

func (v NullablePskPortalTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePskPortalTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


