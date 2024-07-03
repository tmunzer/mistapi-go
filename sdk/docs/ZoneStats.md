# ZoneStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsWaits** | Pointer to [**ZoneStatsAssetsWaits**](ZoneStatsAssetsWaits.md) |  | [optional] 
**ClientsWaits** | Pointer to [**ZoneStatsClientsWaits**](ZoneStatsClientsWaits.md) |  | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] 
**Id** | **string** | id of the zone | 
**MapId** | **string** | map_id of the zone | 
**ModifiedTime** | Pointer to **int32** |  | [optional] 
**Name** | **string** | name of the zone | 
**NumAssets** | Pointer to **int32** | number of assets | [optional] 
**NumClients** | Pointer to **int32** | number of wifi clients (unconnected + connected) | [optional] 
**NumSdkclients** | Pointer to **int32** | number of sdk clients | [optional] 
**OccupancyLimit** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SdkclientsWaits** | Pointer to [**ZoneStatsSdkclientsWaits**](ZoneStatsSdkclientsWaits.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Vertices** | Pointer to [**[]ZoneVertex**](ZoneVertex.md) | vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area | [optional] 
**VerticesM** | Pointer to [**[]ZoneVertexM**](ZoneVertexM.md) |  | [optional] 

## Methods

### NewZoneStats

`func NewZoneStats(id string, mapId string, name string, ) *ZoneStats`

NewZoneStats instantiates a new ZoneStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneStatsWithDefaults

`func NewZoneStatsWithDefaults() *ZoneStats`

NewZoneStatsWithDefaults instantiates a new ZoneStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsWaits

`func (o *ZoneStats) GetAssetsWaits() ZoneStatsAssetsWaits`

GetAssetsWaits returns the AssetsWaits field if non-nil, zero value otherwise.

### GetAssetsWaitsOk

`func (o *ZoneStats) GetAssetsWaitsOk() (*ZoneStatsAssetsWaits, bool)`

GetAssetsWaitsOk returns a tuple with the AssetsWaits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsWaits

`func (o *ZoneStats) SetAssetsWaits(v ZoneStatsAssetsWaits)`

SetAssetsWaits sets AssetsWaits field to given value.

### HasAssetsWaits

`func (o *ZoneStats) HasAssetsWaits() bool`

HasAssetsWaits returns a boolean if a field has been set.

### GetClientsWaits

`func (o *ZoneStats) GetClientsWaits() ZoneStatsClientsWaits`

GetClientsWaits returns the ClientsWaits field if non-nil, zero value otherwise.

### GetClientsWaitsOk

`func (o *ZoneStats) GetClientsWaitsOk() (*ZoneStatsClientsWaits, bool)`

GetClientsWaitsOk returns a tuple with the ClientsWaits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsWaits

`func (o *ZoneStats) SetClientsWaits(v ZoneStatsClientsWaits)`

SetClientsWaits sets ClientsWaits field to given value.

### HasClientsWaits

`func (o *ZoneStats) HasClientsWaits() bool`

HasClientsWaits returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ZoneStats) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ZoneStats) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ZoneStats) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ZoneStats) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *ZoneStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneStats) SetId(v string)`

SetId sets Id field to given value.


### GetMapId

`func (o *ZoneStats) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ZoneStats) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ZoneStats) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetModifiedTime

`func (o *ZoneStats) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *ZoneStats) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *ZoneStats) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *ZoneStats) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *ZoneStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneStats) SetName(v string)`

SetName sets Name field to given value.


### GetNumAssets

`func (o *ZoneStats) GetNumAssets() int32`

GetNumAssets returns the NumAssets field if non-nil, zero value otherwise.

### GetNumAssetsOk

`func (o *ZoneStats) GetNumAssetsOk() (*int32, bool)`

GetNumAssetsOk returns a tuple with the NumAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssets

`func (o *ZoneStats) SetNumAssets(v int32)`

SetNumAssets sets NumAssets field to given value.

### HasNumAssets

`func (o *ZoneStats) HasNumAssets() bool`

HasNumAssets returns a boolean if a field has been set.

### GetNumClients

`func (o *ZoneStats) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *ZoneStats) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *ZoneStats) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *ZoneStats) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetNumSdkclients

`func (o *ZoneStats) GetNumSdkclients() int32`

GetNumSdkclients returns the NumSdkclients field if non-nil, zero value otherwise.

### GetNumSdkclientsOk

`func (o *ZoneStats) GetNumSdkclientsOk() (*int32, bool)`

GetNumSdkclientsOk returns a tuple with the NumSdkclients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSdkclients

`func (o *ZoneStats) SetNumSdkclients(v int32)`

SetNumSdkclients sets NumSdkclients field to given value.

### HasNumSdkclients

`func (o *ZoneStats) HasNumSdkclients() bool`

HasNumSdkclients returns a boolean if a field has been set.

### GetOccupancyLimit

`func (o *ZoneStats) GetOccupancyLimit() int32`

GetOccupancyLimit returns the OccupancyLimit field if non-nil, zero value otherwise.

### GetOccupancyLimitOk

`func (o *ZoneStats) GetOccupancyLimitOk() (*int32, bool)`

GetOccupancyLimitOk returns a tuple with the OccupancyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancyLimit

`func (o *ZoneStats) SetOccupancyLimit(v int32)`

SetOccupancyLimit sets OccupancyLimit field to given value.

### HasOccupancyLimit

`func (o *ZoneStats) HasOccupancyLimit() bool`

HasOccupancyLimit returns a boolean if a field has been set.

### GetOrgId

`func (o *ZoneStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ZoneStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ZoneStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ZoneStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSdkclientsWaits

`func (o *ZoneStats) GetSdkclientsWaits() ZoneStatsSdkclientsWaits`

GetSdkclientsWaits returns the SdkclientsWaits field if non-nil, zero value otherwise.

### GetSdkclientsWaitsOk

`func (o *ZoneStats) GetSdkclientsWaitsOk() (*ZoneStatsSdkclientsWaits, bool)`

GetSdkclientsWaitsOk returns a tuple with the SdkclientsWaits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclientsWaits

`func (o *ZoneStats) SetSdkclientsWaits(v ZoneStatsSdkclientsWaits)`

SetSdkclientsWaits sets SdkclientsWaits field to given value.

### HasSdkclientsWaits

`func (o *ZoneStats) HasSdkclientsWaits() bool`

HasSdkclientsWaits returns a boolean if a field has been set.

### GetSiteId

`func (o *ZoneStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ZoneStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ZoneStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ZoneStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVertices

`func (o *ZoneStats) GetVertices() []ZoneVertex`

GetVertices returns the Vertices field if non-nil, zero value otherwise.

### GetVerticesOk

`func (o *ZoneStats) GetVerticesOk() (*[]ZoneVertex, bool)`

GetVerticesOk returns a tuple with the Vertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertices

`func (o *ZoneStats) SetVertices(v []ZoneVertex)`

SetVertices sets Vertices field to given value.

### HasVertices

`func (o *ZoneStats) HasVertices() bool`

HasVertices returns a boolean if a field has been set.

### GetVerticesM

`func (o *ZoneStats) GetVerticesM() []ZoneVertexM`

GetVerticesM returns the VerticesM field if non-nil, zero value otherwise.

### GetVerticesMOk

`func (o *ZoneStats) GetVerticesMOk() (*[]ZoneVertexM, bool)`

GetVerticesMOk returns a tuple with the VerticesM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticesM

`func (o *ZoneStats) SetVerticesM(v []ZoneVertexM)`

SetVerticesM sets VerticesM field to given value.

### HasVerticesM

`func (o *ZoneStats) HasVerticesM() bool`

HasVerticesM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


