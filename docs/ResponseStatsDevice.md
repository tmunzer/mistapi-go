# ResponseStatsDevice

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

### NewResponseStatsDevice

`func NewResponseStatsDevice(mac string, model string, ) *ResponseStatsDevice`

NewResponseStatsDevice instantiates a new ResponseStatsDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStatsDeviceWithDefaults

`func NewResponseStatsDeviceWithDefaults() *ResponseStatsDevice`

NewResponseStatsDeviceWithDefaults instantiates a new ResponseStatsDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPlacement

`func (o *ResponseStatsDevice) GetAutoPlacement() ApStatsAutoPlacement`

GetAutoPlacement returns the AutoPlacement field if non-nil, zero value otherwise.

### GetAutoPlacementOk

`func (o *ResponseStatsDevice) GetAutoPlacementOk() (*ApStatsAutoPlacement, bool)`

GetAutoPlacementOk returns a tuple with the AutoPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlacement

`func (o *ResponseStatsDevice) SetAutoPlacement(v ApStatsAutoPlacement)`

SetAutoPlacement sets AutoPlacement field to given value.

### HasAutoPlacement

`func (o *ResponseStatsDevice) HasAutoPlacement() bool`

HasAutoPlacement returns a boolean if a field has been set.

### GetBleConfig

`func (o *ResponseStatsDevice) GetBleConfig() ApStatsBleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *ResponseStatsDevice) GetBleConfigOk() (*ApStatsBleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *ResponseStatsDevice) SetBleConfig(v ApStatsBleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *ResponseStatsDevice) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetBleStat

`func (o *ResponseStatsDevice) GetBleStat() ApStatsBleStat`

GetBleStat returns the BleStat field if non-nil, zero value otherwise.

### GetBleStatOk

`func (o *ResponseStatsDevice) GetBleStatOk() (*ApStatsBleStat, bool)`

GetBleStatOk returns a tuple with the BleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleStat

`func (o *ResponseStatsDevice) SetBleStat(v ApStatsBleStat)`

SetBleStat sets BleStat field to given value.

### HasBleStat

`func (o *ResponseStatsDevice) HasBleStat() bool`

HasBleStat returns a boolean if a field has been set.

### GetCertExpiry

`func (o *ResponseStatsDevice) GetCertExpiry() float32`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *ResponseStatsDevice) GetCertExpiryOk() (*float32, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *ResponseStatsDevice) SetCertExpiry(v float32)`

SetCertExpiry sets CertExpiry field to given value.

### HasCertExpiry

`func (o *ResponseStatsDevice) HasCertExpiry() bool`

HasCertExpiry returns a boolean if a field has been set.

### GetEnvStat

`func (o *ResponseStatsDevice) GetEnvStat() ApStatsEnvStat`

GetEnvStat returns the EnvStat field if non-nil, zero value otherwise.

### GetEnvStatOk

`func (o *ResponseStatsDevice) GetEnvStatOk() (*ApStatsEnvStat, bool)`

GetEnvStatOk returns a tuple with the EnvStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvStat

`func (o *ResponseStatsDevice) SetEnvStat(v ApStatsEnvStat)`

SetEnvStat sets EnvStat field to given value.

### HasEnvStat

`func (o *ResponseStatsDevice) HasEnvStat() bool`

HasEnvStat returns a boolean if a field has been set.

### GetEslStat

`func (o *ResponseStatsDevice) GetEslStat() ApStatsEslStat`

GetEslStat returns the EslStat field if non-nil, zero value otherwise.

### GetEslStatOk

`func (o *ResponseStatsDevice) GetEslStatOk() (*ApStatsEslStat, bool)`

GetEslStatOk returns a tuple with the EslStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslStat

`func (o *ResponseStatsDevice) SetEslStat(v ApStatsEslStat)`

SetEslStat sets EslStat field to given value.

### HasEslStat

`func (o *ResponseStatsDevice) HasEslStat() bool`

HasEslStat returns a boolean if a field has been set.

### GetExtIp

`func (o *ResponseStatsDevice) GetExtIp() string`

GetExtIp returns the ExtIp field if non-nil, zero value otherwise.

### GetExtIpOk

`func (o *ResponseStatsDevice) GetExtIpOk() (*string, bool)`

GetExtIpOk returns a tuple with the ExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIp

`func (o *ResponseStatsDevice) SetExtIp(v string)`

SetExtIp sets ExtIp field to given value.

### HasExtIp

`func (o *ResponseStatsDevice) HasExtIp() bool`

HasExtIp returns a boolean if a field has been set.

### GetFwupdate

`func (o *ResponseStatsDevice) GetFwupdate() ApStatsFwupdate`

GetFwupdate returns the Fwupdate field if non-nil, zero value otherwise.

### GetFwupdateOk

`func (o *ResponseStatsDevice) GetFwupdateOk() (*ApStatsFwupdate, bool)`

GetFwupdateOk returns a tuple with the Fwupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwupdate

`func (o *ResponseStatsDevice) SetFwupdate(v ApStatsFwupdate)`

SetFwupdate sets Fwupdate field to given value.

### HasFwupdate

`func (o *ResponseStatsDevice) HasFwupdate() bool`

HasFwupdate returns a boolean if a field has been set.

### GetIotStat

`func (o *ResponseStatsDevice) GetIotStat() map[string]ApStatsIotStatAdditionalProperties`

GetIotStat returns the IotStat field if non-nil, zero value otherwise.

### GetIotStatOk

`func (o *ResponseStatsDevice) GetIotStatOk() (*map[string]ApStatsIotStatAdditionalProperties, bool)`

GetIotStatOk returns a tuple with the IotStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotStat

`func (o *ResponseStatsDevice) SetIotStat(v map[string]ApStatsIotStatAdditionalProperties)`

SetIotStat sets IotStat field to given value.

### HasIotStat

`func (o *ResponseStatsDevice) HasIotStat() bool`

HasIotStat returns a boolean if a field has been set.

### GetIp

`func (o *ResponseStatsDevice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ResponseStatsDevice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ResponseStatsDevice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ResponseStatsDevice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpConfig

`func (o *ResponseStatsDevice) GetIpConfig() ApIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ResponseStatsDevice) GetIpConfigOk() (*ApIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ResponseStatsDevice) SetIpConfig(v ApIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ResponseStatsDevice) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetIpStat

`func (o *ResponseStatsDevice) GetIpStat() IpStat`

GetIpStat returns the IpStat field if non-nil, zero value otherwise.

### GetIpStatOk

`func (o *ResponseStatsDevice) GetIpStatOk() (*IpStat, bool)`

GetIpStatOk returns a tuple with the IpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStat

`func (o *ResponseStatsDevice) SetIpStat(v IpStat)`

SetIpStat sets IpStat field to given value.

### HasIpStat

`func (o *ResponseStatsDevice) HasIpStat() bool`

HasIpStat returns a boolean if a field has been set.

### GetL2tpStat

`func (o *ResponseStatsDevice) GetL2tpStat() map[string]ApStatsL2tpStat`

GetL2tpStat returns the L2tpStat field if non-nil, zero value otherwise.

### GetL2tpStatOk

`func (o *ResponseStatsDevice) GetL2tpStatOk() (*map[string]ApStatsL2tpStat, bool)`

GetL2tpStatOk returns a tuple with the L2tpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2tpStat

`func (o *ResponseStatsDevice) SetL2tpStat(v map[string]ApStatsL2tpStat)`

SetL2tpStat sets L2tpStat field to given value.

### HasL2tpStat

`func (o *ResponseStatsDevice) HasL2tpStat() bool`

HasL2tpStat returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseStatsDevice) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseStatsDevice) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseStatsDevice) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseStatsDevice) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastTrouble

`func (o *ResponseStatsDevice) GetLastTrouble() LastTrouble`

GetLastTrouble returns the LastTrouble field if non-nil, zero value otherwise.

### GetLastTroubleOk

`func (o *ResponseStatsDevice) GetLastTroubleOk() (*LastTrouble, bool)`

GetLastTroubleOk returns a tuple with the LastTrouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrouble

`func (o *ResponseStatsDevice) SetLastTrouble(v LastTrouble)`

SetLastTrouble sets LastTrouble field to given value.

### HasLastTrouble

`func (o *ResponseStatsDevice) HasLastTrouble() bool`

HasLastTrouble returns a boolean if a field has been set.

### GetLed

`func (o *ResponseStatsDevice) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *ResponseStatsDevice) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *ResponseStatsDevice) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *ResponseStatsDevice) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLldpStat

`func (o *ResponseStatsDevice) GetLldpStat() ApStatsLldpStat`

GetLldpStat returns the LldpStat field if non-nil, zero value otherwise.

### GetLldpStatOk

`func (o *ResponseStatsDevice) GetLldpStatOk() (*ApStatsLldpStat, bool)`

GetLldpStatOk returns a tuple with the LldpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpStat

`func (o *ResponseStatsDevice) SetLldpStat(v ApStatsLldpStat)`

SetLldpStat sets LldpStat field to given value.

### HasLldpStat

`func (o *ResponseStatsDevice) HasLldpStat() bool`

HasLldpStat returns a boolean if a field has been set.

### GetLocating

`func (o *ResponseStatsDevice) GetLocating() bool`

GetLocating returns the Locating field if non-nil, zero value otherwise.

### GetLocatingOk

`func (o *ResponseStatsDevice) GetLocatingOk() (*bool, bool)`

GetLocatingOk returns a tuple with the Locating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocating

`func (o *ResponseStatsDevice) SetLocating(v bool)`

SetLocating sets Locating field to given value.

### HasLocating

`func (o *ResponseStatsDevice) HasLocating() bool`

HasLocating returns a boolean if a field has been set.

### GetLocked

`func (o *ResponseStatsDevice) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ResponseStatsDevice) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ResponseStatsDevice) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ResponseStatsDevice) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMac

`func (o *ResponseStatsDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ResponseStatsDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ResponseStatsDevice) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *ResponseStatsDevice) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ResponseStatsDevice) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ResponseStatsDevice) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *ResponseStatsDevice) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMeshDownlinks

`func (o *ResponseStatsDevice) GetMeshDownlinks() map[string]ApStatMeshDownlink`

GetMeshDownlinks returns the MeshDownlinks field if non-nil, zero value otherwise.

### GetMeshDownlinksOk

`func (o *ResponseStatsDevice) GetMeshDownlinksOk() (*map[string]ApStatMeshDownlink, bool)`

GetMeshDownlinksOk returns a tuple with the MeshDownlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshDownlinks

`func (o *ResponseStatsDevice) SetMeshDownlinks(v map[string]ApStatMeshDownlink)`

SetMeshDownlinks sets MeshDownlinks field to given value.

### HasMeshDownlinks

`func (o *ResponseStatsDevice) HasMeshDownlinks() bool`

HasMeshDownlinks returns a boolean if a field has been set.

### GetMeshUplink

`func (o *ResponseStatsDevice) GetMeshUplink() ApStatMeshUplink`

GetMeshUplink returns the MeshUplink field if non-nil, zero value otherwise.

### GetMeshUplinkOk

`func (o *ResponseStatsDevice) GetMeshUplinkOk() (*ApStatMeshUplink, bool)`

GetMeshUplinkOk returns a tuple with the MeshUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshUplink

`func (o *ResponseStatsDevice) SetMeshUplink(v ApStatMeshUplink)`

SetMeshUplink sets MeshUplink field to given value.

### HasMeshUplink

`func (o *ResponseStatsDevice) HasMeshUplink() bool`

HasMeshUplink returns a boolean if a field has been set.

### GetModel

`func (o *ResponseStatsDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResponseStatsDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResponseStatsDevice) SetModel(v string)`

SetModel sets Model field to given value.


### GetMount

`func (o *ResponseStatsDevice) GetMount() string`

GetMount returns the Mount field if non-nil, zero value otherwise.

### GetMountOk

`func (o *ResponseStatsDevice) GetMountOk() (*string, bool)`

GetMountOk returns a tuple with the Mount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMount

`func (o *ResponseStatsDevice) SetMount(v string)`

SetMount sets Mount field to given value.

### HasMount

`func (o *ResponseStatsDevice) HasMount() bool`

HasMount returns a boolean if a field has been set.

### GetName

`func (o *ResponseStatsDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseStatsDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseStatsDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseStatsDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumClients

`func (o *ResponseStatsDevice) GetNumClients() SwitchStatsNumClients`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *ResponseStatsDevice) GetNumClientsOk() (*SwitchStatsNumClients, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *ResponseStatsDevice) SetNumClients(v SwitchStatsNumClients)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *ResponseStatsDevice) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetPortStat

`func (o *ResponseStatsDevice) GetPortStat() map[string]ApStatsPortStat`

GetPortStat returns the PortStat field if non-nil, zero value otherwise.

### GetPortStatOk

`func (o *ResponseStatsDevice) GetPortStatOk() (*map[string]ApStatsPortStat, bool)`

GetPortStatOk returns a tuple with the PortStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStat

`func (o *ResponseStatsDevice) SetPortStat(v map[string]ApStatsPortStat)`

SetPortStat sets PortStat field to given value.

### HasPortStat

`func (o *ResponseStatsDevice) HasPortStat() bool`

HasPortStat returns a boolean if a field has been set.

### GetPowerBudget

`func (o *ResponseStatsDevice) GetPowerBudget() float32`

GetPowerBudget returns the PowerBudget field if non-nil, zero value otherwise.

### GetPowerBudgetOk

`func (o *ResponseStatsDevice) GetPowerBudgetOk() (*float32, bool)`

GetPowerBudgetOk returns a tuple with the PowerBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerBudget

`func (o *ResponseStatsDevice) SetPowerBudget(v float32)`

SetPowerBudget sets PowerBudget field to given value.

### HasPowerBudget

`func (o *ResponseStatsDevice) HasPowerBudget() bool`

HasPowerBudget returns a boolean if a field has been set.

### GetPowerConstrained

`func (o *ResponseStatsDevice) GetPowerConstrained() bool`

GetPowerConstrained returns the PowerConstrained field if non-nil, zero value otherwise.

### GetPowerConstrainedOk

`func (o *ResponseStatsDevice) GetPowerConstrainedOk() (*bool, bool)`

GetPowerConstrainedOk returns a tuple with the PowerConstrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConstrained

`func (o *ResponseStatsDevice) SetPowerConstrained(v bool)`

SetPowerConstrained sets PowerConstrained field to given value.

### HasPowerConstrained

`func (o *ResponseStatsDevice) HasPowerConstrained() bool`

HasPowerConstrained returns a boolean if a field has been set.

### GetPowerOpmode

`func (o *ResponseStatsDevice) GetPowerOpmode() string`

GetPowerOpmode returns the PowerOpmode field if non-nil, zero value otherwise.

### GetPowerOpmodeOk

`func (o *ResponseStatsDevice) GetPowerOpmodeOk() (*string, bool)`

GetPowerOpmodeOk returns a tuple with the PowerOpmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOpmode

`func (o *ResponseStatsDevice) SetPowerOpmode(v string)`

SetPowerOpmode sets PowerOpmode field to given value.

### HasPowerOpmode

`func (o *ResponseStatsDevice) HasPowerOpmode() bool`

HasPowerOpmode returns a boolean if a field has been set.

### GetPowerSrc

`func (o *ResponseStatsDevice) GetPowerSrc() string`

GetPowerSrc returns the PowerSrc field if non-nil, zero value otherwise.

### GetPowerSrcOk

`func (o *ResponseStatsDevice) GetPowerSrcOk() (*string, bool)`

GetPowerSrcOk returns a tuple with the PowerSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSrc

`func (o *ResponseStatsDevice) SetPowerSrc(v string)`

SetPowerSrc sets PowerSrc field to given value.

### HasPowerSrc

`func (o *ResponseStatsDevice) HasPowerSrc() bool`

HasPowerSrc returns a boolean if a field has been set.

### GetRadioConfig

`func (o *ResponseStatsDevice) GetRadioConfig() ApStatsRadioConfig`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *ResponseStatsDevice) GetRadioConfigOk() (*ApStatsRadioConfig, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *ResponseStatsDevice) SetRadioConfig(v ApStatsRadioConfig)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *ResponseStatsDevice) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetRadioStat

`func (o *ResponseStatsDevice) GetRadioStat() ApStatsRadioStat`

GetRadioStat returns the RadioStat field if non-nil, zero value otherwise.

### GetRadioStatOk

`func (o *ResponseStatsDevice) GetRadioStatOk() (*ApStatsRadioStat, bool)`

GetRadioStatOk returns a tuple with the RadioStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioStat

`func (o *ResponseStatsDevice) SetRadioStat(v ApStatsRadioStat)`

SetRadioStat sets RadioStat field to given value.

### HasRadioStat

`func (o *ResponseStatsDevice) HasRadioStat() bool`

HasRadioStat returns a boolean if a field has been set.

### GetRxBps

`func (o *ResponseStatsDevice) GetRxBps() float32`

GetRxBps returns the RxBps field if non-nil, zero value otherwise.

### GetRxBpsOk

`func (o *ResponseStatsDevice) GetRxBpsOk() (*float32, bool)`

GetRxBpsOk returns a tuple with the RxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBps

`func (o *ResponseStatsDevice) SetRxBps(v float32)`

SetRxBps sets RxBps field to given value.

### HasRxBps

`func (o *ResponseStatsDevice) HasRxBps() bool`

HasRxBps returns a boolean if a field has been set.

### GetRxBytes

`func (o *ResponseStatsDevice) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ResponseStatsDevice) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ResponseStatsDevice) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ResponseStatsDevice) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *ResponseStatsDevice) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ResponseStatsDevice) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ResponseStatsDevice) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ResponseStatsDevice) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSerial

`func (o *ResponseStatsDevice) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ResponseStatsDevice) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ResponseStatsDevice) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ResponseStatsDevice) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseStatsDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseStatsDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseStatsDevice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseStatsDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchRedundancy

`func (o *ResponseStatsDevice) GetSwitchRedundancy() ApStatsSwitchRedundancy`

GetSwitchRedundancy returns the SwitchRedundancy field if non-nil, zero value otherwise.

### GetSwitchRedundancyOk

`func (o *ResponseStatsDevice) GetSwitchRedundancyOk() (*ApStatsSwitchRedundancy, bool)`

GetSwitchRedundancyOk returns a tuple with the SwitchRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchRedundancy

`func (o *ResponseStatsDevice) SetSwitchRedundancy(v ApStatsSwitchRedundancy)`

SetSwitchRedundancy sets SwitchRedundancy field to given value.

### HasSwitchRedundancy

`func (o *ResponseStatsDevice) HasSwitchRedundancy() bool`

HasSwitchRedundancy returns a boolean if a field has been set.

### GetTxBps

`func (o *ResponseStatsDevice) GetTxBps() float32`

GetTxBps returns the TxBps field if non-nil, zero value otherwise.

### GetTxBpsOk

`func (o *ResponseStatsDevice) GetTxBpsOk() (*float32, bool)`

GetTxBpsOk returns a tuple with the TxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBps

`func (o *ResponseStatsDevice) SetTxBps(v float32)`

SetTxBps sets TxBps field to given value.

### HasTxBps

`func (o *ResponseStatsDevice) HasTxBps() bool`

HasTxBps returns a boolean if a field has been set.

### GetTxBytes

`func (o *ResponseStatsDevice) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ResponseStatsDevice) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ResponseStatsDevice) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ResponseStatsDevice) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *ResponseStatsDevice) GetTxPkts() float32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ResponseStatsDevice) GetTxPktsOk() (*float32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ResponseStatsDevice) SetTxPkts(v float32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ResponseStatsDevice) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetType

`func (o *ResponseStatsDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseStatsDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseStatsDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseStatsDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *ResponseStatsDevice) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ResponseStatsDevice) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ResponseStatsDevice) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ResponseStatsDevice) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsbStat

`func (o *ResponseStatsDevice) GetUsbStat() ApStatsUsbStat`

GetUsbStat returns the UsbStat field if non-nil, zero value otherwise.

### GetUsbStatOk

`func (o *ResponseStatsDevice) GetUsbStatOk() (*ApStatsUsbStat, bool)`

GetUsbStatOk returns a tuple with the UsbStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbStat

`func (o *ResponseStatsDevice) SetUsbStat(v ApStatsUsbStat)`

SetUsbStat sets UsbStat field to given value.

### HasUsbStat

`func (o *ResponseStatsDevice) HasUsbStat() bool`

HasUsbStat returns a boolean if a field has been set.

### GetVersion

`func (o *ResponseStatsDevice) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseStatsDevice) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseStatsDevice) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseStatsDevice) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetX

`func (o *ResponseStatsDevice) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ResponseStatsDevice) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ResponseStatsDevice) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *ResponseStatsDevice) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ResponseStatsDevice) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ResponseStatsDevice) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ResponseStatsDevice) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *ResponseStatsDevice) HasY() bool`

HasY returns a boolean if a field has been set.

### GetApRedundancy

`func (o *ResponseStatsDevice) GetApRedundancy() ApRedundancy`

GetApRedundancy returns the ApRedundancy field if non-nil, zero value otherwise.

### GetApRedundancyOk

`func (o *ResponseStatsDevice) GetApRedundancyOk() (*ApRedundancy, bool)`

GetApRedundancyOk returns a tuple with the ApRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRedundancy

`func (o *ResponseStatsDevice) SetApRedundancy(v ApRedundancy)`

SetApRedundancy sets ApRedundancy field to given value.

### HasApRedundancy

`func (o *ResponseStatsDevice) HasApRedundancy() bool`

HasApRedundancy returns a boolean if a field has been set.

### GetClients

`func (o *ResponseStatsDevice) GetClients() []SwitchStatsClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ResponseStatsDevice) GetClientsOk() (*[]SwitchStatsClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ResponseStatsDevice) SetClients(v []SwitchStatsClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ResponseStatsDevice) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCpuStat

`func (o *ResponseStatsDevice) GetCpuStat() CpuStat`

GetCpuStat returns the CpuStat field if non-nil, zero value otherwise.

### GetCpuStatOk

`func (o *ResponseStatsDevice) GetCpuStatOk() (*CpuStat, bool)`

GetCpuStatOk returns a tuple with the CpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStat

`func (o *ResponseStatsDevice) SetCpuStat(v CpuStat)`

SetCpuStat sets CpuStat field to given value.

### HasCpuStat

`func (o *ResponseStatsDevice) HasCpuStat() bool`

HasCpuStat returns a boolean if a field has been set.

### GetHasPcap

`func (o *ResponseStatsDevice) GetHasPcap() bool`

GetHasPcap returns the HasPcap field if non-nil, zero value otherwise.

### GetHasPcapOk

`func (o *ResponseStatsDevice) GetHasPcapOk() (*bool, bool)`

GetHasPcapOk returns a tuple with the HasPcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPcap

`func (o *ResponseStatsDevice) SetHasPcap(v bool)`

SetHasPcap sets HasPcap field to given value.

### HasHasPcap

`func (o *ResponseStatsDevice) HasHasPcap() bool`

HasHasPcap returns a boolean if a field has been set.

### GetHostname

`func (o *ResponseStatsDevice) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ResponseStatsDevice) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ResponseStatsDevice) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ResponseStatsDevice) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIfStat

`func (o *ResponseStatsDevice) GetIfStat() map[string]SwitchStatsIfStat`

GetIfStat returns the IfStat field if non-nil, zero value otherwise.

### GetIfStatOk

`func (o *ResponseStatsDevice) GetIfStatOk() (*map[string]SwitchStatsIfStat, bool)`

GetIfStatOk returns a tuple with the IfStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfStat

`func (o *ResponseStatsDevice) SetIfStat(v map[string]SwitchStatsIfStat)`

SetIfStat sets IfStat field to given value.

### HasIfStat

`func (o *ResponseStatsDevice) HasIfStat() bool`

HasIfStat returns a boolean if a field has been set.

### GetModuleStat

`func (o *ResponseStatsDevice) GetModuleStat() []ModuleStat`

GetModuleStat returns the ModuleStat field if non-nil, zero value otherwise.

### GetModuleStatOk

`func (o *ResponseStatsDevice) GetModuleStatOk() (*[]ModuleStat, bool)`

GetModuleStatOk returns a tuple with the ModuleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStat

`func (o *ResponseStatsDevice) SetModuleStat(v []ModuleStat)`

SetModuleStat sets ModuleStat field to given value.

### HasModuleStat

`func (o *ResponseStatsDevice) HasModuleStat() bool`

HasModuleStat returns a boolean if a field has been set.

### GetClusterStat

`func (o *ResponseStatsDevice) GetClusterStat() map[string]GatewayStatsClusterStat`

GetClusterStat returns the ClusterStat field if non-nil, zero value otherwise.

### GetClusterStatOk

`func (o *ResponseStatsDevice) GetClusterStatOk() (*map[string]GatewayStatsClusterStat, bool)`

GetClusterStatOk returns a tuple with the ClusterStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStat

`func (o *ResponseStatsDevice) SetClusterStat(v map[string]GatewayStatsClusterStat)`

SetClusterStat sets ClusterStat field to given value.

### HasClusterStat

`func (o *ResponseStatsDevice) HasClusterStat() bool`

HasClusterStat returns a boolean if a field has been set.

### GetCpu2Stat

`func (o *ResponseStatsDevice) GetCpu2Stat() CpuStat`

GetCpu2Stat returns the Cpu2Stat field if non-nil, zero value otherwise.

### GetCpu2StatOk

`func (o *ResponseStatsDevice) GetCpu2StatOk() (*CpuStat, bool)`

GetCpu2StatOk returns a tuple with the Cpu2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu2Stat

`func (o *ResponseStatsDevice) SetCpu2Stat(v CpuStat)`

SetCpu2Stat sets Cpu2Stat field to given value.

### HasCpu2Stat

`func (o *ResponseStatsDevice) HasCpu2Stat() bool`

HasCpu2Stat returns a boolean if a field has been set.

### GetDhcpd2Stat

`func (o *ResponseStatsDevice) GetDhcpd2Stat() map[string]GatewayStatsDhcpdStatLan`

GetDhcpd2Stat returns the Dhcpd2Stat field if non-nil, zero value otherwise.

### GetDhcpd2StatOk

`func (o *ResponseStatsDevice) GetDhcpd2StatOk() (*map[string]GatewayStatsDhcpdStatLan, bool)`

GetDhcpd2StatOk returns a tuple with the Dhcpd2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpd2Stat

`func (o *ResponseStatsDevice) SetDhcpd2Stat(v map[string]GatewayStatsDhcpdStatLan)`

SetDhcpd2Stat sets Dhcpd2Stat field to given value.

### HasDhcpd2Stat

`func (o *ResponseStatsDevice) HasDhcpd2Stat() bool`

HasDhcpd2Stat returns a boolean if a field has been set.

### GetDhcpdStat

`func (o *ResponseStatsDevice) GetDhcpdStat() map[string]GatewayStatsDhcpdStatLan`

GetDhcpdStat returns the DhcpdStat field if non-nil, zero value otherwise.

### GetDhcpdStatOk

`func (o *ResponseStatsDevice) GetDhcpdStatOk() (*map[string]GatewayStatsDhcpdStatLan, bool)`

GetDhcpdStatOk returns a tuple with the DhcpdStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdStat

`func (o *ResponseStatsDevice) SetDhcpdStat(v map[string]GatewayStatsDhcpdStatLan)`

SetDhcpdStat sets DhcpdStat field to given value.

### HasDhcpdStat

`func (o *ResponseStatsDevice) HasDhcpdStat() bool`

HasDhcpdStat returns a boolean if a field has been set.

### GetIp2Stat

`func (o *ResponseStatsDevice) GetIp2Stat() IpStat`

GetIp2Stat returns the Ip2Stat field if non-nil, zero value otherwise.

### GetIp2StatOk

`func (o *ResponseStatsDevice) GetIp2StatOk() (*IpStat, bool)`

GetIp2StatOk returns a tuple with the Ip2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2Stat

`func (o *ResponseStatsDevice) SetIp2Stat(v IpStat)`

SetIp2Stat sets Ip2Stat field to given value.

### HasIp2Stat

`func (o *ResponseStatsDevice) HasIp2Stat() bool`

HasIp2Stat returns a boolean if a field has been set.

### GetMemory2Stat

`func (o *ResponseStatsDevice) GetMemory2Stat() MemoryStat`

GetMemory2Stat returns the Memory2Stat field if non-nil, zero value otherwise.

### GetMemory2StatOk

`func (o *ResponseStatsDevice) GetMemory2StatOk() (*MemoryStat, bool)`

GetMemory2StatOk returns a tuple with the Memory2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory2Stat

`func (o *ResponseStatsDevice) SetMemory2Stat(v MemoryStat)`

SetMemory2Stat sets Memory2Stat field to given value.

### HasMemory2Stat

`func (o *ResponseStatsDevice) HasMemory2Stat() bool`

HasMemory2Stat returns a boolean if a field has been set.

### GetMemoryStat

`func (o *ResponseStatsDevice) GetMemoryStat() MemoryStat`

GetMemoryStat returns the MemoryStat field if non-nil, zero value otherwise.

### GetMemoryStatOk

`func (o *ResponseStatsDevice) GetMemoryStatOk() (*MemoryStat, bool)`

GetMemoryStatOk returns a tuple with the MemoryStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryStat

`func (o *ResponseStatsDevice) SetMemoryStat(v MemoryStat)`

SetMemoryStat sets MemoryStat field to given value.

### HasMemoryStat

`func (o *ResponseStatsDevice) HasMemoryStat() bool`

HasMemoryStat returns a boolean if a field has been set.

### GetModule2Stat

`func (o *ResponseStatsDevice) GetModule2Stat() []ModuleStat`

GetModule2Stat returns the Module2Stat field if non-nil, zero value otherwise.

### GetModule2StatOk

`func (o *ResponseStatsDevice) GetModule2StatOk() (*[]ModuleStat, bool)`

GetModule2StatOk returns a tuple with the Module2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule2Stat

`func (o *ResponseStatsDevice) SetModule2Stat(v []ModuleStat)`

SetModule2Stat sets Module2Stat field to given value.

### HasModule2Stat

`func (o *ResponseStatsDevice) HasModule2Stat() bool`

HasModule2Stat returns a boolean if a field has been set.

### GetSpu2Stat

`func (o *ResponseStatsDevice) GetSpu2Stat() GatewayStatsSpuStat`

GetSpu2Stat returns the Spu2Stat field if non-nil, zero value otherwise.

### GetSpu2StatOk

`func (o *ResponseStatsDevice) GetSpu2StatOk() (*GatewayStatsSpuStat, bool)`

GetSpu2StatOk returns a tuple with the Spu2Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpu2Stat

`func (o *ResponseStatsDevice) SetSpu2Stat(v GatewayStatsSpuStat)`

SetSpu2Stat sets Spu2Stat field to given value.

### HasSpu2Stat

`func (o *ResponseStatsDevice) HasSpu2Stat() bool`

HasSpu2Stat returns a boolean if a field has been set.

### GetSpuStat

`func (o *ResponseStatsDevice) GetSpuStat() GatewayStatsSpuStat`

GetSpuStat returns the SpuStat field if non-nil, zero value otherwise.

### GetSpuStatOk

`func (o *ResponseStatsDevice) GetSpuStatOk() (*GatewayStatsSpuStat, bool)`

GetSpuStatOk returns a tuple with the SpuStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpuStat

`func (o *ResponseStatsDevice) SetSpuStat(v GatewayStatsSpuStat)`

SetSpuStat sets SpuStat field to given value.

### HasSpuStat

`func (o *ResponseStatsDevice) HasSpuStat() bool`

HasSpuStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


