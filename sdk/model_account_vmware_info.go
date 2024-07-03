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

// checks if the AccountVmwareInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountVmwareInfo{}

// AccountVmwareInfo struct for AccountVmwareInfo
type AccountVmwareInfo struct {
	AccountId *string `json:"account_id,omitempty"`
	// Linked VMware Instance URL
	InstanceUrl *string `json:"instance_url,omitempty"`
	// Is the last data pull for VMware account is successful or not
	LastStatus *string `json:"last_status,omitempty"`
	// Last data pull timestamp, background jobs that pull VMware account data
	LastSync *int32 `json:"last_sync,omitempty"`
	// First name of the user who linked the VMware account
	LinkedBy *string `json:"linked_by,omitempty"`
	// This error is provided when the VMware account fails to fetch token/data
	LinkedTimestamp *int32 `json:"linked_timestamp,omitempty"`
	// Name of the company whose VMware account mist has subscribed to
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountVmwareInfo AccountVmwareInfo

// NewAccountVmwareInfo instantiates a new AccountVmwareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountVmwareInfo() *AccountVmwareInfo {
	this := AccountVmwareInfo{}
	return &this
}

// NewAccountVmwareInfoWithDefaults instantiates a new AccountVmwareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountVmwareInfoWithDefaults() *AccountVmwareInfo {
	this := AccountVmwareInfo{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccountVmwareInfo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetInstanceUrl returns the InstanceUrl field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetInstanceUrl() string {
	if o == nil || IsNil(o.InstanceUrl) {
		var ret string
		return ret
	}
	return *o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetInstanceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceUrl) {
		return nil, false
	}
	return o.InstanceUrl, true
}

// HasInstanceUrl returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasInstanceUrl() bool {
	if o != nil && !IsNil(o.InstanceUrl) {
		return true
	}

	return false
}

// SetInstanceUrl gets a reference to the given string and assigns it to the InstanceUrl field.
func (o *AccountVmwareInfo) SetInstanceUrl(v string) {
	o.InstanceUrl = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetLastStatus() string {
	if o == nil || IsNil(o.LastStatus) {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetLastStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastStatus) {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasLastStatus() bool {
	if o != nil && !IsNil(o.LastStatus) {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *AccountVmwareInfo) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetLastSync() int32 {
	if o == nil || IsNil(o.LastSync) {
		var ret int32
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetLastSyncOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given int32 and assigns it to the LastSync field.
func (o *AccountVmwareInfo) SetLastSync(v int32) {
	o.LastSync = &v
}

// GetLinkedBy returns the LinkedBy field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetLinkedBy() string {
	if o == nil || IsNil(o.LinkedBy) {
		var ret string
		return ret
	}
	return *o.LinkedBy
}

// GetLinkedByOk returns a tuple with the LinkedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetLinkedByOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedBy) {
		return nil, false
	}
	return o.LinkedBy, true
}

// HasLinkedBy returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasLinkedBy() bool {
	if o != nil && !IsNil(o.LinkedBy) {
		return true
	}

	return false
}

// SetLinkedBy gets a reference to the given string and assigns it to the LinkedBy field.
func (o *AccountVmwareInfo) SetLinkedBy(v string) {
	o.LinkedBy = &v
}

// GetLinkedTimestamp returns the LinkedTimestamp field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetLinkedTimestamp() int32 {
	if o == nil || IsNil(o.LinkedTimestamp) {
		var ret int32
		return ret
	}
	return *o.LinkedTimestamp
}

// GetLinkedTimestampOk returns a tuple with the LinkedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetLinkedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkedTimestamp) {
		return nil, false
	}
	return o.LinkedTimestamp, true
}

// HasLinkedTimestamp returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasLinkedTimestamp() bool {
	if o != nil && !IsNil(o.LinkedTimestamp) {
		return true
	}

	return false
}

// SetLinkedTimestamp gets a reference to the given int32 and assigns it to the LinkedTimestamp field.
func (o *AccountVmwareInfo) SetLinkedTimestamp(v int32) {
	o.LinkedTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountVmwareInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountVmwareInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountVmwareInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountVmwareInfo) SetName(v string) {
	o.Name = &v
}

func (o AccountVmwareInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountVmwareInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.InstanceUrl) {
		toSerialize["instance_url"] = o.InstanceUrl
	}
	if !IsNil(o.LastStatus) {
		toSerialize["last_status"] = o.LastStatus
	}
	if !IsNil(o.LastSync) {
		toSerialize["last_sync"] = o.LastSync
	}
	if !IsNil(o.LinkedBy) {
		toSerialize["linked_by"] = o.LinkedBy
	}
	if !IsNil(o.LinkedTimestamp) {
		toSerialize["linked_timestamp"] = o.LinkedTimestamp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountVmwareInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountVmwareInfo := _AccountVmwareInfo{}

	err = json.Unmarshal(data, &varAccountVmwareInfo)

	if err != nil {
		return err
	}

	*o = AccountVmwareInfo(varAccountVmwareInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "instance_url")
		delete(additionalProperties, "last_status")
		delete(additionalProperties, "last_sync")
		delete(additionalProperties, "linked_by")
		delete(additionalProperties, "linked_timestamp")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountVmwareInfo struct {
	value *AccountVmwareInfo
	isSet bool
}

func (v NullableAccountVmwareInfo) Get() *AccountVmwareInfo {
	return v.value
}

func (v *NullableAccountVmwareInfo) Set(val *AccountVmwareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountVmwareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountVmwareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountVmwareInfo(val *AccountVmwareInfo) *NullableAccountVmwareInfo {
	return &NullableAccountVmwareInfo{value: val, isSet: true}
}

func (v NullableAccountVmwareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountVmwareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


