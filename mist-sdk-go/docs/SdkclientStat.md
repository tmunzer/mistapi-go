# SdkclientStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LastSeen** | **float32** | last seen timestamp | 
**MapId** | Pointer to **NullableString** | map_id of the sdk client (if known), or null | [optional] 
**Name** | Pointer to **string** | name of the sdk client (if provided) | [optional] 
**NetworkConnection** | [**SdkclientStatNetworkConnection**](SdkclientStatNetworkConnection.md) |  | 
**Uuid** | **string** | uuid of the sdk client | 
**X** | Pointer to **float32** | x (in pixels) of user location on the map (if known) | [optional] 
**Y** | Pointer to **float32** | y (in pixels) of user location on the map (if known) | [optional] 

## Methods

### NewSdkclientStat

`func NewSdkclientStat(id string, lastSeen float32, networkConnection SdkclientStatNetworkConnection, uuid string, ) *SdkclientStat`

NewSdkclientStat instantiates a new SdkclientStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkclientStatWithDefaults

`func NewSdkclientStatWithDefaults() *SdkclientStat`

NewSdkclientStatWithDefaults instantiates a new SdkclientStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SdkclientStat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdkclientStat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdkclientStat) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeen

`func (o *SdkclientStat) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *SdkclientStat) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *SdkclientStat) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetMapId

`func (o *SdkclientStat) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *SdkclientStat) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *SdkclientStat) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *SdkclientStat) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### SetMapIdNil

`func (o *SdkclientStat) SetMapIdNil(b bool)`

 SetMapIdNil sets the value for MapId to be an explicit nil

### UnsetMapId
`func (o *SdkclientStat) UnsetMapId()`

UnsetMapId ensures that no value is present for MapId, not even an explicit nil
### GetName

`func (o *SdkclientStat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkclientStat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkclientStat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkclientStat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConnection

`func (o *SdkclientStat) GetNetworkConnection() SdkclientStatNetworkConnection`

GetNetworkConnection returns the NetworkConnection field if non-nil, zero value otherwise.

### GetNetworkConnectionOk

`func (o *SdkclientStat) GetNetworkConnectionOk() (*SdkclientStatNetworkConnection, bool)`

GetNetworkConnectionOk returns a tuple with the NetworkConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnection

`func (o *SdkclientStat) SetNetworkConnection(v SdkclientStatNetworkConnection)`

SetNetworkConnection sets NetworkConnection field to given value.


### GetUuid

`func (o *SdkclientStat) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SdkclientStat) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SdkclientStat) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetX

`func (o *SdkclientStat) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SdkclientStat) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SdkclientStat) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *SdkclientStat) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *SdkclientStat) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *SdkclientStat) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *SdkclientStat) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *SdkclientStat) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


