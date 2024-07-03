# AccountOauthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Linked app(zoom/teams/intune) account id | [optional] 
**Company** | Pointer to **string** | Name of the company whose account mist has subscribed to | [optional] [readonly] 
**Error** | Pointer to **string** | This error is provided when the account fails to fetch token/data | [optional] [readonly] 
**Errors** | Pointer to **[]string** |  | [optional] [readonly] 
**LastStatus** | Pointer to **string** | Is the last data pull for account is successful or not | [optional] [readonly] 
**LastSync** | Pointer to **int32** | Last data pull timestamp, background jobs that pull account data | [optional] [readonly] 
**LinkedBy** | Pointer to **string** | First name of the user who linked the account | [optional] [readonly] 
**MaxDailyApiRequests** | Pointer to **int32** | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ | [optional] 

## Methods

### NewAccountOauthInfo

`func NewAccountOauthInfo() *AccountOauthInfo`

NewAccountOauthInfo instantiates a new AccountOauthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOauthInfoWithDefaults

`func NewAccountOauthInfoWithDefaults() *AccountOauthInfo`

NewAccountOauthInfoWithDefaults instantiates a new AccountOauthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountOauthInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountOauthInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountOauthInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountOauthInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCompany

`func (o *AccountOauthInfo) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AccountOauthInfo) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AccountOauthInfo) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AccountOauthInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetError

`func (o *AccountOauthInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountOauthInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountOauthInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AccountOauthInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *AccountOauthInfo) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AccountOauthInfo) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AccountOauthInfo) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AccountOauthInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLastStatus

`func (o *AccountOauthInfo) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *AccountOauthInfo) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *AccountOauthInfo) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *AccountOauthInfo) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastSync

`func (o *AccountOauthInfo) GetLastSync() int32`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AccountOauthInfo) GetLastSyncOk() (*int32, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AccountOauthInfo) SetLastSync(v int32)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AccountOauthInfo) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLinkedBy

`func (o *AccountOauthInfo) GetLinkedBy() string`

GetLinkedBy returns the LinkedBy field if non-nil, zero value otherwise.

### GetLinkedByOk

`func (o *AccountOauthInfo) GetLinkedByOk() (*string, bool)`

GetLinkedByOk returns a tuple with the LinkedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedBy

`func (o *AccountOauthInfo) SetLinkedBy(v string)`

SetLinkedBy sets LinkedBy field to given value.

### HasLinkedBy

`func (o *AccountOauthInfo) HasLinkedBy() bool`

HasLinkedBy returns a boolean if a field has been set.

### GetMaxDailyApiRequests

`func (o *AccountOauthInfo) GetMaxDailyApiRequests() int32`

GetMaxDailyApiRequests returns the MaxDailyApiRequests field if non-nil, zero value otherwise.

### GetMaxDailyApiRequestsOk

`func (o *AccountOauthInfo) GetMaxDailyApiRequestsOk() (*int32, bool)`

GetMaxDailyApiRequestsOk returns a tuple with the MaxDailyApiRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDailyApiRequests

`func (o *AccountOauthInfo) SetMaxDailyApiRequests(v int32)`

SetMaxDailyApiRequests sets MaxDailyApiRequests field to given value.

### HasMaxDailyApiRequests

`func (o *AccountOauthInfo) HasMaxDailyApiRequests() bool`

HasMaxDailyApiRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


