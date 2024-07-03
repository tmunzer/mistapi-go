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

// PskPortalTemplateAlignment defines alignment on portal
type PskPortalTemplateAlignment string

// List of psk_portal_template_alignment
const (
	PSKPORTALTEMPLATEALIGNMENT_EMPTY PskPortalTemplateAlignment = ""
	PSKPORTALTEMPLATEALIGNMENT_LEFT PskPortalTemplateAlignment = "left"
	PSKPORTALTEMPLATEALIGNMENT_CENTER PskPortalTemplateAlignment = "center"
	PSKPORTALTEMPLATEALIGNMENT_RIGHT PskPortalTemplateAlignment = "right"
)

// All allowed values of PskPortalTemplateAlignment enum
var AllowedPskPortalTemplateAlignmentEnumValues = []PskPortalTemplateAlignment{
	"",
	"left",
	"center",
	"right",
}

func (v *PskPortalTemplateAlignment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PskPortalTemplateAlignment(value)
	for _, existing := range AllowedPskPortalTemplateAlignmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PskPortalTemplateAlignment", value)
}

// NewPskPortalTemplateAlignmentFromValue returns a pointer to a valid PskPortalTemplateAlignment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPskPortalTemplateAlignmentFromValue(v string) (*PskPortalTemplateAlignment, error) {
	ev := PskPortalTemplateAlignment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PskPortalTemplateAlignment: valid values are %v", v, AllowedPskPortalTemplateAlignmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PskPortalTemplateAlignment) IsValid() bool {
	for _, existing := range AllowedPskPortalTemplateAlignmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to psk_portal_template_alignment value
func (v PskPortalTemplateAlignment) Ptr() *PskPortalTemplateAlignment {
	return &v
}

type NullablePskPortalTemplateAlignment struct {
	value *PskPortalTemplateAlignment
	isSet bool
}

func (v NullablePskPortalTemplateAlignment) Get() *PskPortalTemplateAlignment {
	return v.value
}

func (v *NullablePskPortalTemplateAlignment) Set(val *PskPortalTemplateAlignment) {
	v.value = val
	v.isSet = true
}

func (v NullablePskPortalTemplateAlignment) IsSet() bool {
	return v.isSet
}

func (v *NullablePskPortalTemplateAlignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePskPortalTemplateAlignment(val *PskPortalTemplateAlignment) *NullablePskPortalTemplateAlignment {
	return &NullablePskPortalTemplateAlignment{value: val, isSet: true}
}

func (v NullablePskPortalTemplateAlignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePskPortalTemplateAlignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

