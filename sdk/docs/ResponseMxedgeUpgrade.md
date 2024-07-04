# ResponseMxedgeUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**Counts** | [**MxedgeUpgradeResponseCounts**](MxedgeUpgradeResponseCounts.md) |  | 
**Id** | **string** |  | 
**Status** | **string** |  | 
**Strategy** | **string** |  | 
**Versions** | **map[string]interface{}** |  | 

## Methods

### NewResponseMxedgeUpgrade

`func NewResponseMxedgeUpgrade(channel string, counts MxedgeUpgradeResponseCounts, id string, status string, strategy string, versions map[string]interface{}, ) *ResponseMxedgeUpgrade`

NewResponseMxedgeUpgrade instantiates a new ResponseMxedgeUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMxedgeUpgradeWithDefaults

`func NewResponseMxedgeUpgradeWithDefaults() *ResponseMxedgeUpgrade`

NewResponseMxedgeUpgradeWithDefaults instantiates a new ResponseMxedgeUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ResponseMxedgeUpgrade) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ResponseMxedgeUpgrade) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ResponseMxedgeUpgrade) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCounts

`func (o *ResponseMxedgeUpgrade) GetCounts() MxedgeUpgradeResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ResponseMxedgeUpgrade) GetCountsOk() (*MxedgeUpgradeResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ResponseMxedgeUpgrade) SetCounts(v MxedgeUpgradeResponseCounts)`

SetCounts sets Counts field to given value.


### GetId

`func (o *ResponseMxedgeUpgrade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseMxedgeUpgrade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseMxedgeUpgrade) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ResponseMxedgeUpgrade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseMxedgeUpgrade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseMxedgeUpgrade) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStrategy

`func (o *ResponseMxedgeUpgrade) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ResponseMxedgeUpgrade) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ResponseMxedgeUpgrade) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetVersions

`func (o *ResponseMxedgeUpgrade) GetVersions() map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ResponseMxedgeUpgrade) GetVersionsOk() (*map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ResponseMxedgeUpgrade) SetVersions(v map[string]interface{})`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


