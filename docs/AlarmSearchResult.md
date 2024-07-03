# AlarmSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Results** | [**[]Alarm**](Alarm.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewAlarmSearchResult

`func NewAlarmSearchResult(end int32, limit int32, results []Alarm, start int32, total int32, ) *AlarmSearchResult`

NewAlarmSearchResult instantiates a new AlarmSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmSearchResultWithDefaults

`func NewAlarmSearchResultWithDefaults() *AlarmSearchResult`

NewAlarmSearchResultWithDefaults instantiates a new AlarmSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *AlarmSearchResult) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AlarmSearchResult) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AlarmSearchResult) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *AlarmSearchResult) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AlarmSearchResult) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AlarmSearchResult) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *AlarmSearchResult) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AlarmSearchResult) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AlarmSearchResult) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *AlarmSearchResult) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPage

`func (o *AlarmSearchResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AlarmSearchResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AlarmSearchResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AlarmSearchResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *AlarmSearchResult) GetResults() []Alarm`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AlarmSearchResult) GetResultsOk() (*[]Alarm, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AlarmSearchResult) SetResults(v []Alarm)`

SetResults sets Results field to given value.


### GetStart

`func (o *AlarmSearchResult) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AlarmSearchResult) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AlarmSearchResult) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *AlarmSearchResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlarmSearchResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlarmSearchResult) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


