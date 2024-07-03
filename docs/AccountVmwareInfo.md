# AccountVmwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**InstanceUrl** | Pointer to **string** | Linked VMware Instance URL | [optional] 
**LastStatus** | Pointer to **string** | Is the last data pull for VMware account is successful or not | [optional] 
**LastSync** | Pointer to **int32** | Last data pull timestamp, background jobs that pull VMware account data | [optional] 
**LinkedBy** | Pointer to **string** | First name of the user who linked the VMware account | [optional] 
**LinkedTimestamp** | Pointer to **int32** | This error is provided when the VMware account fails to fetch token/data | [optional] 
**Name** | Pointer to **string** | Name of the company whose VMware account mist has subscribed to | [optional] 

## Methods

### NewAccountVmwareInfo

`func NewAccountVmwareInfo() *AccountVmwareInfo`

NewAccountVmwareInfo instantiates a new AccountVmwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountVmwareInfoWithDefaults

`func NewAccountVmwareInfoWithDefaults() *AccountVmwareInfo`

NewAccountVmwareInfoWithDefaults instantiates a new AccountVmwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountVmwareInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountVmwareInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountVmwareInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountVmwareInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstanceUrl

`func (o *AccountVmwareInfo) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *AccountVmwareInfo) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *AccountVmwareInfo) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.

### HasInstanceUrl

`func (o *AccountVmwareInfo) HasInstanceUrl() bool`

HasInstanceUrl returns a boolean if a field has been set.

### GetLastStatus

`func (o *AccountVmwareInfo) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *AccountVmwareInfo) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *AccountVmwareInfo) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *AccountVmwareInfo) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastSync

`func (o *AccountVmwareInfo) GetLastSync() int32`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AccountVmwareInfo) GetLastSyncOk() (*int32, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AccountVmwareInfo) SetLastSync(v int32)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AccountVmwareInfo) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLinkedBy

`func (o *AccountVmwareInfo) GetLinkedBy() string`

GetLinkedBy returns the LinkedBy field if non-nil, zero value otherwise.

### GetLinkedByOk

`func (o *AccountVmwareInfo) GetLinkedByOk() (*string, bool)`

GetLinkedByOk returns a tuple with the LinkedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedBy

`func (o *AccountVmwareInfo) SetLinkedBy(v string)`

SetLinkedBy sets LinkedBy field to given value.

### HasLinkedBy

`func (o *AccountVmwareInfo) HasLinkedBy() bool`

HasLinkedBy returns a boolean if a field has been set.

### GetLinkedTimestamp

`func (o *AccountVmwareInfo) GetLinkedTimestamp() int32`

GetLinkedTimestamp returns the LinkedTimestamp field if non-nil, zero value otherwise.

### GetLinkedTimestampOk

`func (o *AccountVmwareInfo) GetLinkedTimestampOk() (*int32, bool)`

GetLinkedTimestampOk returns a tuple with the LinkedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTimestamp

`func (o *AccountVmwareInfo) SetLinkedTimestamp(v int32)`

SetLinkedTimestamp sets LinkedTimestamp field to given value.

### HasLinkedTimestamp

`func (o *AccountVmwareInfo) HasLinkedTimestamp() bool`

HasLinkedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *AccountVmwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountVmwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountVmwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountVmwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


