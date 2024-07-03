# ModelMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Flags** | Pointer to **map[string]int32** | name/val pair objects for location engine to use | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Height** | Pointer to **int32** | when type&#x3D;image, height of the image map | [optional] 
**HeightM** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LatlngBr** | Pointer to [**LatlngBr**](LatlngBr.md) |  | [optional] 
**LatlngTl** | Pointer to [**LatlngTl**](LatlngTl.md) |  | [optional] 
**Locked** | Pointer to **bool** | whether this map is considered locked down | [optional] [default to false]
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the map | [optional] 
**OccupancyLimit** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Orientation** | Pointer to **int32** | orientation of the map, 0 means up is north, 90 means up is west | [optional] [default to 0]
**OriginX** | Pointer to **int32** | the user-annotated x origin, pixels | [optional] 
**OriginY** | Pointer to **int32** | the user-annotated y origin, pixels | [optional] 
**Ppm** | Pointer to **float32** | when type&#x3D;image, pixels per meter | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SitesurveyPath** | Pointer to [**[]MapSitesurveyPathItems**](MapSitesurveyPathItems.md) | sitesurvey_path | [optional] 
**ThumbnailUrl** | Pointer to **string** | when type&#x3D;image, the url for the thumbnail image / preview | [optional] [readonly] 
**Type** | Pointer to [**MapType**](MapType.md) |  | [optional] [default to MAPTYPE_IMAGE]
**Url** | Pointer to **string** | when type&#x3D;image, the url | [optional] [readonly] 
**UseAutoOrientation** | Pointer to **bool** | whether this map uses autooreintation values or ignores them | [optional] [default to false]
**UseAutoPlacement** | Pointer to **bool** | whether this map uses autoplacement values or ignores them | [optional] [default to false]
**View** | Pointer to [**NullableMapView**](MapView.md) |  | [optional] 
**WallPath** | Pointer to [**MapWallPath**](MapWallPath.md) |  | [optional] 
**Wayfinding** | Pointer to [**MapWayfinding**](MapWayfinding.md) |  | [optional] 
**WayfindingPath** | Pointer to [**MapWayfindingPath**](MapWayfindingPath.md) |  | [optional] 
**Width** | Pointer to **int32** | when type&#x3D;image, width of the image map | [optional] 
**WidthM** | Pointer to **float32** |  | [optional] 

## Methods

### NewModelMap

`func NewModelMap() *ModelMap`

NewModelMap instantiates a new ModelMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMapWithDefaults

`func NewModelMapWithDefaults() *ModelMap`

NewModelMapWithDefaults instantiates a new ModelMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *ModelMap) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ModelMap) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ModelMap) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ModelMap) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetFlags

`func (o *ModelMap) GetFlags() map[string]int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ModelMap) GetFlagsOk() (*map[string]int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ModelMap) SetFlags(v map[string]int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ModelMap) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetForSite

`func (o *ModelMap) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *ModelMap) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *ModelMap) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *ModelMap) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeight

`func (o *ModelMap) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelMap) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelMap) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelMap) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetHeightM

`func (o *ModelMap) GetHeightM() float32`

GetHeightM returns the HeightM field if non-nil, zero value otherwise.

### GetHeightMOk

`func (o *ModelMap) GetHeightMOk() (*float32, bool)`

GetHeightMOk returns a tuple with the HeightM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightM

`func (o *ModelMap) SetHeightM(v float32)`

SetHeightM sets HeightM field to given value.

### HasHeightM

`func (o *ModelMap) HasHeightM() bool`

HasHeightM returns a boolean if a field has been set.

### GetId

`func (o *ModelMap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelMap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelMap) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelMap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatlngBr

`func (o *ModelMap) GetLatlngBr() LatlngBr`

GetLatlngBr returns the LatlngBr field if non-nil, zero value otherwise.

### GetLatlngBrOk

`func (o *ModelMap) GetLatlngBrOk() (*LatlngBr, bool)`

GetLatlngBrOk returns a tuple with the LatlngBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlngBr

`func (o *ModelMap) SetLatlngBr(v LatlngBr)`

SetLatlngBr sets LatlngBr field to given value.

### HasLatlngBr

`func (o *ModelMap) HasLatlngBr() bool`

HasLatlngBr returns a boolean if a field has been set.

### GetLatlngTl

`func (o *ModelMap) GetLatlngTl() LatlngTl`

GetLatlngTl returns the LatlngTl field if non-nil, zero value otherwise.

### GetLatlngTlOk

`func (o *ModelMap) GetLatlngTlOk() (*LatlngTl, bool)`

GetLatlngTlOk returns a tuple with the LatlngTl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlngTl

`func (o *ModelMap) SetLatlngTl(v LatlngTl)`

SetLatlngTl sets LatlngTl field to given value.

### HasLatlngTl

`func (o *ModelMap) HasLatlngTl() bool`

HasLatlngTl returns a boolean if a field has been set.

### GetLocked

`func (o *ModelMap) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ModelMap) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ModelMap) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ModelMap) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModifiedTime

`func (o *ModelMap) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *ModelMap) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *ModelMap) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *ModelMap) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *ModelMap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelMap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOccupancyLimit

`func (o *ModelMap) GetOccupancyLimit() int32`

GetOccupancyLimit returns the OccupancyLimit field if non-nil, zero value otherwise.

### GetOccupancyLimitOk

`func (o *ModelMap) GetOccupancyLimitOk() (*int32, bool)`

GetOccupancyLimitOk returns a tuple with the OccupancyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancyLimit

`func (o *ModelMap) SetOccupancyLimit(v int32)`

SetOccupancyLimit sets OccupancyLimit field to given value.

### HasOccupancyLimit

`func (o *ModelMap) HasOccupancyLimit() bool`

HasOccupancyLimit returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelMap) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelMap) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelMap) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelMap) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrientation

`func (o *ModelMap) GetOrientation() int32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ModelMap) GetOrientationOk() (*int32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ModelMap) SetOrientation(v int32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ModelMap) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOriginX

`func (o *ModelMap) GetOriginX() int32`

GetOriginX returns the OriginX field if non-nil, zero value otherwise.

### GetOriginXOk

`func (o *ModelMap) GetOriginXOk() (*int32, bool)`

GetOriginXOk returns a tuple with the OriginX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginX

`func (o *ModelMap) SetOriginX(v int32)`

SetOriginX sets OriginX field to given value.

### HasOriginX

`func (o *ModelMap) HasOriginX() bool`

HasOriginX returns a boolean if a field has been set.

### GetOriginY

`func (o *ModelMap) GetOriginY() int32`

GetOriginY returns the OriginY field if non-nil, zero value otherwise.

### GetOriginYOk

`func (o *ModelMap) GetOriginYOk() (*int32, bool)`

GetOriginYOk returns a tuple with the OriginY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginY

`func (o *ModelMap) SetOriginY(v int32)`

SetOriginY sets OriginY field to given value.

### HasOriginY

`func (o *ModelMap) HasOriginY() bool`

HasOriginY returns a boolean if a field has been set.

### GetPpm

`func (o *ModelMap) GetPpm() float32`

GetPpm returns the Ppm field if non-nil, zero value otherwise.

### GetPpmOk

`func (o *ModelMap) GetPpmOk() (*float32, bool)`

GetPpmOk returns a tuple with the Ppm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpm

`func (o *ModelMap) SetPpm(v float32)`

SetPpm sets Ppm field to given value.

### HasPpm

`func (o *ModelMap) HasPpm() bool`

HasPpm returns a boolean if a field has been set.

### GetSiteId

`func (o *ModelMap) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ModelMap) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ModelMap) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ModelMap) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSitesurveyPath

`func (o *ModelMap) GetSitesurveyPath() []MapSitesurveyPathItems`

GetSitesurveyPath returns the SitesurveyPath field if non-nil, zero value otherwise.

### GetSitesurveyPathOk

`func (o *ModelMap) GetSitesurveyPathOk() (*[]MapSitesurveyPathItems, bool)`

GetSitesurveyPathOk returns a tuple with the SitesurveyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesurveyPath

`func (o *ModelMap) SetSitesurveyPath(v []MapSitesurveyPathItems)`

SetSitesurveyPath sets SitesurveyPath field to given value.

### HasSitesurveyPath

`func (o *ModelMap) HasSitesurveyPath() bool`

HasSitesurveyPath returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *ModelMap) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ModelMap) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ModelMap) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *ModelMap) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetType

`func (o *ModelMap) GetType() MapType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelMap) GetTypeOk() (*MapType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelMap) SetType(v MapType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelMap) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ModelMap) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelMap) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelMap) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelMap) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUseAutoOrientation

`func (o *ModelMap) GetUseAutoOrientation() bool`

GetUseAutoOrientation returns the UseAutoOrientation field if non-nil, zero value otherwise.

### GetUseAutoOrientationOk

`func (o *ModelMap) GetUseAutoOrientationOk() (*bool, bool)`

GetUseAutoOrientationOk returns a tuple with the UseAutoOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAutoOrientation

`func (o *ModelMap) SetUseAutoOrientation(v bool)`

SetUseAutoOrientation sets UseAutoOrientation field to given value.

### HasUseAutoOrientation

`func (o *ModelMap) HasUseAutoOrientation() bool`

HasUseAutoOrientation returns a boolean if a field has been set.

### GetUseAutoPlacement

`func (o *ModelMap) GetUseAutoPlacement() bool`

GetUseAutoPlacement returns the UseAutoPlacement field if non-nil, zero value otherwise.

### GetUseAutoPlacementOk

`func (o *ModelMap) GetUseAutoPlacementOk() (*bool, bool)`

GetUseAutoPlacementOk returns a tuple with the UseAutoPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAutoPlacement

`func (o *ModelMap) SetUseAutoPlacement(v bool)`

SetUseAutoPlacement sets UseAutoPlacement field to given value.

### HasUseAutoPlacement

`func (o *ModelMap) HasUseAutoPlacement() bool`

HasUseAutoPlacement returns a boolean if a field has been set.

### GetView

`func (o *ModelMap) GetView() MapView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ModelMap) GetViewOk() (*MapView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ModelMap) SetView(v MapView)`

SetView sets View field to given value.

### HasView

`func (o *ModelMap) HasView() bool`

HasView returns a boolean if a field has been set.

### SetViewNil

`func (o *ModelMap) SetViewNil(b bool)`

 SetViewNil sets the value for View to be an explicit nil

### UnsetView
`func (o *ModelMap) UnsetView()`

UnsetView ensures that no value is present for View, not even an explicit nil
### GetWallPath

`func (o *ModelMap) GetWallPath() MapWallPath`

GetWallPath returns the WallPath field if non-nil, zero value otherwise.

### GetWallPathOk

`func (o *ModelMap) GetWallPathOk() (*MapWallPath, bool)`

GetWallPathOk returns a tuple with the WallPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallPath

`func (o *ModelMap) SetWallPath(v MapWallPath)`

SetWallPath sets WallPath field to given value.

### HasWallPath

`func (o *ModelMap) HasWallPath() bool`

HasWallPath returns a boolean if a field has been set.

### GetWayfinding

`func (o *ModelMap) GetWayfinding() MapWayfinding`

GetWayfinding returns the Wayfinding field if non-nil, zero value otherwise.

### GetWayfindingOk

`func (o *ModelMap) GetWayfindingOk() (*MapWayfinding, bool)`

GetWayfindingOk returns a tuple with the Wayfinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayfinding

`func (o *ModelMap) SetWayfinding(v MapWayfinding)`

SetWayfinding sets Wayfinding field to given value.

### HasWayfinding

`func (o *ModelMap) HasWayfinding() bool`

HasWayfinding returns a boolean if a field has been set.

### GetWayfindingPath

`func (o *ModelMap) GetWayfindingPath() MapWayfindingPath`

GetWayfindingPath returns the WayfindingPath field if non-nil, zero value otherwise.

### GetWayfindingPathOk

`func (o *ModelMap) GetWayfindingPathOk() (*MapWayfindingPath, bool)`

GetWayfindingPathOk returns a tuple with the WayfindingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayfindingPath

`func (o *ModelMap) SetWayfindingPath(v MapWayfindingPath)`

SetWayfindingPath sets WayfindingPath field to given value.

### HasWayfindingPath

`func (o *ModelMap) HasWayfindingPath() bool`

HasWayfindingPath returns a boolean if a field has been set.

### GetWidth

`func (o *ModelMap) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelMap) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelMap) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelMap) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetWidthM

`func (o *ModelMap) GetWidthM() float32`

GetWidthM returns the WidthM field if non-nil, zero value otherwise.

### GetWidthMOk

`func (o *ModelMap) GetWidthMOk() (*float32, bool)`

GetWidthMOk returns a tuple with the WidthM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthM

`func (o *ModelMap) SetWidthM(v float32)`

SetWidthM sets WidthM field to given value.

### HasWidthM

`func (o *ModelMap) HasWidthM() bool`

HasWidthM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


