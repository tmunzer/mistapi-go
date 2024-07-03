# ConstDeviceAp

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
**HasPoeOut** | Pointer to **bool** |  | [optional] 
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
**Type** | Pointer to **string** |  | [optional] [default to "ap"]
**Unmanaged** | Pointer to **bool** |  | [optional] 
**Vble** | Pointer to [**ConstDeviceApVble**](ConstDeviceApVble.md) |  | [optional] 

## Methods

### NewConstDeviceAp

`func NewConstDeviceAp() *ConstDeviceAp`

NewConstDeviceAp instantiates a new ConstDeviceAp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstDeviceApWithDefaults

`func NewConstDeviceApWithDefaults() *ConstDeviceAp`

NewConstDeviceApWithDefaults instantiates a new ConstDeviceAp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApType

`func (o *ConstDeviceAp) GetApType() string`

GetApType returns the ApType field if non-nil, zero value otherwise.

### GetApTypeOk

`func (o *ConstDeviceAp) GetApTypeOk() (*string, bool)`

GetApTypeOk returns a tuple with the ApType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApType

`func (o *ConstDeviceAp) SetApType(v string)`

SetApType sets ApType field to given value.

### HasApType

`func (o *ConstDeviceAp) HasApType() bool`

HasApType returns a boolean if a field has been set.

### GetBand24

`func (o *ConstDeviceAp) GetBand24() ConstDeviceApBand24`

GetBand24 returns the Band24 field if non-nil, zero value otherwise.

### GetBand24Ok

`func (o *ConstDeviceAp) GetBand24Ok() (*ConstDeviceApBand24, bool)`

GetBand24Ok returns a tuple with the Band24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24

`func (o *ConstDeviceAp) SetBand24(v ConstDeviceApBand24)`

SetBand24 sets Band24 field to given value.

### HasBand24

`func (o *ConstDeviceAp) HasBand24() bool`

HasBand24 returns a boolean if a field has been set.

### GetBand5

`func (o *ConstDeviceAp) GetBand5() ConstDeviceApBand5`

GetBand5 returns the Band5 field if non-nil, zero value otherwise.

### GetBand5Ok

`func (o *ConstDeviceAp) GetBand5Ok() (*ConstDeviceApBand5, bool)`

GetBand5Ok returns a tuple with the Band5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5

`func (o *ConstDeviceAp) SetBand5(v ConstDeviceApBand5)`

SetBand5 sets Band5 field to given value.

### HasBand5

`func (o *ConstDeviceAp) HasBand5() bool`

HasBand5 returns a boolean if a field has been set.

### GetBand6

`func (o *ConstDeviceAp) GetBand6() ConstDeviceApBand6`

GetBand6 returns the Band6 field if non-nil, zero value otherwise.

### GetBand6Ok

`func (o *ConstDeviceAp) GetBand6Ok() (*ConstDeviceApBand6, bool)`

GetBand6Ok returns a tuple with the Band6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6

`func (o *ConstDeviceAp) SetBand6(v ConstDeviceApBand6)`

SetBand6 sets Band6 field to given value.

### HasBand6

`func (o *ConstDeviceAp) HasBand6() bool`

HasBand6 returns a boolean if a field has been set.

### GetBand24Usages

`func (o *ConstDeviceAp) GetBand24Usages() RadioBand24Usage`

GetBand24Usages returns the Band24Usages field if non-nil, zero value otherwise.

### GetBand24UsagesOk

`func (o *ConstDeviceAp) GetBand24UsagesOk() (*RadioBand24Usage, bool)`

GetBand24UsagesOk returns a tuple with the Band24Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Usages

`func (o *ConstDeviceAp) SetBand24Usages(v RadioBand24Usage)`

SetBand24Usages sets Band24Usages field to given value.

### HasBand24Usages

`func (o *ConstDeviceAp) HasBand24Usages() bool`

HasBand24Usages returns a boolean if a field has been set.

### GetCeDfsOk

`func (o *ConstDeviceAp) GetCeDfsOk() bool`

GetCeDfsOk returns the CeDfsOk field if non-nil, zero value otherwise.

### GetCeDfsOkOk

`func (o *ConstDeviceAp) GetCeDfsOkOk() (*bool, bool)`

GetCeDfsOkOk returns a tuple with the CeDfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeDfsOk

`func (o *ConstDeviceAp) SetCeDfsOk(v bool)`

SetCeDfsOk sets CeDfsOk field to given value.

### HasCeDfsOk

`func (o *ConstDeviceAp) HasCeDfsOk() bool`

HasCeDfsOk returns a boolean if a field has been set.

### GetCiscoPace

`func (o *ConstDeviceAp) GetCiscoPace() bool`

GetCiscoPace returns the CiscoPace field if non-nil, zero value otherwise.

### GetCiscoPaceOk

`func (o *ConstDeviceAp) GetCiscoPaceOk() (*bool, bool)`

GetCiscoPaceOk returns a tuple with the CiscoPace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoPace

`func (o *ConstDeviceAp) SetCiscoPace(v bool)`

SetCiscoPace sets CiscoPace field to given value.

### HasCiscoPace

`func (o *ConstDeviceAp) HasCiscoPace() bool`

HasCiscoPace returns a boolean if a field has been set.

### GetDescription

`func (o *ConstDeviceAp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstDeviceAp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstDeviceAp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstDeviceAp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisallowedChannels

`func (o *ConstDeviceAp) GetDisallowedChannels() map[string]map[string][]int32`

GetDisallowedChannels returns the DisallowedChannels field if non-nil, zero value otherwise.

### GetDisallowedChannelsOk

`func (o *ConstDeviceAp) GetDisallowedChannelsOk() (*map[string]map[string][]int32, bool)`

GetDisallowedChannelsOk returns a tuple with the DisallowedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedChannels

`func (o *ConstDeviceAp) SetDisallowedChannels(v map[string]map[string][]int32)`

SetDisallowedChannels sets DisallowedChannels field to given value.

### HasDisallowedChannels

`func (o *ConstDeviceAp) HasDisallowedChannels() bool`

HasDisallowedChannels returns a boolean if a field has been set.

### GetDisplay

`func (o *ConstDeviceAp) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConstDeviceAp) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConstDeviceAp) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConstDeviceAp) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetExtio

`func (o *ConstDeviceAp) GetExtio() map[string]ConstDeviceApExtios`

GetExtio returns the Extio field if non-nil, zero value otherwise.

### GetExtioOk

`func (o *ConstDeviceAp) GetExtioOk() (*map[string]ConstDeviceApExtios, bool)`

GetExtioOk returns a tuple with the Extio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtio

`func (o *ConstDeviceAp) SetExtio(v map[string]ConstDeviceApExtios)`

SetExtio sets Extio field to given value.

### HasExtio

`func (o *ConstDeviceAp) HasExtio() bool`

HasExtio returns a boolean if a field has been set.

### GetFccDfsOk

`func (o *ConstDeviceAp) GetFccDfsOk() bool`

GetFccDfsOk returns the FccDfsOk field if non-nil, zero value otherwise.

### GetFccDfsOkOk

`func (o *ConstDeviceAp) GetFccDfsOkOk() (*bool, bool)`

GetFccDfsOkOk returns a tuple with the FccDfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFccDfsOk

`func (o *ConstDeviceAp) SetFccDfsOk(v bool)`

SetFccDfsOk sets FccDfsOk field to given value.

### HasFccDfsOk

`func (o *ConstDeviceAp) HasFccDfsOk() bool`

HasFccDfsOk returns a boolean if a field has been set.

### GetHas11ax

`func (o *ConstDeviceAp) GetHas11ax() bool`

GetHas11ax returns the Has11ax field if non-nil, zero value otherwise.

### GetHas11axOk

`func (o *ConstDeviceAp) GetHas11axOk() (*bool, bool)`

GetHas11axOk returns a tuple with the Has11ax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas11ax

`func (o *ConstDeviceAp) SetHas11ax(v bool)`

SetHas11ax sets Has11ax field to given value.

### HasHas11ax

`func (o *ConstDeviceAp) HasHas11ax() bool`

HasHas11ax returns a boolean if a field has been set.

### GetHasCompass

`func (o *ConstDeviceAp) GetHasCompass() bool`

GetHasCompass returns the HasCompass field if non-nil, zero value otherwise.

### GetHasCompassOk

`func (o *ConstDeviceAp) GetHasCompassOk() (*bool, bool)`

GetHasCompassOk returns a tuple with the HasCompass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompass

`func (o *ConstDeviceAp) SetHasCompass(v bool)`

SetHasCompass sets HasCompass field to given value.

### HasHasCompass

`func (o *ConstDeviceAp) HasHasCompass() bool`

HasHasCompass returns a boolean if a field has been set.

### GetHasExtAnt

`func (o *ConstDeviceAp) GetHasExtAnt() bool`

GetHasExtAnt returns the HasExtAnt field if non-nil, zero value otherwise.

### GetHasExtAntOk

`func (o *ConstDeviceAp) GetHasExtAntOk() (*bool, bool)`

GetHasExtAntOk returns a tuple with the HasExtAnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExtAnt

`func (o *ConstDeviceAp) SetHasExtAnt(v bool)`

SetHasExtAnt sets HasExtAnt field to given value.

### HasHasExtAnt

`func (o *ConstDeviceAp) HasHasExtAnt() bool`

HasHasExtAnt returns a boolean if a field has been set.

### GetHasExtio

`func (o *ConstDeviceAp) GetHasExtio() bool`

GetHasExtio returns the HasExtio field if non-nil, zero value otherwise.

### GetHasExtioOk

`func (o *ConstDeviceAp) GetHasExtioOk() (*bool, bool)`

GetHasExtioOk returns a tuple with the HasExtio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExtio

`func (o *ConstDeviceAp) SetHasExtio(v bool)`

SetHasExtio sets HasExtio field to given value.

### HasHasExtio

`func (o *ConstDeviceAp) HasHasExtio() bool`

HasHasExtio returns a boolean if a field has been set.

### GetHasHeight

`func (o *ConstDeviceAp) GetHasHeight() bool`

GetHasHeight returns the HasHeight field if non-nil, zero value otherwise.

### GetHasHeightOk

`func (o *ConstDeviceAp) GetHasHeightOk() (*bool, bool)`

GetHasHeightOk returns a tuple with the HasHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHeight

`func (o *ConstDeviceAp) SetHasHeight(v bool)`

SetHasHeight sets HasHeight field to given value.

### HasHasHeight

`func (o *ConstDeviceAp) HasHasHeight() bool`

HasHasHeight returns a boolean if a field has been set.

### GetHasModulePort

`func (o *ConstDeviceAp) GetHasModulePort() bool`

GetHasModulePort returns the HasModulePort field if non-nil, zero value otherwise.

### GetHasModulePortOk

`func (o *ConstDeviceAp) GetHasModulePortOk() (*bool, bool)`

GetHasModulePortOk returns a tuple with the HasModulePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasModulePort

`func (o *ConstDeviceAp) SetHasModulePort(v bool)`

SetHasModulePort sets HasModulePort field to given value.

### HasHasModulePort

`func (o *ConstDeviceAp) HasHasModulePort() bool`

HasHasModulePort returns a boolean if a field has been set.

### GetHasPoeOut

`func (o *ConstDeviceAp) GetHasPoeOut() bool`

GetHasPoeOut returns the HasPoeOut field if non-nil, zero value otherwise.

### GetHasPoeOutOk

`func (o *ConstDeviceAp) GetHasPoeOutOk() (*bool, bool)`

GetHasPoeOutOk returns a tuple with the HasPoeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPoeOut

`func (o *ConstDeviceAp) SetHasPoeOut(v bool)`

SetHasPoeOut sets HasPoeOut field to given value.

### HasHasPoeOut

`func (o *ConstDeviceAp) HasHasPoeOut() bool`

HasHasPoeOut returns a boolean if a field has been set.

### GetHasScanningRadio

`func (o *ConstDeviceAp) GetHasScanningRadio() bool`

GetHasScanningRadio returns the HasScanningRadio field if non-nil, zero value otherwise.

### GetHasScanningRadioOk

`func (o *ConstDeviceAp) GetHasScanningRadioOk() (*bool, bool)`

GetHasScanningRadioOk returns a tuple with the HasScanningRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScanningRadio

`func (o *ConstDeviceAp) SetHasScanningRadio(v bool)`

SetHasScanningRadio sets HasScanningRadio field to given value.

### HasHasScanningRadio

`func (o *ConstDeviceAp) HasHasScanningRadio() bool`

HasHasScanningRadio returns a boolean if a field has been set.

### GetHasSelectableRadio

`func (o *ConstDeviceAp) GetHasSelectableRadio() bool`

GetHasSelectableRadio returns the HasSelectableRadio field if non-nil, zero value otherwise.

### GetHasSelectableRadioOk

`func (o *ConstDeviceAp) GetHasSelectableRadioOk() (*bool, bool)`

GetHasSelectableRadioOk returns a tuple with the HasSelectableRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectableRadio

`func (o *ConstDeviceAp) SetHasSelectableRadio(v bool)`

SetHasSelectableRadio sets HasSelectableRadio field to given value.

### HasHasSelectableRadio

`func (o *ConstDeviceAp) HasHasSelectableRadio() bool`

HasHasSelectableRadio returns a boolean if a field has been set.

### GetHasUsb

`func (o *ConstDeviceAp) GetHasUsb() bool`

GetHasUsb returns the HasUsb field if non-nil, zero value otherwise.

### GetHasUsbOk

`func (o *ConstDeviceAp) GetHasUsbOk() (*bool, bool)`

GetHasUsbOk returns a tuple with the HasUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsb

`func (o *ConstDeviceAp) SetHasUsb(v bool)`

SetHasUsb sets HasUsb field to given value.

### HasHasUsb

`func (o *ConstDeviceAp) HasHasUsb() bool`

HasHasUsb returns a boolean if a field has been set.

### GetHasVble

`func (o *ConstDeviceAp) GetHasVble() bool`

GetHasVble returns the HasVble field if non-nil, zero value otherwise.

### GetHasVbleOk

`func (o *ConstDeviceAp) GetHasVbleOk() (*bool, bool)`

GetHasVbleOk returns a tuple with the HasVble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVble

`func (o *ConstDeviceAp) SetHasVble(v bool)`

SetHasVble sets HasVble field to given value.

### HasHasVble

`func (o *ConstDeviceAp) HasHasVble() bool`

HasHasVble returns a boolean if a field has been set.

### GetHasWifiBand24

`func (o *ConstDeviceAp) GetHasWifiBand24() bool`

GetHasWifiBand24 returns the HasWifiBand24 field if non-nil, zero value otherwise.

### GetHasWifiBand24Ok

`func (o *ConstDeviceAp) GetHasWifiBand24Ok() (*bool, bool)`

GetHasWifiBand24Ok returns a tuple with the HasWifiBand24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWifiBand24

`func (o *ConstDeviceAp) SetHasWifiBand24(v bool)`

SetHasWifiBand24 sets HasWifiBand24 field to given value.

### HasHasWifiBand24

`func (o *ConstDeviceAp) HasHasWifiBand24() bool`

HasHasWifiBand24 returns a boolean if a field has been set.

### GetHasWifiBand5

`func (o *ConstDeviceAp) GetHasWifiBand5() bool`

GetHasWifiBand5 returns the HasWifiBand5 field if non-nil, zero value otherwise.

### GetHasWifiBand5Ok

`func (o *ConstDeviceAp) GetHasWifiBand5Ok() (*bool, bool)`

GetHasWifiBand5Ok returns a tuple with the HasWifiBand5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWifiBand5

`func (o *ConstDeviceAp) SetHasWifiBand5(v bool)`

SetHasWifiBand5 sets HasWifiBand5 field to given value.

### HasHasWifiBand5

`func (o *ConstDeviceAp) HasHasWifiBand5() bool`

HasHasWifiBand5 returns a boolean if a field has been set.

### GetHasWifiBand6

`func (o *ConstDeviceAp) GetHasWifiBand6() bool`

GetHasWifiBand6 returns the HasWifiBand6 field if non-nil, zero value otherwise.

### GetHasWifiBand6Ok

`func (o *ConstDeviceAp) GetHasWifiBand6Ok() (*bool, bool)`

GetHasWifiBand6Ok returns a tuple with the HasWifiBand6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWifiBand6

`func (o *ConstDeviceAp) SetHasWifiBand6(v bool)`

SetHasWifiBand6 sets HasWifiBand6 field to given value.

### HasHasWifiBand6

`func (o *ConstDeviceAp) HasHasWifiBand6() bool`

HasHasWifiBand6 returns a boolean if a field has been set.

### GetMaxPoeOut

`func (o *ConstDeviceAp) GetMaxPoeOut() int32`

GetMaxPoeOut returns the MaxPoeOut field if non-nil, zero value otherwise.

### GetMaxPoeOutOk

`func (o *ConstDeviceAp) GetMaxPoeOutOk() (*int32, bool)`

GetMaxPoeOutOk returns a tuple with the MaxPoeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoeOut

`func (o *ConstDeviceAp) SetMaxPoeOut(v int32)`

SetMaxPoeOut sets MaxPoeOut field to given value.

### HasMaxPoeOut

`func (o *ConstDeviceAp) HasMaxPoeOut() bool`

HasMaxPoeOut returns a boolean if a field has been set.

### GetMaxWlans

`func (o *ConstDeviceAp) GetMaxWlans() int32`

GetMaxWlans returns the MaxWlans field if non-nil, zero value otherwise.

### GetMaxWlansOk

`func (o *ConstDeviceAp) GetMaxWlansOk() (*int32, bool)`

GetMaxWlansOk returns a tuple with the MaxWlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWlans

`func (o *ConstDeviceAp) SetMaxWlans(v int32)`

SetMaxWlans sets MaxWlans field to given value.

### HasMaxWlans

`func (o *ConstDeviceAp) HasMaxWlans() bool`

HasMaxWlans returns a boolean if a field has been set.

### GetModel

`func (o *ConstDeviceAp) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConstDeviceAp) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConstDeviceAp) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConstDeviceAp) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOtherDfsOk

`func (o *ConstDeviceAp) GetOtherDfsOk() bool`

GetOtherDfsOk returns the OtherDfsOk field if non-nil, zero value otherwise.

### GetOtherDfsOkOk

`func (o *ConstDeviceAp) GetOtherDfsOkOk() (*bool, bool)`

GetOtherDfsOkOk returns a tuple with the OtherDfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherDfsOk

`func (o *ConstDeviceAp) SetOtherDfsOk(v bool)`

SetOtherDfsOk sets OtherDfsOk field to given value.

### HasOtherDfsOk

`func (o *ConstDeviceAp) HasOtherDfsOk() bool`

HasOtherDfsOk returns a boolean if a field has been set.

### GetOutdoor

`func (o *ConstDeviceAp) GetOutdoor() bool`

GetOutdoor returns the Outdoor field if non-nil, zero value otherwise.

### GetOutdoorOk

`func (o *ConstDeviceAp) GetOutdoorOk() (*bool, bool)`

GetOutdoorOk returns a tuple with the Outdoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdoor

`func (o *ConstDeviceAp) SetOutdoor(v bool)`

SetOutdoor sets Outdoor field to given value.

### HasOutdoor

`func (o *ConstDeviceAp) HasOutdoor() bool`

HasOutdoor returns a boolean if a field has been set.

### GetRadios

`func (o *ConstDeviceAp) GetRadios() map[string]string`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *ConstDeviceAp) GetRadiosOk() (*map[string]string, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *ConstDeviceAp) SetRadios(v map[string]string)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *ConstDeviceAp) HasRadios() bool`

HasRadios returns a boolean if a field has been set.

### GetSharedScanningRadio

`func (o *ConstDeviceAp) GetSharedScanningRadio() bool`

GetSharedScanningRadio returns the SharedScanningRadio field if non-nil, zero value otherwise.

### GetSharedScanningRadioOk

`func (o *ConstDeviceAp) GetSharedScanningRadioOk() (*bool, bool)`

GetSharedScanningRadioOk returns a tuple with the SharedScanningRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScanningRadio

`func (o *ConstDeviceAp) SetSharedScanningRadio(v bool)`

SetSharedScanningRadio sets SharedScanningRadio field to given value.

### HasSharedScanningRadio

`func (o *ConstDeviceAp) HasSharedScanningRadio() bool`

HasSharedScanningRadio returns a boolean if a field has been set.

### GetType

`func (o *ConstDeviceAp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstDeviceAp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstDeviceAp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstDeviceAp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnmanaged

`func (o *ConstDeviceAp) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *ConstDeviceAp) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *ConstDeviceAp) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *ConstDeviceAp) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetVble

`func (o *ConstDeviceAp) GetVble() ConstDeviceApVble`

GetVble returns the Vble field if non-nil, zero value otherwise.

### GetVbleOk

`func (o *ConstDeviceAp) GetVbleOk() (*ConstDeviceApVble, bool)`

GetVbleOk returns a tuple with the Vble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVble

`func (o *ConstDeviceAp) SetVble(v ConstDeviceApVble)`

SetVble sets Vble field to given value.

### HasVble

`func (o *ConstDeviceAp) HasVble() bool`

HasVble returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


