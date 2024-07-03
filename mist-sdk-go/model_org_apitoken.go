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

// checks if the OrgApitoken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgApitoken{}

// OrgApitoken Org API Token
type OrgApitoken struct {
	// email of the token creator / null if creator is deleted
	CreatedBy NullableString `json:"created_by,omitempty"`
	CreatedTime *int32 `json:"created_time,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	LastUsed NullableInt32 `json:"last_used,omitempty"`
	// name of the token
	Name NullableString `json:"name,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	// list of privileges the token has on the orgs/sites
	Privileges []PrivilegeOrg `json:"privileges,omitempty"`
	// to restrict where the API can be called from
	SrcIps []string `json:"src_ips,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgApitoken OrgApitoken

// NewOrgApitoken instantiates a new OrgApitoken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgApitoken() *OrgApitoken {
	this := OrgApitoken{}
	return &this
}

// NewOrgApitokenWithDefaults instantiates a new OrgApitoken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgApitokenWithDefaults() *OrgApitoken {
	this := OrgApitoken{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgApitoken) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgApitoken) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OrgApitoken) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *OrgApitoken) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *OrgApitoken) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *OrgApitoken) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *OrgApitoken) GetCreatedTime() int32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret int32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgApitoken) GetCreatedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *OrgApitoken) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int32 and assigns it to the CreatedTime field.
func (o *OrgApitoken) SetCreatedTime(v int32) {
	o.CreatedTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgApitoken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgApitoken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgApitoken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgApitoken) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OrgApitoken) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgApitoken) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OrgApitoken) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OrgApitoken) SetKey(v string) {
	o.Key = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgApitoken) GetLastUsed() int32 {
	if o == nil || IsNil(o.LastUsed.Get()) {
		var ret int32
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgApitoken) GetLastUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// HasLastUsed returns a boolean if a field has been set.
func (o *OrgApitoken) HasLastUsed() bool {
	if o != nil && o.LastUsed.IsSet() {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given NullableInt32 and assigns it to the LastUsed field.
func (o *OrgApitoken) SetLastUsed(v int32) {
	o.LastUsed.Set(&v)
}
// SetLastUsedNil sets the value for LastUsed to be an explicit nil
func (o *OrgApitoken) SetLastUsedNil() {
	o.LastUsed.Set(nil)
}

// UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
func (o *OrgApitoken) UnsetLastUsed() {
	o.LastUsed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgApitoken) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgApitoken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrgApitoken) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrgApitoken) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OrgApitoken) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrgApitoken) UnsetName() {
	o.Name.Unset()
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrgApitoken) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgApitoken) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrgApitoken) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrgApitoken) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *OrgApitoken) GetPrivileges() []PrivilegeOrg {
	if o == nil || IsNil(o.Privileges) {
		var ret []PrivilegeOrg
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgApitoken) GetPrivilegesOk() ([]PrivilegeOrg, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *OrgApitoken) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []PrivilegeOrg and assigns it to the Privileges field.
func (o *OrgApitoken) SetPrivileges(v []PrivilegeOrg) {
	o.Privileges = v
}

// GetSrcIps returns the SrcIps field value if set, zero value otherwise.
func (o *OrgApitoken) GetSrcIps() []string {
	if o == nil || IsNil(o.SrcIps) {
		var ret []string
		return ret
	}
	return o.SrcIps
}

// GetSrcIpsOk returns a tuple with the SrcIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgApitoken) GetSrcIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.SrcIps) {
		return nil, false
	}
	return o.SrcIps, true
}

// HasSrcIps returns a boolean if a field has been set.
func (o *OrgApitoken) HasSrcIps() bool {
	if o != nil && !IsNil(o.SrcIps) {
		return true
	}

	return false
}

// SetSrcIps gets a reference to the given []string and assigns it to the SrcIps field.
func (o *OrgApitoken) SetSrcIps(v []string) {
	o.SrcIps = v
}

func (o OrgApitoken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgApitoken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.LastUsed.IsSet() {
		toSerialize["last_used"] = o.LastUsed.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !IsNil(o.SrcIps) {
		toSerialize["src_ips"] = o.SrcIps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgApitoken) UnmarshalJSON(data []byte) (err error) {
	varOrgApitoken := _OrgApitoken{}

	err = json.Unmarshal(data, &varOrgApitoken)

	if err != nil {
		return err
	}

	*o = OrgApitoken(varOrgApitoken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "last_used")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "privileges")
		delete(additionalProperties, "src_ips")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgApitoken struct {
	value *OrgApitoken
	isSet bool
}

func (v NullableOrgApitoken) Get() *OrgApitoken {
	return v.value
}

func (v *NullableOrgApitoken) Set(val *OrgApitoken) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgApitoken) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgApitoken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgApitoken(val *OrgApitoken) *NullableOrgApitoken {
	return &NullableOrgApitoken{value: val, isSet: true}
}

func (v NullableOrgApitoken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgApitoken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

