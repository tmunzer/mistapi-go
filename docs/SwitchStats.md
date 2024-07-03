# SwitchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApRedundancy** | Pointer to [**SwitchStatsApRedundancy**](SwitchStatsApRedundancy.md) |  | [optional] 
**Clients** | Pointer to [**[]SwitchStatsClient**](SwitchStatsClient.md) |  | [optional] 
**CpuStat** | Pointer to [**CpuStat**](CpuStat.md) |  | [optional] 
**HasPcap** | Pointer to **bool** | whether the switch supports packet capture | [optional] 
**Hostname** | Pointer to **string** | hostname reported by the device | [optional] 
**IfStat** | Pointer to [**map[string]SwitchStatsIfStat**](SwitchStatsIfStat.md) | Property key is the interface name | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpStat** | Pointer to [**IpStat**](IpStat.md) |  | [optional] 
**LastSeen** | Pointer to **int32** |  | [optional] 
**LastTrouble** | Pointer to [**LastTrouble**](LastTrouble.md) |  | [optional] 
**Mac** | Pointer to [**MemoryStat**](MemoryStat.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModuleStat** | Pointer to [**[]ModuleStat**](ModuleStat.md) |  | [optional] 
**Name** | Pointer to **string** | device name if configured | [optional] 
**NumClients** | Pointer to [**SwitchStatsNumClients**](SwitchStatsNumClients.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewSwitchStats

`func NewSwitchStats() *SwitchStats`

NewSwitchStats instantiates a new SwitchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchStatsWithDefaults

`func NewSwitchStatsWithDefaults() *SwitchStats`

NewSwitchStatsWithDefaults instantiates a new SwitchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApRedundancy

`func (o *SwitchStats) GetApRedundancy() SwitchStatsApRedundancy`

GetApRedundancy returns the ApRedundancy field if non-nil, zero value otherwise.

### GetApRedundancyOk

`func (o *SwitchStats) GetApRedundancyOk() (*SwitchStatsApRedundancy, bool)`

GetApRedundancyOk returns a tuple with the ApRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRedundancy

`func (o *SwitchStats) SetApRedundancy(v SwitchStatsApRedundancy)`

SetApRedundancy sets ApRedundancy field to given value.

### HasApRedundancy

`func (o *SwitchStats) HasApRedundancy() bool`

HasApRedundancy returns a boolean if a field has been set.

### GetClients

`func (o *SwitchStats) GetClients() []SwitchStatsClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *SwitchStats) GetClientsOk() (*[]SwitchStatsClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *SwitchStats) SetClients(v []SwitchStatsClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *SwitchStats) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCpuStat

`func (o *SwitchStats) GetCpuStat() CpuStat`

GetCpuStat returns the CpuStat field if non-nil, zero value otherwise.

### GetCpuStatOk

`func (o *SwitchStats) GetCpuStatOk() (*CpuStat, bool)`

GetCpuStatOk returns a tuple with the CpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStat

`func (o *SwitchStats) SetCpuStat(v CpuStat)`

SetCpuStat sets CpuStat field to given value.

### HasCpuStat

`func (o *SwitchStats) HasCpuStat() bool`

HasCpuStat returns a boolean if a field has been set.

### GetHasPcap

`func (o *SwitchStats) GetHasPcap() bool`

GetHasPcap returns the HasPcap field if non-nil, zero value otherwise.

### GetHasPcapOk

`func (o *SwitchStats) GetHasPcapOk() (*bool, bool)`

GetHasPcapOk returns a tuple with the HasPcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPcap

`func (o *SwitchStats) SetHasPcap(v bool)`

SetHasPcap sets HasPcap field to given value.

### HasHasPcap

`func (o *SwitchStats) HasHasPcap() bool`

HasHasPcap returns a boolean if a field has been set.

### GetHostname

`func (o *SwitchStats) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SwitchStats) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SwitchStats) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SwitchStats) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIfStat

`func (o *SwitchStats) GetIfStat() map[string]SwitchStatsIfStat`

GetIfStat returns the IfStat field if non-nil, zero value otherwise.

### GetIfStatOk

`func (o *SwitchStats) GetIfStatOk() (*map[string]SwitchStatsIfStat, bool)`

GetIfStatOk returns a tuple with the IfStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfStat

`func (o *SwitchStats) SetIfStat(v map[string]SwitchStatsIfStat)`

SetIfStat sets IfStat field to given value.

### HasIfStat

`func (o *SwitchStats) HasIfStat() bool`

HasIfStat returns a boolean if a field has been set.

### GetIp

`func (o *SwitchStats) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SwitchStats) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SwitchStats) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SwitchStats) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpStat

`func (o *SwitchStats) GetIpStat() IpStat`

GetIpStat returns the IpStat field if non-nil, zero value otherwise.

### GetIpStatOk

`func (o *SwitchStats) GetIpStatOk() (*IpStat, bool)`

GetIpStatOk returns a tuple with the IpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStat

`func (o *SwitchStats) SetIpStat(v IpStat)`

SetIpStat sets IpStat field to given value.

### HasIpStat

`func (o *SwitchStats) HasIpStat() bool`

HasIpStat returns a boolean if a field has been set.

### GetLastSeen

`func (o *SwitchStats) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *SwitchStats) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *SwitchStats) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *SwitchStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastTrouble

`func (o *SwitchStats) GetLastTrouble() LastTrouble`

GetLastTrouble returns the LastTrouble field if non-nil, zero value otherwise.

### GetLastTroubleOk

`func (o *SwitchStats) GetLastTroubleOk() (*LastTrouble, bool)`

GetLastTroubleOk returns a tuple with the LastTrouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrouble

`func (o *SwitchStats) SetLastTrouble(v LastTrouble)`

SetLastTrouble sets LastTrouble field to given value.

### HasLastTrouble

`func (o *SwitchStats) HasLastTrouble() bool`

HasLastTrouble returns a boolean if a field has been set.

### GetMac

`func (o *SwitchStats) GetMac() MemoryStat`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchStats) GetMacOk() (*MemoryStat, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchStats) SetMac(v MemoryStat)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SwitchStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *SwitchStats) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SwitchStats) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SwitchStats) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SwitchStats) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModuleStat

`func (o *SwitchStats) GetModuleStat() []ModuleStat`

GetModuleStat returns the ModuleStat field if non-nil, zero value otherwise.

### GetModuleStatOk

`func (o *SwitchStats) GetModuleStatOk() (*[]ModuleStat, bool)`

GetModuleStatOk returns a tuple with the ModuleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStat

`func (o *SwitchStats) SetModuleStat(v []ModuleStat)`

SetModuleStat sets ModuleStat field to given value.

### HasModuleStat

`func (o *SwitchStats) HasModuleStat() bool`

HasModuleStat returns a boolean if a field has been set.

### GetName

`func (o *SwitchStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SwitchStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumClients

`func (o *SwitchStats) GetNumClients() SwitchStatsNumClients`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *SwitchStats) GetNumClientsOk() (*SwitchStatsNumClients, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *SwitchStats) SetNumClients(v SwitchStatsNumClients)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *SwitchStats) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetSerial

`func (o *SwitchStats) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *SwitchStats) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *SwitchStats) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *SwitchStats) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *SwitchStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwitchStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SwitchStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *SwitchStats) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SwitchStats) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SwitchStats) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SwitchStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *SwitchStats) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SwitchStats) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SwitchStats) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SwitchStats) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


