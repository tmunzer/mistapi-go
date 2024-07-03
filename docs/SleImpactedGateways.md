# SleImpactedGateways

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Failure** | Pointer to **string** |  | [optional] 
**Gateways** | Pointer to [**[]SleImpactedGatewaysGateway**](SleImpactedGatewaysGateway.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSleImpactedGateways

`func NewSleImpactedGateways() *SleImpactedGateways`

NewSleImpactedGateways instantiates a new SleImpactedGateways object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleImpactedGatewaysWithDefaults

`func NewSleImpactedGatewaysWithDefaults() *SleImpactedGateways`

NewSleImpactedGatewaysWithDefaults instantiates a new SleImpactedGateways object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *SleImpactedGateways) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleImpactedGateways) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleImpactedGateways) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *SleImpactedGateways) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetEnd

`func (o *SleImpactedGateways) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleImpactedGateways) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleImpactedGateways) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SleImpactedGateways) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFailure

`func (o *SleImpactedGateways) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *SleImpactedGateways) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *SleImpactedGateways) SetFailure(v string)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *SleImpactedGateways) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetGateways

`func (o *SleImpactedGateways) GetGateways() []SleImpactedGatewaysGateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *SleImpactedGateways) GetGatewaysOk() (*[]SleImpactedGatewaysGateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *SleImpactedGateways) SetGateways(v []SleImpactedGatewaysGateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *SleImpactedGateways) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetLimit

`func (o *SleImpactedGateways) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SleImpactedGateways) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SleImpactedGateways) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SleImpactedGateways) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *SleImpactedGateways) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleImpactedGateways) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleImpactedGateways) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SleImpactedGateways) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPage

`func (o *SleImpactedGateways) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SleImpactedGateways) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SleImpactedGateways) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SleImpactedGateways) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStart

`func (o *SleImpactedGateways) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleImpactedGateways) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleImpactedGateways) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SleImpactedGateways) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotalCount

`func (o *SleImpactedGateways) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SleImpactedGateways) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SleImpactedGateways) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SleImpactedGateways) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


