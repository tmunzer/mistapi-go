# JunosPortConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AeDisableLacp** | Pointer to **bool** | To disable LACP support for the AE interface | [optional] 
**AeIdx** | Pointer to **int32** | Users could force to use the designated AE name | [optional] 
**AeLacpSlow** | Pointer to **bool** | to use fast timeout | [optional] [default to true]
**Aggregated** | Pointer to **bool** |  | [optional] [default to false]
**Critical** | Pointer to **bool** | if want to generate port up/down alarm | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisableAutoneg** | Pointer to **bool** | if &#x60;speed&#x60; and &#x60;duplex&#x60; are specified, whether to disable autonegotiation | [optional] [default to false]
**Duplex** | Pointer to [**JunosPortConfigDuplex**](JunosPortConfigDuplex.md) |  | [optional] [default to JUNOSPORTCONFIGDUPLEX_AUTO]
**DynamicUsage** | Pointer to **NullableString** | Enable dynamic usage for this port. Set to &#x60;dynamic&#x60; to enable. | [optional] 
**Esilag** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **int32** | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation | [optional] [default to 1514]
**NoLocalOverwrite** | Pointer to **bool** | prevent helpdesk to override the port config | [optional] 
**PoeDisabled** | Pointer to **bool** |  | [optional] [default to false]
**Speed** | Pointer to [**JunosPortConfigSpeed**](JunosPortConfigSpeed.md) |  | [optional] [default to JUNOSPORTCONFIGSPEED_AUTO]
**Usage** | **string** | port usage name.   If EVPN is used, use &#x60;evpn_uplink&#x60;or &#x60;evpn_downlink&#x60; | 

## Methods

### NewJunosPortConfig

`func NewJunosPortConfig(usage string, ) *JunosPortConfig`

NewJunosPortConfig instantiates a new JunosPortConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosPortConfigWithDefaults

`func NewJunosPortConfigWithDefaults() *JunosPortConfig`

NewJunosPortConfigWithDefaults instantiates a new JunosPortConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeDisableLacp

`func (o *JunosPortConfig) GetAeDisableLacp() bool`

GetAeDisableLacp returns the AeDisableLacp field if non-nil, zero value otherwise.

### GetAeDisableLacpOk

`func (o *JunosPortConfig) GetAeDisableLacpOk() (*bool, bool)`

GetAeDisableLacpOk returns a tuple with the AeDisableLacp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeDisableLacp

`func (o *JunosPortConfig) SetAeDisableLacp(v bool)`

SetAeDisableLacp sets AeDisableLacp field to given value.

### HasAeDisableLacp

`func (o *JunosPortConfig) HasAeDisableLacp() bool`

HasAeDisableLacp returns a boolean if a field has been set.

### GetAeIdx

`func (o *JunosPortConfig) GetAeIdx() int32`

GetAeIdx returns the AeIdx field if non-nil, zero value otherwise.

### GetAeIdxOk

`func (o *JunosPortConfig) GetAeIdxOk() (*int32, bool)`

GetAeIdxOk returns a tuple with the AeIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeIdx

`func (o *JunosPortConfig) SetAeIdx(v int32)`

SetAeIdx sets AeIdx field to given value.

### HasAeIdx

`func (o *JunosPortConfig) HasAeIdx() bool`

HasAeIdx returns a boolean if a field has been set.

### GetAeLacpSlow

`func (o *JunosPortConfig) GetAeLacpSlow() bool`

GetAeLacpSlow returns the AeLacpSlow field if non-nil, zero value otherwise.

### GetAeLacpSlowOk

`func (o *JunosPortConfig) GetAeLacpSlowOk() (*bool, bool)`

GetAeLacpSlowOk returns a tuple with the AeLacpSlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeLacpSlow

`func (o *JunosPortConfig) SetAeLacpSlow(v bool)`

SetAeLacpSlow sets AeLacpSlow field to given value.

### HasAeLacpSlow

`func (o *JunosPortConfig) HasAeLacpSlow() bool`

HasAeLacpSlow returns a boolean if a field has been set.

### GetAggregated

`func (o *JunosPortConfig) GetAggregated() bool`

GetAggregated returns the Aggregated field if non-nil, zero value otherwise.

### GetAggregatedOk

`func (o *JunosPortConfig) GetAggregatedOk() (*bool, bool)`

GetAggregatedOk returns a tuple with the Aggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregated

`func (o *JunosPortConfig) SetAggregated(v bool)`

SetAggregated sets Aggregated field to given value.

### HasAggregated

`func (o *JunosPortConfig) HasAggregated() bool`

HasAggregated returns a boolean if a field has been set.

### GetCritical

`func (o *JunosPortConfig) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *JunosPortConfig) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *JunosPortConfig) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *JunosPortConfig) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetDescription

`func (o *JunosPortConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JunosPortConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JunosPortConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JunosPortConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableAutoneg

`func (o *JunosPortConfig) GetDisableAutoneg() bool`

GetDisableAutoneg returns the DisableAutoneg field if non-nil, zero value otherwise.

### GetDisableAutonegOk

`func (o *JunosPortConfig) GetDisableAutonegOk() (*bool, bool)`

GetDisableAutonegOk returns a tuple with the DisableAutoneg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoneg

`func (o *JunosPortConfig) SetDisableAutoneg(v bool)`

SetDisableAutoneg sets DisableAutoneg field to given value.

### HasDisableAutoneg

`func (o *JunosPortConfig) HasDisableAutoneg() bool`

HasDisableAutoneg returns a boolean if a field has been set.

### GetDuplex

`func (o *JunosPortConfig) GetDuplex() JunosPortConfigDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *JunosPortConfig) GetDuplexOk() (*JunosPortConfigDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *JunosPortConfig) SetDuplex(v JunosPortConfigDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *JunosPortConfig) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDynamicUsage

`func (o *JunosPortConfig) GetDynamicUsage() string`

GetDynamicUsage returns the DynamicUsage field if non-nil, zero value otherwise.

### GetDynamicUsageOk

`func (o *JunosPortConfig) GetDynamicUsageOk() (*string, bool)`

GetDynamicUsageOk returns a tuple with the DynamicUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUsage

`func (o *JunosPortConfig) SetDynamicUsage(v string)`

SetDynamicUsage sets DynamicUsage field to given value.

### HasDynamicUsage

`func (o *JunosPortConfig) HasDynamicUsage() bool`

HasDynamicUsage returns a boolean if a field has been set.

### SetDynamicUsageNil

`func (o *JunosPortConfig) SetDynamicUsageNil(b bool)`

 SetDynamicUsageNil sets the value for DynamicUsage to be an explicit nil

### UnsetDynamicUsage
`func (o *JunosPortConfig) UnsetDynamicUsage()`

UnsetDynamicUsage ensures that no value is present for DynamicUsage, not even an explicit nil
### GetEsilag

`func (o *JunosPortConfig) GetEsilag() bool`

GetEsilag returns the Esilag field if non-nil, zero value otherwise.

### GetEsilagOk

`func (o *JunosPortConfig) GetEsilagOk() (*bool, bool)`

GetEsilagOk returns a tuple with the Esilag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsilag

`func (o *JunosPortConfig) SetEsilag(v bool)`

SetEsilag sets Esilag field to given value.

### HasEsilag

`func (o *JunosPortConfig) HasEsilag() bool`

HasEsilag returns a boolean if a field has been set.

### GetMtu

`func (o *JunosPortConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *JunosPortConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *JunosPortConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *JunosPortConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNoLocalOverwrite

`func (o *JunosPortConfig) GetNoLocalOverwrite() bool`

GetNoLocalOverwrite returns the NoLocalOverwrite field if non-nil, zero value otherwise.

### GetNoLocalOverwriteOk

`func (o *JunosPortConfig) GetNoLocalOverwriteOk() (*bool, bool)`

GetNoLocalOverwriteOk returns a tuple with the NoLocalOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalOverwrite

`func (o *JunosPortConfig) SetNoLocalOverwrite(v bool)`

SetNoLocalOverwrite sets NoLocalOverwrite field to given value.

### HasNoLocalOverwrite

`func (o *JunosPortConfig) HasNoLocalOverwrite() bool`

HasNoLocalOverwrite returns a boolean if a field has been set.

### GetPoeDisabled

`func (o *JunosPortConfig) GetPoeDisabled() bool`

GetPoeDisabled returns the PoeDisabled field if non-nil, zero value otherwise.

### GetPoeDisabledOk

`func (o *JunosPortConfig) GetPoeDisabledOk() (*bool, bool)`

GetPoeDisabledOk returns a tuple with the PoeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisabled

`func (o *JunosPortConfig) SetPoeDisabled(v bool)`

SetPoeDisabled sets PoeDisabled field to given value.

### HasPoeDisabled

`func (o *JunosPortConfig) HasPoeDisabled() bool`

HasPoeDisabled returns a boolean if a field has been set.

### GetSpeed

`func (o *JunosPortConfig) GetSpeed() JunosPortConfigSpeed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *JunosPortConfig) GetSpeedOk() (*JunosPortConfigSpeed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *JunosPortConfig) SetSpeed(v JunosPortConfigSpeed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *JunosPortConfig) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetUsage

`func (o *JunosPortConfig) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *JunosPortConfig) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *JunosPortConfig) SetUsage(v string)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


