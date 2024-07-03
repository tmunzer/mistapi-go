# UpgradeOrgDeviceUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**DeviceUpgradeStatus**](DeviceUpgradeStatus.md) |  | [optional] 
**Targets** | Pointer to [**UpgradeOrgDeviceTargets**](UpgradeOrgDeviceTargets.md) |  | [optional] 

## Methods

### NewUpgradeOrgDeviceUpgrade

`func NewUpgradeOrgDeviceUpgrade() *UpgradeOrgDeviceUpgrade`

NewUpgradeOrgDeviceUpgrade instantiates a new UpgradeOrgDeviceUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeOrgDeviceUpgradeWithDefaults

`func NewUpgradeOrgDeviceUpgradeWithDefaults() *UpgradeOrgDeviceUpgrade`

NewUpgradeOrgDeviceUpgradeWithDefaults instantiates a new UpgradeOrgDeviceUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpgradeOrgDeviceUpgrade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpgradeOrgDeviceUpgrade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpgradeOrgDeviceUpgrade) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpgradeOrgDeviceUpgrade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *UpgradeOrgDeviceUpgrade) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpgradeOrgDeviceUpgrade) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpgradeOrgDeviceUpgrade) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpgradeOrgDeviceUpgrade) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *UpgradeOrgDeviceUpgrade) GetStatus() DeviceUpgradeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradeOrgDeviceUpgrade) GetStatusOk() (*DeviceUpgradeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradeOrgDeviceUpgrade) SetStatus(v DeviceUpgradeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpgradeOrgDeviceUpgrade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargets

`func (o *UpgradeOrgDeviceUpgrade) GetTargets() UpgradeOrgDeviceTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *UpgradeOrgDeviceUpgrade) GetTargetsOk() (*UpgradeOrgDeviceTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *UpgradeOrgDeviceUpgrade) SetTargets(v UpgradeOrgDeviceTargets)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *UpgradeOrgDeviceUpgrade) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


