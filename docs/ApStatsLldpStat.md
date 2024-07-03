# ApStatsLldpStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | Pointer to **string** |  | [optional] 
**LldpMedSupported** | Pointer to **bool** | whether it support LLDP-MED | [optional] 
**MgmtAddr** | Pointer to **string** | switch’s management address (if advertised), can be IPv4, IPv6, or MAC | [optional] 
**PortDesc** | Pointer to **string** | port description, e.g. “2/20”, “Port 2 on Switch0” | [optional] 
**PowerAllocated** | Pointer to **float32** | in mW, provided/allocated by PSE | [optional] 
**PowerDraw** | Pointer to **float32** | in mW, total power needed by PD | [optional] 
**PowerRequestCount** | Pointer to **int32** | number of negotiations, if it keeps increasing, we don’t have a stable power | [optional] 
**PowerRequested** | Pointer to **float32** | in mW, the current power requested by PD | [optional] 
**SystemDesc** | Pointer to **string** | description provided by switch, e.g. “HP J9729A 2920-48G-POE+ Switch” | [optional] 
**SystemName** | Pointer to **string** | name of the switch, e.g. “TC2-OWL-Stack-01” | [optional] 

## Methods

### NewApStatsLldpStat

`func NewApStatsLldpStat() *ApStatsLldpStat`

NewApStatsLldpStat instantiates a new ApStatsLldpStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsLldpStatWithDefaults

`func NewApStatsLldpStatWithDefaults() *ApStatsLldpStat`

NewApStatsLldpStatWithDefaults instantiates a new ApStatsLldpStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisId

`func (o *ApStatsLldpStat) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ApStatsLldpStat) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ApStatsLldpStat) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ApStatsLldpStat) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetLldpMedSupported

`func (o *ApStatsLldpStat) GetLldpMedSupported() bool`

GetLldpMedSupported returns the LldpMedSupported field if non-nil, zero value otherwise.

### GetLldpMedSupportedOk

`func (o *ApStatsLldpStat) GetLldpMedSupportedOk() (*bool, bool)`

GetLldpMedSupportedOk returns a tuple with the LldpMedSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedSupported

`func (o *ApStatsLldpStat) SetLldpMedSupported(v bool)`

SetLldpMedSupported sets LldpMedSupported field to given value.

### HasLldpMedSupported

`func (o *ApStatsLldpStat) HasLldpMedSupported() bool`

HasLldpMedSupported returns a boolean if a field has been set.

### GetMgmtAddr

`func (o *ApStatsLldpStat) GetMgmtAddr() string`

GetMgmtAddr returns the MgmtAddr field if non-nil, zero value otherwise.

### GetMgmtAddrOk

`func (o *ApStatsLldpStat) GetMgmtAddrOk() (*string, bool)`

GetMgmtAddrOk returns a tuple with the MgmtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtAddr

`func (o *ApStatsLldpStat) SetMgmtAddr(v string)`

SetMgmtAddr sets MgmtAddr field to given value.

### HasMgmtAddr

`func (o *ApStatsLldpStat) HasMgmtAddr() bool`

HasMgmtAddr returns a boolean if a field has been set.

### GetPortDesc

`func (o *ApStatsLldpStat) GetPortDesc() string`

GetPortDesc returns the PortDesc field if non-nil, zero value otherwise.

### GetPortDescOk

`func (o *ApStatsLldpStat) GetPortDescOk() (*string, bool)`

GetPortDescOk returns a tuple with the PortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDesc

`func (o *ApStatsLldpStat) SetPortDesc(v string)`

SetPortDesc sets PortDesc field to given value.

### HasPortDesc

`func (o *ApStatsLldpStat) HasPortDesc() bool`

HasPortDesc returns a boolean if a field has been set.

### GetPowerAllocated

`func (o *ApStatsLldpStat) GetPowerAllocated() float32`

GetPowerAllocated returns the PowerAllocated field if non-nil, zero value otherwise.

### GetPowerAllocatedOk

`func (o *ApStatsLldpStat) GetPowerAllocatedOk() (*float32, bool)`

GetPowerAllocatedOk returns a tuple with the PowerAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAllocated

`func (o *ApStatsLldpStat) SetPowerAllocated(v float32)`

SetPowerAllocated sets PowerAllocated field to given value.

### HasPowerAllocated

`func (o *ApStatsLldpStat) HasPowerAllocated() bool`

HasPowerAllocated returns a boolean if a field has been set.

### GetPowerDraw

`func (o *ApStatsLldpStat) GetPowerDraw() float32`

GetPowerDraw returns the PowerDraw field if non-nil, zero value otherwise.

### GetPowerDrawOk

`func (o *ApStatsLldpStat) GetPowerDrawOk() (*float32, bool)`

GetPowerDrawOk returns a tuple with the PowerDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDraw

`func (o *ApStatsLldpStat) SetPowerDraw(v float32)`

SetPowerDraw sets PowerDraw field to given value.

### HasPowerDraw

`func (o *ApStatsLldpStat) HasPowerDraw() bool`

HasPowerDraw returns a boolean if a field has been set.

### GetPowerRequestCount

`func (o *ApStatsLldpStat) GetPowerRequestCount() int32`

GetPowerRequestCount returns the PowerRequestCount field if non-nil, zero value otherwise.

### GetPowerRequestCountOk

`func (o *ApStatsLldpStat) GetPowerRequestCountOk() (*int32, bool)`

GetPowerRequestCountOk returns a tuple with the PowerRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRequestCount

`func (o *ApStatsLldpStat) SetPowerRequestCount(v int32)`

SetPowerRequestCount sets PowerRequestCount field to given value.

### HasPowerRequestCount

`func (o *ApStatsLldpStat) HasPowerRequestCount() bool`

HasPowerRequestCount returns a boolean if a field has been set.

### GetPowerRequested

`func (o *ApStatsLldpStat) GetPowerRequested() float32`

GetPowerRequested returns the PowerRequested field if non-nil, zero value otherwise.

### GetPowerRequestedOk

`func (o *ApStatsLldpStat) GetPowerRequestedOk() (*float32, bool)`

GetPowerRequestedOk returns a tuple with the PowerRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRequested

`func (o *ApStatsLldpStat) SetPowerRequested(v float32)`

SetPowerRequested sets PowerRequested field to given value.

### HasPowerRequested

`func (o *ApStatsLldpStat) HasPowerRequested() bool`

HasPowerRequested returns a boolean if a field has been set.

### GetSystemDesc

`func (o *ApStatsLldpStat) GetSystemDesc() string`

GetSystemDesc returns the SystemDesc field if non-nil, zero value otherwise.

### GetSystemDescOk

`func (o *ApStatsLldpStat) GetSystemDescOk() (*string, bool)`

GetSystemDescOk returns a tuple with the SystemDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDesc

`func (o *ApStatsLldpStat) SetSystemDesc(v string)`

SetSystemDesc sets SystemDesc field to given value.

### HasSystemDesc

`func (o *ApStatsLldpStat) HasSystemDesc() bool`

HasSystemDesc returns a boolean if a field has been set.

### GetSystemName

`func (o *ApStatsLldpStat) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *ApStatsLldpStat) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *ApStatsLldpStat) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *ApStatsLldpStat) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


