# ResponseDiscoveredSwitches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **float32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]DiscoveredSwitch**](DiscoveredSwitch.md) |  | 
**Start** | **float32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseDiscoveredSwitches

`func NewResponseDiscoveredSwitches(end float32, limit int32, results []DiscoveredSwitch, start float32, total int32, ) *ResponseDiscoveredSwitches`

NewResponseDiscoveredSwitches instantiates a new ResponseDiscoveredSwitches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDiscoveredSwitchesWithDefaults

`func NewResponseDiscoveredSwitchesWithDefaults() *ResponseDiscoveredSwitches`

NewResponseDiscoveredSwitchesWithDefaults instantiates a new ResponseDiscoveredSwitches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseDiscoveredSwitches) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseDiscoveredSwitches) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseDiscoveredSwitches) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseDiscoveredSwitches) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseDiscoveredSwitches) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseDiscoveredSwitches) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseDiscoveredSwitches) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseDiscoveredSwitches) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseDiscoveredSwitches) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseDiscoveredSwitches) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseDiscoveredSwitches) GetResults() []DiscoveredSwitch`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseDiscoveredSwitches) GetResultsOk() (*[]DiscoveredSwitch, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseDiscoveredSwitches) SetResults(v []DiscoveredSwitch)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseDiscoveredSwitches) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseDiscoveredSwitches) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseDiscoveredSwitches) SetStart(v float32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseDiscoveredSwitches) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseDiscoveredSwitches) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseDiscoveredSwitches) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


