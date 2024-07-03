# MxedgeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuStat** | Pointer to [**MxedgeStatsCpuStat**](MxedgeStatsCpuStat.md) |  | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] 
**FipsEnabled** | Pointer to **bool** | alue indicating fips configuration on the device | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpStat** | Pointer to [**MxedgeStatsIpStat**](MxedgeStatsIpStat.md) |  | [optional] 
**LagStat** | Pointer to [**map[string]MxedgeStatsLagStat**](MxedgeStatsLagStat.md) | Stat for LAG (Link Aggregation Group). Property key is the LAG name | [optional] 
**LastSeen** | Pointer to **float32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**MemoryStat** | Pointer to [**MxedgeStatsMemoryStat**](MxedgeStatsMemoryStat.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to **int32** |  | [optional] 
**MxagentRegistered** | Pointer to **bool** |  | [optional] 
**MxclusterId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the tunnel | [optional] 
**NumTunnels** | Pointer to **int32** |  | [optional] 
**OobIpConfig** | Pointer to [**MxedgeOobIpConfig**](MxedgeOobIpConfig.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PortStat** | Pointer to [**map[string]MxedgeStatsPortStat**](MxedgeStatsPortStat.md) |  | [optional] 
**SensorStat** | Pointer to **map[string]interface{}** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**ServiceStat** | Pointer to [**map[string]MxedgeStatsServiceStat**](MxedgeStatsServiceStat.md) | stat for each services | [optional] 
**Services** | Pointer to **[]string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] 
**TuntermIpConfig** | Pointer to [**MxedgeStatsTuntermIpConfig**](MxedgeStatsTuntermIpConfig.md) |  | [optional] 
**TuntermPortConfig** | Pointer to [**MxedgeStatsTuntermPortConfig**](MxedgeStatsTuntermPortConfig.md) |  | [optional] 
**TuntermRegistered** | Pointer to **bool** |  | [optional] 
**TuntermStat** | Pointer to [**MxedgeStatsTuntermStat**](MxedgeStatsTuntermStat.md) |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**VirtualizationType** | Pointer to **string** | Virtualization environment | [optional] 

## Methods

### NewMxedgeStats

`func NewMxedgeStats() *MxedgeStats`

NewMxedgeStats instantiates a new MxedgeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsWithDefaults

`func NewMxedgeStatsWithDefaults() *MxedgeStats`

NewMxedgeStatsWithDefaults instantiates a new MxedgeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuStat

`func (o *MxedgeStats) GetCpuStat() MxedgeStatsCpuStat`

GetCpuStat returns the CpuStat field if non-nil, zero value otherwise.

### GetCpuStatOk

`func (o *MxedgeStats) GetCpuStatOk() (*MxedgeStatsCpuStat, bool)`

GetCpuStatOk returns a tuple with the CpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStat

`func (o *MxedgeStats) SetCpuStat(v MxedgeStatsCpuStat)`

SetCpuStat sets CpuStat field to given value.

### HasCpuStat

`func (o *MxedgeStats) HasCpuStat() bool`

HasCpuStat returns a boolean if a field has been set.

### GetCreatedTime

`func (o *MxedgeStats) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MxedgeStats) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MxedgeStats) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *MxedgeStats) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetFipsEnabled

`func (o *MxedgeStats) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *MxedgeStats) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *MxedgeStats) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *MxedgeStats) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetForSite

`func (o *MxedgeStats) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *MxedgeStats) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *MxedgeStats) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *MxedgeStats) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *MxedgeStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MxedgeStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MxedgeStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MxedgeStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpStat

`func (o *MxedgeStats) GetIpStat() MxedgeStatsIpStat`

GetIpStat returns the IpStat field if non-nil, zero value otherwise.

### GetIpStatOk

`func (o *MxedgeStats) GetIpStatOk() (*MxedgeStatsIpStat, bool)`

GetIpStatOk returns a tuple with the IpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStat

`func (o *MxedgeStats) SetIpStat(v MxedgeStatsIpStat)`

SetIpStat sets IpStat field to given value.

### HasIpStat

`func (o *MxedgeStats) HasIpStat() bool`

HasIpStat returns a boolean if a field has been set.

### GetLagStat

`func (o *MxedgeStats) GetLagStat() map[string]MxedgeStatsLagStat`

GetLagStat returns the LagStat field if non-nil, zero value otherwise.

### GetLagStatOk

`func (o *MxedgeStats) GetLagStatOk() (*map[string]MxedgeStatsLagStat, bool)`

GetLagStatOk returns a tuple with the LagStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagStat

`func (o *MxedgeStats) SetLagStat(v map[string]MxedgeStatsLagStat)`

SetLagStat sets LagStat field to given value.

### HasLagStat

`func (o *MxedgeStats) HasLagStat() bool`

HasLagStat returns a boolean if a field has been set.

### GetLastSeen

`func (o *MxedgeStats) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *MxedgeStats) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *MxedgeStats) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *MxedgeStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *MxedgeStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MxedgeStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MxedgeStats) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MxedgeStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemoryStat

`func (o *MxedgeStats) GetMemoryStat() MxedgeStatsMemoryStat`

GetMemoryStat returns the MemoryStat field if non-nil, zero value otherwise.

### GetMemoryStatOk

`func (o *MxedgeStats) GetMemoryStatOk() (*MxedgeStatsMemoryStat, bool)`

GetMemoryStatOk returns a tuple with the MemoryStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryStat

`func (o *MxedgeStats) SetMemoryStat(v MxedgeStatsMemoryStat)`

SetMemoryStat sets MemoryStat field to given value.

### HasMemoryStat

`func (o *MxedgeStats) HasMemoryStat() bool`

HasMemoryStat returns a boolean if a field has been set.

### GetModel

`func (o *MxedgeStats) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MxedgeStats) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MxedgeStats) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MxedgeStats) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModifiedTime

`func (o *MxedgeStats) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *MxedgeStats) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *MxedgeStats) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *MxedgeStats) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMxagentRegistered

`func (o *MxedgeStats) GetMxagentRegistered() bool`

GetMxagentRegistered returns the MxagentRegistered field if non-nil, zero value otherwise.

### GetMxagentRegisteredOk

`func (o *MxedgeStats) GetMxagentRegisteredOk() (*bool, bool)`

GetMxagentRegisteredOk returns a tuple with the MxagentRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxagentRegistered

`func (o *MxedgeStats) SetMxagentRegistered(v bool)`

SetMxagentRegistered sets MxagentRegistered field to given value.

### HasMxagentRegistered

`func (o *MxedgeStats) HasMxagentRegistered() bool`

HasMxagentRegistered returns a boolean if a field has been set.

### GetMxclusterId

`func (o *MxedgeStats) GetMxclusterId() string`

GetMxclusterId returns the MxclusterId field if non-nil, zero value otherwise.

### GetMxclusterIdOk

`func (o *MxedgeStats) GetMxclusterIdOk() (*string, bool)`

GetMxclusterIdOk returns a tuple with the MxclusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterId

`func (o *MxedgeStats) SetMxclusterId(v string)`

SetMxclusterId sets MxclusterId field to given value.

### HasMxclusterId

`func (o *MxedgeStats) HasMxclusterId() bool`

HasMxclusterId returns a boolean if a field has been set.

### GetName

`func (o *MxedgeStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MxedgeStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MxedgeStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MxedgeStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumTunnels

`func (o *MxedgeStats) GetNumTunnels() int32`

GetNumTunnels returns the NumTunnels field if non-nil, zero value otherwise.

### GetNumTunnelsOk

`func (o *MxedgeStats) GetNumTunnelsOk() (*int32, bool)`

GetNumTunnelsOk returns a tuple with the NumTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTunnels

`func (o *MxedgeStats) SetNumTunnels(v int32)`

SetNumTunnels sets NumTunnels field to given value.

### HasNumTunnels

`func (o *MxedgeStats) HasNumTunnels() bool`

HasNumTunnels returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *MxedgeStats) GetOobIpConfig() MxedgeOobIpConfig`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *MxedgeStats) GetOobIpConfigOk() (*MxedgeOobIpConfig, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *MxedgeStats) SetOobIpConfig(v MxedgeOobIpConfig)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *MxedgeStats) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOrgId

`func (o *MxedgeStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MxedgeStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MxedgeStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MxedgeStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortStat

`func (o *MxedgeStats) GetPortStat() map[string]MxedgeStatsPortStat`

GetPortStat returns the PortStat field if non-nil, zero value otherwise.

### GetPortStatOk

`func (o *MxedgeStats) GetPortStatOk() (*map[string]MxedgeStatsPortStat, bool)`

GetPortStatOk returns a tuple with the PortStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStat

`func (o *MxedgeStats) SetPortStat(v map[string]MxedgeStatsPortStat)`

SetPortStat sets PortStat field to given value.

### HasPortStat

`func (o *MxedgeStats) HasPortStat() bool`

HasPortStat returns a boolean if a field has been set.

### GetSensorStat

`func (o *MxedgeStats) GetSensorStat() map[string]interface{}`

GetSensorStat returns the SensorStat field if non-nil, zero value otherwise.

### GetSensorStatOk

`func (o *MxedgeStats) GetSensorStatOk() (*map[string]interface{}, bool)`

GetSensorStatOk returns a tuple with the SensorStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorStat

`func (o *MxedgeStats) SetSensorStat(v map[string]interface{})`

SetSensorStat sets SensorStat field to given value.

### HasSensorStat

`func (o *MxedgeStats) HasSensorStat() bool`

HasSensorStat returns a boolean if a field has been set.

### GetSerial

`func (o *MxedgeStats) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *MxedgeStats) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *MxedgeStats) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *MxedgeStats) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *MxedgeStats) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *MxedgeStats) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetServiceStat

`func (o *MxedgeStats) GetServiceStat() map[string]MxedgeStatsServiceStat`

GetServiceStat returns the ServiceStat field if non-nil, zero value otherwise.

### GetServiceStatOk

`func (o *MxedgeStats) GetServiceStatOk() (*map[string]MxedgeStatsServiceStat, bool)`

GetServiceStatOk returns a tuple with the ServiceStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStat

`func (o *MxedgeStats) SetServiceStat(v map[string]MxedgeStatsServiceStat)`

SetServiceStat sets ServiceStat field to given value.

### HasServiceStat

`func (o *MxedgeStats) HasServiceStat() bool`

HasServiceStat returns a boolean if a field has been set.

### GetServices

`func (o *MxedgeStats) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MxedgeStats) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MxedgeStats) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *MxedgeStats) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSiteId

`func (o *MxedgeStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MxedgeStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MxedgeStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MxedgeStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *MxedgeStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MxedgeStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MxedgeStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MxedgeStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTuntermIpConfig

`func (o *MxedgeStats) GetTuntermIpConfig() MxedgeStatsTuntermIpConfig`

GetTuntermIpConfig returns the TuntermIpConfig field if non-nil, zero value otherwise.

### GetTuntermIpConfigOk

`func (o *MxedgeStats) GetTuntermIpConfigOk() (*MxedgeStatsTuntermIpConfig, bool)`

GetTuntermIpConfigOk returns a tuple with the TuntermIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermIpConfig

`func (o *MxedgeStats) SetTuntermIpConfig(v MxedgeStatsTuntermIpConfig)`

SetTuntermIpConfig sets TuntermIpConfig field to given value.

### HasTuntermIpConfig

`func (o *MxedgeStats) HasTuntermIpConfig() bool`

HasTuntermIpConfig returns a boolean if a field has been set.

### GetTuntermPortConfig

`func (o *MxedgeStats) GetTuntermPortConfig() MxedgeStatsTuntermPortConfig`

GetTuntermPortConfig returns the TuntermPortConfig field if non-nil, zero value otherwise.

### GetTuntermPortConfigOk

`func (o *MxedgeStats) GetTuntermPortConfigOk() (*MxedgeStatsTuntermPortConfig, bool)`

GetTuntermPortConfigOk returns a tuple with the TuntermPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermPortConfig

`func (o *MxedgeStats) SetTuntermPortConfig(v MxedgeStatsTuntermPortConfig)`

SetTuntermPortConfig sets TuntermPortConfig field to given value.

### HasTuntermPortConfig

`func (o *MxedgeStats) HasTuntermPortConfig() bool`

HasTuntermPortConfig returns a boolean if a field has been set.

### GetTuntermRegistered

`func (o *MxedgeStats) GetTuntermRegistered() bool`

GetTuntermRegistered returns the TuntermRegistered field if non-nil, zero value otherwise.

### GetTuntermRegisteredOk

`func (o *MxedgeStats) GetTuntermRegisteredOk() (*bool, bool)`

GetTuntermRegisteredOk returns a tuple with the TuntermRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermRegistered

`func (o *MxedgeStats) SetTuntermRegistered(v bool)`

SetTuntermRegistered sets TuntermRegistered field to given value.

### HasTuntermRegistered

`func (o *MxedgeStats) HasTuntermRegistered() bool`

HasTuntermRegistered returns a boolean if a field has been set.

### GetTuntermStat

`func (o *MxedgeStats) GetTuntermStat() MxedgeStatsTuntermStat`

GetTuntermStat returns the TuntermStat field if non-nil, zero value otherwise.

### GetTuntermStatOk

`func (o *MxedgeStats) GetTuntermStatOk() (*MxedgeStatsTuntermStat, bool)`

GetTuntermStatOk returns a tuple with the TuntermStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermStat

`func (o *MxedgeStats) SetTuntermStat(v MxedgeStatsTuntermStat)`

SetTuntermStat sets TuntermStat field to given value.

### HasTuntermStat

`func (o *MxedgeStats) HasTuntermStat() bool`

HasTuntermStat returns a boolean if a field has been set.

### GetUptime

`func (o *MxedgeStats) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MxedgeStats) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MxedgeStats) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MxedgeStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVirtualizationType

`func (o *MxedgeStats) GetVirtualizationType() string`

GetVirtualizationType returns the VirtualizationType field if non-nil, zero value otherwise.

### GetVirtualizationTypeOk

`func (o *MxedgeStats) GetVirtualizationTypeOk() (*string, bool)`

GetVirtualizationTypeOk returns a tuple with the VirtualizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationType

`func (o *MxedgeStats) SetVirtualizationType(v string)`

SetVirtualizationType sets VirtualizationType field to given value.

### HasVirtualizationType

`func (o *MxedgeStats) HasVirtualizationType() bool`

HasVirtualizationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


