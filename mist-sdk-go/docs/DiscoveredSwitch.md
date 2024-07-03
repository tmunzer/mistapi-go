# DiscoveredSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adopted** | Pointer to **bool** |  | [optional] 
**ApRedundancy** | Pointer to [**ApRedundancy**](ApRedundancy.md) |  | [optional] 
**Aps** | Pointer to [**[]DiscoveredSwitchAp**](DiscoveredSwitchAp.md) |  | [optional] 
**ChassisId** | Pointer to **[]string** |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Model** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SystemDesc** | Pointer to **string** |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDiscoveredSwitch

`func NewDiscoveredSwitch() *DiscoveredSwitch`

NewDiscoveredSwitch instantiates a new DiscoveredSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredSwitchWithDefaults

`func NewDiscoveredSwitchWithDefaults() *DiscoveredSwitch`

NewDiscoveredSwitchWithDefaults instantiates a new DiscoveredSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdopted

`func (o *DiscoveredSwitch) GetAdopted() bool`

GetAdopted returns the Adopted field if non-nil, zero value otherwise.

### GetAdoptedOk

`func (o *DiscoveredSwitch) GetAdoptedOk() (*bool, bool)`

GetAdoptedOk returns a tuple with the Adopted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdopted

`func (o *DiscoveredSwitch) SetAdopted(v bool)`

SetAdopted sets Adopted field to given value.

### HasAdopted

`func (o *DiscoveredSwitch) HasAdopted() bool`

HasAdopted returns a boolean if a field has been set.

### GetApRedundancy

`func (o *DiscoveredSwitch) GetApRedundancy() ApRedundancy`

GetApRedundancy returns the ApRedundancy field if non-nil, zero value otherwise.

### GetApRedundancyOk

`func (o *DiscoveredSwitch) GetApRedundancyOk() (*ApRedundancy, bool)`

GetApRedundancyOk returns a tuple with the ApRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRedundancy

`func (o *DiscoveredSwitch) SetApRedundancy(v ApRedundancy)`

SetApRedundancy sets ApRedundancy field to given value.

### HasApRedundancy

`func (o *DiscoveredSwitch) HasApRedundancy() bool`

HasApRedundancy returns a boolean if a field has been set.

### GetAps

`func (o *DiscoveredSwitch) GetAps() []DiscoveredSwitchAp`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *DiscoveredSwitch) GetApsOk() (*[]DiscoveredSwitchAp, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *DiscoveredSwitch) SetAps(v []DiscoveredSwitchAp)`

SetAps sets Aps field to given value.

### HasAps

`func (o *DiscoveredSwitch) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetChassisId

`func (o *DiscoveredSwitch) GetChassisId() []string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *DiscoveredSwitch) GetChassisIdOk() (*[]string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *DiscoveredSwitch) SetChassisId(v []string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *DiscoveredSwitch) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetForSite

`func (o *DiscoveredSwitch) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *DiscoveredSwitch) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *DiscoveredSwitch) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *DiscoveredSwitch) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetModel

`func (o *DiscoveredSwitch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiscoveredSwitch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiscoveredSwitch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiscoveredSwitch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrgId

`func (o *DiscoveredSwitch) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DiscoveredSwitch) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DiscoveredSwitch) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DiscoveredSwitch) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteId

`func (o *DiscoveredSwitch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DiscoveredSwitch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DiscoveredSwitch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DiscoveredSwitch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemDesc

`func (o *DiscoveredSwitch) GetSystemDesc() string`

GetSystemDesc returns the SystemDesc field if non-nil, zero value otherwise.

### GetSystemDescOk

`func (o *DiscoveredSwitch) GetSystemDescOk() (*string, bool)`

GetSystemDescOk returns a tuple with the SystemDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDesc

`func (o *DiscoveredSwitch) SetSystemDesc(v string)`

SetSystemDesc sets SystemDesc field to given value.

### HasSystemDesc

`func (o *DiscoveredSwitch) HasSystemDesc() bool`

HasSystemDesc returns a boolean if a field has been set.

### GetSystemName

`func (o *DiscoveredSwitch) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *DiscoveredSwitch) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *DiscoveredSwitch) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *DiscoveredSwitch) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTimestamp

`func (o *DiscoveredSwitch) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscoveredSwitch) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscoveredSwitch) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscoveredSwitch) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVendor

`func (o *DiscoveredSwitch) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DiscoveredSwitch) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DiscoveredSwitch) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DiscoveredSwitch) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *DiscoveredSwitch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DiscoveredSwitch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DiscoveredSwitch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DiscoveredSwitch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


