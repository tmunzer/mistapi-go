# GatewayStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApRedundancy** | Pointer to [**ApRedundancy**](ApRedundancy.md) |  | [optional] 
**ClusterStat** | Pointer to [**map[string]GatewayStatsClusterStat**](GatewayStatsClusterStat.md) |  | [optional] 
**Cpu2Stat** | Pointer to [**CpuStat**](CpuStat.md) |  | [optional] 
**CpuStat** | Pointer to [**CpuStat**](CpuStat.md) |  | [optional] 
**Dhcpd2Stat** | Pointer to [**map[string]GatewayStatsDhcpdStatLan**](GatewayStatsDhcpdStatLan.md) | Property key is the network name | [optional] 
**DhcpdStat** | Pointer to [**map[string]GatewayStatsDhcpdStatLan**](GatewayStatsDhcpdStatLan.md) | Property key is the network name | [optional] 
**Hostname** | Pointer to **string** | hostname reported by the device | [optional] 
**Ip** | Pointer to **string** | IP address | [optional] 
**Ip2Stat** | Pointer to [**IpStat**](IpStat.md) |  | [optional] 
**IpStat** | Pointer to [**IpStat**](IpStat.md) |  | [optional] 
**LastSeen** | Pointer to **float32** | last seen timestamp | [optional] 
**Mac** | **string** | device mac | 
**Memory2Stat** | Pointer to [**MemoryStat**](MemoryStat.md) |  | [optional] 
**MemoryStat** | Pointer to [**MemoryStat**](MemoryStat.md) |  | [optional] 
**Model** | Pointer to **string** | device model | [optional] 
**Module2Stat** | Pointer to [**[]ModuleStat**](ModuleStat.md) |  | [optional] 
**ModuleStat** | Pointer to [**[]ModuleStat**](ModuleStat.md) |  | [optional] 
**Name** | Pointer to **string** | device name if configured | [optional] 
**Serial** | Pointer to **string** | serial | [optional] 
**Spu2Stat** | Pointer to [**GatewayStatsSpuStat**](GatewayStatsSpuStat.md) |  | [optional] 
**SpuStat** | Pointer to [**GatewayStatsSpuStat**](GatewayStatsSpuStat.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayStats

`func NewGatewayStats(mac string, ) *GatewayStats`

NewGatewayStats instantiates a new GatewayStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayStatsWithDefaults

`func NewGatewayStatsWithDefaults() *GatewayStats`

NewGatewayStatsWithDefaults instantiates a new GatewayStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApRedundancy

`func (o *GatewayStats) GetApRedundancy() ApRedundancy`

GetApRedundancy returns the ApRedundancy field if non-nil, zero value otherwise.

### GetApRedundancyOk

`func (o *GatewayStats) GetApRedundancyOk() (*ApRedundancy, bool)`

GetApRedundancyOk returns a tuple with the ApRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRedundancy

`func (o *GatewayStats) SetApRedundancy(v ApRedundancy)`

SetApRedundancy sets ApRedundancy field to given value.

### HasApRedundancy

`func (o *GatewayStats) HasApRedundancy() bool`

HasApRedundancy returns a boolean if a field has been set.

### GetClusterStat

`func (o *GatewayStats) GetClusterStat() map[string]GatewayStatsClusterStat`

GetClusterStat returns the ClusterStat field if non-nil, zero value otherwise.

### GetClusterStatOk

`func (o *GatewayStats) GetClusterStatOk() (*map[string]GatewayStatsClusterStat, bool)`

GetClusterStatOk returns a tuple with the ClusterStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStat

`func (o *GatewayStats) SetClusterStat(v map[string]GatewayStatsClusterStat)`

SetClusterStat sets ClusterStat field to given value.

### HasClusterStat

`func (o *GatewayStats) HasClusterStat() bool`

HasClusterStat returns a boolean if a field has been set.

### GetCpu2Stat

`func (o *GatewayStats) GetCpu2Stat() CpuStat`

GetCpu2Stat returns the Cpu2Stat field if non-nil, zero value otherwise.

### GetCpu2StatOk

`func (o *GatewayStats) GetCpu2StatOk() (*CpuStat, bool)`

GetCpu2StatOk returns a tuple with the Cpu2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu2Stat

`func (o *GatewayStats) SetCpu2Stat(v CpuStat)`

SetCpu2Stat sets Cpu2Stat field to given value.

### HasCpu2Stat

`func (o *GatewayStats) HasCpu2Stat() bool`

HasCpu2Stat returns a boolean if a field has been set.

### GetCpuStat

`func (o *GatewayStats) GetCpuStat() CpuStat`

GetCpuStat returns the CpuStat field if non-nil, zero value otherwise.

### GetCpuStatOk

`func (o *GatewayStats) GetCpuStatOk() (*CpuStat, bool)`

GetCpuStatOk returns a tuple with the CpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStat

`func (o *GatewayStats) SetCpuStat(v CpuStat)`

SetCpuStat sets CpuStat field to given value.

### HasCpuStat

`func (o *GatewayStats) HasCpuStat() bool`

HasCpuStat returns a boolean if a field has been set.

### GetDhcpd2Stat

`func (o *GatewayStats) GetDhcpd2Stat() map[string]GatewayStatsDhcpdStatLan`

GetDhcpd2Stat returns the Dhcpd2Stat field if non-nil, zero value otherwise.

### GetDhcpd2StatOk

`func (o *GatewayStats) GetDhcpd2StatOk() (*map[string]GatewayStatsDhcpdStatLan, bool)`

GetDhcpd2StatOk returns a tuple with the Dhcpd2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpd2Stat

`func (o *GatewayStats) SetDhcpd2Stat(v map[string]GatewayStatsDhcpdStatLan)`

SetDhcpd2Stat sets Dhcpd2Stat field to given value.

### HasDhcpd2Stat

`func (o *GatewayStats) HasDhcpd2Stat() bool`

HasDhcpd2Stat returns a boolean if a field has been set.

### GetDhcpdStat

`func (o *GatewayStats) GetDhcpdStat() map[string]GatewayStatsDhcpdStatLan`

GetDhcpdStat returns the DhcpdStat field if non-nil, zero value otherwise.

### GetDhcpdStatOk

`func (o *GatewayStats) GetDhcpdStatOk() (*map[string]GatewayStatsDhcpdStatLan, bool)`

GetDhcpdStatOk returns a tuple with the DhcpdStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdStat

`func (o *GatewayStats) SetDhcpdStat(v map[string]GatewayStatsDhcpdStatLan)`

SetDhcpdStat sets DhcpdStat field to given value.

### HasDhcpdStat

`func (o *GatewayStats) HasDhcpdStat() bool`

HasDhcpdStat returns a boolean if a field has been set.

### GetHostname

`func (o *GatewayStats) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GatewayStats) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GatewayStats) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GatewayStats) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *GatewayStats) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayStats) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayStats) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayStats) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp2Stat

`func (o *GatewayStats) GetIp2Stat() IpStat`

GetIp2Stat returns the Ip2Stat field if non-nil, zero value otherwise.

### GetIp2StatOk

`func (o *GatewayStats) GetIp2StatOk() (*IpStat, bool)`

GetIp2StatOk returns a tuple with the Ip2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2Stat

`func (o *GatewayStats) SetIp2Stat(v IpStat)`

SetIp2Stat sets Ip2Stat field to given value.

### HasIp2Stat

`func (o *GatewayStats) HasIp2Stat() bool`

HasIp2Stat returns a boolean if a field has been set.

### GetIpStat

`func (o *GatewayStats) GetIpStat() IpStat`

GetIpStat returns the IpStat field if non-nil, zero value otherwise.

### GetIpStatOk

`func (o *GatewayStats) GetIpStatOk() (*IpStat, bool)`

GetIpStatOk returns a tuple with the IpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStat

`func (o *GatewayStats) SetIpStat(v IpStat)`

SetIpStat sets IpStat field to given value.

### HasIpStat

`func (o *GatewayStats) HasIpStat() bool`

HasIpStat returns a boolean if a field has been set.

### GetLastSeen

`func (o *GatewayStats) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GatewayStats) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GatewayStats) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GatewayStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *GatewayStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GatewayStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GatewayStats) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMemory2Stat

`func (o *GatewayStats) GetMemory2Stat() MemoryStat`

GetMemory2Stat returns the Memory2Stat field if non-nil, zero value otherwise.

### GetMemory2StatOk

`func (o *GatewayStats) GetMemory2StatOk() (*MemoryStat, bool)`

GetMemory2StatOk returns a tuple with the Memory2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory2Stat

`func (o *GatewayStats) SetMemory2Stat(v MemoryStat)`

SetMemory2Stat sets Memory2Stat field to given value.

### HasMemory2Stat

`func (o *GatewayStats) HasMemory2Stat() bool`

HasMemory2Stat returns a boolean if a field has been set.

### GetMemoryStat

`func (o *GatewayStats) GetMemoryStat() MemoryStat`

GetMemoryStat returns the MemoryStat field if non-nil, zero value otherwise.

### GetMemoryStatOk

`func (o *GatewayStats) GetMemoryStatOk() (*MemoryStat, bool)`

GetMemoryStatOk returns a tuple with the MemoryStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryStat

`func (o *GatewayStats) SetMemoryStat(v MemoryStat)`

SetMemoryStat sets MemoryStat field to given value.

### HasMemoryStat

`func (o *GatewayStats) HasMemoryStat() bool`

HasMemoryStat returns a boolean if a field has been set.

### GetModel

`func (o *GatewayStats) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GatewayStats) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GatewayStats) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GatewayStats) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModule2Stat

`func (o *GatewayStats) GetModule2Stat() []ModuleStat`

GetModule2Stat returns the Module2Stat field if non-nil, zero value otherwise.

### GetModule2StatOk

`func (o *GatewayStats) GetModule2StatOk() (*[]ModuleStat, bool)`

GetModule2StatOk returns a tuple with the Module2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule2Stat

`func (o *GatewayStats) SetModule2Stat(v []ModuleStat)`

SetModule2Stat sets Module2Stat field to given value.

### HasModule2Stat

`func (o *GatewayStats) HasModule2Stat() bool`

HasModule2Stat returns a boolean if a field has been set.

### GetModuleStat

`func (o *GatewayStats) GetModuleStat() []ModuleStat`

GetModuleStat returns the ModuleStat field if non-nil, zero value otherwise.

### GetModuleStatOk

`func (o *GatewayStats) GetModuleStatOk() (*[]ModuleStat, bool)`

GetModuleStatOk returns a tuple with the ModuleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStat

`func (o *GatewayStats) SetModuleStat(v []ModuleStat)`

SetModuleStat sets ModuleStat field to given value.

### HasModuleStat

`func (o *GatewayStats) HasModuleStat() bool`

HasModuleStat returns a boolean if a field has been set.

### GetName

`func (o *GatewayStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *GatewayStats) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GatewayStats) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GatewayStats) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GatewayStats) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSpu2Stat

`func (o *GatewayStats) GetSpu2Stat() GatewayStatsSpuStat`

GetSpu2Stat returns the Spu2Stat field if non-nil, zero value otherwise.

### GetSpu2StatOk

`func (o *GatewayStats) GetSpu2StatOk() (*GatewayStatsSpuStat, bool)`

GetSpu2StatOk returns a tuple with the Spu2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpu2Stat

`func (o *GatewayStats) SetSpu2Stat(v GatewayStatsSpuStat)`

SetSpu2Stat sets Spu2Stat field to given value.

### HasSpu2Stat

`func (o *GatewayStats) HasSpu2Stat() bool`

HasSpu2Stat returns a boolean if a field has been set.

### GetSpuStat

`func (o *GatewayStats) GetSpuStat() GatewayStatsSpuStat`

GetSpuStat returns the SpuStat field if non-nil, zero value otherwise.

### GetSpuStatOk

`func (o *GatewayStats) GetSpuStatOk() (*GatewayStatsSpuStat, bool)`

GetSpuStatOk returns a tuple with the SpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpuStat

`func (o *GatewayStats) SetSpuStat(v GatewayStatsSpuStat)`

SetSpuStat sets SpuStat field to given value.

### HasSpuStat

`func (o *GatewayStats) HasSpuStat() bool`

HasSpuStat returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GatewayStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *GatewayStats) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GatewayStats) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GatewayStats) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GatewayStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *GatewayStats) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewayStats) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewayStats) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GatewayStats) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


