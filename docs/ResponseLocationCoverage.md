# ResponseLocationCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeamsMeans** | **[][]float32** | list of [x, y, mean]s, x/y are in meters (UI would need to use map.ppm to calulate the pixel location from top-left). | 
**End** | **int32** |  | 
**Gridsize** | **float32** | the size of grid, in meter | 
**ResultDef** | **[]string** | list of names annotating the fields in results | 
**Results** | **[][]float32** | list of results, see result_def. | 
**Start** | **int32** |  | 

## Methods

### NewResponseLocationCoverage

`func NewResponseLocationCoverage(beamsMeans [][]float32, end int32, gridsize float32, resultDef []string, results [][]float32, start int32, ) *ResponseLocationCoverage`

NewResponseLocationCoverage instantiates a new ResponseLocationCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLocationCoverageWithDefaults

`func NewResponseLocationCoverageWithDefaults() *ResponseLocationCoverage`

NewResponseLocationCoverageWithDefaults instantiates a new ResponseLocationCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeamsMeans

`func (o *ResponseLocationCoverage) GetBeamsMeans() [][]float32`

GetBeamsMeans returns the BeamsMeans field if non-nil, zero value otherwise.

### GetBeamsMeansOk

`func (o *ResponseLocationCoverage) GetBeamsMeansOk() (*[][]float32, bool)`

GetBeamsMeansOk returns a tuple with the BeamsMeans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamsMeans

`func (o *ResponseLocationCoverage) SetBeamsMeans(v [][]float32)`

SetBeamsMeans sets BeamsMeans field to given value.


### GetEnd

`func (o *ResponseLocationCoverage) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseLocationCoverage) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseLocationCoverage) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetGridsize

`func (o *ResponseLocationCoverage) GetGridsize() float32`

GetGridsize returns the Gridsize field if non-nil, zero value otherwise.

### GetGridsizeOk

`func (o *ResponseLocationCoverage) GetGridsizeOk() (*float32, bool)`

GetGridsizeOk returns a tuple with the Gridsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridsize

`func (o *ResponseLocationCoverage) SetGridsize(v float32)`

SetGridsize sets Gridsize field to given value.


### GetResultDef

`func (o *ResponseLocationCoverage) GetResultDef() []string`

GetResultDef returns the ResultDef field if non-nil, zero value otherwise.

### GetResultDefOk

`func (o *ResponseLocationCoverage) GetResultDefOk() (*[]string, bool)`

GetResultDefOk returns a tuple with the ResultDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDef

`func (o *ResponseLocationCoverage) SetResultDef(v []string)`

SetResultDef sets ResultDef field to given value.


### GetResults

`func (o *ResponseLocationCoverage) GetResults() [][]float32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseLocationCoverage) GetResultsOk() (*[][]float32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseLocationCoverage) SetResults(v [][]float32)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseLocationCoverage) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseLocationCoverage) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseLocationCoverage) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


