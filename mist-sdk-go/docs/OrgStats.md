# OrgStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmtemplateId** | **string** |  | 
**AllowMist** | **bool** |  | 
**CreatedTime** | **float32** |  | 
**Id** | **string** |  | 
**ModifiedTime** | **float32** |  | 
**MspId** | **string** |  | 
**Name** | **string** |  | 
**NumDevices** | **int32** |  | 
**NumDevicesConnected** | **int32** |  | 
**NumDevicesDisconnected** | **int32** |  | 
**NumInventory** | **int32** |  | 
**NumSites** | **int32** |  | 
**OrggroupIds** | **[]string** |  | 
**SessionExpiry** | **int32** |  | 
**Sle** | [**[]OrgSleStat**](OrgSleStat.md) |  | 

## Methods

### NewOrgStats

`func NewOrgStats(alarmtemplateId string, allowMist bool, createdTime float32, id string, modifiedTime float32, mspId string, name string, numDevices int32, numDevicesConnected int32, numDevicesDisconnected int32, numInventory int32, numSites int32, orggroupIds []string, sessionExpiry int32, sle []OrgSleStat, ) *OrgStats`

NewOrgStats instantiates a new OrgStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgStatsWithDefaults

`func NewOrgStatsWithDefaults() *OrgStats`

NewOrgStatsWithDefaults instantiates a new OrgStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmtemplateId

`func (o *OrgStats) GetAlarmtemplateId() string`

GetAlarmtemplateId returns the AlarmtemplateId field if non-nil, zero value otherwise.

### GetAlarmtemplateIdOk

`func (o *OrgStats) GetAlarmtemplateIdOk() (*string, bool)`

GetAlarmtemplateIdOk returns a tuple with the AlarmtemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmtemplateId

`func (o *OrgStats) SetAlarmtemplateId(v string)`

SetAlarmtemplateId sets AlarmtemplateId field to given value.


### GetAllowMist

`func (o *OrgStats) GetAllowMist() bool`

GetAllowMist returns the AllowMist field if non-nil, zero value otherwise.

### GetAllowMistOk

`func (o *OrgStats) GetAllowMistOk() (*bool, bool)`

GetAllowMistOk returns a tuple with the AllowMist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMist

`func (o *OrgStats) SetAllowMist(v bool)`

SetAllowMist sets AllowMist field to given value.


### GetCreatedTime

`func (o *OrgStats) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OrgStats) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OrgStats) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetId

`func (o *OrgStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgStats) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedTime

`func (o *OrgStats) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *OrgStats) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *OrgStats) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.


### GetMspId

`func (o *OrgStats) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *OrgStats) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *OrgStats) SetMspId(v string)`

SetMspId sets MspId field to given value.


### GetName

`func (o *OrgStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgStats) SetName(v string)`

SetName sets Name field to given value.


### GetNumDevices

`func (o *OrgStats) GetNumDevices() int32`

GetNumDevices returns the NumDevices field if non-nil, zero value otherwise.

### GetNumDevicesOk

`func (o *OrgStats) GetNumDevicesOk() (*int32, bool)`

GetNumDevicesOk returns a tuple with the NumDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevices

`func (o *OrgStats) SetNumDevices(v int32)`

SetNumDevices sets NumDevices field to given value.


### GetNumDevicesConnected

`func (o *OrgStats) GetNumDevicesConnected() int32`

GetNumDevicesConnected returns the NumDevicesConnected field if non-nil, zero value otherwise.

### GetNumDevicesConnectedOk

`func (o *OrgStats) GetNumDevicesConnectedOk() (*int32, bool)`

GetNumDevicesConnectedOk returns a tuple with the NumDevicesConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevicesConnected

`func (o *OrgStats) SetNumDevicesConnected(v int32)`

SetNumDevicesConnected sets NumDevicesConnected field to given value.


### GetNumDevicesDisconnected

`func (o *OrgStats) GetNumDevicesDisconnected() int32`

GetNumDevicesDisconnected returns the NumDevicesDisconnected field if non-nil, zero value otherwise.

### GetNumDevicesDisconnectedOk

`func (o *OrgStats) GetNumDevicesDisconnectedOk() (*int32, bool)`

GetNumDevicesDisconnectedOk returns a tuple with the NumDevicesDisconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevicesDisconnected

`func (o *OrgStats) SetNumDevicesDisconnected(v int32)`

SetNumDevicesDisconnected sets NumDevicesDisconnected field to given value.


### GetNumInventory

`func (o *OrgStats) GetNumInventory() int32`

GetNumInventory returns the NumInventory field if non-nil, zero value otherwise.

### GetNumInventoryOk

`func (o *OrgStats) GetNumInventoryOk() (*int32, bool)`

GetNumInventoryOk returns a tuple with the NumInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInventory

`func (o *OrgStats) SetNumInventory(v int32)`

SetNumInventory sets NumInventory field to given value.


### GetNumSites

`func (o *OrgStats) GetNumSites() int32`

GetNumSites returns the NumSites field if non-nil, zero value otherwise.

### GetNumSitesOk

`func (o *OrgStats) GetNumSitesOk() (*int32, bool)`

GetNumSitesOk returns a tuple with the NumSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSites

`func (o *OrgStats) SetNumSites(v int32)`

SetNumSites sets NumSites field to given value.


### GetOrggroupIds

`func (o *OrgStats) GetOrggroupIds() []string`

GetOrggroupIds returns the OrggroupIds field if non-nil, zero value otherwise.

### GetOrggroupIdsOk

`func (o *OrgStats) GetOrggroupIdsOk() (*[]string, bool)`

GetOrggroupIdsOk returns a tuple with the OrggroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrggroupIds

`func (o *OrgStats) SetOrggroupIds(v []string)`

SetOrggroupIds sets OrggroupIds field to given value.


### GetSessionExpiry

`func (o *OrgStats) GetSessionExpiry() int32`

GetSessionExpiry returns the SessionExpiry field if non-nil, zero value otherwise.

### GetSessionExpiryOk

`func (o *OrgStats) GetSessionExpiryOk() (*int32, bool)`

GetSessionExpiryOk returns a tuple with the SessionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiry

`func (o *OrgStats) SetSessionExpiry(v int32)`

SetSessionExpiry sets SessionExpiry field to given value.


### GetSle

`func (o *OrgStats) GetSle() []OrgSleStat`

GetSle returns the Sle field if non-nil, zero value otherwise.

### GetSleOk

`func (o *OrgStats) GetSleOk() (*[]OrgSleStat, bool)`

GetSleOk returns a tuple with the Sle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSle

`func (o *OrgStats) SetSle(v []OrgSleStat)`

SetSle sets Sle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


