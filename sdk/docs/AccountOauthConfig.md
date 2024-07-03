# AccountOauthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Linked app(zoom/teams/intune) account id | 
**MaxDailyApiRequests** | Pointer to **int32** | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ | [optional] 

## Methods

### NewAccountOauthConfig

`func NewAccountOauthConfig(accountId string, ) *AccountOauthConfig`

NewAccountOauthConfig instantiates a new AccountOauthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOauthConfigWithDefaults

`func NewAccountOauthConfigWithDefaults() *AccountOauthConfig`

NewAccountOauthConfigWithDefaults instantiates a new AccountOauthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountOauthConfig) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountOauthConfig) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountOauthConfig) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetMaxDailyApiRequests

`func (o *AccountOauthConfig) GetMaxDailyApiRequests() int32`

GetMaxDailyApiRequests returns the MaxDailyApiRequests field if non-nil, zero value otherwise.

### GetMaxDailyApiRequestsOk

`func (o *AccountOauthConfig) GetMaxDailyApiRequestsOk() (*int32, bool)`

GetMaxDailyApiRequestsOk returns a tuple with the MaxDailyApiRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDailyApiRequests

`func (o *AccountOauthConfig) SetMaxDailyApiRequests(v int32)`

SetMaxDailyApiRequests sets MaxDailyApiRequests field to given value.

### HasMaxDailyApiRequests

`func (o *AccountOauthConfig) HasMaxDailyApiRequests() bool`

HasMaxDailyApiRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


