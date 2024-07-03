# UnconnectedClientStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** | mac address of the AP that heard the client | 
**LastSeen** | **float32** | last seen timestamp | 
**Mac** | **string** | mac address of the (unconnected) client | 
**Manufacture** | **string** | device manufacture, through fingerprinting or OUI | 
**MapId** | Pointer to **NullableString** | map_id of the client (if known), or null | [optional] 
**Rssi** | **int32** | client RSSI observered by the AP that heard the client (in dBm) | 
**X** | Pointer to **int32** | x (in pixels) of user location on the map (if known) | [optional] 
**Y** | **int32** | y (in pixels) of user location on the map (if known) | 

## Methods

### NewUnconnectedClientStat

`func NewUnconnectedClientStat(apMac string, lastSeen float32, mac string, manufacture string, rssi int32, y int32, ) *UnconnectedClientStat`

NewUnconnectedClientStat instantiates a new UnconnectedClientStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconnectedClientStatWithDefaults

`func NewUnconnectedClientStatWithDefaults() *UnconnectedClientStat`

NewUnconnectedClientStatWithDefaults instantiates a new UnconnectedClientStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *UnconnectedClientStat) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *UnconnectedClientStat) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *UnconnectedClientStat) SetApMac(v string)`

SetApMac sets ApMac field to given value.


### GetLastSeen

`func (o *UnconnectedClientStat) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *UnconnectedClientStat) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *UnconnectedClientStat) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetMac

`func (o *UnconnectedClientStat) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *UnconnectedClientStat) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *UnconnectedClientStat) SetMac(v string)`

SetMac sets Mac field to given value.


### GetManufacture

`func (o *UnconnectedClientStat) GetManufacture() string`

GetManufacture returns the Manufacture field if non-nil, zero value otherwise.

### GetManufactureOk

`func (o *UnconnectedClientStat) GetManufactureOk() (*string, bool)`

GetManufactureOk returns a tuple with the Manufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacture

`func (o *UnconnectedClientStat) SetManufacture(v string)`

SetManufacture sets Manufacture field to given value.


### GetMapId

`func (o *UnconnectedClientStat) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *UnconnectedClientStat) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *UnconnectedClientStat) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *UnconnectedClientStat) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### SetMapIdNil

`func (o *UnconnectedClientStat) SetMapIdNil(b bool)`

 SetMapIdNil sets the value for MapId to be an explicit nil

### UnsetMapId
`func (o *UnconnectedClientStat) UnsetMapId()`

UnsetMapId ensures that no value is present for MapId, not even an explicit nil
### GetRssi

`func (o *UnconnectedClientStat) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *UnconnectedClientStat) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *UnconnectedClientStat) SetRssi(v int32)`

SetRssi sets Rssi field to given value.


### GetX

`func (o *UnconnectedClientStat) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *UnconnectedClientStat) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *UnconnectedClientStat) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *UnconnectedClientStat) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *UnconnectedClientStat) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *UnconnectedClientStat) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *UnconnectedClientStat) SetY(v int32)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


