# StatsDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPlacement** | Pointer to [**ApStatsAutoPlacement**](ApStatsAutoPlacement.md) |  | [optional] 
**BleConfig** | Pointer to [**ApStatsBleConfig**](ApStatsBleConfig.md) |  | [optional] 
**BleStat** | Pointer to [**ApStatsBleStat**](ApStatsBleStat.md) |  | [optional] 
**CertExpiry** | Pointer to **float32** |  | [optional] 
**EnvStat** | Pointer to [**ApStatsEnvStat**](ApStatsEnvStat.md) |  | [optional] 
**EslStat** | Pointer to [**ApStatsEslStat**](ApStatsEslStat.md) |  | [optional] 
**ExtIp** | Pointer to **string** |  | [optional] 
**Fwupdate** | Pointer to [**ApStatsFwupdate**](ApStatsFwupdate.md) |  | [optional] 
**IotStat** | Pointer to [**map[string]ApStatsIotStatAdditionalProperties**](ApStatsIotStatAdditionalProperties.md) |  | [optional] 
**Ip** | Pointer to **string** | IP address | [optional] 
**IpConfig** | Pointer to [**ApIpConfig**](ApIpConfig.md) |  | [optional] 
**IpStat** | Pointer to [**IpStat**](IpStat.md) |  | [optional] 
**L2tpStat** | Pointer to [**map[string]ApStatsL2tpStat**](ApStatsL2tpStat.md) | l2tp tunnel status (key is the wxtunnel_id) | [optional] 
**LastSeen** | Pointer to **float32** | last seen timestamp | [optional] 
**LastTrouble** | Pointer to [**LastTrouble**](LastTrouble.md) |  | [optional] 
**Led** | Pointer to [**ApLed**](ApLed.md) |  | [optional] 
**LldpStat** | Pointer to [**ApStatsLldpStat**](ApStatsLldpStat.md) |  | [optional] 
**Locating** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** | whether this AP is considered locked (placement / orientation has been vetted) | [optional] 
**Mac** | **string** | device mac | 
**MapId** | Pointer to **string** |  | [optional] 
**MeshDownlinks** | Pointer to [**map[string]ApStatMeshDownlink**](ApStatMeshDownlink.md) | Property key is the mesh downlink id (e.g &#x60;00000000-0000-0000-1000-5c5b35000010&#x60;) | [optional] 
**MeshUplink** | Pointer to [**ApStatMeshUplink**](ApStatMeshUplink.md) |  | [optional] 
**Model** | **string** | device model | 
**Mount** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | device name if configured | [optional] 
**NumClients** | Pointer to [**SwitchStatsNumClients**](SwitchStatsNumClients.md) |  | [optional] 
**PortStat** | Pointer to [**map[string]ApStatsPortStat**](ApStatsPortStat.md) | Property key is the port name (e.g. &#x60;eth0&#x60;) | [optional] 
**PowerBudget** | Pointer to **float32** | in mW, surplus if positive or deficit if negative | [optional] 
**PowerConstrained** | Pointer to **bool** | whether insufficient power | [optional] 
**PowerOpmode** | Pointer to **string** | constrained mode | [optional] 
**PowerSrc** | Pointer to **string** | DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown) | [optional] 
**RadioConfig** | Pointer to [**ApStatsRadioConfig**](ApStatsRadioConfig.md) |  | [optional] 
**RadioStat** | Pointer to [**ApStatsRadioStat**](ApStatsRadioStat.md) |  | [optional] 
**RxBps** | Pointer to **float32** |  | [optional] 
**RxBytes** | Pointer to **int32** |  | [optional] 
**RxPkts** | Pointer to **int32** |  | [optional] 
**Serial** | Pointer to **string** | serial | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SwitchRedundancy** | Pointer to [**ApStatsSwitchRedundancy**](ApStatsSwitchRedundancy.md) |  | [optional] 
**TxBps** | Pointer to **float32** |  | [optional] 
**TxBytes** | Pointer to **float32** |  | [optional] 
**TxPkts** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **float32** |  | [optional] 
**UsbStat** | Pointer to [**ApStatsUsbStat**](ApStatsUsbStat.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float32** |  | [optional] 
**Y** | Pointer to **float32** |  | [optional] 
**ApRedundancy** | Pointer to [**ApRedundancy**](ApRedundancy.md) |  | [optional] 
**Clients** | Pointer to [**[]SwitchStatsClient**](SwitchStatsClient.md) |  | [optional] 
**CpuStat** | Pointer to [**CpuStat**](CpuStat.md) |  | [optional] 
**HasPcap** | Pointer to **bool** | whether the switch supports packet capture | [optional] 
**Hostname** | Pointer to **string** | hostname reported by the device | [optional] 
**IfStat** | Pointer to [**map[string]SwitchStatsIfStat**](SwitchStatsIfStat.md) | Property key is the interface name | [optional] 
**ModuleStat** | Pointer to [**[]ModuleStat**](ModuleStat.md) |  | [optional] 
**ClusterStat** | Pointer to [**map[string]GatewayStatsClusterStat**](GatewayStatsClusterStat.md) |  | [optional] 
**Cpu2Stat** | Pointer to [**CpuStat**](CpuStat.md) |  | [optional] 
**Dhcpd2Stat** | Pointer to [**map[string]GatewayStatsDhcpdStatLan**](GatewayStatsDhcpdStatLan.md) | Property key is the network name | [optional] 
**DhcpdStat** | Pointer to [**map[string]GatewayStatsDhcpdStatLan**](GatewayStatsDhcpdStatLan.md) | Property key is the network name | [optional] 
**Ip2Stat** | Pointer to [**IpStat**](IpStat.md) |  | [optional] 
**Memory2Stat** | Pointer to [**MemoryStat**](MemoryStat.md) |  | [optional] 
**MemoryStat** | Pointer to [**MemoryStat**](MemoryStat.md) |  | [optional] 
**Module2Stat** | Pointer to [**[]ModuleStat**](ModuleStat.md) |  | [optional] 
**Spu2Stat** | Pointer to [**GatewayStatsSpuStat**](GatewayStatsSpuStat.md) |  | [optional] 
**SpuStat** | Pointer to [**GatewayStatsSpuStat**](GatewayStatsSpuStat.md) |  | [optional] 

## Methods

### NewStatsDevice

`func NewStatsDevice(mac string, model string, ) *StatsDevice`

NewStatsDevice instantiates a new StatsDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsDeviceWithDefaults

`func NewStatsDeviceWithDefaults() *StatsDevice`

NewStatsDeviceWithDefaults instantiates a new StatsDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPlacement

`func (o *StatsDevice) GetAutoPlacement() ApStatsAutoPlacement`

GetAutoPlacement returns the AutoPlacement field if non-nil, zero value otherwise.

### GetAutoPlacementOk

`func (o *StatsDevice) GetAutoPlacementOk() (*ApStatsAutoPlacement, bool)`

GetAutoPlacementOk returns a tuple with the AutoPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlacement

`func (o *StatsDevice) SetAutoPlacement(v ApStatsAutoPlacement)`

SetAutoPlacement sets AutoPlacement field to given value.

### HasAutoPlacement

`func (o *StatsDevice) HasAutoPlacement() bool`

HasAutoPlacement returns a boolean if a field has been set.

### GetBleConfig

`func (o *StatsDevice) GetBleConfig() ApStatsBleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *StatsDevice) GetBleConfigOk() (*ApStatsBleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *StatsDevice) SetBleConfig(v ApStatsBleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *StatsDevice) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetBleStat

`func (o *StatsDevice) GetBleStat() ApStatsBleStat`

GetBleStat returns the BleStat field if non-nil, zero value otherwise.

### GetBleStatOk

`func (o *StatsDevice) GetBleStatOk() (*ApStatsBleStat, bool)`

GetBleStatOk returns a tuple with the BleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleStat

`func (o *StatsDevice) SetBleStat(v ApStatsBleStat)`

SetBleStat sets BleStat field to given value.

### HasBleStat

`func (o *StatsDevice) HasBleStat() bool`

HasBleStat returns a boolean if a field has been set.

### GetCertExpiry

`func (o *StatsDevice) GetCertExpiry() float32`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *StatsDevice) GetCertExpiryOk() (*float32, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *StatsDevice) SetCertExpiry(v float32)`

SetCertExpiry sets CertExpiry field to given value.

### HasCertExpiry

`func (o *StatsDevice) HasCertExpiry() bool`

HasCertExpiry returns a boolean if a field has been set.

### GetEnvStat

`func (o *StatsDevice) GetEnvStat() ApStatsEnvStat`

GetEnvStat returns the EnvStat field if non-nil, zero value otherwise.

### GetEnvStatOk

`func (o *StatsDevice) GetEnvStatOk() (*ApStatsEnvStat, bool)`

GetEnvStatOk returns a tuple with the EnvStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvStat

`func (o *StatsDevice) SetEnvStat(v ApStatsEnvStat)`

SetEnvStat sets EnvStat field to given value.

### HasEnvStat

`func (o *StatsDevice) HasEnvStat() bool`

HasEnvStat returns a boolean if a field has been set.

### GetEslStat

`func (o *StatsDevice) GetEslStat() ApStatsEslStat`

GetEslStat returns the EslStat field if non-nil, zero value otherwise.

### GetEslStatOk

`func (o *StatsDevice) GetEslStatOk() (*ApStatsEslStat, bool)`

GetEslStatOk returns a tuple with the EslStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslStat

`func (o *StatsDevice) SetEslStat(v ApStatsEslStat)`

SetEslStat sets EslStat field to given value.

### HasEslStat

`func (o *StatsDevice) HasEslStat() bool`

HasEslStat returns a boolean if a field has been set.

### GetExtIp

`func (o *StatsDevice) GetExtIp() string`

GetExtIp returns the ExtIp field if non-nil, zero value otherwise.

### GetExtIpOk

`func (o *StatsDevice) GetExtIpOk() (*string, bool)`

GetExtIpOk returns a tuple with the ExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIp

`func (o *StatsDevice) SetExtIp(v string)`

SetExtIp sets ExtIp field to given value.

### HasExtIp

`func (o *StatsDevice) HasExtIp() bool`

HasExtIp returns a boolean if a field has been set.

### GetFwupdate

`func (o *StatsDevice) GetFwupdate() ApStatsFwupdate`

GetFwupdate returns the Fwupdate field if non-nil, zero value otherwise.

### GetFwupdateOk

`func (o *StatsDevice) GetFwupdateOk() (*ApStatsFwupdate, bool)`

GetFwupdateOk returns a tuple with the Fwupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwupdate

`func (o *StatsDevice) SetFwupdate(v ApStatsFwupdate)`

SetFwupdate sets Fwupdate field to given value.

### HasFwupdate

`func (o *StatsDevice) HasFwupdate() bool`

HasFwupdate returns a boolean if a field has been set.

### GetIotStat

`func (o *StatsDevice) GetIotStat() map[string]ApStatsIotStatAdditionalProperties`

GetIotStat returns the IotStat field if non-nil, zero value otherwise.

### GetIotStatOk

`func (o *StatsDevice) GetIotStatOk() (*map[string]ApStatsIotStatAdditionalProperties, bool)`

GetIotStatOk returns a tuple with the IotStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotStat

`func (o *StatsDevice) SetIotStat(v map[string]ApStatsIotStatAdditionalProperties)`

SetIotStat sets IotStat field to given value.

### HasIotStat

`func (o *StatsDevice) HasIotStat() bool`

HasIotStat returns a boolean if a field has been set.

### GetIp

`func (o *StatsDevice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *StatsDevice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *StatsDevice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *StatsDevice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpConfig

`func (o *StatsDevice) GetIpConfig() ApIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *StatsDevice) GetIpConfigOk() (*ApIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *StatsDevice) SetIpConfig(v ApIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *StatsDevice) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetIpStat

`func (o *StatsDevice) GetIpStat() IpStat`

GetIpStat returns the IpStat field if non-nil, zero value otherwise.

### GetIpStatOk

`func (o *StatsDevice) GetIpStatOk() (*IpStat, bool)`

GetIpStatOk returns a tuple with the IpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStat

`func (o *StatsDevice) SetIpStat(v IpStat)`

SetIpStat sets IpStat field to given value.

### HasIpStat

`func (o *StatsDevice) HasIpStat() bool`

HasIpStat returns a boolean if a field has been set.

### GetL2tpStat

`func (o *StatsDevice) GetL2tpStat() map[string]ApStatsL2tpStat`

GetL2tpStat returns the L2tpStat field if non-nil, zero value otherwise.

### GetL2tpStatOk

`func (o *StatsDevice) GetL2tpStatOk() (*map[string]ApStatsL2tpStat, bool)`

GetL2tpStatOk returns a tuple with the L2tpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2tpStat

`func (o *StatsDevice) SetL2tpStat(v map[string]ApStatsL2tpStat)`

SetL2tpStat sets L2tpStat field to given value.

### HasL2tpStat

`func (o *StatsDevice) HasL2tpStat() bool`

HasL2tpStat returns a boolean if a field has been set.

### GetLastSeen

`func (o *StatsDevice) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *StatsDevice) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *StatsDevice) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *StatsDevice) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastTrouble

`func (o *StatsDevice) GetLastTrouble() LastTrouble`

GetLastTrouble returns the LastTrouble field if non-nil, zero value otherwise.

### GetLastTroubleOk

`func (o *StatsDevice) GetLastTroubleOk() (*LastTrouble, bool)`

GetLastTroubleOk returns a tuple with the LastTrouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrouble

`func (o *StatsDevice) SetLastTrouble(v LastTrouble)`

SetLastTrouble sets LastTrouble field to given value.

### HasLastTrouble

`func (o *StatsDevice) HasLastTrouble() bool`

HasLastTrouble returns a boolean if a field has been set.

### GetLed

`func (o *StatsDevice) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *StatsDevice) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *StatsDevice) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *StatsDevice) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLldpStat

`func (o *StatsDevice) GetLldpStat() ApStatsLldpStat`

GetLldpStat returns the LldpStat field if non-nil, zero value otherwise.

### GetLldpStatOk

`func (o *StatsDevice) GetLldpStatOk() (*ApStatsLldpStat, bool)`

GetLldpStatOk returns a tuple with the LldpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpStat

`func (o *StatsDevice) SetLldpStat(v ApStatsLldpStat)`

SetLldpStat sets LldpStat field to given value.

### HasLldpStat

`func (o *StatsDevice) HasLldpStat() bool`

HasLldpStat returns a boolean if a field has been set.

### GetLocating

`func (o *StatsDevice) GetLocating() bool`

GetLocating returns the Locating field if non-nil, zero value otherwise.

### GetLocatingOk

`func (o *StatsDevice) GetLocatingOk() (*bool, bool)`

GetLocatingOk returns a tuple with the Locating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocating

`func (o *StatsDevice) SetLocating(v bool)`

SetLocating sets Locating field to given value.

### HasLocating

`func (o *StatsDevice) HasLocating() bool`

HasLocating returns a boolean if a field has been set.

### GetLocked

`func (o *StatsDevice) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *StatsDevice) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *StatsDevice) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *StatsDevice) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMac

`func (o *StatsDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *StatsDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *StatsDevice) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *StatsDevice) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *StatsDevice) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *StatsDevice) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *StatsDevice) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMeshDownlinks

`func (o *StatsDevice) GetMeshDownlinks() map[string]ApStatMeshDownlink`

GetMeshDownlinks returns the MeshDownlinks field if non-nil, zero value otherwise.

### GetMeshDownlinksOk

`func (o *StatsDevice) GetMeshDownlinksOk() (*map[string]ApStatMeshDownlink, bool)`

GetMeshDownlinksOk returns a tuple with the MeshDownlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshDownlinks

`func (o *StatsDevice) SetMeshDownlinks(v map[string]ApStatMeshDownlink)`

SetMeshDownlinks sets MeshDownlinks field to given value.

### HasMeshDownlinks

`func (o *StatsDevice) HasMeshDownlinks() bool`

HasMeshDownlinks returns a boolean if a field has been set.

### GetMeshUplink

`func (o *StatsDevice) GetMeshUplink() ApStatMeshUplink`

GetMeshUplink returns the MeshUplink field if non-nil, zero value otherwise.

### GetMeshUplinkOk

`func (o *StatsDevice) GetMeshUplinkOk() (*ApStatMeshUplink, bool)`

GetMeshUplinkOk returns a tuple with the MeshUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshUplink

`func (o *StatsDevice) SetMeshUplink(v ApStatMeshUplink)`

SetMeshUplink sets MeshUplink field to given value.

### HasMeshUplink

`func (o *StatsDevice) HasMeshUplink() bool`

HasMeshUplink returns a boolean if a field has been set.

### GetModel

`func (o *StatsDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StatsDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StatsDevice) SetModel(v string)`

SetModel sets Model field to given value.


### GetMount

`func (o *StatsDevice) GetMount() string`

GetMount returns the Mount field if non-nil, zero value otherwise.

### GetMountOk

`func (o *StatsDevice) GetMountOk() (*string, bool)`

GetMountOk returns a tuple with the Mount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMount

`func (o *StatsDevice) SetMount(v string)`

SetMount sets Mount field to given value.

### HasMount

`func (o *StatsDevice) HasMount() bool`

HasMount returns a boolean if a field has been set.

### GetName

`func (o *StatsDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatsDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatsDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatsDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumClients

`func (o *StatsDevice) GetNumClients() SwitchStatsNumClients`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *StatsDevice) GetNumClientsOk() (*SwitchStatsNumClients, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *StatsDevice) SetNumClients(v SwitchStatsNumClients)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *StatsDevice) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetPortStat

`func (o *StatsDevice) GetPortStat() map[string]ApStatsPortStat`

GetPortStat returns the PortStat field if non-nil, zero value otherwise.

### GetPortStatOk

`func (o *StatsDevice) GetPortStatOk() (*map[string]ApStatsPortStat, bool)`

GetPortStatOk returns a tuple with the PortStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStat

`func (o *StatsDevice) SetPortStat(v map[string]ApStatsPortStat)`

SetPortStat sets PortStat field to given value.

### HasPortStat

`func (o *StatsDevice) HasPortStat() bool`

HasPortStat returns a boolean if a field has been set.

### GetPowerBudget

`func (o *StatsDevice) GetPowerBudget() float32`

GetPowerBudget returns the PowerBudget field if non-nil, zero value otherwise.

### GetPowerBudgetOk

`func (o *StatsDevice) GetPowerBudgetOk() (*float32, bool)`

GetPowerBudgetOk returns a tuple with the PowerBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerBudget

`func (o *StatsDevice) SetPowerBudget(v float32)`

SetPowerBudget sets PowerBudget field to given value.

### HasPowerBudget

`func (o *StatsDevice) HasPowerBudget() bool`

HasPowerBudget returns a boolean if a field has been set.

### GetPowerConstrained

`func (o *StatsDevice) GetPowerConstrained() bool`

GetPowerConstrained returns the PowerConstrained field if non-nil, zero value otherwise.

### GetPowerConstrainedOk

`func (o *StatsDevice) GetPowerConstrainedOk() (*bool, bool)`

GetPowerConstrainedOk returns a tuple with the PowerConstrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConstrained

`func (o *StatsDevice) SetPowerConstrained(v bool)`

SetPowerConstrained sets PowerConstrained field to given value.

### HasPowerConstrained

`func (o *StatsDevice) HasPowerConstrained() bool`

HasPowerConstrained returns a boolean if a field has been set.

### GetPowerOpmode

`func (o *StatsDevice) GetPowerOpmode() string`

GetPowerOpmode returns the PowerOpmode field if non-nil, zero value otherwise.

### GetPowerOpmodeOk

`func (o *StatsDevice) GetPowerOpmodeOk() (*string, bool)`

GetPowerOpmodeOk returns a tuple with the PowerOpmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOpmode

`func (o *StatsDevice) SetPowerOpmode(v string)`

SetPowerOpmode sets PowerOpmode field to given value.

### HasPowerOpmode

`func (o *StatsDevice) HasPowerOpmode() bool`

HasPowerOpmode returns a boolean if a field has been set.

### GetPowerSrc

`func (o *StatsDevice) GetPowerSrc() string`

GetPowerSrc returns the PowerSrc field if non-nil, zero value otherwise.

### GetPowerSrcOk

`func (o *StatsDevice) GetPowerSrcOk() (*string, bool)`

GetPowerSrcOk returns a tuple with the PowerSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSrc

`func (o *StatsDevice) SetPowerSrc(v string)`

SetPowerSrc sets PowerSrc field to given value.

### HasPowerSrc

`func (o *StatsDevice) HasPowerSrc() bool`

HasPowerSrc returns a boolean if a field has been set.

### GetRadioConfig

`func (o *StatsDevice) GetRadioConfig() ApStatsRadioConfig`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *StatsDevice) GetRadioConfigOk() (*ApStatsRadioConfig, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *StatsDevice) SetRadioConfig(v ApStatsRadioConfig)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *StatsDevice) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetRadioStat

`func (o *StatsDevice) GetRadioStat() ApStatsRadioStat`

GetRadioStat returns the RadioStat field if non-nil, zero value otherwise.

### GetRadioStatOk

`func (o *StatsDevice) GetRadioStatOk() (*ApStatsRadioStat, bool)`

GetRadioStatOk returns a tuple with the RadioStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioStat

`func (o *StatsDevice) SetRadioStat(v ApStatsRadioStat)`

SetRadioStat sets RadioStat field to given value.

### HasRadioStat

`func (o *StatsDevice) HasRadioStat() bool`

HasRadioStat returns a boolean if a field has been set.

### GetRxBps

`func (o *StatsDevice) GetRxBps() float32`

GetRxBps returns the RxBps field if non-nil, zero value otherwise.

### GetRxBpsOk

`func (o *StatsDevice) GetRxBpsOk() (*float32, bool)`

GetRxBpsOk returns a tuple with the RxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBps

`func (o *StatsDevice) SetRxBps(v float32)`

SetRxBps sets RxBps field to given value.

### HasRxBps

`func (o *StatsDevice) HasRxBps() bool`

HasRxBps returns a boolean if a field has been set.

### GetRxBytes

`func (o *StatsDevice) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *StatsDevice) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *StatsDevice) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *StatsDevice) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *StatsDevice) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *StatsDevice) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *StatsDevice) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *StatsDevice) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSerial

`func (o *StatsDevice) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StatsDevice) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StatsDevice) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StatsDevice) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *StatsDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatsDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatsDevice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatsDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchRedundancy

`func (o *StatsDevice) GetSwitchRedundancy() ApStatsSwitchRedundancy`

GetSwitchRedundancy returns the SwitchRedundancy field if non-nil, zero value otherwise.

### GetSwitchRedundancyOk

`func (o *StatsDevice) GetSwitchRedundancyOk() (*ApStatsSwitchRedundancy, bool)`

GetSwitchRedundancyOk returns a tuple with the SwitchRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchRedundancy

`func (o *StatsDevice) SetSwitchRedundancy(v ApStatsSwitchRedundancy)`

SetSwitchRedundancy sets SwitchRedundancy field to given value.

### HasSwitchRedundancy

`func (o *StatsDevice) HasSwitchRedundancy() bool`

HasSwitchRedundancy returns a boolean if a field has been set.

### GetTxBps

`func (o *StatsDevice) GetTxBps() float32`

GetTxBps returns the TxBps field if non-nil, zero value otherwise.

### GetTxBpsOk

`func (o *StatsDevice) GetTxBpsOk() (*float32, bool)`

GetTxBpsOk returns a tuple with the TxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBps

`func (o *StatsDevice) SetTxBps(v float32)`

SetTxBps sets TxBps field to given value.

### HasTxBps

`func (o *StatsDevice) HasTxBps() bool`

HasTxBps returns a boolean if a field has been set.

### GetTxBytes

`func (o *StatsDevice) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *StatsDevice) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *StatsDevice) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *StatsDevice) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *StatsDevice) GetTxPkts() float32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *StatsDevice) GetTxPktsOk() (*float32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *StatsDevice) SetTxPkts(v float32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *StatsDevice) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetType

`func (o *StatsDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatsDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatsDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatsDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *StatsDevice) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *StatsDevice) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *StatsDevice) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *StatsDevice) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsbStat

`func (o *StatsDevice) GetUsbStat() ApStatsUsbStat`

GetUsbStat returns the UsbStat field if non-nil, zero value otherwise.

### GetUsbStatOk

`func (o *StatsDevice) GetUsbStatOk() (*ApStatsUsbStat, bool)`

GetUsbStatOk returns a tuple with the UsbStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbStat

`func (o *StatsDevice) SetUsbStat(v ApStatsUsbStat)`

SetUsbStat sets UsbStat field to given value.

### HasUsbStat

`func (o *StatsDevice) HasUsbStat() bool`

HasUsbStat returns a boolean if a field has been set.

### GetVersion

`func (o *StatsDevice) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatsDevice) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatsDevice) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StatsDevice) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetX

`func (o *StatsDevice) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *StatsDevice) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *StatsDevice) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *StatsDevice) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *StatsDevice) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *StatsDevice) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *StatsDevice) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *StatsDevice) HasY() bool`

HasY returns a boolean if a field has been set.

### GetApRedundancy

`func (o *StatsDevice) GetApRedundancy() ApRedundancy`

GetApRedundancy returns the ApRedundancy field if non-nil, zero value otherwise.

### GetApRedundancyOk

`func (o *StatsDevice) GetApRedundancyOk() (*ApRedundancy, bool)`

GetApRedundancyOk returns a tuple with the ApRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRedundancy

`func (o *StatsDevice) SetApRedundancy(v ApRedundancy)`

SetApRedundancy sets ApRedundancy field to given value.

### HasApRedundancy

`func (o *StatsDevice) HasApRedundancy() bool`

HasApRedundancy returns a boolean if a field has been set.

### GetClients

`func (o *StatsDevice) GetClients() []SwitchStatsClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *StatsDevice) GetClientsOk() (*[]SwitchStatsClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *StatsDevice) SetClients(v []SwitchStatsClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *StatsDevice) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCpuStat

`func (o *StatsDevice) GetCpuStat() CpuStat`

GetCpuStat returns the CpuStat field if non-nil, zero value otherwise.

### GetCpuStatOk

`func (o *StatsDevice) GetCpuStatOk() (*CpuStat, bool)`

GetCpuStatOk returns a tuple with the CpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStat

`func (o *StatsDevice) SetCpuStat(v CpuStat)`

SetCpuStat sets CpuStat field to given value.

### HasCpuStat

`func (o *StatsDevice) HasCpuStat() bool`

HasCpuStat returns a boolean if a field has been set.

### GetHasPcap

`func (o *StatsDevice) GetHasPcap() bool`

GetHasPcap returns the HasPcap field if non-nil, zero value otherwise.

### GetHasPcapOk

`func (o *StatsDevice) GetHasPcapOk() (*bool, bool)`

GetHasPcapOk returns a tuple with the HasPcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPcap

`func (o *StatsDevice) SetHasPcap(v bool)`

SetHasPcap sets HasPcap field to given value.

### HasHasPcap

`func (o *StatsDevice) HasHasPcap() bool`

HasHasPcap returns a boolean if a field has been set.

### GetHostname

`func (o *StatsDevice) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *StatsDevice) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *StatsDevice) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *StatsDevice) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIfStat

`func (o *StatsDevice) GetIfStat() map[string]SwitchStatsIfStat`

GetIfStat returns the IfStat field if non-nil, zero value otherwise.

### GetIfStatOk

`func (o *StatsDevice) GetIfStatOk() (*map[string]SwitchStatsIfStat, bool)`

GetIfStatOk returns a tuple with the IfStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfStat

`func (o *StatsDevice) SetIfStat(v map[string]SwitchStatsIfStat)`

SetIfStat sets IfStat field to given value.

### HasIfStat

`func (o *StatsDevice) HasIfStat() bool`

HasIfStat returns a boolean if a field has been set.

### GetModuleStat

`func (o *StatsDevice) GetModuleStat() []ModuleStat`

GetModuleStat returns the ModuleStat field if non-nil, zero value otherwise.

### GetModuleStatOk

`func (o *StatsDevice) GetModuleStatOk() (*[]ModuleStat, bool)`

GetModuleStatOk returns a tuple with the ModuleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStat

`func (o *StatsDevice) SetModuleStat(v []ModuleStat)`

SetModuleStat sets ModuleStat field to given value.

### HasModuleStat

`func (o *StatsDevice) HasModuleStat() bool`

HasModuleStat returns a boolean if a field has been set.

### GetClusterStat

`func (o *StatsDevice) GetClusterStat() map[string]GatewayStatsClusterStat`

GetClusterStat returns the ClusterStat field if non-nil, zero value otherwise.

### GetClusterStatOk

`func (o *StatsDevice) GetClusterStatOk() (*map[string]GatewayStatsClusterStat, bool)`

GetClusterStatOk returns a tuple with the ClusterStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStat

`func (o *StatsDevice) SetClusterStat(v map[string]GatewayStatsClusterStat)`

SetClusterStat sets ClusterStat field to given value.

### HasClusterStat

`func (o *StatsDevice) HasClusterStat() bool`

HasClusterStat returns a boolean if a field has been set.

### GetCpu2Stat

`func (o *StatsDevice) GetCpu2Stat() CpuStat`

GetCpu2Stat returns the Cpu2Stat field if non-nil, zero value otherwise.

### GetCpu2StatOk

`func (o *StatsDevice) GetCpu2StatOk() (*CpuStat, bool)`

GetCpu2StatOk returns a tuple with the Cpu2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu2Stat

`func (o *StatsDevice) SetCpu2Stat(v CpuStat)`

SetCpu2Stat sets Cpu2Stat field to given value.

### HasCpu2Stat

`func (o *StatsDevice) HasCpu2Stat() bool`

HasCpu2Stat returns a boolean if a field has been set.

### GetDhcpd2Stat

`func (o *StatsDevice) GetDhcpd2Stat() map[string]GatewayStatsDhcpdStatLan`

GetDhcpd2Stat returns the Dhcpd2Stat field if non-nil, zero value otherwise.

### GetDhcpd2StatOk

`func (o *StatsDevice) GetDhcpd2StatOk() (*map[string]GatewayStatsDhcpdStatLan, bool)`

GetDhcpd2StatOk returns a tuple with the Dhcpd2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpd2Stat

`func (o *StatsDevice) SetDhcpd2Stat(v map[string]GatewayStatsDhcpdStatLan)`

SetDhcpd2Stat sets Dhcpd2Stat field to given value.

### HasDhcpd2Stat

`func (o *StatsDevice) HasDhcpd2Stat() bool`

HasDhcpd2Stat returns a boolean if a field has been set.

### GetDhcpdStat

`func (o *StatsDevice) GetDhcpdStat() map[string]GatewayStatsDhcpdStatLan`

GetDhcpdStat returns the DhcpdStat field if non-nil, zero value otherwise.

### GetDhcpdStatOk

`func (o *StatsDevice) GetDhcpdStatOk() (*map[string]GatewayStatsDhcpdStatLan, bool)`

GetDhcpdStatOk returns a tuple with the DhcpdStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdStat

`func (o *StatsDevice) SetDhcpdStat(v map[string]GatewayStatsDhcpdStatLan)`

SetDhcpdStat sets DhcpdStat field to given value.

### HasDhcpdStat

`func (o *StatsDevice) HasDhcpdStat() bool`

HasDhcpdStat returns a boolean if a field has been set.

### GetIp2Stat

`func (o *StatsDevice) GetIp2Stat() IpStat`

GetIp2Stat returns the Ip2Stat field if non-nil, zero value otherwise.

### GetIp2StatOk

`func (o *StatsDevice) GetIp2StatOk() (*IpStat, bool)`

GetIp2StatOk returns a tuple with the Ip2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2Stat

`func (o *StatsDevice) SetIp2Stat(v IpStat)`

SetIp2Stat sets Ip2Stat field to given value.

### HasIp2Stat

`func (o *StatsDevice) HasIp2Stat() bool`

HasIp2Stat returns a boolean if a field has been set.

### GetMemory2Stat

`func (o *StatsDevice) GetMemory2Stat() MemoryStat`

GetMemory2Stat returns the Memory2Stat field if non-nil, zero value otherwise.

### GetMemory2StatOk

`func (o *StatsDevice) GetMemory2StatOk() (*MemoryStat, bool)`

GetMemory2StatOk returns a tuple with the Memory2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory2Stat

`func (o *StatsDevice) SetMemory2Stat(v MemoryStat)`

SetMemory2Stat sets Memory2Stat field to given value.

### HasMemory2Stat

`func (o *StatsDevice) HasMemory2Stat() bool`

HasMemory2Stat returns a boolean if a field has been set.

### GetMemoryStat

`func (o *StatsDevice) GetMemoryStat() MemoryStat`

GetMemoryStat returns the MemoryStat field if non-nil, zero value otherwise.

### GetMemoryStatOk

`func (o *StatsDevice) GetMemoryStatOk() (*MemoryStat, bool)`

GetMemoryStatOk returns a tuple with the MemoryStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryStat

`func (o *StatsDevice) SetMemoryStat(v MemoryStat)`

SetMemoryStat sets MemoryStat field to given value.

### HasMemoryStat

`func (o *StatsDevice) HasMemoryStat() bool`

HasMemoryStat returns a boolean if a field has been set.

### GetModule2Stat

`func (o *StatsDevice) GetModule2Stat() []ModuleStat`

GetModule2Stat returns the Module2Stat field if non-nil, zero value otherwise.

### GetModule2StatOk

`func (o *StatsDevice) GetModule2StatOk() (*[]ModuleStat, bool)`

GetModule2StatOk returns a tuple with the Module2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule2Stat

`func (o *StatsDevice) SetModule2Stat(v []ModuleStat)`

SetModule2Stat sets Module2Stat field to given value.

### HasModule2Stat

`func (o *StatsDevice) HasModule2Stat() bool`

HasModule2Stat returns a boolean if a field has been set.

### GetSpu2Stat

`func (o *StatsDevice) GetSpu2Stat() GatewayStatsSpuStat`

GetSpu2Stat returns the Spu2Stat field if non-nil, zero value otherwise.

### GetSpu2StatOk

`func (o *StatsDevice) GetSpu2StatOk() (*GatewayStatsSpuStat, bool)`

GetSpu2StatOk returns a tuple with the Spu2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpu2Stat

`func (o *StatsDevice) SetSpu2Stat(v GatewayStatsSpuStat)`

SetSpu2Stat sets Spu2Stat field to given value.

### HasSpu2Stat

`func (o *StatsDevice) HasSpu2Stat() bool`

HasSpu2Stat returns a boolean if a field has been set.

### GetSpuStat

`func (o *StatsDevice) GetSpuStat() GatewayStatsSpuStat`

GetSpuStat returns the SpuStat field if non-nil, zero value otherwise.

### GetSpuStatOk

`func (o *StatsDevice) GetSpuStatOk() (*GatewayStatsSpuStat, bool)`

GetSpuStatOk returns a tuple with the SpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpuStat

`func (o *StatsDevice) SetSpuStat(v GatewayStatsSpuStat)`

SetSpuStat sets SpuStat field to given value.

### HasSpuStat

`func (o *StatsDevice) HasSpuStat() bool`

HasSpuStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


