# SwitchPortStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if interface is active/inactive | [optional] 
**AuthState** | Pointer to [**SwitchPortStatsAuthState**](SwitchPortStatsAuthState.md) |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**FullDuplex** | Pointer to **bool** | indicates full or half duplex | [optional] 
**Jitter** | Pointer to **float32** | Last sampled jitter of the interface | [optional] 
**Latency** | Pointer to **float32** | Last sampled latency of the interface | [optional] 
**Loss** | Pointer to **float32** | Last sampled loss of the interface | [optional] 
**LteIccid** | Pointer to **NullableString** | LTE ICCID value, Check for null/empty | [optional] 
**LteImei** | Pointer to **NullableString** | LTE IMEI value, Check for null/empty | [optional] 
**LteImsi** | Pointer to **NullableString** | LTE IMSI value, Check for null/empty | [optional] 
**Mac** | **string** |  | [readonly] 
**MacCount** | Pointer to **int32** | Number of mac addresses in the forwarding table | [optional] 
**MacLimit** | Pointer to **int32** | Limit on number of dynamically learned macs | [optional] 
**NeighborMac** | **string** | chassis identifier of the chassis type listed | [readonly] 
**NeighborPortDesc** | Pointer to **string** | description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | [optional] [readonly] 
**NeighborSystemName** | Pointer to **string** | name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | [optional] [readonly] 
**OrgId** | **string** |  | [readonly] 
**PoeDisabled** | Pointer to **bool** | is the POE configured not be disabled. | [optional] [readonly] 
**PoeMode** | Pointer to [**SwitchPortStatsPoeMode**](SwitchPortStatsPoeMode.md) |  | [optional] 
**PoeOn** | Pointer to **bool** | is the device attached to POE | [optional] 
**PortId** | **string** |  | [readonly] 
**PortMac** | **string** | interface mac address | [readonly] 
**PortUsage** | Pointer to [**SwitchPortStatsPortUsage**](SwitchPortStatsPortUsage.md) |  | [optional] 
**PowerDraw** | Pointer to **float32** | Amount of power being used by the interface at the time the command is executed. Unit in watts. | [optional] 
**RxBcastPkts** | Pointer to **int32** | Broadcast input packets | [optional] 
**RxBps** | Pointer to **int32** | Input rate | [optional] 
**RxBytes** | **int32** | rx bytes | [readonly] 
**RxErrors** | Pointer to **int32** | Input errors | [optional] 
**RxMcastPkts** | Pointer to **int32** | Multicast input packets | [optional] 
**RxPkts** | **int32** | rx packets | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Speed** | Pointer to **int32** | port speed | [optional] [readonly] 
**StpRole** | Pointer to [**SwitchPortStatsStpRole**](SwitchPortStatsStpRole.md) |  | [optional] 
**StpState** | Pointer to [**SwitchPortStatsStpState**](SwitchPortStatsStpState.md) |  | [optional] 
**TxBcastPkts** | Pointer to **int32** | Broadcast output packets | [optional] 
**TxBps** | Pointer to **int32** | Output rate | [optional] 
**TxBytes** | **int32** | tx bytes | [readonly] 
**TxErrors** | Pointer to **int32** | Output errors | [optional] 
**TxMcastPkts** | Pointer to **int32** | Multicast output packets | [optional] 
**TxPkts** | **int32** | tx packets | [readonly] 
**Type** | Pointer to [**SwitchPortStatsType**](SwitchPortStatsType.md) |  | [optional] 
**Unconfigured** | Pointer to **bool** | indicates if interface is unconfigured | [optional] [readonly] 
**Up** | Pointer to **bool** | indicates if interface is up | [optional] [readonly] 
**XcvrModel** | Pointer to **string** | Optic Slot ModelName, Check for null/empty | [optional] 
**XcvrPartNumber** | Pointer to **string** | Optic Slot Partnumber, Check for null/empty | [optional] 
**XcvrSerial** | Pointer to **string** | Optic Slot SerialNumber, Check for null/empty | [optional] 

## Methods

### NewSwitchPortStats

`func NewSwitchPortStats(mac string, neighborMac string, orgId string, portId string, portMac string, rxBytes int32, rxPkts int32, siteId string, txBytes int32, txPkts int32, ) *SwitchPortStats`

NewSwitchPortStats instantiates a new SwitchPortStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortStatsWithDefaults

`func NewSwitchPortStatsWithDefaults() *SwitchPortStats`

NewSwitchPortStatsWithDefaults instantiates a new SwitchPortStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SwitchPortStats) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SwitchPortStats) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SwitchPortStats) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SwitchPortStats) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAuthState

`func (o *SwitchPortStats) GetAuthState() SwitchPortStatsAuthState`

GetAuthState returns the AuthState field if non-nil, zero value otherwise.

### GetAuthStateOk

`func (o *SwitchPortStats) GetAuthStateOk() (*SwitchPortStatsAuthState, bool)`

GetAuthStateOk returns a tuple with the AuthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthState

`func (o *SwitchPortStats) SetAuthState(v SwitchPortStatsAuthState)`

SetAuthState sets AuthState field to given value.

### HasAuthState

`func (o *SwitchPortStats) HasAuthState() bool`

HasAuthState returns a boolean if a field has been set.

### GetForSite

`func (o *SwitchPortStats) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *SwitchPortStats) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *SwitchPortStats) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *SwitchPortStats) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetFullDuplex

`func (o *SwitchPortStats) GetFullDuplex() bool`

GetFullDuplex returns the FullDuplex field if non-nil, zero value otherwise.

### GetFullDuplexOk

`func (o *SwitchPortStats) GetFullDuplexOk() (*bool, bool)`

GetFullDuplexOk returns a tuple with the FullDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDuplex

`func (o *SwitchPortStats) SetFullDuplex(v bool)`

SetFullDuplex sets FullDuplex field to given value.

### HasFullDuplex

`func (o *SwitchPortStats) HasFullDuplex() bool`

HasFullDuplex returns a boolean if a field has been set.

### GetJitter

`func (o *SwitchPortStats) GetJitter() float32`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *SwitchPortStats) GetJitterOk() (*float32, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *SwitchPortStats) SetJitter(v float32)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *SwitchPortStats) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetLatency

`func (o *SwitchPortStats) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *SwitchPortStats) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *SwitchPortStats) SetLatency(v float32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *SwitchPortStats) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLoss

`func (o *SwitchPortStats) GetLoss() float32`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *SwitchPortStats) GetLossOk() (*float32, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *SwitchPortStats) SetLoss(v float32)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *SwitchPortStats) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLteIccid

`func (o *SwitchPortStats) GetLteIccid() string`

GetLteIccid returns the LteIccid field if non-nil, zero value otherwise.

### GetLteIccidOk

`func (o *SwitchPortStats) GetLteIccidOk() (*string, bool)`

GetLteIccidOk returns a tuple with the LteIccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteIccid

`func (o *SwitchPortStats) SetLteIccid(v string)`

SetLteIccid sets LteIccid field to given value.

### HasLteIccid

`func (o *SwitchPortStats) HasLteIccid() bool`

HasLteIccid returns a boolean if a field has been set.

### SetLteIccidNil

`func (o *SwitchPortStats) SetLteIccidNil(b bool)`

 SetLteIccidNil sets the value for LteIccid to be an explicit nil

### UnsetLteIccid
`func (o *SwitchPortStats) UnsetLteIccid()`

UnsetLteIccid ensures that no value is present for LteIccid, not even an explicit nil
### GetLteImei

`func (o *SwitchPortStats) GetLteImei() string`

GetLteImei returns the LteImei field if non-nil, zero value otherwise.

### GetLteImeiOk

`func (o *SwitchPortStats) GetLteImeiOk() (*string, bool)`

GetLteImeiOk returns a tuple with the LteImei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteImei

`func (o *SwitchPortStats) SetLteImei(v string)`

SetLteImei sets LteImei field to given value.

### HasLteImei

`func (o *SwitchPortStats) HasLteImei() bool`

HasLteImei returns a boolean if a field has been set.

### SetLteImeiNil

`func (o *SwitchPortStats) SetLteImeiNil(b bool)`

 SetLteImeiNil sets the value for LteImei to be an explicit nil

### UnsetLteImei
`func (o *SwitchPortStats) UnsetLteImei()`

UnsetLteImei ensures that no value is present for LteImei, not even an explicit nil
### GetLteImsi

`func (o *SwitchPortStats) GetLteImsi() string`

GetLteImsi returns the LteImsi field if non-nil, zero value otherwise.

### GetLteImsiOk

`func (o *SwitchPortStats) GetLteImsiOk() (*string, bool)`

GetLteImsiOk returns a tuple with the LteImsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteImsi

`func (o *SwitchPortStats) SetLteImsi(v string)`

SetLteImsi sets LteImsi field to given value.

### HasLteImsi

`func (o *SwitchPortStats) HasLteImsi() bool`

HasLteImsi returns a boolean if a field has been set.

### SetLteImsiNil

`func (o *SwitchPortStats) SetLteImsiNil(b bool)`

 SetLteImsiNil sets the value for LteImsi to be an explicit nil

### UnsetLteImsi
`func (o *SwitchPortStats) UnsetLteImsi()`

UnsetLteImsi ensures that no value is present for LteImsi, not even an explicit nil
### GetMac

`func (o *SwitchPortStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchPortStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchPortStats) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMacCount

`func (o *SwitchPortStats) GetMacCount() int32`

GetMacCount returns the MacCount field if non-nil, zero value otherwise.

### GetMacCountOk

`func (o *SwitchPortStats) GetMacCountOk() (*int32, bool)`

GetMacCountOk returns a tuple with the MacCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacCount

`func (o *SwitchPortStats) SetMacCount(v int32)`

SetMacCount sets MacCount field to given value.

### HasMacCount

`func (o *SwitchPortStats) HasMacCount() bool`

HasMacCount returns a boolean if a field has been set.

### GetMacLimit

`func (o *SwitchPortStats) GetMacLimit() int32`

GetMacLimit returns the MacLimit field if non-nil, zero value otherwise.

### GetMacLimitOk

`func (o *SwitchPortStats) GetMacLimitOk() (*int32, bool)`

GetMacLimitOk returns a tuple with the MacLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLimit

`func (o *SwitchPortStats) SetMacLimit(v int32)`

SetMacLimit sets MacLimit field to given value.

### HasMacLimit

`func (o *SwitchPortStats) HasMacLimit() bool`

HasMacLimit returns a boolean if a field has been set.

### GetNeighborMac

`func (o *SwitchPortStats) GetNeighborMac() string`

GetNeighborMac returns the NeighborMac field if non-nil, zero value otherwise.

### GetNeighborMacOk

`func (o *SwitchPortStats) GetNeighborMacOk() (*string, bool)`

GetNeighborMacOk returns a tuple with the NeighborMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborMac

`func (o *SwitchPortStats) SetNeighborMac(v string)`

SetNeighborMac sets NeighborMac field to given value.


### GetNeighborPortDesc

`func (o *SwitchPortStats) GetNeighborPortDesc() string`

GetNeighborPortDesc returns the NeighborPortDesc field if non-nil, zero value otherwise.

### GetNeighborPortDescOk

`func (o *SwitchPortStats) GetNeighborPortDescOk() (*string, bool)`

GetNeighborPortDescOk returns a tuple with the NeighborPortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborPortDesc

`func (o *SwitchPortStats) SetNeighborPortDesc(v string)`

SetNeighborPortDesc sets NeighborPortDesc field to given value.

### HasNeighborPortDesc

`func (o *SwitchPortStats) HasNeighborPortDesc() bool`

HasNeighborPortDesc returns a boolean if a field has been set.

### GetNeighborSystemName

`func (o *SwitchPortStats) GetNeighborSystemName() string`

GetNeighborSystemName returns the NeighborSystemName field if non-nil, zero value otherwise.

### GetNeighborSystemNameOk

`func (o *SwitchPortStats) GetNeighborSystemNameOk() (*string, bool)`

GetNeighborSystemNameOk returns a tuple with the NeighborSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborSystemName

`func (o *SwitchPortStats) SetNeighborSystemName(v string)`

SetNeighborSystemName sets NeighborSystemName field to given value.

### HasNeighborSystemName

`func (o *SwitchPortStats) HasNeighborSystemName() bool`

HasNeighborSystemName returns a boolean if a field has been set.

### GetOrgId

`func (o *SwitchPortStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SwitchPortStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SwitchPortStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetPoeDisabled

`func (o *SwitchPortStats) GetPoeDisabled() bool`

GetPoeDisabled returns the PoeDisabled field if non-nil, zero value otherwise.

### GetPoeDisabledOk

`func (o *SwitchPortStats) GetPoeDisabledOk() (*bool, bool)`

GetPoeDisabledOk returns a tuple with the PoeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisabled

`func (o *SwitchPortStats) SetPoeDisabled(v bool)`

SetPoeDisabled sets PoeDisabled field to given value.

### HasPoeDisabled

`func (o *SwitchPortStats) HasPoeDisabled() bool`

HasPoeDisabled returns a boolean if a field has been set.

### GetPoeMode

`func (o *SwitchPortStats) GetPoeMode() SwitchPortStatsPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *SwitchPortStats) GetPoeModeOk() (*SwitchPortStatsPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *SwitchPortStats) SetPoeMode(v SwitchPortStatsPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *SwitchPortStats) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### GetPoeOn

`func (o *SwitchPortStats) GetPoeOn() bool`

GetPoeOn returns the PoeOn field if non-nil, zero value otherwise.

### GetPoeOnOk

`func (o *SwitchPortStats) GetPoeOnOk() (*bool, bool)`

GetPoeOnOk returns a tuple with the PoeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeOn

`func (o *SwitchPortStats) SetPoeOn(v bool)`

SetPoeOn sets PoeOn field to given value.

### HasPoeOn

`func (o *SwitchPortStats) HasPoeOn() bool`

HasPoeOn returns a boolean if a field has been set.

### GetPortId

`func (o *SwitchPortStats) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SwitchPortStats) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SwitchPortStats) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetPortMac

`func (o *SwitchPortStats) GetPortMac() string`

GetPortMac returns the PortMac field if non-nil, zero value otherwise.

### GetPortMacOk

`func (o *SwitchPortStats) GetPortMacOk() (*string, bool)`

GetPortMacOk returns a tuple with the PortMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMac

`func (o *SwitchPortStats) SetPortMac(v string)`

SetPortMac sets PortMac field to given value.


### GetPortUsage

`func (o *SwitchPortStats) GetPortUsage() SwitchPortStatsPortUsage`

GetPortUsage returns the PortUsage field if non-nil, zero value otherwise.

### GetPortUsageOk

`func (o *SwitchPortStats) GetPortUsageOk() (*SwitchPortStatsPortUsage, bool)`

GetPortUsageOk returns a tuple with the PortUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsage

`func (o *SwitchPortStats) SetPortUsage(v SwitchPortStatsPortUsage)`

SetPortUsage sets PortUsage field to given value.

### HasPortUsage

`func (o *SwitchPortStats) HasPortUsage() bool`

HasPortUsage returns a boolean if a field has been set.

### GetPowerDraw

`func (o *SwitchPortStats) GetPowerDraw() float32`

GetPowerDraw returns the PowerDraw field if non-nil, zero value otherwise.

### GetPowerDrawOk

`func (o *SwitchPortStats) GetPowerDrawOk() (*float32, bool)`

GetPowerDrawOk returns a tuple with the PowerDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDraw

`func (o *SwitchPortStats) SetPowerDraw(v float32)`

SetPowerDraw sets PowerDraw field to given value.

### HasPowerDraw

`func (o *SwitchPortStats) HasPowerDraw() bool`

HasPowerDraw returns a boolean if a field has been set.

### GetRxBcastPkts

`func (o *SwitchPortStats) GetRxBcastPkts() int32`

GetRxBcastPkts returns the RxBcastPkts field if non-nil, zero value otherwise.

### GetRxBcastPktsOk

`func (o *SwitchPortStats) GetRxBcastPktsOk() (*int32, bool)`

GetRxBcastPktsOk returns a tuple with the RxBcastPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBcastPkts

`func (o *SwitchPortStats) SetRxBcastPkts(v int32)`

SetRxBcastPkts sets RxBcastPkts field to given value.

### HasRxBcastPkts

`func (o *SwitchPortStats) HasRxBcastPkts() bool`

HasRxBcastPkts returns a boolean if a field has been set.

### GetRxBps

`func (o *SwitchPortStats) GetRxBps() int32`

GetRxBps returns the RxBps field if non-nil, zero value otherwise.

### GetRxBpsOk

`func (o *SwitchPortStats) GetRxBpsOk() (*int32, bool)`

GetRxBpsOk returns a tuple with the RxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBps

`func (o *SwitchPortStats) SetRxBps(v int32)`

SetRxBps sets RxBps field to given value.

### HasRxBps

`func (o *SwitchPortStats) HasRxBps() bool`

HasRxBps returns a boolean if a field has been set.

### GetRxBytes

`func (o *SwitchPortStats) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *SwitchPortStats) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *SwitchPortStats) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.


### GetRxErrors

`func (o *SwitchPortStats) GetRxErrors() int32`

GetRxErrors returns the RxErrors field if non-nil, zero value otherwise.

### GetRxErrorsOk

`func (o *SwitchPortStats) GetRxErrorsOk() (*int32, bool)`

GetRxErrorsOk returns a tuple with the RxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrors

`func (o *SwitchPortStats) SetRxErrors(v int32)`

SetRxErrors sets RxErrors field to given value.

### HasRxErrors

`func (o *SwitchPortStats) HasRxErrors() bool`

HasRxErrors returns a boolean if a field has been set.

### GetRxMcastPkts

`func (o *SwitchPortStats) GetRxMcastPkts() int32`

GetRxMcastPkts returns the RxMcastPkts field if non-nil, zero value otherwise.

### GetRxMcastPktsOk

`func (o *SwitchPortStats) GetRxMcastPktsOk() (*int32, bool)`

GetRxMcastPktsOk returns a tuple with the RxMcastPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMcastPkts

`func (o *SwitchPortStats) SetRxMcastPkts(v int32)`

SetRxMcastPkts sets RxMcastPkts field to given value.

### HasRxMcastPkts

`func (o *SwitchPortStats) HasRxMcastPkts() bool`

HasRxMcastPkts returns a boolean if a field has been set.

### GetRxPkts

`func (o *SwitchPortStats) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *SwitchPortStats) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *SwitchPortStats) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.


### GetSiteId

`func (o *SwitchPortStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SwitchPortStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SwitchPortStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSpeed

`func (o *SwitchPortStats) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *SwitchPortStats) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *SwitchPortStats) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *SwitchPortStats) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStpRole

`func (o *SwitchPortStats) GetStpRole() SwitchPortStatsStpRole`

GetStpRole returns the StpRole field if non-nil, zero value otherwise.

### GetStpRoleOk

`func (o *SwitchPortStats) GetStpRoleOk() (*SwitchPortStatsStpRole, bool)`

GetStpRoleOk returns a tuple with the StpRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpRole

`func (o *SwitchPortStats) SetStpRole(v SwitchPortStatsStpRole)`

SetStpRole sets StpRole field to given value.

### HasStpRole

`func (o *SwitchPortStats) HasStpRole() bool`

HasStpRole returns a boolean if a field has been set.

### GetStpState

`func (o *SwitchPortStats) GetStpState() SwitchPortStatsStpState`

GetStpState returns the StpState field if non-nil, zero value otherwise.

### GetStpStateOk

`func (o *SwitchPortStats) GetStpStateOk() (*SwitchPortStatsStpState, bool)`

GetStpStateOk returns a tuple with the StpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpState

`func (o *SwitchPortStats) SetStpState(v SwitchPortStatsStpState)`

SetStpState sets StpState field to given value.

### HasStpState

`func (o *SwitchPortStats) HasStpState() bool`

HasStpState returns a boolean if a field has been set.

### GetTxBcastPkts

`func (o *SwitchPortStats) GetTxBcastPkts() int32`

GetTxBcastPkts returns the TxBcastPkts field if non-nil, zero value otherwise.

### GetTxBcastPktsOk

`func (o *SwitchPortStats) GetTxBcastPktsOk() (*int32, bool)`

GetTxBcastPktsOk returns a tuple with the TxBcastPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBcastPkts

`func (o *SwitchPortStats) SetTxBcastPkts(v int32)`

SetTxBcastPkts sets TxBcastPkts field to given value.

### HasTxBcastPkts

`func (o *SwitchPortStats) HasTxBcastPkts() bool`

HasTxBcastPkts returns a boolean if a field has been set.

### GetTxBps

`func (o *SwitchPortStats) GetTxBps() int32`

GetTxBps returns the TxBps field if non-nil, zero value otherwise.

### GetTxBpsOk

`func (o *SwitchPortStats) GetTxBpsOk() (*int32, bool)`

GetTxBpsOk returns a tuple with the TxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBps

`func (o *SwitchPortStats) SetTxBps(v int32)`

SetTxBps sets TxBps field to given value.

### HasTxBps

`func (o *SwitchPortStats) HasTxBps() bool`

HasTxBps returns a boolean if a field has been set.

### GetTxBytes

`func (o *SwitchPortStats) GetTxBytes() int32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *SwitchPortStats) GetTxBytesOk() (*int32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *SwitchPortStats) SetTxBytes(v int32)`

SetTxBytes sets TxBytes field to given value.


### GetTxErrors

`func (o *SwitchPortStats) GetTxErrors() int32`

GetTxErrors returns the TxErrors field if non-nil, zero value otherwise.

### GetTxErrorsOk

`func (o *SwitchPortStats) GetTxErrorsOk() (*int32, bool)`

GetTxErrorsOk returns a tuple with the TxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrors

`func (o *SwitchPortStats) SetTxErrors(v int32)`

SetTxErrors sets TxErrors field to given value.

### HasTxErrors

`func (o *SwitchPortStats) HasTxErrors() bool`

HasTxErrors returns a boolean if a field has been set.

### GetTxMcastPkts

`func (o *SwitchPortStats) GetTxMcastPkts() int32`

GetTxMcastPkts returns the TxMcastPkts field if non-nil, zero value otherwise.

### GetTxMcastPktsOk

`func (o *SwitchPortStats) GetTxMcastPktsOk() (*int32, bool)`

GetTxMcastPktsOk returns a tuple with the TxMcastPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMcastPkts

`func (o *SwitchPortStats) SetTxMcastPkts(v int32)`

SetTxMcastPkts sets TxMcastPkts field to given value.

### HasTxMcastPkts

`func (o *SwitchPortStats) HasTxMcastPkts() bool`

HasTxMcastPkts returns a boolean if a field has been set.

### GetTxPkts

`func (o *SwitchPortStats) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *SwitchPortStats) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *SwitchPortStats) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.


### GetType

`func (o *SwitchPortStats) GetType() SwitchPortStatsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchPortStats) GetTypeOk() (*SwitchPortStatsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchPortStats) SetType(v SwitchPortStatsType)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchPortStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnconfigured

`func (o *SwitchPortStats) GetUnconfigured() bool`

GetUnconfigured returns the Unconfigured field if non-nil, zero value otherwise.

### GetUnconfiguredOk

`func (o *SwitchPortStats) GetUnconfiguredOk() (*bool, bool)`

GetUnconfiguredOk returns a tuple with the Unconfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfigured

`func (o *SwitchPortStats) SetUnconfigured(v bool)`

SetUnconfigured sets Unconfigured field to given value.

### HasUnconfigured

`func (o *SwitchPortStats) HasUnconfigured() bool`

HasUnconfigured returns a boolean if a field has been set.

### GetUp

`func (o *SwitchPortStats) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *SwitchPortStats) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *SwitchPortStats) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *SwitchPortStats) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetXcvrModel

`func (o *SwitchPortStats) GetXcvrModel() string`

GetXcvrModel returns the XcvrModel field if non-nil, zero value otherwise.

### GetXcvrModelOk

`func (o *SwitchPortStats) GetXcvrModelOk() (*string, bool)`

GetXcvrModelOk returns a tuple with the XcvrModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcvrModel

`func (o *SwitchPortStats) SetXcvrModel(v string)`

SetXcvrModel sets XcvrModel field to given value.

### HasXcvrModel

`func (o *SwitchPortStats) HasXcvrModel() bool`

HasXcvrModel returns a boolean if a field has been set.

### GetXcvrPartNumber

`func (o *SwitchPortStats) GetXcvrPartNumber() string`

GetXcvrPartNumber returns the XcvrPartNumber field if non-nil, zero value otherwise.

### GetXcvrPartNumberOk

`func (o *SwitchPortStats) GetXcvrPartNumberOk() (*string, bool)`

GetXcvrPartNumberOk returns a tuple with the XcvrPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcvrPartNumber

`func (o *SwitchPortStats) SetXcvrPartNumber(v string)`

SetXcvrPartNumber sets XcvrPartNumber field to given value.

### HasXcvrPartNumber

`func (o *SwitchPortStats) HasXcvrPartNumber() bool`

HasXcvrPartNumber returns a boolean if a field has been set.

### GetXcvrSerial

`func (o *SwitchPortStats) GetXcvrSerial() string`

GetXcvrSerial returns the XcvrSerial field if non-nil, zero value otherwise.

### GetXcvrSerialOk

`func (o *SwitchPortStats) GetXcvrSerialOk() (*string, bool)`

GetXcvrSerialOk returns a tuple with the XcvrSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcvrSerial

`func (o *SwitchPortStats) SetXcvrSerial(v string)`

SetXcvrSerial sets XcvrSerial field to given value.

### HasXcvrSerial

`func (o *SwitchPortStats) HasXcvrSerial() bool`

HasXcvrSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


