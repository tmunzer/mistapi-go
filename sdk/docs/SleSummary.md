# SleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifiers** | [**[]SleClassifier**](SleClassifier.md) |  | 
**End** | **float32** |  | 
**Events** | **[]map[string]interface{}** |  | 
**Impact** | [**SleSummaryImpact**](SleSummaryImpact.md) |  | 
**Sle** | [**SleSummarySle**](SleSummarySle.md) |  | 
**Start** | **float32** |  | 

## Methods

### NewSleSummary

`func NewSleSummary(classifiers []SleClassifier, end float32, events []map[string]interface{}, impact SleSummaryImpact, sle SleSummarySle, start float32, ) *SleSummary`

NewSleSummary instantiates a new SleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleSummaryWithDefaults

`func NewSleSummaryWithDefaults() *SleSummary`

NewSleSummaryWithDefaults instantiates a new SleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifiers

`func (o *SleSummary) GetClassifiers() []SleClassifier`

GetClassifiers returns the Classifiers field if non-nil, zero value otherwise.

### GetClassifiersOk

`func (o *SleSummary) GetClassifiersOk() (*[]SleClassifier, bool)`

GetClassifiersOk returns a tuple with the Classifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifiers

`func (o *SleSummary) SetClassifiers(v []SleClassifier)`

SetClassifiers sets Classifiers field to given value.


### GetEnd

`func (o *SleSummary) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleSummary) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleSummary) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetEvents

`func (o *SleSummary) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SleSummary) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SleSummary) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.


### GetImpact

`func (o *SleSummary) GetImpact() SleSummaryImpact`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *SleSummary) GetImpactOk() (*SleSummaryImpact, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *SleSummary) SetImpact(v SleSummaryImpact)`

SetImpact sets Impact field to given value.


### GetSle

`func (o *SleSummary) GetSle() SleSummarySle`

GetSle returns the Sle field if non-nil, zero value otherwise.

### GetSleOk

`func (o *SleSummary) GetSleOk() (*SleSummarySle, bool)`

GetSleOk returns a tuple with the Sle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSle

`func (o *SleSummary) SetSle(v SleSummarySle)`

SetSle sets Sle field to given value.


### GetStart

`func (o *SleSummary) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleSummary) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleSummary) SetStart(v float32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


