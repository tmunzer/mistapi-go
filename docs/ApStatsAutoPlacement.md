# ApStatsAutoPlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**ApStatsAutoPlacementInfo**](ApStatsAutoPlacementInfo.md) |  | [optional] 
**RecommendedAnchor** | Pointer to **bool** | Flag to represent if AP is recommended as an anchor by auto placement service | [optional] 
**Status** | Pointer to **string** | Basic Placement Status | [optional] 
**StatusDetail** | Pointer to **string** | Additional info about placement status | [optional] 
**UseAutoPlacement** | Pointer to **bool** | Flag to represent if auto_placement values are currently utilized | [optional] 
**X** | Pointer to **float32** | X Autoplaced Position in pixels | [optional] 
**XM** | Pointer to **float32** | X Autoplaced Position in meters | [optional] 
**Y** | Pointer to **float32** | Y Autoplaced Position in pixels | [optional] 
**YM** | Pointer to **float32** | X Autoplaced Position in meters | [optional] 

## Methods

### NewApStatsAutoPlacement

`func NewApStatsAutoPlacement() *ApStatsAutoPlacement`

NewApStatsAutoPlacement instantiates a new ApStatsAutoPlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsAutoPlacementWithDefaults

`func NewApStatsAutoPlacementWithDefaults() *ApStatsAutoPlacement`

NewApStatsAutoPlacementWithDefaults instantiates a new ApStatsAutoPlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApStatsAutoPlacement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApStatsAutoPlacement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApStatsAutoPlacement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApStatsAutoPlacement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *ApStatsAutoPlacement) GetInfo() ApStatsAutoPlacementInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ApStatsAutoPlacement) GetInfoOk() (*ApStatsAutoPlacementInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ApStatsAutoPlacement) SetInfo(v ApStatsAutoPlacementInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ApStatsAutoPlacement) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetRecommendedAnchor

`func (o *ApStatsAutoPlacement) GetRecommendedAnchor() bool`

GetRecommendedAnchor returns the RecommendedAnchor field if non-nil, zero value otherwise.

### GetRecommendedAnchorOk

`func (o *ApStatsAutoPlacement) GetRecommendedAnchorOk() (*bool, bool)`

GetRecommendedAnchorOk returns a tuple with the RecommendedAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedAnchor

`func (o *ApStatsAutoPlacement) SetRecommendedAnchor(v bool)`

SetRecommendedAnchor sets RecommendedAnchor field to given value.

### HasRecommendedAnchor

`func (o *ApStatsAutoPlacement) HasRecommendedAnchor() bool`

HasRecommendedAnchor returns a boolean if a field has been set.

### GetStatus

`func (o *ApStatsAutoPlacement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApStatsAutoPlacement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApStatsAutoPlacement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApStatsAutoPlacement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *ApStatsAutoPlacement) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *ApStatsAutoPlacement) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *ApStatsAutoPlacement) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *ApStatsAutoPlacement) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetUseAutoPlacement

`func (o *ApStatsAutoPlacement) GetUseAutoPlacement() bool`

GetUseAutoPlacement returns the UseAutoPlacement field if non-nil, zero value otherwise.

### GetUseAutoPlacementOk

`func (o *ApStatsAutoPlacement) GetUseAutoPlacementOk() (*bool, bool)`

GetUseAutoPlacementOk returns a tuple with the UseAutoPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAutoPlacement

`func (o *ApStatsAutoPlacement) SetUseAutoPlacement(v bool)`

SetUseAutoPlacement sets UseAutoPlacement field to given value.

### HasUseAutoPlacement

`func (o *ApStatsAutoPlacement) HasUseAutoPlacement() bool`

HasUseAutoPlacement returns a boolean if a field has been set.

### GetX

`func (o *ApStatsAutoPlacement) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ApStatsAutoPlacement) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ApStatsAutoPlacement) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *ApStatsAutoPlacement) HasX() bool`

HasX returns a boolean if a field has been set.

### GetXM

`func (o *ApStatsAutoPlacement) GetXM() float32`

GetXM returns the XM field if non-nil, zero value otherwise.

### GetXMOk

`func (o *ApStatsAutoPlacement) GetXMOk() (*float32, bool)`

GetXMOk returns a tuple with the XM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXM

`func (o *ApStatsAutoPlacement) SetXM(v float32)`

SetXM sets XM field to given value.

### HasXM

`func (o *ApStatsAutoPlacement) HasXM() bool`

HasXM returns a boolean if a field has been set.

### GetY

`func (o *ApStatsAutoPlacement) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ApStatsAutoPlacement) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ApStatsAutoPlacement) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *ApStatsAutoPlacement) HasY() bool`

HasY returns a boolean if a field has been set.

### GetYM

`func (o *ApStatsAutoPlacement) GetYM() float32`

GetYM returns the YM field if non-nil, zero value otherwise.

### GetYMOk

`func (o *ApStatsAutoPlacement) GetYMOk() (*float32, bool)`

GetYMOk returns a tuple with the YM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYM

`func (o *ApStatsAutoPlacement) SetYM(v float32)`

SetYM sets YM field to given value.

### HasYM

`func (o *ApStatsAutoPlacement) HasYM() bool`

HasYM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


