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

// WebhookDeliveryDistinct webhook topic
type WebhookDeliveryDistinct string

// List of webhook_delivery_distinct
const (
	WEBHOOKDELIVERYDISTINCT_EMPTY WebhookDeliveryDistinct = ""
	WEBHOOKDELIVERYDISTINCT_STATUS WebhookDeliveryDistinct = "status"
	WEBHOOKDELIVERYDISTINCT_TOPIC WebhookDeliveryDistinct = "topic"
	WEBHOOKDELIVERYDISTINCT_STATUS_CODE WebhookDeliveryDistinct = "status_code"
	WEBHOOKDELIVERYDISTINCT_WEBHOOK_ID WebhookDeliveryDistinct = "webhook_id"
)

// All allowed values of WebhookDeliveryDistinct enum
var AllowedWebhookDeliveryDistinctEnumValues = []WebhookDeliveryDistinct{
	"",
	"status",
	"topic",
	"status_code",
	"webhook_id",
}

func (v *WebhookDeliveryDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookDeliveryDistinct(value)
	for _, existing := range AllowedWebhookDeliveryDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookDeliveryDistinct", value)
}

// NewWebhookDeliveryDistinctFromValue returns a pointer to a valid WebhookDeliveryDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookDeliveryDistinctFromValue(v string) (*WebhookDeliveryDistinct, error) {
	ev := WebhookDeliveryDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookDeliveryDistinct: valid values are %v", v, AllowedWebhookDeliveryDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookDeliveryDistinct) IsValid() bool {
	for _, existing := range AllowedWebhookDeliveryDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_delivery_distinct value
func (v WebhookDeliveryDistinct) Ptr() *WebhookDeliveryDistinct {
	return &v
}

type NullableWebhookDeliveryDistinct struct {
	value *WebhookDeliveryDistinct
	isSet bool
}

func (v NullableWebhookDeliveryDistinct) Get() *WebhookDeliveryDistinct {
	return v.value
}

func (v *NullableWebhookDeliveryDistinct) Set(val *WebhookDeliveryDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryDistinct(val *WebhookDeliveryDistinct) *NullableWebhookDeliveryDistinct {
	return &NullableWebhookDeliveryDistinct{value: val, isSet: true}
}

func (v NullableWebhookDeliveryDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

