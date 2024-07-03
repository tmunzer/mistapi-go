# SdkstatsWirelessClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LastSeen** | **float32** | last seen timestamp | 
**MapId** | Pointer to **NullableString** | map_id of the sdk client (if known), or null | [optional] 
**Name** | Pointer to **string** | name of the sdk client (if provided) | [optional] 
**NetworkConnection** | Pointer to [**SdkclientStatNetworkConnection**](SdkclientStatNetworkConnection.md) |  | [optional] 
**Uuid** | **string** | uuid of the sdk client | 
**Vbeacons** | Pointer to [**[]SdkclientWirelessStatsVbeacon**](SdkclientWirelessStatsVbeacon.md) | list of beacon_id’s of the sdk client is in and since when (if known) | [optional] 
**X** | Pointer to **float32** | x (in pixels) of user location on the map (if known) | [optional] 
**Y** | Pointer to **float32** | y (in pixels) of user location on the map (if known) | [optional] 
**Zones** | Pointer to [**[]SdkclientWirelessStatsZone**](SdkclientWirelessStatsZone.md) | list of zone_id’s of the sdk client is in and since when (if known) | [optional] 

## Methods

### NewSdkstatsWirelessClient

`func NewSdkstatsWirelessClient(id string, lastSeen float32, uuid string, ) *SdkstatsWirelessClient`

NewSdkstatsWirelessClient instantiates a new SdkstatsWirelessClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkstatsWirelessClientWithDefaults

`func NewSdkstatsWirelessClientWithDefaults() *SdkstatsWirelessClient`

NewSdkstatsWirelessClientWithDefaults instantiates a new SdkstatsWirelessClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SdkstatsWirelessClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdkstatsWirelessClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdkstatsWirelessClient) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeen

`func (o *SdkstatsWirelessClient) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *SdkstatsWirelessClient) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *SdkstatsWirelessClient) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetMapId

`func (o *SdkstatsWirelessClient) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *SdkstatsWirelessClient) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *SdkstatsWirelessClient) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *SdkstatsWirelessClient) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### SetMapIdNil

`func (o *SdkstatsWirelessClient) SetMapIdNil(b bool)`

 SetMapIdNil sets the value for MapId to be an explicit nil

### UnsetMapId
`func (o *SdkstatsWirelessClient) UnsetMapId()`

UnsetMapId ensures that no value is present for MapId, not even an explicit nil
### GetName

`func (o *SdkstatsWirelessClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkstatsWirelessClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkstatsWirelessClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkstatsWirelessClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConnection

`func (o *SdkstatsWirelessClient) GetNetworkConnection() SdkclientStatNetworkConnection`

GetNetworkConnection returns the NetworkConnection field if non-nil, zero value otherwise.

### GetNetworkConnectionOk

`func (o *SdkstatsWirelessClient) GetNetworkConnectionOk() (*SdkclientStatNetworkConnection, bool)`

GetNetworkConnectionOk returns a tuple with the NetworkConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnection

`func (o *SdkstatsWirelessClient) SetNetworkConnection(v SdkclientStatNetworkConnection)`

SetNetworkConnection sets NetworkConnection field to given value.

### HasNetworkConnection

`func (o *SdkstatsWirelessClient) HasNetworkConnection() bool`

HasNetworkConnection returns a boolean if a field has been set.

### GetUuid

`func (o *SdkstatsWirelessClient) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SdkstatsWirelessClient) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SdkstatsWirelessClient) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetVbeacons

`func (o *SdkstatsWirelessClient) GetVbeacons() []SdkclientWirelessStatsVbeacon`

GetVbeacons returns the Vbeacons field if non-nil, zero value otherwise.

### GetVbeaconsOk

`func (o *SdkstatsWirelessClient) GetVbeaconsOk() (*[]SdkclientWirelessStatsVbeacon, bool)`

GetVbeaconsOk returns a tuple with the Vbeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVbeacons

`func (o *SdkstatsWirelessClient) SetVbeacons(v []SdkclientWirelessStatsVbeacon)`

SetVbeacons sets Vbeacons field to given value.

### HasVbeacons

`func (o *SdkstatsWirelessClient) HasVbeacons() bool`

HasVbeacons returns a boolean if a field has been set.

### GetX

`func (o *SdkstatsWirelessClient) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SdkstatsWirelessClient) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SdkstatsWirelessClient) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *SdkstatsWirelessClient) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *SdkstatsWirelessClient) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *SdkstatsWirelessClient) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *SdkstatsWirelessClient) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *SdkstatsWirelessClient) HasY() bool`

HasY returns a boolean if a field has been set.

### GetZones

`func (o *SdkstatsWirelessClient) GetZones() []SdkclientWirelessStatsZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *SdkstatsWirelessClient) GetZonesOk() (*[]SdkclientWirelessStatsZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *SdkstatsWirelessClient) SetZones(v []SdkclientWirelessStatsZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *SdkstatsWirelessClient) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


