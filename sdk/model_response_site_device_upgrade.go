/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// checks if the ResponseSiteDeviceUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSiteDeviceUpgrade{}

// ResponseSiteDeviceUpgrade struct for ResponseSiteDeviceUpgrade
type ResponseSiteDeviceUpgrade struct {
	Counts *ResponseSiteDeviceUpgradeCounts `json:"counts,omitempty"`
	// current canary or rrm phase in progress
	CurrentPhase *int32 `json:"current_phase,omitempty"`
	// whether to allow local AP-to-AP FW upgrade
	EnableP2p *bool `json:"enable_p2p,omitempty"`
	// whether to force upgrade when requested version is same as running version
	Force *bool `json:"force,omitempty"`
	// unique id for the upgrade
	Id string `json:"id"`
	// percentage of failures allowed
	MaxFailurePercentage *int32 `json:"max_failure_percentage,omitempty"`
	// number of failures allowed within a canary phase or serial rollout
	MaxFailures []int32 `json:"max_failures,omitempty"`
	// reboot start time in epoch
	RebootAt *int32 `json:"reboot_at,omitempty"`
	// firmware download start time in epoch
	StartTime *float32 `json:"start_time,omitempty"`
	Status *DeviceUpgradeStatus `json:"status,omitempty"`
	Strategy *DeviceUpgradeStrategy `json:"strategy,omitempty"`
	// version to upgrade to
	TargetVersion *string `json:"target_version,omitempty"`
	// a dictionary of rrm phase number to devices part of that phase
	UpgradePlan map[string]interface{} `json:"upgrade_plan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSiteDeviceUpgrade ResponseSiteDeviceUpgrade

// NewResponseSiteDeviceUpgrade instantiates a new ResponseSiteDeviceUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSiteDeviceUpgrade(id string) *ResponseSiteDeviceUpgrade {
	this := ResponseSiteDeviceUpgrade{}
	this.Id = id
	var strategy DeviceUpgradeStrategy = DEVICEUPGRADESTRATEGY_BIG_BANG
	this.Strategy = &strategy
	return &this
}

// NewResponseSiteDeviceUpgradeWithDefaults instantiates a new ResponseSiteDeviceUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSiteDeviceUpgradeWithDefaults() *ResponseSiteDeviceUpgrade {
	this := ResponseSiteDeviceUpgrade{}
	var strategy DeviceUpgradeStrategy = DEVICEUPGRADESTRATEGY_BIG_BANG
	this.Strategy = &strategy
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetCounts() ResponseSiteDeviceUpgradeCounts {
	if o == nil || IsNil(o.Counts) {
		var ret ResponseSiteDeviceUpgradeCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetCountsOk() (*ResponseSiteDeviceUpgradeCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given ResponseSiteDeviceUpgradeCounts and assigns it to the Counts field.
func (o *ResponseSiteDeviceUpgrade) SetCounts(v ResponseSiteDeviceUpgradeCounts) {
	o.Counts = &v
}

// GetCurrentPhase returns the CurrentPhase field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetCurrentPhase() int32 {
	if o == nil || IsNil(o.CurrentPhase) {
		var ret int32
		return ret
	}
	return *o.CurrentPhase
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetCurrentPhaseOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPhase) {
		return nil, false
	}
	return o.CurrentPhase, true
}

// HasCurrentPhase returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasCurrentPhase() bool {
	if o != nil && !IsNil(o.CurrentPhase) {
		return true
	}

	return false
}

// SetCurrentPhase gets a reference to the given int32 and assigns it to the CurrentPhase field.
func (o *ResponseSiteDeviceUpgrade) SetCurrentPhase(v int32) {
	o.CurrentPhase = &v
}

// GetEnableP2p returns the EnableP2p field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetEnableP2p() bool {
	if o == nil || IsNil(o.EnableP2p) {
		var ret bool
		return ret
	}
	return *o.EnableP2p
}

// GetEnableP2pOk returns a tuple with the EnableP2p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetEnableP2pOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableP2p) {
		return nil, false
	}
	return o.EnableP2p, true
}

// HasEnableP2p returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasEnableP2p() bool {
	if o != nil && !IsNil(o.EnableP2p) {
		return true
	}

	return false
}

// SetEnableP2p gets a reference to the given bool and assigns it to the EnableP2p field.
func (o *ResponseSiteDeviceUpgrade) SetEnableP2p(v bool) {
	o.EnableP2p = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *ResponseSiteDeviceUpgrade) SetForce(v bool) {
	o.Force = &v
}

// GetId returns the Id field value
func (o *ResponseSiteDeviceUpgrade) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseSiteDeviceUpgrade) SetId(v string) {
	o.Id = v
}

// GetMaxFailurePercentage returns the MaxFailurePercentage field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetMaxFailurePercentage() int32 {
	if o == nil || IsNil(o.MaxFailurePercentage) {
		var ret int32
		return ret
	}
	return *o.MaxFailurePercentage
}

// GetMaxFailurePercentageOk returns a tuple with the MaxFailurePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetMaxFailurePercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxFailurePercentage) {
		return nil, false
	}
	return o.MaxFailurePercentage, true
}

// HasMaxFailurePercentage returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasMaxFailurePercentage() bool {
	if o != nil && !IsNil(o.MaxFailurePercentage) {
		return true
	}

	return false
}

// SetMaxFailurePercentage gets a reference to the given int32 and assigns it to the MaxFailurePercentage field.
func (o *ResponseSiteDeviceUpgrade) SetMaxFailurePercentage(v int32) {
	o.MaxFailurePercentage = &v
}

// GetMaxFailures returns the MaxFailures field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetMaxFailures() []int32 {
	if o == nil || IsNil(o.MaxFailures) {
		var ret []int32
		return ret
	}
	return o.MaxFailures
}

// GetMaxFailuresOk returns a tuple with the MaxFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetMaxFailuresOk() ([]int32, bool) {
	if o == nil || IsNil(o.MaxFailures) {
		return nil, false
	}
	return o.MaxFailures, true
}

// HasMaxFailures returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasMaxFailures() bool {
	if o != nil && !IsNil(o.MaxFailures) {
		return true
	}

	return false
}

// SetMaxFailures gets a reference to the given []int32 and assigns it to the MaxFailures field.
func (o *ResponseSiteDeviceUpgrade) SetMaxFailures(v []int32) {
	o.MaxFailures = v
}

// GetRebootAt returns the RebootAt field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetRebootAt() int32 {
	if o == nil || IsNil(o.RebootAt) {
		var ret int32
		return ret
	}
	return *o.RebootAt
}

// GetRebootAtOk returns a tuple with the RebootAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetRebootAtOk() (*int32, bool) {
	if o == nil || IsNil(o.RebootAt) {
		return nil, false
	}
	return o.RebootAt, true
}

// HasRebootAt returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasRebootAt() bool {
	if o != nil && !IsNil(o.RebootAt) {
		return true
	}

	return false
}

// SetRebootAt gets a reference to the given int32 and assigns it to the RebootAt field.
func (o *ResponseSiteDeviceUpgrade) SetRebootAt(v int32) {
	o.RebootAt = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetStartTime() float32 {
	if o == nil || IsNil(o.StartTime) {
		var ret float32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetStartTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given float32 and assigns it to the StartTime field.
func (o *ResponseSiteDeviceUpgrade) SetStartTime(v float32) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetStatus() DeviceUpgradeStatus {
	if o == nil || IsNil(o.Status) {
		var ret DeviceUpgradeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetStatusOk() (*DeviceUpgradeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeviceUpgradeStatus and assigns it to the Status field.
func (o *ResponseSiteDeviceUpgrade) SetStatus(v DeviceUpgradeStatus) {
	o.Status = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetStrategy() DeviceUpgradeStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret DeviceUpgradeStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetStrategyOk() (*DeviceUpgradeStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given DeviceUpgradeStrategy and assigns it to the Strategy field.
func (o *ResponseSiteDeviceUpgrade) SetStrategy(v DeviceUpgradeStrategy) {
	o.Strategy = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetTargetVersion() string {
	if o == nil || IsNil(o.TargetVersion) {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetTargetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetVersion) {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasTargetVersion() bool {
	if o != nil && !IsNil(o.TargetVersion) {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *ResponseSiteDeviceUpgrade) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// GetUpgradePlan returns the UpgradePlan field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgrade) GetUpgradePlan() map[string]interface{} {
	if o == nil || IsNil(o.UpgradePlan) {
		var ret map[string]interface{}
		return ret
	}
	return o.UpgradePlan
}

// GetUpgradePlanOk returns a tuple with the UpgradePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgrade) GetUpgradePlanOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UpgradePlan) {
		return map[string]interface{}{}, false
	}
	return o.UpgradePlan, true
}

// HasUpgradePlan returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgrade) HasUpgradePlan() bool {
	if o != nil && !IsNil(o.UpgradePlan) {
		return true
	}

	return false
}

// SetUpgradePlan gets a reference to the given map[string]interface{} and assigns it to the UpgradePlan field.
func (o *ResponseSiteDeviceUpgrade) SetUpgradePlan(v map[string]interface{}) {
	o.UpgradePlan = v
}

func (o ResponseSiteDeviceUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSiteDeviceUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !IsNil(o.CurrentPhase) {
		toSerialize["current_phase"] = o.CurrentPhase
	}
	if !IsNil(o.EnableP2p) {
		toSerialize["enable_p2p"] = o.EnableP2p
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.MaxFailurePercentage) {
		toSerialize["max_failure_percentage"] = o.MaxFailurePercentage
	}
	if !IsNil(o.MaxFailures) {
		toSerialize["max_failures"] = o.MaxFailures
	}
	if !IsNil(o.RebootAt) {
		toSerialize["reboot_at"] = o.RebootAt
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.TargetVersion) {
		toSerialize["target_version"] = o.TargetVersion
	}
	if !IsNil(o.UpgradePlan) {
		toSerialize["upgrade_plan"] = o.UpgradePlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSiteDeviceUpgrade) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varResponseSiteDeviceUpgrade := _ResponseSiteDeviceUpgrade{}

	err = json.Unmarshal(data, &varResponseSiteDeviceUpgrade)

	if err != nil {
		return err
	}

	*o = ResponseSiteDeviceUpgrade(varResponseSiteDeviceUpgrade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "counts")
		delete(additionalProperties, "current_phase")
		delete(additionalProperties, "enable_p2p")
		delete(additionalProperties, "force")
		delete(additionalProperties, "id")
		delete(additionalProperties, "max_failure_percentage")
		delete(additionalProperties, "max_failures")
		delete(additionalProperties, "reboot_at")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "status")
		delete(additionalProperties, "strategy")
		delete(additionalProperties, "target_version")
		delete(additionalProperties, "upgrade_plan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSiteDeviceUpgrade struct {
	value *ResponseSiteDeviceUpgrade
	isSet bool
}

func (v NullableResponseSiteDeviceUpgrade) Get() *ResponseSiteDeviceUpgrade {
	return v.value
}

func (v *NullableResponseSiteDeviceUpgrade) Set(val *ResponseSiteDeviceUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSiteDeviceUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSiteDeviceUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSiteDeviceUpgrade(val *ResponseSiteDeviceUpgrade) *NullableResponseSiteDeviceUpgrade {
	return &NullableResponseSiteDeviceUpgrade{value: val, isSet: true}
}

func (v NullableResponseSiteDeviceUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSiteDeviceUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


