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

// checks if the Alarm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Alarm{}

// Alarm additional information per alarm type
type Alarm struct {
	// UUID of the admin who acked the alarm
	AckAdminId *string `json:"ack_admin_id,omitempty"`
	// Name & Email ID of the admin who acked the alarm
	AckAdminName *string `json:"ack_admin_name,omitempty"`
	// Whether the alarm is acked or not
	Acked *bool `json:"acked,omitempty"`
	// Epoch (seconds) when the alarm was acked
	AckedTime *int32 `json:"acked_time,omitempty"`
	// additional information: List of MACs of the APs
	Aps []string `json:"aps,omitempty"`
	// List of BSSIDs
	Bssids []string `json:"bssids,omitempty"`
	// Number of incident within an alarm window
	Count int32 `json:"count"`
	// additional information: List of MACs of the gateways
	Gateways []string `json:"gateways,omitempty"`
	// Group of the alarm
	Group string `json:"group"`
	// additional information: List of Hostnames of the devices (AP/Switch/Gateway)
	Hostnames []string `json:"hostnames,omitempty"`
	// UUID of the alarm
	Id string `json:"id"`
	// Epoch (seconds) of the last incident/alarm within an alarm window
	LastSeen int32 `json:"last_seen"`
	// Text describing the alarm
	Note *string `json:"note,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	// Severity of the alarm
	Severity string `json:"severity"`
	SiteId *string `json:"site_id,omitempty"`
	// List of SSIDs
	Ssids []string `json:"ssids,omitempty"`
	// additional information: List of MACs of the switches
	Switches []string `json:"switches,omitempty"`
	// Epoch (seconds) of the first incident/alarm
	Timestamp int32 `json:"timestamp"`
	// Key-name of the alarm type
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _Alarm Alarm

// NewAlarm instantiates a new Alarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarm(count int32, group string, id string, lastSeen int32, severity string, timestamp int32, type_ string) *Alarm {
	this := Alarm{}
	this.Count = count
	this.Group = group
	this.Id = id
	this.LastSeen = lastSeen
	this.Severity = severity
	this.Timestamp = timestamp
	this.Type = type_
	return &this
}

// NewAlarmWithDefaults instantiates a new Alarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmWithDefaults() *Alarm {
	this := Alarm{}
	return &this
}

// GetAckAdminId returns the AckAdminId field value if set, zero value otherwise.
func (o *Alarm) GetAckAdminId() string {
	if o == nil || IsNil(o.AckAdminId) {
		var ret string
		return ret
	}
	return *o.AckAdminId
}

// GetAckAdminIdOk returns a tuple with the AckAdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetAckAdminIdOk() (*string, bool) {
	if o == nil || IsNil(o.AckAdminId) {
		return nil, false
	}
	return o.AckAdminId, true
}

// HasAckAdminId returns a boolean if a field has been set.
func (o *Alarm) HasAckAdminId() bool {
	if o != nil && !IsNil(o.AckAdminId) {
		return true
	}

	return false
}

// SetAckAdminId gets a reference to the given string and assigns it to the AckAdminId field.
func (o *Alarm) SetAckAdminId(v string) {
	o.AckAdminId = &v
}

// GetAckAdminName returns the AckAdminName field value if set, zero value otherwise.
func (o *Alarm) GetAckAdminName() string {
	if o == nil || IsNil(o.AckAdminName) {
		var ret string
		return ret
	}
	return *o.AckAdminName
}

// GetAckAdminNameOk returns a tuple with the AckAdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetAckAdminNameOk() (*string, bool) {
	if o == nil || IsNil(o.AckAdminName) {
		return nil, false
	}
	return o.AckAdminName, true
}

// HasAckAdminName returns a boolean if a field has been set.
func (o *Alarm) HasAckAdminName() bool {
	if o != nil && !IsNil(o.AckAdminName) {
		return true
	}

	return false
}

// SetAckAdminName gets a reference to the given string and assigns it to the AckAdminName field.
func (o *Alarm) SetAckAdminName(v string) {
	o.AckAdminName = &v
}

// GetAcked returns the Acked field value if set, zero value otherwise.
func (o *Alarm) GetAcked() bool {
	if o == nil || IsNil(o.Acked) {
		var ret bool
		return ret
	}
	return *o.Acked
}

// GetAckedOk returns a tuple with the Acked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetAckedOk() (*bool, bool) {
	if o == nil || IsNil(o.Acked) {
		return nil, false
	}
	return o.Acked, true
}

// HasAcked returns a boolean if a field has been set.
func (o *Alarm) HasAcked() bool {
	if o != nil && !IsNil(o.Acked) {
		return true
	}

	return false
}

// SetAcked gets a reference to the given bool and assigns it to the Acked field.
func (o *Alarm) SetAcked(v bool) {
	o.Acked = &v
}

// GetAckedTime returns the AckedTime field value if set, zero value otherwise.
func (o *Alarm) GetAckedTime() int32 {
	if o == nil || IsNil(o.AckedTime) {
		var ret int32
		return ret
	}
	return *o.AckedTime
}

// GetAckedTimeOk returns a tuple with the AckedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetAckedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AckedTime) {
		return nil, false
	}
	return o.AckedTime, true
}

// HasAckedTime returns a boolean if a field has been set.
func (o *Alarm) HasAckedTime() bool {
	if o != nil && !IsNil(o.AckedTime) {
		return true
	}

	return false
}

// SetAckedTime gets a reference to the given int32 and assigns it to the AckedTime field.
func (o *Alarm) SetAckedTime(v int32) {
	o.AckedTime = &v
}

// GetAps returns the Aps field value if set, zero value otherwise.
func (o *Alarm) GetAps() []string {
	if o == nil || IsNil(o.Aps) {
		var ret []string
		return ret
	}
	return o.Aps
}

// GetApsOk returns a tuple with the Aps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetApsOk() ([]string, bool) {
	if o == nil || IsNil(o.Aps) {
		return nil, false
	}
	return o.Aps, true
}

// HasAps returns a boolean if a field has been set.
func (o *Alarm) HasAps() bool {
	if o != nil && !IsNil(o.Aps) {
		return true
	}

	return false
}

// SetAps gets a reference to the given []string and assigns it to the Aps field.
func (o *Alarm) SetAps(v []string) {
	o.Aps = v
}

// GetBssids returns the Bssids field value if set, zero value otherwise.
func (o *Alarm) GetBssids() []string {
	if o == nil || IsNil(o.Bssids) {
		var ret []string
		return ret
	}
	return o.Bssids
}

// GetBssidsOk returns a tuple with the Bssids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetBssidsOk() ([]string, bool) {
	if o == nil || IsNil(o.Bssids) {
		return nil, false
	}
	return o.Bssids, true
}

// HasBssids returns a boolean if a field has been set.
func (o *Alarm) HasBssids() bool {
	if o != nil && !IsNil(o.Bssids) {
		return true
	}

	return false
}

// SetBssids gets a reference to the given []string and assigns it to the Bssids field.
func (o *Alarm) SetBssids(v []string) {
	o.Bssids = v
}

// GetCount returns the Count field value
func (o *Alarm) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *Alarm) SetCount(v int32) {
	o.Count = v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *Alarm) GetGateways() []string {
	if o == nil || IsNil(o.Gateways) {
		var ret []string
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetGatewaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *Alarm) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []string and assigns it to the Gateways field.
func (o *Alarm) SetGateways(v []string) {
	o.Gateways = v
}

// GetGroup returns the Group field value
func (o *Alarm) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *Alarm) SetGroup(v string) {
	o.Group = v
}

// GetHostnames returns the Hostnames field value if set, zero value otherwise.
func (o *Alarm) GetHostnames() []string {
	if o == nil || IsNil(o.Hostnames) {
		var ret []string
		return ret
	}
	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetHostnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostnames) {
		return nil, false
	}
	return o.Hostnames, true
}

// HasHostnames returns a boolean if a field has been set.
func (o *Alarm) HasHostnames() bool {
	if o != nil && !IsNil(o.Hostnames) {
		return true
	}

	return false
}

// SetHostnames gets a reference to the given []string and assigns it to the Hostnames field.
func (o *Alarm) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetId returns the Id field value
func (o *Alarm) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Alarm) SetId(v string) {
	o.Id = v
}

// GetLastSeen returns the LastSeen field value
func (o *Alarm) GetLastSeen() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetLastSeenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *Alarm) SetLastSeen(v int32) {
	o.LastSeen = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *Alarm) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *Alarm) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *Alarm) SetNote(v string) {
	o.Note = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Alarm) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Alarm) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Alarm) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSeverity returns the Severity field value
func (o *Alarm) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Alarm) SetSeverity(v string) {
	o.Severity = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Alarm) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Alarm) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Alarm) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSsids returns the Ssids field value if set, zero value otherwise.
func (o *Alarm) GetSsids() []string {
	if o == nil || IsNil(o.Ssids) {
		var ret []string
		return ret
	}
	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetSsidsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ssids) {
		return nil, false
	}
	return o.Ssids, true
}

// HasSsids returns a boolean if a field has been set.
func (o *Alarm) HasSsids() bool {
	if o != nil && !IsNil(o.Ssids) {
		return true
	}

	return false
}

// SetSsids gets a reference to the given []string and assigns it to the Ssids field.
func (o *Alarm) SetSsids(v []string) {
	o.Ssids = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *Alarm) GetSwitches() []string {
	if o == nil || IsNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetSwitchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *Alarm) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *Alarm) SetSwitches(v []string) {
	o.Switches = v
}

// GetTimestamp returns the Timestamp field value
func (o *Alarm) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Alarm) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetType returns the Type field value
func (o *Alarm) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Alarm) SetType(v string) {
	o.Type = v
}

func (o Alarm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alarm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AckAdminId) {
		toSerialize["ack_admin_id"] = o.AckAdminId
	}
	if !IsNil(o.AckAdminName) {
		toSerialize["ack_admin_name"] = o.AckAdminName
	}
	if !IsNil(o.Acked) {
		toSerialize["acked"] = o.Acked
	}
	if !IsNil(o.AckedTime) {
		toSerialize["acked_time"] = o.AckedTime
	}
	if !IsNil(o.Aps) {
		toSerialize["aps"] = o.Aps
	}
	if !IsNil(o.Bssids) {
		toSerialize["bssids"] = o.Bssids
	}
	toSerialize["count"] = o.Count
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	toSerialize["group"] = o.Group
	if !IsNil(o.Hostnames) {
		toSerialize["hostnames"] = o.Hostnames
	}
	toSerialize["id"] = o.Id
	toSerialize["last_seen"] = o.LastSeen
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	toSerialize["severity"] = o.Severity
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Ssids) {
		toSerialize["ssids"] = o.Ssids
	}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Alarm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"group",
		"id",
		"last_seen",
		"severity",
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

	varAlarm := _Alarm{}

	err = json.Unmarshal(data, &varAlarm)

	if err != nil {
		return err
	}

	*o = Alarm(varAlarm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ack_admin_id")
		delete(additionalProperties, "ack_admin_name")
		delete(additionalProperties, "acked")
		delete(additionalProperties, "acked_time")
		delete(additionalProperties, "aps")
		delete(additionalProperties, "bssids")
		delete(additionalProperties, "count")
		delete(additionalProperties, "gateways")
		delete(additionalProperties, "group")
		delete(additionalProperties, "hostnames")
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "note")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "ssids")
		delete(additionalProperties, "switches")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlarm struct {
	value *Alarm
	isSet bool
}

func (v NullableAlarm) Get() *Alarm {
	return v.value
}

func (v *NullableAlarm) Set(val *Alarm) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarm(val *Alarm) *NullableAlarm {
	return &NullableAlarm{value: val, isSet: true}
}

func (v NullableAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


