# SwitchSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtIp** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**LastHostname** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**NumMembers** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewSwitchSearch

`func NewSwitchSearch() *SwitchSearch`

NewSwitchSearch instantiates a new SwitchSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchSearchWithDefaults

`func NewSwitchSearchWithDefaults() *SwitchSearch`

NewSwitchSearchWithDefaults instantiates a new SwitchSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtIp

`func (o *SwitchSearch) GetExtIp() string`

GetExtIp returns the ExtIp field if non-nil, zero value otherwise.

### GetExtIpOk

`func (o *SwitchSearch) GetExtIpOk() (*string, bool)`

GetExtIpOk returns a tuple with the ExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIp

`func (o *SwitchSearch) SetExtIp(v string)`

SetExtIp sets ExtIp field to given value.

### HasExtIp

`func (o *SwitchSearch) HasExtIp() bool`

HasExtIp returns a boolean if a field has been set.

### GetHostname

`func (o *SwitchSearch) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SwitchSearch) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SwitchSearch) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SwitchSearch) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *SwitchSearch) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SwitchSearch) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SwitchSearch) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SwitchSearch) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastHostname

`func (o *SwitchSearch) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *SwitchSearch) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *SwitchSearch) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *SwitchSearch) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetMac

`func (o *SwitchSearch) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchSearch) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchSearch) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SwitchSearch) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *SwitchSearch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SwitchSearch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SwitchSearch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SwitchSearch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNumMembers

`func (o *SwitchSearch) GetNumMembers() int32`

GetNumMembers returns the NumMembers field if non-nil, zero value otherwise.

### GetNumMembersOk

`func (o *SwitchSearch) GetNumMembersOk() (*int32, bool)`

GetNumMembersOk returns a tuple with the NumMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMembers

`func (o *SwitchSearch) SetNumMembers(v int32)`

SetNumMembers sets NumMembers field to given value.

### HasNumMembers

`func (o *SwitchSearch) HasNumMembers() bool`

HasNumMembers returns a boolean if a field has been set.

### GetOrgId

`func (o *SwitchSearch) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SwitchSearch) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SwitchSearch) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SwitchSearch) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteId

`func (o *SwitchSearch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SwitchSearch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SwitchSearch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SwitchSearch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SwitchSearch) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SwitchSearch) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SwitchSearch) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SwitchSearch) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *SwitchSearch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchSearch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchSearch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchSearch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *SwitchSearch) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SwitchSearch) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SwitchSearch) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SwitchSearch) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *SwitchSearch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SwitchSearch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SwitchSearch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SwitchSearch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


