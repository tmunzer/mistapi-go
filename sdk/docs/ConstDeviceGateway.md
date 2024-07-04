# ConstDeviceGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Defaults** | Pointer to **map[string]string** | Object Key is the interface type name (e.g. \&quot;lan_ports\&quot;, \&quot;wan_ports\&quot;, ...) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Experimental** | Pointer to **bool** |  | [optional] [default to false]
**FansPluggable** | Pointer to **bool** |  | [optional] [default to true]
**HaNode0Fpc** | Pointer to **int32** |  | [optional] 
**HaNode1Fpc** | Pointer to **int32** |  | [optional] 
**HasBgp** | Pointer to **bool** |  | [optional] [default to false]
**HasFxp0** | Pointer to **bool** |  | [optional] [default to true]
**HasHaControl** | Pointer to **bool** |  | [optional] [default to false]
**HasHaData** | Pointer to **bool** |  | [optional] [default to false]
**HasIrb** | Pointer to **bool** |  | [optional] [default to false]
**HasPoeOut** | Pointer to **bool** |  | [optional] [default to true]
**HasSnapshot** | Pointer to **bool** |  | [optional] [default to true]
**IrbDisabledByDefault** | Pointer to **bool** |  | [optional] [default to false]
**Model** | Pointer to **string** |  | [optional] 
**NumberFans** | Pointer to **int32** |  | [optional] 
**OcDevice** | Pointer to **bool** |  | [optional] [default to false]
**Pic** | Pointer to **map[string]string** | Object Key is the PIC number | [optional] 
**Ports** | Pointer to [**ConstDeviceGatewayPorts**](ConstDeviceGatewayPorts.md) |  | [optional] 
**SubRequired** | Pointer to **string** |  | [optional] 
**T128Device** | Pointer to **bool** |  | [optional] [default to false]
**Type** | Pointer to **string** |  | [optional] [default to "gateway"]

## Methods

### NewConstDeviceGateway

`func NewConstDeviceGateway() *ConstDeviceGateway`

NewConstDeviceGateway instantiates a new ConstDeviceGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstDeviceGatewayWithDefaults

`func NewConstDeviceGatewayWithDefaults() *ConstDeviceGateway`

NewConstDeviceGatewayWithDefaults instantiates a new ConstDeviceGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaults

`func (o *ConstDeviceGateway) GetDefaults() map[string]string`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *ConstDeviceGateway) GetDefaultsOk() (*map[string]string, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *ConstDeviceGateway) SetDefaults(v map[string]string)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *ConstDeviceGateway) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDescription

`func (o *ConstDeviceGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstDeviceGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstDeviceGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstDeviceGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExperimental

`func (o *ConstDeviceGateway) GetExperimental() bool`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *ConstDeviceGateway) GetExperimentalOk() (*bool, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *ConstDeviceGateway) SetExperimental(v bool)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *ConstDeviceGateway) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### GetFansPluggable

`func (o *ConstDeviceGateway) GetFansPluggable() bool`

GetFansPluggable returns the FansPluggable field if non-nil, zero value otherwise.

### GetFansPluggableOk

`func (o *ConstDeviceGateway) GetFansPluggableOk() (*bool, bool)`

GetFansPluggableOk returns a tuple with the FansPluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFansPluggable

`func (o *ConstDeviceGateway) SetFansPluggable(v bool)`

SetFansPluggable sets FansPluggable field to given value.

### HasFansPluggable

`func (o *ConstDeviceGateway) HasFansPluggable() bool`

HasFansPluggable returns a boolean if a field has been set.

### GetHaNode0Fpc

`func (o *ConstDeviceGateway) GetHaNode0Fpc() int32`

GetHaNode0Fpc returns the HaNode0Fpc field if non-nil, zero value otherwise.

### GetHaNode0FpcOk

`func (o *ConstDeviceGateway) GetHaNode0FpcOk() (*int32, bool)`

GetHaNode0FpcOk returns a tuple with the HaNode0Fpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNode0Fpc

`func (o *ConstDeviceGateway) SetHaNode0Fpc(v int32)`

SetHaNode0Fpc sets HaNode0Fpc field to given value.

### HasHaNode0Fpc

`func (o *ConstDeviceGateway) HasHaNode0Fpc() bool`

HasHaNode0Fpc returns a boolean if a field has been set.

### GetHaNode1Fpc

`func (o *ConstDeviceGateway) GetHaNode1Fpc() int32`

GetHaNode1Fpc returns the HaNode1Fpc field if non-nil, zero value otherwise.

### GetHaNode1FpcOk

`func (o *ConstDeviceGateway) GetHaNode1FpcOk() (*int32, bool)`

GetHaNode1FpcOk returns a tuple with the HaNode1Fpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNode1Fpc

`func (o *ConstDeviceGateway) SetHaNode1Fpc(v int32)`

SetHaNode1Fpc sets HaNode1Fpc field to given value.

### HasHaNode1Fpc

`func (o *ConstDeviceGateway) HasHaNode1Fpc() bool`

HasHaNode1Fpc returns a boolean if a field has been set.

### GetHasBgp

`func (o *ConstDeviceGateway) GetHasBgp() bool`

GetHasBgp returns the HasBgp field if non-nil, zero value otherwise.

### GetHasBgpOk

`func (o *ConstDeviceGateway) GetHasBgpOk() (*bool, bool)`

GetHasBgpOk returns a tuple with the HasBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBgp

`func (o *ConstDeviceGateway) SetHasBgp(v bool)`

SetHasBgp sets HasBgp field to given value.

### HasHasBgp

`func (o *ConstDeviceGateway) HasHasBgp() bool`

HasHasBgp returns a boolean if a field has been set.

### GetHasFxp0

`func (o *ConstDeviceGateway) GetHasFxp0() bool`

GetHasFxp0 returns the HasFxp0 field if non-nil, zero value otherwise.

### GetHasFxp0Ok

`func (o *ConstDeviceGateway) GetHasFxp0Ok() (*bool, bool)`

GetHasFxp0Ok returns a tuple with the HasFxp0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFxp0

`func (o *ConstDeviceGateway) SetHasFxp0(v bool)`

SetHasFxp0 sets HasFxp0 field to given value.

### HasHasFxp0

`func (o *ConstDeviceGateway) HasHasFxp0() bool`

HasHasFxp0 returns a boolean if a field has been set.

### GetHasHaControl

`func (o *ConstDeviceGateway) GetHasHaControl() bool`

GetHasHaControl returns the HasHaControl field if non-nil, zero value otherwise.

### GetHasHaControlOk

`func (o *ConstDeviceGateway) GetHasHaControlOk() (*bool, bool)`

GetHasHaControlOk returns a tuple with the HasHaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHaControl

`func (o *ConstDeviceGateway) SetHasHaControl(v bool)`

SetHasHaControl sets HasHaControl field to given value.

### HasHasHaControl

`func (o *ConstDeviceGateway) HasHasHaControl() bool`

HasHasHaControl returns a boolean if a field has been set.

### GetHasHaData

`func (o *ConstDeviceGateway) GetHasHaData() bool`

GetHasHaData returns the HasHaData field if non-nil, zero value otherwise.

### GetHasHaDataOk

`func (o *ConstDeviceGateway) GetHasHaDataOk() (*bool, bool)`

GetHasHaDataOk returns a tuple with the HasHaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHaData

`func (o *ConstDeviceGateway) SetHasHaData(v bool)`

SetHasHaData sets HasHaData field to given value.

### HasHasHaData

`func (o *ConstDeviceGateway) HasHasHaData() bool`

HasHasHaData returns a boolean if a field has been set.

### GetHasIrb

`func (o *ConstDeviceGateway) GetHasIrb() bool`

GetHasIrb returns the HasIrb field if non-nil, zero value otherwise.

### GetHasIrbOk

`func (o *ConstDeviceGateway) GetHasIrbOk() (*bool, bool)`

GetHasIrbOk returns a tuple with the HasIrb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIrb

`func (o *ConstDeviceGateway) SetHasIrb(v bool)`

SetHasIrb sets HasIrb field to given value.

### HasHasIrb

`func (o *ConstDeviceGateway) HasHasIrb() bool`

HasHasIrb returns a boolean if a field has been set.

### GetHasPoeOut

`func (o *ConstDeviceGateway) GetHasPoeOut() bool`

GetHasPoeOut returns the HasPoeOut field if non-nil, zero value otherwise.

### GetHasPoeOutOk

`func (o *ConstDeviceGateway) GetHasPoeOutOk() (*bool, bool)`

GetHasPoeOutOk returns a tuple with the HasPoeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPoeOut

`func (o *ConstDeviceGateway) SetHasPoeOut(v bool)`

SetHasPoeOut sets HasPoeOut field to given value.

### HasHasPoeOut

`func (o *ConstDeviceGateway) HasHasPoeOut() bool`

HasHasPoeOut returns a boolean if a field has been set.

### GetHasSnapshot

`func (o *ConstDeviceGateway) GetHasSnapshot() bool`

GetHasSnapshot returns the HasSnapshot field if non-nil, zero value otherwise.

### GetHasSnapshotOk

`func (o *ConstDeviceGateway) GetHasSnapshotOk() (*bool, bool)`

GetHasSnapshotOk returns a tuple with the HasSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshot

`func (o *ConstDeviceGateway) SetHasSnapshot(v bool)`

SetHasSnapshot sets HasSnapshot field to given value.

### HasHasSnapshot

`func (o *ConstDeviceGateway) HasHasSnapshot() bool`

HasHasSnapshot returns a boolean if a field has been set.

### GetIrbDisabledByDefault

`func (o *ConstDeviceGateway) GetIrbDisabledByDefault() bool`

GetIrbDisabledByDefault returns the IrbDisabledByDefault field if non-nil, zero value otherwise.

### GetIrbDisabledByDefaultOk

`func (o *ConstDeviceGateway) GetIrbDisabledByDefaultOk() (*bool, bool)`

GetIrbDisabledByDefaultOk returns a tuple with the IrbDisabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrbDisabledByDefault

`func (o *ConstDeviceGateway) SetIrbDisabledByDefault(v bool)`

SetIrbDisabledByDefault sets IrbDisabledByDefault field to given value.

### HasIrbDisabledByDefault

`func (o *ConstDeviceGateway) HasIrbDisabledByDefault() bool`

HasIrbDisabledByDefault returns a boolean if a field has been set.

### GetModel

`func (o *ConstDeviceGateway) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConstDeviceGateway) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConstDeviceGateway) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConstDeviceGateway) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNumberFans

`func (o *ConstDeviceGateway) GetNumberFans() int32`

GetNumberFans returns the NumberFans field if non-nil, zero value otherwise.

### GetNumberFansOk

`func (o *ConstDeviceGateway) GetNumberFansOk() (*int32, bool)`

GetNumberFansOk returns a tuple with the NumberFans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFans

`func (o *ConstDeviceGateway) SetNumberFans(v int32)`

SetNumberFans sets NumberFans field to given value.

### HasNumberFans

`func (o *ConstDeviceGateway) HasNumberFans() bool`

HasNumberFans returns a boolean if a field has been set.

### GetOcDevice

`func (o *ConstDeviceGateway) GetOcDevice() bool`

GetOcDevice returns the OcDevice field if non-nil, zero value otherwise.

### GetOcDeviceOk

`func (o *ConstDeviceGateway) GetOcDeviceOk() (*bool, bool)`

GetOcDeviceOk returns a tuple with the OcDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcDevice

`func (o *ConstDeviceGateway) SetOcDevice(v bool)`

SetOcDevice sets OcDevice field to given value.

### HasOcDevice

`func (o *ConstDeviceGateway) HasOcDevice() bool`

HasOcDevice returns a boolean if a field has been set.

### GetPic

`func (o *ConstDeviceGateway) GetPic() map[string]string`

GetPic returns the Pic field if non-nil, zero value otherwise.

### GetPicOk

`func (o *ConstDeviceGateway) GetPicOk() (*map[string]string, bool)`

GetPicOk returns a tuple with the Pic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPic

`func (o *ConstDeviceGateway) SetPic(v map[string]string)`

SetPic sets Pic field to given value.

### HasPic

`func (o *ConstDeviceGateway) HasPic() bool`

HasPic returns a boolean if a field has been set.

### GetPorts

`func (o *ConstDeviceGateway) GetPorts() ConstDeviceGatewayPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ConstDeviceGateway) GetPortsOk() (*ConstDeviceGatewayPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ConstDeviceGateway) SetPorts(v ConstDeviceGatewayPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ConstDeviceGateway) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSubRequired

`func (o *ConstDeviceGateway) GetSubRequired() string`

GetSubRequired returns the SubRequired field if non-nil, zero value otherwise.

### GetSubRequiredOk

`func (o *ConstDeviceGateway) GetSubRequiredOk() (*string, bool)`

GetSubRequiredOk returns a tuple with the SubRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRequired

`func (o *ConstDeviceGateway) SetSubRequired(v string)`

SetSubRequired sets SubRequired field to given value.

### HasSubRequired

`func (o *ConstDeviceGateway) HasSubRequired() bool`

HasSubRequired returns a boolean if a field has been set.

### GetT128Device

`func (o *ConstDeviceGateway) GetT128Device() bool`

GetT128Device returns the T128Device field if non-nil, zero value otherwise.

### GetT128DeviceOk

`func (o *ConstDeviceGateway) GetT128DeviceOk() (*bool, bool)`

GetT128DeviceOk returns a tuple with the T128Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT128Device

`func (o *ConstDeviceGateway) SetT128Device(v bool)`

SetT128Device sets T128Device field to given value.

### HasT128Device

`func (o *ConstDeviceGateway) HasT128Device() bool`

HasT128Device returns a boolean if a field has been set.

### GetType

`func (o *ConstDeviceGateway) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstDeviceGateway) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstDeviceGateway) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstDeviceGateway) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


