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

// checks if the EventsSkyatp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsSkyatp{}

// EventsSkyatp SkyATP events
type EventsSkyatp struct {
	DeviceMac string `json:"device_mac"`
	ForSite *bool `json:"for_site,omitempty"`
	Ip string `json:"ip"`
	Mac string `json:"mac"`
	OrgId string `json:"org_id"`
	SiteId string `json:"site_id"`
	ThreatLevel int32 `json:"threat_level"`
	Timestamp float32 `json:"timestamp"`
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _EventsSkyatp EventsSkyatp

// NewEventsSkyatp instantiates a new EventsSkyatp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsSkyatp(deviceMac string, ip string, mac string, orgId string, siteId string, threatLevel int32, timestamp float32, type_ string) *EventsSkyatp {
	this := EventsSkyatp{}
	this.DeviceMac = deviceMac
	this.Ip = ip
	this.Mac = mac
	this.OrgId = orgId
	this.SiteId = siteId
	this.ThreatLevel = threatLevel
	this.Timestamp = timestamp
	this.Type = type_
	return &this
}

// NewEventsSkyatpWithDefaults instantiates a new EventsSkyatp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsSkyatpWithDefaults() *EventsSkyatp {
	this := EventsSkyatp{}
	return &this
}

// GetDeviceMac returns the DeviceMac field value
func (o *EventsSkyatp) GetDeviceMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetDeviceMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceMac, true
}

// SetDeviceMac sets field value
func (o *EventsSkyatp) SetDeviceMac(v string) {
	o.DeviceMac = v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *EventsSkyatp) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *EventsSkyatp) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *EventsSkyatp) SetForSite(v bool) {
	o.ForSite = &v
}

// GetIp returns the Ip field value
func (o *EventsSkyatp) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *EventsSkyatp) SetIp(v string) {
	o.Ip = v
}

// GetMac returns the Mac field value
func (o *EventsSkyatp) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *EventsSkyatp) SetMac(v string) {
	o.Mac = v
}

// GetOrgId returns the OrgId field value
func (o *EventsSkyatp) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *EventsSkyatp) SetOrgId(v string) {
	o.OrgId = v
}

// GetSiteId returns the SiteId field value
func (o *EventsSkyatp) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *EventsSkyatp) SetSiteId(v string) {
	o.SiteId = v
}

// GetThreatLevel returns the ThreatLevel field value
func (o *EventsSkyatp) GetThreatLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetThreatLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreatLevel, true
}

// SetThreatLevel sets field value
func (o *EventsSkyatp) SetThreatLevel(v int32) {
	o.ThreatLevel = v
}

// GetTimestamp returns the Timestamp field value
func (o *EventsSkyatp) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EventsSkyatp) SetTimestamp(v float32) {
	o.Timestamp = v
}

// GetType returns the Type field value
func (o *EventsSkyatp) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventsSkyatp) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventsSkyatp) SetType(v string) {
	o.Type = v
}

func (o EventsSkyatp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsSkyatp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_mac"] = o.DeviceMac
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	toSerialize["ip"] = o.Ip
	toSerialize["mac"] = o.Mac
	toSerialize["org_id"] = o.OrgId
	toSerialize["site_id"] = o.SiteId
	toSerialize["threat_level"] = o.ThreatLevel
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventsSkyatp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device_mac",
		"ip",
		"mac",
		"org_id",
		"site_id",
		"threat_level",
		"timestamp",
		"type",
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

	varEventsSkyatp := _EventsSkyatp{}

	err = json.Unmarshal(data, &varEventsSkyatp)

	if err != nil {
		return err
	}

	*o = EventsSkyatp(varEventsSkyatp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_mac")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "threat_level")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventsSkyatp struct {
	value *EventsSkyatp
	isSet bool
}

func (v NullableEventsSkyatp) Get() *EventsSkyatp {
	return v.value
}

func (v *NullableEventsSkyatp) Set(val *EventsSkyatp) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsSkyatp) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsSkyatp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsSkyatp(val *EventsSkyatp) *NullableEventsSkyatp {
	return &NullableEventsSkyatp{value: val, isSet: true}
}

func (v NullableEventsSkyatp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsSkyatp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


