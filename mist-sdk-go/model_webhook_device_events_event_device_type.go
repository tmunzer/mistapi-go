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

// WebhookDeviceEventsEventDeviceType the model 'WebhookDeviceEventsEventDeviceType'
type WebhookDeviceEventsEventDeviceType string

// List of webhook_device_events_event_device_type
const (
	WEBHOOKDEVICEEVENTSEVENTDEVICETYPE_EMPTY WebhookDeviceEventsEventDeviceType = ""
	WEBHOOKDEVICEEVENTSEVENTDEVICETYPE_AP WebhookDeviceEventsEventDeviceType = "ap"
	WEBHOOKDEVICEEVENTSEVENTDEVICETYPE_SWITCH WebhookDeviceEventsEventDeviceType = "switch"
	WEBHOOKDEVICEEVENTSEVENTDEVICETYPE_GATEWAY WebhookDeviceEventsEventDeviceType = "gateway"
)

// All allowed values of WebhookDeviceEventsEventDeviceType enum
var AllowedWebhookDeviceEventsEventDeviceTypeEnumValues = []WebhookDeviceEventsEventDeviceType{
	"",
	"ap",
	"switch",
	"gateway",
}

func (v *WebhookDeviceEventsEventDeviceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookDeviceEventsEventDeviceType(value)
	for _, existing := range AllowedWebhookDeviceEventsEventDeviceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookDeviceEventsEventDeviceType", value)
}

// NewWebhookDeviceEventsEventDeviceTypeFromValue returns a pointer to a valid WebhookDeviceEventsEventDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookDeviceEventsEventDeviceTypeFromValue(v string) (*WebhookDeviceEventsEventDeviceType, error) {
	ev := WebhookDeviceEventsEventDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookDeviceEventsEventDeviceType: valid values are %v", v, AllowedWebhookDeviceEventsEventDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookDeviceEventsEventDeviceType) IsValid() bool {
	for _, existing := range AllowedWebhookDeviceEventsEventDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_device_events_event_device_type value
func (v WebhookDeviceEventsEventDeviceType) Ptr() *WebhookDeviceEventsEventDeviceType {
	return &v
}

type NullableWebhookDeviceEventsEventDeviceType struct {
	value *WebhookDeviceEventsEventDeviceType
	isSet bool
}

func (v NullableWebhookDeviceEventsEventDeviceType) Get() *WebhookDeviceEventsEventDeviceType {
	return v.value
}

func (v *NullableWebhookDeviceEventsEventDeviceType) Set(val *WebhookDeviceEventsEventDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeviceEventsEventDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeviceEventsEventDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeviceEventsEventDeviceType(val *WebhookDeviceEventsEventDeviceType) *NullableWebhookDeviceEventsEventDeviceType {
	return &NullableWebhookDeviceEventsEventDeviceType{value: val, isSet: true}
}

func (v NullableWebhookDeviceEventsEventDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeviceEventsEventDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

