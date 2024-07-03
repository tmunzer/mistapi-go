# DiscoveredSwitchMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adopted** | Pointer to **bool** |  | [optional] 
**Aps** | Pointer to [**[]DiscoveredSwitchMetricAp**](DiscoveredSwitchMetricAp.md) |  | [optional] 
**ChassisId** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**MgmtAddr** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Scope** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SystemDesc** | Pointer to **string** |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDiscoveredSwitchMetric

`func NewDiscoveredSwitchMetric() *DiscoveredSwitchMetric`

NewDiscoveredSwitchMetric instantiates a new DiscoveredSwitchMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredSwitchMetricWithDefaults

`func NewDiscoveredSwitchMetricWithDefaults() *DiscoveredSwitchMetric`

NewDiscoveredSwitchMetricWithDefaults instantiates a new DiscoveredSwitchMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdopted

`func (o *DiscoveredSwitchMetric) GetAdopted() bool`

GetAdopted returns the Adopted field if non-nil, zero value otherwise.

### GetAdoptedOk

`func (o *DiscoveredSwitchMetric) GetAdoptedOk() (*bool, bool)`

GetAdoptedOk returns a tuple with the Adopted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdopted

`func (o *DiscoveredSwitchMetric) SetAdopted(v bool)`

SetAdopted sets Adopted field to given value.

### HasAdopted

`func (o *DiscoveredSwitchMetric) HasAdopted() bool`

HasAdopted returns a boolean if a field has been set.

### GetAps

`func (o *DiscoveredSwitchMetric) GetAps() []DiscoveredSwitchMetricAp`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *DiscoveredSwitchMetric) GetApsOk() (*[]DiscoveredSwitchMetricAp, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *DiscoveredSwitchMetric) SetAps(v []DiscoveredSwitchMetricAp)`

SetAps sets Aps field to given value.

### HasAps

`func (o *DiscoveredSwitchMetric) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetChassisId

`func (o *DiscoveredSwitchMetric) GetChassisId() []string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *DiscoveredSwitchMetric) GetChassisIdOk() (*[]string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *DiscoveredSwitchMetric) SetChassisId(v []string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *DiscoveredSwitchMetric) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetHostname

`func (o *DiscoveredSwitchMetric) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DiscoveredSwitchMetric) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DiscoveredSwitchMetric) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DiscoveredSwitchMetric) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMgmtAddr

`func (o *DiscoveredSwitchMetric) GetMgmtAddr() string`

GetMgmtAddr returns the MgmtAddr field if non-nil, zero value otherwise.

### GetMgmtAddrOk

`func (o *DiscoveredSwitchMetric) GetMgmtAddrOk() (*string, bool)`

GetMgmtAddrOk returns a tuple with the MgmtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtAddr

`func (o *DiscoveredSwitchMetric) SetMgmtAddr(v string)`

SetMgmtAddr sets MgmtAddr field to given value.

### HasMgmtAddr

`func (o *DiscoveredSwitchMetric) HasMgmtAddr() bool`

HasMgmtAddr returns a boolean if a field has been set.

### GetModel

`func (o *DiscoveredSwitchMetric) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiscoveredSwitchMetric) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiscoveredSwitchMetric) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiscoveredSwitchMetric) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrgId

`func (o *DiscoveredSwitchMetric) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DiscoveredSwitchMetric) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DiscoveredSwitchMetric) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DiscoveredSwitchMetric) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetScope

`func (o *DiscoveredSwitchMetric) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DiscoveredSwitchMetric) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DiscoveredSwitchMetric) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DiscoveredSwitchMetric) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScore

`func (o *DiscoveredSwitchMetric) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DiscoveredSwitchMetric) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DiscoveredSwitchMetric) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *DiscoveredSwitchMetric) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSiteId

`func (o *DiscoveredSwitchMetric) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DiscoveredSwitchMetric) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DiscoveredSwitchMetric) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DiscoveredSwitchMetric) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemDesc

`func (o *DiscoveredSwitchMetric) GetSystemDesc() string`

GetSystemDesc returns the SystemDesc field if non-nil, zero value otherwise.

### GetSystemDescOk

`func (o *DiscoveredSwitchMetric) GetSystemDescOk() (*string, bool)`

GetSystemDescOk returns a tuple with the SystemDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDesc

`func (o *DiscoveredSwitchMetric) SetSystemDesc(v string)`

SetSystemDesc sets SystemDesc field to given value.

### HasSystemDesc

`func (o *DiscoveredSwitchMetric) HasSystemDesc() bool`

HasSystemDesc returns a boolean if a field has been set.

### GetSystemName

`func (o *DiscoveredSwitchMetric) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *DiscoveredSwitchMetric) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *DiscoveredSwitchMetric) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *DiscoveredSwitchMetric) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTimestamp

`func (o *DiscoveredSwitchMetric) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscoveredSwitchMetric) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscoveredSwitchMetric) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscoveredSwitchMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *DiscoveredSwitchMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredSwitchMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredSwitchMetric) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredSwitchMetric) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *DiscoveredSwitchMetric) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DiscoveredSwitchMetric) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DiscoveredSwitchMetric) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DiscoveredSwitchMetric) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *DiscoveredSwitchMetric) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DiscoveredSwitchMetric) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DiscoveredSwitchMetric) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DiscoveredSwitchMetric) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


