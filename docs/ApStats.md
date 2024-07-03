# ApStats

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
**Ip** | Pointer to **string** |  | [optional] 
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
**Name** | Pointer to **string** |  | [optional] 
**NumClients** | Pointer to **int32** | how many wireless clients are currently connected | [optional] 
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
**Status** | Pointer to [**ApStatsStatus**](ApStatsStatus.md) |  | [optional] 
**SwitchRedundancy** | Pointer to [**ApStatsSwitchRedundancy**](ApStatsSwitchRedundancy.md) |  | [optional] 
**TxBps** | Pointer to **float32** |  | [optional] 
**TxBytes** | Pointer to **float32** |  | [optional] 
**TxPkts** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** | device type, ap / ble | [optional] 
**Uptime** | Pointer to **float32** | how long, in seconds, has the device been up (or rebooted) | [optional] 
**UsbStat** | Pointer to [**ApStatsUsbStat**](ApStatsUsbStat.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float32** |  | [optional] 
**Y** | Pointer to **float32** |  | [optional] 

## Methods

### NewApStats

`func NewApStats(mac string, model string, ) *ApStats`

NewApStats instantiates a new ApStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsWithDefaults

`func NewApStatsWithDefaults() *ApStats`

NewApStatsWithDefaults instantiates a new ApStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPlacement

`func (o *ApStats) GetAutoPlacement() ApStatsAutoPlacement`

GetAutoPlacement returns the AutoPlacement field if non-nil, zero value otherwise.

### GetAutoPlacementOk

`func (o *ApStats) GetAutoPlacementOk() (*ApStatsAutoPlacement, bool)`

GetAutoPlacementOk returns a tuple with the AutoPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlacement

`func (o *ApStats) SetAutoPlacement(v ApStatsAutoPlacement)`

SetAutoPlacement sets AutoPlacement field to given value.

### HasAutoPlacement

`func (o *ApStats) HasAutoPlacement() bool`

HasAutoPlacement returns a boolean if a field has been set.

### GetBleConfig

`func (o *ApStats) GetBleConfig() ApStatsBleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *ApStats) GetBleConfigOk() (*ApStatsBleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *ApStats) SetBleConfig(v ApStatsBleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *ApStats) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetBleStat

`func (o *ApStats) GetBleStat() ApStatsBleStat`

GetBleStat returns the BleStat field if non-nil, zero value otherwise.

### GetBleStatOk

`func (o *ApStats) GetBleStatOk() (*ApStatsBleStat, bool)`

GetBleStatOk returns a tuple with the BleStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleStat

`func (o *ApStats) SetBleStat(v ApStatsBleStat)`

SetBleStat sets BleStat field to given value.

### HasBleStat

`func (o *ApStats) HasBleStat() bool`

HasBleStat returns a boolean if a field has been set.

### GetCertExpiry

`func (o *ApStats) GetCertExpiry() float32`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *ApStats) GetCertExpiryOk() (*float32, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *ApStats) SetCertExpiry(v float32)`

SetCertExpiry sets CertExpiry field to given value.

### HasCertExpiry

`func (o *ApStats) HasCertExpiry() bool`

HasCertExpiry returns a boolean if a field has been set.

### GetEnvStat

`func (o *ApStats) GetEnvStat() ApStatsEnvStat`

GetEnvStat returns the EnvStat field if non-nil, zero value otherwise.

### GetEnvStatOk

`func (o *ApStats) GetEnvStatOk() (*ApStatsEnvStat, bool)`

GetEnvStatOk returns a tuple with the EnvStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvStat

`func (o *ApStats) SetEnvStat(v ApStatsEnvStat)`

SetEnvStat sets EnvStat field to given value.

### HasEnvStat

`func (o *ApStats) HasEnvStat() bool`

HasEnvStat returns a boolean if a field has been set.

### GetEslStat

`func (o *ApStats) GetEslStat() ApStatsEslStat`

GetEslStat returns the EslStat field if non-nil, zero value otherwise.

### GetEslStatOk

`func (o *ApStats) GetEslStatOk() (*ApStatsEslStat, bool)`

GetEslStatOk returns a tuple with the EslStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslStat

`func (o *ApStats) SetEslStat(v ApStatsEslStat)`

SetEslStat sets EslStat field to given value.

### HasEslStat

`func (o *ApStats) HasEslStat() bool`

HasEslStat returns a boolean if a field has been set.

### GetExtIp

`func (o *ApStats) GetExtIp() string`

GetExtIp returns the ExtIp field if non-nil, zero value otherwise.

### GetExtIpOk

`func (o *ApStats) GetExtIpOk() (*string, bool)`

GetExtIpOk returns a tuple with the ExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIp

`func (o *ApStats) SetExtIp(v string)`

SetExtIp sets ExtIp field to given value.

### HasExtIp

`func (o *ApStats) HasExtIp() bool`

HasExtIp returns a boolean if a field has been set.

### GetFwupdate

`func (o *ApStats) GetFwupdate() ApStatsFwupdate`

GetFwupdate returns the Fwupdate field if non-nil, zero value otherwise.

### GetFwupdateOk

`func (o *ApStats) GetFwupdateOk() (*ApStatsFwupdate, bool)`

GetFwupdateOk returns a tuple with the Fwupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwupdate

`func (o *ApStats) SetFwupdate(v ApStatsFwupdate)`

SetFwupdate sets Fwupdate field to given value.

### HasFwupdate

`func (o *ApStats) HasFwupdate() bool`

HasFwupdate returns a boolean if a field has been set.

### GetIotStat

`func (o *ApStats) GetIotStat() map[string]ApStatsIotStatAdditionalProperties`

GetIotStat returns the IotStat field if non-nil, zero value otherwise.

### GetIotStatOk

`func (o *ApStats) GetIotStatOk() (*map[string]ApStatsIotStatAdditionalProperties, bool)`

GetIotStatOk returns a tuple with the IotStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotStat

`func (o *ApStats) SetIotStat(v map[string]ApStatsIotStatAdditionalProperties)`

SetIotStat sets IotStat field to given value.

### HasIotStat

`func (o *ApStats) HasIotStat() bool`

HasIotStat returns a boolean if a field has been set.

### GetIp

`func (o *ApStats) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApStats) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApStats) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApStats) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpConfig

`func (o *ApStats) GetIpConfig() ApIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ApStats) GetIpConfigOk() (*ApIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ApStats) SetIpConfig(v ApIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ApStats) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetIpStat

`func (o *ApStats) GetIpStat() IpStat`

GetIpStat returns the IpStat field if non-nil, zero value otherwise.

### GetIpStatOk

`func (o *ApStats) GetIpStatOk() (*IpStat, bool)`

GetIpStatOk returns a tuple with the IpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStat

`func (o *ApStats) SetIpStat(v IpStat)`

SetIpStat sets IpStat field to given value.

### HasIpStat

`func (o *ApStats) HasIpStat() bool`

HasIpStat returns a boolean if a field has been set.

### GetL2tpStat

`func (o *ApStats) GetL2tpStat() map[string]ApStatsL2tpStat`

GetL2tpStat returns the L2tpStat field if non-nil, zero value otherwise.

### GetL2tpStatOk

`func (o *ApStats) GetL2tpStatOk() (*map[string]ApStatsL2tpStat, bool)`

GetL2tpStatOk returns a tuple with the L2tpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2tpStat

`func (o *ApStats) SetL2tpStat(v map[string]ApStatsL2tpStat)`

SetL2tpStat sets L2tpStat field to given value.

### HasL2tpStat

`func (o *ApStats) HasL2tpStat() bool`

HasL2tpStat returns a boolean if a field has been set.

### GetLastSeen

`func (o *ApStats) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ApStats) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ApStats) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ApStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastTrouble

`func (o *ApStats) GetLastTrouble() LastTrouble`

GetLastTrouble returns the LastTrouble field if non-nil, zero value otherwise.

### GetLastTroubleOk

`func (o *ApStats) GetLastTroubleOk() (*LastTrouble, bool)`

GetLastTroubleOk returns a tuple with the LastTrouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrouble

`func (o *ApStats) SetLastTrouble(v LastTrouble)`

SetLastTrouble sets LastTrouble field to given value.

### HasLastTrouble

`func (o *ApStats) HasLastTrouble() bool`

HasLastTrouble returns a boolean if a field has been set.

### GetLed

`func (o *ApStats) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *ApStats) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *ApStats) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *ApStats) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLldpStat

`func (o *ApStats) GetLldpStat() ApStatsLldpStat`

GetLldpStat returns the LldpStat field if non-nil, zero value otherwise.

### GetLldpStatOk

`func (o *ApStats) GetLldpStatOk() (*ApStatsLldpStat, bool)`

GetLldpStatOk returns a tuple with the LldpStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpStat

`func (o *ApStats) SetLldpStat(v ApStatsLldpStat)`

SetLldpStat sets LldpStat field to given value.

### HasLldpStat

`func (o *ApStats) HasLldpStat() bool`

HasLldpStat returns a boolean if a field has been set.

### GetLocating

`func (o *ApStats) GetLocating() bool`

GetLocating returns the Locating field if non-nil, zero value otherwise.

### GetLocatingOk

`func (o *ApStats) GetLocatingOk() (*bool, bool)`

GetLocatingOk returns a tuple with the Locating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocating

`func (o *ApStats) SetLocating(v bool)`

SetLocating sets Locating field to given value.

### HasLocating

`func (o *ApStats) HasLocating() bool`

HasLocating returns a boolean if a field has been set.

### GetLocked

`func (o *ApStats) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ApStats) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ApStats) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ApStats) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMac

`func (o *ApStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApStats) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *ApStats) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ApStats) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ApStats) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *ApStats) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMeshDownlinks

`func (o *ApStats) GetMeshDownlinks() map[string]ApStatMeshDownlink`

GetMeshDownlinks returns the MeshDownlinks field if non-nil, zero value otherwise.

### GetMeshDownlinksOk

`func (o *ApStats) GetMeshDownlinksOk() (*map[string]ApStatMeshDownlink, bool)`

GetMeshDownlinksOk returns a tuple with the MeshDownlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshDownlinks

`func (o *ApStats) SetMeshDownlinks(v map[string]ApStatMeshDownlink)`

SetMeshDownlinks sets MeshDownlinks field to given value.

### HasMeshDownlinks

`func (o *ApStats) HasMeshDownlinks() bool`

HasMeshDownlinks returns a boolean if a field has been set.

### GetMeshUplink

`func (o *ApStats) GetMeshUplink() ApStatMeshUplink`

GetMeshUplink returns the MeshUplink field if non-nil, zero value otherwise.

### GetMeshUplinkOk

`func (o *ApStats) GetMeshUplinkOk() (*ApStatMeshUplink, bool)`

GetMeshUplinkOk returns a tuple with the MeshUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshUplink

`func (o *ApStats) SetMeshUplink(v ApStatMeshUplink)`

SetMeshUplink sets MeshUplink field to given value.

### HasMeshUplink

`func (o *ApStats) HasMeshUplink() bool`

HasMeshUplink returns a boolean if a field has been set.

### GetModel

`func (o *ApStats) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApStats) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApStats) SetModel(v string)`

SetModel sets Model field to given value.


### GetMount

`func (o *ApStats) GetMount() string`

GetMount returns the Mount field if non-nil, zero value otherwise.

### GetMountOk

`func (o *ApStats) GetMountOk() (*string, bool)`

GetMountOk returns a tuple with the Mount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMount

`func (o *ApStats) SetMount(v string)`

SetMount sets Mount field to given value.

### HasMount

`func (o *ApStats) HasMount() bool`

HasMount returns a boolean if a field has been set.

### GetName

`func (o *ApStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumClients

`func (o *ApStats) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *ApStats) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *ApStats) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *ApStats) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetPortStat

`func (o *ApStats) GetPortStat() map[string]ApStatsPortStat`

GetPortStat returns the PortStat field if non-nil, zero value otherwise.

### GetPortStatOk

`func (o *ApStats) GetPortStatOk() (*map[string]ApStatsPortStat, bool)`

GetPortStatOk returns a tuple with the PortStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStat

`func (o *ApStats) SetPortStat(v map[string]ApStatsPortStat)`

SetPortStat sets PortStat field to given value.

### HasPortStat

`func (o *ApStats) HasPortStat() bool`

HasPortStat returns a boolean if a field has been set.

### GetPowerBudget

`func (o *ApStats) GetPowerBudget() float32`

GetPowerBudget returns the PowerBudget field if non-nil, zero value otherwise.

### GetPowerBudgetOk

`func (o *ApStats) GetPowerBudgetOk() (*float32, bool)`

GetPowerBudgetOk returns a tuple with the PowerBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerBudget

`func (o *ApStats) SetPowerBudget(v float32)`

SetPowerBudget sets PowerBudget field to given value.

### HasPowerBudget

`func (o *ApStats) HasPowerBudget() bool`

HasPowerBudget returns a boolean if a field has been set.

### GetPowerConstrained

`func (o *ApStats) GetPowerConstrained() bool`

GetPowerConstrained returns the PowerConstrained field if non-nil, zero value otherwise.

### GetPowerConstrainedOk

`func (o *ApStats) GetPowerConstrainedOk() (*bool, bool)`

GetPowerConstrainedOk returns a tuple with the PowerConstrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConstrained

`func (o *ApStats) SetPowerConstrained(v bool)`

SetPowerConstrained sets PowerConstrained field to given value.

### HasPowerConstrained

`func (o *ApStats) HasPowerConstrained() bool`

HasPowerConstrained returns a boolean if a field has been set.

### GetPowerOpmode

`func (o *ApStats) GetPowerOpmode() string`

GetPowerOpmode returns the PowerOpmode field if non-nil, zero value otherwise.

### GetPowerOpmodeOk

`func (o *ApStats) GetPowerOpmodeOk() (*string, bool)`

GetPowerOpmodeOk returns a tuple with the PowerOpmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOpmode

`func (o *ApStats) SetPowerOpmode(v string)`

SetPowerOpmode sets PowerOpmode field to given value.

### HasPowerOpmode

`func (o *ApStats) HasPowerOpmode() bool`

HasPowerOpmode returns a boolean if a field has been set.

### GetPowerSrc

`func (o *ApStats) GetPowerSrc() string`

GetPowerSrc returns the PowerSrc field if non-nil, zero value otherwise.

### GetPowerSrcOk

`func (o *ApStats) GetPowerSrcOk() (*string, bool)`

GetPowerSrcOk returns a tuple with the PowerSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSrc

`func (o *ApStats) SetPowerSrc(v string)`

SetPowerSrc sets PowerSrc field to given value.

### HasPowerSrc

`func (o *ApStats) HasPowerSrc() bool`

HasPowerSrc returns a boolean if a field has been set.

### GetRadioConfig

`func (o *ApStats) GetRadioConfig() ApStatsRadioConfig`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *ApStats) GetRadioConfigOk() (*ApStatsRadioConfig, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *ApStats) SetRadioConfig(v ApStatsRadioConfig)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *ApStats) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetRadioStat

`func (o *ApStats) GetRadioStat() ApStatsRadioStat`

GetRadioStat returns the RadioStat field if non-nil, zero value otherwise.

### GetRadioStatOk

`func (o *ApStats) GetRadioStatOk() (*ApStatsRadioStat, bool)`

GetRadioStatOk returns a tuple with the RadioStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioStat

`func (o *ApStats) SetRadioStat(v ApStatsRadioStat)`

SetRadioStat sets RadioStat field to given value.

### HasRadioStat

`func (o *ApStats) HasRadioStat() bool`

HasRadioStat returns a boolean if a field has been set.

### GetRxBps

`func (o *ApStats) GetRxBps() float32`

GetRxBps returns the RxBps field if non-nil, zero value otherwise.

### GetRxBpsOk

`func (o *ApStats) GetRxBpsOk() (*float32, bool)`

GetRxBpsOk returns a tuple with the RxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBps

`func (o *ApStats) SetRxBps(v float32)`

SetRxBps sets RxBps field to given value.

### HasRxBps

`func (o *ApStats) HasRxBps() bool`

HasRxBps returns a boolean if a field has been set.

### GetRxBytes

`func (o *ApStats) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ApStats) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ApStats) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ApStats) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *ApStats) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ApStats) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ApStats) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ApStats) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSerial

`func (o *ApStats) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ApStats) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ApStats) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ApStats) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *ApStats) GetStatus() ApStatsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApStats) GetStatusOk() (*ApStatsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApStats) SetStatus(v ApStatsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchRedundancy

`func (o *ApStats) GetSwitchRedundancy() ApStatsSwitchRedundancy`

GetSwitchRedundancy returns the SwitchRedundancy field if non-nil, zero value otherwise.

### GetSwitchRedundancyOk

`func (o *ApStats) GetSwitchRedundancyOk() (*ApStatsSwitchRedundancy, bool)`

GetSwitchRedundancyOk returns a tuple with the SwitchRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchRedundancy

`func (o *ApStats) SetSwitchRedundancy(v ApStatsSwitchRedundancy)`

SetSwitchRedundancy sets SwitchRedundancy field to given value.

### HasSwitchRedundancy

`func (o *ApStats) HasSwitchRedundancy() bool`

HasSwitchRedundancy returns a boolean if a field has been set.

### GetTxBps

`func (o *ApStats) GetTxBps() float32`

GetTxBps returns the TxBps field if non-nil, zero value otherwise.

### GetTxBpsOk

`func (o *ApStats) GetTxBpsOk() (*float32, bool)`

GetTxBpsOk returns a tuple with the TxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBps

`func (o *ApStats) SetTxBps(v float32)`

SetTxBps sets TxBps field to given value.

### HasTxBps

`func (o *ApStats) HasTxBps() bool`

HasTxBps returns a boolean if a field has been set.

### GetTxBytes

`func (o *ApStats) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ApStats) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ApStats) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ApStats) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *ApStats) GetTxPkts() float32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ApStats) GetTxPktsOk() (*float32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ApStats) SetTxPkts(v float32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ApStats) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetType

`func (o *ApStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *ApStats) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ApStats) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ApStats) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ApStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsbStat

`func (o *ApStats) GetUsbStat() ApStatsUsbStat`

GetUsbStat returns the UsbStat field if non-nil, zero value otherwise.

### GetUsbStatOk

`func (o *ApStats) GetUsbStatOk() (*ApStatsUsbStat, bool)`

GetUsbStatOk returns a tuple with the UsbStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbStat

`func (o *ApStats) SetUsbStat(v ApStatsUsbStat)`

SetUsbStat sets UsbStat field to given value.

### HasUsbStat

`func (o *ApStats) HasUsbStat() bool`

HasUsbStat returns a boolean if a field has been set.

### GetVersion

`func (o *ApStats) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApStats) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApStats) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApStats) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetX

`func (o *ApStats) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ApStats) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ApStats) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *ApStats) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ApStats) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ApStats) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ApStats) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *ApStats) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


