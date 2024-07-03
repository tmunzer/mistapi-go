# ConstNacEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** |  | [optional] 
**Bssid** | Pointer to **string** |  | [optional] 
**CertCn** | Pointer to **string** |  | [optional] 
**CertExpiry** | Pointer to **int32** |  | [optional] 
**CertIssuer** | Pointer to **string** |  | [optional] 
**CertSanUpn** | Pointer to **[]string** |  | [optional] 
**CertSerial** | Pointer to **string** |  | [optional] 
**CertSubject** | Pointer to **string** |  | [optional] 
**EapType** | Pointer to **string** |  | [optional] 
**NasVendor** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RandomMac** | Pointer to **bool** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Ssid** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Wcid** | Pointer to **string** |  | [optional] 

## Methods

### NewConstNacEvent

`func NewConstNacEvent() *ConstNacEvent`

NewConstNacEvent instantiates a new ConstNacEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstNacEventWithDefaults

`func NewConstNacEventWithDefaults() *ConstNacEvent`

NewConstNacEventWithDefaults instantiates a new ConstNacEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *ConstNacEvent) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *ConstNacEvent) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *ConstNacEvent) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *ConstNacEvent) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetBssid

`func (o *ConstNacEvent) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *ConstNacEvent) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *ConstNacEvent) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *ConstNacEvent) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetCertCn

`func (o *ConstNacEvent) GetCertCn() string`

GetCertCn returns the CertCn field if non-nil, zero value otherwise.

### GetCertCnOk

`func (o *ConstNacEvent) GetCertCnOk() (*string, bool)`

GetCertCnOk returns a tuple with the CertCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertCn

`func (o *ConstNacEvent) SetCertCn(v string)`

SetCertCn sets CertCn field to given value.

### HasCertCn

`func (o *ConstNacEvent) HasCertCn() bool`

HasCertCn returns a boolean if a field has been set.

### GetCertExpiry

`func (o *ConstNacEvent) GetCertExpiry() int32`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *ConstNacEvent) GetCertExpiryOk() (*int32, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *ConstNacEvent) SetCertExpiry(v int32)`

SetCertExpiry sets CertExpiry field to given value.

### HasCertExpiry

`func (o *ConstNacEvent) HasCertExpiry() bool`

HasCertExpiry returns a boolean if a field has been set.

### GetCertIssuer

`func (o *ConstNacEvent) GetCertIssuer() string`

GetCertIssuer returns the CertIssuer field if non-nil, zero value otherwise.

### GetCertIssuerOk

`func (o *ConstNacEvent) GetCertIssuerOk() (*string, bool)`

GetCertIssuerOk returns a tuple with the CertIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuer

`func (o *ConstNacEvent) SetCertIssuer(v string)`

SetCertIssuer sets CertIssuer field to given value.

### HasCertIssuer

`func (o *ConstNacEvent) HasCertIssuer() bool`

HasCertIssuer returns a boolean if a field has been set.

### GetCertSanUpn

`func (o *ConstNacEvent) GetCertSanUpn() []string`

GetCertSanUpn returns the CertSanUpn field if non-nil, zero value otherwise.

### GetCertSanUpnOk

`func (o *ConstNacEvent) GetCertSanUpnOk() (*[]string, bool)`

GetCertSanUpnOk returns a tuple with the CertSanUpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSanUpn

`func (o *ConstNacEvent) SetCertSanUpn(v []string)`

SetCertSanUpn sets CertSanUpn field to given value.

### HasCertSanUpn

`func (o *ConstNacEvent) HasCertSanUpn() bool`

HasCertSanUpn returns a boolean if a field has been set.

### GetCertSerial

`func (o *ConstNacEvent) GetCertSerial() string`

GetCertSerial returns the CertSerial field if non-nil, zero value otherwise.

### GetCertSerialOk

`func (o *ConstNacEvent) GetCertSerialOk() (*string, bool)`

GetCertSerialOk returns a tuple with the CertSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSerial

`func (o *ConstNacEvent) SetCertSerial(v string)`

SetCertSerial sets CertSerial field to given value.

### HasCertSerial

`func (o *ConstNacEvent) HasCertSerial() bool`

HasCertSerial returns a boolean if a field has been set.

### GetCertSubject

`func (o *ConstNacEvent) GetCertSubject() string`

GetCertSubject returns the CertSubject field if non-nil, zero value otherwise.

### GetCertSubjectOk

`func (o *ConstNacEvent) GetCertSubjectOk() (*string, bool)`

GetCertSubjectOk returns a tuple with the CertSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSubject

`func (o *ConstNacEvent) SetCertSubject(v string)`

SetCertSubject sets CertSubject field to given value.

### HasCertSubject

`func (o *ConstNacEvent) HasCertSubject() bool`

HasCertSubject returns a boolean if a field has been set.

### GetEapType

`func (o *ConstNacEvent) GetEapType() string`

GetEapType returns the EapType field if non-nil, zero value otherwise.

### GetEapTypeOk

`func (o *ConstNacEvent) GetEapTypeOk() (*string, bool)`

GetEapTypeOk returns a tuple with the EapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapType

`func (o *ConstNacEvent) SetEapType(v string)`

SetEapType sets EapType field to given value.

### HasEapType

`func (o *ConstNacEvent) HasEapType() bool`

HasEapType returns a boolean if a field has been set.

### GetNasVendor

`func (o *ConstNacEvent) GetNasVendor() string`

GetNasVendor returns the NasVendor field if non-nil, zero value otherwise.

### GetNasVendorOk

`func (o *ConstNacEvent) GetNasVendorOk() (*string, bool)`

GetNasVendorOk returns a tuple with the NasVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasVendor

`func (o *ConstNacEvent) SetNasVendor(v string)`

SetNasVendor sets NasVendor field to given value.

### HasNasVendor

`func (o *ConstNacEvent) HasNasVendor() bool`

HasNasVendor returns a boolean if a field has been set.

### GetOrgId

`func (o *ConstNacEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConstNacEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConstNacEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConstNacEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRandomMac

`func (o *ConstNacEvent) GetRandomMac() bool`

GetRandomMac returns the RandomMac field if non-nil, zero value otherwise.

### GetRandomMacOk

`func (o *ConstNacEvent) GetRandomMacOk() (*bool, bool)`

GetRandomMacOk returns a tuple with the RandomMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomMac

`func (o *ConstNacEvent) SetRandomMac(v bool)`

SetRandomMac sets RandomMac field to given value.

### HasRandomMac

`func (o *ConstNacEvent) HasRandomMac() bool`

HasRandomMac returns a boolean if a field has been set.

### GetSiteId

`func (o *ConstNacEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConstNacEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConstNacEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ConstNacEvent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsid

`func (o *ConstNacEvent) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ConstNacEvent) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ConstNacEvent) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ConstNacEvent) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimestamp

`func (o *ConstNacEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConstNacEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConstNacEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConstNacEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *ConstNacEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstNacEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstNacEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstNacEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *ConstNacEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConstNacEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConstNacEvent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConstNacEvent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWcid

`func (o *ConstNacEvent) GetWcid() string`

GetWcid returns the Wcid field if non-nil, zero value otherwise.

### GetWcidOk

`func (o *ConstNacEvent) GetWcidOk() (*string, bool)`

GetWcidOk returns a tuple with the Wcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWcid

`func (o *ConstNacEvent) SetWcid(v string)`

SetWcid sets Wcid field to given value.

### HasWcid

`func (o *ConstNacEvent) HasWcid() bool`

HasWcid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


