# AccountJamfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | customer account client id | [optional] 
**Error** | Pointer to **string** | This error is provided when the Jamf account fails to fetch token/data | [optional] [readonly] 
**InstanceUrl** | Pointer to **string** | customer account Jamf instance URL | [optional] 
**LastStatus** | Pointer to **string** | Is the last data pull for Jamf account is successful or not | [optional] [readonly] 
**LastSync** | Pointer to **int32** | Last data pull timestamp, background jobs that pull Jamf account data | [optional] [readonly] 
**LinkedBy** | Pointer to **string** | First name of the user who linked the Jamf account | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the company whose Jamf account mist has subscribed to | [optional] [readonly] 
**SmartgroupName** | Pointer to **string** | smart group membership for determining compliance status | [optional] 

## Methods

### NewAccountJamfInfo

`func NewAccountJamfInfo() *AccountJamfInfo`

NewAccountJamfInfo instantiates a new AccountJamfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountJamfInfoWithDefaults

`func NewAccountJamfInfoWithDefaults() *AccountJamfInfo`

NewAccountJamfInfoWithDefaults instantiates a new AccountJamfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountJamfInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountJamfInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountJamfInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AccountJamfInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetError

`func (o *AccountJamfInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountJamfInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountJamfInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AccountJamfInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInstanceUrl

`func (o *AccountJamfInfo) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *AccountJamfInfo) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *AccountJamfInfo) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.

### HasInstanceUrl

`func (o *AccountJamfInfo) HasInstanceUrl() bool`

HasInstanceUrl returns a boolean if a field has been set.

### GetLastStatus

`func (o *AccountJamfInfo) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *AccountJamfInfo) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *AccountJamfInfo) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *AccountJamfInfo) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastSync

`func (o *AccountJamfInfo) GetLastSync() int32`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AccountJamfInfo) GetLastSyncOk() (*int32, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AccountJamfInfo) SetLastSync(v int32)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AccountJamfInfo) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLinkedBy

`func (o *AccountJamfInfo) GetLinkedBy() string`

GetLinkedBy returns the LinkedBy field if non-nil, zero value otherwise.

### GetLinkedByOk

`func (o *AccountJamfInfo) GetLinkedByOk() (*string, bool)`

GetLinkedByOk returns a tuple with the LinkedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedBy

`func (o *AccountJamfInfo) SetLinkedBy(v string)`

SetLinkedBy sets LinkedBy field to given value.

### HasLinkedBy

`func (o *AccountJamfInfo) HasLinkedBy() bool`

HasLinkedBy returns a boolean if a field has been set.

### GetName

`func (o *AccountJamfInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountJamfInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountJamfInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountJamfInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSmartgroupName

`func (o *AccountJamfInfo) GetSmartgroupName() string`

GetSmartgroupName returns the SmartgroupName field if non-nil, zero value otherwise.

### GetSmartgroupNameOk

`func (o *AccountJamfInfo) GetSmartgroupNameOk() (*string, bool)`

GetSmartgroupNameOk returns a tuple with the SmartgroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartgroupName

`func (o *AccountJamfInfo) SetSmartgroupName(v string)`

SetSmartgroupName sets SmartgroupName field to given value.

### HasSmartgroupName

`func (o *AccountJamfInfo) HasSmartgroupName() bool`

HasSmartgroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


