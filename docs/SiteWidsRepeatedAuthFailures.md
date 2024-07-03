# SiteWidsRepeatedAuthFailures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | window where a trigger will be detected and action to be taken (in seconds) | [optional] 
**Threshold** | Pointer to **int32** | count of events to trigger | [optional] 

## Methods

### NewSiteWidsRepeatedAuthFailures

`func NewSiteWidsRepeatedAuthFailures() *SiteWidsRepeatedAuthFailures`

NewSiteWidsRepeatedAuthFailures instantiates a new SiteWidsRepeatedAuthFailures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWidsRepeatedAuthFailuresWithDefaults

`func NewSiteWidsRepeatedAuthFailuresWithDefaults() *SiteWidsRepeatedAuthFailures`

NewSiteWidsRepeatedAuthFailuresWithDefaults instantiates a new SiteWidsRepeatedAuthFailures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *SiteWidsRepeatedAuthFailures) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SiteWidsRepeatedAuthFailures) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SiteWidsRepeatedAuthFailures) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SiteWidsRepeatedAuthFailures) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetThreshold

`func (o *SiteWidsRepeatedAuthFailures) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SiteWidsRepeatedAuthFailures) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SiteWidsRepeatedAuthFailures) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SiteWidsRepeatedAuthFailures) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


