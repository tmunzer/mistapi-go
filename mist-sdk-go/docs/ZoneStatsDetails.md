# ZoneStatsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to **[]string** | list of ble assets currently in the zone and when they entered | [optional] 
**ClientWaits** | [**ZoneStatsDetailsClientWaits**](ZoneStatsDetailsClientWaits.md) |  | 
**Clients** | Pointer to **[]string** | list of clients currently in the zone and when they entered | [optional] 
**Id** | **string** | id of the zone | 
**MapId** | **string** | map_id of the zone | 
**Name** | **string** | name of the zone | 
**NumClients** | **int32** |  | 
**NumSdkclients** | **int32** | sdkclient wait time right now | 
**Sdkclients** | Pointer to **[]string** | list of sdkclients currently in the zone and when they entered | [optional] 

## Methods

### NewZoneStatsDetails

`func NewZoneStatsDetails(clientWaits ZoneStatsDetailsClientWaits, id string, mapId string, name string, numClients int32, numSdkclients int32, ) *ZoneStatsDetails`

NewZoneStatsDetails instantiates a new ZoneStatsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneStatsDetailsWithDefaults

`func NewZoneStatsDetailsWithDefaults() *ZoneStatsDetails`

NewZoneStatsDetailsWithDefaults instantiates a new ZoneStatsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *ZoneStatsDetails) GetAssets() []string`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ZoneStatsDetails) GetAssetsOk() (*[]string, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ZoneStatsDetails) SetAssets(v []string)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ZoneStatsDetails) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetClientWaits

`func (o *ZoneStatsDetails) GetClientWaits() ZoneStatsDetailsClientWaits`

GetClientWaits returns the ClientWaits field if non-nil, zero value otherwise.

### GetClientWaitsOk

`func (o *ZoneStatsDetails) GetClientWaitsOk() (*ZoneStatsDetailsClientWaits, bool)`

GetClientWaitsOk returns a tuple with the ClientWaits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientWaits

`func (o *ZoneStatsDetails) SetClientWaits(v ZoneStatsDetailsClientWaits)`

SetClientWaits sets ClientWaits field to given value.


### GetClients

`func (o *ZoneStatsDetails) GetClients() []string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ZoneStatsDetails) GetClientsOk() (*[]string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ZoneStatsDetails) SetClients(v []string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ZoneStatsDetails) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetId

`func (o *ZoneStatsDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneStatsDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneStatsDetails) SetId(v string)`

SetId sets Id field to given value.


### GetMapId

`func (o *ZoneStatsDetails) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ZoneStatsDetails) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ZoneStatsDetails) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetName

`func (o *ZoneStatsDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneStatsDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneStatsDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNumClients

`func (o *ZoneStatsDetails) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *ZoneStatsDetails) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *ZoneStatsDetails) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.


### GetNumSdkclients

`func (o *ZoneStatsDetails) GetNumSdkclients() int32`

GetNumSdkclients returns the NumSdkclients field if non-nil, zero value otherwise.

### GetNumSdkclientsOk

`func (o *ZoneStatsDetails) GetNumSdkclientsOk() (*int32, bool)`

GetNumSdkclientsOk returns a tuple with the NumSdkclients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSdkclients

`func (o *ZoneStatsDetails) SetNumSdkclients(v int32)`

SetNumSdkclients sets NumSdkclients field to given value.


### GetSdkclients

`func (o *ZoneStatsDetails) GetSdkclients() []string`

GetSdkclients returns the Sdkclients field if non-nil, zero value otherwise.

### GetSdkclientsOk

`func (o *ZoneStatsDetails) GetSdkclientsOk() (*[]string, bool)`

GetSdkclientsOk returns a tuple with the Sdkclients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclients

`func (o *ZoneStatsDetails) SetSdkclients(v []string)`

SetSdkclients sets Sdkclients field to given value.

### HasSdkclients

`func (o *ZoneStatsDetails) HasSdkclients() bool`

HasSdkclients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


