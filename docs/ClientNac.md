# ClientNac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **[]string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**CertCn** | Pointer to **[]string** |  | [optional] 
**CertIssuer** | Pointer to **[]string** |  | [optional] 
**IdpId** | Pointer to **string** |  | [optional] 
**IdpRole** | Pointer to **[]string** |  | [optional] 
**LastAp** | Pointer to **string** |  | [optional] 
**LastCertCn** | Pointer to **string** |  | [optional] 
**LastCertExpiry** | Pointer to **int32** |  | [optional] 
**LastCertIssuer** | Pointer to **string** |  | [optional] 
**LastNacruleId** | Pointer to **string** |  | [optional] 
**LastNacruleName** | Pointer to **string** |  | [optional] 
**LastNasVendor** | Pointer to **string** |  | [optional] 
**LastSsid** | Pointer to **string** |  | [optional] 
**LastStatus** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**NacruleId** | Pointer to **[]string** |  | [optional] 
**NacruleMatched** | Pointer to **bool** |  | [optional] 
**NacruleName** | Pointer to **[]string** |  | [optional] 
**NasVendor** | Pointer to **[]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RandomMac** | Pointer to **bool** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Ssid** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewClientNac

`func NewClientNac() *ClientNac`

NewClientNac instantiates a new ClientNac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientNacWithDefaults

`func NewClientNacWithDefaults() *ClientNac`

NewClientNacWithDefaults instantiates a new ClientNac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *ClientNac) GetAp() []string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *ClientNac) GetApOk() (*[]string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *ClientNac) SetAp(v []string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *ClientNac) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetAuthType

`func (o *ClientNac) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ClientNac) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ClientNac) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ClientNac) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCertCn

`func (o *ClientNac) GetCertCn() []string`

GetCertCn returns the CertCn field if non-nil, zero value otherwise.

### GetCertCnOk

`func (o *ClientNac) GetCertCnOk() (*[]string, bool)`

GetCertCnOk returns a tuple with the CertCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertCn

`func (o *ClientNac) SetCertCn(v []string)`

SetCertCn sets CertCn field to given value.

### HasCertCn

`func (o *ClientNac) HasCertCn() bool`

HasCertCn returns a boolean if a field has been set.

### GetCertIssuer

`func (o *ClientNac) GetCertIssuer() []string`

GetCertIssuer returns the CertIssuer field if non-nil, zero value otherwise.

### GetCertIssuerOk

`func (o *ClientNac) GetCertIssuerOk() (*[]string, bool)`

GetCertIssuerOk returns a tuple with the CertIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuer

`func (o *ClientNac) SetCertIssuer(v []string)`

SetCertIssuer sets CertIssuer field to given value.

### HasCertIssuer

`func (o *ClientNac) HasCertIssuer() bool`

HasCertIssuer returns a boolean if a field has been set.

### GetIdpId

`func (o *ClientNac) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *ClientNac) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *ClientNac) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *ClientNac) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetIdpRole

`func (o *ClientNac) GetIdpRole() []string`

GetIdpRole returns the IdpRole field if non-nil, zero value otherwise.

### GetIdpRoleOk

`func (o *ClientNac) GetIdpRoleOk() (*[]string, bool)`

GetIdpRoleOk returns a tuple with the IdpRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpRole

`func (o *ClientNac) SetIdpRole(v []string)`

SetIdpRole sets IdpRole field to given value.

### HasIdpRole

`func (o *ClientNac) HasIdpRole() bool`

HasIdpRole returns a boolean if a field has been set.

### GetLastAp

`func (o *ClientNac) GetLastAp() string`

GetLastAp returns the LastAp field if non-nil, zero value otherwise.

### GetLastApOk

`func (o *ClientNac) GetLastApOk() (*string, bool)`

GetLastApOk returns a tuple with the LastAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAp

`func (o *ClientNac) SetLastAp(v string)`

SetLastAp sets LastAp field to given value.

### HasLastAp

`func (o *ClientNac) HasLastAp() bool`

HasLastAp returns a boolean if a field has been set.

### GetLastCertCn

`func (o *ClientNac) GetLastCertCn() string`

GetLastCertCn returns the LastCertCn field if non-nil, zero value otherwise.

### GetLastCertCnOk

`func (o *ClientNac) GetLastCertCnOk() (*string, bool)`

GetLastCertCnOk returns a tuple with the LastCertCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertCn

`func (o *ClientNac) SetLastCertCn(v string)`

SetLastCertCn sets LastCertCn field to given value.

### HasLastCertCn

`func (o *ClientNac) HasLastCertCn() bool`

HasLastCertCn returns a boolean if a field has been set.

### GetLastCertExpiry

`func (o *ClientNac) GetLastCertExpiry() int32`

GetLastCertExpiry returns the LastCertExpiry field if non-nil, zero value otherwise.

### GetLastCertExpiryOk

`func (o *ClientNac) GetLastCertExpiryOk() (*int32, bool)`

GetLastCertExpiryOk returns a tuple with the LastCertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertExpiry

`func (o *ClientNac) SetLastCertExpiry(v int32)`

SetLastCertExpiry sets LastCertExpiry field to given value.

### HasLastCertExpiry

`func (o *ClientNac) HasLastCertExpiry() bool`

HasLastCertExpiry returns a boolean if a field has been set.

### GetLastCertIssuer

`func (o *ClientNac) GetLastCertIssuer() string`

GetLastCertIssuer returns the LastCertIssuer field if non-nil, zero value otherwise.

### GetLastCertIssuerOk

`func (o *ClientNac) GetLastCertIssuerOk() (*string, bool)`

GetLastCertIssuerOk returns a tuple with the LastCertIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertIssuer

`func (o *ClientNac) SetLastCertIssuer(v string)`

SetLastCertIssuer sets LastCertIssuer field to given value.

### HasLastCertIssuer

`func (o *ClientNac) HasLastCertIssuer() bool`

HasLastCertIssuer returns a boolean if a field has been set.

### GetLastNacruleId

`func (o *ClientNac) GetLastNacruleId() string`

GetLastNacruleId returns the LastNacruleId field if non-nil, zero value otherwise.

### GetLastNacruleIdOk

`func (o *ClientNac) GetLastNacruleIdOk() (*string, bool)`

GetLastNacruleIdOk returns a tuple with the LastNacruleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNacruleId

`func (o *ClientNac) SetLastNacruleId(v string)`

SetLastNacruleId sets LastNacruleId field to given value.

### HasLastNacruleId

`func (o *ClientNac) HasLastNacruleId() bool`

HasLastNacruleId returns a boolean if a field has been set.

### GetLastNacruleName

`func (o *ClientNac) GetLastNacruleName() string`

GetLastNacruleName returns the LastNacruleName field if non-nil, zero value otherwise.

### GetLastNacruleNameOk

`func (o *ClientNac) GetLastNacruleNameOk() (*string, bool)`

GetLastNacruleNameOk returns a tuple with the LastNacruleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNacruleName

`func (o *ClientNac) SetLastNacruleName(v string)`

SetLastNacruleName sets LastNacruleName field to given value.

### HasLastNacruleName

`func (o *ClientNac) HasLastNacruleName() bool`

HasLastNacruleName returns a boolean if a field has been set.

### GetLastNasVendor

`func (o *ClientNac) GetLastNasVendor() string`

GetLastNasVendor returns the LastNasVendor field if non-nil, zero value otherwise.

### GetLastNasVendorOk

`func (o *ClientNac) GetLastNasVendorOk() (*string, bool)`

GetLastNasVendorOk returns a tuple with the LastNasVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNasVendor

`func (o *ClientNac) SetLastNasVendor(v string)`

SetLastNasVendor sets LastNasVendor field to given value.

### HasLastNasVendor

`func (o *ClientNac) HasLastNasVendor() bool`

HasLastNasVendor returns a boolean if a field has been set.

### GetLastSsid

`func (o *ClientNac) GetLastSsid() string`

GetLastSsid returns the LastSsid field if non-nil, zero value otherwise.

### GetLastSsidOk

`func (o *ClientNac) GetLastSsidOk() (*string, bool)`

GetLastSsidOk returns a tuple with the LastSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSsid

`func (o *ClientNac) SetLastSsid(v string)`

SetLastSsid sets LastSsid field to given value.

### HasLastSsid

`func (o *ClientNac) HasLastSsid() bool`

HasLastSsid returns a boolean if a field has been set.

### GetLastStatus

`func (o *ClientNac) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ClientNac) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ClientNac) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ClientNac) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetMac

`func (o *ClientNac) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientNac) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientNac) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientNac) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNacruleId

`func (o *ClientNac) GetNacruleId() []string`

GetNacruleId returns the NacruleId field if non-nil, zero value otherwise.

### GetNacruleIdOk

`func (o *ClientNac) GetNacruleIdOk() (*[]string, bool)`

GetNacruleIdOk returns a tuple with the NacruleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacruleId

`func (o *ClientNac) SetNacruleId(v []string)`

SetNacruleId sets NacruleId field to given value.

### HasNacruleId

`func (o *ClientNac) HasNacruleId() bool`

HasNacruleId returns a boolean if a field has been set.

### GetNacruleMatched

`func (o *ClientNac) GetNacruleMatched() bool`

GetNacruleMatched returns the NacruleMatched field if non-nil, zero value otherwise.

### GetNacruleMatchedOk

`func (o *ClientNac) GetNacruleMatchedOk() (*bool, bool)`

GetNacruleMatchedOk returns a tuple with the NacruleMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacruleMatched

`func (o *ClientNac) SetNacruleMatched(v bool)`

SetNacruleMatched sets NacruleMatched field to given value.

### HasNacruleMatched

`func (o *ClientNac) HasNacruleMatched() bool`

HasNacruleMatched returns a boolean if a field has been set.

### GetNacruleName

`func (o *ClientNac) GetNacruleName() []string`

GetNacruleName returns the NacruleName field if non-nil, zero value otherwise.

### GetNacruleNameOk

`func (o *ClientNac) GetNacruleNameOk() (*[]string, bool)`

GetNacruleNameOk returns a tuple with the NacruleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacruleName

`func (o *ClientNac) SetNacruleName(v []string)`

SetNacruleName sets NacruleName field to given value.

### HasNacruleName

`func (o *ClientNac) HasNacruleName() bool`

HasNacruleName returns a boolean if a field has been set.

### GetNasVendor

`func (o *ClientNac) GetNasVendor() []string`

GetNasVendor returns the NasVendor field if non-nil, zero value otherwise.

### GetNasVendorOk

`func (o *ClientNac) GetNasVendorOk() (*[]string, bool)`

GetNasVendorOk returns a tuple with the NasVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasVendor

`func (o *ClientNac) SetNasVendor(v []string)`

SetNasVendor sets NasVendor field to given value.

### HasNasVendor

`func (o *ClientNac) HasNasVendor() bool`

HasNasVendor returns a boolean if a field has been set.

### GetOrgId

`func (o *ClientNac) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ClientNac) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ClientNac) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ClientNac) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRandomMac

`func (o *ClientNac) GetRandomMac() bool`

GetRandomMac returns the RandomMac field if non-nil, zero value otherwise.

### GetRandomMacOk

`func (o *ClientNac) GetRandomMacOk() (*bool, bool)`

GetRandomMacOk returns a tuple with the RandomMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomMac

`func (o *ClientNac) SetRandomMac(v bool)`

SetRandomMac sets RandomMac field to given value.

### HasRandomMac

`func (o *ClientNac) HasRandomMac() bool`

HasRandomMac returns a boolean if a field has been set.

### GetSiteId

`func (o *ClientNac) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ClientNac) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ClientNac) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ClientNac) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsid

`func (o *ClientNac) GetSsid() []string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientNac) GetSsidOk() (*[]string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientNac) SetSsid(v []string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientNac) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimestamp

`func (o *ClientNac) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClientNac) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClientNac) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ClientNac) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *ClientNac) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientNac) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientNac) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientNac) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


