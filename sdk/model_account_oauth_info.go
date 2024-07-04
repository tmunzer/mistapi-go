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
)

// checks if the AccountOauthInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountOauthInfo{}

// AccountOauthInfo OAuth linked apps (zoom/teams/intune) account details
type AccountOauthInfo struct {
	// Linked app(zoom/teams/intune) account id
	AccountId *string `json:"account_id,omitempty"`
	// Name of the company whose account mist has subscribed to
	Company *string `json:"company,omitempty"`
	// This error is provided when the account fails to fetch token/data
	Error *string `json:"error,omitempty"`
	Errors []string `json:"errors,omitempty"`
	// Is the last data pull for account is successful or not
	LastStatus *string `json:"last_status,omitempty"`
	// Last data pull timestamp, background jobs that pull account data
	LastSync *int32 `json:"last_sync,omitempty"`
	// First name of the user who linked the account
	LinkedBy *string `json:"linked_by,omitempty"`
	// Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/
	MaxDailyApiRequests *int32 `json:"max_daily_api_requests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountOauthInfo AccountOauthInfo

// NewAccountOauthInfo instantiates a new AccountOauthInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOauthInfo() *AccountOauthInfo {
	this := AccountOauthInfo{}
	return &this
}

// NewAccountOauthInfoWithDefaults instantiates a new AccountOauthInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOauthInfoWithDefaults() *AccountOauthInfo {
	this := AccountOauthInfo{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccountOauthInfo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AccountOauthInfo) SetCompany(v string) {
	o.Company = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AccountOauthInfo) SetError(v string) {
	o.Error = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AccountOauthInfo) SetErrors(v []string) {
	o.Errors = v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetLastStatus() string {
	if o == nil || IsNil(o.LastStatus) {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetLastStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastStatus) {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasLastStatus() bool {
	if o != nil && !IsNil(o.LastStatus) {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *AccountOauthInfo) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetLastSync() int32 {
	if o == nil || IsNil(o.LastSync) {
		var ret int32
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetLastSyncOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given int32 and assigns it to the LastSync field.
func (o *AccountOauthInfo) SetLastSync(v int32) {
	o.LastSync = &v
}

// GetLinkedBy returns the LinkedBy field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetLinkedBy() string {
	if o == nil || IsNil(o.LinkedBy) {
		var ret string
		return ret
	}
	return *o.LinkedBy
}

// GetLinkedByOk returns a tuple with the LinkedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetLinkedByOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedBy) {
		return nil, false
	}
	return o.LinkedBy, true
}

// HasLinkedBy returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasLinkedBy() bool {
	if o != nil && !IsNil(o.LinkedBy) {
		return true
	}

	return false
}

// SetLinkedBy gets a reference to the given string and assigns it to the LinkedBy field.
func (o *AccountOauthInfo) SetLinkedBy(v string) {
	o.LinkedBy = &v
}

// GetMaxDailyApiRequests returns the MaxDailyApiRequests field value if set, zero value otherwise.
func (o *AccountOauthInfo) GetMaxDailyApiRequests() int32 {
	if o == nil || IsNil(o.MaxDailyApiRequests) {
		var ret int32
		return ret
	}
	return *o.MaxDailyApiRequests
}

// GetMaxDailyApiRequestsOk returns a tuple with the MaxDailyApiRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOauthInfo) GetMaxDailyApiRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDailyApiRequests) {
		return nil, false
	}
	return o.MaxDailyApiRequests, true
}

// HasMaxDailyApiRequests returns a boolean if a field has been set.
func (o *AccountOauthInfo) HasMaxDailyApiRequests() bool {
	if o != nil && !IsNil(o.MaxDailyApiRequests) {
		return true
	}

	return false
}

// SetMaxDailyApiRequests gets a reference to the given int32 and assigns it to the MaxDailyApiRequests field.
func (o *AccountOauthInfo) SetMaxDailyApiRequests(v int32) {
	o.MaxDailyApiRequests = &v
}

func (o AccountOauthInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountOauthInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
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
	if !IsNil(o.MaxDailyApiRequests) {
		toSerialize["max_daily_api_requests"] = o.MaxDailyApiRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountOauthInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountOauthInfo := _AccountOauthInfo{}

	err = json.Unmarshal(data, &varAccountOauthInfo)

	if err != nil {
		return err
	}

	*o = AccountOauthInfo(varAccountOauthInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "company")
		delete(additionalProperties, "error")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "last_status")
		delete(additionalProperties, "last_sync")
		delete(additionalProperties, "linked_by")
		delete(additionalProperties, "max_daily_api_requests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountOauthInfo struct {
	value *AccountOauthInfo
	isSet bool
}

func (v NullableAccountOauthInfo) Get() *AccountOauthInfo {
	return v.value
}

func (v *NullableAccountOauthInfo) Set(val *AccountOauthInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOauthInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOauthInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOauthInfo(val *AccountOauthInfo) *NullableAccountOauthInfo {
	return &NullableAccountOauthInfo{value: val, isSet: true}
}

func (v NullableAccountOauthInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOauthInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


