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

// checks if the ResponseSiteDeviceUpgradeCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSiteDeviceUpgradeCounts{}

// ResponseSiteDeviceUpgradeCounts struct for ResponseSiteDeviceUpgradeCounts
type ResponseSiteDeviceUpgradeCounts struct {
	// count of devices which cloud has requested to download firmware
	DownloadRequested *int32 `json:"download_requested,omitempty"`
	// count of ap’s which have the firmware downloaded
	Downloaded *int32 `json:"downloaded,omitempty"`
	// count of devices which have failed to upgrade
	Failed *int32 `json:"failed,omitempty"`
	// count of devices which are rebooting
	RebootInProgress *int32 `json:"reboot_in_progress,omitempty"`
	// count of devices which have rebooted successfully
	Rebooted *int32 `json:"rebooted,omitempty"`
	// count of devices which cloud has scheduled an upgrade for
	Scheduled *int32 `json:"scheduled,omitempty"`
	// count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade
	Skipped *int32 `json:"skipped,omitempty"`
	// count of devices part of this upgrade
	Total *int32 `json:"total,omitempty"`
	// count of devices which have upgraded successfully
	Upgraded *int32 `json:"upgraded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSiteDeviceUpgradeCounts ResponseSiteDeviceUpgradeCounts

// NewResponseSiteDeviceUpgradeCounts instantiates a new ResponseSiteDeviceUpgradeCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSiteDeviceUpgradeCounts() *ResponseSiteDeviceUpgradeCounts {
	this := ResponseSiteDeviceUpgradeCounts{}
	return &this
}

// NewResponseSiteDeviceUpgradeCountsWithDefaults instantiates a new ResponseSiteDeviceUpgradeCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSiteDeviceUpgradeCountsWithDefaults() *ResponseSiteDeviceUpgradeCounts {
	this := ResponseSiteDeviceUpgradeCounts{}
	return &this
}

// GetDownloadRequested returns the DownloadRequested field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetDownloadRequested() int32 {
	if o == nil || IsNil(o.DownloadRequested) {
		var ret int32
		return ret
	}
	return *o.DownloadRequested
}

// GetDownloadRequestedOk returns a tuple with the DownloadRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetDownloadRequestedOk() (*int32, bool) {
	if o == nil || IsNil(o.DownloadRequested) {
		return nil, false
	}
	return o.DownloadRequested, true
}

// HasDownloadRequested returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasDownloadRequested() bool {
	if o != nil && !IsNil(o.DownloadRequested) {
		return true
	}

	return false
}

// SetDownloadRequested gets a reference to the given int32 and assigns it to the DownloadRequested field.
func (o *ResponseSiteDeviceUpgradeCounts) SetDownloadRequested(v int32) {
	o.DownloadRequested = &v
}

// GetDownloaded returns the Downloaded field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetDownloaded() int32 {
	if o == nil || IsNil(o.Downloaded) {
		var ret int32
		return ret
	}
	return *o.Downloaded
}

// GetDownloadedOk returns a tuple with the Downloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetDownloadedOk() (*int32, bool) {
	if o == nil || IsNil(o.Downloaded) {
		return nil, false
	}
	return o.Downloaded, true
}

// HasDownloaded returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasDownloaded() bool {
	if o != nil && !IsNil(o.Downloaded) {
		return true
	}

	return false
}

// SetDownloaded gets a reference to the given int32 and assigns it to the Downloaded field.
func (o *ResponseSiteDeviceUpgradeCounts) SetDownloaded(v int32) {
	o.Downloaded = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *ResponseSiteDeviceUpgradeCounts) SetFailed(v int32) {
	o.Failed = &v
}

// GetRebootInProgress returns the RebootInProgress field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetRebootInProgress() int32 {
	if o == nil || IsNil(o.RebootInProgress) {
		var ret int32
		return ret
	}
	return *o.RebootInProgress
}

// GetRebootInProgressOk returns a tuple with the RebootInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetRebootInProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.RebootInProgress) {
		return nil, false
	}
	return o.RebootInProgress, true
}

// HasRebootInProgress returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasRebootInProgress() bool {
	if o != nil && !IsNil(o.RebootInProgress) {
		return true
	}

	return false
}

// SetRebootInProgress gets a reference to the given int32 and assigns it to the RebootInProgress field.
func (o *ResponseSiteDeviceUpgradeCounts) SetRebootInProgress(v int32) {
	o.RebootInProgress = &v
}

// GetRebooted returns the Rebooted field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetRebooted() int32 {
	if o == nil || IsNil(o.Rebooted) {
		var ret int32
		return ret
	}
	return *o.Rebooted
}

// GetRebootedOk returns a tuple with the Rebooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetRebootedOk() (*int32, bool) {
	if o == nil || IsNil(o.Rebooted) {
		return nil, false
	}
	return o.Rebooted, true
}

// HasRebooted returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasRebooted() bool {
	if o != nil && !IsNil(o.Rebooted) {
		return true
	}

	return false
}

// SetRebooted gets a reference to the given int32 and assigns it to the Rebooted field.
func (o *ResponseSiteDeviceUpgradeCounts) SetRebooted(v int32) {
	o.Rebooted = &v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetScheduled() int32 {
	if o == nil || IsNil(o.Scheduled) {
		var ret int32
		return ret
	}
	return *o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetScheduledOk() (*int32, bool) {
	if o == nil || IsNil(o.Scheduled) {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasScheduled() bool {
	if o != nil && !IsNil(o.Scheduled) {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given int32 and assigns it to the Scheduled field.
func (o *ResponseSiteDeviceUpgradeCounts) SetScheduled(v int32) {
	o.Scheduled = &v
}

// GetSkipped returns the Skipped field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetSkipped() int32 {
	if o == nil || IsNil(o.Skipped) {
		var ret int32
		return ret
	}
	return *o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetSkippedOk() (*int32, bool) {
	if o == nil || IsNil(o.Skipped) {
		return nil, false
	}
	return o.Skipped, true
}

// HasSkipped returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasSkipped() bool {
	if o != nil && !IsNil(o.Skipped) {
		return true
	}

	return false
}

// SetSkipped gets a reference to the given int32 and assigns it to the Skipped field.
func (o *ResponseSiteDeviceUpgradeCounts) SetSkipped(v int32) {
	o.Skipped = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponseSiteDeviceUpgradeCounts) SetTotal(v int32) {
	o.Total = &v
}

// GetUpgraded returns the Upgraded field value if set, zero value otherwise.
func (o *ResponseSiteDeviceUpgradeCounts) GetUpgraded() int32 {
	if o == nil || IsNil(o.Upgraded) {
		var ret int32
		return ret
	}
	return *o.Upgraded
}

// GetUpgradedOk returns a tuple with the Upgraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSiteDeviceUpgradeCounts) GetUpgradedOk() (*int32, bool) {
	if o == nil || IsNil(o.Upgraded) {
		return nil, false
	}
	return o.Upgraded, true
}

// HasUpgraded returns a boolean if a field has been set.
func (o *ResponseSiteDeviceUpgradeCounts) HasUpgraded() bool {
	if o != nil && !IsNil(o.Upgraded) {
		return true
	}

	return false
}

// SetUpgraded gets a reference to the given int32 and assigns it to the Upgraded field.
func (o *ResponseSiteDeviceUpgradeCounts) SetUpgraded(v int32) {
	o.Upgraded = &v
}

func (o ResponseSiteDeviceUpgradeCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSiteDeviceUpgradeCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadRequested) {
		toSerialize["download_requested"] = o.DownloadRequested
	}
	if !IsNil(o.Downloaded) {
		toSerialize["downloaded"] = o.Downloaded
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.RebootInProgress) {
		toSerialize["reboot_in_progress"] = o.RebootInProgress
	}
	if !IsNil(o.Rebooted) {
		toSerialize["rebooted"] = o.Rebooted
	}
	if !IsNil(o.Scheduled) {
		toSerialize["scheduled"] = o.Scheduled
	}
	if !IsNil(o.Skipped) {
		toSerialize["skipped"] = o.Skipped
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Upgraded) {
		toSerialize["upgraded"] = o.Upgraded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSiteDeviceUpgradeCounts) UnmarshalJSON(data []byte) (err error) {
	varResponseSiteDeviceUpgradeCounts := _ResponseSiteDeviceUpgradeCounts{}

	err = json.Unmarshal(data, &varResponseSiteDeviceUpgradeCounts)

	if err != nil {
		return err
	}

	*o = ResponseSiteDeviceUpgradeCounts(varResponseSiteDeviceUpgradeCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "download_requested")
		delete(additionalProperties, "downloaded")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "reboot_in_progress")
		delete(additionalProperties, "rebooted")
		delete(additionalProperties, "scheduled")
		delete(additionalProperties, "skipped")
		delete(additionalProperties, "total")
		delete(additionalProperties, "upgraded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSiteDeviceUpgradeCounts struct {
	value *ResponseSiteDeviceUpgradeCounts
	isSet bool
}

func (v NullableResponseSiteDeviceUpgradeCounts) Get() *ResponseSiteDeviceUpgradeCounts {
	return v.value
}

func (v *NullableResponseSiteDeviceUpgradeCounts) Set(val *ResponseSiteDeviceUpgradeCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSiteDeviceUpgradeCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSiteDeviceUpgradeCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSiteDeviceUpgradeCounts(val *ResponseSiteDeviceUpgradeCounts) *NullableResponseSiteDeviceUpgradeCounts {
	return &NullableResponseSiteDeviceUpgradeCounts{value: val, isSet: true}
}

func (v NullableResponseSiteDeviceUpgradeCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSiteDeviceUpgradeCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


