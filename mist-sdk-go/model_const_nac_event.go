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

// checks if the ConstNacEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstNacEvent{}

// ConstNacEvent struct for ConstNacEvent
type ConstNacEvent struct {
	Ap *string `json:"ap,omitempty"`
	Bssid *string `json:"bssid,omitempty"`
	CertCn *string `json:"cert_cn,omitempty"`
	CertExpiry *int32 `json:"cert_expiry,omitempty"`
	CertIssuer *string `json:"cert_issuer,omitempty"`
	CertSanUpn []string `json:"cert_san_upn,omitempty"`
	CertSerial *string `json:"cert_serial,omitempty"`
	CertSubject *string `json:"cert_subject,omitempty"`
	EapType *string `json:"eap_type,omitempty"`
	NasVendor *string `json:"nas_vendor,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	RandomMac *bool `json:"random_mac,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	Ssid *string `json:"ssid,omitempty"`
	Timestamp *float32 `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
	Username *string `json:"username,omitempty"`
	Wcid *string `json:"wcid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstNacEvent ConstNacEvent

// NewConstNacEvent instantiates a new ConstNacEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstNacEvent() *ConstNacEvent {
	this := ConstNacEvent{}
	return &this
}

// NewConstNacEventWithDefaults instantiates a new ConstNacEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstNacEventWithDefaults() *ConstNacEvent {
	this := ConstNacEvent{}
	return &this
}

// GetAp returns the Ap field value if set, zero value otherwise.
func (o *ConstNacEvent) GetAp() string {
	if o == nil || IsNil(o.Ap) {
		var ret string
		return ret
	}
	return *o.Ap
}

// GetApOk returns a tuple with the Ap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetApOk() (*string, bool) {
	if o == nil || IsNil(o.Ap) {
		return nil, false
	}
	return o.Ap, true
}

// HasAp returns a boolean if a field has been set.
func (o *ConstNacEvent) HasAp() bool {
	if o != nil && !IsNil(o.Ap) {
		return true
	}

	return false
}

// SetAp gets a reference to the given string and assigns it to the Ap field.
func (o *ConstNacEvent) SetAp(v string) {
	o.Ap = &v
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *ConstNacEvent) GetBssid() string {
	if o == nil || IsNil(o.Bssid) {
		var ret string
		return ret
	}
	return *o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetBssidOk() (*string, bool) {
	if o == nil || IsNil(o.Bssid) {
		return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *ConstNacEvent) HasBssid() bool {
	if o != nil && !IsNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given string and assigns it to the Bssid field.
func (o *ConstNacEvent) SetBssid(v string) {
	o.Bssid = &v
}

// GetCertCn returns the CertCn field value if set, zero value otherwise.
func (o *ConstNacEvent) GetCertCn() string {
	if o == nil || IsNil(o.CertCn) {
		var ret string
		return ret
	}
	return *o.CertCn
}

// GetCertCnOk returns a tuple with the CertCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetCertCnOk() (*string, bool) {
	if o == nil || IsNil(o.CertCn) {
		return nil, false
	}
	return o.CertCn, true
}

// HasCertCn returns a boolean if a field has been set.
func (o *ConstNacEvent) HasCertCn() bool {
	if o != nil && !IsNil(o.CertCn) {
		return true
	}

	return false
}

// SetCertCn gets a reference to the given string and assigns it to the CertCn field.
func (o *ConstNacEvent) SetCertCn(v string) {
	o.CertCn = &v
}

// GetCertExpiry returns the CertExpiry field value if set, zero value otherwise.
func (o *ConstNacEvent) GetCertExpiry() int32 {
	if o == nil || IsNil(o.CertExpiry) {
		var ret int32
		return ret
	}
	return *o.CertExpiry
}

// GetCertExpiryOk returns a tuple with the CertExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetCertExpiryOk() (*int32, bool) {
	if o == nil || IsNil(o.CertExpiry) {
		return nil, false
	}
	return o.CertExpiry, true
}

// HasCertExpiry returns a boolean if a field has been set.
func (o *ConstNacEvent) HasCertExpiry() bool {
	if o != nil && !IsNil(o.CertExpiry) {
		return true
	}

	return false
}

// SetCertExpiry gets a reference to the given int32 and assigns it to the CertExpiry field.
func (o *ConstNacEvent) SetCertExpiry(v int32) {
	o.CertExpiry = &v
}

// GetCertIssuer returns the CertIssuer field value if set, zero value otherwise.
func (o *ConstNacEvent) GetCertIssuer() string {
	if o == nil || IsNil(o.CertIssuer) {
		var ret string
		return ret
	}
	return *o.CertIssuer
}

// GetCertIssuerOk returns a tuple with the CertIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetCertIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.CertIssuer) {
		return nil, false
	}
	return o.CertIssuer, true
}

// HasCertIssuer returns a boolean if a field has been set.
func (o *ConstNacEvent) HasCertIssuer() bool {
	if o != nil && !IsNil(o.CertIssuer) {
		return true
	}

	return false
}

// SetCertIssuer gets a reference to the given string and assigns it to the CertIssuer field.
func (o *ConstNacEvent) SetCertIssuer(v string) {
	o.CertIssuer = &v
}

// GetCertSanUpn returns the CertSanUpn field value if set, zero value otherwise.
func (o *ConstNacEvent) GetCertSanUpn() []string {
	if o == nil || IsNil(o.CertSanUpn) {
		var ret []string
		return ret
	}
	return o.CertSanUpn
}

// GetCertSanUpnOk returns a tuple with the CertSanUpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetCertSanUpnOk() ([]string, bool) {
	if o == nil || IsNil(o.CertSanUpn) {
		return nil, false
	}
	return o.CertSanUpn, true
}

// HasCertSanUpn returns a boolean if a field has been set.
func (o *ConstNacEvent) HasCertSanUpn() bool {
	if o != nil && !IsNil(o.CertSanUpn) {
		return true
	}

	return false
}

// SetCertSanUpn gets a reference to the given []string and assigns it to the CertSanUpn field.
func (o *ConstNacEvent) SetCertSanUpn(v []string) {
	o.CertSanUpn = v
}

// GetCertSerial returns the CertSerial field value if set, zero value otherwise.
func (o *ConstNacEvent) GetCertSerial() string {
	if o == nil || IsNil(o.CertSerial) {
		var ret string
		return ret
	}
	return *o.CertSerial
}

// GetCertSerialOk returns a tuple with the CertSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetCertSerialOk() (*string, bool) {
	if o == nil || IsNil(o.CertSerial) {
		return nil, false
	}
	return o.CertSerial, true
}

// HasCertSerial returns a boolean if a field has been set.
func (o *ConstNacEvent) HasCertSerial() bool {
	if o != nil && !IsNil(o.CertSerial) {
		return true
	}

	return false
}

// SetCertSerial gets a reference to the given string and assigns it to the CertSerial field.
func (o *ConstNacEvent) SetCertSerial(v string) {
	o.CertSerial = &v
}

// GetCertSubject returns the CertSubject field value if set, zero value otherwise.
func (o *ConstNacEvent) GetCertSubject() string {
	if o == nil || IsNil(o.CertSubject) {
		var ret string
		return ret
	}
	return *o.CertSubject
}

// GetCertSubjectOk returns a tuple with the CertSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetCertSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.CertSubject) {
		return nil, false
	}
	return o.CertSubject, true
}

// HasCertSubject returns a boolean if a field has been set.
func (o *ConstNacEvent) HasCertSubject() bool {
	if o != nil && !IsNil(o.CertSubject) {
		return true
	}

	return false
}

// SetCertSubject gets a reference to the given string and assigns it to the CertSubject field.
func (o *ConstNacEvent) SetCertSubject(v string) {
	o.CertSubject = &v
}

// GetEapType returns the EapType field value if set, zero value otherwise.
func (o *ConstNacEvent) GetEapType() string {
	if o == nil || IsNil(o.EapType) {
		var ret string
		return ret
	}
	return *o.EapType
}

// GetEapTypeOk returns a tuple with the EapType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetEapTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EapType) {
		return nil, false
	}
	return o.EapType, true
}

// HasEapType returns a boolean if a field has been set.
func (o *ConstNacEvent) HasEapType() bool {
	if o != nil && !IsNil(o.EapType) {
		return true
	}

	return false
}

// SetEapType gets a reference to the given string and assigns it to the EapType field.
func (o *ConstNacEvent) SetEapType(v string) {
	o.EapType = &v
}

// GetNasVendor returns the NasVendor field value if set, zero value otherwise.
func (o *ConstNacEvent) GetNasVendor() string {
	if o == nil || IsNil(o.NasVendor) {
		var ret string
		return ret
	}
	return *o.NasVendor
}

// GetNasVendorOk returns a tuple with the NasVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetNasVendorOk() (*string, bool) {
	if o == nil || IsNil(o.NasVendor) {
		return nil, false
	}
	return o.NasVendor, true
}

// HasNasVendor returns a boolean if a field has been set.
func (o *ConstNacEvent) HasNasVendor() bool {
	if o != nil && !IsNil(o.NasVendor) {
		return true
	}

	return false
}

// SetNasVendor gets a reference to the given string and assigns it to the NasVendor field.
func (o *ConstNacEvent) SetNasVendor(v string) {
	o.NasVendor = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ConstNacEvent) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ConstNacEvent) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ConstNacEvent) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRandomMac returns the RandomMac field value if set, zero value otherwise.
func (o *ConstNacEvent) GetRandomMac() bool {
	if o == nil || IsNil(o.RandomMac) {
		var ret bool
		return ret
	}
	return *o.RandomMac
}

// GetRandomMacOk returns a tuple with the RandomMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetRandomMacOk() (*bool, bool) {
	if o == nil || IsNil(o.RandomMac) {
		return nil, false
	}
	return o.RandomMac, true
}

// HasRandomMac returns a boolean if a field has been set.
func (o *ConstNacEvent) HasRandomMac() bool {
	if o != nil && !IsNil(o.RandomMac) {
		return true
	}

	return false
}

// SetRandomMac gets a reference to the given bool and assigns it to the RandomMac field.
func (o *ConstNacEvent) SetRandomMac(v bool) {
	o.RandomMac = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ConstNacEvent) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ConstNacEvent) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *ConstNacEvent) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *ConstNacEvent) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *ConstNacEvent) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *ConstNacEvent) SetSsid(v string) {
	o.Ssid = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ConstNacEvent) GetTimestamp() float32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret float32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ConstNacEvent) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float32 and assigns it to the Timestamp field.
func (o *ConstNacEvent) SetTimestamp(v float32) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConstNacEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConstNacEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConstNacEvent) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ConstNacEvent) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ConstNacEvent) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ConstNacEvent) SetUsername(v string) {
	o.Username = &v
}

// GetWcid returns the Wcid field value if set, zero value otherwise.
func (o *ConstNacEvent) GetWcid() string {
	if o == nil || IsNil(o.Wcid) {
		var ret string
		return ret
	}
	return *o.Wcid
}

// GetWcidOk returns a tuple with the Wcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstNacEvent) GetWcidOk() (*string, bool) {
	if o == nil || IsNil(o.Wcid) {
		return nil, false
	}
	return o.Wcid, true
}

// HasWcid returns a boolean if a field has been set.
func (o *ConstNacEvent) HasWcid() bool {
	if o != nil && !IsNil(o.Wcid) {
		return true
	}

	return false
}

// SetWcid gets a reference to the given string and assigns it to the Wcid field.
func (o *ConstNacEvent) SetWcid(v string) {
	o.Wcid = &v
}

func (o ConstNacEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstNacEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ap) {
		toSerialize["ap"] = o.Ap
	}
	if !IsNil(o.Bssid) {
		toSerialize["bssid"] = o.Bssid
	}
	if !IsNil(o.CertCn) {
		toSerialize["cert_cn"] = o.CertCn
	}
	if !IsNil(o.CertExpiry) {
		toSerialize["cert_expiry"] = o.CertExpiry
	}
	if !IsNil(o.CertIssuer) {
		toSerialize["cert_issuer"] = o.CertIssuer
	}
	if !IsNil(o.CertSanUpn) {
		toSerialize["cert_san_upn"] = o.CertSanUpn
	}
	if !IsNil(o.CertSerial) {
		toSerialize["cert_serial"] = o.CertSerial
	}
	if !IsNil(o.CertSubject) {
		toSerialize["cert_subject"] = o.CertSubject
	}
	if !IsNil(o.EapType) {
		toSerialize["eap_type"] = o.EapType
	}
	if !IsNil(o.NasVendor) {
		toSerialize["nas_vendor"] = o.NasVendor
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.RandomMac) {
		toSerialize["random_mac"] = o.RandomMac
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Wcid) {
		toSerialize["wcid"] = o.Wcid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstNacEvent) UnmarshalJSON(data []byte) (err error) {
	varConstNacEvent := _ConstNacEvent{}

	err = json.Unmarshal(data, &varConstNacEvent)

	if err != nil {
		return err
	}

	*o = ConstNacEvent(varConstNacEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap")
		delete(additionalProperties, "bssid")
		delete(additionalProperties, "cert_cn")
		delete(additionalProperties, "cert_expiry")
		delete(additionalProperties, "cert_issuer")
		delete(additionalProperties, "cert_san_upn")
		delete(additionalProperties, "cert_serial")
		delete(additionalProperties, "cert_subject")
		delete(additionalProperties, "eap_type")
		delete(additionalProperties, "nas_vendor")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "random_mac")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "type")
		delete(additionalProperties, "username")
		delete(additionalProperties, "wcid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstNacEvent struct {
	value *ConstNacEvent
	isSet bool
}

func (v NullableConstNacEvent) Get() *ConstNacEvent {
	return v.value
}

func (v *NullableConstNacEvent) Set(val *ConstNacEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableConstNacEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableConstNacEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstNacEvent(val *ConstNacEvent) *NullableConstNacEvent {
	return &NullableConstNacEvent{value: val, isSet: true}
}

func (v NullableConstNacEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstNacEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


