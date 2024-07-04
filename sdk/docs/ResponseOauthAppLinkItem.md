# ResponseOauthAppLinkItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | customer account client id | [optional] 
**Error** | Pointer to **string** | This error is provided when the account fails to fetch token/data | [optional] [readonly] 
**InstanceUrl** | Pointer to **string** | Linked VMware Instance URL | [optional] 
**LastStatus** | Pointer to **string** | Is the last data pull for VMware account is successful or not | [optional] 
**LastSync** | Pointer to **int32** | Last data pull timestamp, background jobs that pull VMware account data | [optional] 
**LinkedBy** | Pointer to **string** | First name of the user who linked the VMware account | [optional] 
**Name** | Pointer to **string** | Name of the company whose VMware account mist has subscribed to | [optional] 
**SmartgroupName** | Pointer to **string** | smart group membership for determining compliance status | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** | Name of the company whose account mist has subscribed to | [optional] [readonly] 
**Errors** | Pointer to **[]string** |  | [optional] [readonly] 
**MaxDailyApiRequests** | Pointer to **int32** | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ | [optional] 
**LinkedTimestamp** | Pointer to **int32** | This error is provided when the VMware account fails to fetch token/data | [optional] 

## Methods

### NewResponseOauthAppLinkItem

`func NewResponseOauthAppLinkItem() *ResponseOauthAppLinkItem`

NewResponseOauthAppLinkItem instantiates a new ResponseOauthAppLinkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOauthAppLinkItemWithDefaults

`func NewResponseOauthAppLinkItemWithDefaults() *ResponseOauthAppLinkItem`

NewResponseOauthAppLinkItemWithDefaults instantiates a new ResponseOauthAppLinkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ResponseOauthAppLinkItem) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ResponseOauthAppLinkItem) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ResponseOauthAppLinkItem) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ResponseOauthAppLinkItem) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetError

`func (o *ResponseOauthAppLinkItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseOauthAppLinkItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseOauthAppLinkItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseOauthAppLinkItem) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInstanceUrl

`func (o *ResponseOauthAppLinkItem) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *ResponseOauthAppLinkItem) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *ResponseOauthAppLinkItem) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.

### HasInstanceUrl

`func (o *ResponseOauthAppLinkItem) HasInstanceUrl() bool`

HasInstanceUrl returns a boolean if a field has been set.

### GetLastStatus

`func (o *ResponseOauthAppLinkItem) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ResponseOauthAppLinkItem) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ResponseOauthAppLinkItem) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ResponseOauthAppLinkItem) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastSync

`func (o *ResponseOauthAppLinkItem) GetLastSync() int32`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ResponseOauthAppLinkItem) GetLastSyncOk() (*int32, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ResponseOauthAppLinkItem) SetLastSync(v int32)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ResponseOauthAppLinkItem) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLinkedBy

`func (o *ResponseOauthAppLinkItem) GetLinkedBy() string`

GetLinkedBy returns the LinkedBy field if non-nil, zero value otherwise.

### GetLinkedByOk

`func (o *ResponseOauthAppLinkItem) GetLinkedByOk() (*string, bool)`

GetLinkedByOk returns a tuple with the LinkedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedBy

`func (o *ResponseOauthAppLinkItem) SetLinkedBy(v string)`

SetLinkedBy sets LinkedBy field to given value.

### HasLinkedBy

`func (o *ResponseOauthAppLinkItem) HasLinkedBy() bool`

HasLinkedBy returns a boolean if a field has been set.

### GetName

`func (o *ResponseOauthAppLinkItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseOauthAppLinkItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseOauthAppLinkItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseOauthAppLinkItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSmartgroupName

`func (o *ResponseOauthAppLinkItem) GetSmartgroupName() string`

GetSmartgroupName returns the SmartgroupName field if non-nil, zero value otherwise.

### GetSmartgroupNameOk

`func (o *ResponseOauthAppLinkItem) GetSmartgroupNameOk() (*string, bool)`

GetSmartgroupNameOk returns a tuple with the SmartgroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartgroupName

`func (o *ResponseOauthAppLinkItem) SetSmartgroupName(v string)`

SetSmartgroupName sets SmartgroupName field to given value.

### HasSmartgroupName

`func (o *ResponseOauthAppLinkItem) HasSmartgroupName() bool`

HasSmartgroupName returns a boolean if a field has been set.

### GetAccountId

`func (o *ResponseOauthAppLinkItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResponseOauthAppLinkItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResponseOauthAppLinkItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResponseOauthAppLinkItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCompany

`func (o *ResponseOauthAppLinkItem) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ResponseOauthAppLinkItem) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ResponseOauthAppLinkItem) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ResponseOauthAppLinkItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetErrors

`func (o *ResponseOauthAppLinkItem) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ResponseOauthAppLinkItem) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ResponseOauthAppLinkItem) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ResponseOauthAppLinkItem) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMaxDailyApiRequests

`func (o *ResponseOauthAppLinkItem) GetMaxDailyApiRequests() int32`

GetMaxDailyApiRequests returns the MaxDailyApiRequests field if non-nil, zero value otherwise.

### GetMaxDailyApiRequestsOk

`func (o *ResponseOauthAppLinkItem) GetMaxDailyApiRequestsOk() (*int32, bool)`

GetMaxDailyApiRequestsOk returns a tuple with the MaxDailyApiRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDailyApiRequests

`func (o *ResponseOauthAppLinkItem) SetMaxDailyApiRequests(v int32)`

SetMaxDailyApiRequests sets MaxDailyApiRequests field to given value.

### HasMaxDailyApiRequests

`func (o *ResponseOauthAppLinkItem) HasMaxDailyApiRequests() bool`

HasMaxDailyApiRequests returns a boolean if a field has been set.

### GetLinkedTimestamp

`func (o *ResponseOauthAppLinkItem) GetLinkedTimestamp() int32`

GetLinkedTimestamp returns the LinkedTimestamp field if non-nil, zero value otherwise.

### GetLinkedTimestampOk

`func (o *ResponseOauthAppLinkItem) GetLinkedTimestampOk() (*int32, bool)`

GetLinkedTimestampOk returns a tuple with the LinkedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTimestamp

`func (o *ResponseOauthAppLinkItem) SetLinkedTimestamp(v int32)`

SetLinkedTimestamp sets LinkedTimestamp field to given value.

### HasLinkedTimestamp

`func (o *ResponseOauthAppLinkItem) HasLinkedTimestamp() bool`

HasLinkedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


