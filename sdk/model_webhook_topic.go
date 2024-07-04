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

// WebhookTopic the model 'WebhookTopic'
type WebhookTopic string

// List of webhook_topic
const (
	WEBHOOKTOPIC_EMPTY WebhookTopic = ""
	WEBHOOKTOPIC_ALARMS WebhookTopic = "alarms"
	WEBHOOKTOPIC_AUDITS WebhookTopic = "audits"
	WEBHOOKTOPIC_CLIENT_INFO WebhookTopic = "client-info"
	WEBHOOKTOPIC_CLIENT_JOIN WebhookTopic = "client-join"
	WEBHOOKTOPIC_CLIENT_LATENCY WebhookTopic = "client-latency"
	WEBHOOKTOPIC_CLIENT_SESSIONS WebhookTopic = "client-sessions"
	WEBHOOKTOPIC_DEVICE_EVENTS WebhookTopic = "device_events"
	WEBHOOKTOPIC_DEVICE_UPDOWNS WebhookTopic = "device-updowns"
	WEBHOOKTOPIC_MXEDGE_EVENTS WebhookTopic = "mxedge_events"
	WEBHOOKTOPIC_NAC_ACCOUNTING WebhookTopic = "nac-accounting"
	WEBHOOKTOPIC_NAC_EVENTS WebhookTopic = "nac_events"
	WEBHOOKTOPIC_OCCUPANCY_ALERTS WebhookTopic = "occupancy-alerts"
	WEBHOOKTOPIC_RSSIZONE WebhookTopic = "rssizone"
	WEBHOOKTOPIC_SDKCLIENT_SCAN_DATA WebhookTopic = "sdkclient_scan_data"
	WEBHOOKTOPIC_SITE_SLE WebhookTopic = "site_sle"
	WEBHOOKTOPIC_VBEACON WebhookTopic = "vbeacon"
	WEBHOOKTOPIC_ZONE WebhookTopic = "zone"
	WEBHOOKTOPIC_LOCATION WebhookTopic = "location"
	WEBHOOKTOPIC_LOCATION_ASSET WebhookTopic = "location_asset"
	WEBHOOKTOPIC_LOCATION_CENTRAK WebhookTopic = "location_centrak"
	WEBHOOKTOPIC_LOCATION_CLIENT WebhookTopic = "location_client"
	WEBHOOKTOPIC_LOCATION_UNCLIENT WebhookTopic = "location_unclient"
	WEBHOOKTOPIC_LOCATION_SDK WebhookTopic = "location_sdk"
	WEBHOOKTOPIC_ASSET_RAW WebhookTopic = "asset-raw"
	WEBHOOKTOPIC_ASSET_RAW_RSSI WebhookTopic = "asset-raw-rssi"
	WEBHOOKTOPIC_DISCOVERED_RAW_RSSI WebhookTopic = "discovered-raw-rssi"
	WEBHOOKTOPIC_WIFI_CONN_RAW WebhookTopic = "wifi-conn-raw"
	WEBHOOKTOPIC_WIFI_UNCONN_RAW WebhookTopic = "wifi-unconn-raw"
)

// All allowed values of WebhookTopic enum
var AllowedWebhookTopicEnumValues = []WebhookTopic{
	"",
	"alarms",
	"audits",
	"client-info",
	"client-join",
	"client-latency",
	"client-sessions",
	"device_events",
	"device-updowns",
	"mxedge_events",
	"nac-accounting",
	"nac_events",
	"occupancy-alerts",
	"rssizone",
	"sdkclient_scan_data",
	"site_sle",
	"vbeacon",
	"zone",
	"location",
	"location_asset",
	"location_centrak",
	"location_client",
	"location_unclient",
	"location_sdk",
	"asset-raw",
	"asset-raw-rssi",
	"discovered-raw-rssi",
	"wifi-conn-raw",
	"wifi-unconn-raw",
}

func (v *WebhookTopic) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookTopic(value)
	for _, existing := range AllowedWebhookTopicEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookTopic", value)
}

// NewWebhookTopicFromValue returns a pointer to a valid WebhookTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookTopicFromValue(v string) (*WebhookTopic, error) {
	ev := WebhookTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookTopic: valid values are %v", v, AllowedWebhookTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookTopic) IsValid() bool {
	for _, existing := range AllowedWebhookTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_topic value
func (v WebhookTopic) Ptr() *WebhookTopic {
	return &v
}

type NullableWebhookTopic struct {
	value *WebhookTopic
	isSet bool
}

func (v NullableWebhookTopic) Get() *WebhookTopic {
	return v.value
}

func (v *NullableWebhookTopic) Set(val *WebhookTopic) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTopic(val *WebhookTopic) *NullableWebhookTopic {
	return &NullableWebhookTopic{value: val, isSet: true}
}

func (v NullableWebhookTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
