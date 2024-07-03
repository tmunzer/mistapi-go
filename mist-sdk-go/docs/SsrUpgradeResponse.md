# SsrUpgradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**Counts** | [**SsrUpgradeResponseCounts**](SsrUpgradeResponseCounts.md) |  | 
**DeviceType** | **string** |  | 
**Id** | **string** |  | 
**Status** | **string** |  | 
**Strategy** | **string** |  | 
**Versions** | **map[string]string** |  | 

## Methods

### NewSsrUpgradeResponse

`func NewSsrUpgradeResponse(channel string, counts SsrUpgradeResponseCounts, deviceType string, id string, status string, strategy string, versions map[string]string, ) *SsrUpgradeResponse`

NewSsrUpgradeResponse instantiates a new SsrUpgradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsrUpgradeResponseWithDefaults

`func NewSsrUpgradeResponseWithDefaults() *SsrUpgradeResponse`

NewSsrUpgradeResponseWithDefaults instantiates a new SsrUpgradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SsrUpgradeResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SsrUpgradeResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SsrUpgradeResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCounts

`func (o *SsrUpgradeResponse) GetCounts() SsrUpgradeResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *SsrUpgradeResponse) GetCountsOk() (*SsrUpgradeResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *SsrUpgradeResponse) SetCounts(v SsrUpgradeResponseCounts)`

SetCounts sets Counts field to given value.


### GetDeviceType

`func (o *SsrUpgradeResponse) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SsrUpgradeResponse) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SsrUpgradeResponse) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetId

`func (o *SsrUpgradeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsrUpgradeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsrUpgradeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *SsrUpgradeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SsrUpgradeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SsrUpgradeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStrategy

`func (o *SsrUpgradeResponse) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *SsrUpgradeResponse) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *SsrUpgradeResponse) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetVersions

`func (o *SsrUpgradeResponse) GetVersions() map[string]string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *SsrUpgradeResponse) GetVersionsOk() (*map[string]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *SsrUpgradeResponse) SetVersions(v map[string]string)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


