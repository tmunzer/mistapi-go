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

// checks if the AssetFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetFilter{}

// AssetFilter Asset Filter
type AssetFilter struct {
	ApMac *string `json:"ap_mac,omitempty"`
	Beam *int32 `json:"beam,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	// whether the asset filter is disabled
	Disasbled *bool `json:"disasbled,omitempty"`
	// eddystone uid namespace used to filter assets
	EddystoneUidNamespace *string `json:"eddystone_uid_namespace,omitempty"`
	// eddystone url used to filter assets
	EddystoneUrl *string `json:"eddystone_url,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	// ibeacon major value used to filter assets
	IbeaconMajor *int32 `json:"ibeacon_major,omitempty"`
	// ibeacon uuid used to filter assets
	IbeaconUuid *string `json:"ibeacon_uuid,omitempty"`
	Id *string `json:"id,omitempty"`
	// ble manufacturing-specific company-id used to filter assets
	MfgCompanyId *int32 `json:"mfg_company_id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	Rssi *int32 `json:"rssi,omitempty"`
	// ble service data uuid used to filter assets
	ServiceUuid *string `json:"service_uuid,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetFilter AssetFilter

// NewAssetFilter instantiates a new AssetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetFilter(name string) *AssetFilter {
	this := AssetFilter{}
	var disabled bool = false
	this.Disabled = &disabled
	this.Name = name
	return &this
}

// NewAssetFilterWithDefaults instantiates a new AssetFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetFilterWithDefaults() *AssetFilter {
	this := AssetFilter{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetApMac returns the ApMac field value if set, zero value otherwise.
func (o *AssetFilter) GetApMac() string {
	if o == nil || IsNil(o.ApMac) {
		var ret string
		return ret
	}
	return *o.ApMac
}

// GetApMacOk returns a tuple with the ApMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetApMacOk() (*string, bool) {
	if o == nil || IsNil(o.ApMac) {
		return nil, false
	}
	return o.ApMac, true
}

// HasApMac returns a boolean if a field has been set.
func (o *AssetFilter) HasApMac() bool {
	if o != nil && !IsNil(o.ApMac) {
		return true
	}

	return false
}

// SetApMac gets a reference to the given string and assigns it to the ApMac field.
func (o *AssetFilter) SetApMac(v string) {
	o.ApMac = &v
}

// GetBeam returns the Beam field value if set, zero value otherwise.
func (o *AssetFilter) GetBeam() int32 {
	if o == nil || IsNil(o.Beam) {
		var ret int32
		return ret
	}
	return *o.Beam
}

// GetBeamOk returns a tuple with the Beam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetBeamOk() (*int32, bool) {
	if o == nil || IsNil(o.Beam) {
		return nil, false
	}
	return o.Beam, true
}

// HasBeam returns a boolean if a field has been set.
func (o *AssetFilter) HasBeam() bool {
	if o != nil && !IsNil(o.Beam) {
		return true
	}

	return false
}

// SetBeam gets a reference to the given int32 and assigns it to the Beam field.
func (o *AssetFilter) SetBeam(v int32) {
	o.Beam = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *AssetFilter) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *AssetFilter) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *AssetFilter) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AssetFilter) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AssetFilter) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AssetFilter) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDisasbled returns the Disasbled field value if set, zero value otherwise.
func (o *AssetFilter) GetDisasbled() bool {
	if o == nil || IsNil(o.Disasbled) {
		var ret bool
		return ret
	}
	return *o.Disasbled
}

// GetDisasbledOk returns a tuple with the Disasbled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetDisasbledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disasbled) {
		return nil, false
	}
	return o.Disasbled, true
}

// HasDisasbled returns a boolean if a field has been set.
func (o *AssetFilter) HasDisasbled() bool {
	if o != nil && !IsNil(o.Disasbled) {
		return true
	}

	return false
}

// SetDisasbled gets a reference to the given bool and assigns it to the Disasbled field.
func (o *AssetFilter) SetDisasbled(v bool) {
	o.Disasbled = &v
}

// GetEddystoneUidNamespace returns the EddystoneUidNamespace field value if set, zero value otherwise.
func (o *AssetFilter) GetEddystoneUidNamespace() string {
	if o == nil || IsNil(o.EddystoneUidNamespace) {
		var ret string
		return ret
	}
	return *o.EddystoneUidNamespace
}

// GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetEddystoneUidNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUidNamespace) {
		return nil, false
	}
	return o.EddystoneUidNamespace, true
}

// HasEddystoneUidNamespace returns a boolean if a field has been set.
func (o *AssetFilter) HasEddystoneUidNamespace() bool {
	if o != nil && !IsNil(o.EddystoneUidNamespace) {
		return true
	}

	return false
}

// SetEddystoneUidNamespace gets a reference to the given string and assigns it to the EddystoneUidNamespace field.
func (o *AssetFilter) SetEddystoneUidNamespace(v string) {
	o.EddystoneUidNamespace = &v
}

// GetEddystoneUrl returns the EddystoneUrl field value if set, zero value otherwise.
func (o *AssetFilter) GetEddystoneUrl() string {
	if o == nil || IsNil(o.EddystoneUrl) {
		var ret string
		return ret
	}
	return *o.EddystoneUrl
}

// GetEddystoneUrlOk returns a tuple with the EddystoneUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetEddystoneUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EddystoneUrl) {
		return nil, false
	}
	return o.EddystoneUrl, true
}

// HasEddystoneUrl returns a boolean if a field has been set.
func (o *AssetFilter) HasEddystoneUrl() bool {
	if o != nil && !IsNil(o.EddystoneUrl) {
		return true
	}

	return false
}

// SetEddystoneUrl gets a reference to the given string and assigns it to the EddystoneUrl field.
func (o *AssetFilter) SetEddystoneUrl(v string) {
	o.EddystoneUrl = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *AssetFilter) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *AssetFilter) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *AssetFilter) SetForSite(v bool) {
	o.ForSite = &v
}

// GetIbeaconMajor returns the IbeaconMajor field value if set, zero value otherwise.
func (o *AssetFilter) GetIbeaconMajor() int32 {
	if o == nil || IsNil(o.IbeaconMajor) {
		var ret int32
		return ret
	}
	return *o.IbeaconMajor
}

// GetIbeaconMajorOk returns a tuple with the IbeaconMajor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetIbeaconMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.IbeaconMajor) {
		return nil, false
	}
	return o.IbeaconMajor, true
}

// HasIbeaconMajor returns a boolean if a field has been set.
func (o *AssetFilter) HasIbeaconMajor() bool {
	if o != nil && !IsNil(o.IbeaconMajor) {
		return true
	}

	return false
}

// SetIbeaconMajor gets a reference to the given int32 and assigns it to the IbeaconMajor field.
func (o *AssetFilter) SetIbeaconMajor(v int32) {
	o.IbeaconMajor = &v
}

// GetIbeaconUuid returns the IbeaconUuid field value if set, zero value otherwise.
func (o *AssetFilter) GetIbeaconUuid() string {
	if o == nil || IsNil(o.IbeaconUuid) {
		var ret string
		return ret
	}
	return *o.IbeaconUuid
}

// GetIbeaconUuidOk returns a tuple with the IbeaconUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetIbeaconUuidOk() (*string, bool) {
	if o == nil || IsNil(o.IbeaconUuid) {
		return nil, false
	}
	return o.IbeaconUuid, true
}

// HasIbeaconUuid returns a boolean if a field has been set.
func (o *AssetFilter) HasIbeaconUuid() bool {
	if o != nil && !IsNil(o.IbeaconUuid) {
		return true
	}

	return false
}

// SetIbeaconUuid gets a reference to the given string and assigns it to the IbeaconUuid field.
func (o *AssetFilter) SetIbeaconUuid(v string) {
	o.IbeaconUuid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetFilter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetFilter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetFilter) SetId(v string) {
	o.Id = &v
}

// GetMfgCompanyId returns the MfgCompanyId field value if set, zero value otherwise.
func (o *AssetFilter) GetMfgCompanyId() int32 {
	if o == nil || IsNil(o.MfgCompanyId) {
		var ret int32
		return ret
	}
	return *o.MfgCompanyId
}

// GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetMfgCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MfgCompanyId) {
		return nil, false
	}
	return o.MfgCompanyId, true
}

// HasMfgCompanyId returns a boolean if a field has been set.
func (o *AssetFilter) HasMfgCompanyId() bool {
	if o != nil && !IsNil(o.MfgCompanyId) {
		return true
	}

	return false
}

// SetMfgCompanyId gets a reference to the given int32 and assigns it to the MfgCompanyId field.
func (o *AssetFilter) SetMfgCompanyId(v int32) {
	o.MfgCompanyId = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *AssetFilter) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *AssetFilter) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *AssetFilter) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *AssetFilter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssetFilter) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AssetFilter) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AssetFilter) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *AssetFilter) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *AssetFilter) GetRssi() int32 {
	if o == nil || IsNil(o.Rssi) {
		var ret int32
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetRssiOk() (*int32, bool) {
	if o == nil || IsNil(o.Rssi) {
		return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *AssetFilter) HasRssi() bool {
	if o != nil && !IsNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given int32 and assigns it to the Rssi field.
func (o *AssetFilter) SetRssi(v int32) {
	o.Rssi = &v
}

// GetServiceUuid returns the ServiceUuid field value if set, zero value otherwise.
func (o *AssetFilter) GetServiceUuid() string {
	if o == nil || IsNil(o.ServiceUuid) {
		var ret string
		return ret
	}
	return *o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetServiceUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUuid) {
		return nil, false
	}
	return o.ServiceUuid, true
}

// HasServiceUuid returns a boolean if a field has been set.
func (o *AssetFilter) HasServiceUuid() bool {
	if o != nil && !IsNil(o.ServiceUuid) {
		return true
	}

	return false
}

// SetServiceUuid gets a reference to the given string and assigns it to the ServiceUuid field.
func (o *AssetFilter) SetServiceUuid(v string) {
	o.ServiceUuid = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *AssetFilter) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFilter) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *AssetFilter) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *AssetFilter) SetSiteId(v string) {
	o.SiteId = &v
}

func (o AssetFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApMac) {
		toSerialize["ap_mac"] = o.ApMac
	}
	if !IsNil(o.Beam) {
		toSerialize["beam"] = o.Beam
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Disasbled) {
		toSerialize["disasbled"] = o.Disasbled
	}
	if !IsNil(o.EddystoneUidNamespace) {
		toSerialize["eddystone_uid_namespace"] = o.EddystoneUidNamespace
	}
	if !IsNil(o.EddystoneUrl) {
		toSerialize["eddystone_url"] = o.EddystoneUrl
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.IbeaconMajor) {
		toSerialize["ibeacon_major"] = o.IbeaconMajor
	}
	if !IsNil(o.IbeaconUuid) {
		toSerialize["ibeacon_uuid"] = o.IbeaconUuid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MfgCompanyId) {
		toSerialize["mfg_company_id"] = o.MfgCompanyId
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	if !IsNil(o.ServiceUuid) {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAssetFilter := _AssetFilter{}

	err = json.Unmarshal(data, &varAssetFilter)

	if err != nil {
		return err
	}

	*o = AssetFilter(varAssetFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_mac")
		delete(additionalProperties, "beam")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "disasbled")
		delete(additionalProperties, "eddystone_uid_namespace")
		delete(additionalProperties, "eddystone_url")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "ibeacon_major")
		delete(additionalProperties, "ibeacon_uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mfg_company_id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "rssi")
		delete(additionalProperties, "service_uuid")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetFilter struct {
	value *AssetFilter
	isSet bool
}

func (v NullableAssetFilter) Get() *AssetFilter {
	return v.value
}

func (v *NullableAssetFilter) Set(val *AssetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetFilter(val *AssetFilter) *NullableAssetFilter {
	return &NullableAssetFilter{value: val, isSet: true}
}

func (v NullableAssetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


