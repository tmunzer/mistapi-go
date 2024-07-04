# ConstDeviceSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Defaults** | Pointer to [**ConstDeviceSwitchDefault**](ConstDeviceSwitchDefault.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**EvolvedOs** | Pointer to **bool** |  | [optional] [default to false]
**EvpnRiType** | Pointer to **string** |  | [optional] 
**Experimental** | Pointer to **bool** |  | [optional] [default to false]
**FansPluggable** | Pointer to **bool** |  | [optional] [default to false]
**HasBgp** | Pointer to **bool** |  | [optional] [default to false]
**HasEts** | Pointer to **bool** |  | [optional] [default to false]
**HasEvpn** | Pointer to **bool** |  | [optional] [default to false]
**HasIrb** | Pointer to **bool** |  | [optional] [default to false]
**HasPoeOut** | Pointer to **bool** |  | [optional] [default to false]
**HasSnapshot** | Pointer to **bool** |  | [optional] [default to true]
**HasVc** | Pointer to **bool** |  | [optional] [default to true]
**Model** | Pointer to **string** |  | [optional] 
**Modular** | Pointer to **bool** |  | [optional] [default to false]
**NoShapingRate** | Pointer to **bool** |  | [optional] [default to false]
**NumberFans** | Pointer to **int32** |  | [optional] 
**OcDevice** | Pointer to **bool** |  | [optional] [default to false]
**OobInterface** | Pointer to **string** |  | [optional] 
**PacketActionDropOnly** | Pointer to **bool** |  | [optional] [default to false]
**Pic** | Pointer to **map[string]string** | Object Key is the PIC number | [optional] 
**SubRequired** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "switch"]

## Methods

### NewConstDeviceSwitch

`func NewConstDeviceSwitch() *ConstDeviceSwitch`

NewConstDeviceSwitch instantiates a new ConstDeviceSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstDeviceSwitchWithDefaults

`func NewConstDeviceSwitchWithDefaults() *ConstDeviceSwitch`

NewConstDeviceSwitchWithDefaults instantiates a new ConstDeviceSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ConstDeviceSwitch) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ConstDeviceSwitch) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ConstDeviceSwitch) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ConstDeviceSwitch) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDefaults

`func (o *ConstDeviceSwitch) GetDefaults() ConstDeviceSwitchDefault`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *ConstDeviceSwitch) GetDefaultsOk() (*ConstDeviceSwitchDefault, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *ConstDeviceSwitch) SetDefaults(v ConstDeviceSwitchDefault)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *ConstDeviceSwitch) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDescription

`func (o *ConstDeviceSwitch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstDeviceSwitch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstDeviceSwitch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstDeviceSwitch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplay

`func (o *ConstDeviceSwitch) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConstDeviceSwitch) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConstDeviceSwitch) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConstDeviceSwitch) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetEvolvedOs

`func (o *ConstDeviceSwitch) GetEvolvedOs() bool`

GetEvolvedOs returns the EvolvedOs field if non-nil, zero value otherwise.

### GetEvolvedOsOk

`func (o *ConstDeviceSwitch) GetEvolvedOsOk() (*bool, bool)`

GetEvolvedOsOk returns a tuple with the EvolvedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvolvedOs

`func (o *ConstDeviceSwitch) SetEvolvedOs(v bool)`

SetEvolvedOs sets EvolvedOs field to given value.

### HasEvolvedOs

`func (o *ConstDeviceSwitch) HasEvolvedOs() bool`

HasEvolvedOs returns a boolean if a field has been set.

### GetEvpnRiType

`func (o *ConstDeviceSwitch) GetEvpnRiType() string`

GetEvpnRiType returns the EvpnRiType field if non-nil, zero value otherwise.

### GetEvpnRiTypeOk

`func (o *ConstDeviceSwitch) GetEvpnRiTypeOk() (*string, bool)`

GetEvpnRiTypeOk returns a tuple with the EvpnRiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnRiType

`func (o *ConstDeviceSwitch) SetEvpnRiType(v string)`

SetEvpnRiType sets EvpnRiType field to given value.

### HasEvpnRiType

`func (o *ConstDeviceSwitch) HasEvpnRiType() bool`

HasEvpnRiType returns a boolean if a field has been set.

### GetExperimental

`func (o *ConstDeviceSwitch) GetExperimental() bool`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *ConstDeviceSwitch) GetExperimentalOk() (*bool, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *ConstDeviceSwitch) SetExperimental(v bool)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *ConstDeviceSwitch) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### GetFansPluggable

`func (o *ConstDeviceSwitch) GetFansPluggable() bool`

GetFansPluggable returns the FansPluggable field if non-nil, zero value otherwise.

### GetFansPluggableOk

`func (o *ConstDeviceSwitch) GetFansPluggableOk() (*bool, bool)`

GetFansPluggableOk returns a tuple with the FansPluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFansPluggable

`func (o *ConstDeviceSwitch) SetFansPluggable(v bool)`

SetFansPluggable sets FansPluggable field to given value.

### HasFansPluggable

`func (o *ConstDeviceSwitch) HasFansPluggable() bool`

HasFansPluggable returns a boolean if a field has been set.

### GetHasBgp

`func (o *ConstDeviceSwitch) GetHasBgp() bool`

GetHasBgp returns the HasBgp field if non-nil, zero value otherwise.

### GetHasBgpOk

`func (o *ConstDeviceSwitch) GetHasBgpOk() (*bool, bool)`

GetHasBgpOk returns a tuple with the HasBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBgp

`func (o *ConstDeviceSwitch) SetHasBgp(v bool)`

SetHasBgp sets HasBgp field to given value.

### HasHasBgp

`func (o *ConstDeviceSwitch) HasHasBgp() bool`

HasHasBgp returns a boolean if a field has been set.

### GetHasEts

`func (o *ConstDeviceSwitch) GetHasEts() bool`

GetHasEts returns the HasEts field if non-nil, zero value otherwise.

### GetHasEtsOk

`func (o *ConstDeviceSwitch) GetHasEtsOk() (*bool, bool)`

GetHasEtsOk returns a tuple with the HasEts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEts

`func (o *ConstDeviceSwitch) SetHasEts(v bool)`

SetHasEts sets HasEts field to given value.

### HasHasEts

`func (o *ConstDeviceSwitch) HasHasEts() bool`

HasHasEts returns a boolean if a field has been set.

### GetHasEvpn

`func (o *ConstDeviceSwitch) GetHasEvpn() bool`

GetHasEvpn returns the HasEvpn field if non-nil, zero value otherwise.

### GetHasEvpnOk

`func (o *ConstDeviceSwitch) GetHasEvpnOk() (*bool, bool)`

GetHasEvpnOk returns a tuple with the HasEvpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEvpn

`func (o *ConstDeviceSwitch) SetHasEvpn(v bool)`

SetHasEvpn sets HasEvpn field to given value.

### HasHasEvpn

`func (o *ConstDeviceSwitch) HasHasEvpn() bool`

HasHasEvpn returns a boolean if a field has been set.

### GetHasIrb

`func (o *ConstDeviceSwitch) GetHasIrb() bool`

GetHasIrb returns the HasIrb field if non-nil, zero value otherwise.

### GetHasIrbOk

`func (o *ConstDeviceSwitch) GetHasIrbOk() (*bool, bool)`

GetHasIrbOk returns a tuple with the HasIrb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIrb

`func (o *ConstDeviceSwitch) SetHasIrb(v bool)`

SetHasIrb sets HasIrb field to given value.

### HasHasIrb

`func (o *ConstDeviceSwitch) HasHasIrb() bool`

HasHasIrb returns a boolean if a field has been set.

### GetHasPoeOut

`func (o *ConstDeviceSwitch) GetHasPoeOut() bool`

GetHasPoeOut returns the HasPoeOut field if non-nil, zero value otherwise.

### GetHasPoeOutOk

`func (o *ConstDeviceSwitch) GetHasPoeOutOk() (*bool, bool)`

GetHasPoeOutOk returns a tuple with the HasPoeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPoeOut

`func (o *ConstDeviceSwitch) SetHasPoeOut(v bool)`

SetHasPoeOut sets HasPoeOut field to given value.

### HasHasPoeOut

`func (o *ConstDeviceSwitch) HasHasPoeOut() bool`

HasHasPoeOut returns a boolean if a field has been set.

### GetHasSnapshot

`func (o *ConstDeviceSwitch) GetHasSnapshot() bool`

GetHasSnapshot returns the HasSnapshot field if non-nil, zero value otherwise.

### GetHasSnapshotOk

`func (o *ConstDeviceSwitch) GetHasSnapshotOk() (*bool, bool)`

GetHasSnapshotOk returns a tuple with the HasSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshot

`func (o *ConstDeviceSwitch) SetHasSnapshot(v bool)`

SetHasSnapshot sets HasSnapshot field to given value.

### HasHasSnapshot

`func (o *ConstDeviceSwitch) HasHasSnapshot() bool`

HasHasSnapshot returns a boolean if a field has been set.

### GetHasVc

`func (o *ConstDeviceSwitch) GetHasVc() bool`

GetHasVc returns the HasVc field if non-nil, zero value otherwise.

### GetHasVcOk

`func (o *ConstDeviceSwitch) GetHasVcOk() (*bool, bool)`

GetHasVcOk returns a tuple with the HasVc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVc

`func (o *ConstDeviceSwitch) SetHasVc(v bool)`

SetHasVc sets HasVc field to given value.

### HasHasVc

`func (o *ConstDeviceSwitch) HasHasVc() bool`

HasHasVc returns a boolean if a field has been set.

### GetModel

`func (o *ConstDeviceSwitch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConstDeviceSwitch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConstDeviceSwitch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConstDeviceSwitch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModular

`func (o *ConstDeviceSwitch) GetModular() bool`

GetModular returns the Modular field if non-nil, zero value otherwise.

### GetModularOk

`func (o *ConstDeviceSwitch) GetModularOk() (*bool, bool)`

GetModularOk returns a tuple with the Modular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModular

`func (o *ConstDeviceSwitch) SetModular(v bool)`

SetModular sets Modular field to given value.

### HasModular

`func (o *ConstDeviceSwitch) HasModular() bool`

HasModular returns a boolean if a field has been set.

### GetNoShapingRate

`func (o *ConstDeviceSwitch) GetNoShapingRate() bool`

GetNoShapingRate returns the NoShapingRate field if non-nil, zero value otherwise.

### GetNoShapingRateOk

`func (o *ConstDeviceSwitch) GetNoShapingRateOk() (*bool, bool)`

GetNoShapingRateOk returns a tuple with the NoShapingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShapingRate

`func (o *ConstDeviceSwitch) SetNoShapingRate(v bool)`

SetNoShapingRate sets NoShapingRate field to given value.

### HasNoShapingRate

`func (o *ConstDeviceSwitch) HasNoShapingRate() bool`

HasNoShapingRate returns a boolean if a field has been set.

### GetNumberFans

`func (o *ConstDeviceSwitch) GetNumberFans() int32`

GetNumberFans returns the NumberFans field if non-nil, zero value otherwise.

### GetNumberFansOk

`func (o *ConstDeviceSwitch) GetNumberFansOk() (*int32, bool)`

GetNumberFansOk returns a tuple with the NumberFans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFans

`func (o *ConstDeviceSwitch) SetNumberFans(v int32)`

SetNumberFans sets NumberFans field to given value.

### HasNumberFans

`func (o *ConstDeviceSwitch) HasNumberFans() bool`

HasNumberFans returns a boolean if a field has been set.

### GetOcDevice

`func (o *ConstDeviceSwitch) GetOcDevice() bool`

GetOcDevice returns the OcDevice field if non-nil, zero value otherwise.

### GetOcDeviceOk

`func (o *ConstDeviceSwitch) GetOcDeviceOk() (*bool, bool)`

GetOcDeviceOk returns a tuple with the OcDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcDevice

`func (o *ConstDeviceSwitch) SetOcDevice(v bool)`

SetOcDevice sets OcDevice field to given value.

### HasOcDevice

`func (o *ConstDeviceSwitch) HasOcDevice() bool`

HasOcDevice returns a boolean if a field has been set.

### GetOobInterface

`func (o *ConstDeviceSwitch) GetOobInterface() string`

GetOobInterface returns the OobInterface field if non-nil, zero value otherwise.

### GetOobInterfaceOk

`func (o *ConstDeviceSwitch) GetOobInterfaceOk() (*string, bool)`

GetOobInterfaceOk returns a tuple with the OobInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobInterface

`func (o *ConstDeviceSwitch) SetOobInterface(v string)`

SetOobInterface sets OobInterface field to given value.

### HasOobInterface

`func (o *ConstDeviceSwitch) HasOobInterface() bool`

HasOobInterface returns a boolean if a field has been set.

### GetPacketActionDropOnly

`func (o *ConstDeviceSwitch) GetPacketActionDropOnly() bool`

GetPacketActionDropOnly returns the PacketActionDropOnly field if non-nil, zero value otherwise.

### GetPacketActionDropOnlyOk

`func (o *ConstDeviceSwitch) GetPacketActionDropOnlyOk() (*bool, bool)`

GetPacketActionDropOnlyOk returns a tuple with the PacketActionDropOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketActionDropOnly

`func (o *ConstDeviceSwitch) SetPacketActionDropOnly(v bool)`

SetPacketActionDropOnly sets PacketActionDropOnly field to given value.

### HasPacketActionDropOnly

`func (o *ConstDeviceSwitch) HasPacketActionDropOnly() bool`

HasPacketActionDropOnly returns a boolean if a field has been set.

### GetPic

`func (o *ConstDeviceSwitch) GetPic() map[string]string`

GetPic returns the Pic field if non-nil, zero value otherwise.

### GetPicOk

`func (o *ConstDeviceSwitch) GetPicOk() (*map[string]string, bool)`

GetPicOk returns a tuple with the Pic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPic

`func (o *ConstDeviceSwitch) SetPic(v map[string]string)`

SetPic sets Pic field to given value.

### HasPic

`func (o *ConstDeviceSwitch) HasPic() bool`

HasPic returns a boolean if a field has been set.

### GetSubRequired

`func (o *ConstDeviceSwitch) GetSubRequired() string`

GetSubRequired returns the SubRequired field if non-nil, zero value otherwise.

### GetSubRequiredOk

`func (o *ConstDeviceSwitch) GetSubRequiredOk() (*string, bool)`

GetSubRequiredOk returns a tuple with the SubRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRequired

`func (o *ConstDeviceSwitch) SetSubRequired(v string)`

SetSubRequired sets SubRequired field to given value.

### HasSubRequired

`func (o *ConstDeviceSwitch) HasSubRequired() bool`

HasSubRequired returns a boolean if a field has been set.

### GetType

`func (o *ConstDeviceSwitch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstDeviceSwitch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstDeviceSwitch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstDeviceSwitch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


