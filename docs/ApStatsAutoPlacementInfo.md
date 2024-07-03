# ApStatsAutoPlacementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNumber** | Pointer to **int32** | All APs sharing a given cluster number can be placed relative to each other | [optional] 
**OrientationStats** | Pointer to **int32** | The orientation of an AP | [optional] 
**ProbabilitySurface** | Pointer to [**ApStatsAutoPlacementInfoProbabilitySurface**](ApStatsAutoPlacementInfoProbabilitySurface.md) |  | [optional] 

## Methods

### NewApStatsAutoPlacementInfo

`func NewApStatsAutoPlacementInfo() *ApStatsAutoPlacementInfo`

NewApStatsAutoPlacementInfo instantiates a new ApStatsAutoPlacementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsAutoPlacementInfoWithDefaults

`func NewApStatsAutoPlacementInfoWithDefaults() *ApStatsAutoPlacementInfo`

NewApStatsAutoPlacementInfoWithDefaults instantiates a new ApStatsAutoPlacementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterNumber

`func (o *ApStatsAutoPlacementInfo) GetClusterNumber() int32`

GetClusterNumber returns the ClusterNumber field if non-nil, zero value otherwise.

### GetClusterNumberOk

`func (o *ApStatsAutoPlacementInfo) GetClusterNumberOk() (*int32, bool)`

GetClusterNumberOk returns a tuple with the ClusterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNumber

`func (o *ApStatsAutoPlacementInfo) SetClusterNumber(v int32)`

SetClusterNumber sets ClusterNumber field to given value.

### HasClusterNumber

`func (o *ApStatsAutoPlacementInfo) HasClusterNumber() bool`

HasClusterNumber returns a boolean if a field has been set.

### GetOrientationStats

`func (o *ApStatsAutoPlacementInfo) GetOrientationStats() int32`

GetOrientationStats returns the OrientationStats field if non-nil, zero value otherwise.

### GetOrientationStatsOk

`func (o *ApStatsAutoPlacementInfo) GetOrientationStatsOk() (*int32, bool)`

GetOrientationStatsOk returns a tuple with the OrientationStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientationStats

`func (o *ApStatsAutoPlacementInfo) SetOrientationStats(v int32)`

SetOrientationStats sets OrientationStats field to given value.

### HasOrientationStats

`func (o *ApStatsAutoPlacementInfo) HasOrientationStats() bool`

HasOrientationStats returns a boolean if a field has been set.

### GetProbabilitySurface

`func (o *ApStatsAutoPlacementInfo) GetProbabilitySurface() ApStatsAutoPlacementInfoProbabilitySurface`

GetProbabilitySurface returns the ProbabilitySurface field if non-nil, zero value otherwise.

### GetProbabilitySurfaceOk

`func (o *ApStatsAutoPlacementInfo) GetProbabilitySurfaceOk() (*ApStatsAutoPlacementInfoProbabilitySurface, bool)`

GetProbabilitySurfaceOk returns a tuple with the ProbabilitySurface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilitySurface

`func (o *ApStatsAutoPlacementInfo) SetProbabilitySurface(v ApStatsAutoPlacementInfoProbabilitySurface)`

SetProbabilitySurface sets ProbabilitySurface field to given value.

### HasProbabilitySurface

`func (o *ApStatsAutoPlacementInfo) HasProbabilitySurface() bool`

HasProbabilitySurface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


