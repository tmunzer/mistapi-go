# VpnPeerStatSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **float32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]VpnPeerStat**](VpnPeerStat.md) |  | 
**Start** | **float32** |  | 
**Total** | **int32** |  | 

## Methods

### NewVpnPeerStatSearch

`func NewVpnPeerStatSearch(end float32, limit int32, results []VpnPeerStat, start float32, total int32, ) *VpnPeerStatSearch`

NewVpnPeerStatSearch instantiates a new VpnPeerStatSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPeerStatSearchWithDefaults

`func NewVpnPeerStatSearchWithDefaults() *VpnPeerStatSearch`

NewVpnPeerStatSearchWithDefaults instantiates a new VpnPeerStatSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *VpnPeerStatSearch) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *VpnPeerStatSearch) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *VpnPeerStatSearch) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *VpnPeerStatSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VpnPeerStatSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VpnPeerStatSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *VpnPeerStatSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *VpnPeerStatSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *VpnPeerStatSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *VpnPeerStatSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *VpnPeerStatSearch) GetResults() []VpnPeerStat`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *VpnPeerStatSearch) GetResultsOk() (*[]VpnPeerStat, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *VpnPeerStatSearch) SetResults(v []VpnPeerStat)`

SetResults sets Results field to given value.


### GetStart

`func (o *VpnPeerStatSearch) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *VpnPeerStatSearch) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *VpnPeerStatSearch) SetStart(v float32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *VpnPeerStatSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VpnPeerStatSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VpnPeerStatSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


