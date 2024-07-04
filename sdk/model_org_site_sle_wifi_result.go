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

// checks if the OrgSiteSleWifiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSiteSleWifiResult{}

// OrgSiteSleWifiResult struct for OrgSiteSleWifiResult
type OrgSiteSleWifiResult struct {
	ApAvailability *float32 `json:"ap-availability,omitempty"`
	ApHealth *float32 `json:"ap-health,omitempty"`
	Capacity *float32 `json:"capacity,omitempty"`
	Coverage *float32 `json:"coverage,omitempty"`
	NumAps *float32 `json:"num_aps,omitempty"`
	NumClients *float32 `json:"num_clients,omitempty"`
	Roaming *float32 `json:"roaming,omitempty"`
	SiteId string `json:"site_id"`
	SuccessfulConnect *float32 `json:"successful-connect,omitempty"`
	Throughput *float32 `json:"throughput,omitempty"`
	TimeToConnect *float32 `json:"time-to-connect,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSiteSleWifiResult OrgSiteSleWifiResult

// NewOrgSiteSleWifiResult instantiates a new OrgSiteSleWifiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSiteSleWifiResult(siteId string) *OrgSiteSleWifiResult {
	this := OrgSiteSleWifiResult{}
	this.SiteId = siteId
	return &this
}

// NewOrgSiteSleWifiResultWithDefaults instantiates a new OrgSiteSleWifiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSiteSleWifiResultWithDefaults() *OrgSiteSleWifiResult {
	this := OrgSiteSleWifiResult{}
	return &this
}

// GetApAvailability returns the ApAvailability field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetApAvailability() float32 {
	if o == nil || IsNil(o.ApAvailability) {
		var ret float32
		return ret
	}
	return *o.ApAvailability
}

// GetApAvailabilityOk returns a tuple with the ApAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetApAvailabilityOk() (*float32, bool) {
	if o == nil || IsNil(o.ApAvailability) {
		return nil, false
	}
	return o.ApAvailability, true
}

// HasApAvailability returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasApAvailability() bool {
	if o != nil && !IsNil(o.ApAvailability) {
		return true
	}

	return false
}

// SetApAvailability gets a reference to the given float32 and assigns it to the ApAvailability field.
func (o *OrgSiteSleWifiResult) SetApAvailability(v float32) {
	o.ApAvailability = &v
}

// GetApHealth returns the ApHealth field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetApHealth() float32 {
	if o == nil || IsNil(o.ApHealth) {
		var ret float32
		return ret
	}
	return *o.ApHealth
}

// GetApHealthOk returns a tuple with the ApHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetApHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.ApHealth) {
		return nil, false
	}
	return o.ApHealth, true
}

// HasApHealth returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasApHealth() bool {
	if o != nil && !IsNil(o.ApHealth) {
		return true
	}

	return false
}

// SetApHealth gets a reference to the given float32 and assigns it to the ApHealth field.
func (o *OrgSiteSleWifiResult) SetApHealth(v float32) {
	o.ApHealth = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetCapacity() float32 {
	if o == nil || IsNil(o.Capacity) {
		var ret float32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given float32 and assigns it to the Capacity field.
func (o *OrgSiteSleWifiResult) SetCapacity(v float32) {
	o.Capacity = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetCoverage() float32 {
	if o == nil || IsNil(o.Coverage) {
		var ret float32
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetCoverageOk() (*float32, bool) {
	if o == nil || IsNil(o.Coverage) {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasCoverage() bool {
	if o != nil && !IsNil(o.Coverage) {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given float32 and assigns it to the Coverage field.
func (o *OrgSiteSleWifiResult) SetCoverage(v float32) {
	o.Coverage = &v
}

// GetNumAps returns the NumAps field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetNumAps() float32 {
	if o == nil || IsNil(o.NumAps) {
		var ret float32
		return ret
	}
	return *o.NumAps
}

// GetNumApsOk returns a tuple with the NumAps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetNumApsOk() (*float32, bool) {
	if o == nil || IsNil(o.NumAps) {
		return nil, false
	}
	return o.NumAps, true
}

// HasNumAps returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasNumAps() bool {
	if o != nil && !IsNil(o.NumAps) {
		return true
	}

	return false
}

// SetNumAps gets a reference to the given float32 and assigns it to the NumAps field.
func (o *OrgSiteSleWifiResult) SetNumAps(v float32) {
	o.NumAps = &v
}

// GetNumClients returns the NumClients field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetNumClients() float32 {
	if o == nil || IsNil(o.NumClients) {
		var ret float32
		return ret
	}
	return *o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetNumClientsOk() (*float32, bool) {
	if o == nil || IsNil(o.NumClients) {
		return nil, false
	}
	return o.NumClients, true
}

// HasNumClients returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasNumClients() bool {
	if o != nil && !IsNil(o.NumClients) {
		return true
	}

	return false
}

// SetNumClients gets a reference to the given float32 and assigns it to the NumClients field.
func (o *OrgSiteSleWifiResult) SetNumClients(v float32) {
	o.NumClients = &v
}

// GetRoaming returns the Roaming field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetRoaming() float32 {
	if o == nil || IsNil(o.Roaming) {
		var ret float32
		return ret
	}
	return *o.Roaming
}

// GetRoamingOk returns a tuple with the Roaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetRoamingOk() (*float32, bool) {
	if o == nil || IsNil(o.Roaming) {
		return nil, false
	}
	return o.Roaming, true
}

// HasRoaming returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasRoaming() bool {
	if o != nil && !IsNil(o.Roaming) {
		return true
	}

	return false
}

// SetRoaming gets a reference to the given float32 and assigns it to the Roaming field.
func (o *OrgSiteSleWifiResult) SetRoaming(v float32) {
	o.Roaming = &v
}

// GetSiteId returns the SiteId field value
func (o *OrgSiteSleWifiResult) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *OrgSiteSleWifiResult) SetSiteId(v string) {
	o.SiteId = v
}

// GetSuccessfulConnect returns the SuccessfulConnect field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetSuccessfulConnect() float32 {
	if o == nil || IsNil(o.SuccessfulConnect) {
		var ret float32
		return ret
	}
	return *o.SuccessfulConnect
}

// GetSuccessfulConnectOk returns a tuple with the SuccessfulConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetSuccessfulConnectOk() (*float32, bool) {
	if o == nil || IsNil(o.SuccessfulConnect) {
		return nil, false
	}
	return o.SuccessfulConnect, true
}

// HasSuccessfulConnect returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasSuccessfulConnect() bool {
	if o != nil && !IsNil(o.SuccessfulConnect) {
		return true
	}

	return false
}

// SetSuccessfulConnect gets a reference to the given float32 and assigns it to the SuccessfulConnect field.
func (o *OrgSiteSleWifiResult) SetSuccessfulConnect(v float32) {
	o.SuccessfulConnect = &v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetThroughput() float32 {
	if o == nil || IsNil(o.Throughput) {
		var ret float32
		return ret
	}
	return *o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetThroughputOk() (*float32, bool) {
	if o == nil || IsNil(o.Throughput) {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasThroughput() bool {
	if o != nil && !IsNil(o.Throughput) {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given float32 and assigns it to the Throughput field.
func (o *OrgSiteSleWifiResult) SetThroughput(v float32) {
	o.Throughput = &v
}

// GetTimeToConnect returns the TimeToConnect field value if set, zero value otherwise.
func (o *OrgSiteSleWifiResult) GetTimeToConnect() float32 {
	if o == nil || IsNil(o.TimeToConnect) {
		var ret float32
		return ret
	}
	return *o.TimeToConnect
}

// GetTimeToConnectOk returns a tuple with the TimeToConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWifiResult) GetTimeToConnectOk() (*float32, bool) {
	if o == nil || IsNil(o.TimeToConnect) {
		return nil, false
	}
	return o.TimeToConnect, true
}

// HasTimeToConnect returns a boolean if a field has been set.
func (o *OrgSiteSleWifiResult) HasTimeToConnect() bool {
	if o != nil && !IsNil(o.TimeToConnect) {
		return true
	}

	return false
}

// SetTimeToConnect gets a reference to the given float32 and assigns it to the TimeToConnect field.
func (o *OrgSiteSleWifiResult) SetTimeToConnect(v float32) {
	o.TimeToConnect = &v
}

func (o OrgSiteSleWifiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSiteSleWifiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApAvailability) {
		toSerialize["ap-availability"] = o.ApAvailability
	}
	if !IsNil(o.ApHealth) {
		toSerialize["ap-health"] = o.ApHealth
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.Coverage) {
		toSerialize["coverage"] = o.Coverage
	}
	if !IsNil(o.NumAps) {
		toSerialize["num_aps"] = o.NumAps
	}
	if !IsNil(o.NumClients) {
		toSerialize["num_clients"] = o.NumClients
	}
	if !IsNil(o.Roaming) {
		toSerialize["roaming"] = o.Roaming
	}
	toSerialize["site_id"] = o.SiteId
	if !IsNil(o.SuccessfulConnect) {
		toSerialize["successful-connect"] = o.SuccessfulConnect
	}
	if !IsNil(o.Throughput) {
		toSerialize["throughput"] = o.Throughput
	}
	if !IsNil(o.TimeToConnect) {
		toSerialize["time-to-connect"] = o.TimeToConnect
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSiteSleWifiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"site_id",
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

	varOrgSiteSleWifiResult := _OrgSiteSleWifiResult{}

	err = json.Unmarshal(data, &varOrgSiteSleWifiResult)

	if err != nil {
		return err
	}

	*o = OrgSiteSleWifiResult(varOrgSiteSleWifiResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap-availability")
		delete(additionalProperties, "ap-health")
		delete(additionalProperties, "capacity")
		delete(additionalProperties, "coverage")
		delete(additionalProperties, "num_aps")
		delete(additionalProperties, "num_clients")
		delete(additionalProperties, "roaming")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "successful-connect")
		delete(additionalProperties, "throughput")
		delete(additionalProperties, "time-to-connect")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSiteSleWifiResult struct {
	value *OrgSiteSleWifiResult
	isSet bool
}

func (v NullableOrgSiteSleWifiResult) Get() *OrgSiteSleWifiResult {
	return v.value
}

func (v *NullableOrgSiteSleWifiResult) Set(val *OrgSiteSleWifiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSiteSleWifiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSiteSleWifiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSiteSleWifiResult(val *OrgSiteSleWifiResult) *NullableOrgSiteSleWifiResult {
	return &NullableOrgSiteSleWifiResult{value: val, isSet: true}
}

func (v NullableOrgSiteSleWifiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSiteSleWifiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

