# BeaconStatsItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatteryVoltage** | Pointer to **float32** | battery voltage, in mV | [optional] 
**EddystoneInstance** | Pointer to **string** |  | [optional] 
**EddystoneNamespace** | Pointer to **string** |  | [optional] 
**LastSeen** | **float32** |  | 
**Mac** | **string** |  | 
**MapId** | **string** |  | 
**Name** | **string** |  | 
**Power** | **int32** |  | 
**Type** | **string** |  | 
**X** | **float32** |  | 
**Y** | **float32** |  | 

## Methods

### NewBeaconStatsItems

`func NewBeaconStatsItems(lastSeen float32, mac string, mapId string, name string, power int32, type_ string, x float32, y float32, ) *BeaconStatsItems`

NewBeaconStatsItems instantiates a new BeaconStatsItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeaconStatsItemsWithDefaults

`func NewBeaconStatsItemsWithDefaults() *BeaconStatsItems`

NewBeaconStatsItemsWithDefaults instantiates a new BeaconStatsItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatteryVoltage

`func (o *BeaconStatsItems) GetBatteryVoltage() float32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *BeaconStatsItems) GetBatteryVoltageOk() (*float32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *BeaconStatsItems) SetBatteryVoltage(v float32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *BeaconStatsItems) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### GetEddystoneInstance

`func (o *BeaconStatsItems) GetEddystoneInstance() string`

GetEddystoneInstance returns the EddystoneInstance field if non-nil, zero value otherwise.

### GetEddystoneInstanceOk

`func (o *BeaconStatsItems) GetEddystoneInstanceOk() (*string, bool)`

GetEddystoneInstanceOk returns a tuple with the EddystoneInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneInstance

`func (o *BeaconStatsItems) SetEddystoneInstance(v string)`

SetEddystoneInstance sets EddystoneInstance field to given value.

### HasEddystoneInstance

`func (o *BeaconStatsItems) HasEddystoneInstance() bool`

HasEddystoneInstance returns a boolean if a field has been set.

### GetEddystoneNamespace

`func (o *BeaconStatsItems) GetEddystoneNamespace() string`

GetEddystoneNamespace returns the EddystoneNamespace field if non-nil, zero value otherwise.

### GetEddystoneNamespaceOk

`func (o *BeaconStatsItems) GetEddystoneNamespaceOk() (*string, bool)`

GetEddystoneNamespaceOk returns a tuple with the EddystoneNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneNamespace

`func (o *BeaconStatsItems) SetEddystoneNamespace(v string)`

SetEddystoneNamespace sets EddystoneNamespace field to given value.

### HasEddystoneNamespace

`func (o *BeaconStatsItems) HasEddystoneNamespace() bool`

HasEddystoneNamespace returns a boolean if a field has been set.

### GetLastSeen

`func (o *BeaconStatsItems) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *BeaconStatsItems) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *BeaconStatsItems) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetMac

`func (o *BeaconStatsItems) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BeaconStatsItems) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BeaconStatsItems) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *BeaconStatsItems) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *BeaconStatsItems) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *BeaconStatsItems) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetName

`func (o *BeaconStatsItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BeaconStatsItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BeaconStatsItems) SetName(v string)`

SetName sets Name field to given value.


### GetPower

`func (o *BeaconStatsItems) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *BeaconStatsItems) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *BeaconStatsItems) SetPower(v int32)`

SetPower sets Power field to given value.


### GetType

`func (o *BeaconStatsItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BeaconStatsItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BeaconStatsItems) SetType(v string)`

SetType sets Type field to given value.


### GetX

`func (o *BeaconStatsItems) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *BeaconStatsItems) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *BeaconStatsItems) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *BeaconStatsItems) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *BeaconStatsItems) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *BeaconStatsItems) SetY(v float32)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


