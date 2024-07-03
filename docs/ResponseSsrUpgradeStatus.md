# ResponseSsrUpgradeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**DeviceType** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Status** | **string** |  | 
**Targets** | [**ResponseSsrUpgradeStatusTargets**](ResponseSsrUpgradeStatusTargets.md) |  | 
**Versions** | **map[string]interface{}** |  | 

## Methods

### NewResponseSsrUpgradeStatus

`func NewResponseSsrUpgradeStatus(channel string, id string, status string, targets ResponseSsrUpgradeStatusTargets, versions map[string]interface{}, ) *ResponseSsrUpgradeStatus`

NewResponseSsrUpgradeStatus instantiates a new ResponseSsrUpgradeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSsrUpgradeStatusWithDefaults

`func NewResponseSsrUpgradeStatusWithDefaults() *ResponseSsrUpgradeStatus`

NewResponseSsrUpgradeStatusWithDefaults instantiates a new ResponseSsrUpgradeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ResponseSsrUpgradeStatus) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ResponseSsrUpgradeStatus) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ResponseSsrUpgradeStatus) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetDeviceType

`func (o *ResponseSsrUpgradeStatus) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ResponseSsrUpgradeStatus) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ResponseSsrUpgradeStatus) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ResponseSsrUpgradeStatus) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetId

`func (o *ResponseSsrUpgradeStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseSsrUpgradeStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseSsrUpgradeStatus) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ResponseSsrUpgradeStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseSsrUpgradeStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseSsrUpgradeStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargets

`func (o *ResponseSsrUpgradeStatus) GetTargets() ResponseSsrUpgradeStatusTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ResponseSsrUpgradeStatus) GetTargetsOk() (*ResponseSsrUpgradeStatusTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ResponseSsrUpgradeStatus) SetTargets(v ResponseSsrUpgradeStatusTargets)`

SetTargets sets Targets field to given value.


### GetVersions

`func (o *ResponseSsrUpgradeStatus) GetVersions() map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ResponseSsrUpgradeStatus) GetVersionsOk() (*map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ResponseSsrUpgradeStatus) SetVersions(v map[string]interface{})`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


