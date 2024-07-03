# ConstDeviceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApType** | Pointer to **string** |  | [optional] 
**Band24** | Pointer to [**ConstDeviceApBand24**](ConstDeviceApBand24.md) |  | [optional] 
**Band5** | Pointer to [**ConstDeviceApBand5**](ConstDeviceApBand5.md) |  | [optional] 
**Band6** | Pointer to [**ConstDeviceApBand6**](ConstDeviceApBand6.md) |  | [optional] 
**Band24Usages** | Pointer to [**RadioBand24Usage**](RadioBand24Usage.md) |  | [optional] 
**CeDfsOk** | Pointer to **bool** |  | [optional] 
**CiscoPace** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisallowedChannels** | Pointer to [**map[string]map[string][]int32**](map.md) | Property key is a list of country codes (e.g. \&quot;GB, DE\&quot;) | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Extio** | Pointer to [**map[string]ConstDeviceApExtios**](ConstDeviceApExtios.md) | Property key is the GPIO port name (e.g. \&quot;D0\&quot;, \&quot;A1\&quot;) | [optional] 
**FccDfsOk** | Pointer to **bool** |  | [optional] 
**Has11ax** | Pointer to **bool** |  | [optional] 
**HasCompass** | Pointer to **bool** |  | [optional] 
**HasExtAnt** | Pointer to **bool** |  | [optional] 
**HasExtio** | Pointer to **bool** |  | [optional] 
**HasHeight** | Pointer to **bool** |  | [optional] 
**HasModulePort** | Pointer to **bool** |  | [optional] 
**HasPoeOut** | Pointer to **bool** |  | [optional] [default to true]
**HasScanningRadio** | Pointer to **bool** |  | [optional] 
**HasSelectableRadio** | Pointer to **bool** |  | [optional] 
**HasUsb** | Pointer to **bool** |  | [optional] 
**HasVble** | Pointer to **bool** |  | [optional] 
**HasWifiBand24** | Pointer to **bool** |  | [optional] 
**HasWifiBand5** | Pointer to **bool** |  | [optional] 
**HasWifiBand6** | Pointer to **bool** |  | [optional] 
**MaxPoeOut** | Pointer to **int32** |  | [optional] 
**MaxWlans** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**OtherDfsOk** | Pointer to **bool** |  | [optional] 
**Outdoor** | Pointer to **bool** |  | [optional] 
**Radios** | Pointer to **map[string]string** | Property key is the radio number (e.g. r0, r1, ...). Property value is the RF band (e.g. \&quot;24\&quot;, \&quot;5\&quot;, ...) | [optional] 
**SharedScanningRadio** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "gateway"]
**Unmanaged** | Pointer to **bool** |  | [optional] 
**Vble** | Pointer to [**ConstDeviceApVble**](ConstDeviceApVble.md) |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Defaults** | Pointer to **map[string]string** | Object Key is the interface type name (e.g. \&quot;lan_ports\&quot;, \&quot;wan_ports\&quot;, ...) | [optional] 
**EvolvedOs** | Pointer to **bool** |  | [optional] [default to false]
**EvpnRiType** | Pointer to **string** |  | [optional] 
**Experimental** | Pointer to **bool** |  | [optional] [default to false]
**FansPluggable** | Pointer to **bool** |  | [optional] [default to true]
**HasBgp** | Pointer to **bool** |  | [optional] [default to false]
**HasEts** | Pointer to **bool** |  | [optional] [default to false]
**HasEvpn** | Pointer to **bool** |  | [optional] [default to false]
**HasIrb** | Pointer to **bool** |  | [optional] [default to false]
**HasSnapshot** | Pointer to **bool** |  | [optional] [default to true]
**HasVc** | Pointer to **bool** |  | [optional] [default to true]
**Modular** | Pointer to **bool** |  | [optional] [default to false]
**NoShapingRate** | Pointer to **bool** |  | [optional] [default to false]
**NumberFans** | Pointer to **int32** |  | [optional] 
**OcDevice** | Pointer to **bool** |  | [optional] [default to false]
**OobInterface** | Pointer to **string** |  | [optional] 
**PacketActionDropOnly** | Pointer to **bool** |  | [optional] [default to false]
**Pic** | Pointer to **map[string]string** | Object Key is the PIC number | [optional] 
**SubRequired** | Pointer to **string** |  | [optional] 
**HaNode0Fpc** | Pointer to **int32** |  | [optional] 
**HaNode1Fpc** | Pointer to **int32** |  | [optional] 
**HasFxp0** | Pointer to **bool** |  | [optional] [default to true]
**HasHaControl** | Pointer to **bool** |  | [optional] [default to false]
**HasHaData** | Pointer to **bool** |  | [optional] [default to false]
**IrbDisabledByDefault** | Pointer to **bool** |  | [optional] [default to false]
**Ports** | Pointer to [**ConstDeviceGatewayPorts**](ConstDeviceGatewayPorts.md) |  | [optional] 
**T128Device** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewConstDeviceModel

`func NewConstDeviceModel() *ConstDeviceModel`

NewConstDeviceModel instantiates a new ConstDeviceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstDeviceModelWithDefaults

`func NewConstDeviceModelWithDefaults() *ConstDeviceModel`

NewConstDeviceModelWithDefaults instantiates a new ConstDeviceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApType

`func (o *ConstDeviceModel) GetApType() string`

GetApType returns the ApType field if non-nil, zero value otherwise.

### GetApTypeOk

`func (o *ConstDeviceModel) GetApTypeOk() (*string, bool)`

GetApTypeOk returns a tuple with the ApType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApType

`func (o *ConstDeviceModel) SetApType(v string)`

SetApType sets ApType field to given value.

### HasApType

`func (o *ConstDeviceModel) HasApType() bool`

HasApType returns a boolean if a field has been set.

### GetBand24

`func (o *ConstDeviceModel) GetBand24() ConstDeviceApBand24`

GetBand24 returns the Band24 field if non-nil, zero value otherwise.

### GetBand24Ok

`func (o *ConstDeviceModel) GetBand24Ok() (*ConstDeviceApBand24, bool)`

GetBand24Ok returns a tuple with the Band24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24

`func (o *ConstDeviceModel) SetBand24(v ConstDeviceApBand24)`

SetBand24 sets Band24 field to given value.

### HasBand24

`func (o *ConstDeviceModel) HasBand24() bool`

HasBand24 returns a boolean if a field has been set.

### GetBand5

`func (o *ConstDeviceModel) GetBand5() ConstDeviceApBand5`

GetBand5 returns the Band5 field if non-nil, zero value otherwise.

### GetBand5Ok

`func (o *ConstDeviceModel) GetBand5Ok() (*ConstDeviceApBand5, bool)`

GetBand5Ok returns a tuple with the Band5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5

`func (o *ConstDeviceModel) SetBand5(v ConstDeviceApBand5)`

SetBand5 sets Band5 field to given value.

### HasBand5

`func (o *ConstDeviceModel) HasBand5() bool`

HasBand5 returns a boolean if a field has been set.

### GetBand6

`func (o *ConstDeviceModel) GetBand6() ConstDeviceApBand6`

GetBand6 returns the Band6 field if non-nil, zero value otherwise.

### GetBand6Ok

`func (o *ConstDeviceModel) GetBand6Ok() (*ConstDeviceApBand6, bool)`

GetBand6Ok returns a tuple with the Band6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6

`func (o *ConstDeviceModel) SetBand6(v ConstDeviceApBand6)`

SetBand6 sets Band6 field to given value.

### HasBand6

`func (o *ConstDeviceModel) HasBand6() bool`

HasBand6 returns a boolean if a field has been set.

### GetBand24Usages

`func (o *ConstDeviceModel) GetBand24Usages() RadioBand24Usage`

GetBand24Usages returns the Band24Usages field if non-nil, zero value otherwise.

### GetBand24UsagesOk

`func (o *ConstDeviceModel) GetBand24UsagesOk() (*RadioBand24Usage, bool)`

GetBand24UsagesOk returns a tuple with the Band24Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Usages

`func (o *ConstDeviceModel) SetBand24Usages(v RadioBand24Usage)`

SetBand24Usages sets Band24Usages field to given value.

### HasBand24Usages

`func (o *ConstDeviceModel) HasBand24Usages() bool`

HasBand24Usages returns a boolean if a field has been set.

### GetCeDfsOk

`func (o *ConstDeviceModel) GetCeDfsOk() bool`

GetCeDfsOk returns the CeDfsOk field if non-nil, zero value otherwise.

### GetCeDfsOkOk

`func (o *ConstDeviceModel) GetCeDfsOkOk() (*bool, bool)`

GetCeDfsOkOk returns a tuple with the CeDfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeDfsOk

`func (o *ConstDeviceModel) SetCeDfsOk(v bool)`

SetCeDfsOk sets CeDfsOk field to given value.

### HasCeDfsOk

`func (o *ConstDeviceModel) HasCeDfsOk() bool`

HasCeDfsOk returns a boolean if a field has been set.

### GetCiscoPace

`func (o *ConstDeviceModel) GetCiscoPace() bool`

GetCiscoPace returns the CiscoPace field if non-nil, zero value otherwise.

### GetCiscoPaceOk

`func (o *ConstDeviceModel) GetCiscoPaceOk() (*bool, bool)`

GetCiscoPaceOk returns a tuple with the CiscoPace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoPace

`func (o *ConstDeviceModel) SetCiscoPace(v bool)`

SetCiscoPace sets CiscoPace field to given value.

### HasCiscoPace

`func (o *ConstDeviceModel) HasCiscoPace() bool`

HasCiscoPace returns a boolean if a field has been set.

### GetDescription

`func (o *ConstDeviceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstDeviceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstDeviceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstDeviceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisallowedChannels

`func (o *ConstDeviceModel) GetDisallowedChannels() map[string]map[string][]int32`

GetDisallowedChannels returns the DisallowedChannels field if non-nil, zero value otherwise.

### GetDisallowedChannelsOk

`func (o *ConstDeviceModel) GetDisallowedChannelsOk() (*map[string]map[string][]int32, bool)`

GetDisallowedChannelsOk returns a tuple with the DisallowedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedChannels

`func (o *ConstDeviceModel) SetDisallowedChannels(v map[string]map[string][]int32)`

SetDisallowedChannels sets DisallowedChannels field to given value.

### HasDisallowedChannels

`func (o *ConstDeviceModel) HasDisallowedChannels() bool`

HasDisallowedChannels returns a boolean if a field has been set.

### GetDisplay

`func (o *ConstDeviceModel) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConstDeviceModel) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConstDeviceModel) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConstDeviceModel) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetExtio

`func (o *ConstDeviceModel) GetExtio() map[string]ConstDeviceApExtios`

GetExtio returns the Extio field if non-nil, zero value otherwise.

### GetExtioOk

`func (o *ConstDeviceModel) GetExtioOk() (*map[string]ConstDeviceApExtios, bool)`

GetExtioOk returns a tuple with the Extio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtio

`func (o *ConstDeviceModel) SetExtio(v map[string]ConstDeviceApExtios)`

SetExtio sets Extio field to given value.

### HasExtio

`func (o *ConstDeviceModel) HasExtio() bool`

HasExtio returns a boolean if a field has been set.

### GetFccDfsOk

`func (o *ConstDeviceModel) GetFccDfsOk() bool`

GetFccDfsOk returns the FccDfsOk field if non-nil, zero value otherwise.

### GetFccDfsOkOk

`func (o *ConstDeviceModel) GetFccDfsOkOk() (*bool, bool)`

GetFccDfsOkOk returns a tuple with the FccDfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFccDfsOk

`func (o *ConstDeviceModel) SetFccDfsOk(v bool)`

SetFccDfsOk sets FccDfsOk field to given value.

### HasFccDfsOk

`func (o *ConstDeviceModel) HasFccDfsOk() bool`

HasFccDfsOk returns a boolean if a field has been set.

### GetHas11ax

`func (o *ConstDeviceModel) GetHas11ax() bool`

GetHas11ax returns the Has11ax field if non-nil, zero value otherwise.

### GetHas11axOk

`func (o *ConstDeviceModel) GetHas11axOk() (*bool, bool)`

GetHas11axOk returns a tuple with the Has11ax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas11ax

`func (o *ConstDeviceModel) SetHas11ax(v bool)`

SetHas11ax sets Has11ax field to given value.

### HasHas11ax

`func (o *ConstDeviceModel) HasHas11ax() bool`

HasHas11ax returns a boolean if a field has been set.

### GetHasCompass

`func (o *ConstDeviceModel) GetHasCompass() bool`

GetHasCompass returns the HasCompass field if non-nil, zero value otherwise.

### GetHasCompassOk

`func (o *ConstDeviceModel) GetHasCompassOk() (*bool, bool)`

GetHasCompassOk returns a tuple with the HasCompass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompass

`func (o *ConstDeviceModel) SetHasCompass(v bool)`

SetHasCompass sets HasCompass field to given value.

### HasHasCompass

`func (o *ConstDeviceModel) HasHasCompass() bool`

HasHasCompass returns a boolean if a field has been set.

### GetHasExtAnt

`func (o *ConstDeviceModel) GetHasExtAnt() bool`

GetHasExtAnt returns the HasExtAnt field if non-nil, zero value otherwise.

### GetHasExtAntOk

`func (o *ConstDeviceModel) GetHasExtAntOk() (*bool, bool)`

GetHasExtAntOk returns a tuple with the HasExtAnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExtAnt

`func (o *ConstDeviceModel) SetHasExtAnt(v bool)`

SetHasExtAnt sets HasExtAnt field to given value.

### HasHasExtAnt

`func (o *ConstDeviceModel) HasHasExtAnt() bool`

HasHasExtAnt returns a boolean if a field has been set.

### GetHasExtio

`func (o *ConstDeviceModel) GetHasExtio() bool`

GetHasExtio returns the HasExtio field if non-nil, zero value otherwise.

### GetHasExtioOk

`func (o *ConstDeviceModel) GetHasExtioOk() (*bool, bool)`

GetHasExtioOk returns a tuple with the HasExtio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExtio

`func (o *ConstDeviceModel) SetHasExtio(v bool)`

SetHasExtio sets HasExtio field to given value.

### HasHasExtio

`func (o *ConstDeviceModel) HasHasExtio() bool`

HasHasExtio returns a boolean if a field has been set.

### GetHasHeight

`func (o *ConstDeviceModel) GetHasHeight() bool`

GetHasHeight returns the HasHeight field if non-nil, zero value otherwise.

### GetHasHeightOk

`func (o *ConstDeviceModel) GetHasHeightOk() (*bool, bool)`

GetHasHeightOk returns a tuple with the HasHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHeight

`func (o *ConstDeviceModel) SetHasHeight(v bool)`

SetHasHeight sets HasHeight field to given value.

### HasHasHeight

`func (o *ConstDeviceModel) HasHasHeight() bool`

HasHasHeight returns a boolean if a field has been set.

### GetHasModulePort

`func (o *ConstDeviceModel) GetHasModulePort() bool`

GetHasModulePort returns the HasModulePort field if non-nil, zero value otherwise.

### GetHasModulePortOk

`func (o *ConstDeviceModel) GetHasModulePortOk() (*bool, bool)`

GetHasModulePortOk returns a tuple with the HasModulePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasModulePort

`func (o *ConstDeviceModel) SetHasModulePort(v bool)`

SetHasModulePort sets HasModulePort field to given value.

### HasHasModulePort

`func (o *ConstDeviceModel) HasHasModulePort() bool`

HasHasModulePort returns a boolean if a field has been set.

### GetHasPoeOut

`func (o *ConstDeviceModel) GetHasPoeOut() bool`

GetHasPoeOut returns the HasPoeOut field if non-nil, zero value otherwise.

### GetHasPoeOutOk

`func (o *ConstDeviceModel) GetHasPoeOutOk() (*bool, bool)`

GetHasPoeOutOk returns a tuple with the HasPoeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPoeOut

`func (o *ConstDeviceModel) SetHasPoeOut(v bool)`

SetHasPoeOut sets HasPoeOut field to given value.

### HasHasPoeOut

`func (o *ConstDeviceModel) HasHasPoeOut() bool`

HasHasPoeOut returns a boolean if a field has been set.

### GetHasScanningRadio

`func (o *ConstDeviceModel) GetHasScanningRadio() bool`

GetHasScanningRadio returns the HasScanningRadio field if non-nil, zero value otherwise.

### GetHasScanningRadioOk

`func (o *ConstDeviceModel) GetHasScanningRadioOk() (*bool, bool)`

GetHasScanningRadioOk returns a tuple with the HasScanningRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScanningRadio

`func (o *ConstDeviceModel) SetHasScanningRadio(v bool)`

SetHasScanningRadio sets HasScanningRadio field to given value.

### HasHasScanningRadio

`func (o *ConstDeviceModel) HasHasScanningRadio() bool`

HasHasScanningRadio returns a boolean if a field has been set.

### GetHasSelectableRadio

`func (o *ConstDeviceModel) GetHasSelectableRadio() bool`

GetHasSelectableRadio returns the HasSelectableRadio field if non-nil, zero value otherwise.

### GetHasSelectableRadioOk

`func (o *ConstDeviceModel) GetHasSelectableRadioOk() (*bool, bool)`

GetHasSelectableRadioOk returns a tuple with the HasSelectableRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectableRadio

`func (o *ConstDeviceModel) SetHasSelectableRadio(v bool)`

SetHasSelectableRadio sets HasSelectableRadio field to given value.

### HasHasSelectableRadio

`func (o *ConstDeviceModel) HasHasSelectableRadio() bool`

HasHasSelectableRadio returns a boolean if a field has been set.

### GetHasUsb

`func (o *ConstDeviceModel) GetHasUsb() bool`

GetHasUsb returns the HasUsb field if non-nil, zero value otherwise.

### GetHasUsbOk

`func (o *ConstDeviceModel) GetHasUsbOk() (*bool, bool)`

GetHasUsbOk returns a tuple with the HasUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsb

`func (o *ConstDeviceModel) SetHasUsb(v bool)`

SetHasUsb sets HasUsb field to given value.

### HasHasUsb

`func (o *ConstDeviceModel) HasHasUsb() bool`

HasHasUsb returns a boolean if a field has been set.

### GetHasVble

`func (o *ConstDeviceModel) GetHasVble() bool`

GetHasVble returns the HasVble field if non-nil, zero value otherwise.

### GetHasVbleOk

`func (o *ConstDeviceModel) GetHasVbleOk() (*bool, bool)`

GetHasVbleOk returns a tuple with the HasVble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVble

`func (o *ConstDeviceModel) SetHasVble(v bool)`

SetHasVble sets HasVble field to given value.

### HasHasVble

`func (o *ConstDeviceModel) HasHasVble() bool`

HasHasVble returns a boolean if a field has been set.

### GetHasWifiBand24

`func (o *ConstDeviceModel) GetHasWifiBand24() bool`

GetHasWifiBand24 returns the HasWifiBand24 field if non-nil, zero value otherwise.

### GetHasWifiBand24Ok

`func (o *ConstDeviceModel) GetHasWifiBand24Ok() (*bool, bool)`

GetHasWifiBand24Ok returns a tuple with the HasWifiBand24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWifiBand24

`func (o *ConstDeviceModel) SetHasWifiBand24(v bool)`

SetHasWifiBand24 sets HasWifiBand24 field to given value.

### HasHasWifiBand24

`func (o *ConstDeviceModel) HasHasWifiBand24() bool`

HasHasWifiBand24 returns a boolean if a field has been set.

### GetHasWifiBand5

`func (o *ConstDeviceModel) GetHasWifiBand5() bool`

GetHasWifiBand5 returns the HasWifiBand5 field if non-nil, zero value otherwise.

### GetHasWifiBand5Ok

`func (o *ConstDeviceModel) GetHasWifiBand5Ok() (*bool, bool)`

GetHasWifiBand5Ok returns a tuple with the HasWifiBand5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWifiBand5

`func (o *ConstDeviceModel) SetHasWifiBand5(v bool)`

SetHasWifiBand5 sets HasWifiBand5 field to given value.

### HasHasWifiBand5

`func (o *ConstDeviceModel) HasHasWifiBand5() bool`

HasHasWifiBand5 returns a boolean if a field has been set.

### GetHasWifiBand6

`func (o *ConstDeviceModel) GetHasWifiBand6() bool`

GetHasWifiBand6 returns the HasWifiBand6 field if non-nil, zero value otherwise.

### GetHasWifiBand6Ok

`func (o *ConstDeviceModel) GetHasWifiBand6Ok() (*bool, bool)`

GetHasWifiBand6Ok returns a tuple with the HasWifiBand6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWifiBand6

`func (o *ConstDeviceModel) SetHasWifiBand6(v bool)`

SetHasWifiBand6 sets HasWifiBand6 field to given value.

### HasHasWifiBand6

`func (o *ConstDeviceModel) HasHasWifiBand6() bool`

HasHasWifiBand6 returns a boolean if a field has been set.

### GetMaxPoeOut

`func (o *ConstDeviceModel) GetMaxPoeOut() int32`

GetMaxPoeOut returns the MaxPoeOut field if non-nil, zero value otherwise.

### GetMaxPoeOutOk

`func (o *ConstDeviceModel) GetMaxPoeOutOk() (*int32, bool)`

GetMaxPoeOutOk returns a tuple with the MaxPoeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoeOut

`func (o *ConstDeviceModel) SetMaxPoeOut(v int32)`

SetMaxPoeOut sets MaxPoeOut field to given value.

### HasMaxPoeOut

`func (o *ConstDeviceModel) HasMaxPoeOut() bool`

HasMaxPoeOut returns a boolean if a field has been set.

### GetMaxWlans

`func (o *ConstDeviceModel) GetMaxWlans() int32`

GetMaxWlans returns the MaxWlans field if non-nil, zero value otherwise.

### GetMaxWlansOk

`func (o *ConstDeviceModel) GetMaxWlansOk() (*int32, bool)`

GetMaxWlansOk returns a tuple with the MaxWlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWlans

`func (o *ConstDeviceModel) SetMaxWlans(v int32)`

SetMaxWlans sets MaxWlans field to given value.

### HasMaxWlans

`func (o *ConstDeviceModel) HasMaxWlans() bool`

HasMaxWlans returns a boolean if a field has been set.

### GetModel

`func (o *ConstDeviceModel) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConstDeviceModel) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConstDeviceModel) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConstDeviceModel) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOtherDfsOk

`func (o *ConstDeviceModel) GetOtherDfsOk() bool`

GetOtherDfsOk returns the OtherDfsOk field if non-nil, zero value otherwise.

### GetOtherDfsOkOk

`func (o *ConstDeviceModel) GetOtherDfsOkOk() (*bool, bool)`

GetOtherDfsOkOk returns a tuple with the OtherDfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherDfsOk

`func (o *ConstDeviceModel) SetOtherDfsOk(v bool)`

SetOtherDfsOk sets OtherDfsOk field to given value.

### HasOtherDfsOk

`func (o *ConstDeviceModel) HasOtherDfsOk() bool`

HasOtherDfsOk returns a boolean if a field has been set.

### GetOutdoor

`func (o *ConstDeviceModel) GetOutdoor() bool`

GetOutdoor returns the Outdoor field if non-nil, zero value otherwise.

### GetOutdoorOk

`func (o *ConstDeviceModel) GetOutdoorOk() (*bool, bool)`

GetOutdoorOk returns a tuple with the Outdoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdoor

`func (o *ConstDeviceModel) SetOutdoor(v bool)`

SetOutdoor sets Outdoor field to given value.

### HasOutdoor

`func (o *ConstDeviceModel) HasOutdoor() bool`

HasOutdoor returns a boolean if a field has been set.

### GetRadios

`func (o *ConstDeviceModel) GetRadios() map[string]string`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *ConstDeviceModel) GetRadiosOk() (*map[string]string, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *ConstDeviceModel) SetRadios(v map[string]string)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *ConstDeviceModel) HasRadios() bool`

HasRadios returns a boolean if a field has been set.

### GetSharedScanningRadio

`func (o *ConstDeviceModel) GetSharedScanningRadio() bool`

GetSharedScanningRadio returns the SharedScanningRadio field if non-nil, zero value otherwise.

### GetSharedScanningRadioOk

`func (o *ConstDeviceModel) GetSharedScanningRadioOk() (*bool, bool)`

GetSharedScanningRadioOk returns a tuple with the SharedScanningRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScanningRadio

`func (o *ConstDeviceModel) SetSharedScanningRadio(v bool)`

SetSharedScanningRadio sets SharedScanningRadio field to given value.

### HasSharedScanningRadio

`func (o *ConstDeviceModel) HasSharedScanningRadio() bool`

HasSharedScanningRadio returns a boolean if a field has been set.

### GetType

`func (o *ConstDeviceModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstDeviceModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstDeviceModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstDeviceModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnmanaged

`func (o *ConstDeviceModel) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *ConstDeviceModel) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *ConstDeviceModel) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *ConstDeviceModel) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetVble

`func (o *ConstDeviceModel) GetVble() ConstDeviceApVble`

GetVble returns the Vble field if non-nil, zero value otherwise.

### GetVbleOk

`func (o *ConstDeviceModel) GetVbleOk() (*ConstDeviceApVble, bool)`

GetVbleOk returns a tuple with the Vble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVble

`func (o *ConstDeviceModel) SetVble(v ConstDeviceApVble)`

SetVble sets Vble field to given value.

### HasVble

`func (o *ConstDeviceModel) HasVble() bool`

HasVble returns a boolean if a field has been set.

### GetAlias

`func (o *ConstDeviceModel) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ConstDeviceModel) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ConstDeviceModel) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ConstDeviceModel) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDefaults

`func (o *ConstDeviceModel) GetDefaults() map[string]string`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *ConstDeviceModel) GetDefaultsOk() (*map[string]string, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *ConstDeviceModel) SetDefaults(v map[string]string)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *ConstDeviceModel) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetEvolvedOs

`func (o *ConstDeviceModel) GetEvolvedOs() bool`

GetEvolvedOs returns the EvolvedOs field if non-nil, zero value otherwise.

### GetEvolvedOsOk

`func (o *ConstDeviceModel) GetEvolvedOsOk() (*bool, bool)`

GetEvolvedOsOk returns a tuple with the EvolvedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvolvedOs

`func (o *ConstDeviceModel) SetEvolvedOs(v bool)`

SetEvolvedOs sets EvolvedOs field to given value.

### HasEvolvedOs

`func (o *ConstDeviceModel) HasEvolvedOs() bool`

HasEvolvedOs returns a boolean if a field has been set.

### GetEvpnRiType

`func (o *ConstDeviceModel) GetEvpnRiType() string`

GetEvpnRiType returns the EvpnRiType field if non-nil, zero value otherwise.

### GetEvpnRiTypeOk

`func (o *ConstDeviceModel) GetEvpnRiTypeOk() (*string, bool)`

GetEvpnRiTypeOk returns a tuple with the EvpnRiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnRiType

`func (o *ConstDeviceModel) SetEvpnRiType(v string)`

SetEvpnRiType sets EvpnRiType field to given value.

### HasEvpnRiType

`func (o *ConstDeviceModel) HasEvpnRiType() bool`

HasEvpnRiType returns a boolean if a field has been set.

### GetExperimental

`func (o *ConstDeviceModel) GetExperimental() bool`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *ConstDeviceModel) GetExperimentalOk() (*bool, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *ConstDeviceModel) SetExperimental(v bool)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *ConstDeviceModel) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### GetFansPluggable

`func (o *ConstDeviceModel) GetFansPluggable() bool`

GetFansPluggable returns the FansPluggable field if non-nil, zero value otherwise.

### GetFansPluggableOk

`func (o *ConstDeviceModel) GetFansPluggableOk() (*bool, bool)`

GetFansPluggableOk returns a tuple with the FansPluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFansPluggable

`func (o *ConstDeviceModel) SetFansPluggable(v bool)`

SetFansPluggable sets FansPluggable field to given value.

### HasFansPluggable

`func (o *ConstDeviceModel) HasFansPluggable() bool`

HasFansPluggable returns a boolean if a field has been set.

### GetHasBgp

`func (o *ConstDeviceModel) GetHasBgp() bool`

GetHasBgp returns the HasBgp field if non-nil, zero value otherwise.

### GetHasBgpOk

`func (o *ConstDeviceModel) GetHasBgpOk() (*bool, bool)`

GetHasBgpOk returns a tuple with the HasBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBgp

`func (o *ConstDeviceModel) SetHasBgp(v bool)`

SetHasBgp sets HasBgp field to given value.

### HasHasBgp

`func (o *ConstDeviceModel) HasHasBgp() bool`

HasHasBgp returns a boolean if a field has been set.

### GetHasEts

`func (o *ConstDeviceModel) GetHasEts() bool`

GetHasEts returns the HasEts field if non-nil, zero value otherwise.

### GetHasEtsOk

`func (o *ConstDeviceModel) GetHasEtsOk() (*bool, bool)`

GetHasEtsOk returns a tuple with the HasEts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEts

`func (o *ConstDeviceModel) SetHasEts(v bool)`

SetHasEts sets HasEts field to given value.

### HasHasEts

`func (o *ConstDeviceModel) HasHasEts() bool`

HasHasEts returns a boolean if a field has been set.

### GetHasEvpn

`func (o *ConstDeviceModel) GetHasEvpn() bool`

GetHasEvpn returns the HasEvpn field if non-nil, zero value otherwise.

### GetHasEvpnOk

`func (o *ConstDeviceModel) GetHasEvpnOk() (*bool, bool)`

GetHasEvpnOk returns a tuple with the HasEvpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEvpn

`func (o *ConstDeviceModel) SetHasEvpn(v bool)`

SetHasEvpn sets HasEvpn field to given value.

### HasHasEvpn

`func (o *ConstDeviceModel) HasHasEvpn() bool`

HasHasEvpn returns a boolean if a field has been set.

### GetHasIrb

`func (o *ConstDeviceModel) GetHasIrb() bool`

GetHasIrb returns the HasIrb field if non-nil, zero value otherwise.

### GetHasIrbOk

`func (o *ConstDeviceModel) GetHasIrbOk() (*bool, bool)`

GetHasIrbOk returns a tuple with the HasIrb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIrb

`func (o *ConstDeviceModel) SetHasIrb(v bool)`

SetHasIrb sets HasIrb field to given value.

### HasHasIrb

`func (o *ConstDeviceModel) HasHasIrb() bool`

HasHasIrb returns a boolean if a field has been set.

### GetHasSnapshot

`func (o *ConstDeviceModel) GetHasSnapshot() bool`

GetHasSnapshot returns the HasSnapshot field if non-nil, zero value otherwise.

### GetHasSnapshotOk

`func (o *ConstDeviceModel) GetHasSnapshotOk() (*bool, bool)`

GetHasSnapshotOk returns a tuple with the HasSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshot

`func (o *ConstDeviceModel) SetHasSnapshot(v bool)`

SetHasSnapshot sets HasSnapshot field to given value.

### HasHasSnapshot

`func (o *ConstDeviceModel) HasHasSnapshot() bool`

HasHasSnapshot returns a boolean if a field has been set.

### GetHasVc

`func (o *ConstDeviceModel) GetHasVc() bool`

GetHasVc returns the HasVc field if non-nil, zero value otherwise.

### GetHasVcOk

`func (o *ConstDeviceModel) GetHasVcOk() (*bool, bool)`

GetHasVcOk returns a tuple with the HasVc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVc

`func (o *ConstDeviceModel) SetHasVc(v bool)`

SetHasVc sets HasVc field to given value.

### HasHasVc

`func (o *ConstDeviceModel) HasHasVc() bool`

HasHasVc returns a boolean if a field has been set.

### GetModular

`func (o *ConstDeviceModel) GetModular() bool`

GetModular returns the Modular field if non-nil, zero value otherwise.

### GetModularOk

`func (o *ConstDeviceModel) GetModularOk() (*bool, bool)`

GetModularOk returns a tuple with the Modular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModular

`func (o *ConstDeviceModel) SetModular(v bool)`

SetModular sets Modular field to given value.

### HasModular

`func (o *ConstDeviceModel) HasModular() bool`

HasModular returns a boolean if a field has been set.

### GetNoShapingRate

`func (o *ConstDeviceModel) GetNoShapingRate() bool`

GetNoShapingRate returns the NoShapingRate field if non-nil, zero value otherwise.

### GetNoShapingRateOk

`func (o *ConstDeviceModel) GetNoShapingRateOk() (*bool, bool)`

GetNoShapingRateOk returns a tuple with the NoShapingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShapingRate

`func (o *ConstDeviceModel) SetNoShapingRate(v bool)`

SetNoShapingRate sets NoShapingRate field to given value.

### HasNoShapingRate

`func (o *ConstDeviceModel) HasNoShapingRate() bool`

HasNoShapingRate returns a boolean if a field has been set.

### GetNumberFans

`func (o *ConstDeviceModel) GetNumberFans() int32`

GetNumberFans returns the NumberFans field if non-nil, zero value otherwise.

### GetNumberFansOk

`func (o *ConstDeviceModel) GetNumberFansOk() (*int32, bool)`

GetNumberFansOk returns a tuple with the NumberFans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFans

`func (o *ConstDeviceModel) SetNumberFans(v int32)`

SetNumberFans sets NumberFans field to given value.

### HasNumberFans

`func (o *ConstDeviceModel) HasNumberFans() bool`

HasNumberFans returns a boolean if a field has been set.

### GetOcDevice

`func (o *ConstDeviceModel) GetOcDevice() bool`

GetOcDevice returns the OcDevice field if non-nil, zero value otherwise.

### GetOcDeviceOk

`func (o *ConstDeviceModel) GetOcDeviceOk() (*bool, bool)`

GetOcDeviceOk returns a tuple with the OcDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcDevice

`func (o *ConstDeviceModel) SetOcDevice(v bool)`

SetOcDevice sets OcDevice field to given value.

### HasOcDevice

`func (o *ConstDeviceModel) HasOcDevice() bool`

HasOcDevice returns a boolean if a field has been set.

### GetOobInterface

`func (o *ConstDeviceModel) GetOobInterface() string`

GetOobInterface returns the OobInterface field if non-nil, zero value otherwise.

### GetOobInterfaceOk

`func (o *ConstDeviceModel) GetOobInterfaceOk() (*string, bool)`

GetOobInterfaceOk returns a tuple with the OobInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobInterface

`func (o *ConstDeviceModel) SetOobInterface(v string)`

SetOobInterface sets OobInterface field to given value.

### HasOobInterface

`func (o *ConstDeviceModel) HasOobInterface() bool`

HasOobInterface returns a boolean if a field has been set.

### GetPacketActionDropOnly

`func (o *ConstDeviceModel) GetPacketActionDropOnly() bool`

GetPacketActionDropOnly returns the PacketActionDropOnly field if non-nil, zero value otherwise.

### GetPacketActionDropOnlyOk

`func (o *ConstDeviceModel) GetPacketActionDropOnlyOk() (*bool, bool)`

GetPacketActionDropOnlyOk returns a tuple with the PacketActionDropOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketActionDropOnly

`func (o *ConstDeviceModel) SetPacketActionDropOnly(v bool)`

SetPacketActionDropOnly sets PacketActionDropOnly field to given value.

### HasPacketActionDropOnly

`func (o *ConstDeviceModel) HasPacketActionDropOnly() bool`

HasPacketActionDropOnly returns a boolean if a field has been set.

### GetPic

`func (o *ConstDeviceModel) GetPic() map[string]string`

GetPic returns the Pic field if non-nil, zero value otherwise.

### GetPicOk

`func (o *ConstDeviceModel) GetPicOk() (*map[string]string, bool)`

GetPicOk returns a tuple with the Pic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPic

`func (o *ConstDeviceModel) SetPic(v map[string]string)`

SetPic sets Pic field to given value.

### HasPic

`func (o *ConstDeviceModel) HasPic() bool`

HasPic returns a boolean if a field has been set.

### GetSubRequired

`func (o *ConstDeviceModel) GetSubRequired() string`

GetSubRequired returns the SubRequired field if non-nil, zero value otherwise.

### GetSubRequiredOk

`func (o *ConstDeviceModel) GetSubRequiredOk() (*string, bool)`

GetSubRequiredOk returns a tuple with the SubRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRequired

`func (o *ConstDeviceModel) SetSubRequired(v string)`

SetSubRequired sets SubRequired field to given value.

### HasSubRequired

`func (o *ConstDeviceModel) HasSubRequired() bool`

HasSubRequired returns a boolean if a field has been set.

### GetHaNode0Fpc

`func (o *ConstDeviceModel) GetHaNode0Fpc() int32`

GetHaNode0Fpc returns the HaNode0Fpc field if non-nil, zero value otherwise.

### GetHaNode0FpcOk

`func (o *ConstDeviceModel) GetHaNode0FpcOk() (*int32, bool)`

GetHaNode0FpcOk returns a tuple with the HaNode0Fpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNode0Fpc

`func (o *ConstDeviceModel) SetHaNode0Fpc(v int32)`

SetHaNode0Fpc sets HaNode0Fpc field to given value.

### HasHaNode0Fpc

`func (o *ConstDeviceModel) HasHaNode0Fpc() bool`

HasHaNode0Fpc returns a boolean if a field has been set.

### GetHaNode1Fpc

`func (o *ConstDeviceModel) GetHaNode1Fpc() int32`

GetHaNode1Fpc returns the HaNode1Fpc field if non-nil, zero value otherwise.

### GetHaNode1FpcOk

`func (o *ConstDeviceModel) GetHaNode1FpcOk() (*int32, bool)`

GetHaNode1FpcOk returns a tuple with the HaNode1Fpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNode1Fpc

`func (o *ConstDeviceModel) SetHaNode1Fpc(v int32)`

SetHaNode1Fpc sets HaNode1Fpc field to given value.

### HasHaNode1Fpc

`func (o *ConstDeviceModel) HasHaNode1Fpc() bool`

HasHaNode1Fpc returns a boolean if a field has been set.

### GetHasFxp0

`func (o *ConstDeviceModel) GetHasFxp0() bool`

GetHasFxp0 returns the HasFxp0 field if non-nil, zero value otherwise.

### GetHasFxp0Ok

`func (o *ConstDeviceModel) GetHasFxp0Ok() (*bool, bool)`

GetHasFxp0Ok returns a tuple with the HasFxp0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFxp0

`func (o *ConstDeviceModel) SetHasFxp0(v bool)`

SetHasFxp0 sets HasFxp0 field to given value.

### HasHasFxp0

`func (o *ConstDeviceModel) HasHasFxp0() bool`

HasHasFxp0 returns a boolean if a field has been set.

### GetHasHaControl

`func (o *ConstDeviceModel) GetHasHaControl() bool`

GetHasHaControl returns the HasHaControl field if non-nil, zero value otherwise.

### GetHasHaControlOk

`func (o *ConstDeviceModel) GetHasHaControlOk() (*bool, bool)`

GetHasHaControlOk returns a tuple with the HasHaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHaControl

`func (o *ConstDeviceModel) SetHasHaControl(v bool)`

SetHasHaControl sets HasHaControl field to given value.

### HasHasHaControl

`func (o *ConstDeviceModel) HasHasHaControl() bool`

HasHasHaControl returns a boolean if a field has been set.

### GetHasHaData

`func (o *ConstDeviceModel) GetHasHaData() bool`

GetHasHaData returns the HasHaData field if non-nil, zero value otherwise.

### GetHasHaDataOk

`func (o *ConstDeviceModel) GetHasHaDataOk() (*bool, bool)`

GetHasHaDataOk returns a tuple with the HasHaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHaData

`func (o *ConstDeviceModel) SetHasHaData(v bool)`

SetHasHaData sets HasHaData field to given value.

### HasHasHaData

`func (o *ConstDeviceModel) HasHasHaData() bool`

HasHasHaData returns a boolean if a field has been set.

### GetIrbDisabledByDefault

`func (o *ConstDeviceModel) GetIrbDisabledByDefault() bool`

GetIrbDisabledByDefault returns the IrbDisabledByDefault field if non-nil, zero value otherwise.

### GetIrbDisabledByDefaultOk

`func (o *ConstDeviceModel) GetIrbDisabledByDefaultOk() (*bool, bool)`

GetIrbDisabledByDefaultOk returns a tuple with the IrbDisabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrbDisabledByDefault

`func (o *ConstDeviceModel) SetIrbDisabledByDefault(v bool)`

SetIrbDisabledByDefault sets IrbDisabledByDefault field to given value.

### HasIrbDisabledByDefault

`func (o *ConstDeviceModel) HasIrbDisabledByDefault() bool`

HasIrbDisabledByDefault returns a boolean if a field has been set.

### GetPorts

`func (o *ConstDeviceModel) GetPorts() ConstDeviceGatewayPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ConstDeviceModel) GetPortsOk() (*ConstDeviceGatewayPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ConstDeviceModel) SetPorts(v ConstDeviceGatewayPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ConstDeviceModel) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetT128Device

`func (o *ConstDeviceModel) GetT128Device() bool`

GetT128Device returns the T128Device field if non-nil, zero value otherwise.

### GetT128DeviceOk

`func (o *ConstDeviceModel) GetT128DeviceOk() (*bool, bool)`

GetT128DeviceOk returns a tuple with the T128Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT128Device

`func (o *ConstDeviceModel) SetT128Device(v bool)`

SetT128Device sets T128Device field to given value.

### HasT128Device

`func (o *ConstDeviceModel) HasT128Device() bool`

HasT128Device returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


