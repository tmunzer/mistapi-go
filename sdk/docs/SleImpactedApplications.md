# SleImpactedApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]SleImpactedApplicationsApp**](SleImpactedApplicationsApp.md) |  | [optional] 
**Classifier** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Failure** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSleImpactedApplications

`func NewSleImpactedApplications() *SleImpactedApplications`

NewSleImpactedApplications instantiates a new SleImpactedApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleImpactedApplicationsWithDefaults

`func NewSleImpactedApplicationsWithDefaults() *SleImpactedApplications`

NewSleImpactedApplicationsWithDefaults instantiates a new SleImpactedApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *SleImpactedApplications) GetApps() []SleImpactedApplicationsApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SleImpactedApplications) GetAppsOk() (*[]SleImpactedApplicationsApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SleImpactedApplications) SetApps(v []SleImpactedApplicationsApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SleImpactedApplications) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetClassifier

`func (o *SleImpactedApplications) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleImpactedApplications) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleImpactedApplications) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *SleImpactedApplications) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetEnd

`func (o *SleImpactedApplications) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleImpactedApplications) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleImpactedApplications) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SleImpactedApplications) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFailure

`func (o *SleImpactedApplications) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *SleImpactedApplications) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *SleImpactedApplications) SetFailure(v string)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *SleImpactedApplications) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetLimit

`func (o *SleImpactedApplications) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SleImpactedApplications) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SleImpactedApplications) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SleImpactedApplications) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *SleImpactedApplications) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleImpactedApplications) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleImpactedApplications) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SleImpactedApplications) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPage

`func (o *SleImpactedApplications) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SleImpactedApplications) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SleImpactedApplications) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SleImpactedApplications) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStart

`func (o *SleImpactedApplications) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleImpactedApplications) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleImpactedApplications) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SleImpactedApplications) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotalCount

`func (o *SleImpactedApplications) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SleImpactedApplications) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SleImpactedApplications) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SleImpactedApplications) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


